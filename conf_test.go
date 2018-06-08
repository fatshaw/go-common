package go_common

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

type RemoteConfig struct {
	Eureka struct {
		Address string `yaml:address`
	}
	Server struct {
		Port int `yaml:"port"`
	}
}

func TestGetConfFromConfigserver(t *testing.T) {
	os.Setenv("GO_ENV", "dev")
	tests := []struct {
		name string
	}{
		{name: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := RemoteConfig{}
			GetAppConfig("simulator-go", &c)
			assert.Equal(t, 8080, c.Server.Port, "server port should be equal")
			assert.Equal(t, "http://eureka-tertiary.local/eureka,http://eureka-secondary.local/eureka/,http://eureka-primary.local/eureka/",
				c.Eureka.Address, "eureka address should be equal")

		})
	}
}
