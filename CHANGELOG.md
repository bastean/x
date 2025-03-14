# Changelog

## [0.4.0](https://github.com/bastean/x/compare/v0.3.0...v0.4.0) (2025-03-14)

### Chores

- **release:** tools/v0.2.2 ([346f39b](https://github.com/bastean/x/commit/346f39b1635f3e3062bc8c92e6557ebcaf000a90))
- **task:** add go dependencies download ([bff3086](https://github.com/bastean/x/commit/bff308676c456155f2a5ad96aeba7a2498e8ac74))
- **task:** add vet, shadow, errcheck, ineffassign and gocyclo ([5385123](https://github.com/bastean/x/commit/5385123c39240ee40cfce6de1e4084133af56fab))
- **vscode:** add extensions and settings ([fc1a177](https://github.com/bastean/x/commit/fc1a1771b538766e113a5d526ef9f38b36032c9a))

### Continuous Integration

- **github:** add upload of test reports ([f8769b9](https://github.com/bastean/x/commit/f8769b99999427a29967dd840d017519b8f2fa84))

### Bug Fixes

- **tools:** handle unchecked errors ([af929f6](https://github.com/bastean/x/commit/af929f64afef71ae84341b4f948a0c5f02b64be4))

## [0.3.0](https://github.com/bastean/x/compare/v0.2.0...v0.3.0) (2025-03-12)

### Chores

- **deps:** upgrade ([d7caad8](https://github.com/bastean/x/commit/d7caad8c4be5629d728c79588e8737afef61e79d))
- **release:** tools/v0.2.1 ([9232ba2](https://github.com/bastean/x/commit/9232ba2f9de7192843cbaecbf539b8503b4e508b))
- **task:** add module workflow with git branches ([93ae3bb](https://github.com/bastean/x/commit/93ae3bb36f00c837b6c77ba1186d7b9d62e593cd))

### Continuous Integration

- add scans for vulnerabilities and misconfigurations ([d407869](https://github.com/bastean/x/commit/d40786922d5eac8bcef3c3a84f1e252fc40a9aae))

### Documentation

- **readme:** add workflow with git branches ([a108006](https://github.com/bastean/x/commit/a1080060943ab35a18404f5859a618f469e83947))

### Styles

- **tools:** rename file main to release ([a9c9912](https://github.com/bastean/x/commit/a9c991200adf92f3acf7e9f3b43260788f60c6c6))

### Tests

- reuse setup across all test cases in the suite ([fd82f44](https://github.com/bastean/x/commit/fd82f44f1220650e16663998db8c9a386f1a0da9))
- use data race detector ([8b12128](https://github.com/bastean/x/commit/8b1212893f646c95443ff3924bca74521e619fe3))

## [0.2.0](https://github.com/bastean/x/compare/v0.1.0...v0.2.0) (2025-03-08)

### Chores

- **deps:** upgrade ([4d18e1b](https://github.com/bastean/x/commit/4d18e1b489a34342a02c163c7398dd7c6a9b802c))
- **release:** tools/v0.1.0 ([131198a](https://github.com/bastean/x/commit/131198a370ffc1a5bcfecd20b69b0c106b1564fb))
- **release:** tools/v0.2.0 ([8f4d007](https://github.com/bastean/x/commit/8f4d0071d7bc52d728fb99372edf691b304e735b))
- **task:** add tasks to manage workspace modules ([6ea3e0d](https://github.com/bastean/x/commit/6ea3e0d1d4d5a5abaeb58068f0ebed9c8ca23a7f))
- **work:** drop example module ([6489559](https://github.com/bastean/x/commit/64895593ce502be9a14f7195612c75e57414fbf5))

### Documentation

- **readme:** add workflow with modules ([ea34cb9](https://github.com/bastean/x/commit/ea34cb9de6433f1138e161c46ead4ace7f721aa2))
- **readme:** center module table ([02f321b](https://github.com/bastean/x/commit/02f321b63b50545050eeae806169169b623d4175))

### New Features

- **tools:** add release package ([5302b5a](https://github.com/bastean/x/commit/5302b5ab132abaf823bc165ce5e7172344395c0f))

### Bug Fixes

- **ci:** use multi-module unit test ([55d0d4d](https://github.com/bastean/x/commit/55d0d4d57f294c63969cf0b4a050d09e3007a8e3))
- **task:** run tests using bash to detect pipeline errors ([e193e09](https://github.com/bastean/x/commit/e193e09ef439afb1ef477b1411c81b5fca09c050))
- **tools:** change split separator to avoid flaky tests ([ec92207](https://github.com/bastean/x/commit/ec92207b825548b93945d339d794f84f0386001a))
- **tools:** use error returned in latest tag ([0ea1182](https://github.com/bastean/x/commit/0ea1182e361cb1ca55b907ed4704f1196e1ef401))

### Tests

- **tools:** add integration test to exec ([a63a3a5](https://github.com/bastean/x/commit/a63a3a5e370e9a46f3dde656b8658c532d6cf946))

## 0.1.0 (2025-03-04)

### Chores

- **deps:** upgrade ([99fd27e](https://github.com/bastean/x/commit/99fd27e6a5503fb7502a7242c9e302900b836d00))
- **task:** turn off work mode in build ([b73adef](https://github.com/bastean/x/commit/b73adef0c6fead4b76feb319501a76520e4e5cdb))
- update codexgo to x ([f312bfa](https://github.com/bastean/x/commit/f312bfa66458203462f1562862c07cf41acab687))

### Continuous Integration

- **github:** add workflow to release modules ([98b4c31](https://github.com/bastean/x/commit/98b4c3193702b04577046cb60eb951405254853f))
- **github:** use local scanners to find secrets ([33a3d78](https://github.com/bastean/x/commit/33a3d78a483e530050a74c01d2457becbb7a4c2d))
- **github:** use minor versions for repository releases ([4b65c98](https://github.com/bastean/x/commit/4b65c98c0f9d25581690487aebfefac29c66da2a))

### New Features

- **genesis:** codexgo ([9261231](https://github.com/bastean/x/commit/92612318da2ad64bbab4a114c110d9737fc8d6c5))
