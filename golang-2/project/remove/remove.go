package remove

import (
	"fmt"
	sorting "newmod/sortingFile"
	"os"
	"strconv"
	"strings"
)

func Remove(duplicateAmount map[sorting.CopyFiles]int, yes string) {
	yes = strings.ToUpper(yes)

	for i := range duplicateAmount {

		switch yes {
		case "Y":
			{

				err := os.Remove(i.Path)
				if err != nil {
					err = fmt.Errorf("ошибка при удалении дубликатов %s", err)
					return
				}
				fmt.Println("Дубликаты файлов успешно удалены")

			}
		case "N":
			{
				fmt.Println("Отмена")
				return
			}
		case yes:
			{

				if _, err := strconv.Atoi(yes); err == nil {
					fmt.Printf("%q Вы ввели число, операция прервана", yes)
					return
				}
				if yes != "N" {
					fmt.Printf("Вы ввели неверную букву: '%s', операция прервана", yes)
					return
				} else if yes != "Y" {
					fmt.Printf("Вы ввели неверную букву: '%s', операция прервана", yes)
					return
				}

			}

		}

	}

}
