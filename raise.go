package workers

import (
  "encoding/json"
  "fmt"
  "os"
  "net/http"
)

func Raise(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf-8")

  value := FormValue("test")

  /*Configure(map[string]string {
    "server": os.Getenv("REDIS_URL"),
    "database": os.Getenv("REDIS_DATABASE"),
    "pool": os.Getenv("REDIS_POOL"),
    "process": os.Getenv("REDIS_PROCESS"),
  })

  Enqueue("plan-one", "Add", []int{1, 2})*/

  stats := map[string]interface{}{
    "raised": value,
  }

  body, _ := json.MarshalIndent(stats, "", "  ")
  fmt.Fprintln(w, string(body))
}
