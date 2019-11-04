package model

type BaseModel struct {
	CreatedAt int `gorm:"default:0" json:"created_at"`
	UpdatedAt int `gorm:"default:0" json:"updated_at"`
}

//type DeletedAt struct {
//	DeletedAt time.Time `gorm:"datetime;index" json:"deleted_at"`
//}

type Application struct {
	Id          int    `gorm:"primary_key,AUTO_INCREMENT" json:"id"`
	StartTime   int    `gorm:"default:0" json:"start_time"`
	Polling     int    `gorm:"default:0" json:"polling"`
	Designation string `gorm:"type:varchar(30);not null" json:"designation" `
	Alias       string `gorm:"type:varchar(20);unique_index;not null" json:"alias" `
	BaseModel
	Article []Article `json:"-" ` // One-To-Many (拥有多个 - article表的ApplicationId作外键)
}
