package version

import (
	"io"
	"os"
	"testing"
)

func TestPrintInfo(t *testing.T) {
	cases := []struct {
		name       string
		appVersion string
		buildTime  string
		commitHash string
		expected   string
	}{
		{
			name:       "Test with valid inputs",
			appVersion: "1.0",
			buildTime:  "2024-02-18",
			commitHash: "abcdef",
			expected:   "App Version: 1.0\nBuild Time: 2024-02-18\nCommit Hash: abcdef\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			AppVersion = c.appVersion
			BuildTime = c.buildTime
			CommitHash = c.commitHash

			o := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintInfo()

			w.Close()
			os.Stdout = o

			out, _ := io.ReadAll(r)

			if got := string(out); got != c.expected {
				t.Errorf("PrintInfo() = %v, want %v", got, c.expected)
			}
		})
	}
}
