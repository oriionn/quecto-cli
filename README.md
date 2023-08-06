# Quecto-CLI
Quecto-cli is a simple solution for use Quecto in a command line environment.

## Installation
```bash
npm install -g quecto-cli
```

## Usage
```bash
quecto <command> [options]
```

## Commands
### shorten
Shorten a URL.

#### Arguments
- `<link>`: The URL to shorten.

#### Options
- `-i, --instance <instance>`: The instance to use. (default: `https://s.oriondev.fr`) 

### unshorten

#### Arguments
- `<link>`: The URL to unshorten.

#### Options
- `-i, --instance <instance>`: The instance to use. (default: `https://s.oriondev.fr`)

### ivi
Know if a domain is a instance of Quecto.

#### Arguments
- `<instance>`: The instance to check.

### help
Display help for a command.

#### Arguments
- `<command>`: The command to display help for.

## License
[GPL3](https://github.com/oriionn/quecto-cli/blob/main/LICENSE)

## Contributors
[![Contributors](https://contrib.rocks/image?repo=oriionn/quecto-cli)]