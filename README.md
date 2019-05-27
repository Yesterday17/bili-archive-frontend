# bili-archive-frontend

## 简介

[bili-archive-frontend](https://github.com/Yesterday17/bili-archive-frontend) 是对应 [bili-archive](https://github.com/Yesterday17/bili-archive) 后端的一个前端实现。  
从`2019年5月27日`起，`bili-archive-frontend`不再作为`bili-archive`的一部分存在，而是`bili-archive`后端对应的一个前端实现，你可以参考该项目以及[后端接口文档](https://github.com/Yesterday17/bili-archive/docs)实现你自己的版本。

## 如何使用

运行对应平台的可执行文件即可。（请勿关闭 `Console`）

## 如何构建

1. `clone`该项目

```bash
git clone https://github.com/Yesterday17/bili-archive-frontend.git
cd ./bili-archive-frontend
```

2. 获得 `statik`

```bash
go get github.com/rakyll/statik
```

3. 编译前端

```bash
# yarn install 的 install 可以省略
yarn

# 构建
yarn build
```

4. 打包前端

```bash
statik -src=./dist -f
```

5. 构建

```bash
set CGO_ENABLED=0

# 选择架构(amd64, 386, arm)
set GOARCH=amd64

# 选择平台(windows, linux, darwin, freebsd)
set GOOS=windows
go build -o ./build/bili_archive_front_end_windows_x64.exe

set GOOS=darwin
go build -o ./build/bili_archive_front_end_darwin_amd64

set GOOS=linux
go build -o ./build/bili_archive_front_end_linux_amd64
```

## 协助开发

```
yarn run serve
```

### Lints

```
yarn run lint
```
