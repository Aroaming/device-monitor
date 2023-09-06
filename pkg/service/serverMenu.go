package service

import (
	"fmt"
	"log"
	"net"
	"zmos-underground-monitor/pkg/prometheus"
)

// node_exporter指标
const (
	temperatureQuery     = `node_thermal_zone_temp{type="center-thermal",zone="4"}`                                                                   //                 node_thermal_zone_temp{type="center-thermal",zone="4"} 42.538
	memoryUsageQuery     = `100*(1-(node_memory_MemAvailable_bytes/node_memory_MemTotal_bytes))`                                                      //内存指标使用率    100 * (1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes))
	diskUsageQuery       = `(1-node_filesystem_avail_bytes{fstype=~"ext4|xfs"}/node_filesystem_size_bytes{fstype=~"ext4|xfs"})*100`                   //硬盘使用率       (1 - node_filesystem_avail_bytes{fstype=~"ext4|xfs"} / node_filesystem_size_bytes{fstype=~"ext4|xfs"}) * 100
	cpuUsageQuery        = `(1-sum(rate(node_cpu_seconds_total{mode="idle"}[1m]))by(instance)/sum(rate(node_cpu_seconds_total[1m]))by(instance))*100` //cpu使用率       (1 - sum(rate(node_cpu_seconds_total{mode="idle"}[1m])) by (instance) / sum(rate(node_cpu_seconds_total[1m])) by (instance) ) * 100
	npuUsageQuery        = ""                                                                                                                         //npu使用率
	gpuUsageQuery        = ""                                                                                                                         //gpu使用率
	networkUploadQuery   = `rate(node_network_transmit_bytes_total{device="eth0"}[1m])`                                                               //网络上行速率 rate(node_network_transmit_bytes_total{device="eth0"}[1m])
	networkDownloadQuery = `rate(node_network_receive_bytes_total{device="eth0"}[1m])`                                                                //网络下行速率 rate(node_network_receive_bytes_total{device="eth0"}[1m])

	IP = "" //IP
)

type ServerMenu struct {
	Temperature     float64 `json:"temperature"`     //温度
	MemoryUsage     float64 `json:"memoryUsage"`     //内存使用率
	DiskUsage       float64 `json:"diskUsage"`       //硬盘使用率
	CpuUsage        float64 `json:"cpuUsage"`        //cpu使用率
	NpuUsage        float64 `json:"npuUsage"`        //npu使用率
	GpuUsage        float64 `json:"gpuUsage"`        //gpu使用率
	NetworkUpload   float64 `json:"networkUpload"`   //网络上行速率
	NetworkDownload float64 `json:"networkDownload"` //网络下行速率
	IP              string  `json:"ip"`              // ip
}

func NewServerMenu() *ServerMenu {
	return &ServerMenu{}
}

func (s *ServerMenu) SetData() {
	s.setTemperature()
	s.setMemoryUsage()
	s.setDiskUsage()
	s.setCpuUsage()
	s.setGpuUsage()
	s.setNpuUsage()
	s.setNetworkUsage()
	s.setLocalIPByName("eth0")
}

func (s *ServerMenu) setTemperature() {
	value, err := prometheus.GetPoint(temperatureQuery)
	if err != nil {
		log.Printf("获取温度指标错误：%v\n", err)
		return
	}
	s.Temperature = value
}

func (s *ServerMenu) setMemoryUsage() {
	value, err := prometheus.GetPoint(memoryUsageQuery)
	if err != nil {
		log.Printf("获取内存使用率指标错误：%v\n", err)
		return
	}
	s.MemoryUsage = value
}

func (s *ServerMenu) setDiskUsage() {
	value, err := prometheus.GetPoint(diskUsageQuery)
	if err != nil {
		log.Printf("获取硬盘使用率指标错误：%v\n", err)
		return
	}
	s.DiskUsage = value
}

func (s *ServerMenu) setCpuUsage() {
	value, err := prometheus.GetPoint(cpuUsageQuery)
	if err != nil {
		log.Printf("获取CPU使用率指标错误：%v\n", err)
		return
	}
	s.CpuUsage = value
}

func (s *ServerMenu) setNpuUsage() {
	value, err := prometheus.GetPoint(npuUsageQuery)
	if err != nil {
		log.Printf("获取NPU使用率指标错误：%v\n", err)
		return
	}
	s.NpuUsage = value
}

func (s *ServerMenu) setGpuUsage() {
	value, err := prometheus.GetPoint(gpuUsageQuery)
	if err != nil {
		log.Printf("获取GPU使用率指标错误：%v\n", err)
		return
	}
	s.GpuUsage = value
}

func (s *ServerMenu) setNetworkUsage() {
	value, err := prometheus.GetPoint(networkDownloadQuery)
	if err != nil {
		log.Printf("获取网络上传速率指标错误：%v\n", err)
		return
	}
	s.NetworkDownload = value

	value, err = prometheus.GetPoint(networkUploadQuery)
	if err != nil {
		log.Printf("获取网络下载速率指标错误：%v\n", err)
		return
	}
	s.NetworkUpload = value
}

// 先获取eth0网卡IP
func (s *ServerMenu) setLocalIPByName(name string) {
	inter, err := net.InterfaceByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	addrs, err := inter.Addrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				s.IP = ipnet.IP.String()
			}
		}
	}
}
