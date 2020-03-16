package app

import (
	"Superhero/Handler"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	DB1     *sql.DB
}
/*func (a *App) Initial(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	_, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	//a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.Routes()
}*/

func (a *App) Initialize() {
	//connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	var err error
	a.DB, err = sql.Open ("mysql", "root:jamespdh@tcp(127.0.0.1:3306)/superheromovie")
	a.DB1, err =  sql.Open ("mysql", "root:jamespdh@tcp(127.0.0.1:3306)/toptenlist")

	//a.DB , err = sql.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	//a.DB = .DBMigrate(db)
	a.Router = mux.NewRouter()
	a.Routes()
}
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
func (a *App) Routes() {
	a.Router.HandleFunc("/Movie", a.SelectMov).Methods("GET")
	a.Router.HandleFunc("/", a.SelectData).Methods("GET")
	a.Router.HandleFunc("/Top", a.SelectTop).Methods("GET")
	a.Router.HandleFunc("/Top/Marvel", a.SelectMar).Methods("GET")
	a.Router.HandleFunc("/Movie/{movie_name}", a.UpdateMov).Methods("PUT")
	a.Router.HandleFunc("/Movie/{movie_name}", a.DeleteMov).Methods("DELETE")
	a.Router.HandleFunc("/Org", a.InsertOrganization).Methods("POST")
	a.Router.HandleFunc("/Movie", a.InsertMovieDetail).Methods("POST")
	/*a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")*/
}

func (a *App) InsertOrganization(w http.ResponseWriter, r *http.Request)  {
	Handler.CreateOrganization(a.DB, w, r)
}
func (a *App) InsertMovieDetail(w http.ResponseWriter, r *http.Request)  {
	Handler.CreatePost(a.DB, w, r)
}
func (a *App) SelectData(w http.ResponseWriter, r *http.Request)  {
	Handler.SelectAll(a.DB,w, r)
}
func (a *App) SelectTop(w http.ResponseWriter, r *http.Request)  {
	Handler.Selecttopten(a.DB1,w, r)
}
func (a *App) SelectMar(w http.ResponseWriter, r *http.Request)  {
	Handler.SelectMarvel(a.DB1 ,w, r)
}
func (a *App) SelectMov(w http.ResponseWriter, r *http.Request)  {
	Handler.SelectMovie(a.DB ,w, r)
}
func (a *App) UpdateMov(w http.ResponseWriter, r *http.Request)  {
	Handler.UpdateMovie(a.DB ,w, r)
}
func (a *App) DeleteMov(w http.ResponseWriter, r *http.Request)  {
	Handler.DeleteMovie(a.DB ,w, r)
}