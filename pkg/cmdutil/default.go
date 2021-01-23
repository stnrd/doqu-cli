package cmdutil

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/stnrd/doqu-cli/internal/config"
	"github.com/stnrd/doqu-cli/pkg/iostreams"
)

func NewFactory(appVersion string) *Factory {
	io := iostreams.System()

	var cachedConfig config.Config
	var configError error
	configFunc := func() (config.Config, error) {
		if cachedConfig != nil || configError != nil {
			return cachedConfig, configError
		}
		cachedConfig, configError = config.ParseDefaultConfig()
		if errors.Is(configError, os.ErrNotExist) {
			cachedConfig = config.NewBlankConfig()
			configError = nil
		}
		return cachedConfig, configError
	}

	return &Factory{
		IOStreams: io,
		Config:    configFunc,
		HttpClient: func() (*http.Client, error) {
			cfg, err := configFunc()
			if err != nil {
				return nil, err
			}
			fmt.Println(cfg)

			// return NewHTTPClient(io, cfg, appVersion, true), nil
			return &http.Client{}, nil
		},
	}

}
