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

- `xsops ls`: List secrets in the current directory.
- `xsops get <KEY>`: Get a secret by key.
- `xsops set <KEY>`: Set a secret by key. There are multiple ways to set a secret
  value.
  - `--stdin`: Read the secret value from standard input.
  - `--value`: Set the secret value directly in the command line.
  - `--file`: Read the secret value from a file.
  - `--env`: Read the secret value from an environment variable.
- `xsops rm  <KEY>`: Remove a secret by key.
- `xsops init`: Ensure an age file exists, creates a .sops.yaml
   file and creates xsops.secrets.json file in the specified directory if it does
   not exist. If no directory is specified, it defaults user's data home directory.
   The data home directory is typically `~/.local/share/xsops` on Linux,
   `~/Library/Application Support/xsops/data` on macOS, and `%APPDATA%\xsops\data` on Windows.
- `xsops ensure <url> <key>`: Ensure a secret exists by key, if it does not exist,
   it will be created using a cryptographically secure random value that defaults
   to NIST standards.
- `xsops edit <url>`: Allows editing of the secrets file in a text editor and then saves
    the changes back to the file when the file is closed.

## Global Flags

- `debug` or `-d`: Enable debug output.
- `vault` or `-v`: Specify a custom vault URI or relative path to use for storing secrets.
  - ./xsops.secrets.json
  - . - current working directory, assumes the file is named `xsops.secrets.json`.
  - `default` - special value that uses the home data directory for the vault.
  - `./child/xsops.secrets.json` - relative path to the secrets file.
  - `sops://full/path/to/secrets/file.json` - absolute path to the secrets file.
  - `file://full/path/to/secrets/file.json` - absolute path to the secrets file.

`$XSOPS_VAULT` environment variable can also be used to specify the vault URI. The `--vault`
flag will override the environment variable if both are set.

## Other References

- [sops](https://github.com/getsops/sops)
- [age](https://github.com/FiloSottile/age)
