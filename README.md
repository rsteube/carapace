# carapace

[![CircleCI](https://circleci.com/gh/rsteube/carapace.svg?style=svg)](https://circleci.com/gh/rsteube/carapace)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/rsteube/carapace)](https://pkg.go.dev/github.com/rsteube/carapace)
[![documentation](https://img.shields.io/badge/&zwnj;-documentation-blue?logo=gitbook)](https://rsteube.github.io/carapace/)
[![GoReportCard](https://goreportcard.com/badge/github.com/rsteube/carapace)](https://goreportcard.com/report/github.com/rsteube/carapace)
[![Coverage Status](https://coveralls.io/repos/github/rsteube/carapace/badge.svg?branch=master)](https://coveralls.io/github/rsteube/carapace?branch=master)

Command argument completion generator for [cobra]. You can read more about it here: _[A pragmatic approach to shell completion](https://dev.to/rsteube/a-pragmatic-approach-to-shell-completion-4gp0)_.


Supported shells:
- [Bash](https://www.gnu.org/software/bash/)
- [Elvish](https://elv.sh/)
- [Fish](https://fishshell.com/)
- [Ion](https://doc.redox-os.org/ion-manual/) ([experimental](https://github.com/rsteube/carapace/issues/88))
- [Nushell](https://www.nushell.sh/)
- [Oil](http://www.oilshell.org/)
- [Powershell](https://microsoft.com/powershell)
- [Tcsh](https://www.tcsh.org/) ([experimental](https://github.com/rsteube/carapace/issues/331))
- [Xonsh](https://xon.sh/)
- [Zsh](https://www.zsh.org/)

## Usage

Calling `carapace.Gen` on the root command is sufficient to enable completion script generation using the [hidden command](https://rsteube.github.io/carapace/carapace/gen/hiddenSubcommand.html).

```go
import (
    "github.com/rsteube/carapace"
)

carapace.Gen(rootCmd)
```

## Example

An example implementation can be found in the [example](./example/) folder.


## Standalone Mode

Carapace can also be used to provide completion for arbitrary commands as well.
See [rsteube/carapace-bin](https://github.com/rsteube/carapace-bin) for examples.
There is also a binary to parse flags from gnu help pages at [caraparse](https://github.com/rsteube/carapace-bin/tree/master/cmd/caraparse).

## Related Projects

- [carapace-bin](https://github.com/rsteube/carapace-bin) multi-shell multi-command argument completer
- [carapace-pflag](https://github.com/rsteube/carapace-pflag) Drop-in replacement for spf13/pflag with support for non-posix variants
- [carapace-spec](https://github.com/rsteube/carapace-spec) define simple completions using a spec file
- [carapace-spec-clap](https://github.com/rsteube/carapace-spec-clap) spec generation for clap-rs/clap

[cobra]:https://github.com/spf13/cobra
