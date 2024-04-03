package data

import (
	"encoding/csv"
	"gar/app/ent"
	entity "gar/ent"
	"gar/searcher"
	hasher "github.com/leemcloughlin/gofarmhash"
	"google.golang.org/protobuf/proto"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type CsvDatasource struct {
	file string
}

func NewCsvSource(locationCsv string) *CsvDatasource {
	return &CsvDatasource{locationCsv}
}

func (c *CsvDatasource) BuildIndexes(indexer *searcher.Indexer, numWorkers int, workerId int) {
	file, err := os.Open(c.file)

	if err != nil {
		slog.Error("open csv file failed: {}", err)
		return
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	location, _ := time.LoadLocation("Asia/Shanghai")

	reader := csv.NewReader(file)
	processed := 0

	for {
		record, err := reader.Read()

		if err != nil {
			if err != io.EOF {
				slog.Error("read csv file failed: {}", err)
			}
			break
		}

		if len(record) < 10 {
			continue
		}

		docId := strings.TrimPrefix(record[0], "https://www.bilibili.com/video/")

		if numWorkers > 0 && int(hasher.Hash32WithSeed([]byte(docId), 0))%numWorkers != workerId {
			continue
		}

		video := &ent.BiliBiliVideo{
			Id:     docId,
			Title:  record[1],
			Author: record[3],
		}

		if len(record[2]) > 4 {
			t, err := time.ParseInLocation("2006/1/2 15:4", record[2], location)

			if err == nil {
				video.PostAt = t.Unix()
			} else {
				slog.Warn("parse {} to time failed: {}", record[2], err)
			}
		}

		n, _ := strconv.Atoi(record[4])
		video.View = uint32(n)
		n, _ = strconv.Atoi(record[5])
		video.Like = uint32(n)
		n, _ = strconv.Atoi(record[6])
		video.Coin = uint32(n)
		n, _ = strconv.Atoi(record[7])
		video.Favorite = uint32(n)
		n, _ = strconv.Atoi(record[8])
		video.Share = uint32(n)

		keywords := strings.Split(record[9], ",")

		if len(keywords) > 0 {
			for _, word := range keywords {
				word = strings.TrimSpace(word)
				if len(word) > 0 {
					video.Keywords = append(video.Keywords, strings.ToLower(word))
				}
			}
		}

		addIndex(video, indexer)
		processed++
	}
}

func addIndex(video *ent.BiliBiliVideo, indexer *searcher.Indexer) {

	doc := entity.Document{
		Id: video.Id,
	}

	bytes, err := proto.Marshal(video)

	if err == nil {
		doc.Bytes = bytes
	} else {
		slog.Error("Serialize video failed: {}", err)
		return
	}

	keywords := make([]*entity.Keyword, 0, len(video.Keywords))

	for _, keyword := range video.Keywords {
		keywords = append(keywords, &entity.Keyword{
			Field: "content",
			Word:  strings.ToLower(keyword),
		})
	}

	if len(video.Author) > 0 {
		keywords = append(keywords, &entity.Keyword{
			Field: "author",
			Word:  strings.ToLower(strings.TrimSpace(video.Author)),
		})
	}

	doc.Keywords = keywords

	n, err := indexer.Add(&doc)

	if err == nil {
		slog.Info("add {} indexes", n)
	} else {
		slog.Error("add index error: {}", err)
	}
}
