# meroshare-cli

A command-line tool for various utilities including timezone operations.

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd meroshare-cli
```

2. Build the application:
```bash
go build -o meroshare-cli
```

3. (Optional) Install globally:
```bash
go install
```

## Usage

### Timezone Command

Display the current date in a specified timezone.

```bash
./meroshare-cli timezone [timezone]
```

#### Examples

```bash
# Get current date in UTC
./meroshare-cli timezone UTC

# Get current date in New York timezone
./meroshare-cli timezone America/New_York

# Get current date in Tokyo timezone
./meroshare-cli timezone Asia/Tokyo

# Get current date with custom format
./meroshare-cli timezone UTC --date "2006-01-02 15:04:05"
```

#### Options

- `--date`: Specify the date format (default is RFC3339 date format)

#### Common Timezones

- `UTC` - Coordinated Universal Time
- `America/New_York` - Eastern Time
- `America/Los_Angeles` - Pacific Time
- `Europe/London` - Greenwich Mean Time
- `Asia/Tokyo` - Japan Standard Time
- `Asia/Kolkata` - India Standard Time

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o meroshare-cli
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests
5. Submit a pull request

## License

[Add your license information here]