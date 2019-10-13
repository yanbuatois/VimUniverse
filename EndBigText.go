package main

type EndBigText struct {
	*BigText
}

var LostText string
var WonText string

func init() {
	WonText = "Congratulations!\n\nYou successfully went through numerous traps in the Vim Universe, and you have managed to get out from Vim.\nYou can now return to your terminal with the \":q\" command.\nThank you for playing"
	LostText = "You lost...\n\nUnfortunately, the Vim Universe has succeeded to hold you prisoner, with its awful traps.\nIf you dare, you can retry with the \":retry\" command."
}

func NewEndBigText(text string) *EndBigText {
	retour := new(EndBigText)
	retour.BigText = NewBigText(text)

	return retour
}
