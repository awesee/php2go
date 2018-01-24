## [用Go语言实现PHP内置函数](https://openset.github.io/php2go/)

这是一个用Go语言开发的辅助库，尤其适用于熟悉PHP内置函数的开发者，将会实现PHP内置函数。

### 下载安装

```shell

go get github.com/openset/php2go/php
```

### 关于命名

PHP下划线命名改为大驼峰命名。

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

### 贡献代码

这是一个开源项目，期待你的加入。

## Contributors

[Your contributions are always welcome!](https://github.com/openset/php2go/graphs/contributors)

## LICENSE

Released under [MIT](https://github.com/openset/php2go/blob/master/LICENSE) LICENSE
