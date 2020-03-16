package Model

type Organization struct {
	OName string `json:"o_name"`
	Address string `json:"address"`
	Owner string `json:"owner"`
}

type Iron struct {
	MovieName string `json:"movie_name"`
	MainLead string `json:"main_lead"`
	Director string `json:"director"`
	Budget string   `json:"budget"`
	DBO string		`json:"dbo"`
	OBO string 		`json:"obo"`
	WBO string		`json:"wbo"`
	Des string		`json:"des"`
	Rank string		`json:"rank"`
	Tit string 		`json:"tit"`
	OName string	`json:"o_name"`
}
type MMovie struct {
	Budget string   `json:"budget"`
	Rank string		`json:"rank"`
	MainLead string `json:"main_lead"`
	Director string `json:"director"`
	OName string	`json:"o_name"`
}
type Twodatatwo struct {
	TopTen TopTen
	MMovie MMovie
}
type TopTen struct {
	Top int `json:"top"`
	MovieName string `json:"movie_name"`
	Total string   `json:"total"`
	ReleaseDate string	`json:"release_date"`
}
type Twodata struct {
	Iron Iron
	Detail Detail

}
type Detail struct {
	Pos int			`json:"pos"`
	About string	`json:"about"`
	Credit	string	`json:"credit"`
}

/*func (o *Organization) getOrganization(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM organization").Scan()
}*/
