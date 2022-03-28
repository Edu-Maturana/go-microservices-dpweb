package entitie

type Product struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required"`
	Image       string `json:"image" bson:"image" validate:"required"`
	Stock       uint8  `json:"stock" bson:"stock" validate:"required" gte:"0"`
	Price       uint16 `json:"price" bson:"price" validate:"required" gte:"0"`
}
