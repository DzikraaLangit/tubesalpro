package main
import ("fmt")

func main(){
	var logreg int
	var usn, pass string
	var menu int

	early(&logreg)
	if logreg == 1 {
		login(usn, pass)
	} else if logreg == 2 {
		register(&usn, &pass)
		login(usn, pass)
	}
	menu=1
	for menu!=0{
	inter(&menu)
	fmt.Scan(&menu)
	fmt.Scanln()
	clearScreen()
	inmenu(menu)
	if menu != 0{
		backtomenu()
	}
	}
}

