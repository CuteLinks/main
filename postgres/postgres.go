package postgres

import (
	"cutelinks/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var connStr = config.GetConfig().DB_CONN_STR

func GetUrl(tinyUrl string) string{
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	var fullurl string
	err = db.QueryRow("select fullurl from link where tinyurl=$1",tinyUrl).Scan(&fullurl)
	if err != nil {
		log.Println(err)
	}
	return fullurl
}
func SaveUrl(tinyUrl string, fullUrl string) string{
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	stmt, err := db.Prepare("insert into link values ($1, $2)")
	if err != nil {
		log.Println(err)
	}
	res, err := stmt.Exec(tinyUrl, fullUrl)
	if res!=nil{
		rowCnt, err:=res.RowsAffected()
		if rowCnt==1 {
			return tinyUrl
		}
		if err != nil {
			log.Println(err)
		}
	}
	return ""
}