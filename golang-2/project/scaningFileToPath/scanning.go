package scanning

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ListDirByWalk(path string) (infoOfAllFiles map[string]os.FileInfo, err error) {

	infoOfAllFiles = make(map[string]os.FileInfo)

	err = filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		//  возвращаем название папки
		if info.IsDir() {
			log.Printf("[Директория  %s]\n", wPath)
			return nil
		}
		// Выводится название файла рзмер и путь
		if wPath != path {

			log.Printf("Имя файла %s, Размер %d, Путь %s\n", info.Name(), info.Size(), wPath)
			infoOfAllFiles[wPath] = info
		}

		return nil

	})
	if err != nil {
		err = fmt.Errorf("ошибка рекурсивного обхода выбранной директории %s", err)
	}

	for i, v := range infoOfAllFiles {
		log.Printf("Список всех обнаруженных файлов: Имя файла %s, Размер %d Kb, Путь %s\n", v.Name(), v.Size(), i)

	}

	return infoOfAllFiles, nil
}
