
package main
import (
"log"
"net/http"
"github.com/mohamedabdelghani/myfirstgoapp/routes"
"github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()
  routes.RegisterPersonRoutes(router)
  http.Handle("/", router)
  log.Fatal(http.ListenAndServe("localhost:8080", router))
}


