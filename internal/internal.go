package internal

func init() {
	bootLogger()
	bootConfig()
	bootCache()
	bootServer()

	Version = &ServerVersion{
		UIBackendVersion:  backendVersion,
		UIFrontendVersion: frontendVersion,
	}
}

var (
	backendVersion  = "1.0.0"
	frontendVersion = "1.0.0"
)

type ServerVersion struct {
	UIBackendVersion  string `json:"ui_backend_version"`
	UIFrontendVersion string `json:"ui_frontend_version"`
	ServerVersion     string `json:"server_version"`
}

var Version *ServerVersion
