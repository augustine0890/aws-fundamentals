package storage

type Bucket struct {
	Name         string
	CreationDate string
}

type BucketObject struct {
	Name string
	Size int64
}