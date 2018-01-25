## 贡献指南

我们欢迎广大开发者贡献大家的智慧，让我们共同让它变得更完美.

## 代码格式

  1. 函数名称严格按照 PHP下划线 -> 大驼峰命名
     - array() -> Array()
     - array_change_key_case() -> ArrayChangeKeyCase()

  2. 函数名后第一行留空
  3. 缩紧统一4个空格
  4. 每次提交一个函数，可以添加

示例:

```go

// +------------------------------------------------------------
// | @desc      Description of function do
// | @param     args  ...Type
// |
// | @author    Openset <openset.wang@gmail.com>
// | @link      https://github.com/openset
// | @date      2018/01/01
// |
// | @return    Type
// +------------------------------------------------------------
func FuncName(args ...Type) Type {
    //此处空一行
	return args
}

```

### 贡献代码

1. Fork [openset/php2go](https://github.com/openset/php2go) 到本地.
2. 创建新的分支：

    ```shell
    $ git checkout -b new_feature
    ```

3. 编写代码。
4. Push 到你的分支:

    ```shell
    $ git push origin new_feature
    ```

5. 创建 Pull Request 并描述你完成的功能或者做出的修改。
