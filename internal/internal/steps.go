package internal

import (
	"github.com/hetue/gradle/internal/internal/step"
	"github.com/pangum/pangu"
)

type Steps struct {
	pangu.Get

	Todo *step.Todo
}
