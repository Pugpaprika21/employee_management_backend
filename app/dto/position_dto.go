package dto

type PositionRequestBody struct {
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Salary           float64 `json:"salary"`
	Status           string  `json:"status"`
	Location         string  `json:"location"`
	Responsibilities string  `json:"responsibilities"`
}

type PositionFetchRecord struct {
	ID               uint    `gorm:"id"`
	Name             string  `gorm:"name"`
	Description      string  `gorm:"description"`
	Salary           float64 `gorm:"salary"`
	Status           string  `gorm:"status"`
	Location         string  `gorm:"location"`
	Responsibilities string  `gorm:"responsibilities"`
	CreatedAt        string  `gorm:"created_at"`
	UpdatedAt        string  `gorm:"updated_at"`
	DeletedAt        string  `gorm:"deleted_at"`
}

type PositionResponesData struct {
	ID               uint    `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Salary           float64 `json:"salary"`
	Status           string  `json:"status"`
	Location         string  `json:"location"`
	Responsibilities string  `json:"responsibilities"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	DeletedAt        string  `json:"deleted_at"`
}

type PositionRespones struct {
	Message    string      `json:"message"`
	StatusBool bool        `json:"status_bool"`
	Data       interface{} `json:"data,omitempty"`
}
