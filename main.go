package mensa_wue

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func getRawMenu(mensa string) *html.Node {
	var endpoint = "https://www.swerk-wue.de/wuerzburg/essen-trinken/mensen-speiseplaene/" + mensa + "/menu"

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil
	}
	defer response.Body.Close()

	node, err := html.Parse(response.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return nil
	}

	return node
}

func processData(node *html.Node) Menu {

	menu := Menu{}

	hero := getNextChildWithClass(node, "hero")
	menu.Mensa = getNextChildOfElementType(hero, "h1").FirstChild.Data

	for _, m := range getAllChildrenWithClass(node, "day-menu") {
		dayMenu := DayMenu{}
		dayMenu.Date = getValueByKey(m.Attr, "data-day")

		entries := getNextChildWithClass(m, "day-menu-entries")
		for _, e := range getDirectChildrenOfElementType(entries, "article") {
			food := Food{}

			types := getNextChildWithClass(e, "food-type")
			for _, t := range getDirectChildrenOfElementType(types, "span") {
				food.Types = append(food.Types, getValueByKey(t.Attr, "title"))
			}

			food.Name = getNextChildOfElementType(e, "h5").FirstChild.Data
			prices := getNextChildWithClass(e, "price")
			food.Price = Price{
				Students: getValueByKey(prices.Attr, "data-price-student") + "€",
				Servants: getValueByKey(prices.Attr, "data-price-servant") + "€",
				Guests:   getValueByKey(prices.Attr, "data-price-guest") + "€",
			}

			food.Info.IsClimatePlate = getNextChildWithClass(e, "climate-plate") != nil
			energy := getNextChildWithText(e, " Brennwert: ")
			if energy != nil {
				food.Info.Energy = energy.Parent.NextSibling.NextSibling.FirstChild.Data
			}

			dayMenu.Options = append(dayMenu.Options, food)
		}

		menu.Menus = append(menu.Menus, dayMenu)
	}

	return menu
}

func getMenu(mensa string) []byte {
	rawMenu := getRawMenu(mensa)
	menu := processData(rawMenu)

	json, err := json.Marshal(menu)
	if err != nil {
		fmt.Println("Error processing menu:", err)
		return nil
	}

	return json
}
