<!-- TOC -->

- [1. BSON](#1-bson)
    - [1.1. 什么是 BSON](#11-什么是-bson)
    - [1.2. 数据类型和语法编辑](#12-数据类型和语法编辑)
    - [1.3. 效率](#13-效率)

<!-- /TOC -->

# 1. BSON

## 1.1. 什么是 BSON

BSON(/ˈbiːsən/) 是一种计算机数据交换格式, 主要被用作 MongoDB 数据库中的数据存储和网络传输格式。它是一种二进制表示形式, 能用来表示简单数据结构、关联数组 (MongoDB 中称为 "对象" 或 "文档") 以及 MongoDB 中的各种数据类型。BSON 之名缘于 JSON, 含义为 Binary JSON(二进制 JSON)。

## 1.2. 数据类型和语法编辑

BSON 文档 (对象) 由一个有序的元素列表构成。每个元素由一个字段名、一个类型和一个值组成。字段名为字符串。类型包括: 

- string
- integer(32 或 64 位)
- double(64 位 IEEE 754 浮点数)
- decimal128(128 位 IEEE 754-2008 浮点数; Binary Integer Decimal 变体), 适合作为任意精度为 34 个十进制数字的数字载体, 最大值近似 10
- date(整数, 自 UNIX 时间的毫秒数)
- byte array(二进制数组)
- 布尔(true 或 false)
- null
- BSON 对象
- BSON 数组
- JavaScript 代码
- MD5 二进制数据
- 正则表达式(Perl 兼容的正则表达式, 即 PCRE, 版本 8.41, 含 UTF-8 支持; 与 Python 不完全兼容)

BSON 的类型名义上是 JSON 类型的一个超集 (JSON 没有 date 或字节数组类型), 但一个例外是没有像 JSON 那样的通用 "数字"(number) 类型。

## 1.3. 效率

与 JSON 相比, BSON 着眼于提高存储和扫描效率。BSON 文档中的大型元素以长度字段为前缀以便于扫描。在某些情况下, 由于长度前缀和显式数组索引的存在, BSON 使用的空间会多于 JSON。
