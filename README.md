# whatdidido

A simple CLI tool to show your git commits since midnight.

## Installation

```bash
go install github.com/ktkennychow/whatdidido@latest
```

Or build from source:

```bash
go build -o whatdidido
```

## Usage

```bash
whatdidido
```

This runs: `git log --since=midnight --author=ktkennychow --no-merges --pretty=format:"%s"`

