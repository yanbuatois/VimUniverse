package main

func NewBoatItem () (*Item) {
	item := new(Item)

	item.Icon = '⌴'
	item.Id = "boat"
	item.Name = "Boat"
	item.Description = "You can float on the water to discover new areas."
	item.UniqueUsage = false
	item.RareItem = true
	item.UseItem = func() bool {
		return false
	}
	item.OnPickBehaviour = true
	item.pickItem = func() {
		TheGame.Player.CanFloat = true
	}

	return item
}