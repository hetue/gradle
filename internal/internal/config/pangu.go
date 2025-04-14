package config

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newGradle,

		newBinary,
		newSource,
	).Build().Apply()
}
