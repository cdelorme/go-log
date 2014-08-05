
# go-logger

Yet another logging utility package for golang.


## alternatives

While other libraries exist, including golangs own [log](http://golang.org/pkg/log/) package, and third party utilities like [go-logging](https://github.com/op/go-logging), they didn't agree with my use cases.  They do supply a [syslog extension](http://golang.org/pkg/log/syslog/), but it requires that you supply it with a writer, adding to the complexity of using it without writing an abstraction to cover it.

First, golang's own package only supports three methods that don't match [log-level standards](http://en.wikipedia.org/wiki/Syslog#Internet_standards); `printf`, `panic`, and `fatal`, two of which directly impact the flow of the application by throwing `panic()` or `os.Exit(1)`.

The go-logging library is actually fantastic if you want a feature-filled logging library complete with abstractions and extensible tested code.  My only complaint is the complexity; for something as simple as logging we shouldn't need all that much extra around it.


## sales pitch

My library differs in that it aims for utmost simplicity while returning control to the developer.

It:

- supports linux-standard log-level message types
- has only three options; `logLevel`, `silent`, and `nocolor`
- lets the developer control async operations via goroutines (ex. `go logger.Error(...)`)
- returns a `LogMessage` object if you wanted to do more with it

It does not:

- run operational code on behalf of the developer (like throwing `panic()` or `os.Exit(1)`)
- have wildly extensible abstraction layers or a myriad of extra features
- include test files which increase the binary size (and golang binaries are far from "tiny")


For size comparison, I build a simple main package with one line of log output including each library, and the respective sizes of the executables with the golang 1.3 compiler were:

- 2.8M go-logging
- 1.8M go-logger
- 1.7M golang log

So the **golang log** package wins in smallest size, but not utility, while _my go-logger_ package shaved off over a megabyte of executable size by not having the test files, additional classes, and abstraction layers.


## usage

Using my library is simple:

    import "github.com/cdelorme/go-logger"

The package name is `log`, and you can get a new logger and begin using it with:

    logger := log.Logger{Level: "Warning"}
    logger.Warning("message %v", object)