package Handler

import (
	"Superhero/Model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

func Selecttopten (db *sql.DB, w http.ResponseWriter, r *http.Request){
	//tag := Model.TopTen{}
	data := []Model.TopTen{}

	//Select S.*, T.position, T.about , T.credit_review from superheromovie.movie S Inner Join topcollection.tcmarvelmovie T
	//On S.movie_name = T.movie_name

	sel, err := db.Query(" Select * FROM top10 ORDER By top10.top")
	if err != nil{
		panic(err)
	}

	for sel.Next() {
		var tag Model.TopTen
		// for each row, scan the result into our tag composite object
		err = sel.Scan(&tag.Top, &tag.MovieName, &tag.Total, &tag.ReleaseDate)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data = append(data, tag)
		//fmt.Println(b )

		// and then print out the tag's Name attribute

	}
	for a := range data{
		fmt.Println(data[a])
	}
	//sel.Scan(&data)

	respondJSON(w, http.StatusOK, data)

}

func SelectMarvel (db *sql.DB, w http.ResponseWriter, r *http.Request){
	tag := Model.Twodatatwo{}
	data := []Model.Twodatatwo{}

	//Select S.*, T.position, T.about , T.credit_review from superheromovie.movie S Inner Join topcollection.tcmarvelmovie T
	//On S.movie_name = T.movie_name

	sel, err := db.Query(" Select S.*, T.budget , T.ranking , T.main_lead, T.director , T.oname from toptenlist.top10 S " +
		" Inner Join superheromovie.movie T " +
		"On S.movie_name = T.movie_name ORDER BY S.top asc" )
	if err != nil{
		panic(err)
	}

	for sel.Next() {
		//	var tag Model.Twodata
		// for each row, scan the result into our tag composite object
		err = sel.Scan(&tag.TopTen.Top, &tag.TopTen.MovieName, &tag.TopTen.Total, &tag.TopTen.ReleaseDate,&tag.MMovie.Budget,&tag.MMovie.Rank,&tag.MMovie.MainLead,&tag.MMovie.Director,&tag.MMovie.OName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data = append(data, tag)
		//fmt.Println(b )

		// and then print out the tag's Name attribute


		//log.Print(tag.Iron.MovieName,tag.Iron.MainLead,tag.Iron.Director,tag.Iron.Budget,tag.Iron.DBO,tag.Iron.OBO,tag.Iron.WBO,tag.Iron.Des,tag.Iron.Rank,tag.Iron.Tit,tag.Iron.OName,tag.Detail.Pos,tag.Detail.About,tag.Detail.Credit)
	}

	for a := range data{
		fmt.Println(data[a])
	}
	//sel.Scan(&data)

	respondJSON(w, http.StatusOK, data)

}
func SelectMovie (db *sql.DB, w http.ResponseWriter, r *http.Request){

	data := []Model.Iron{}
	mov := Model.Iron{}

	//Select S.*, T.position, T.about , T.credit_review from superheromovie.movie S Inner Join topcollection.tcmarvelmovie T
	//On S.movie_name = T.movie_name

	sel, err := db.Query(" Select * from superheromovie.movie  " )
	if err != nil{
		panic(err)
	}

	for sel.Next() {
		//	var tag Model.Twodata
		// for each row, scan the result into our tag composite object
		err = sel.Scan(&mov.MovieName,&mov.MainLead,&mov.Director,
			&mov.Budget,&mov.DBO,&mov.OBO,&mov.WBO,&mov.Des,&mov.Rank,&mov.Tit,&mov.OName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data = append(data, mov)
		//fmt.Println(b )

		// and then print out the tag's Name attribute


		//log.Print(tag.Iron.MovieName,tag.Iron.MainLead,tag.Iron.Director,tag.Iron.Budget,tag.Iron.DBO,tag.Iron.OBO,tag.Iron.WBO,tag.Iron.Des,tag.Iron.Rank,tag.Iron.Tit,tag.Iron.OName,tag.Detail.Pos,tag.Detail.About,tag.Detail.Credit)
	}

	/*for a := range data{
		fmt.Println(data[a])
	}*/
	//sel.Scan(&data)

	respondJSON(w, http.StatusOK, data)

}
func SelectAll (db *sql.DB, w http.ResponseWriter, r *http.Request){
	tag := Model.Twodata{}
	data := []Model.Twodata{}

	//Select S.*, T.position, T.about , T.credit_review from superheromovie.movie S Inner Join topcollection.tcmarvelmovie T
	//On S.movie_name = T.movie_name"
	sel, err := db.Query(" Select S.*, T.position, T.about , T.credit_review from superheromovie.movie S " +
		" Inner Join topcollection.tcmarvelmovie T " +
		"On S.movie_name = T.movie_name ORDER BY T.position asc" )
	if err != nil{
		panic(err)
	}

	for sel.Next() {
	//	var tag Model.Twodata
		// for each row, scan the result into our tag composite object
		err = sel.Scan(&tag.Iron.MovieName,&tag.Iron.MainLead, &tag.Iron.Director,&tag.Iron.Budget,&tag.Iron.DBO,&tag.Iron.OBO,&tag.Iron.WBO,&tag.Iron.Des,&tag.Iron.Rank,&tag.Iron.Tit,&tag.Iron.OName,&tag.Detail.Pos,&tag.Detail.About,&tag.Detail.Credit)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data = append(data, tag)
		//fmt.Println(b )

		// and then print out the tag's Name attribute


		log.Print(tag.Iron.MovieName,tag.Iron.MainLead,tag.Iron.Director,tag.Iron.Budget,tag.Iron.DBO,tag.Iron.OBO,tag.Iron.WBO,tag.Iron.Des,tag.Iron.Rank,tag.Iron.Tit,tag.Iron.OName,tag.Detail.Pos,tag.Detail.About,tag.Detail.Credit)
	}

	for a := range data{
		fmt.Println(data[a])
	}
	//sel.Scan(&data)

	respondJSON(w, http.StatusOK, data)

}
func CreateOrganization(db *sql.DB, w http.ResponseWriter, r *http.Request)  {
	org := Model.Organization{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&org); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	/*defer db.Close()*/
	insert, err :=db.Query("insert into organization values ('"+ org.OName + "', '"+ org.Address +"', '"+ org.Owner +"')")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	//db.Query("insert into organization values ('Marvel Studio', BurbankWalt Disney Studios')")
	respondJSON(w, http.StatusCreated, org)

}
func CreatePost(db *sql.DB, w http.ResponseWriter, r *http.Request)  {
	org := Model.Iron{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&org); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	//defer db.Close()
	insert, err :=db.Query(" insert into movie  values ('"+ org.MovieName + "', '"+ org.MainLead +"', '" + org.Director + "' , '"+ org.Budget +"', '"+ org.DBO + "', '"+ org.OBO +"', '"+ org.WBO +"', '"+ org.Des + "', '"+ org.Rank +"' , '" + org.Tit + "' , '" + org.OName + "'  ) ")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	//db.Query("insert into organization values ('Marvel Studio', BurbankWalt Disney Studios')")
	respondJSON(w, http.StatusCreated, org)

}

func UpdateMovie(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")*/
	 params := mux.Vars(r)
  stmt, err := db.Prepare("UPDATE movie SET main_lead = ? WHERE movie_name = ?")
  if err != nil {
    panic(err.Error())
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)
  newTitle := keyVal["main_lead"]
  _, err = stmt.Exec(newTitle, params["movie_name"])
  if err != nil {
    panic(err.Error())
  }
  respondJSON(w, http.StatusCreated, newTitle)
}
func DeleteMovie(db *sql.DB, w http.ResponseWriter, r *http.Request) {
 /* w.Header().Set("Content-Type", "application/json")*/
  params := mux.Vars(r)
  stmt, err := db.Prepare("DELETE FROM movie WHERE movie_name = ?")
  if err != nil {
    panic(err.Error())
  }
  _, err = stmt.Exec(params["movie_name"])
  if err != nil {
    panic(err.Error())
  }
  respondJSON(w, http.StatusCreated, params)
}