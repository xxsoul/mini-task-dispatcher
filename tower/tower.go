package tower

import (
	"database/sql"
	"flag"
	"net/http"
	"task-airport/tower/protocols/response"
	"task-airport/tower/tools"

	_ "github.com/go-sql-driver/mysql"
)

var sqlConn *sql.DB // 数据库链接

func init() {
	dbConStr := ""
	flag.StringVar(&dbConStr, "-db", "none", "sql connect string")
	flag.Parse()

	http.HandleFunc("", handleGetTask)
	sqlConn, _ = sql.Open("mysql", dbConStr)
}

// handleGetTask 获取任务信息
func handleGetTask(rw http.ResponseWriter, req *http.Request) {
	if !tools.CheckHttpMethod(req, http.MethodGet) {
		tools.WriteHttpJsonResponse(&rw, response.NewResponse[any]("METHOD_NOT_ALLOWED", "http方法不合法", nil))
		return
	}

	urlQuery := req.URL.Query()
	seqNo := req.URL.Query().Get("seqNo")
}
