package data

import "gar/searcher"

type Source interface {
	BuildIndexes(indexer *searcher.Indexer, numWorkers int, workerId int)
}
