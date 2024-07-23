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

type advRect struct {
	p1, p2 Vec2I
	rot    float64
	cen    Vec2I
}

func (r *advRect) Center() Vec2I {
	centerX := (r.p1.x + r.p2.x) / 2
	centerY := (r.p1.y + r.p2.y) / 2
	return Vec2I{centerX, centerY}
}

// RotatePoint rotates a point around a center by an angle (in degrees)
func (point Vec2F) RotatePoint(center Vec2F, angle float64) Vec2F {
	rad := angle * (math.Pi / 180)
	translated := point.Sub(center)

	rotated := Vec2F{
		x: translated.x*math.Cos(rad) - translated.y*math.Sin(rad),
		y: translated.x*math.Sin(rad) + translated.y*math.Cos(rad),
	}

	return rotated.Add(center)
}

func (r *advRect) Set(p1, p2 Vec2I) {
	r.p1 = p1
	r.p2 = p2
	r.cen = r.Center()
}

func (r *advRect) SetWH(x1, y1, w, h int32) {
	r.p1 = Vec2I{x1, y1}
	r.p2 = r.p1.Add(Vec2I{w, h})
	r.cen = r.Center()
}

func (r *advRect) Rotate() {
	r.p1 = r.p1.ToVec2F().RotatePoint(r.cen.ToVec2F(), r.rot).ToVec2I()
	r.p2 = r.p2.ToVec2F().RotatePoint(r.cen.ToVec2F(), r.rot).ToVec2I()
}

func (r *advRect) GetSdlRect() sdl.Rect {
	r.Rotate()
	return sdl.Rect{X: r.p1.x, Y: r.p1.y, W: r.p2.x - r.p1.x, H: r.p2.y - r.p1.y}
}

type Vec2I struct {
	x int32
	y int32
}

type Vec2F struct {
	x, y float64
}

type SnekBody struct {
	rot    float64
	pos    Vec2I
	secRot float64
	dir    Vec2I
}

func (v *SnekBody) Scale(scalar int32) SnekBody {
	return SnekBody{v.rot, v.pos.Scale(scalar), -1, Vec2I{0, 0}}
}

func (v Vec2F) Add(other Vec2F) Vec2F {
	return Vec2F{v.x + other.x, v.y + other.y}
}

func (v Vec2F) Sub(other Vec2F) Vec2F {
	return Vec2F{v.x - other.x, v.y - other.y}
}

// Convert Vec2I to Vec2F for precise calculations
func (v Vec2I) ToVec2F() Vec2F {
	return Vec2F{x: float64(v.x), y: float64(v.y)}
}

// Convert Vec2F back to Vec2I, rounding to the nearest integer
func (v Vec2F) ToVec2I() Vec2I {
	return Vec2I{x: int32(math.Round(v.x)), y: int32(math.Round(v.y))}
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
	Body             []SnekBody
	currentThickness float32
}

func (s *Snake) getRects() []sdl.Rect {
	var rects []sdl.Rect
	for _, rect := range s.Body {
		//fmt.Println(rect.secRot)
		rects = append(rects, getRotated(s.currentThickness, rect)...)
	}
	return rects
}

func getRotated(currentThickness float32, sb SnekBody) []sdl.Rect {
	pos := sb.pos
	ang1 := sb.rot
	ang2 := sb.secRot
	thick := int32(snakeSize * currentThickness)
	missingThick := (snakeSize - thick) / 2

	rotated := advRect{}
	rotated.SetWH((pos.x*snakeSize)+int32(missingThick), pos.y*snakeSize, thick, snakeSize)

	if ang2 != -1 && currentThickness != 1 {
		//fmt.Println("ang1: ", ang1, " ang2: ", ang2)
		//fmt.Println("rotating2")
		rotated1 := advRect{}
		rotated1.SetWH((pos.x*snakeSize)+int32(missingThick)-1, pos.y*snakeSize, thick, snakeSize-missingThick)
		//rotated1.cen = rotated.Center()

		rotated2 := advRect{}
		rotated2.SetWH((pos.x*snakeSize)+int32(missingThick), pos.y*snakeSize, thick, missingThick+1)
		//rotated2.cen = rotated.Center()
		fmt.Println(rotated1.GetSdlRect(), rotated2.GetSdlRect())
		//180
		//90
		//270 = 0,1
		//0
		//whackyAdd := Vec2I{0, 0}
		switch ang1 {
		case 0:
			rotated1.cen = rotated.cen.Add(Vec2I{-1, 0})
		case 90:
			rotated1.cen = rotated.cen.Add(Vec2I{0, -1})
		case 180:
			rotated1.cen = rotated.cen.Add(Vec2I{1, 0})
		case 270:
			//whackyAdd = Vec2I{0, 1}
		case 360:
			rotated1.cen = rotated.cen.Add(Vec2I{-1, 0})
		}
		rotated1.cen = rotated.cen //.Add(Vec2I{0, 1})
		rotated2.cen = rotated.cen
		fmt.Println(ang1 + 180)
		rotated1.rot = float64(int32(ang1+180) % 360)
		rotated2.rot = ang2
		//fmt.Println(rotated1.GetSdlRect(), rotated2.GetSdlRect())
		return []sdl.Rect{rotated1.GetSdlRect(), rotated2.GetSdlRect()}
	}

	rotated.rot = ang1
	//fmt.Println(rotated.GetSdlRect())
	return []sdl.Rect{rotated.GetSdlRect()}
}

func calculateAngleWithUp(vec Vec2F) float64 {
	upVec := Vec2F{0, -1} // Up vector
	dotProduct := vec.x*upVec.x + vec.y*upVec.y
	magnitudeVec := math.Sqrt(vec.x*vec.x + vec.y*vec.y)
	magnitudeUpVec := math.Sqrt(upVec.x*upVec.x + upVec.y*upVec.y)
	angleRadians := math.Acos(dotProduct / (magnitudeVec * magnitudeUpVec))

	// Adjust for 360 degrees
	if vec.x < 0 {
		angleRadians = 2*math.Pi - angleRadians
	}

	angleDegrees := angleRadians * (180 / math.Pi)
	fmt.Println(angleDegrees)
	return angleDegrees
}

func (s *Snake) Move(dir Vec2I) {
	if len(s.Body) == 0 {
		return
	}

	// Move the head
	head := s.Body[0]
	var newX = head.pos.x + int32(dir.x)
	var newY = head.pos.y + int32(dir.y)
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
	rot := calculateAngleWithUp(Vec2F{float64(dir.x), float64(dir.y)})
	if head.rot != rot {
		fmt.Println("rotating")
		s.Body[0].secRot = rot
	}
	newHead := SnekBody{rot, Vec2I{newX, newY}, -1, dir}
	s.Body = append([]SnekBody{newHead}, s.Body[:len(s.Body)-1]...)
}

func (s *Snake) CollidesWith(rect sdl.Rect) bool {
	head := s.Body[0]
	if head.pos.x <= rect.X+rect.W && head.pos.x+snakeSize >= rect.X && head.pos.y <= rect.Y+rect.H && head.pos.y+snakeSize >= rect.Y {
		return true
	}
	return false
}

func (s *Snake) CollidesWithGlobSpace(rect sdl.Rect) bool {
	head := s.Body[0]
	head = head.Scale(snakeSize)
	if head.pos.x <= rect.X+rect.W && head.pos.x+snakeSize >= rect.X && head.pos.y <= rect.Y+rect.H && head.pos.y+snakeSize >= rect.Y {
		return true
	}
	return false
}

func (s *Snake) CollidesWithSelf() bool {
	head := s.Body[0]
	for i := 1; i < len(s.Body); i++ {
		if head.pos.x == s.Body[i].pos.x && head.pos.y == s.Body[i].pos.y {
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
	if s.currentThickness > 0.3 {
		s.currentThickness -= .05
	}
	tail := s.Body[len(s.Body)-1]
	newTail := SnekBody{tail.rot, tail.pos, -1, tail.dir}
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
	Snake := Snake{Body: []SnekBody{{0.0, Vec2I{x: 0, y: 0}, -1, Vec2I{-3, -3}}}, currentThickness: 1}
	//var moveSpeed float64 = 0.1
	var moveDir Vec2I = Vec2I{0, 0}
	var count = 0

	food := getRandomFood(&Snake)
	foodNum := 0
	death := false
	pause := true
	deathcounter := 0
	funCounter := 0
	test := advRect{}
	test.SetWH(400, 400, 50, 10)
	testC := 0.0
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
			test.rot = testC + 90
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
			rec := test.GetSdlRect()
			renderer.FillRect(&rec)
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
