package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	// ゲームパッド番号。今回は0番で決め打ちする。
	const gamepadID = 0
	// 現在押されているボタンを取得する。
	buttons := []ebiten.GamepadButton{}
	for b := ebiten.GamepadButton0; b <= ebiten.GamepadButtonMax; b++ {
		if ebiten.IsGamepadButtonPressed(gamepadID, b) {
			buttons = append(buttons, b)
		}
	}
	// 現在の軸の状態を取得する。
	axes := make([]float64, ebiten.GamepadAxisNum(gamepadID))
	for i := range axes {
		axes[i] = ebiten.GamepadAxis(gamepadID, i)
	}
	if ebiten.IsRunningSlowly() {
		return nil
	}
	// ボタンの状態を文字列化する。
	msg := "Buttons: "
	for _, b := range buttons {
		msg += fmt.Sprintf("%d ", b)
	}
	// 軸の状態を文字列化する。
	// 軸の傾きは-1から1の浮動小数点で表される。
	msg += "\nAxes:\n"
	for _, a := range axes {
		msg += fmt.Sprintf("  %0.6f\n", a)
	}
	ebitenutil.DebugPrint(screen, msg)
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Gamepad"); err != nil {
		log.Fatal(err)
	}
}
