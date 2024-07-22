package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

func DrawNumber(renderer *sdl.Renderer, num int, startX, startY int, color sdl.Color) {

	digits := map[int][][]int{
		0: {
			{1, 1, 1},
			{1, 0, 1},
			{1, 0, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		1: {
			{0, 1, 0},
			{1, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
		2: {
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
		},
		3: {
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},
		4: {
			{1, 0, 1},
			{1, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{0, 0, 1},
		},
		5: {
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},
		6: {
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		7: {
			{1, 1, 1},
			{0, 0, 1},
			{0, 0, 1},
			{0, 0, 1},
			{0, 0, 1},
		},
		8: {
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		9: {
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},
	}

	numStr := strconv.Itoa(num) // Convert the number to a string
	numDigits := len(numStr)    // Get the number of digits in the number
	digitWidth := 4             // Width of each digit in pixels, adjust as needed
	spacing := 30               // Spacing between digits in terms of 'blocks'

	// Calculate the total width of the number including spacing between digits
	totalWidth := numDigits*digitWidth + (numDigits-1)*spacing

	// Adjust startX so the number grows to the left
	adjustedStartX := startX - totalWidth

	for i, runeValue := range numStr {
		digit := int(runeValue - '0') // Convert each rune to its integer value

		// Correctly calculate the X position for this digit
		digitX := adjustedStartX + i*(digitWidth+spacing)

		// Assume digits is a predefined map of digit patterns
		digitPattern := digits[digit]

		// Set the draw color
		renderer.SetDrawColor(color.R, color.G, color.B, color.A)

		// Draw the digit
		for y, row := range digitPattern {
			for x, val := range row {
				if val == 1 {
					renderer.FillRect(&sdl.Rect{X: int32(digitX + x*10), Y: int32(startY + y*10), W: 10, H: 10})
				}
			}
		}
	}
}

func DrawNumberFun(renderer *sdl.Renderer, num int, startX, startY int, color [3]sdl.Color, pass int) {

	digits := map[int][][]int{
		0: {
			{1, 1, 1},
			{1, 0, 1},
			{1, 0, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		1: {
			{0, 1, 0},
			{1, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
			{1, 1, 1},
		},
		2: {
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
		},
		3: {
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},
		4: {
			{1, 0, 1},
			{1, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{0, 0, 1},
		},
		5: {
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},
		6: {
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		7: {
			{1, 1, 1},
			{0, 0, 1},
			{0, 0, 1},
			{0, 0, 1},
			{0, 0, 1},
		},
		8: {
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
		9: {
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},
	}

	numStr := strconv.Itoa(num) // Convert the number to a string
	numDigits := len(numStr)    // Get the number of digits in the number
	digitWidth := 4             // Width of each digit in pixels, adjust as needed
	spacing := 30               // Spacing between digits in terms of 'blocks'

	// Calculate the total width of the number including spacing between digits
	totalWidth := numDigits*digitWidth + (numDigits-1)*spacing

	// Adjust startX so the number grows to the left
	adjustedStartX := startX - totalWidth

	for i, runeValue := range numStr {
		digit := int(runeValue - '0') // Convert each rune to its integer value

		// Correctly calculate the X position for this digit
		digitX := adjustedStartX + i*(digitWidth+spacing)

		// Assume digits is a predefined map of digit patterns
		digitPattern := digits[digit]

		// Set the draw color
		col := color[(i+pass)%len(color)]
		renderer.SetDrawColor(col.R, col.G, col.B, col.A)

		// Draw the digit
		for y, row := range digitPattern {
			for x, val := range row {
				if val == 1 {
					renderer.FillRect(&sdl.Rect{X: int32(digitX + x*10), Y: int32(startY + y*10), W: 10, H: 10})
				}
			}
		}
	}
}

func DrawPauseSign(renderer *sdl.Renderer, windowWidth, windowHeight int32) {
	// Calculate the size and position of the pause sign
	barWidth := windowWidth / 8   // Each bar is 1/8th the width of the window
	barHeight := windowHeight / 2 // Each bar's height is half the height of the window
	gap := barWidth / 2           // Gap between the two bars

	// Calculate the starting X position for the first bar so that the sign is centered
	startX := (windowWidth - 2*barWidth - gap) / 2
	startY := (windowHeight - barHeight) / 2 // Start Y position to center the bars vertically

	// Set the draw color to white (or any color you prefer)
	renderer.SetDrawColor(255, 255, 255, 255) // RGBA

	// Draw the first bar
	firstBar := sdl.Rect{X: startX, Y: startY, W: barWidth, H: barHeight}
	renderer.FillRect(&firstBar)

	// Draw the second bar, positioned with the gap from the first bar
	secondBar := sdl.Rect{X: startX + barWidth + gap, Y: startY, W: barWidth, H: barHeight}
	renderer.FillRect(&secondBar)
}

const (
	windowWidth  = 800 //x
	windowHeight = 600 //y
	snakeSize    = 50
)

func setPixel(x, y int, c sdl.Color, pixels []byte) {
	index := (y*windowWidth + x) * 4 // 4 bytes per pixel (RGBA)

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.R
		pixels[index+1] = c.G
		pixels[index+2] = c.B
		pixels[index+3] = c.A
	}
}

type Vec2I struct {
	x int32
	y int32
}

// Add adds another Vec2I to this one and returns the result.
func (v Vec2I) Add(other Vec2I) Vec2I {
	return Vec2I{v.x + other.x, v.y + other.y}
}

// Subtract subtracts another Vec2I from this one and returns the result.
func (v Vec2I) Subtract(other Vec2I) Vec2I {
	return Vec2I{v.x - other.x, v.y - other.y}
}

// Scale multiplies the vector by a scalar and returns the result.
func (v Vec2I) Scale(scalar int32) Vec2I {
	return Vec2I{v.x * scalar, v.y * scalar}
}

func (v Vec2I) ScaleF(scalar float32) Vec2I {
	return Vec2I{int32(float32(v.x) * scalar), int32(float32(v.y) * scalar)}
}

func (v Vec2I) ScaleVec(other Vec2I) Vec2I {
	return Vec2I{v.x * other.x, v.y * other.y}
}

func (v Vec2I) SetZeroComponents(value int32) Vec2I {
	var x, y int32
	if v.x == 0 {
		x = value
	}
	if v.y == 0 {
		y = value
	}
	return Vec2I{x, y}
}

func ExponentialDecay(x float64) float64 {
	a := 100.0 // Initial value
	b := 0.95  // Base value for a slow decay, closer to 1 makes the decrement smaller

	return a * math.Pow(b, x)
}

type Snake struct {
	Body             []Vec2I
	currentThickness float32
}

func (s *Snake) getRects() []sdl.Rect {
	var rects []sdl.Rect
	for _, rect := range s.Body {
		/*
			prevOrientation, nextOrientation := Vec2I{0, 0}, Vec2I{0, 0}
			if idx != 0 {
				prevOrientation = s.Body[idx-1].Subtract(rect)
			}
			if idx != len(s.Body)-1 {
				nextOrientation = s.Body[idx+1].Subtract(rect)
			}
		*/
		prevPos := rect.Scale(snakeSize)
		firstRec := sdl.Rect{X: prevPos.x, Y: prevPos.y, W: snakeSize, H: snakeSize}
		rects = append(rects, firstRec)
	}
	return rects
}

func (s *Snake) Move(dir Vec2I) {
	if len(s.Body) == 0 {
		return
	}

	// Move the head
	head := s.Body[0]
	var newX = head.x + int32(dir.x)
	var newY = head.y + int32(dir.y)
	if newX*snakeSize < 0 {
		newX = int32((windowWidth - snakeSize) / snakeSize)
	} else if newX*snakeSize >= int32(windowWidth) {
		newX = 0
	}
	if newY*snakeSize < 0 {
		newY = int32((windowHeight - snakeSize) / snakeSize)
	} else if newY*snakeSize >= int32(windowHeight) {
		newY = 0
	}
	newHead := Vec2I{newX, newY}
	s.Body = append([]Vec2I{newHead}, s.Body[:len(s.Body)-1]...)
}

func (s *Snake) CollidesWith(rect sdl.Rect) bool {
	head := s.Body[0]
	if head.x <= rect.X+rect.W && head.x+snakeSize >= rect.X && head.y <= rect.Y+rect.H && head.y+snakeSize >= rect.Y {
		return true
	}
	return false
}

func (s *Snake) CollidesWithGlobSpace(rect sdl.Rect) bool {
	head := s.Body[0]
	head = head.Scale(snakeSize)
	if head.x <= rect.X+rect.W && head.x+snakeSize >= rect.X && head.y <= rect.Y+rect.H && head.y+snakeSize >= rect.Y {
		return true
	}
	return false
}

func (s *Snake) CollidesWithSelf() bool {
	head := s.Body[0]
	for i := 1; i < len(s.Body); i++ {
		if head.x == s.Body[i].x && head.y == s.Body[i].y {
			return true
		}
	}
	return false
}

func (s *Snake) Grow() {
	if len(s.Body) == 0 {
		return
	}

	// Grow the snake
	tail := s.Body[len(s.Body)-1]
	newTail := Vec2I{x: tail.x, y: tail.y}
	s.Body = append(s.Body, newTail)
}

func getRandomFood(s *Snake) sdl.Rect {
	randomX := rand.Intn((windowWidth / snakeSize)) * snakeSize
	randomY := rand.Intn((windowHeight / snakeSize)) * snakeSize
	for _, rect := range s.getRects() {
		if rect.X == int32(randomX) && rect.Y == int32(randomY) {
			return getRandomFood(s)
		}
	}
	return sdl.Rect{X: int32(randomX), Y: int32(randomY), W: 30, H: 50}
}

func numberOfDigits(num int) int {
	if num == 0 {
		return 1
	}
	return int(math.Log10(float64(num))) + 1
}

func trySDLInit() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	return nil
}

func main() {

	rand.Seed(time.Now().UnixNano())

	os.Setenv("SDL_VIDEODRIVER", "wayland")
	if err := trySDLInit(); err != nil {
		// Fallback to X11 for Linux or Windows for Windows OS
		if runtime.GOOS == "linux" {
			os.Setenv("SDL_VIDEODRIVER", "x11")
		} else if runtime.GOOS == "windows" {
			os.Setenv("SDL_VIDEODRIVER", "windows")
		}
		if err := trySDLInit(); err != nil {
			panic(err) // Handle error if both attempts fail
		}
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("GO-Snek", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, windowWidth, windowHeight)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	pixels := make([]byte, windowWidth*windowHeight*4) // 4 bytes per pixel (RGBA)

	for y := 0; y < windowHeight; y++ {
		for x := 0; x < windowWidth; x++ {
			setPixel(x, y, sdl.Color{R: 255, G: 0, B: 0, A: 255}, pixels) // Example: setting all pixels to red
		}
	}

	texture.Update(nil, unsafe.Pointer(&pixels[0]), windowWidth*4)
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	running := true
	Snake := Snake{Body: []Vec2I{{x: 0, y: 0}}, currentThickness: 1}
	//var moveSpeed float64 = 0.1
	var moveDir Vec2I = Vec2I{0, 0}
	var count = 0

	food := getRandomFood(&Snake)
	foodNum := 0
	death := false
	pause := true
	deathcounter := 0
	funCounter := 0
	for running {
		// Handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Sym {
					case sdl.K_UP:
						pause = false
						if moveDir.y == 1 {
							break
						}
						moveDir = Vec2I{0, -1}
					case sdl.K_w:
						pause = false
						if moveDir.y == 1 {
							break
						}
						moveDir = Vec2I{0, -1}
					case sdl.K_DOWN:
						pause = false
						if moveDir.y == -1 {
							break
						}
						moveDir = Vec2I{0, 1}
					case sdl.K_s:
						pause = false
						if moveDir.y == -1 {
							break
						}
						moveDir = Vec2I{0, 1}
					case sdl.K_LEFT:
						pause = false
						if moveDir.x == 1 {
							break
						}
						moveDir = Vec2I{-1, 0}
					case sdl.K_a:
						pause = false
						if moveDir.x == 1 {
							break
						}
						moveDir = Vec2I{-1, 0}
					case sdl.K_RIGHT:
						pause = false
						if moveDir.x == -1 {
							break
						}
						moveDir = Vec2I{1, 0}
					case sdl.K_d:
						pause = false
						if moveDir.x == -1 {
							break
						}
						moveDir = Vec2I{1, 0}
					case sdl.K_ESCAPE:
						pause = false
						death = true
					case sdl.K_p:
						pause = true
					}
				}
			}
		}
		// Clear the renderer
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		// Inside your game loop
		// Calculate the dynamic speed based on the snake's size
		dynamicSpeed := 49 - math.Pow(float64(foodNum), 1.0/3.0)*10

		// Ensure dynamicSpeed does not fall below a minimum threshold
		if dynamicSpeed < 11 {
			dynamicSpeed = 11
		}

		// Dynamic drawing operations here
		// For example, drawing a line that changes position over time

		if count == int(dynamicSpeed) {
			count = 0
			Snake.Move(moveDir)
			if Snake.CollidesWithSelf() {
				death = true
			}
			if Snake.CollidesWithGlobSpace(food) {
				Snake.Grow()
				foodNum++
				food = getRandomFood(&Snake)
			}
		}
		if !pause && !death {
			count++
		}
		if pause {
			DrawPauseSign(renderer, windowWidth, windowHeight)
		}
		for _, rect := range Snake.getRects() {
			renderer.SetDrawColor(102, 180, 71, 255)
			renderer.FillRect(&food)
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.FillRect(&rect)
		}
		DrawNumber(renderer, foodNum*10, 750, 30, sdl.Color{R: 255, G: 0, B: 0, A: 128})
		texture.Update(nil, unsafe.Pointer(&pixels[0]), windowWidth*4)

		if death {
			if deathcounter%30 == 0 {
				funCounter++
			}
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			var colors = [3]sdl.Color{
				{R: 255, G: 0, B: 0, A: 255}, // Red
				{R: 0, G: 255, B: 0, A: 255}, // Green
				{R: 0, G: 0, B: 255, A: 255}, // Blue
			}
			DrawNumberFun(renderer, foodNum*10, (windowWidth/2)+numberOfDigits(foodNum*10)*13, (windowHeight/2)-50, colors, funCounter)
			texture.Update(nil, unsafe.Pointer(&pixels[0]), windowWidth*4)
			deathcounter++
		}
		renderer.Present()
		if deathcounter >= 300 {
			fmt.Println("Game over! Score: ", foodNum*10)
			running = false
		}
		// Delay to cap frame rate
		sdl.Delay(16)
	}
}
