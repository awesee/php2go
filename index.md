## [用Go语言实现PHP内置函数](https://openset.github.io/php2go)

[![Build Status](https://travis-ci.org/openset/php2go.svg?branch=master)](https://travis-ci.org/openset/php2go){:target="_blank"}
[![GitHub issues](https://img.shields.io/github/issues/openset/php2go.svg)](https://github.com/openset/php2go/issues){:target="_blank"}
[![GitHub pull requests](https://img.shields.io/github/issues-pr/openset/php2go.svg)](https://github.com/openset/php2go/pulls){:target="_blank"}
[![GitHub commits](https://img.shields.io/github/commits-since/openset/php2go/latest.svg)](https://github.com/openset/php2go/commits/master){:target="_blank"}
[![GitHub contributors](https://img.shields.io/github/contributors/openset/php2go.svg)](https://github.com/openset/php2go/graphs/contributors){:target="_blank"}
[![license](https://img.shields.io/github/license/openset/php2go.svg)](https://github.com/openset/php2go/blob/master/LICENSE){:target="_blank"}
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/openset/php2go.svg?colorB=green)](https://github.com/openset/php2go/archive/master.zip){:target="_blank"}

[English](https://github.com/openset/go4php){:target="_blank"} | [简体中文](https://github.com/openset/php2go){:target="_blank"}

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

[More](https://github.com/openset/php2go/blob/master/main.go){:target="_blank"}

### 项目进度

[TODO List](https://github.com/openset/php2go/blob/master/TODO.md){:target="_blank"}

### 贡献代码

[贡献指南](https://github.com/openset/php2go/blob/master/.github/CONTRIBUTING.md){:target="_blank"}

## Contributors

[Your contributions are always welcome!](https://github.com/openset/php2go/graphs/contributors){:target="_blank"}

## LICENSE

Released under [MIT](https://github.com/openset/php2go/blob/master/LICENSE){:target="_blank"} LICENSE
