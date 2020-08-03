## tidraw

tidraw 是由 golang 编写的 cli 工具。 在与 tidb 连接并运行一段时间后，tidarw 可以在 tidb dashboard 的 keyvisualizer 面板上显示指定图片所描述的灰阶图像。

## 目录

- 项目结构
- 开发环境
- 快速使用
- 样例展示

## 项目结构
```
├─ assets    // 图片资源
├─ cmd       // 程序入口
├─ examples  // 运行截图
└─ pkg
    ├─ model // 数据库操作
    └─ pixel // 像素操作
```
## 开发环境
go版本：14.4

tidb版本：v4.0.4 (playground)

操作系统：CentOS 7.5 (4C 8G)


## 快速使用

```
go build

./tidraw -host localhost -port 4000 -file assets/butterfly.jpg
```

## 样例展示
![example](https://github.com/kqzh/tidraw/blob/master/examples/example.jpg)

