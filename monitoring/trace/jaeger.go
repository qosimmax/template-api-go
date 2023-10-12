package trace

import (
	"fmt"
	"io"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"template-api-go/config"
)

// InitGlobalTracer creates the global tracer object
func InitGlobalTracer(config *config.Config) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		Reporter: &jaegercfg.ReporterConfig{
			LocalAgentHostPort: fmt.Sprintf("%s:%s", config.JaegerAgentHost, config.JaegerAgentPort),
		},
		Sampler: &jaegercfg.SamplerConfig{
			Type:  config.JaegerSamplerType,
			Param: config.JaegerSamplerParam,
		},
	}

	jLogger := jaegerlog.StdLogger

	closer, err := cfg.InitGlobalTracer(
		"template-api-go",
		jaegercfg.Logger(jLogger),
	)

	if err != nil {
		return nil, fmt.Errorf("error initializing global tracer: %w", err)
	}

	return closer, nil
}
