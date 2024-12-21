package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// https://gorm.io/gen/index.html
func main() {
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath:  "./internal/repository/po",
		OutPath:       "./internal/repository/query",
		OutFile:       "",
		FieldNullable: true,
	})

	gormdb, err := gorm.Open(mysql.Open("root:mysql@(localhost)/easybills?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	g.UseDB(gormdb) // reuse your gorm gormgen
	g.WithOpts(gen.FieldIgnore("create_at", "update_at"))

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(
		g.GenerateModel("user"),
	)

	// Generate the code
	g.Execute()

}
