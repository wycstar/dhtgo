package orm

import "container/list"

type Metadata struct {
	ID string `bson:"_id" json:"id"`
	Name string `bson:"n" json:"name"`
	TotalSize int64 `bson:"l" json:"size"`
	TotalLength int32 `bson:"s" json:"length"`
	Enable bool `bson:"e" json:"enable"`
	FileList list.List `bson:"f" json:"file_list"`

	CreatedAt int64 `bson:"t" json:"created_at"`
	UpdatedAt int64 `bson:"m" json:"updated_at"`
}
