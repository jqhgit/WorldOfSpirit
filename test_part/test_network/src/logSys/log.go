package logSys

func init(){
	log_sys.Init(console)
}

type logSystem struct{
	devices [] log_device_interface
}
var log_sys logSystem

func (l *logSystem) Init(logDevType int) {
	switch {
	case logDevType == console:
		l.devices = append(l.devices, new(log_console))
	case logDevType == file:
	case logDevType == database:
	case logDevType == empty:
		return
	default:
		return
	}
}

func LogDebug(a ...interface{}) {
	for _, dev:=range log_sys.devices {
		dev.Debug(a)
	}
}