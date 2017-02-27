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

## Configuration and running

Create a configuration file like this:

```
{
	"dbLocation": "/var/db/versioncheck/db.json",
	"notificator": "/usr/bin/versioncheck-notify.sh",
	"upstreams": [
		{
			"name": "Autoenv",
			"module": "github",
			"params": {
				"namespace": "kennethreitz",
				"project": "autoenv"
			},
			"tags": [ "aur", "dotfiles" ]
		}
	]
}
```

The directory containing the database file must be writeable and must exist.
If the database location is omitted, the value from above is used as default.

See below for inforamtion about the nofificator and the upstreams.

After creating the notifier, you can run versioncheck like this:

```
$ $GOPATH/bin/versioncheck /path/to/configuration.json
```

## About the notificator

The notificator must be an executable file written in a language that can parse arguments.
It will be called like this:

```
$ notificator [name of the upstream] [module] [old version] [new version] [tag]...
```

Each tag is an own parameter.
Stdout and Stderr will be redirected to the terminal.
An example notificator can be found in the [doc directory](doc/notificator-example.sh).

## About upstreams

An upstream is a project that versioncheck should track.
When versioncheck discovers a new upstream that wasn't there the last time, it the current version is silently added to the database.
If you remove an upstream, the old version is silently removed from the database.

An upstream has these attributes:

- name. This name is a unique identifier for your upstream. When changed, the upstream is removed and re-added with the new name.
- module. The module that is used to find the current version of the project.
- parameters. These parameters configure the module and are module-specific.
- tags. Tags are for your own information. They are forwarded to the notificator and can be parsed in any way.

## Implementing own modules

TODO

TODO Document existing modules
TODO Error handling
