package ModelForGin

type CreateIceCream struct {
	Flavor string `json:"flavor" binding:"required"`
}
