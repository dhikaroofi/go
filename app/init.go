package app

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"time"
)

// App has router and db instances
type Mojoo struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *Mojoo) Initialize(route *mux.Router) {
	config := initConfig()
	db, err := gorm.Open(mysql.Open(config.getDSN()), &gorm.Config{
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("Could not connect database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Could not connect database")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	a.DB = db
	a.Router = route
	a.setRouters()
	log.Println("app server is running")
}

func (a *Mojoo) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}
func (a *Mojoo) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *Mojoo) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *Mojoo) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
func (a *Mojoo) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *Mojoo) guest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

//func (a *Mojoo) guard(handler RequestHandlerFunction) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		if err := helper.AuthorizeRole(a.DB, r, "admin"); err != nil {
//			helper.RespondJSONError(w, http.StatusUnauthorized, err)
//			return
//		}
//		handler(a.DB, w, r)
//	}
//}
