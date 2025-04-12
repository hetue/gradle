package internal

import (
	"github.com/hetue/core"
	"github.com/hetue/gradle/internal/internal"
)

func New(params internal.Steps) []core.Step {
	return []core.Step{
		params.Todo,
	}
}
