package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newGradle,

		newBinary,
		newSource,
	).Build().Apply()
}
