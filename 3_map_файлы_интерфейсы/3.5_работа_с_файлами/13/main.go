package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Открываем архив
	zipFile, err := os.Open("task.zip")
	if err != nil {
		fmt.Println("Ошибка при открытии архива:", err)
		return
	}
	defer zipFile.Close()

	// Получаем информацию о файле архива
	fileInfo, err := zipFile.Stat()
	if err != nil {
		fmt.Println("Ошибка при получении информации о файле архива:", err)
		return
	}

	// Читаем архив
	zipReader, err := zip.NewReader(zipFile, fileInfo.Size())
	if err != nil {
		fmt.Println("Ошибка при чтении архива:", err)
		return
	}

	// Пройдемся по всем файлам в архиве
	for _, file := range zipReader.File {
		if !file.FileInfo().IsDir() && filepath.Ext(file.Name) == ".txt" {
			// Открываем файл в архиве
			rc, err := file.Open()
			if err != nil {
				fmt.Println("Ошибка при открытии файла в архиве:", err)
				continue
			}
			defer rc.Close()

			// Читаем CSV файл
			reader := csv.NewReader(rc)
			records, err := reader.ReadAll()
			if err != nil {
				fmt.Println("Ошибка при чтении CSV файла:", err)
				continue
			}

			// Проверяем, что файл содержит нужное количество строк и столбцов
			if len(records) > 4 && len(records[4]) > 2 {
				// Выводим значение на 5 строке и 3 позиции (индексы 4 и 2)
				fmt.Println("Значение на 5 строке и 3 позиции:", records[4][2])
				return
			}
		}
	}

	fmt.Println("Непустой CSV файл не найден")
}
