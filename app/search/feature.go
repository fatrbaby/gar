package search

import "slices"

const (
	Information = 1 << iota
	Society
	Hotspot
	Life
	Knowledge
	Universal
	Game
	Comprehensive
	Daily
	FilmAndTelevision
	Anime
	Tech
	Entertainment
	Programming
)

type Category struct {
	Bit  uint64 `json:"bit"`
	Text string `json:"text"`
}

var Categories = []Category{
	{Bit: Information, Text: "资讯"},
	{Bit: Society, Text: "社会"},
	{Bit: Hotspot, Text: "热点"},
	{Bit: Life, Text: "生活"},
	{Bit: Knowledge, Text: "知识"},
	{Bit: Universal, Text: "环球"},
	{Bit: Game, Text: "游戏"},
	{Bit: Comprehensive, Text: "综合"},
	{Bit: Daily, Text: "日常"},
	{Bit: FilmAndTelevision, Text: "影视"},
	{Bit: Anime, Text: "动漫"},
	{Bit: Tech, Text: "科技"},
	{Bit: Entertainment, Text: "娱乐"},
	{Bit: Programming, Text: "编程"},
}

func CategoryFeatures(keywords []string) uint64 {
	var feature uint64

	for _, category := range Categories {
		if slices.Contains(keywords, category.Text) {
			feature |= category.Bit
		}
	}

	return feature
}
