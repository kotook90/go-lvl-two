package dir

import (
	"fmt"
	"os"
	"strconv"
)

func Dir() string {
	var path string
	path = "/newfolder"
	err1 := fmt.Errorf("Не могу определить текущую директорию, ")
	err2 := fmt.Errorf("Не могу создать новую папку, ")

	//Определяю текущую директорию и в ней создаю новую папку(гарантированно пустую)
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err1, err)
		os.Exit(1)
	}
	resultPath := pwd + path
	err = os.Mkdir(resultPath, 999)
	if err != nil {
		fmt.Println(err2, err)
		os.Exit(1)
	}
	return resultPath
}
func Create(resultPath string, n int) {

	err3 := fmt.Errorf("Не могу создать новый файл, ")
	err4 := fmt.Errorf("Не могу закрыть файл, ")
	file, err := os.Create(resultPath + "/" + strconv.Itoa(n))
	if err != nil {
		fmt.Println(err3, err)
		os.Exit(1)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err4, err)
			os.Exit(1)

		}
	}()

}
