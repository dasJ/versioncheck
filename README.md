# versioncheck

versioncheck is a tool supposed to make it easier for package maintainers to keep up with upstream updates.
It's supposed to be called in a cron job and will check if any versions were changed.
If something changed, a custom script is invoked.

## Installation

After setting up your `GOPATH`, run:

```
$ go get github.com/dasJ/versioncheck
```

versioncheck can be found in `$GOPATH/bin`.


TODO How to configure
TODO How to extend
TODO Document existing modules
