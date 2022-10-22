# Introduction

Go library

- `gox` - general
- `ut` - unit testing
- `logx` - logging
- `errx` - wrap error
- `listx` - slice comparison
- `linklist` - linklist
- `cachex` - cache
- `api` - api response helper
- `httpx` - http client
- `randcode` - random code or id generator

## Versioning:
- [Semantic Versioning](https://semver.org/). For example, if the current version is `v1.0.0`:
    - backward compatible bugfix, use `v1.0.1`
    - backward compatible change, use `v1.1.0`
    - not backward compatible, use `v2.0.0`

### Create Tag
- To create a new version: `git tag v0.0.0`
- To push a new version: `git push origin v0.0.0`
- To delete tag (in case if mistakenly created): `git tag -d v0.0.0` and `git push --delete origin v0.0.0`

### Delete all Tags
To delete remote tags (before deleting local tags) simply do:
```
git tag -l | xargs -n 1 git push --delete origin
```
and then delete the local copies:
```
git tag | xargs git tag -d
```