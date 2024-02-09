package models

type Shortening struct {
	Alias     string `bson:"alias"`
	Url       string `bson:"url"`
	Author    string `bson:"author"`
	Clicks    int    `bson:"clicks"`
	Timestamp int64  `bson:"timestamp"`
}
