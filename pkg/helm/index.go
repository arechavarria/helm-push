package helm

import (
	"github.com/ghodss/yaml"
	"helm.sh/helm/v3/pkg/repo"
)

type (
	// Index represents the index file in a chart repository
	Index struct {
		*repo.IndexFile
		ServerInfo ServerInfo `json:"serverInfo"`
	}

	// IndexDownloader is a function to download the index
	IndexDownloader func() ([]byte, error)
)

// LoadIndex loads an index file
func LoadIndex(data []byte) (*Index, error) {
	i := &Index{}
	if err := yaml.Unmarshal(data, i); err != nil {
		return i, err
	}
	i.SortEntries()
	return i, nil
}
