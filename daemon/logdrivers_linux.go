package daemon // import "github.com/fdurand/moby/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/fdurand/moby/daemon/logger/awslogs"
	_ "github.com/fdurand/moby/daemon/logger/fluentd"
	_ "github.com/fdurand/moby/daemon/logger/gcplogs"
	_ "github.com/fdurand/moby/daemon/logger/gelf"
	_ "github.com/fdurand/moby/daemon/logger/journald"
	_ "github.com/fdurand/moby/daemon/logger/jsonfilelog"
	_ "github.com/fdurand/moby/daemon/logger/local"
	_ "github.com/fdurand/moby/daemon/logger/logentries"
	_ "github.com/fdurand/moby/daemon/logger/loggerutils/cache"
	_ "github.com/fdurand/moby/daemon/logger/splunk"
	_ "github.com/fdurand/moby/daemon/logger/syslog"
)
