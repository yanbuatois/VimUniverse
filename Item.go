package main

type Item struct {
	Icon            rune
	Id              string
	Name            string
	Description     string
	maxStack        int
	UniqueUsage     bool
	RareItem        bool
	UseItem         func() bool
	OnPickBehaviour bool
	pickItem        func()
	currentStack    int
}

//type ItemDbStruct struct {
//	Database [](*Item)
//}

func (item *Item) GetMaxStack() int {
	if item.RareItem {
		return 1
	} else {
		return item.maxStack
	}
}

func (item *Item) GetCurrentStack() int {
	if item.RareItem {
		return 1
	} else {
		return item.currentStack
	}
}

func (item *Item) PickItem() {
	if item.OnPickBehaviour {
		item.pickItem()
	}
}

//func (itemDb *ItemDbStruct) GetItemById (id string) (*Item) {
//	return ContainsItem(itemDb.Database, id)
//}
