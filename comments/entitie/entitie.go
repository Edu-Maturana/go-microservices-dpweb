package entitie

type Comment struct {
	ID        string `json:"id" bson:"_id" validate:"required"`
	ProductID string `json:"product_id" bson:"product_id" validate:"required"`
	UserID    string `json:"user_id" bson:"user_id" validate:"required"`
	Body      string `json:"body" bson:"body" validate:"required, min=10, max=200"`
	CreatedAt string `json:"created_at" bson:"created_at" validate:"required"`
}
