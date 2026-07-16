package dto

type CreateRequest struct {
	Name     string `json:"name" validate:"required" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"password" validate:"required,min=6" gorm:"type:varchar(100);not null"`
}
