package main

import (
	"image/color"
	"machine"
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"

	fonts "github.com/NekoMaru76/gba-snake/fonts"
	"tinygo.org/x/tinyfont"
)

var (
	//KeyCodes / Buttons
	keyDOWN      = uint16(895)
	keyUP        = uint16(959)
	keyLEFT      = uint16(991)
	keyRIGHT     = uint16(1007)
	keyLSHOULDER = uint16(511)
	keyRSHOULDER = uint16(767)
	keyA         = uint16(1022)
	keyB         = uint16(1021)
	keySTART     = uint16(1015)
	keySELECT    = uint16(1019)

	// Register display
	regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))

	// Register keypad
	regKEYPAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000130)))

	// Display from machine
	display = machine.Display

	// Screen resolution
	screenWidth, screenHeight = display.Size()

	// Colors
	black = color.RGBA{}
	white = color.RGBA{255, 255, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
	red   = color.RGBA{255, 0, 0, 255}

	// Google colors
	gBlue   = color.RGBA{66, 163, 244, 255}
	gRed    = color.RGBA{219, 68, 55, 255}
	gYellow = color.RGBA{244, 160, 0, 255}
	gGreen  = color.RGBA{15, 157, 88, 255}

	// Coordinates
	x int16 = 100 //TODO: horizontally center
	y int16 = 100 //TODO: vertically center
)

func drawStartScreen() {

}

func clearScreen() {

}

func update(interrupt.Interrupt) {
	switch keyValue := regKEYPAD.Get(); keyValue {
	// Start the "game"
	case keySTART:
		// Clear display
		clearScreen()
		// Display gopher
		tinyfont.DrawChar(display, &fonts.NotoSansBold58pt, x, y, 'B', green)
	// Go back to Menu
	case keySELECT:
		clearScreen()
		drawGophers()
	// Gopher go to the right
	case keyRIGHT:
		// Clear display
		clearScreen()
		x = x + 10
		// display gopher at right
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	// Gopher go to the left
	case keyLEFT:
		// Clear display
		clearScreen()
		x = x - 10
		// display gopher at right
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	// Gopher go to the down
	case keyDOWN:
		// Clear display
		clearScreen()
		y = y + 10
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	// Gopher go to the up
	case keyUP:
		// Clear display
		clearScreen()
		y = y - 10
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	// Gopher jump
	case keyA:
		// Clear display
		clearScreen()
		// Display the gopher up
		y = y - 20
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
		// Clear the display
		clearScreen()
		// Display the gopher down
		y = y + 20
		tinyfont.DrawChar(display, &fonts.Regular58pt, x, y, 'B', green)
	}
}

func main() {
	// Set up the display
	display.Configure()

	// Register display status
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	drawStartScreen()

	// Creates an interrupt that will call the "update" fonction below, hardware way to display things on the screen
	interrupt.New(machine.IRQ_VBLANK, update).Enable()

	// Infinite loop to avoid exiting the application
	for {
	}
}
