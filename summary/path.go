package summary

import (
	"fmt"
	"strings"
)

// Template describes how to assign a path to a K8s YAML.
type Template struct{}

// SuggestPath for a K8s YAML file.
func (t *Template) SuggestPath(summaries []Summary) (string, error) {
	if len(summaries) != 1 {
		return "", fmt.Errorf(
			"expected exactly one K8s object in YAML, got %d",
			len(summaries),
		)
	}
	s := summaries[0]

	if s.Kind == "Namespace" {
		return fmt.Sprintf("namespace/%s.yaml", s.Name), nil
	}

	ns := s.Namespace
	if len(ns) == 0 {
		ns = "default"
	}

	return fmt.Sprintf(
		"%s/%s/%s.yaml",
		ns,
		strings.ToLower(s.Kind),
		s.Name,
	), nil
}
