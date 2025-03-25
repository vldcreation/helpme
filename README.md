# Helpme CLI Tool

A CLI tool for finding and setting up development templates, running tests, and managing dependencies.

## Features

- Find and generate code examples from golang built in (IN PROGRESS)
- Set up project templates (NOT STARTED)
- Generate randoom, secure and memorable passwords (DONE)
- Pull Go dependencies (DONE)
- Run tests with sample inputs and outputs (IN PROGRESS | DONE FOR CODEFORCE and it's similar)

## Prerequisites
- Go 1.21 or higher

## Installation
1. Clone the repository:
```bash
git clone https://github.com/vldcreation/helpme.git
```

2. Navigate to the project directory:
```bash
cd helpme
```

3. Build and install the CLI:
```bash
go install
```

## Usage

### Find Command
Search for code examples in different programming languages.

```bash
helpme find [function_name] --lang [language] [flags]

Flags:
  -l, --lang string    Language to search (go/javascript)
  -p, --pkg string     Package name (optional)
  -s, --save          Save example to a file
  -e, --exec          Run the saved example file
  -d, --dir string    Directory to save the example file (default ".")
```

### Setup Command
Set up project templates.

```bash
helpme setup [flags]
```

### Generate Password Command
Generate secure passwords with customizable options.

```bash
helpme generate-password [flags]
```

### Pull Command
Pull Go dependencies from repositories.

```bash
helpme pull [flags]

Flags:
  -H, --host string     Hostname of the repository (e.g. github.com)
  -u, --user string     Username of the repository
  -r, --repo string     Repository name
  -b, --branch string   Branch name of the repository
```

### Run Test Command
Run tests with sample inputs and compare outputs.

```bash
helpme runtest [flags]

Flags:
  -F, --file string     Filepath of file to execute (e.g. mypackage/a.go)
  -f, --func string     Function name to invoke (e.g: MyFunc)
  -D, --debug_out      Print debug output
  -i, --input string    Input path sample (.in file)
  -o, --output string   Output path sample (.out file)
```

## Examples

### Finding a Code Example
```bash
helpme find strings.Join --lang go --save
```

### Running a Test
```bash
helpme runtest -F mycode.go -f TestFunc -i test.in -o test.out
```

### Pulling a Repository
```bash
helpme pull -u username -r repo-name -H github.com
```

### Tracking Clipboard Content
```bash
helpme trackclipboard -C /path/to/track.yaml
```
sample config:
```yaml
app:
  channel: "telegram" # local | telegram
  idle: "20s"
  debug: true
file:
  path: "~/Documents/track"
  name: "" # leave empty to use current date
telegram:
  token: "<telegram_bot_token>"
  chat_id: "<telegram_chat_id>"
```

### Share Files Over Network
Share your workspace directory with other devices on the same network.

```bash
helpme sharefile -D /path/to/workspace -P 9000
```

Flags:
  -D, --dir string   Root directory of workspace to share
  -P, --port string  Port number for the file server (default "9000")

Once started, other devices on the same network can access the shared files through their web browser using the displayed local IP address and port number.

## Development

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Create your utility function to [helpme-package](https://github.com/vldcreation/helpme-package)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Update the depdencies, you can leverage helpme to pull it's own depdencies, see ```make pull r=pkg``` for example
5. Push to the branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

## Disclaimer
This tool is provided as-is and without any warranty. Use it at your own risk.

## License

This project is licensed under the MIT License.