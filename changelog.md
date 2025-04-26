## 修改记录 Modify the record


#### 0.0.1
> 1. 9223372036854776000 改为 9223372036854775807. json Unmarshal 中使用了 strconv.ParseFloat 精度丢失
> 2. model_configuration.go 中 Configuration 改为 ModelConfiguration