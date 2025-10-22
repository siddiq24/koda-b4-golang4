package app

import "fmt"

func Title(title string) {
	lgt := len(title)
	space := "   "
	dashN := lgt + 6 + 10
	dash := ""
	for range dashN{
		dash += "="
	}

	ClearScreen()
	fmt.Println(dash)
	fmt.Printf("=====%s%s%s=====\n", space, title, space)
	fmt.Printf("%s\n\n",dash)
}