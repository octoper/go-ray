## Debug your Go with Ray to fix problems faster

![Latest Version](https://img.shields.io/github/v/tag/octoper/go-ray)
![Go version](https://img.shields.io/github/go-mod/go-version/octoper/go-ray)
![Tests](https://github.com/octoper/go-ray/workflows/Tests/badge.svg)
[![Golang CI Lint](https://github.com/octoper/go-ray/actions/workflows/linter.yml/badge.svg)](https://github.com/octoper/go-ray/actions/workflows/linter.yml)
![License](https://img.shields.io/github/license/octoper/go-ray)
[![Go Reference](https://pkg.go.dev/badge/github.com/octoper/go-ray.svg)](https://pkg.go.dev/github.com/octoper/go-ray)
[![Buy us a tree](https://img.shields.io/badge/Treeware-%F0%9F%8C%B3-lightgreen)](https://plant.treeware.earth/octoper/go-ray)

This module can be installed in any Go application to send messages to the Ray app.

## Install

When using Go Modules, you do not need to install anything to start using Ray with your Go program. Import the module
and the go will automatically download the latest version of the module when you next build your program.

```go
import (
    "github.com/octoper/go-ray"
)
```

With or without Go Modules, to use the latest version of the SDK, run:

`go get github.com/octoper/go-ray`

Consult the [Go documentation on Modules](https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies) for more information on how to manage dependencies.

## Documentation

You can find the full documentation on [our documentation site](https://spatie.be/docs/ray).

## Testing

```bash
go test -v
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](.github/CONTRIBUTING.md) for details.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/octoper/go-ray. This project is intended to be
a safe, welcoming space for collaboration, and contributors are expected to adhere to
the [code of conduct](https://github.com/spatie/ray/blob/master/CODE_OF_CONDUCT.md).

## Contributing

Please see [CODE_OF_CONDUCT](.github/CODE_OF_CONDUCT.md) for details.

## Security Vulnerabilities

Please review [our security policy](SECURITY.md) on how to report security vulnerabilities.

## Credits

- [Vaggelis Yfantis](https://github.com/octoper)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

## Treeware

You're free to use this package, but if it makes it to your production environment you are required to buy the world a tree.

It’s now common knowledge that one of the best tools to tackle the climate crisis and keep our temperatures from rising above 1.5C is to <a href="https://www.bbc.co.uk/news/science-environment-48870920">plant trees</a>. If you support this package and contribute to the Treeware forest you’ll be creating employment for local families and restoring wildlife habitats.

You can buy trees here [offset.earth/treeware](https://plant.treeware.earth/octoper/go-ray)

Read more about Treeware at [treeware.earth](http://treeware.earth)
