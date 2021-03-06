package models

import "gorm.io/gorm"

type Companies struct {
	gorm.Model
	Name        string `gorm:"not null"`
	LogoURL     string
	Description string `gorm:"not null"`
	Type        string `gorm:"not null"`
	WebsiteURL  string `gorm:"not null"`
	Email       string `gorm:"not null"`
	AddressID uint   `gorm:"not null"`
	Addresses Addresses `gorm:"foreignKey:AddressID;references:id"`
}
