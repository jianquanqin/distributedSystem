package registry

type Registration struct {
	ServiceName      ServiceName //can use string, but customize type be better
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string
	HeartBeatURL     string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portal")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
