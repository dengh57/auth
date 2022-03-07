package lib

type Model struct {
	ID                int64 `gorm:"primaryKey;autoIncrement"`
	RequestDefinition string
	PolicyDefinition  string
	RoleDefinition    string
	PolicyEffect      string
	Matchers          string
}

type Rule struct {
	ID      int64  `gorm:"primaryKey;autoIncrement"`
	Ptype   string `gorm:"size:100"`
	V0      string `gorm:"size:100"`
	V1      string `gorm:"size:100"`
	V2      string `gorm:"size:100"`
	V3      string `gorm:"size:100"`
	V4      string `gorm:"size:100"`
	V5      string `gorm:"size:100"`
	V6      string `gorm:"size:100"`
	V7      string `gorm:"size:100"`
	V8      string `gorm:"size:100"`
	V9      string `gorm:"size:100"`
	ModelID string `gorm:"size:100"`
}
