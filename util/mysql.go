package main

import (
	"github.com/chenshone/tiktok-lite/dal"
	"gorm.io/gen"
	"gorm.io/gen/field"
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
	user := g.GenerateModel("user")
	video := g.GenerateModel("video", gen.FieldRelate(field.BelongsTo, "Author", user, &field.RelateConfig{
		GORMTag: "foreignKey:UserID",
	}))
	favorite := g.GenerateModel("favorite")
	g.ApplyBasic(user, video, favorite)
	g.Execute()
}
