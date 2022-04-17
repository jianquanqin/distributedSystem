package registry

type Registration struct {
	ServiceName ServiceName //can use string, but customize type be better
	ServiceURL  string
}

type ServiceName string

const (
	LogService = ServiceName("LogService")
)
