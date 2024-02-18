package version

import "fmt"

var (
	AppVersion string
	BuildTime  string
	CommitHash string
)

func PrintInfo() {
	fmt.Printf("App Version: %s\n", AppVersion)
	fmt.Printf("Build Time: %s\n", BuildTime)
	fmt.Printf("Commit Hash: %s\n", CommitHash)
}
