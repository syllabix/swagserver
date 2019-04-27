package swagserver

import (
	"html/template"
	"net/http"
	"testing"
)

func Test_isStaticFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "swagger css",
			args: args{
				path: "apidocs/swagger-ui.css",
			},
			want: true,
		},
		{
			name: "invalid file",
			args: args{
				path: "apidocs/swagger-ui.mp3",
			},
			want: false,
		},
		{
			name: "favicon",
			args: args{
				path: "/favicon-32x32.png",
			},
			want: true,
		},
		{
			name: "swagger-ui-bundle.js",
			args: args{
				path: "/swagger-ui-bundle.js",
			},
			want: true,
		},
		{
			name: "swagger-ui-bundle.js",
			args: args{
				path: "/swagger-ui-standalone-preset.js",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStaticFile(tt.args.path); got != tt.want {
				t.Errorf("isStaticFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler_ServeHTTP(t *testing.T) {
	type fields struct {
		fileserver http.Handler
		tmpl       *template.Template
		specURL    string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &handler{
				fileserver: tt.fields.fileserver,
				tmpl:       tt.fields.tmpl,
				specURL:    tt.fields.specURL,
			}
			s.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}
