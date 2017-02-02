package main

// import "fmt"
import "github.com/nsf/termbox-go"
import "time"

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

  for {
    select {
    case e := <- eventQueue:
      if e.Type == termbox.EventKey && (e.Ch == 'q' || e.Key == termbox.KeyEsc) {
        return
      }
    default:
      termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

      drawTetromino(TetrominoI, 0, 0, 0)
      drawTetromino(TetrominoJ, 1, 1, 0)
      drawTetromino(TetrominoL, 3, 1, 0)
      drawTetromino(TetrominoO, 5, 2, 0)
      drawTetromino(TetrominoS, 7, 1, 0)
      drawTetromino(TetrominoT, 9, 2, 0)
      drawTetromino(TetrominoZ, 12, 1, 1)
      // drawBoard()

      termbox.Flush()

      time.Sleep(10 * time.Millisecond)
    }
  }
}
