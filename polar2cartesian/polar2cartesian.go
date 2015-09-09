// polar2cartesian.go
package main

import (
	"fmt"
	"bufio"
	"math"
	"os"
	"runtime"
)

type polar struct{
	redius float64
	o float64
}

type cartersian struct{
	x ,y float64
}
var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " +
	"or %s to quit."

func init(){
	if runtime.GOOS == "windows"{
		prompt = fmt.Sprintf(prompt,"Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt,"Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions,answers)
}

func createSolver(questions chan polar) chan cartersian{
	answers := make(chan cartersian)
	go func (){
		for{
			polarCoord := <- questions
			o := polarCoord.o *math.Pi /180.0
			x := polarCoord.redius * math.Cos(o)
			y := polarCoord.redius * math.Sin(o)
			answers <- cartersian{x,y}
		}
	}()
	return answers
}

const result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"

func interact(questions chan polar,answers chan cartersian){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for{
		fmt.Printf("Radius and angle :")
		line ,err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius , o float64
		if _ , err := fmt.Sscanf(line ,"%f %f",&radius,&o);err!=nil{
			fmt.Fprintln(os.Stderr,"invalid input")
			continue
		}
		questions <- polar{radius,o}
		coord := <- answers
		fmt.Printf(result,radius,o,coord.x,coord.y)
		
	}
	fmt.Println()
}


