package main
import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}
func insert(t Todo) (int, error) {
	db, err := sql.Open("postgres","postgres://ornvuxai:9GFXmiWs7r29tjJvN9HfkRsFfaigC47v@suleiman.db.elephantsql.com:5432/ornvuxai")
	if err != nil {
		log.Fatal("Connect to database error",err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1,$2) RETURNING id",t.Title,t.Status)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0,err
	}
	return id,nil
}
func query(filter string) (Todo,error) {
	db, err := sql.Open("postgres","postgres://ornvuxai:9GFXmiWs7r29tjJvN9HfkRsFfaigC47v@suleiman.db.elephantsql.com:5432/ornvuxai")
	if err != nil {
		log.Fatal("Connect to database error",err)
	}
	defer db.Close()

		db.Prepare("SELECT id,title,status FROM todos WHERE status=$1")
		if err != nil {
			log.Fatal("can't prepare statement", err)
		}
		row := stmt.QueryRow(filter)
		var id int
		var title, status string
		err = row.Scan(&id,&title,&status)
		if err != nil {
			log.Fatal("can't scan row into variable", err)
			return nil,err
		}
		return Todo{id,title,status},nil
	
}
func queryAll(filter string) (Todo,error) {
	db, err := sql.Open("postgres","postgres://ornvuxai:9GFXmiWs7r29tjJvN9HfkRsFfaigC47v@suleiman.db.elephantsql.com:5432/ornvuxai")
	if err != nil {
		log.Fatal("Connect to database error",err)
	}
	defer db.Close()
	if filter != "" {
		db.Prepare("SELECT id,title,status FROM todos WHERE status=$1")
		if err != nil {
			log.Fatal("can't prepare statement", err)
		}
		row := stmt.QueryRow(filter)
		var id int
		var title, status string
		err = row.Scan(&id,&title,&status)
		if err != nil {
			log.Fatal("can't scan row into variable", err)
		}
		return Todo{id,title,status},nil
	}else{
		db.Prepare("SELECT id,title,status FROM todos")
		if err != nil {
			log.Fatal("can't prepare statement", err)
		}
		rows := stmt.Query()
	}
	
}
func postTodoHandler(c *gin.Context)  {
	t := Todo{}
	err := c.ShouldBindJSON(&t)
	if(err!=nil){
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}
	id,err := insert(t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}
	t.ID = id
	c.JSON(http.StatusCreated,t)

}
func getTodoHandler(c *gin.Context)  {
	status, ok := c.GetQuery("status")
	if ok != nil {
		rows, err := query("")
	}
	else{
		rows, err := query(status)
	}
}
func main() {
	r := gin.Default()
	r.GET("/todos",getTodoHandler)
	r.POST("/todos",postTodoHandler)
	// r.GET("/todos/:id",getByIdHandler)
	// r.PUT("/todos/:id",updateByIdHandler)
	r.Run(":1234")
}