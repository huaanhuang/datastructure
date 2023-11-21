package limiter

// ILimiter 限频接口
type ILimiter interface {
	// Allow 检查是否允许请求
	Allow() bool
}
