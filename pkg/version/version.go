// Package version provides version information for the pentagi application.
package version

import (
	"fmt"
	"runtime"
)

var (
	// Version is the current version of pentagi.
	// This is set at build time via -ldflags.
	Version = "0.1.0"

	// GitCommit is the git commit hash at build time.
	// This is set at build time via -ldflags.
	GitCommit = "unknown"

	// BuildDate is the date when the binary was built.
	// This is set at build time via -ldflags.
	BuildDate = "unknown"

	// GoVersion is the version of Go used to build the binary.
	GoVersion = runtime.Version()

	// Platform is the OS/Arch combination of the build target.
	Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

// Info holds the version information for the application.
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

// Get returns the current version information.
func Get() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		GoVersion: GoVersion,
		Platform:  Platform,
	}
}

// String returns a human-readable version string.
func (i Info) String() string {
	return fmt.Sprintf(
		"pentagi version %s (commit: %s, built: %s, go: %s, platform: %s)",
		i.Version,
		i.GitCommit,
		i.BuildDate,
		i.GoVersion,
		i.Platform,
	)
}
