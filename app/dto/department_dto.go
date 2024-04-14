package dto

type DepartmentRequestBody struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type DepartmentFetchRecord struct {
	ID        uint   `gorm:"id"`
	Name      string `gorm:"name"`
	Location  string `gorm:"location"`
	CreatedAt string `gorm:"created_at"`
	UpdatedAt string `gorm:"updated_at"`
	DeletedAt string `gorm:"deleted_at"`
}

type DepartmentResponesData struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type DepartmentRespones struct {
	Message    string      `json:"message"`
	StatusBool bool        `json:"status_bool"`
	Data       interface{} `json:"data,omitempty"`
}
