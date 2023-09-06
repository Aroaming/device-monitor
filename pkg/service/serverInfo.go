package service

// node_exporter指标
const (
	temperatureExplosionQuery = ``
	temperatureSafeQuery      = ``
	cpusQuery                 = `100-(avg by(cpu)(irate(node_cpu_seconds_total{mode="idle"}[5m]))*100)`
	memsQuery                 = `100 *(1-(sum by(instance)(node_memory_MemAvailable_bytes)/sum by(instance)(node_memory_MemTotal_bytes)))`
	disksQuery                = `100*(1-(sum by(device)(node_filesystem_avail_bytes)/sum by(device)(node_filesystem_size_bytes)))`
	npusQuery                 = ``
)

type ServerInfo struct {
	TemperatureExplosion float64   `json:"temperatureExplosion"` //隔爆腔 温度
	TemperatureSafe      float64   `json:"temperatureSafe"`      //本安腔 温度
	CPUs                 []float64 `json:"cpus"`                 //每个cpu使用率
	Mems                 []float64 `json:"mems"`                 //单个内存的使用率
	Disks                []float64 `json:"disks"`                //单个磁盘使用率
	NPUs                 []float64 `json:"npus"`                 //每个npu使用率
	IP                   string    `json:"ip"`                   //ip
}

func NewServerInfo() *ServerInfo {
	return &ServerInfo{}
}

//TODO
