package step

import (
	"context"

	"github.com/goexl/args"
	"github.com/hetue/boot"
	"github.com/hetue/gradle/internal/internal/config"
	"github.com/hetue/gradle/internal/internal/step/internal/command"
)

var _ boot.Step = (*Clean)(nil)

type Clean struct {
	source *config.Source
	gradle *command.Gradle
}

func newClean(source *config.Source, gradle *command.Gradle) *Clean {
	return &Clean{
		source: source,
		gradle: gradle,
	}
}

func (c *Clean) Name() string {
	return "清理"
}

func (c *Clean) Runnable() bool {
	return true
}

func (c *Clean) Retryable() bool { // 不重试
	return false
}

func (c *Clean) Asyncable() bool { // 不异步
	return false
}

func (c *Clean) Run(ctx *context.Context) (err error) {
	return c.gradle.Exec(ctx, args.New().Build().Subcommand("clean").Build())
}
