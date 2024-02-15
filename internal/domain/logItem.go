package domain

type LogItem struct {
	ID     string `json:"id" bson:"_id, omitempty"`
	Action string `json:"action" bson:"action"`
	Status string `json:"status" bson:"status"`
}
