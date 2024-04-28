

## cobra cli

### 初始化

```shell
cobra-cli init --viper --author "dreamsxin@qq.com" --license "Apache-2.0"
```

### 创建 init 命令

```shell
cobra-cli add serve
```

## GORM

```shell
go install gorm.io/gen/tools/gentool@latest
```

### 从数据库生成代码

```shell
gentool -dsn "root:123456@tcp(localhost:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local" -tables "user" -outPath "./modules/hello/query" -outFile "query.go" -onlyModel -modelPkgName "models"
```
