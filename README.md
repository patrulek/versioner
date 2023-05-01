# versioner - tool to generate app's version info

## Overview

This repo aims to provide tool and package to manage Golang application's version.

:warning: This is only proof of concept. Do not use in production. :warning:

## Quick start

Install `versioner` tool:

```console
go install github.com/patrulek/versioner
```

In your application's main package insert such line:

```go
//go:generate versioner $PWD
```

And in a place where you need to use a version info:

```go
import github.com/patrulek/versioner/version

func yourFunc() {
    // use version string
    versionStr := version.Short()

    // use Version struct
    versionObj := version.Full()
}
```

Then if you want to embed version information, execute `go generate` before building application. Tool will generate version information based on current status of git repository, if such exist. If no git repository was found in current dir, this will set `v0.0.0-<buildtime>` as a version.

The cool thing is that even if you forget about executing `go generate` and you already imported `version` package, compilation will still pass without any error.

Another cool thing is that it uses `go-git` package to get repository information, so there's no need for `git` dependency.

## Changelog

- **v0.0.1 - 01.05.2023**: Proof of Concept