package config

import (
	"os"
	"reflect"
	"testing"

	provides "github.com/zx06/ddnss/providers"
)

func Test_getAllenv(t *testing.T) {
	tests := []struct {
		name string
		want [][2]string
	}{
		{
			name: "getAllenv",
			want: [][2]string{
				{"GO_ENV", "test"},
				{"GO_ENV_TEST", "test"},
				{"GO_ENV_TEST_TEST", "test"},
			},
		},
	}
	os.Clearenv()
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.want {
				_ = os.Setenv(ENV_PREFIX+v[0], v[1])
			}
			if got := getAllenv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllenv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCfg(t *testing.T) {
	tests := []struct {
		name string
		want map[string]*CfgItem
	}{
		{
			name: "getAllCfg",
			want: map[string]*CfgItem{
				"test": {
					Type:   provides.ProviderType_DYNU,
					Domain: "test.com",
					ApiKey: "testkey",
				},
			},
		},
	}
	os.Clearenv()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.want {
				_ = os.Setenv(ENV_PREFIX+k+"_"+TYPE_KEY, v.Type)
				_ = os.Setenv(ENV_PREFIX+k+"_"+DOMAIN_KEY, v.Domain)
				_ = os.Setenv(ENV_PREFIX+k+"_"+APIKEY_KEY, v.ApiKey)

			}
			if got := GetEnvCfg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCfg() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
