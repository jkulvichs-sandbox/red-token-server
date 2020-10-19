package storage

// ConfStorage - DataBases configuration
type ConfStorage struct {
	SQLite *ConfSQLiteDataBase `yaml:"sqlite,omitempty"`
}

// ConfStorage - Storage configuration
type ConfSQLiteDataBase struct {
	Path   string  `yaml:"file"`
	Cache  *string `yaml:"cache,omitempty"`
	Memory *string `yaml:"memory,omitempty"`
}
