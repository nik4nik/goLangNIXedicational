package main
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("mysql", "root:rdfr@/productdb")
    if err != nil {
        log.Fatal(err)
    }
    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
}
func main() {
	posts(jsonplaceholder(7, "posts?userId="))
	fmt.Scanln()
}
func posts(data []map[string]interface{}) {
	for _, e := range data {
		res, err := db.Exec("insert into productdb.posts (user_id, title, body) values (?, ?, ?)", int(e["userId"].(float64)), e["title"], e["body"])
		if err != nil{
			panic(err)
		}
		id, _ := res.LastInsertId()
		go comments(jsonplaceholder(int(id), "comments?postId="))
	}
	fmt.Println("All posts are saved to the database")
}
func comments(data []map[string]interface{}) {
	for _, e := range data {
		_, err := db.Exec("insert into productdb.comments (post_id, name, email, body) values (?, ?, ?, ?)", int(e["postId"].(float64)), e["name"], e["email"], e["body"])
		if err != nil{
			panic(err)
		}
	}
}
func jsonplaceholder(n int, str string)(data []map[string]interface{}) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/" + str + strconv.Itoa(n))
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		panic(err)
	}
	return
}