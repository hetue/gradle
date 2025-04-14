package step

import (
	"context"

	"github.com/goexl/args"
	"github.com/hetue/core"
	"github.com/hetue/gradle/internal/internal/config"
	"github.com/hetue/gradle/internal/internal/step/internal/command"
)

var _ core.Step = (*Build)(nil)

type Build struct {
	source *config.Source
	gradle *command.Gradle
}

func newBuild(source *config.Source, gradle *command.Gradle) *Build {
	return &Build{
		source: source,
		gradle: gradle,
	}
}

func (b *Build) Name() string {
	return "编译"
}

func (b *Build) Runnable() bool {
	return true
}

func (b *Build) Retryable() bool { // 不重试
	return false
}

func (b *Build) Asyncable() bool { // 不异步
	return false
}

func (b *Build) Run(ctx *context.Context) (err error) {
	return b.gradle.Exec(ctx, args.New().Build().Subcommand("build").Build())
}
