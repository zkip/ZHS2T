# ZHS2T
使文件夹中的所有文件的文本内容进行简体/繁体转换。

## 安装

### Fork
```
git clone https://github.com/zkip/ZHS2T
```

#### 安装依赖
```
go get github.com/liuzl/gocc
```

#### 编译
```
go build .
```

## 参数说明
```
Usage of zstw:
	-h
		help
	-c string
		s2t / s2tw / s2twp / s2hk / t2s / t2tw / tw2s / tw2sp / t2hk / hk2s
	-i string
		Input direction
	-o string
		Out direction
```