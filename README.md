# gabby

Gabby is a helpful assistant for your organization's directory and calendar systems, that currently supports Office 365.

## Installing

Gabby is written in Go. You can install Gabby by running: `go install github.com/joefitzgerald/gabby/cmd/gabby@latest`.

## Usage

```
$ gabby
Usage: gabby <command>

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  calendar impact
    Perform an impact analysis of events over a given time period.

  person photo <ids> ...
    Get User Photo for ID

  person name <ids> ...
    Get User Name for ID

Run "gabby <command> --help" for more information on a command.
```
