package entity

import (
	//"fmt"
	//"time"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	datebase, err := gorm.Open(sqlite.Open("SA-Project.db"))
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	datebase.AutoMigrate(
		&Building{},
		&Room{},
		&Device{},
		&User{},
		&Role{},
		&Room_has_Device{},
	)
	db = datebase

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	db.Model(&Role{}).Create(&Role{Name: "User"})
	db.Model(&Role{}).Create(&Role{Name: "Tech"})
	db.Model(&Role{}).Create(&Role{Name: "Admin"})
	
	var r_user, r_tech, r_admin Role
	db.Raw("SELECT * FROM roles WHERE name = ?", "User").Scan(&r_user)
	db.Raw("SELECT * FROM roles WHERE name = ?", "Tech").Scan(&r_tech)
	db.Raw("SELECT * FROM roles WHERE name = ?", "Admin").Scan(&r_admin)

	db.Model(&User{}).Create(&User{
		Name:     "Teppharit",
		Email:    "Teppharitn@gmail.com",
		Password: string(password),
		Role: r_user,
	})
	db.Model(&User{}).Create(&User{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(password),
		Role: r_admin,
		
	})

	var teppharit User
	var name User
	db.Raw("SELECT * FROM users WHERE email = ?", "Teppharit@gmail.com").Scan(&teppharit)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)

	// Building data
	// F01
	db.Model(&Building{}).Create(&Building{
		Name: "F01",
	})
	db.Model(&Building{}).Create(&Building{
		Name: "F02",
	})
	db.Model(&Building{}).Create(&Building{
		Name: "F03",
	})
	db.Model(&Building{}).Create(&Building{
		Name: "F04",
	})
	var building1, building2, building3, building4 Building
	db.Raw("SELECT * FROM buildings WHERE name = ?", "F01").Scan(&building1)
	db.Raw("SELECT * FROM buildings WHERE name = ?", "F02").Scan(&building2)
	db.Raw("SELECT * FROM buildings WHERE name = ?", "F03").Scan(&building3)
	db.Raw("SELECT * FROM buildings WHERE name = ?", "F04").Scan(&building4)

	// Room data
	db.Model(&Room{}).Create(&Room{
		Name:     "B101",
		Building: building1,
	})
	db.Model(&Room{}).Create(&Room{
		Name:     "B102",
		Building: building1,
	})
	// F02
	db.Model(&Room{}).Create(&Room{
		Name:     "B201",
		Building: building2,
	})
	db.Model(&Room{}).Create(&Room{
		Name:     "B202",
		Building: building2,
	})
	// F03
	db.Model(&Room{}).Create(&Room{
		Name:     "B302",
		Building: building3,
	})
	db.Model(&Room{}).Create(&Room{
		Name:     "B302",
		Building: building3,
	})
	// F04
	db.Model(&Room{}).Create(&Room{
		Name:     "B401",
		Building: building4,
	})
	db.Model(&Room{}).Create(&Room{
		Name:     "B402",
		Building: building4,
	})

	var RoomB101, RoomB102 Room
	var RoomB201, RoomB202 Room
	var RoomB301, RoomB302 Room
	var RoomB401, RoomB402 Room
	db.Raw("SELECT * FROM rooms WHERE name = ?", "B101").Scan(&RoomB101)
	db.Raw("SELECT * FROM rooms WHERE name = ?", "B102").Scan(&RoomB102)

	db.Raw("SELECT * FROM rooms WHERE name = ?", "B201").Scan(&RoomB201)
	db.Raw("SELECT * FROM rooms WHERE name = ?", "B202").Scan(&RoomB202)

	db.Raw("SELECT * FROM rooms WHERE name = ?", "B301").Scan(&RoomB301)
	db.Raw("SELECT * FROM rooms WHERE name = ?", "B302").Scan(&RoomB302)

	db.Raw("SELECT * FROM rooms WHERE name = ?", "B401").Scan(&RoomB401)
	db.Raw("SELECT * FROM rooms WHERE name = ?", "B402").Scan(&RoomB402)

	//Device data
	db.Model(&Device{}).Create(&Device{
		Name: "com1",
	})
	db.Model(&Device{}).Create(&Device{
		Name: "com2",
	})
	db.Model(&Device{}).Create(&Device{
		Name: "com3",
	})
	db.Model(&Device{}).Create(&Device{
		Name: "com4",
	})
	var D1, D2, D3, D4 Device
	db.Raw("SELECT * FROM buildings WHERE name = ?", "com1").Scan(&D1)
	db.Raw("SELECT * FROM buildings WHERE name = ?", "com2").Scan(&D2)
	db.Raw("SELECT * FROM buildings WHERE name = ?", "com3").Scan(&D3)
	db.Raw("SELECT * FROM buildings WHERE name = ?", "com4").Scan(&D4)

}
