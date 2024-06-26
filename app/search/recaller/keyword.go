package recaller

import (
	"gar/app/ent"
	"gar/app/search/bus"
	ent2 "gar/ent"
	"google.golang.org/protobuf/proto"
	"strings"
)

type KeywordRecaller struct {
}

func (k *KeywordRecaller) Recall(context *bus.Context) []*ent.BiliBiliVideo {
	r := context.Request
	indexer := context.Indexer

	if r == nil || indexer == nil {
		return nil
	}

	keywords := r.Keywords
	query := &ent2.TermQuery{}

	if len(keywords) > 0 {
		for _, keyword := range keywords {
			query = query.And(ent2.NewTermQuery("content", strings.ToLower(keyword)))
		}
	}

	if len(r.Author) > 0 {
		query = query.And(ent2.NewTermQuery("author", strings.ToLower(r.Author)))
	}

	orFlags := []uint64{ent.CategoryFeatures(r.Categories)}
	docs := indexer.Search(query, 0, 0, orFlags)
	videos := make([]*ent.BiliBiliVideo, 0, len(docs))

	for _, doc := range docs {
		var video *ent.BiliBiliVideo
		if err := proto.Unmarshal(doc.Bytes, video); err != nil {
			videos = append(videos, video)
		}
	}

	return videos
}
