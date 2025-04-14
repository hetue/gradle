package internal

import (
	"github.com/hetue/core"
	"github.com/hetue/gradle/internal/internal/config"
	"github.com/pangum/pangu"
)

type Gradle struct {
	pangu.Get

	Command *core.Command
	Binary  *config.Binary
	Source  *config.Source

	Runtime *core.Runtime
}
