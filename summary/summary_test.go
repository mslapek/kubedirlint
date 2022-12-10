package summary

import (
	"reflect"
	"testing"
)

func TestSummarize(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []Summary
		wantErr bool
	}{
		{
			name: "empty",
			data: []byte{},
			want: nil,
		},
		{
			name: "deployment without namespace",
			data: []byte(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3`),
			want: []Summary{{
				Name: "nginx-deployment",
				Kind: "Deployment",
			}},
		},
		{
			name: "deployment with namespace",
			data: []byte(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: accounting
  labels:
    app: nginx
spec:
  replicas: 3`),
			want: []Summary{{
				Name:      "nginx-deployment",
				Namespace: "accounting",
				Kind:      "Deployment",
			}},
		},
		{
			name: "namespace",
			data: []byte(`
apiVersion: v1
kind: Namespace
metadata:
  name: development
  labels:
    section: finance`),
			want: []Summary{{
				Name: "development",
				Kind: "Namespace",
			}},
		},
		{
			name: "two objects",
			data: []byte(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tax-calc
  namespace: accounting
spec:
  replicas: 3

---

apiVersion: apps/v1
kind: Service
metadata:
  name: web
  namespace: development
spec:
  type: NodePort
`),
			want: []Summary{
				{
					Name:      "tax-calc",
					Namespace: "accounting",
					Kind:      "Deployment",
				},
				{
					Name:      "web",
					Namespace: "development",
					Kind:      "Service",
				},
			},
		},
		{
			name: "buggy yaml",
			data: []byte(`
apiVersion apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Summarize(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Summarize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Summarize() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
