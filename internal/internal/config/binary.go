package config

type Binary struct {
	Gradle string `default:"/usr/bin/gradle" json:"gradle,omitempty"`
}

func newBinary(gradle *Gradle) *Binary {
	return gradle.Binary
}
