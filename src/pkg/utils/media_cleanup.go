package utils

import (
	"os"
	"path/filepath"
	"time"

	"github.com/aldinokemal/go-whatsapp-web-multidevice/config"
	"github.com/sirupsen/logrus"
)

// CleanupOldMediaFiles removes media files older than the specified number of days
func CleanupOldMediaFiles(retentionDays int) error {
	if retentionDays <= 0 {
		logrus.Debug("Media cleanup disabled (retention days = 0)")
		return nil
	}

	mediaPath := config.PathMedia
	if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
		logrus.Warnf("Media path does not exist: %s", mediaPath)
		return nil
	}

	cutoffTime := time.Now().AddDate(0, 0, -retentionDays)
	logrus.Infof("Starting media cleanup: removing files older than %d days (before %s)", 
		retentionDays, cutoffTime.Format("2006-01-02 15:04:05"))

	var totalFiles, deletedFiles, failedFiles int
	var totalSize, deletedSize int64

	err := filepath.Walk(mediaPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Warnf("Error accessing path %s: %v", path, err)
			return nil // Continue walking
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		totalFiles++
		totalSize += info.Size()

		// Check if file is older than cutoff time
		if info.ModTime().Before(cutoffTime) {
			size := info.Size()
			if err := os.Remove(path); err != nil {
				logrus.Warnf("Failed to delete file %s: %v", path, err)
				failedFiles++
			} else {
				deletedFiles++
				deletedSize += size
				logrus.Debugf("Deleted old media file: %s (age: %s, size: %d bytes)", 
					path, time.Since(info.ModTime()).Round(time.Hour), size)
			}
		}

		return nil
	})

	if err != nil {
		logrus.Errorf("Error during media cleanup: %v", err)
		return err
	}

	logrus.Infof("Media cleanup completed: scanned %d files (%.2f MB), deleted %d files (%.2f MB), failed %d",
		totalFiles, float64(totalSize)/(1024*1024),
		deletedFiles, float64(deletedSize)/(1024*1024),
		failedFiles)

	// Clean up empty directories
	cleanupEmptyDirectories(mediaPath)

	return nil
}

// cleanupEmptyDirectories removes empty subdirectories in the media path
func cleanupEmptyDirectories(rootPath string) {
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() || path == rootPath {
			return nil
		}

		// Try to remove the directory (will only succeed if empty)
		if err := os.Remove(path); err == nil {
			logrus.Debugf("Removed empty directory: %s", path)
		}

		return nil
	})
}

// StartMediaCleanupScheduler starts a background goroutine that periodically cleans up old media files
func StartMediaCleanupScheduler(retentionDays, intervalHours int) {
	if retentionDays <= 0 {
		logrus.Info("Media cleanup scheduler disabled (retention days = 0)")
		return
	}

	if intervalHours <= 0 {
		intervalHours = 24 // Default to 24 hours
	}

	logrus.Infof("Starting media cleanup scheduler: retention=%d days, interval=%d hours", 
		retentionDays, intervalHours)

	go func() {
		// Run immediately on startup
		time.Sleep(30 * time.Second) // Wait 30 seconds after startup
		if err := CleanupOldMediaFiles(retentionDays); err != nil {
			logrus.Errorf("Initial media cleanup failed: %v", err)
		}

		// Then run periodically
		ticker := time.NewTicker(time.Duration(intervalHours) * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			if err := CleanupOldMediaFiles(retentionDays); err != nil {
				logrus.Errorf("Scheduled media cleanup failed: %v", err)
			}
		}
	}()
}
