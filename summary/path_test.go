package summary

import "testing"

func TestTemplate_SuggestPath(t *testing.T) {
	tests := []struct {
		name      string
		tr        *Template
		summaries []Summary
		want      string
		wantErr   string
	}{
		{
			name:      "no summaries",
			tr:        &Template{},
			summaries: []Summary{},
			wantErr:   "expected exactly one K8s object in YAML, got 0",
		},
		{
			name:      "many summaries",
			tr:        &Template{},
			summaries: []Summary{{}, {}},
			wantErr:   "expected exactly one K8s object in YAML, got 2",
		},
		{
			name: "namespace",
			tr:   &Template{},
			summaries: []Summary{{
				Name: "accounting",
				Kind: "Namespace",
			}},
			want: "namespace/accounting.yaml",
		},
		{
			name: "deployment",
			tr:   &Template{},
			summaries: []Summary{{
				Name:      "tax-compute",
				Namespace: "accounting",
				Kind:      "Deployment",
			}},
			want: "accounting/deployment/tax-compute.yaml",
		},
		{
			name: "deployment in default namespace",
			tr:   &Template{},
			summaries: []Summary{{
				Name: "tax-compute",
				Kind: "Deployment",
			}},
			want: "default/deployment/tax-compute.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.SuggestPath(tt.summaries)
			if (err == nil && tt.wantErr != "") ||
				(err != nil && err.Error() != tt.wantErr) {
				t.Errorf(
					"Template.SuggestPath() error = %q, wantErr %q",
					err, tt.wantErr,
				)
				return
			}
			if got != tt.want {
				t.Errorf("Template.SuggestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
