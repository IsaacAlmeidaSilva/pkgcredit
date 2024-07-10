package correlationid

import (
	"context"
	"github.com/google/uuid"
)

const CorrelationKey = "correlation-id"

var CorrelationsTypes = []string{
	"correlation-id",
	"correlation_id",
	"correlationId",
	"correlationID",
	"correlation-Id",
	"correlation-Id",
	"correlation-ID",
}

func GetFromContext(ctx context.Context) uuid.UUID {
	value := ctx.Value(CorrelationKey)
	if value != nil {
		switch value.(type) {
		case uuid.UUID:
			return value.(uuid.UUID)
		case string:
			valueUUID, err := uuid.Parse(value.(string))
			if err != nil {
				return uuid.Nil
			}
			return valueUUID
		}
	}
	newUUID := uuid.New()
	return newUUID
}

func SetOnContext(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, CorrelationKey, value)
}
