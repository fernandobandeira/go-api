package presenters

import (
	"api/api/middlewares"
	"api/entities"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func (p *presenters) Error(w http.ResponseWriter, r *http.Request, err error) {
	span := trace.SpanFromContext(r.Context())
	span.RecordError(err)

	switch e := err.(type) {
	case entities.StatusError:
		// We can retrieve the status here and write out a specific
		// HTTP status code.

		p.logger.Error().
			Err(err).
			Str("caller", err.(entities.StatusError).Caller).
			Str("request-id", middlewares.GetReqID(r.Context())).
			Str("trace.id", span.SpanContext().TraceID().String()).
			Msg("error")

		http.Error(w, http.StatusText(e.Code), e.Code)
		return
	default:
		// Any error types we don't specifically look out for default
		// to serving a HTTP Internal Server Error

		p.logger.Error().
			Err(err).
			Str("request-id", middlewares.GetReqID(r.Context())).
			Str("trace.id", span.SpanContext().TraceID().String()).
			Msg("unhandled error")

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
