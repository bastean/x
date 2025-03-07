<h1 align="center">

<!-- [![README Logo](https://raw.githubusercontent.com/bastean/x/main/assets/readme/logo.png)](https://github.com/bastean) -->

[![README Logo](assets/readme/logo.png)](https://github.com/bastean/x)

</h1>

<br />

<div align="center">

[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://github.com/commitizen/cz-cli)
[![Release It!](https://img.shields.io/badge/%F0%9F%93%A6%F0%9F%9A%80-release--it-orange.svg)](https://github.com/release-it/release-it)

</div>

<div align="center">

[![Upgrade workflow](https://github.com/bastean/x/actions/workflows/upgrade.yml/badge.svg)](https://github.com/bastean/x/actions/workflows/upgrade.yml)
[![CI workflow](https://github.com/bastean/x/actions/workflows/ci.yml/badge.svg)](https://github.com/bastean/x/actions/workflows/ci.yml)
[![Release workflow](https://github.com/bastean/x/actions/workflows/release.yml/badge.svg)](https://github.com/bastean/x/actions/workflows/release.yml)

</div>

<div align="center">

[![GitHub Releases](https://img.shields.io/github/v/release/bastean/x.svg)](https://github.com/bastean/x/releases)

</div>

## Modules

| Module         | Reference                                                                                                                 | Status                                                                                                                                     | Latest                                                                                    |
| -------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------- |
| [tools](tools) | [![Go Reference](https://pkg.go.dev/badge/github.com/bastean/x/tools.svg)](https://pkg.go.dev/github.com/bastean/x/tools) | [![Go Report Card](https://goreportcard.com/badge/github.com/bastean/x/tools)](https://goreportcard.com/report/github.com/bastean/x/tools) | ![Version](https://img.shields.io/github/v/tag/bastean/x?filter=tools%2Fv*&label=release) |

### Workflow

> [!IMPORTANT]
> To add or remove a module within [go.work](go.work), we must use the following tasks to synchronize the workflow of [module.yml](.github/workflows/module.yml) with the new changes in the workspace.

Add new module to workspace

```bash
task work-use-"<module>"
```

Remove module from workspace

> [!WARNING]
> This task will also delete the specified module folder.

```bash
task work-drop-"<module>"
```

## First Steps

### Clone

#### HTTPS

```bash
git clone https://github.com/bastean/x.git && cd x
```

#### SSH

```bash
git clone git@github.com:bastean/x.git && cd x
```

### Initialize

#### Locally

1. System Requirements

   - [Go](https://go.dev/doc/install)
   - [Task](https://taskfile.dev/installation)

2. Run

   ```bash
   task init
   ```

### Run

#### Tests

##### Unit (Single-Module)

```bash
cd <module> && task test-unit
```

##### Unit (Multi-Module)

```bash
task test-units
```

##### Integration (Single-Module)

```bash
cd <module> && task test-integration
```

##### Integration (Multi-Module)

```bash
task test-integrations
```

##### Acceptance (Single-Module)

```bash
cd <module> && task test-acceptance
```

##### Acceptance (Multi-Module)

```bash
task test-acceptances
```

##### Unit / Integration / Acceptance (Multi-Module)

```bash
task tests
```

## Tech Stack

#### Base

- [Go](https://go.dev)

#### Please see

- [package.json](package.json)

## Contributing

Contributions and Feedback are always welcome!

- [Open a new issue](https://github.com/bastean/x/issues/new/choose)

## License

[MIT](LICENSE)
