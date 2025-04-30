# Templie

Templie is a CLI tool to manage templates. It allows you to create and manage templates and their
categories.

## Features

* Initialize new templates with categories and optionally with a working directory.
* Preview templates before using them in a tree-like structure.
* List and manage existing templates.
* Sync templates.
* Configuration with reset/show/update.
* Written in Go.

## Installation

### Using Go install

1. Make sure you have Go installed (Go 1.18 or later is recommended).
2. Run the following command to install Templie:
    ```bash
    go install github.com/stankomichal/templie@latest
    ```

### From GitHub Release

1. Visit the [releases page](https://github.com/stankomichal/templie/releases) of the project.
2. Download the appropriate binary for your platform (Linux, macOS, Windows).
3. After downloading, move the binary to a directory in your `PATH` (e.g., `/usr/local/bin` for Linux/macOS).

### From Source

1. Make sure you have Go installed (Go 1.18 or later is recommended).
2. Clone the repository:
    ```bash
    git clone https://github.com/stankomichal/templie.git
    cd templie
    ```
3. Run the following command to build and install the tool:
    ```bash
    go install
    ```

## Usage

Once installed, you can start using Templie from the command line. The general syntax for commands is:

```bash
templie <command> [arguments]
```

### Initialize a template

To initialize a new template, use the `init` command. This command allows you to create a new template and optionally.
add categories and copy content from the current working directory.

```bash
templie template init <template-name> [options]
```

### List templates

To list all available templates, use the `list` command. This command will display a tree-like structure template with
categories.

```bash
templie template list
```

### Preview a template

To preview a specific template, use the `preview` command. This command will show you the structure of the template.

```bash
templie template preview <template-name>
```

### Global Flags

* `-v`, `--version`: Show version.
* `-h`, `--help`: Show help information for the command.

## Configuration

Use the `config` command to manage the configuration of Templie. You can reset, show, or update the configuration.

* To show the current configuration: `templie config display`

* To show a variable: `templie config show <variable>`

* To update a variable: `templie config update <variable> <value>`

* To reset a variable: `templie config reset <variable>`

* To show the default configuration: `templie config default`

## Autocompletion

Templie supports autocompletion for bash, zsh, and fish shells. Use the `completion` command to generate the completion
script:

```bash
templie completion <shell>
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.