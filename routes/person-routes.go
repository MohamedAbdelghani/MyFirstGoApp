package routes
import (
	"github.com/mohamedabdelghani/myfirstgogrudapp/controllers"
	"github.com/gorilla/mux"
)

var RegisterPersonRoutes = func(router *mux.Router) {
	router.HandleFunc("/AddPerson/", controllers.Add).Methods("POST")
	router.HandleFunc("/GetById?{id}", controllers.GetById).Methods("GET")
  router.HandleFunc("/GetByWeight?{weight}", controllers.GetByWeight).Methods("GET")
	router.HandleFunc("/GetByHeight?{height}", controllers.GetByHeight).Methods("GET")
	router.HandleFunc("/EditPerson/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/DeletePerson/{id}", controllers.Delete).Methods("DELETE")
}