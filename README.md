# docker

使用Go生成`docker-compose.yml`的内容。

### 使用方法

导入Go包：

```go
import "gitee.com/xiaochengtech/docker"
```

构造`docker-compose.yml`配置文件对象，并赋值内容：

```go
item := docker.Yml{
    XXX: yyy,
    ....
}
```

调用编解码函数：

```go
// 编码
content := docker.MarshalYaml(item)

// 解码
var newItem docker.Yml
err := docker.UnmarshalYaml(content, &newItem)
```

### 资料

* [docker-compose配置文件格式说明](https://docs.docker.com/compose/compose-file/)
