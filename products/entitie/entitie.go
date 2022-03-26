package entitie

type Product struct {
	ID          string `json:"id" bson:"_id" validate:"required"`
	Name        string `json:"name" bson:"name" validate:"required, min=3, max=50"`
	Description string `json:"description" bson:"description" validate:"required, min=10, max=200"`
	Image       string `json:"image" bson:"image" validate:"required, url"`
	Stock       uint8  `json:"stock" bson:"stock" validate:"required" gte:"0"`
	Price       uint16 `json:"price" bson:"price" validate:"required" gte:"0"`
}
