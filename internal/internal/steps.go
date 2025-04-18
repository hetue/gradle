package internal

import (
	"github.com/harluo/di"
	"github.com/hetue/gradle/internal/internal/step"
)

type Steps struct {
	di.Get

	Todo *step.Clean
}
