package logSys

//logimp type
const (
	empty = iota	// 不记录
	console  		// 控制台
	file			// 文件
	database		// 数据库
)

type log_device_interface interface {
	Debug(a ...interface{})
}
