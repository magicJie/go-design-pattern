package visitor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompresssor_Visit(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr string
	}{
		{
			name: "pdf",
			path: "./xx.pdf",
		},
		{
			name: "ppt",
			path: "./xx.ppt",
		},
		{
			name:    "404",
			path:    "./xx.xx",
			wantErr: "not found file type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := NewResourceFile(tt.path)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			compressor := &Compressor{}
			f.Accept(compressor)
		})
	}
}

func TestCompresssor_Visit2(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr string
	}{
		{
			name: "pdf",
			path: "./xx.pdf",
		},
		{
			name: "ppt",
			path: "./xx.ppt",
		},
		{
			name:    "404",
			path:    "./xx.xx",
			wantErr: "not found file type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := NewResourceFile(tt.path)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			compressor := &Compressor{}
			compressor.Visit(f)
		})
	}
}
