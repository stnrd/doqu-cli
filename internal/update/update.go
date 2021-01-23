package update

import (
	"fmt"

	"github.com/stnrd/doqu-cli/api"
)

// ReleaseInfo stores information about a release
type ReleaseInfo struct {
	Version string `json:"tag_name"`
	URL     string `json:"html_url"`
}

// CheckForUpdate checks whether this software has had a newer release on GitHub
func CheckForUpdate(client *api.Client, stateFilePath, repo, currentVersion string) (*ReleaseInfo, error) {
	releaseInfo, err := getLatestReleaseInfo(client, repo)
	if err != nil {
		return nil, err
	}

	fmt.Println(releaseInfo)

	// if versionGreaterThan(releaseInfo.Version, currentVersion) {
	// 	return releaseInfo, nil
	// }

	return nil, nil
}

func getLatestReleaseInfo(client *api.Client, repo string) (*ReleaseInfo, error) {
	return &ReleaseInfo{}, nil
}
