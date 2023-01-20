package main

import (
	"github.com/chenshone/tiktok-lite/dal"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../dal/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
		//FieldNullable: true,
	})

	err := dal.InitDB()
	if err != nil {
		panic(err)
	}
	g.UseDB(dal.DB)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
