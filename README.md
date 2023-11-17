# tmplx

tmplx is a simple template renderer based on the [Go template engine](https://golang.org/pkg/text/template/). It is designed to be used in a shell pipeline. It reads a template from disk, parses it, reads data from ***env***, ***json file*** or ***yaml file***, and renders the template to ***stdout*** (dry-run) or to a ***file***.

## Installation

We provide pre-built binaries for Linux, MacOS and Windows. You can download them from the [releases page](https://github.com/leonsteinhaeuser/tmplx/releases). Alternatively, you can build the binary yourself using the [build instructions](#build), using the [Docker image](#docker) or using Go directly:

### Go

```bash
go install github.com/leonsteinhaeuser/tmplx@latest
```

### Docker

You can also use the Docker image to run tmplx. The image is based on [Alpine Linux](https://alpinelinux.org/) and is only a few MB in size.

## Build

To build the binary yourself, you need to have Go installed. Then, clone the repository and run `go build -o tmplx .` in the root directory of the repository.

## Usage

| Short Flag | Long Short | Type | Default | Description |
| -----------| -----------| ---- | --------| ----------- |
| `-d` | `--dry-run` | bool | If set, the output will not be written to the output file. |
| `-f` | `--format` | string | The format of the template data. Valid values are: env, json, yaml. (default "env") |
| `-h` | `--help` | |help for tmpls |
| `-o` | `--output` | string | The path to the output file. (default "output.txt") |
| `-s` | `--source` | string | The path to the source file containing template data. Valid values are: <file>.<json|yaml>. If format is env, this flag caries the  |prefix for the environment variables. (default "TMPLX_")
| `-t` | `--template` | string | The path to the template file. (default "template.tmpl") |
| `-v` | `--version` | | version for tmpls |

### Example

The following example shows how to use tmplx to render a template using data from environment variables.

### Setup environment variables

```bash
export TMPLX_USER_ACCOUNT_0="user0:pass0"
export TMPLX_USER_ACCOUNT_1="user1:pass1"
export TMPLX_USER_ACCOUNT_2="user2:pass2"
export TMPLX_USER_ACCOUNT_3="user3:pass3"
export TMPLX_USER_ACCOUNT_4="user4:pass4"
export TMPLX_USER_ACCOUNT_5="user5:pass5"
export TMPLX_USER_ACCOUNT_6="user6:pass6"
export TMPLX_USER_ACCOUNT_7="user7:pass7"
export TMPLX_USER_ACCOUNT_8="user8:pass8"
export TMPLX_USER_ACCOUNT_9="user9:pass9"
export TMPLX_USER_ACCOUNT_10="user10:pass10"
```

### Execute tmplx

```bash
tmplx -t _tests/template.tmpl --dry-run
```
