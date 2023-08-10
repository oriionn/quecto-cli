# Quecto-CLI
Quecto-cli is a simple solution for use Quecto in a command line environment.

## Installation
```bash
npm install -g quecto-cli
```

![Installation](docs/npm%20install.gif)

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
- `-p, --password <password>`: The password for the shortened url. (default: `null`)

![Shorten](docs/quecto%20shorten.gif)

### unshorten

#### Arguments
- `<link>`: The URL to unshorten.

#### Options
- `-i, --instance <instance>`: The instance to use. (default: `https://s.oriondev.fr`)
- `-p, --password <password>`: The password for the shortened url. (default: `null`)

![Unshorten](docs/quecto%20unshorten.gif)

### ivi
Know if a domain is a instance of Quecto.

#### Arguments
- `<instance>`: The instance to check.

![ivi valid](docs/quecto%20ivi%20valid.gif)
![ivi invalid](docs/quecto%20ivi%20invalid.gif)

### help
Display help for a command.

#### Arguments
- `<command>`: The command to display help for.

![Help](docs/quecto%20help.gif)

## License
[GPL3](https://github.com/oriionn/quecto-cli/blob/main/LICENSE)

## Contributors
[![Contributors](https://contrib.rocks/image?repo=oriionn/quecto-cli)]()