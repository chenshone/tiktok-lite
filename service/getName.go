package service

import (
	"context"
	"fmt"
	"github.com/chenshone/tiktok-lite/dal/query"
)

var q = query.Q

func GetName(id int, ctx context.Context) {
	t := q.Person
	do := t.WithContext(context.Background())
	data, err := do.Where(t.ID.Eq(int32(id))).Find()
	if err != nil {
		panic(err)
	}
	fmt.Println(data[0].Name)
}
