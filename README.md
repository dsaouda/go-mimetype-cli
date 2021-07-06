[![Go Report Card](https://goreportcard.com/badge/github.com/dsaouda/go-mimetype-cli)](https://goreportcard.com/report/github.com/dsaouda/go-mimetype-cli)

# go-mimetype-cli
Cli to project github.com/gabriel-vasile/mimetype

# Usage

```bash
>> go-mimetype-cli.exe -file README.md`
>> {"mime":"text/plain; charset=utf-8","extension":".txt"}
```

## Motivation
Simple way to check a file's type by looking at bytes. On linux you will prefer to use the file command. On windows this cli can be an option.
