package main

//
// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
//
// 	"gioui.org/app"
// 	"gioui.org/layout"
// 	"gioui.org/op"
// 	"gioui.org/unit"
// 	"gioui.org/widget"
// 	"gioui.org/widget/material"
//
// 	"gioui.org/x/explorer"
//
// 	"frog_programming_language/frog"
// )
//
// var fileContent []byte
//
// func main() {
// 	go func() {
// 		var w app.Window
// 		if err := frameLoop(&w); err != nil {
// 			log.Fatal(err)
// 		}
// 		os.Exit(0)
// 	}()
// 	app.Main()
// }
//
// func frameLoop(w *app.Window) error {
// 	th := material.NewTheme()
// 	var ops op.Ops
//
// 	var (
// 		// TODO: add symantic checking
// 		// symanticButton widget.Clickable for checking symantic frog code
//
// 		// TODO: add logo image
// 		// at the bottom
//
// 		contentEditor widget.Editor
// 		loadButton    widget.Clickable
// 		lexButton     widget.Clickable
// 		parseButton   widget.Clickable
// 		runButton     widget.Clickable
// 		// symanticButton widget.Clickable
// 		// logoImage     widget.Image
// 	)
// 	exp := explorer.NewExplorer(w)
//
// 	for {
// 		e := w.Event()
// 		exp.ListenEvents(e)
//
// 		switch e := e.(type) {
// 		case app.DestroyEvent:
// 			return e.Err
// 		case app.FrameEvent:
// 			gtx := app.NewContext(&ops, e)
//
// 			if loadButton.Clicked(gtx) {
// 				go func() {
// 					reader, err := exp.ChooseFile(".frg")
// 					if err != nil {
// 						log.Printf("Error choosing file: %v", err)
// 						return
// 					}
// 					if reader == nil {
// 						log.Println("File selection cancelled.")
// 						return
// 					}
// 					defer reader.Close()
//
// 					fileContent, err = io.ReadAll(reader)
// 					if err != nil {
// 						log.Printf("Error reading file: %v", err)
// 						contentEditor.SetText(fmt.Sprintf("Error reading file: %s", err))
// 					} else {
// 						contentEditor.SetText(string(fileContent))
// 					}
// 					w.Invalidate()
// 				}()
// 			}
//
// 			if lexButton.Clicked(gtx) {
// 				if fileContent == nil {
// 					contentEditor.SetText("Please load a file first.")
// 				} else {
// 					l := frog.NewLexer(string(fileContent))
// 					tokens := l.GetAllTokens()
// 					var result strings.Builder
// 					for _, token := range tokens {
// 						result.WriteString(fmt.Sprintf("%s\n", token.String()))
// 					}
// 					contentEditor.SetText(result.String())
// 				}
// 			}
//
// 			if parseButton.Clicked(gtx) {
// 				if fileContent == nil {
// 					contentEditor.SetText("Please load a file first.")
// 				} else {
// 					l := frog.NewLexer(string(fileContent))
// 					p := frog.NewParser(l)
// 					program := p.ParseProgram()
// 					if p.IsThereAnyErrors() {
// 						var errs strings.Builder
// 						for _, err := range p.Errors() {
// 							errs.WriteString(err + "\n")
// 						}
// 						contentEditor.SetText(errs.String())
// 					} else {
// 						contentEditor.SetText(program.String())
// 					}
// 				}
// 			}
//
// 			if runButton.Clicked(gtx) {
// 				if fileContent == nil {
// 					contentEditor.SetText("Please load a file first.")
// 				} else if strings.Contains(string(fileContent), "FRG_Input") {
// 					contentEditor.SetText("Input statements are not supported in the GUI.")
// 				} else {
// 					l := frog.NewLexer(string(fileContent))
// 					p := frog.NewParser(l)
// 					program := p.ParseProgram()
// 					if p.IsThereAnyErrors() {
// 						var errs strings.Builder
// 						for _, err := range p.Errors() {
// 							errs.WriteString(err + "\n")
// 						}
// 						contentEditor.SetText(errs.String())
// 					} else {
// 						oldStdout := os.Stdout
// 						r, wFile, _ := os.Pipe()
// 						os.Stdout = wFile
//
// 						env := frog.NewEnvironment()
// 						evaluated := frog.Eval(program, env)
//
// 						wFile.Close()
// 						os.Stdout = oldStdout
//
// 						var buf bytes.Buffer
// 						io.Copy(&buf, r)
//
// 						if evaluated != nil && evaluated.Type() == "ERROR" {
// 							contentEditor.SetText(evaluated.Inspect())
// 						} else {
// 							contentEditor.SetText(buf.String())
// 						}
// 					}
// 				}
// 			}
// 			padding := layout.UniformInset(unit.Dp(16))
// 			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
// 				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 					return padding.Layout(gtx, material.H4(th, "Frog Language GUI").Layout)
// 				}),
//
// 				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 					return padding.Layout(gtx, material.Button(th, &loadButton, "Upload Frog File").Layout)
// 				}),
//
// 				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 					return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
// 						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
// 							layout.Flexed(1, material.Button(th, &lexButton, "Lexical Analysis").Layout),
// 							layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
// 							layout.Flexed(1, material.Button(th, &parseButton, "Syntaxical Analysis").Layout),
// 							layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
// 							layout.Flexed(1, material.Button(th, &runButton, "Run").Layout),
// 						)
// 					})
// 				}),
//
// 				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
// 					return padding.Layout(gtx, material.Editor(th, &contentEditor, "").Layout)
// 				}),
// 			)
// 			e.Frame(gtx.Ops)
// 		}
// 	}
// }
