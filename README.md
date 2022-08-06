# Introduction

Go library

- `ut` - helpers for unit testing
- `arrayx` - helpers for array comparison

## Versioning:
- We use [Semantic Versioning](https://semver.org/). For example, if the current version is `ut/v1.0.0`:
    - if you are doing a backward compatible bugfix, next version should be `ut/v1.0.1`
    - if you are doing a backward compatible change, next version should be `ut/v1.1.0`
    - if you are doing a change which is not backward compatible, next version should be `ut/v2.0.0`
- To create a new version: `git tag -a ut/v1.0.1 -m "ut/v1.0.1"` or `git tag -a ut/v1.0.1 -m "ut/v1.0.1"`
- To push a new version: `git push origin ut/v1.0.1`
- To delete tag (in case if mistakenly created): `git tag -d ut/v1.0.1` and `git push --delete origin ut/v1.0.1`
