package config

type Source struct {
	// 目录
	Dir string `default:"." validate:"required" json:"dir,omitempty"`
}

func newSource(gradle *Gradle) *Source {
	return gradle.Source
}
