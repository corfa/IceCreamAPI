package ModelForGin

type IceCream struct {
	Flavor string `json:"flavor" binding:"required"`
}
