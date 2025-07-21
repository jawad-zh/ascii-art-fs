package asciiartoutput

import (
	"os"
)

func Count()int{
	var count int
	if len(os.Args)== 4{
		count = 3
	}else if len(os.Args)==3{
		count = 2
	}else{
		count = 1
	}
	return count
}