package config

import "io"

type Configuration struct {
	AllFiles                              []io.Reader
	VerboseOutput                         bool
	IgnoreContainerCpuLimitRequirement    bool
	IgnoreContainerMemoryLimitRequirement bool
	IgnoredTests                          map[string]struct{}
	EnabledOptionalTests                  map[string]struct{}
}
