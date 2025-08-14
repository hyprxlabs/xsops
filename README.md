# xsops

Uses sops as a local secret store.  Xsops wraps the sops command-line
tool and enables you to manage secrets using json files.

As long as the key remains out of a git repository, you can use xsops
to manage secrets in a local file in a git repository the same way
you would use a `.env` file.

## Installation

```bash
go install github.com/hyprxlabs/xsops@latest
```

## Usage

```bash
xsops [command] [flags]
```

## Commands

- `xsops ls <uri>`: List secrets in the current directory.
- `xsops get <uri> <key>`: Get a secret by key.
- `xsops set <uri> <key>`: Set a secret by key.
- `xsops rm <uri> <key>`: Remove a secret by key.
- `xsops init [directory]`: Ensure an age file exists, creates a .sops.yaml
   file and creates xsops.secrets.json file in the specified directory if it does
   not exist. If no directory is specified, it defaults user's data home directory.
- `xsops ensure <url> <key>`: Ensure a secret exists by key, if it does not exist,
   it will be created using a cryptographically secure random value that defaults
   to NIST standards.
- `xsops edit <url>`: Allows editing of the secrets file in a text editor and then saves
    the changes back to the file when the file is closed.

The `uri` can be a relative file path, absolute file path, or a URL using the file://
or sops:// scheme. The `key` is the name of the secret you want to manage.

IF the `uri` is set to `default` or `'', it will use the home data directory.

If the `uri` is set to `.`, it will use the current working directory and assume
the file is named `xsops.secrets.json`.
