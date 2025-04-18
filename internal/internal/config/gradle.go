package config

import (
	"github.com/harluo/boot"
)

type Gradle struct {
	Binary *Binary `default:"{}" json:"binary,omitempty"`
	Source *Source `default:"{}" json:"source,omitempty"`
}

func newGradle(config *boot.Config) (gradle *Gradle, err error) {
	gradle = new(Gradle)
	err = config.Build().Get(gradle)

	return
}
