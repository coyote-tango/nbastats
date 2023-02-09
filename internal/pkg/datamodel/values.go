package datamodel

type TeamInfo struct {
	Name string
	Id   int
}

const NBA = "00"
const ABA = "01"

var teams = map[string]TeamInfo{
	"ATL": {Name: "Atlanta Hawks", Id: 1610612737},
	"BOS": {Name: "Boston Celtics", Id: 1610612738},
	"BKN": {Name: "Brooklyn Nets", Id: 1610612751},
	"CHA": {Name: "Charlotte Hornets", Id: 1610612766},
	"CHI": {Name: "Chicago Bulls", Id: 1610612741},
	"CLE": {Name: "Cleveland Cavaliers", Id: 1610612739},
	"DAL": {Name: "Dallas Mavericks", Id: 1610612742},
	"DEN": {Name: "Denver Nuggets", Id: 1610612743},
	"DET": {Name: "Detroit Pistons", Id: 1610612765},
	"GSW": {Name: "Golden State Warriors", Id: 1610612744},
	"HOU": {Name: "Houston Rockets", Id: 1610612745},
	"IND": {Name: "Indiana Pacers", Id: 1610612754},
	"LAC": {Name: "Los Angeles Clippers", Id: 1610612746},
	"LAL": {Name: "Los Angeles Lakers", Id: 1610612747},
	"MEM": {Name: "Memphis Grizzlies", Id: 1610612763},
	"MIA": {Name: "Miami Heat", Id: 1610612748},
	"MIL": {Name: "Milwaukee Bucks", Id: 1610612749},
	"MIN": {Name: "Minnesota Timberwolves", Id: 1610612750},
	"NOP": {Name: "New Orleans Pelicans", Id: 1610612740},
	"NYK": {Name: "New York Knicks", Id: 1610612752},
	"OKC": {Name: "Oklahoma City Thunder", Id: 1610612760},
	"ORL": {Name: "Orlando Magic", Id: 1610612753},
	"PHI": {Name: "Philadelphia 76ers", Id: 1610612755},
	"PHX": {Name: "Phoenix Suns", Id: 1610612756},
	"POR": {Name: "Portland Trail Blazers", Id: 1610612757},
	"SAC": {Name: "Sacramento Kings", Id: 1610612758},
	"SAS": {Name: "San Antonio Spurs", Id: 1610612759},
	"TOR": {Name: "Toronto Raptors", Id: 1610612761},
	"UTA": {Name: "Utah Jazz", Id: 1610612762},
	"WAS": {Name: "Washington Wizards", Id: 1610612764},
}
