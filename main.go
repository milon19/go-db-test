package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type ResPartner struct {
	bun.BaseModel `bun:"select:sb_partner"`
	ID            int64  `bun:"id"`
	Name          string `bun:"name"`
}

type Order struct {
	bun.BaseModel    `bun:"select:sb_order"`
	ID               int64       `bun:"id"`
	GenOid           string      `bun:"gen_oid"`
	ContactPartner   *ResPartner `bun:"rel:belongs-to"`
	ContactPartnerID int64       `bun:"contact_partner"`
}

func main() {
	r := gin.Default()

	sqldb, err := sql.Open(sqliteshim.ShimName, "file:db.sqlite3")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	ctx := context.Background()
	order := make([]Order, 0)
	err = db.NewSelect().Model(&order).
		Relation("ContactPartner").
		Scan(ctx)

	if err != nil {
		fmt.Println(err)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": order,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
