package mensa_menu_wuerzburg

import (
	"slices"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func getNextChildOfElementType(node *html.Node, element string) *html.Node {
	if node.Data == element {
		return node
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result := getNextChildOfElementType(c, element)
		if result != nil {
			return result
		}
	}

	return nil
}

func getDirectChildrenOfElementType(node *html.Node, element string) []*html.Node {
	children := []*html.Node{}

	for n := range node.ChildNodes() {
		if n.Data == element {
			children = append(children, n)
		}
	}

	return children
}

func getNextChildWithClass(node *html.Node, class string) *html.Node {
	classes := strings.Split(getValueByKey(node.Attr, "class"), " ")
	if slices.Contains(classes, class) {
		return node
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result := getNextChildWithClass(c, class)
		if result != nil {
			return result
		}
	}

	return nil
}

func getAllChildrenWithClass(node *html.Node, class string) []*html.Node {
	children := []*html.Node{}

	for n := range node.ChildNodes() {

		children = append(children, getAllChildrenWithClass(n, class)...)

		classes := strings.Split(getValueByKey(n.Attr, "class"), " ")
		if slices.Contains(classes, class) {
			children = append(children, n)
		}
	}

	return children
}

func getNextChildWithText(node *html.Node, text string) *html.Node {
	if node.FirstChild.Data == text {
		return node
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result := getNextChildOfElementType(c, text)
		if result != nil {
			return result
		}
	}

	return nil
}

func getValueByKey(attributes []html.Attribute, key string) string {
	for _, a := range attributes {
		if a.Key == key {
			return a.Val
		}
	}

	return ""
}

func getDateByCalendarDay(day int) Date {
	berlinTime, _ := time.LoadLocation("Europe/Berlin")
	t := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, berlinTime)
	t = t.AddDate(0, 0, int(day))

	return Date{
		Day:   t.Day(),
		Month: int(t.Month()),
	}
}
