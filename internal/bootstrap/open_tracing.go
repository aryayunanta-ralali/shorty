package bootstrap

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	tracerCfg "github.com/aryayunanta-ralali/shorty/pkg/tracer"
)

// RegistryOpenTracing setup
func RegistryOpenTracing(cfg *appctx.Config) opentracing.Tracer {

	if !cfg.APM.Enable {
		return opentracing.NoopTracer{}
	}

	lf := logger.NewFields(logger.EventName("TracerInitiated"))
	logger.Debug(fmt.Sprint("apm address : ", cfg.APM.Address), lf...)
	tr := opentracer.New(
		tracer.WithAgentAddr(cfg.APM.Address),
		tracer.WithService(cfg.APM.Name),
		tracer.WithGlobalTag("env", cfg.App.Env),
	)

	tracerCfg.New(cfg.App.AppName)
	opentracing.SetGlobalTracer(tr)
	return tr

}