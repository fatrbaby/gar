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

var features = map[uint64]string{
	Information:       "资讯",
	Society:           "社会",
	Hotspot:           "热点",
	Life:              "生活",
	Knowledge:         "知识",
	Universal:         "环球",
	Game:              "游戏",
	Comprehensive:     "综合",
	Daily:             "日常",
	FilmAndTelevision: "影视",
	Anime:             "动漫",
	Tech:              "科技",
	Entertainment:     "娱乐",
	Programming:       "编程",
}

func SimpleFeatures(keywords []string) uint64 {
	var feature uint64

	for f, category := range features {
		if slices.Contains(keywords, category) {
			feature |= f
		}
	}

	return feature
}
