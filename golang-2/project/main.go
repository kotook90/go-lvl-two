
package main

import (
	"fmt"
	"newmod/remove"
	scanning "newmod/scaningFileToPath"
	sorting "newmod/sortingFile"
)

func main() {

	infoOfAllFiles, err := scanning.ListDirByWalk("/home/anton/projects/golang-2/project")

	if err != nil {
		err = fmt.Errorf("произошла ошибка выполнения! %s", err)

	}

	InfoOfCopy, err := sorting.Sorting(infoOfAllFiles)
	if err != nil {
		err = fmt.Errorf("произошла ошибка выполнения! %s", err)
	}

	copyList := sorting.DuplicateList(InfoOfCopy)

	fmt.Println("Если Вы хотите удалить дубликаты файлов, введите Y, для отмены введите N")
	var y string
	_, _ = fmt.Scan(&y)

	remove.Remove(copyList, y)
	if err != nil {
		err = fmt.Errorf("ошибка удаления дубликатов %s", err)
	}

}
