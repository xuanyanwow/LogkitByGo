# 数据上报格式说明

使用TCP上报，数据协议为Goframe框架提供的

> gtcp模块提供了简单轻量级数据交互协议，效率非常高，协议格式如下：
>
> 数据长度(16bit)数据字段(变长)


[goframe框架文档](https://goframe.org/net/gtcp/conn/pkg)

## 上报规则 

使用json定义行为

```json
{
	"action": "what you want to do?",
	"data": ".... something in here"
}
```

action可以如下行为

| 行为    | 解释            |
| ------- | --------------- |
| report  | 上报单条日志    |
| reports | 批量上报日志    |
| api     | 上报单条api统计 |
| apis    | 批量上报api统计 |



## 示例

其中其他字段为固定，对应数据表字段

### 上报日志

```json
{
	"action": "report",
	"data": {
		"project_id": 4, // 项目id
		"log_category": "Api/PayBase/v2", // 日志类型
		"log_point": "curl_begin", // 日志记录点 (比如数据提交前)
		"log_sn": "666", // 日志单号 后续可以根据订单号筛选
		"log_data": "{\"name\":\"siam\",\"age\":22}", // 日志内容，可以普通字符串或者json字符串
		"log_from": "siam", // 日志来自哪个平台
	}
}
```

```json
{
	"action": "reports",
	"data": [{
		"project_id": 4,
		"log_category": "Api/PayBase/v2",
		"log_point": "curl_begin",
		"log_sn": "666",
		"log_data": "{\"name\":\"siam\",\"age\":22}",
		"log_from": "siam"
	}]
}
```



### 上报api统计

```json
{
	"action": "api",
	"data": {
		"project_id": 5, // 项目id
		"api_category": "wechat", // api类型
		"api_method": "refund", // api方法
		"is_success": 1, // 是否成功 1成功 0失败
		"consume_time": 15, //  消耗的时间 单位为毫秒 ms
		"user_from": "payment", // 用户来源平台
		"user_identify": "20200313001", // 用户标识 可以是第三方客户的调用订单号等等
		"api_response": "响应内容 可选", // api响应的内容 做记录 
		"api_param": "{\"test\":\"tttt\"}" // api提交的参数 后续可以根据标识查询出来做用户的调用分析，快速协助用户对接
	}
}
```

```json
{
	"action": "apis",
	"data": [{
		"project_id": 5,
		"api_category": "wechat",
		"api_method": "refund",
		"is_success": 1,
		"consume_time": 15,
		"user_from": "payment",
		"user_identify": "20200313001",
		"api_response": "响应内容 可选",
		"api_param": "{\"test\":\"tttt\"}"
	}],
}
```

