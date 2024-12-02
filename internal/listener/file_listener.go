package listener

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sudoku-solver/internal/handlers"
	"time"
)

const (
	rootPhoneDir = "/run/user/1000/gvfs/mtp:host=Xiaomi_Redmi_Note_10S_SKSGBUE6QO9559TO/Внутренний общий накопитель/DCIM/Screenshots/"
)

type FileListener struct {
	ImageHandler *handlers.ImageHandler
}

func NewFileListener(imageHandler *handlers.ImageHandler) *FileListener {
	return &FileListener{
		ImageHandler: imageHandler,
	}
}

func (listener *FileListener) Listen() {
	useAdbListen(listener.ImageHandler)
}

func useAdbListen(handler *handlers.ImageHandler) {
	adbDirToWatch := "/sdcard/DCIM/Screenshots/"

	// Создаем хранилище для уже существующих файлов
	existingFiles := make(map[string]struct{})

	for {
		// Получаем список файлов в директории через ADB
		cmd := exec.Command("adb", "shell", "ls", adbDirToWatch)
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			fmt.Println("failed to run ADB:", err)
			return
		}

		// Разбираем вывод (имена файлов)
		files := bytes.Split(out.Bytes(), []byte{'\n'})
		for _, file := range files {
			if len(file) > 0 {
				fileName := string(file)
				if _, found := existingFiles[fileName]; !found {
					// Файл новый, добавляем его в хранилище и обрабатываем
					existingFiles[fileName] = struct{}{}
					openFile, err := os.OpenFile(rootPhoneDir+fileName, os.O_RDONLY|os.O_CREATE, 0666)
					script := handler.GetAndroidShellScriptFromFile(openFile)

					if script == "" {
						continue
					}

					err = exec.Command("adb", "shell", script).Run()
					if err != nil {
						continue
					}
					os.Remove(rootPhoneDir + fileName)
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}

//амиксин
