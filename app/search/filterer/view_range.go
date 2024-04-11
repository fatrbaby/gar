package filterer

import (
	"gar/app/ent"
	"gar/app/search/bus"
)

type ViewRangeFilterer struct{}

func (v *ViewRangeFilterer) Apply(context *bus.Context) {
	r := context.Request

	if r == nil {
		return
	}

	if r.ViewsFrom > r.ViewsTo {
		return
	}

	videos := make([]*ent.BiliBiliVideo, 0, len(context.Results))

	for _, video := range context.Results {
		if video.View > uint32(r.ViewsFrom) && video.View <= uint32(r.ViewsTo) {
			videos = append(videos, video)
		}
	}

	context.Results = videos
}
