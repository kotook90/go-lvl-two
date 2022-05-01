package sorting

import (
	"fmt"
	"log"
	"os"
)

type CopyFiles struct {
	Name string
	Size int64
	Path string
}

var LetFile []CopyFiles

func Sorting(infoOfAllFiles map[string]os.FileInfo) (InfoOfCopy []CopyFiles, err error) {

	//тут я перезаписываю все имена, размеры и пути найденных файлов из мапы в слайс. Ключ мапы это путь к файлу(string)

	for k, v := range infoOfAllFiles {
		LetFile = append(LetFile, CopyFiles{v.Name(), v.Size(), k}) //letFile это слайс, который содержит имя, размер и путь всех найденных файлов

	}
	InfoOfCopy = LetFile
	//fmt.Println(InfoOfCopy)

	return InfoOfCopy, nil
}

func DuplicateList(InfoOfCopy []CopyFiles) (duplicateAmount map[CopyFiles]int) {

	//var copyList2 []CopyFiles
	duplicateAmount = make(map[CopyFiles]int)
	for i := 0; i < len(InfoOfCopy); i++ {
		for j := i + 1; j < len(InfoOfCopy); j++ {
			if InfoOfCopy[i].Name == InfoOfCopy[j].Name && InfoOfCopy[i].Size == InfoOfCopy[j].Size {

				duplicateAmount[InfoOfCopy[j]]++
			}
		}
		//return duplicateAmount
	}

	//duplicateAmount2 = duplicateAmount

	if len(duplicateAmount) == 0 {
		fmt.Println("Дубликаты не найдены")
		os.Exit(1)
	}
	log.Printf("Найдены следующие дубликаты")

	for i, v := range duplicateAmount {
		fmt.Printf("Имя %s, Размер. %dKb - Дубликат номер: %d, Расположение дубликата: %s\n", i.Name, i.Size, v, i.Path)
	}

	return duplicateAmount
}
