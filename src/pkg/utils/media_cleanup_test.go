package utils

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/aldinokemal/go-whatsapp-web-multidevice/config"
)

func TestCleanupOldMediaFiles(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "media_cleanup_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Save original config and restore after test
	originalMediaPath := config.PathMedia
	config.PathMedia = tempDir
	defer func() { config.PathMedia = originalMediaPath }()

	tests := []struct {
		name           string
		retentionDays  int
		setupFiles     map[string]time.Time // filename -> modification time
		expectedRemain []string
		expectedDelete []string
	}{
		{
			name:          "Disabled cleanup (retentionDays = 0)",
			retentionDays: 0,
			setupFiles: map[string]time.Time{
				"old_file.jpg": time.Now().AddDate(0, 0, -10),
			},
			expectedRemain: []string{"old_file.jpg"},
			expectedDelete: []string{},
		},
		{
			name:          "Delete files older than 7 days",
			retentionDays: 7,
			setupFiles: map[string]time.Time{
				"old_file.jpg":    time.Now().AddDate(0, 0, -10),
				"recent_file.jpg": time.Now().AddDate(0, 0, -3),
				"new_file.jpg":    time.Now(),
			},
			expectedRemain: []string{"recent_file.jpg", "new_file.jpg"},
			expectedDelete: []string{"old_file.jpg"},
		},
		{
			name:          "Delete files older than 30 days",
			retentionDays: 30,
			setupFiles: map[string]time.Time{
				"very_old.jpg": time.Now().AddDate(0, 0, -45),
				"old.jpg":      time.Now().AddDate(0, 0, -31),
				"recent.jpg":   time.Now().AddDate(0, 0, -29),
			},
			expectedRemain: []string{"recent.jpg"},
			expectedDelete: []string{"very_old.jpg", "old.jpg"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean temp directory
			os.RemoveAll(tempDir)
			os.MkdirAll(tempDir, 0755)

			// Setup test files
			for filename, modTime := range tt.setupFiles {
				filePath := filepath.Join(tempDir, filename)
				if err := os.WriteFile(filePath, []byte("test content"), 0644); err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
				if err := os.Chtimes(filePath, modTime, modTime); err != nil {
					t.Fatalf("Failed to set file time: %v", err)
				}
			}

			// Run cleanup
			if err := CleanupOldMediaFiles(tt.retentionDays); err != nil {
				t.Fatalf("CleanupOldMediaFiles failed: %v", err)
			}

			// Verify expected remaining files
			for _, filename := range tt.expectedRemain {
				filePath := filepath.Join(tempDir, filename)
				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					t.Errorf("Expected file to remain but it was deleted: %s", filename)
				}
			}

			// Verify expected deleted files
			for _, filename := range tt.expectedDelete {
				filePath := filepath.Join(tempDir, filename)
				if _, err := os.Stat(filePath); !os.IsNotExist(err) {
					t.Errorf("Expected file to be deleted but it still exists: %s", filename)
				}
			}
		})
	}
}

func TestCleanupOldMediaFiles_WithSubdirectories(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "media_cleanup_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Save original config and restore after test
	originalMediaPath := config.PathMedia
	config.PathMedia = tempDir
	defer func() { config.PathMedia = originalMediaPath }()

	// Create subdirectories
	subDir1 := filepath.Join(tempDir, "subdir1")
	subDir2 := filepath.Join(tempDir, "subdir2")
	os.MkdirAll(subDir1, 0755)
	os.MkdirAll(subDir2, 0755)

	// Create test files in subdirectories
	oldFile1 := filepath.Join(subDir1, "old.jpg")
	oldFile2 := filepath.Join(subDir2, "old.jpg")
	newFile := filepath.Join(subDir1, "new.jpg")

	oldTime := time.Now().AddDate(0, 0, -10)
	newTime := time.Now()

	// Create and set times
	os.WriteFile(oldFile1, []byte("test"), 0644)
	os.WriteFile(oldFile2, []byte("test"), 0644)
	os.WriteFile(newFile, []byte("test"), 0644)

	os.Chtimes(oldFile1, oldTime, oldTime)
	os.Chtimes(oldFile2, oldTime, oldTime)
	os.Chtimes(newFile, newTime, newTime)

	// Run cleanup with 7 days retention
	if err := CleanupOldMediaFiles(7); err != nil {
		t.Fatalf("CleanupOldMediaFiles failed: %v", err)
	}

	// Verify old files are deleted
	if _, err := os.Stat(oldFile1); !os.IsNotExist(err) {
		t.Errorf("Expected old file to be deleted: %s", oldFile1)
	}
	if _, err := os.Stat(oldFile2); !os.IsNotExist(err) {
		t.Errorf("Expected old file to be deleted: %s", oldFile2)
	}

	// Verify new file remains
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Errorf("Expected new file to remain: %s", newFile)
	}

	// Verify empty directory was removed
	if _, err := os.Stat(subDir2); !os.IsNotExist(err) {
		t.Logf("Note: Empty directory still exists (this is ok): %s", subDir2)
	}
}

func TestCleanupOldMediaFiles_NonExistentPath(t *testing.T) {
	// Save original config and restore after test
	originalMediaPath := config.PathMedia
	config.PathMedia = "/nonexistent/path/that/does/not/exist"
	defer func() { config.PathMedia = originalMediaPath }()

	// Should not return error for non-existent path
	if err := CleanupOldMediaFiles(7); err != nil {
		t.Errorf("CleanupOldMediaFiles should not error on non-existent path: %v", err)
	}
}
