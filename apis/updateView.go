package apis

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiangyu123/create_view/database"
)

func empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func isNotEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return !(v.Interface() == reflect.Zero(v.Type()).Interface())
}

func create_view(c *gin.Context, view_name string, tablename string) {
	sql_statements := fmt.Sprintf("create view %s as select * from %s", view_name, tablename)
	log.Println(sql_statements)
	success, err := database.SqlDB.Exec(sql_statements)
	if err != nil {
		status_code := 400
		content := gin.H{"msg": fmt.Sprintf("faild create view name %s for %s", view_name, tablename), "code": status_code}
		c.JSON(200, content)
	}

	if isNotEmpty(success) {
		status_code := 200
		content := gin.H{"msg": fmt.Sprintf("success create view name %s for %s", view_name, tablename), "code": status_code}
		c.JSON(200, content)
	}
}

func UpdateView(c *gin.Context) {
	encodeTableName := c.Query("entb")
	if !empty(encodeTableName) {
		create_view(c, "enview", encodeTableName)
	}

	secretTableName := c.Query("sectb")
	if !empty(secretTableName) {
		create_view(c, "secview", secretTableName)
	}
}
