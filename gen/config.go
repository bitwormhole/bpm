package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigBPM(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
