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
whatdidido [author] [since]
whatdidido config [key] [value]
```

### Running git log

- `whatdidido` - Uses config defaults (or hardcoded defaults: `--author=ktkennychow --since=midnight`)
- `whatdidido <author>` - Uses provided author, config/default since
- `whatdidido <author> <since>` - Uses both provided values (overrides config)

Examples:

```bash
whatdidido
whatdidido johndoe
whatdidido johndoe "1 day ago"
```

### Configuration

Persistent configuration is stored in `~/.config/whatdidido/config.json`.

```bash
whatdidido config                    # Show current config
whatdidido config author <name>     # Set author
whatdidido config since <time>      # Set since
```

Examples:

```bash
whatdidido config author johndoe
whatdidido config since "1 day ago"
whatdidido config                    # Shows: Author: johndoe, Since: 1 day ago
```

This runs: `git log --since=<since> --author=<author> --no-merges --pretty=format:"%s"`
