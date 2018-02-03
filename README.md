## [用Go语言实现PHP内置函数](https://openset.github.io/php2go)

[![Build Status](https://travis-ci.org/openset/php2go.svg?branch=master)](https://travis-ci.org/openset/php2go)
[![GitHub issues](https://img.shields.io/github/issues/openset/php2go.svg)](https://github.com/openset/php2go/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/openset/php2go.svg)](https://github.com/openset/php2go/pulls)
[![GitHub commits](https://img.shields.io/github/commits-since/openset/php2go/latest.svg)](https://github.com/openset/php2go/commits/master)
[![GitHub contributors](https://img.shields.io/github/contributors/openset/php2go.svg)](https://github.com/openset/php2go/graphs/contributors)
[![license](https://img.shields.io/github/license/openset/php2go.svg)](https://github.com/openset/php2go/blob/master/LICENSE)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/openset/php2go.svg?colorB=green)](https://github.com/openset/php2go/archive/master.zip)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fopenset%2Fphp2go.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fopenset%2Fphp2go?ref=badge_shield)

[English](https://github.com/openset/go4php) | [简体中文](https://github.com/openset/php2go)

这是一个用Go语言开发的辅助库，尤其适用于熟悉PHP内置函数的开发者，将会实现PHP内置函数。

### 下载安装

```shell

go get github.com/openset/php2go/php

```

### 关于命名

PHP下划线命名转为Go大驼峰命名。

### Example:

```go

package main

import (
    "github.com/openset/php2go/php"
)

func main() {

    php.Echo("Hello ", "world!\n")

}

```

[More](https://github.com/openset/php2go/blob/master/main.go)

### 项目进度

[TODO List](https://github.com/openset/php2go/blob/master/TODO.md)

### 贡献代码

[贡献指南](https://github.com/openset/php2go/blob/master/.github/CONTRIBUTING.md)

## Contributors

[Your contributions are always welcome!](https://github.com/openset/php2go/graphs/contributors)

## LICENSE

Released under [MIT](https://github.com/openset/php2go/blob/master/LICENSE) LICENSE


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fopenset%2Fphp2go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fopenset%2Fphp2go?ref=badge_large)