# Atlased To PhaserJS Parser

> Simple CLI program to convert your json files generated with Atlasesd (https://witnessmonolith.itch.io/atlased) to a format that can be read by PhaserJS. By no means is a complete implementation, will be adding features as needed. Pull requests welcomed!

Binary files:
Head to: https://github.com/GAZ082/atlased_phaserjs_parser/releases

## Log

- 2024-01-24
  - Initial release.

## Building

```bash
env GOOS=linux GOARCH=amd64 go build -o atlased_phaserjs_parser_linux_x64 main.go
env GOOS=windows GOARCH=amd64 go build -o atlased_phaserjs_parser_win_x64 main.go
```
