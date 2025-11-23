package utils

import "testing"

func TestDetermineMediaExtension(t *testing.T) {
	tests := []struct {
		name       string
		filename   string
		mimeType   string
		wantSuffix string
	}{
		{
			name:       "DocxFromFilename",
			filename:   "report.docx",
			mimeType:   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			wantSuffix: ".docx",
		},
		{
			name:       "XlsxFromMime",
			filename:   "",
			mimeType:   "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
			wantSuffix: ".xlsx",
		},
		{
			name:       "PptxFromMime",
			filename:   "",
			mimeType:   "application/vnd.openxmlformats-officedocument.presentationml.presentation",
			wantSuffix: ".pptx",
		},
		{
			name:       "ZipFallback",
			filename:   "",
			mimeType:   "application/zip",
			wantSuffix: ".zip",
		},
		{
			name:       "ExeFromFilename",
			filename:   "installer.exe",
			mimeType:   "application/octet-stream",
			wantSuffix: ".exe",
		},
		{
			name:       "OggAudio",
			filename:   "",
			mimeType:   "audio/ogg",
			wantSuffix: ".ogg",
		},
		{
			name:       "OggAudioWithCodec",
			filename:   "",
			mimeType:   "audio/ogg; codecs=opus",
			wantSuffix: ".ogg",
		},
		{
			name:       "Mp3Audio",
			filename:   "",
			mimeType:   "audio/mpeg",
			wantSuffix: ".mp3",
		},
		{
			name:       "M4aAudio",
			filename:   "",
			mimeType:   "audio/mp4",
			wantSuffix: ".m4a",
		},
		{
			name:       "WavAudio",
			filename:   "",
			mimeType:   "audio/wav",
			wantSuffix: ".wav",
		},
		{
			name:       "WebmAudio",
			filename:   "",
			mimeType:   "audio/webm; codecs=opus",
			wantSuffix: ".webm",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := determineMediaExtension(tt.filename, tt.mimeType)
			if got != tt.wantSuffix {
				t.Fatalf("determineMediaExtension() = %q, want %q", got, tt.wantSuffix)
			}
		})
	}
}
