package mensa_menu_wuerzburg

type mensas struct {
	JOSEF_SCHNEIDER_STRASSE string
	ROENTGENRING            string
	STUDENTENHAUS           string
	HUBLAND_NORD            string
	HUBLAND_SUED            string
}

var Mensa = mensas{
	JOSEF_SCHNEIDER_STRASSE: "mensa-josef-schneider-strasse-wuerzburg",
	ROENTGENRING:            "mensa-roentgenring-wuerzburg",
	STUDENTENHAUS:           "mensa-am-studentenhaus-wuerzburg",
	HUBLAND_NORD:            "mensateria-campus-hubland-nord-wuerzburg",
	HUBLAND_SUED:            "mensa-campus-hubland-sued",
}

type price struct {
	Students string `json:"students"`
	Servants string `json:"servants"`
	Guests   string `json:"guests"`
}

type info struct {
	IsClimatePlate bool   `json:"isClimatePlate"`
	Energy         string `json:"energy"`
}

type food struct {
	Name  string   `json:"name"`
	Price price    `json:"price"`
	Types []string `json:"types"`
	Info  info     `json:"info"`
}

type dayMenu struct {
	Date    string `json:"date"`
	Options []food `json:"options"`
}

type menu struct {
	Mensa string    `json:"mensa"`
	Menus []dayMenu `json:"menus"`
}
