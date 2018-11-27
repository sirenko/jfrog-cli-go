package utils

import (
	"net/url"
	"path"

	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/httpclient"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

type artifactStatsJsonResponse struct {
	URI                  string `json:"uri"`
	DownloadCount        int    `json:"downloadCount"`
	LastDownloaded       int64  `json:"lastDownloaded"`
	LastDownloadedBy     string `json:"lastDownloadedBy"`
	RemoteDownloadCount  int    `json:"remoteDownloadCount"`
	RemoteLastDownloaded int    `json:"remoteLastDownloaded"`
}

type TopDownloadsService struct {
	httpClient *httpclient.HttpClient
	ArtDetails auth.ArtifactoryDetails
}

func NewTopDownloadsService(client *httpclient.HttpClient) *TopDownloadsService {
	return &TopDownloadsService{httpClient: client}
}

func (s *TopDownloadsService) GetArtifactoryDetails() auth.ArtifactoryDetails {
	return s.ArtDetails
}

func (s *TopDownloadsService) SetArtifactoryDetails(rt auth.ArtifactoryDetails) {
	s.ArtDetails = rt
}

func (s *TopDownloadsService) GetJfrogHttpClient() *httpclient.HttpClient {
	return s.httpClient
}

func (s *TopDownloadsService) IsDryRun() bool {
	return false
}

// GetAftifactDownloads returns downloads counter for a given 'artifact' in the 'repo'
func (s *TopDownloadsService) GetAftifactDownloadCount(artifactPath string) (int, error) {
	u, err := url.Parse(s.ArtDetails.GetUrl())
	if err != nil {
		return 20, err
	}
	u.Path = path.Join(u.Path, "/artifactory/api/storage", artifactPath)
	u.RawQuery = "stats"

	log.Debug(u)

	return 0, nil
	//
	// resp, err := c.httpClient.Do(req)
	// if err != nil {
	// 	log.Fatalf(`client failed to query artifactory: %+v`, err)
	// 	return -1, err
	// }
	// defer resp.Body.Close()
	// if resp.StatusCode != http.StatusOK {
	// 	log.Fatalf(`artifactory returned unexpected response status: %d - %s`, resp.StatusCode, resp.Status)
	// 	return -1, err
	// }
	// statsResp := new(artifactStatsJsonResponse)
	// if err := json.NewDecoder(resp.Body).Decode(statsResp); err != nil {
	// 	log.Fatalf(`failed to decode json response from artifactory: %+v`, err)
	// 	return -1, err
	// }
	// // fmt.Println(statsResp)
	// return statsResp.DownloadCount, nil
}
