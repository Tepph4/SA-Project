package entity

import (

	//"time"
  
	"gorm.io/gorm"
  
  )
    
 type Building struct{
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	// 1 มีได้หลาย ห้อง
	Rooms	[]Room	`gorm:"foreignkey:BuildingID"`
 }  
 
 type Room struct{
	gorm.Model	
	Name string
	// BuildingID ทำหน้าที่เป็น Fk
	BuildingID *uint
	Building Building `gorm:"references:id"`

	Room_has_Devices []Room_has_Device `gorm:"foreignkey:RoomID"`
	
 }

 type User struct{
	gorm.Model
	Name string
	Email string `gorm:"uniqueIndex"`
	Password string `json:"-"`

	RoleID *uint
	Role Role `gorm:"references:id"`
	Room_has_Devices []Room_has_Device `gorm:"foreignKey:UserID"`
 }
 
 type Role struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	User []User `gorm:"foreignkey:RoleID"`
}
 type Device struct{
	gorm.Model	
	Name string  `gorm:"uniqueIndex"`
	Room_has_Devices []Room_has_Device `gorm:"foreignKey:DeviceID"`
 }

 type Room_has_Device struct{
	gorm.Model
	// RoomID ทำหน้าที่เป็น Fk
	RoomID *uint
	Room Room `gorm:"references:id"`
	// DeviceID ทำหน้าที่เป็น Fk
	DeviceID *uint
	Device Device `gorm:"references:id"`
	// UserID ทำหน้าที่เป็น Fk
	UserID *uint
	User User `gorm:"references:id"`
 }