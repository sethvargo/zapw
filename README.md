# 🕵️‍♂️ Zapw

[![GitHub Actions](https://img.shields.io/github/workflow/status/sethvargo/zapw/Test?style=flat-square)](https://github.com/sethvargo/zapw/actions?query=workflow%3ATest)

Finds common errors for [zap][zap] using static analysis. [Why?](#why)


## Installation

Run the following command to download and install `zapw` onto your system:

```sh
go get -u github.com/sethvargo/zapw/cmd/zapw
```

This will make the `zapw` CLI command available on your machine.


## Usage

Zapw does not presently have any flags or customization. Run it against any Go
project on your machine:

```sh
cd my-project/
zapw ./...
```

When the tool finds a violation, it reports it using the default analysis output:

```text
/path/to/pkg.go:13:2: zap.SugaredLogger must have an even number of "With" elements
/path/to/pkg.go:19:3: zap.SugaredLogger requires keys to be strings (got int)
```

See also: [caveats](#caveats).


## Why?

[Zap][zap] is an excellent logger, and I've used it in a variety of projects.
When used properly, it's fantastic! When used improperly... you might have a bad
day, specifically if you're using the `SugaredLogger`. That logger provides a
really nice interface for adding structured data to log messages without needing
to build fields or think about types. It's notably slower (due to the use of
reflection), so there's often a tradeoff between dev-time optimization and
runtime optimization.

If you do choose to use the `SugaredLogger`, there's two [well-documented
caveats][zap-with-caveats] which I'll summarize here:

`With` fields (e.g. `With`, `Debugw`, `Errorw`) accept a mix of loosely-typed
key-value pairs. When processing pairs, the first element of the pair is used as
the field key and the second as the field value.

1.  **Field keys must be strings.** In development, passing a non-string key
     panics. In production, an separate error is logged and execution continues.

    ```go
    logger.Warnw("oh no", 42, "universe")
                       // ^ invalid, keys must be strings
    ```

1.  **Orphaned keys are invalid.** Passing an orphaned key triggers similar
    behavior: panics in development and errors in production.

    ```go
    logger.Warn("hello", "world")
                      // ^ invalid, must be key-value pairs (orphaned key)
    ```

Emphasizing a few items from above:

1.  The system behaves differently in development vs production. This means
    there's a divergence between what's running on your local machine during
    coding and what's running in production.

1.  In development, the logger only panics if that line is executed. Without
    ~100% test coverage, you cannot be absolutely certain the cases described
    above do not exist in your codebase.

1.  In production, invalid log messages are **dropped** and an error is logged
    **instead**. That means you might be missing critical log information.

To mitigate this, [zapw][zapw] (this project) uses static analysis to find
instances of the violations listed above. It's very fast (~1.2s on 2.2M LOC) and
can easily be integrated into your CI/CD processes.


## Caveats

Zapw only finds instances that meet the following criteria:

-   Function name is a known `With` function (e.g. `With`, `Debugw`, `Errorw`)

-   Function receiver is `go.uber.org/zap.SugaredLogger`

Concretely, custom types that embed and expose the logger directly will be
found. Custom types that embed the logger privately, or functions that consume
the logger, will not be found.

```go
// good
type okEmbed struct {
  *zap.SugaredLogger
}

// bad - zapw will not find misuse of this struct
type badEmbed struct {
  logger *zap.SugaredLogger
}

func (e *badEmbed) Warnw(s string, args ...interface{}) {
  e.logger.Warnw(s, args...)
}

// bad - zapw will not find misuse of this function
func doLog(l *zap.SugaredLogger, s string, args ...interface{}) {
  l.Warnw(s, args...)
}
```

These problems are both possible to solve, but they dramatically increase the
code complexity and runtime of the analysis tool.


## License

```text
Copyright 2020 Seth Vargo

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

Zapw is not affiliated with the [zap][zap] project.




[zap]: https://github.com/uber-go/zap
[zap-with-caveats]: https://pkg.go.dev/go.uber.org/zap#SugaredLogger.With
[zapw]: https://github.com/sethvargo/zapw
