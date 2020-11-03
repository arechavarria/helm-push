package helm

import (
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
)

type (
	// Chart is a helm package that contains metadata
	Chart struct {
		V3 *chart.Chart
	}
)

// SetVersion overrides the chart version
func (c *Chart) SetVersion(version string) {
	c.V3.Metadata.Version = version
}

// GetChartByName returns a chart by "name", which can be
// either a directory or .tgz package
func GetChartByName(name string) (*Chart, error) {
	c := &Chart{}
	v3c, err := loader.Load(name)
	if err != nil {
		return nil, err
	}
	c.V3 = v3c
	return c, nil
}

// CreateChartPackage creates a new .tgz package in directory
func CreateChartPackage(c *Chart, outDir string) (string, error) {
	return chartutil.Save(c.V3, outDir)
}
