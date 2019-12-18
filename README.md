[![GitHub license](https://img.shields.io/github/license/luizvnasc/gonfig)](https://github.com/luizvnasc/gonfig/blob/master/LICENSE)
[![Build Status](https://travis-ci.com/luizvnasc/gonfig.svg?branch=master)](https://travis-ci.com/luizvnasc/gonfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/luizvnasc/gonfig)](https://goreportcard.com/report/github.com/luizvnasc/gonfig)
[![Coverage Status](https://coveralls.io/repos/github/luizvnasc/gonfig/badge.svg?branch=master)](https://coveralls.io/github/luizvnasc/gonfig?branch=master)

# Gonfig

A simple module to load configurations file to a struct.

## Getting Started

To install this package run:

```
go get github.com/luizvnasc/gonfig
```

### Example

``` golang

package main

import (
	"fmt"
	"github.com/luizvnasc/gonfig"
)

//Config struct to store the app configuration
type Config struct {
	Version     string `toml:"version"`
	Description string `toml:"desc"`
	Redis       struct {
		Host string `toml:"host"`
		Port uint   `toml:"port"`
	} `toml:"redis"`
}

func main() {
	var config Config
	gonfig.Load(&config,"config.toml")
	fmt.Printf("%v", config)
}
```

If you want to use environment variables:

```go
package main

import (
	"fmt"

	"github.com/luizvnasc/gonfig"
)

//Config struct to store the app configuration
type Config struct {
	Version     string `env:"VERSION"`
	Description string `env:"DESCRIPTION"`
	Redis       struct {
		Host string `env:"REDIS_HOST"`
		Port uint   `env:"REDIS_PORT"`
	}
}

func main() {
	var config Config
	gonfig.Load(&config)
	fmt.Printf("%v", config)
}
```

You can see more examples [here](https://github.com/luizvnasc/gonfig/tree/master/examples).

## Supported formats

|        Format         |   is supported?    |
| :-------------------: | :----------------: |
|         json          | :heavy_check_mark: |
|          xml          | :heavy_check_mark: |
|         yaml          | :heavy_check_mark: |
|         toml          | :heavy_check_mark: |
| environment variables | :heavy_check_mark: |

## Dependencies

|    Dependency    |                             Repository                             |
| :--------------: | :----------------------------------------------------------------: |
| gopkg.in/yaml.v3 | [https://github.com/go-yaml/yaml](https://github.com/go-yaml/yaml) |
|     go-toml      |    [github.com/pelletier/go-toml](github.com/pelletier/go-toml)    |
|      go-env      |      [github.com/luizvnasc/goenv](github.com/luizvnasc/goenv)      |

## Authors
* Luiz Augusto Volpi Nascimento - Initial work - [@luizvnasc](https://github.com/luizvnasc)

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/luizvnasc/gonfig/blob/master/LICENSE) file for details
