# whatdidido

A simple CLI tool to show you easy to read git commits history so you can tell your manager you are not slacking.
Or you are just too lazy to memorize git flags.

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
whatdidido check [--author=<name>] [--since=<time>]
whatdidido config [--author=<name>] [--since=<time>]
```

### Flags

- `--author, -a`  Git author name (e.g., "johndoe")
- `--since, -s`   Time specification (e.g., "1 day ago", "midnight")

### Running git log

Use the `check` command with optional flags to view commits.

Examples:

```bash
whatdidido check
whatdidido check --author johndoe
whatdidido check -a johndoe --since "1 day ago"
whatdidido check --since "1 day ago"
whatdidido check -s "1 day ago"
```

### Configuration

Persistent configuration is stored in `~/.config/whatdidido/config.json`.

Use the `config` command with flags to set or view configuration.

Examples:

```bash
whatdidido config                    # Show current config
whatdidido config --author johndoe  # Set author
whatdidido config --since "1 day ago"  # Set since
```

This runs: `git log --since=<since> --author=<author> --no-merges --pretty=format:"%s"`

## Testing

Run the test suite:

```bash
go test
```

Or run the comprehensive test script:

```bash
./test.sh
```

The test suite covers:

- Configuration loading and saving
- Flag parsing and validation
- Edge cases like empty flags
- Error handling
