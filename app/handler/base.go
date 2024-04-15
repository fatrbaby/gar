package handler

import "gar/searcher"

type Handler struct {
	indexer *searcher.Indexer
}

func New(indexer *searcher.Indexer) *Handler {
	return &Handler{indexer: indexer}
}
