package downloader

import "context"

type Downloader interface {
	Download(ctx context.Context) error
}

func Download(downloader Downloader) error {
	ctx := context.Background()
	return downloader.Download(ctx)
}
