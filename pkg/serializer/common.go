// serializer 基础序列化器
package serializer

// Response 是公共的响应结构体
type Response struct {
	// 响应状态码
	Status int `json:"status"`
	// 返回数据
	Data interface{} `json:"data"`
	// 返回的消息
	Message string `json:"message"`
	// 返回的错误
	Error string `json:"error"`
}

// TokenData 是登录响应数据的结构体
type TokenData struct {
	User  interface{} `json:"user"`
	Token string
}
