package db

type Config struct {
	Protocol      string `json:"protocol,omitempty"`
	ConnectionURL string `json:"connetionURL"`
	Port          string `json:"port"`
	User          string `json:"user,omitempty"`
	Password      string `json:"password,omitempty"`
	DBName        string `json:"dbName"`
}

type Database struct {
	Priority         int    `json:"priority"`
	TotalSpaceGB     int    `json:"total_space_GB"`
	AvailableSpaceGB int    `json:"available_space_GB"`
	DBProvider       string `json:"db_provider"`
	Config           Config `json:"config"`
}

type Data struct {
	Project  string     `json:"project"`
	Database []Database `json:"database"`
}
