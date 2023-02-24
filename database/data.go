package database

// For string a cricket players basic information
type Player struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Country string  `json:"country"`
	Batting Batting `json:"batting"`
	Bowling Bowling `json:"bowling"`
}

// For a players batting information
type Batting struct {
	Innings int `json:"innings"`
	Runs    int `json:"runs"`
}

// For a players bowling information
type Bowling struct {
	Innings int `json:"innings"`
	Wickets int `json:"wickets"`
}

var Data = []Player{
	{
		ID:      "1",
		Name:    "Babar Azam",
		Age:     28,
		Country: "Pakistan",
		Batting: Batting{
			Innings: 160,
			Runs:    3696,
		},
		Bowling: Bowling{
			Innings: 155,
			Wickets: 2,
		},
	},
	{
		ID:      "2",
		Name:    "Sakib Al Hasan",
		Age:     30,
		Country: "Bangladesh",
		Batting: Batting{
			Innings: 165,
			Runs:    1445,
		},
		Bowling: Bowling{
			Innings: 156,
			Wickets: 500,
		},
	},
	{
		ID:      "3",
		Name:    "David Warner",
		Age:     35,
		Country: "Australia",
		Batting: Batting{
			Innings: 170,
			Runs:    34000,
		},
		Bowling: Bowling{
			Innings: 15,
			Wickets: 25,
		},
	},
	{
		ID:      "4",
		Name:    "Tamim Iqbal",
		Age:     33,
		Country: "Bangladesh",
		Batting: Batting{
			Innings: 190,
			Runs:    3676,
		},
		Bowling: Bowling{
			Innings: 19,
			Wickets: 5,
		},
	},
}
