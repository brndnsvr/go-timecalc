# go-timecalc

`timecalc` is a tiny Go CLI that converts decimal hours into minutes and seconds.

## Usage

Pass the hour value as a positional argument:

```sh
timecalc 1.5
```

Output:

```text
1.5 hours = 90 minutes = 5400 seconds
```

If you omit the argument, `timecalc` prompts for it:

```sh
timecalc
```

## Build

Build a local binary:

```sh
make build
```

Run tests:

```sh
make test
```

Install using the platform default selected by the `Makefile`:

```sh
make install-local
```

On Apple Silicon Macs, `make install-local` defaults to `/opt/homebrew/bin`. On other systems it uses `/usr/local/bin` when available, otherwise `~/.local/bin`.

Install to a specific location by overriding `BINDIR`:

```sh
make install-local BINDIR=/usr/local/bin
```

Other common destinations:

```sh
make install-local BINDIR=/opt/homebrew/bin
```

```sh
make install-local BINDIR=$HOME/.local/bin
```

If `~/.local/bin` is already in your `PATH`, `timecalc` will be available as:

```text
~/.local/bin/timecalc
```

Writing to `/usr/local/bin` may require `sudo` on some Macs:

```sh
sudo make install-local BINDIR=/usr/local/bin
```

## Input

`timecalc` accepts zero-or-greater decimal hour values, such as `0`, `0.25`, `1`, or `1.5`.
