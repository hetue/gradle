package internal

import (
	"github.com/harluo/di"
	"github.com/hetue/boot"
	"github.com/hetue/gradle/internal/internal/config"
)

type Gradle struct {
	di.Get

	Command *boot.Command
	Binary  *config.Binary
	Source  *config.Source

	Runtime *boot.Runtime
}
