// Package summary reads K8s YAML files and gives back a summary.
// Provides routines operating on the summaries.
package summary

import (
	"bytes"
	"io"

	"gopkg.in/yaml.v3"
)

// Summary of a K8s object.
type Summary struct {
	Name, Namespace, Kind string
}

// Summarize K8s YAML file.
func Summarize(data []byte) ([]Summary, error) {
	var ss []Summary
	var err error

	var y struct {
		Kind     string `yaml:"kind"`
		Metadata struct {
			Name      string `yaml:"name"`
			Namespace string `yaml:"namespace"`
		} `yaml:"metadata"`
	}

	d := yaml.NewDecoder(bytes.NewReader(data))
	for err = d.Decode(&y); err == nil; err = d.Decode(&y) {
		ss = append(ss, Summary{
			Name:      y.Metadata.Name,
			Namespace: y.Metadata.Namespace,
			Kind:      y.Kind,
		})
	}

	if err != io.EOF {
		return nil, err
	}

	return ss, nil
}
