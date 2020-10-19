package storage

// ConfStorage - Storage's configuration
type ConfStorage struct {
	SQLite *ConfSQLiteDataBase `yaml:"sqlite,omitempty"`
}

// ConfSQLiteDataBase - SQLite configuration
type ConfSQLiteDataBase struct {
	Path   string  `yaml:"file"`
	Cache  *string `yaml:"cache,omitempty"`
	Memory *string `yaml:"memory,omitempty"`
}
