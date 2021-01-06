package main

import (
         "log"
         "github.com/hajimehoshi/ebiten/v2"
         "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func update(screen *ebiten.Image) error {
         if ebiten.IsDrawingSkipped() {
                 return nil
         }
         ebitenutil.DebugPrint(screen, "Hello, World!")
         return nil
}

func main() {
         if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
                 log.Fatal(err)
         }
}
