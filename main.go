package main

import (
	// "bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Display struct {
	ID     string
	Width  int
	Height int
	PosX   int
	PosY   int
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

func main() {
	switcherCmd := exec.Command("./c/switcher")
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

    if (numDisplays == 0) {
        fmt.Println("No displays found")
        return
    }

	fmt.Println("Displays found:", numDisplays)

	chunks := strings.Split(switcherOut, "\n\n")

	displays := make([]Display, 0, numDisplays)

	for _, chunk := range chunks[1:] {

		// Initialize a new display
		display := Display{
			ID:     "",
			Width:  0,
			Height: 0,
			PosX:   0,
			PosY:   0,
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
			displays = append(displays, display)
		}

	}

    fmt.Println(displays)
}
