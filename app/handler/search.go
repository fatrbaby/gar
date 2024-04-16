package handler

import (
	"gar/app/ent"
	ent2 "gar/ent"
	"gar/shortcut"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/proto"
	"strings"
)

func (h *Handler) Search(c *fiber.Ctx) error {
	b := new(ent.RequestBody)

	if err := c.BodyParser(b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	keywords := shortcut.CleanKeywords(b.Keywords)

	if len(keywords) == 0 && len(b.Author) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "keyword or author are required",
		})
	}

	query := &ent2.TermQuery{}

	if len(keywords) > 0 {
		for _, keyword := range keywords {
			query = query.And(ent2.NewTermQuery("content", strings.ToLower(keyword)))
		}
	}

	if len(b.Author) > 0 {
		query = query.And(ent2.NewTermQuery("author", strings.ToLower(b.Author)))
	}

	orFlag := []uint64{ent.CategoryFeatures(b.Categories)}
	docs := h.indexer.Search(query, 0, 0, orFlag)
	videos := make([]ent.BiliBiliVideo, 0, len(docs))

	for _, doc := range docs {
		var video ent.BiliBiliVideo
		if err := proto.Unmarshal(doc.Bytes, &video); err == nil {
			if video.View >= uint32(b.ViewsFrom) && (b.ViewsTo <= 0 || video.View <= uint32(b.ViewsTo)) {
				videos = append(videos, video)
			}
		}
	}

	return c.JSON(videos)
}
