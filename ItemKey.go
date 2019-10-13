package main

func NewKeyItem() *Item {
	item := new(Item)

	item.Icon = 'K'
	item.Id = "key"
	item.Name = "Key"
	item.Description = "You can open doors."
	item.UniqueUsage = true
	item.RareItem = false
	item.UseItem = func() bool {
		return true
	}
	item.currentStack = 1
	item.maxStack = 0

	return item
}
