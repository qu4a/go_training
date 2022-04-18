package main

import (
	"gadget"
)

type Play interface { //оперделяет тип интерфейся
	Play(string) //должен содрежать метод Play c параметром стринг
	Stop()
}

func playList(device Play, songs []string) { //допустимо любое значение, поддерживающее Plyaer, не только TapePlayer
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixplay := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Play = gadget.TapePlayer{} //объявляем переменную для хранения любого значения, поддреживающего Player
	playList(player, mixplay)
	player = gadget.TapeRecorder{}
	playList(player, mixplay)
}
