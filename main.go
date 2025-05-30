package main
import ("fmt")

func main(){
	var logreg int
	var usn, pass string
	var menu int
	var data tabCatatan
	var lastID int

	early(&logreg)
	if logreg == 1 {
		if !login(usn, pass) {
			early(&logreg)
			if logreg == 2 {
				register(&usn, &pass)
				login(usn, pass)
			}
		} 
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
	inmenu(menu, &data, &lastID)
	if menu != 0{
		backtomenu()
	}
	}
}

