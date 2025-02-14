package spec

const (
	ServerBinConfigPath        = "/etc/hstream/config.yaml"
	ServerDefaultImage         = "hstreamdb/hstream"
	ServerDefaultContainerName = "deploy_hserver"
	ServerDefaultBinPath       = "/usr/local/bin/hstream-server-sid"
	ServerDefaultCfgDir        = "/hstream/deploy/hserver"
	ServerDefaultDataDir       = "/hstream/data/hserver"
)

type HServerSpec struct {
	Host              string       `yaml:"host"`
	AdvertisedAddress string       `yaml:"advertised_address"`
	Port              int          `yaml:"port" default:"6570"`
	InternalPort      int          `yaml:"internal_port" default:"6571"`
	Image             string       `yaml:"image"`
	SSHPort           int          `yaml:"ssh_port" default:"22"`
	RemoteCfgPath     string       `yaml:"remote_config_path"`
	DataDir           string       `yaml:"data_dir"`
	Opts              ServerOpts   `yaml:"server_config"`
	ContainerCfg      ContainerCfg `yaml:"container_config"`

	// optional, when scaling HServer cluster without given admin node info
	StoreAdminHost string `yaml:"store_admin_host"`
	StoreAdminPort int    `yaml:"store_admin_port"`
}

func (h *HServerSpec) SetDefaultDataDir() {
	h.DataDir = ServerDefaultDataDir
}

func (h *HServerSpec) SetDefaultImage() {
	h.Image = ServerDefaultImage
}

func (h *HServerSpec) SetDefaultRemoteCfgPath() {
	h.RemoteCfgPath = ServerDefaultCfgDir
}

type ServerOpts struct {
	ServerLogLevel string `yaml:"server_log_level" default:"info"`
	StoreLogLevel  string `yaml:"store_log_level" default:"info"`
	Compression    string `yaml:"compression" default:"lz4"`
}
