package main

import (
	"fmt"
)

func heightChecker(heights []int) int {
	var copy []int
	copy = append(copy, heights...)
	count := 0
	for i := 0; i < len(heights); i++ {
		for j := i + 1; j > 0 && j != len(heights); j-- {
			a := heights[j]
			if a < heights[j-1] {
				b := heights[j-1]
				heights[j-1] = a
				heights[j] = b
			} else {
				break
			}
		}
	}
	for i, v := range heights {
		if v != copy[i] {
			count++
		}
	}
	return count
}

func main() {
	h := []int{1, 1, 4, 2, 1, 3}
	a := []int{5, 1, 2, 3, 4}
	b := []int{2, 1, 2, 1, 1, 2, 2, 1}

	fmt.Println(heightChecker(h))
	fmt.Println(heightChecker(a))
	fmt.Println(heightChecker(b))
}

// func main() {
// 	var s, sep string

// 	for i := 0; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	counts := make(map[string]int)

// 	input := bufio.NewScanner(os.Stdin)

// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// var palette = []color.Color{color.White, color.Black}
// const (
//     whiteIndex = 0 // first color in palette
//     blackIndex = 1 // next color in palette
// )
// func main() {
//     lissajous(os.Stdout)
// }
// func lissajous(out io.Writer) {
//     const (
// cycles  = 5     // number of complete x oscillator revolutions
// res     = 0.001 // angular resolution
// size    = 100
// nframes = 64
// delay   = 8
// // image canvas covers [-size..+size]
// // number of animation frames
// // delay between frames in 10ms units
// )
// freq := rand.Float64() * 3.0 // relative frequency of y oscillator
// anim := gif.GIF{LoopCount: nframes}
// phase := 0.0 // phase difference
// for i := 0; i < nframes; i++ {
//     rect := image.Rect(0, 0, 2*size+1, 2*size+1)
//     img := image.NewPaletted(rect, palette)
//     for t := 0.0; t < cycles*2*math.Pi; t += res {
//         x := math.Sin(t)
//         y := math.Sin(t*freq + phase)
//         img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
// blackIndex)
//     }
//     phase += 0.1
//     anim.Delay = append(anim.Delay, delay)
//     anim.Image = append(anim.Image, img)
// }
// gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
// func heightChecker(heights []int) int {

// 	sorted := make([]int, len(heights))
// 	copy(sorted, heights)
// 	sort.Ints(sorted)
// 	difference := 0
// 	for i, v := range heights {
// 		if v != sorted[i] {
// 			difference++
// 		}
// 	}
// 	return difference
// }
