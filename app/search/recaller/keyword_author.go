package recaller

import (
	"gar/app/ent"
	"gar/app/search/bus"
	ent2 "gar/ent"
	"google.golang.org/protobuf/proto"
	"strings"
)

type KeywordAndAuthorRecaller struct {
}

func (k *KeywordAndAuthorRecaller) Recall(context *bus.Context) []*ent.BiliBiliVideo {
	r := context.Request
	indexer := context.Indexer

	if r == nil || indexer == nil {
		return nil
	}

	keywords := r.Keywords
	query := &ent2.TermQuery{}

	if len(keywords) > 0 {
		for _, keyword := range keywords {
			query = query.And(ent2.NewTermQuery("conrent", keyword))
		}
	}

	value := context.Upstream.Value(bus.UN("username"))

	if value != nil {
		if author, ok := value.(string); ok {
			if len(author) > 0 {
				query = query.And(ent2.NewTermQuery("author", strings.ToLower(author)))
			}
		}
	}

	orFlags := []uint64{ent.CategoryFeatures(r.Categories)}
	docs := indexer.Search(query, 0, 0, orFlags)
	videos := make([]*ent.BiliBiliVideo, 0, len(docs))
	for _, doc := range docs {
		var video ent.BiliBiliVideo
		if err := proto.Unmarshal(doc.Bytes, &video); err != nil {
			videos = append(videos, &video)
		}
	}

	return videos
}
