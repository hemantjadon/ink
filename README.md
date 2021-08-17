# ink

Package `ink` is a simple yet flexible logging facade for `go` projects. The
idea behind `ink` is to separate the logging API from the actual logging backend.
The logging backends can have their own implementations to handle logs, and
can facilitate variety of features, but the application code is not concerned
with any of this.

## The philosophy

Each package comes with their own philosophies and way to do certain things.  
`ink` is no different, and has some ideas which are core to the package.

* Logging backends must be as transparent as possible with minimal alterations.
* API must be consistent for each backend and should not depend upon the features provided by individual backends.
* Minimal overhead in terms of allocations and processing.
* API should encourage structured logging.
* API should encourage levelled logging.

## Usage

```go
package main

import (
	"log"
	"time"
	
	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/contrib/log/inklog"
)

func main() {
	// Actual logging backend of your system.
	logger := log.Default()

	ink.Log(inklog.Wrap(logger)).Info(
		"detected system health",
		ink.String("name", "ink"),
		ink.Uint64("uptime_days", 9),
		ink.Int64("temperature_celsius", -25),
		ink.Float64("load_factor", 0.275),
		ink.Bool("is_active", true),
		ink.Duration("timeout", 5*time.Second),
		ink.Strings("traits", []string{"primary", "master"}),
	)

	// Output:
	// INFO detected system health name=ink uptime_days=9 temperature_celsius=-25 load_factor=0.275 is_active=true timeout=5s traits=[primary master]	
}
```

For more examples and documentation see [pkg.go.dev](https://pkg.go.dev/hemantjadon/ink)

### Is this middleware to logging library really worth it in go?

The simple answer to this question, like any other questions in Software 
Engineering is *it depends*. For isolated applications, the standard library's 
`log` package is more than enough, or you can choose between many great 
open-source logging libraries.

The complexity arises when you are working in a team on multiple projects, which
have different logging libraries, then the log statements of each project become
different, and figuring out what features a particular logging library supports 
becomes another task. Now having one standard logging library across all the 
projects in the team is a good solution, but it locks the application with a 
logging library, and in future if you want to change it (for performance or any 
other reason), the complexity arises that now the actual code change is needed 
across the whole application supporting the API of the new library, which is 
difficult to do if not completely infeasible.

This is where `ink` comes in, with a central **opinionated** API, with good 
compatibility major logging backends out of the box, and easily extensible
for custom backends, everything without changing a single line in the core 
application code.

### But we don't change the logging backends, then what's the benefit?

The core benefit is the *API*, which is the core idea behind this project. The
logging API should be as simple as possible but yet provide the all commonly
expected features to the client. The API of `ink` intends to be just that.

Using `ink` code will look self similar and yet will provide the flexibility
of different logging backends.
