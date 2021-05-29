package presenters

import (
	"api/interfaces"

	"github.com/rs/zerolog"
)

type presenters struct {
	logger zerolog.Logger
}

func NewPresenters(logger zerolog.Logger) interfaces.Presenters {
	return &presenters{logger: logger}
}
