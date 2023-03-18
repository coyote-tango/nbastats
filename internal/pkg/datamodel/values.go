package datamodel

type TeamInfo struct {
	Name string
	Id   int
	City string
}

const NBA = "00"
const ABA = "01"
const CURRENT = "2022-23"

var teams = map[string]TeamInfo{
	"ATL": {Id: 1610612737, City: "Atlanta", Name: "Hawks"},
	"BOS": {Id: 1610612738, City: "Boston", Name: "Celtics"},
	"BKN": {Id: 1610612751, City: "Brooklyn", Name: "Nets"},
	"CHA": {Id: 1610612766, City: "Charlotte", Name: "Hornets"},
	"CHI": {Id: 1610612741, City: "Chicago", Name: "Bulls"},
	"CLE": {Id: 1610612739, City: "Cleveland", Name: "Cavaliers"},
	"DAL": {Id: 1610612742, City: "Dallas", Name: "Mavericks"},
	"DEN": {Id: 1610612743, City: "Denver", Name: "Nuggets"},
	"DET": {Id: 1610612765, City: "Detroit", Name: "Pistons"},
	"GSW": {Id: 1610612744, City: "Golden State", Name: "Warriors"},
	"HOU": {Id: 1610612745, City: "Houston", Name: "Rockets"},
	"IND": {Id: 1610612754, City: "Indiana", Name: "Pacers"},
	"LAC": {Id: 1610612746, City: "Los Angeles", Name: "Clippers"},
	"LAL": {Id: 1610612747, City: "Los Angeles", Name: "Lakers"},
	"MEM": {Id: 1610612763, City: "Memphis", Name: "Grizzlies"},
	"MIA": {Id: 1610612748, City: "Miami", Name: "Heat"},
	"MIL": {Id: 1610612749, City: "Milwaukee", Name: "Bucks"},
	"MIN": {Id: 1610612750, City: "Minnesota", Name: "Timberwolves"},
	"NOP": {Id: 1610612740, City: "New Orleans", Name: "Pelicans"},
	"NYK": {Id: 1610612752, City: "New York", Name: "Knicks"},
	"OKC": {Id: 1610612760, City: "Oklahoma City", Name: "Thunder"},
	"ORL": {Id: 1610612753, City: "Orlando", Name: "Magic"},
	"PHI": {Id: 1610612755, City: "Philadelphia", Name: "76ers"},
	"PHX": {Id: 1610612756, City: "Phoenix", Name: "Suns"},
	"POR": {Id: 1610612757, City: "Portland", Name: "Trail Blazers"},
	"SAC": {Id: 1610612758, City: "Sacramento", Name: "Kings"},
	"SAS": {Id: 1610612759, City: "San Antonio", Name: "Spurs"},
	"TOR": {Id: 1610612761, City: "Toronto", Name: "Raptors"},
	"UTA": {Id: 1610612762, City: "Utah", Name: "Jazz"},
	"WAS": {Id: 1610612764, City: "Washington", Name: "Wizards"},
}
