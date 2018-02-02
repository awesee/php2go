## [用Go语言实现PHP内置函数](https://openset.github.io/php2go/)

[![Build Status](https://travis-ci.org/openset/php2go.svg?branch=master)](https://travis-ci.org/openset/php2go)
[![GitHub issues](https://img.shields.io/github/issues/openset/php2go.svg?style=plastic)](https://github.com/openset/php2go/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/openset/php2go.svg?style=plastic)](https://github.com/openset/php2go/pulls)
[![GitHub commits](https://img.shields.io/github/commits-since/openset/php2go/latest.svg?style=plastic)](https://github.com/openset/php2go/commits/master)[![GitHub release](https://img.shields.io/github/release/openset/php2go.svg?style=plastic)](https://github.com/openset/php2go/releases)
[![GitHub tag](https://img.shields.io/github/tag/openset/php2go.svg?style=plastic)](https://github.com/openset/php2go/tags)
[![license](https://img.shields.io/github/license/openset/php2go.svg)](https://github.com/openset/php2go/blob/master/LICENSE)

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
