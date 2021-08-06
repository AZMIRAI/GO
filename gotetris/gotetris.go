// gotetris.go

package main

import (
    "fmt"
    "strings"
    "time"

    "github.com/nsf/termbox-go"
)

const backgroundColor = termbox.ColorBlack
const instructionColor = termbox.ColorWhite
const defaultMarginWidth = 2
const defaultMarginHeight = 1
const titleStartX = defaultMarginWidth
const titleStartY = defaultMarginHeight
const titleHeight = 1
const titleEndY = titleStartY + titleHeight
const boardStartX = defaultMarginWidth
const boardStartY = titleEndY + defaultMarginHeight
const boardWidth = 10
const boardHeight = 16
const boardEndX = boardStartX + boardWidth
const boardEndY = boardStartY + boardHeight
const instructionStartX = boardEndX + defaultMarginWidth
const instructionStartY = boardStartY

const title = "Tetris Written in Go Lang"

var instructions = []string{
    "Goal: Fill in 5 lines1",
    "",
    "\u2190    Left",
    "\u2192    Right",
    "\u2191    Rotate",
    "\u2193    Drop faster",
    "s         Start",
    "p         Pause",
    "esc       Exit",
    "",
    "Level: %v",
    "Lines: %v",
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
    for _, c := range msg {
        termbox.SetCell(x, y, c, fg, bg)
        x++
    }
}

func draw() {
    termbox.Clear(backgroundColor, backgroundColor)
    tbprint(titleStartX, titleStartY, instructionColor, backgroundColor, title)
    for y := boardStartY; y < boardEndY; y++ {
        for x := boardStartX; x < boardEndX; x++ {
            termbox.SetCell(x, y, ' ', termbox.ColorGreen, termbox.ColorGreen)
        }
    }
    for i, instruction := range instructions {
        if strings.HasPrefix(instruction, "Level:") {
            instruction = fmt.Sprintf(instruction, 0)
        } else if strings.HasPrefix(instruction, "Lines:") {
            instruction = fmt.Sprintf(instruction, 0)
        }
        tbprint(instructionStartX, instructionStartY+i, instructionColor, backgroundColor, instruction)
    }
    termbox.Flush()
}

func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    eventQueue := make(chan termbox.Event)
    go func() {
        for {
            eventQueue <- termbox.PollEvent()
        }
    }()

    draw()

loop:
    for {
        select {
        case ev := <-eventQueue:
            if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
                break loop
            }
        default:
            draw()
            time.Sleep(10 * time.Millisecond)
        }
    }
}
