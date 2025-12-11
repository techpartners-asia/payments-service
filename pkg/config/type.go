package configPkg

type (
	Config struct {
		App   App   `json:"app"`
		DB    DB    `json:"db"`
		Redis Redis `json:"redis"`
	}
	App struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Port    string `json:"port"`
		Env     string `json:"env"`
	}
	DB struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Timezone string `json:"timezone"`
	}
	Redis struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	}
)
