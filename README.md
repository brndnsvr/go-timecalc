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

Install to `/usr/local/bin/timecalc`:

```sh
make install-local
```

## Input

`timecalc` accepts zero-or-greater decimal hour values, such as `0`, `0.25`, `1`, or `1.5`.
