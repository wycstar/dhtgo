package model

type File struct {
	Name string `bson:"n" json:"name"`
	Size int64  `bson:"s" json:"size"`
}

type Metadata struct {
	ID          string `bson:"_id" json:"id"`
	Name        string `bson:"n" json:"name"`
	TotalSize   int64  `bson:"s" json:"size"`
	TotalLength int32  `bson:"l" json:"length"`
	Enable      bool   `bson:"e" json:"enable"`
	FileList    []File `bson:"f" json:"file_list"`
	Heat        int32  `bson:"c" json:"heat"`
	CreatedAt   int64  `bson:"t" json:"created_at"`
	UpdatedAt   int64  `bson:"m" json:"updated_at"`
}

func Upsert() {
	getClient()
}
