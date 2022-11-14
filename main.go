package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"tic-cat-dog-client/src/core"
	"tic-cat-dog-client/src/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := game.Init()
	gameAdapter := core.GameAdapter{Game: g}
	if err := ebiten.RunGame(&gameAdapter); err != nil {
		log.Fatal(err)
	}
	//connection, err := net.Dial("tcp", "127.0.0.1:7171")
	//
	//if err != nil {
	//	core.HandleError(err)
	//}
	//
	//var wg sync.WaitGroup
	//clientUUID, _ := uuid.NewRandom()
	//client := core.Client{Connection: connection, Id: clientUUID}
	//
	//wg.Add(1)
	//go func() {
	//	client.Connect()
	//	wg.Done()
	//}()
	//
	//wg.Wait()

}

func allZero(bytes []byte) bool {
	for _, b := range bytes {
		if b != 0 {
			return false
		}
	}

	return true
}
