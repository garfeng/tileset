# 介绍

我之前写过一个去除黑边的程序，但该程序只能处理边缘处透明度被邻近图块影响的情况，无法处理边缘颜色互相扩散导致的问题。

比如：相邻两个图块，A（左）是白色地砖，B（右）是青石板，用waifu2x放大后，A的右侧一定会变绿，B的左侧一定会变白。

该程序一起解决这两个问题。

## 编译

### 运行环境

1. Qt 5.5 以上
2. Golang
3. waifu2x-caffe （需要其中的waifu2x-caffe-cui，运行库，以及models）

### 界面

用Qt打开 tileset_gui，然后编译。

### 算法核心

``` shell
go install github.com/garfeng/tileset/tilesetCore
```

### 合并

将下列文件复制到同一目录：

1. tileset_gui编译出的exe文件，携带各种qt库文件
2. $GOPATH/bin/目录下的tilesetCore.exe
3. waifu2x-caffe目录下的所有文件

最后你的目录看起来是这样的：

```
program
 --+
   |-- tilesetCore.exe
   |-- tileset_gui.exe
   |-- waifu2x-caffe-cui.exe
   |-- Qtxxx.dll
   |-- models/

```