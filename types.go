package mensa_wue

type Mensas struct {
	JOSEF_SCHNEIDER_STRASSE string
	ROENTGENRING            string
	STUDENTENHAUS           string
	HUBLAND_NORD            string
	HUBLAND_SUED            string
}

var Mensa = Mensas{
	JOSEF_SCHNEIDER_STRASSE: "mensa-josef-schneider-strasse-wuerzburg",
	ROENTGENRING:            "mensa-roentgenring-wuerzburg",
	STUDENTENHAUS:           "mensa-am-studentenhaus-wuerzburg",
	HUBLAND_NORD:            "mensateria-campus-hubland-nord-wuerzburg",
	HUBLAND_SUED:            "mensa-campus-hubland-sued",
}

type Price struct {
	Students string `json:"students"`
	Servants string `json:"servants"`
	Guests   string `json:"guests"`
}

type Info struct {
	IsClimatePlate bool   `json:"isClimatePlate"`
	Energy         string `json:"energy"`
}

type Food struct {
	Name  string   `json:"name"`
	Price Price    `json:"price"`
	Types []string `json:"types"`
	Info  Info     `json:"info"`
}

type DayMenu struct {
	Date    string `json:"date"`
	Options []Food `json:"options"`
}

type Menu struct {
	Mensa string    `json:"mensa"`
	Menus []DayMenu `json:"menus"`
}
