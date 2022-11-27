# kubedirlint

A simple program that checks the placement of
[Kubernetes YAML files](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/)
in your project.

`kubedirlint` expects files to follow the following pattern:

```
<namespace>/<kind>/<name>.yaml
```

For instance:

```
accounting/deployment/tax-compute.yaml
```

All violations are reported to the stderr.
If violations are present, the program exits with an error code.

## Installation

Run this to build and install the application:

```
go install
```

It should place the application in `~/go/bin/kubedirlint`.
The whole application is self-contained in one executable file.

You might consider adding the directory `~/go/bin` to the `PATH`
environment variable.

## Usage

Change the directory to the catalog with your project.
Then just run the application:

```
cd myproject
kubedirlint
```

If everything is OK, it will print:

```
all 5 files are OK
```

But if something is wrong...

```
got 1 error
file accounting/tax.yml has wrong path, should be accounting/service/tax-calc.yaml
```

## Design Principles

Keep it simple.

Read-only operations.

Automated tests are cool.

Print all errors nicely to the stderr.

No network connections.

## Contributing

Issues are tracked on [GitHub](https://github.com/mslapek/kubedirlint/issues).

## Changelog

Please see [CHANGELOG.md](CHANGELOG.md).
