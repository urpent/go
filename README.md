# Introduction

Go library

- `gox` - general
- `ut` - unit testing
- `logx` - logging
- `errx` - wrap error
- `listx` - slice comparison
- `cachex` - cache

## Versioning:
- [Semantic Versioning](https://semver.org/). For example, if the current version is `errx/v1.0.0`:
    - if you are doing a backward compatible bugfix, next version should be `errx/v1.0.1`
    - if you are doing a backward compatible change, next version should be `errx/v1.1.0`
    - if you are doing a change which is not backward compatible, next version should be `errx/v2.0.0`
- To create a new version: `git tag -a errx/v0.0.0 -m "errx/v0.0.0"`
- To push a new version: `git push origin errx/v1.0.1`
- To delete tag (in case if mistakenly created): `git tag -d errx/v1.0.1` and `git push --delete origin errx/v1.0.1`
