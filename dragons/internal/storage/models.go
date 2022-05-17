package storage

type Bucket struct {
	Name         string
	CreationDate string
}

type BucketObject struct {
	Name string
	Size int64
}

type BucketResponse struct {
	Name    string
	Success bool
	Action  string
}

type ItemResponse struct {
	Bucket  string
	File    string
	Success bool
	Action  string
}
