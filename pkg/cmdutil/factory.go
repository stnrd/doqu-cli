package cmdutil

import (
	"net/http"

	"github.com/stnrd/doqu-cli/internal/config"
	"github.com/stnrd/doqu-cli/pkg/iostreams"
)

type Factory struct {
	IOStreams  *iostreams.IOStreams
	HttpClient func() (*http.Client, error)
	Config     func() (config.Config, error)
}
