//go:generate statik -src=./tmp
package swagserver

import (
	"reflect"
	"testing"

	_ "github.com/syllabix/swagserver/statik"
)

func Test_setConfig(t *testing.T) {
	type args struct {
		configurations []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			"Handles single custom config",
			args{
				configurations: []Config{
					Config{
						URLPath:        "/v1/my-docs",
						ShouldRender:   false,
						SwaggerSpecURL: "/swagger.json",
					},
				},
			},
			Config{
				URLPath:        "/v1/my-docs",
				ShouldRender:   false,
				SwaggerSpecURL: "/swagger.json",
			},
		},
		{
			"Handles multiple configs provided to variadic argument",
			args{
				configurations: []Config{
					Config{
						URLPath:        "/my-path-to-docs",
						ShouldRender:   true,
						SwaggerSpecURL: "/awesome-docs.json",
					},
					Config{
						URLPath:        "/v1/my-docs",
						ShouldRender:   false,
						SwaggerSpecURL: "/swagger.json",
					},
				},
			},
			Config{
				URLPath:        "/my-path-to-docs",
				ShouldRender:   true,
				SwaggerSpecURL: "/awesome-docs.json",
			},
		},
		{
			"Falls back to default if no args are provided",
			args{
				configurations: []Config{},
			},
			Config{
				URLPath:        "/swagger",
				ShouldRender:   true,
				SwaggerSpecURL: "/swagger.json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setConfig(tt.args.configurations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
