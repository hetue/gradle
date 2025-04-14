package config

import (
	"github.com/pangum/pangu"
)

type Gradle struct {
	Binary *Binary `default:"{}" json:"binary,omitempty"`
	Source *Source `default:"{}" json:"source,omitempty"`
}

func newGradle(config *pangu.Config) (gradle *Gradle, err error) {
	gradle = new(Gradle)
	err = config.Build().Get(gradle)

	return
}
