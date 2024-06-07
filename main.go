package main

import (
	// "bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Display struct {
	ID        string
	Width     int
	Height    int
	PosX      int
	PosY      int
	MidX      int
	MidY      int
	IsBuiltin bool
}

func analyzeBytes(bytes []byte) {
	for i, b := range bytes {
		fmt.Printf("Byte: %d: %d (%c)\n", i, b, b)
	}
}

func parseDisplayID(display *Display, line string) {
	_, err := fmt.Sscanf(line, "Display ID: %s\n", &display.ID)

	if err != nil {
		fmt.Println("Error parsing display ID: ", err)
		panic(err)
	}
}

func parseDisplayResolution(display *Display, line string) {
	_, err := fmt.Sscanf(line, "Resolution: %dx%d\n", &display.Width, &display.Height)

	if err != nil {
		fmt.Println("Error parsing display resolution: ", err)
		panic(err)
	}
}

func parseDisplayPosition(display *Display, line string) {
	_, err := fmt.Sscanf(line, "Position: (%d, %d)\n", &display.PosX, &display.PosY)

	if err != nil {
		fmt.Println("Error parsing display position:", err)
		panic(err)
	}
}

func assignMidpoint(display *Display) {
	display.MidX = display.PosX + display.Width/2
	display.MidY = display.PosY + display.Height/2
}

func main() {
	// Check if executable exists
	_, err := exec.LookPath("./c/switcher_c")

	if err != nil {
		fmt.Println("Error: switcher_c executable not found.")
		fmt.Println("Please compile the C program first by running make in the c directory.")
		return
	}

	switcherCmd := exec.Command("./c/switcher_c")
	switcherBytes, err := switcherCmd.Output()

	if err != nil {
		panic(err)
	}

	switcherOut := string(switcherBytes)

	var numDisplays int

	_, err = fmt.Sscanf(
		switcherOut,
		"Number of active displays: %d",
		&numDisplays,
	)

	if err != nil {
		fmt.Println("Error parsing number of displays: ", err)
		return
	}

	if numDisplays == 0 {
		fmt.Println("No displays found")
		return
	}

	fmt.Println("Displays found:", numDisplays)

	chunks := strings.Split(switcherOut, "\n\n")

	displays := make([]Display, 0, numDisplays)

	for _, chunk := range chunks[1:] {

		// Initialize a new display
		display := Display{
			ID:        "",
			Width:     0,
			Height:    0,
			PosX:      0,
			PosY:      0,
			IsBuiltin: false,
		}

		lines := strings.Split(chunk, "\n")

		for _, line := range lines {
			if strings.HasPrefix(line, "Display ID:") {
				parseDisplayID(&display, line)
			} else if strings.HasPrefix(line, "Resolution:") {
				parseDisplayResolution(&display, line)
			} else if strings.HasPrefix(line, "Position:") {
				parseDisplayPosition(&display, line)
			}
		}

		if display.ID != "" {
			if display.PosX == 0 && display.PosY == 0 {
				display.IsBuiltin = true
			}

			// Assign midpoint
			display.MidX = display.PosX + display.Width/2
			display.MidY = display.PosY + display.Height/2

			displays = append(displays, display)
		}

	}

	fmt.Println(displays)

	// Find the primary display
	var primaryDisplay *Display

	for _, display := range displays {
		if display.IsBuiltin {
			primaryDisplay = &display
			break
		}
	}

	if primaryDisplay == nil {
		panic("Error: Primary display not found.")
	}

	singleClick := fmt.Sprintf("c:%d,%d", primaryDisplay.MidX, primaryDisplay.MidY)

	// Make two single clicks on the primary display
	exec.Command("cliclick", singleClick, singleClick).Run()
}
