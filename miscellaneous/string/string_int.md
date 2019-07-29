总结了 golang 中字符串和各种 int 类型之间的相互转换方式

<!-- TOC -->

- [1. string 转成 int](#1-string-转成-int)
- [2. int 转成 string](#2-int-转成-string)

<!-- /TOC -->

# 1. string 转成 int

- string 转成 int:

```Go
int, err := strconv.Atoi(string)
```

- string 转成 int64:

```Go
int64, err := strconv.ParseInt(string, 10, 64)
```

# 2. int 转成 string

- int 转成 string:

```Go
string := strconv.Itoa(int)
```

- int64 转成 string:

```Go
string := strconv.FormatInt(int64,10)
```
