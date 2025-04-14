package command

import (
	"context"

	"github.com/goexl/args"
	"github.com/goexl/log"
	"github.com/hetue/core"
	"github.com/hetue/gradle/internal/internal/config"
	"github.com/hetue/gradle/internal/internal/step/internal/command/internal"
)

type Gradle struct {
	command *core.Command
	binary  *config.Binary
	source  *config.Source

	runtime *core.Runtime
	logger  log.Logger
}

func NewGradle(gradle internal.Gradle) *Gradle {
	return &Gradle{
		command: gradle.Command,
		binary:  gradle.Binary,
		source:  gradle.Source,

		runtime: gradle.Runtime,
	}
}

func (g *Gradle) Exec(ctx *context.Context, arguments *args.Arguments) (err error) {
	if g.runtime.Verbose { // 全栈日志
		arguments.Rebuild().Flag("stacktrace")
	}
	if g.logger.Enabled(log.LevelDebug) { // 调试日志
		arguments.Rebuild().Flag("debug")
	}
	command := g.command.New(g.binary.Gradle).Arguments(arguments).Dir(g.source.Dir)
	_, err = command.Context(*ctx).Build().Exec()

	return
}
