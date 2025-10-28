package mirrornode

type Config struct {
	BASE_URL string
	Network  string
}

var configInstance *Config = nil

func GetConfig() *Config {
	return configInstance
}

func InititaLizeMirronode(baseurl string, network string) {
	configInstance = &Config{
		BASE_URL: baseurl,
		Network:  network,
	}
}
