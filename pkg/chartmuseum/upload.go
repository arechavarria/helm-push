package chartmuseum

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

// UploadChartPackage uploads a chart package to ChartMuseum (POST /api/charts)
func (client *Client) UploadChartPackage(chartPackagePath string, force bool) (*http.Response, error) {
	u, err := url.Parse(client.opts.url)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(chartPackagePath)
	payload := bufio.NewReader(f)

	u.Path = path.Join(u.Path, filepath.Base(chartPackagePath))
	req, err := http.NewRequest("PUT", u.String(), payload)
	if err != nil {
		return nil, err
	}

	// Add ?force to request querystring to force an upload if chart version already exists
	if force {
		req.URL.RawQuery = "force"
	}

	if client.opts.accessToken != "" {
		if client.opts.authHeader != "" {
			req.Header.Set(client.opts.authHeader, client.opts.accessToken)
		} else {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.opts.accessToken))
		}
	} else if client.opts.username != "" && client.opts.password != "" {
		req.SetBasicAuth(client.opts.username, client.opts.password)
	}

	return client.Do(req)
}