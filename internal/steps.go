package internal

import (
	"github.com/hetue/boot"
	"github.com/hetue/gradle/internal/internal"
)

func New(params internal.Steps) []boot.Step {
	return []boot.Step{
		params.Todo,
	}
}
