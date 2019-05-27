# bili-archive-frontend

## 如何使用

运行可执行文件即可。  
（请勿关闭 `Console`）

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
yarn install
yarn build
```

4. 打包前端

```bash
statik -src=./public -f
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

### Compiles and hot-reloads for development

```
yarn run serve
```

### Lints and fixes files

```
yarn run lint
```

### Customize configuration

See [Configuration Reference](https://cli.vuejs.org/config/).
