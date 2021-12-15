package main

import (
	"fmt"
	"time"
)

func main() {
	type placeholder [5]string

	zero:= placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓ ▓",
		"▓ ▓",
		"▓▓▓",
	}
	one:= placeholder{
		"▓▓ ",
		" ▓ ",
		" ▓ ",
		" ▓ ",
		"▓▓▓",
	}
	two:= placeholder{
		"▓▓▓",
		"  ▓",
		"▓▓▓",
		"▓  ",
		"▓▓▓",
	}

	three:= placeholder{
		"▓▓▓",
		"  ▓",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}
	four:= placeholder{
		"▓ ▓",
		"▓ ▓",
		"▓▓▓",
		"  ▓",
		"  ▓",
	}
	five:= placeholder{
		"▓▓▓",
		"▓  ",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}
	six:= placeholder{
		"▓▓▓",
		"▓  ",
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
	}
	seven:= placeholder{
		"▓▓▓",
		"  ▓",
		"  ▓",
		"  ▓",
		"  ▓",
	}
	eight:= placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
	}
	nine:= placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}

	//colon := placeholder{
//		" ",
//		"░",
//		" ",
//		"░",
//		" ",
//	}

    //digits is a mutlidimentional array = digits := [rows][cols] [10][5]
	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

    //Print 9
    for n := range nine{
        fmt.Print(nine[n], " ")
        fmt.Println()
    }

    fmt.Println()

    //line = 0,1,2,3,4 (from [5]string)
    for line := range digits[0]{
        for digit := range digits{
            //multidemeetional i.e. digits[0][0]
            fmt.Print(digits[digit][line], " ")
        }
        fmt.Println()
   }

   now := time.Now()
   hour, min, sec := now.Hour(), now.Minute(), now.Second()
   fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

}







