package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"

	"frog_programming_language/frog"
)

var fileContent []byte
var screenWidth = int32(1000)
var screenHeight = int32(700)

const buttonWidth = 180
const buttonHeight = 40
const buttonPadding = 10
const editorHeight = 400
const editorWidth = 900

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(screenWidth, screenHeight, "Frog Language Raylib GUI")
	rl.SetTargetFPS(60)

	// Define button positions
	buttonX := int32(50)
	loadButton := rl.NewRectangle(float32(buttonX), 50, buttonWidth, buttonHeight)
	lexButton := rl.NewRectangle(float32(buttonX), 100, buttonWidth, buttonHeight)
	parseButton := rl.NewRectangle(float32(buttonX), 150, buttonWidth, buttonHeight)
	runButton := rl.NewRectangle(float32(buttonX), 200, buttonWidth, buttonHeight)

	// Text editor area
	editorRect := rl.NewRectangle(50, 260, editorWidth, editorHeight)

	// Font size and scrolling for text display
	fontSize := float32(12)
	textScroll := float32(0)

	// Content for text editor
	contentText := "Welcome to Frog Language Editor\nClick 'Upload Frog File' to begin\n"
	editorLines := strings.Split(contentText, "\n")

	for !rl.WindowShouldClose() {
		mousePosition := rl.GetMousePosition()

		// Handle button clicks
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if rl.CheckCollisionPointRec(mousePosition, loadButton) {
				// For now, using a simple text input for file path
				// In a more advanced implementation, you could integrate with OS-specific file dialogs
				// or create a custom file browser within Raylib
				selectedFile := showFileDialog()
				if selectedFile != "" {
					content, err := os.ReadFile(selectedFile)
					if err != nil {
						contentText = fmt.Sprintf("Error reading file: %v\n", err)
					} else {
						fileContent = content
						contentText = string(content)
					}
					editorLines = strings.Split(contentText, "\n")
				}
			} else if rl.CheckCollisionPointRec(mousePosition, lexButton) {
				if fileContent == nil {
					contentText = "Please load a file first."
				} else {
					l := frog.NewLexer(string(fileContent))
					tokens := l.GetAllTokens()
					var result strings.Builder
					for _, token := range tokens {
						result.WriteString(fmt.Sprintf("%s\n", token.String()))
					}
					contentText = result.String()
					editorLines = strings.Split(contentText, "\n")
				}
			} else if rl.CheckCollisionPointRec(mousePosition, parseButton) {
				if fileContent == nil {
					contentText = "Please load a file first."
				} else {
					l := frog.NewLexer(string(fileContent))
					p := frog.NewParser(l)
					program := p.ParseProgram()
					if p.IsThereAnyErrors() {
						var errs strings.Builder
						for _, err := range p.Errors() {
							errs.WriteString(err + "\n")
						}
						contentText = errs.String()
						editorLines = strings.Split(contentText, "\n")
					} else {
						contentText = program.String()
						editorLines = strings.Split(contentText, "\n")
					}
				}
			} else if rl.CheckCollisionPointRec(mousePosition, runButton) {
				if fileContent == nil {
					contentText = "Please load a file first."
				} else if strings.Contains(string(fileContent), "FRG_Input") {
					contentText = "Input statements are not supported in the GUI."
				} else {
					l := frog.NewLexer(string(fileContent))
					p := frog.NewParser(l)
					program := p.ParseProgram()
					if p.IsThereAnyErrors() {
						var errs strings.Builder
						for _, err := range p.Errors() {
							errs.WriteString(err + "\n")
						}
						contentText = errs.String()
						editorLines = strings.Split(contentText, "\n")
					} else {
						oldStdout := os.Stdout
						r, wFile, _ := os.Pipe()
						os.Stdout = wFile

						env := frog.NewEnvironment()
						evaluated := frog.Eval(program, env)

						wFile.Close()
						os.Stdout = oldStdout

						var buf bytes.Buffer
						io.Copy(&buf, r)

						if evaluated != nil && evaluated.Type() == "ERROR" {
							contentText = evaluated.Inspect()
						} else {
							contentText = buf.String()
						}
						editorLines = strings.Split(contentText, "\n")
					}
				}
			}
		}

		// Handle text editor scrolling
		if rl.CheckCollisionPointRec(mousePosition, editorRect) {
			scrollAmount := rl.GetMouseWheelMove() * 20
			textScroll += float32(scrollAmount)
		}

		// Draw everything
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw title
		titleText := "Frog Language Raylib GUI"
		titleFontSize := float32(24)
		titleBounds := rl.MeasureTextEx(rl.GetFontDefault(), titleText, titleFontSize, 0)
		rl.DrawText(titleText, int32(screenWidth/2)-int32(titleBounds.X/2), 10, int32(titleFontSize), rl.Black)

		// Draw buttons
		// Load button
		buttonColor := rl.LightGray
		if rl.CheckCollisionPointRec(mousePosition, loadButton) {
			buttonColor = rl.Gray
		}
		rl.DrawRectangleRec(loadButton, buttonColor)
		rl.DrawRectangleLinesEx(loadButton, 1, rl.Black)
		buttonTextSize := rl.MeasureTextEx(rl.GetFontDefault(), "Upload Frog File", fontSize, 0)
		rl.DrawText("Upload Frog File",
			int32(loadButton.X+(loadButton.Width-buttonTextSize.X)/2),
			int32(loadButton.Y+(loadButton.Height-buttonTextSize.Y)/2),
			int32(fontSize), rl.Black)

		// Lex button
		buttonColor = rl.LightGray
		if rl.CheckCollisionPointRec(mousePosition, lexButton) {
			buttonColor = rl.Gray
		}
		rl.DrawRectangleRec(lexButton, buttonColor)
		rl.DrawRectangleLinesEx(lexButton, 1, rl.Black)
		buttonTextSize = rl.MeasureTextEx(rl.GetFontDefault(), "Lexical Analysis", fontSize, 0)
		rl.DrawText("Lexical Analysis",
			int32(lexButton.X+(lexButton.Width-buttonTextSize.X)/2),
			int32(lexButton.Y+(lexButton.Height-buttonTextSize.Y)/2),
			int32(fontSize), rl.Black)

		// Parse button
		buttonColor = rl.LightGray
		if rl.CheckCollisionPointRec(mousePosition, parseButton) {
			buttonColor = rl.Gray
		}
		rl.DrawRectangleRec(parseButton, buttonColor)
		rl.DrawRectangleLinesEx(parseButton, 1, rl.Black)
		buttonTextSize = rl.MeasureTextEx(rl.GetFontDefault(), "Syntaxical Analysis", fontSize, 0)
		rl.DrawText("Syntaxical Analysis",
			int32(parseButton.X+(parseButton.Width-buttonTextSize.X)/2),
			int32(parseButton.Y+(parseButton.Height-buttonTextSize.Y)/2),
			int32(fontSize), rl.Black)

		// Run button
		buttonColor = rl.LightGray
		if rl.CheckCollisionPointRec(mousePosition, runButton) {
			buttonColor = rl.Gray
		}
		rl.DrawRectangleRec(runButton, buttonColor)
		rl.DrawRectangleLinesEx(runButton, 1, rl.Black)
		buttonTextSize = rl.MeasureTextEx(rl.GetFontDefault(), "Run", fontSize, 0)
		rl.DrawText("Run",
			int32(runButton.X+(runButton.Width-buttonTextSize.X)/2),
			int32(runButton.Y+(runButton.Height-buttonTextSize.Y)/2),
			int32(fontSize), rl.Black)

		// Draw text editor area
		rl.DrawRectangleRec(editorRect, rl.White)
		rl.DrawRectangleLinesEx(editorRect, 1, rl.Black)

		// Draw content in text editor with scrolling
		yOffset := float32(editorRect.Y) + 10 - textScroll
		for _, line := range editorLines {
			if yOffset > editorRect.Y && yOffset < editorRect.Y+editorRect.Height {
				rl.DrawText(line, int32(editorRect.X)+5, int32(yOffset), int32(fontSize), rl.Black)
			}
			yOffset += fontSize + 5
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// Since Raylib doesn't have a built-in file dialog, we'll create a simple implementation
// that allows user to specify a file path in console for now
func showFileDialog() string {
	// Note: In a real implementation, you might want to use external libraries
	// or OS-specific calls to show a native file dialog

	// For now, we'll just return an empty string to handle the case
	// where the user doesn't specify a file
	fmt.Print("Enter path to .frg file (or press Enter to cancel): ")
	var path string
	fmt.Scanln(&path)

	if path == "" {
		return ""
	}

	// Validate that it's a .frg file
	if !strings.HasSuffix(path, ".frg") {
		fmt.Println("Please select a .frg file")
		return ""
	}

	return path
}
