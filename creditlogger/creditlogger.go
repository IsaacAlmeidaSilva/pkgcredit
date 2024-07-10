package creditlogger

import (
	"context"
	"github.com/IsaacAlmeidaSilva/pkgcredit/correlationid"
	"github.com/OlaIsaac/horcrux/modules/logger"
	"os"
	"time"
)

const loggerContext = "logger_ctx"

func setMinimalRequirementFields(ctx context.Context, l *logger.Logger) {
	correlationID := correlationid.GetFromContext(ctx)

	l.SetField(
		logger.Field(correlationid.CorrelationKey, correlationID),
		logger.Field("timestamp", time.Now()),
		logger.Field("application-version", os.Getenv("GIT_TAG")),
	)
}

func setKeyValueFields(input map[string]any, l *logger.Logger) {
	for key, value := range input {
		l.SetField(
			logger.Field(key, value),
		)
	}
}

func AddLoggerToCtx(ctx context.Context, input map[string]any) context.Context {
	l := logger.GetFromCtx(ctx)

	setMinimalRequirementFields(ctx, l)
	setKeyValueFields(input, l)

	ctx = context.WithValue(ctx, loggerContext, l)
	return ctx
}

func GetCreditLogger(ctx context.Context) *logger.Logger {
	if ctxLog, ok := ctx.Value(loggerContext).(*logger.Logger); ok {
		return ctxLog
	} else {
		l := logger.GetFromCtx(ctx)
		l.Warn(ctx, "logger not found in ctx")
		return l
	}
}
