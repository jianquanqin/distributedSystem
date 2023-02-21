package registry

//创建注册结构体，后续每一个服务都是它的一个实例

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string
	HeartBeatURL     string
}

type ServiceName string

//将有限数量的服务名称计作常量

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portald")
)

//定义一个变量，表示服务

type patchEntry struct {
	Name ServiceName
	URL  string
}

//再定义一个变量，表示服务的增减情况

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
