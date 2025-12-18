# Rush02-Go

Rush02-Go is a Go program that converts numbers into their English word representation using a dictionary file.

## Usage

```sh
go run ./ex00 <number>
```

Example:
```sh
go run ./ex00 123456
```

## Features

- Converts numbers (including large numbers) to English words.
- Uses a customizable dictionary file (`numbers.en.dict`).
- Handles edge cases like leading zeros and unsupported numbers.

## Files

- `ex00/rush02.go` — Main Go source code.
- `ex00/numbers.en.dict` — Dictionary mapping numbers to English words.

## Dictionary Format

Each line in `numbers.en.dict` should be:
```
<number>: <word>
```
Example:
```
1000: thousand
15: fifteen
```

## License

MIT License