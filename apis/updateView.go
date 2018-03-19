package apis

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiangyu123/create_view/database"
)

type view_infos struct {
	enview    string
	secview   string
	tncview   string
	entbname  string
	sectbname string
	tnctbname string
}

func prepare() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	prepare()
}

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

func del_real_viewname(viewname string) {
	sql_statement := fmt.Sprintf("drop view %s", viewname)
	log.Println(sql_statement)
	database.SqlDB.Exec(sql_statement)
}

func del_view(c *gin.Context, viewsets view_infos) {
	del_real_viewname(viewsets.enview)
	del_real_viewname(viewsets.secview)
}

func create_view(c *gin.Context, viewnames view_infos) {
	tx, err := database.SqlDB.Begin()
	if err != nil {
		log.Fatal("db translation begin faild")
	}
	var message string
	var err_code int16
	var content gin.H
	var counter int

	sql1 := fmt.Sprintf("create view %s as select * from %s", viewnames.enview, viewnames.entbname)
	sql2 := fmt.Sprintf("create view %s as select * from %s", viewnames.secview, viewnames.sectbname)
	sql3 := fmt.Sprintf("create view %s as select * from %s", viewnames.tncview, viewnames.tnctbname)
	log.Println(sql1)
	log.Println(sql2)
	log.Println(sql3)
	if _, err := tx.Exec(sql1); err != nil {
		message = fmt.Sprintf("faild create view name %s for %s", viewnames.enview, viewnames.entbname)
		err_code = 203
		counter++
		log.Println(message)
	}
	if _, err := tx.Exec(sql2); err != nil {
		message = fmt.Sprintf("faild create view name %s for %s", viewnames.secview, viewnames.sectbname)
		err_code = 204
		counter++
		log.Println(message)
	}
	if _, err := tx.Exec(sql3); err != nil {
		message = fmt.Sprintf("faild create view name %s for %s", viewnames.tncview, viewnames.tnctbname)
		err_code = 205
		counter++
		log.Println(message)
	}

	if err_code == 203 || err_code == 204 || err_code == 205 {
		tx.Rollback()
		if counter == 3 {
			message = fmt.Sprintf("faild create both views named %s, %s and %s", viewnames.enview, viewnames.secview, viewnames.tncview)
		}
		content = gin.H{"msg": message, "code": 400}
	} else if err := tx.Commit(); err != nil {
		status_code := 400
		content = gin.H{"msg": "faild create all view", "code": status_code}
		c.JSON(200, content)
		tx.Rollback()
	} else {
		status_code := 200
		content = gin.H{"msg": "success create view", "code": status_code}
	}
	c.JSON(200, content)
}

func UpdateView(c *gin.Context) {
	encodeTableName := c.Query("entb")
	secretTableName := c.Query("sectb")
	tncTableName := c.Query("tnctb")
	viewinfos := view_infos{"tags_relation_encrypt_view", "tags_company_encrypt_chain_view", "tags_encrypt_classify_view", encodeTableName, secretTableName, tncTableName}
	if !empty(encodeTableName) && !empty(secretTableName) && !empty(tncTableName) {
		del_view(c, viewinfos)
		create_view(c, viewinfos)
	} else {
		status_code := 400
		content := gin.H{"msg": "miss a field", "code": status_code}
		c.JSON(200, content)
	}
}
