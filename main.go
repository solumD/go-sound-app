package main

import (
	//"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	//Настройки приложения
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())

	//Верхний текст
	topText := canvas.NewImageFromFile("Изображения\\TopText.png")
	topText.SetMinSize(fyne.Size{Width: 700, Height: 50})

	//Нижний текст
	bottomText := canvas.NewImageFromFile("Изображения\\bottomText.png")
	bottomText.SetMinSize(fyne.Size{Width: 700, Height: 50})

	//Разделитель верхнего текста и поля с кнопками
	topSeparator := canvas.NewImageFromFile("Изображения\\topSeparator.png")
	topSeparator.SetMinSize(fyne.Size{Width: 700, Height: 50})

	//Разделитель нижнего текст и поля с кнопками
	bottomSeparator := canvas.NewImageFromFile("Изображения\\bottomSeparator.png")
	bottomSeparator.SetMinSize(fyne.Size{Width: 700, Height: 50})

	//Кнопки выбора АудиоКниги
	btnKniga1 := widget.NewButton("Книга 1", func() {

	})
	btnKniga2 := widget.NewButton("Книга 2", func() {

	})
	btnKniga3 := widget.NewButton("Книга 3", func() {

	})
	btnKniga4 := widget.NewButton("Книга 4", func() {

	})

	//Ряды с книгами
	firstBooksRow := container.NewGridWrap( //первый ряд
		fyne.NewSize(348, 150),
		btnKniga1,
		btnKniga2,
	)
	secondBooksRow := container.NewGridWrap( //второй ряд
		fyne.NewSize(348, 150),
		btnKniga3,
		btnKniga4,
	)

	//Контейнер с элементами главного окна
	mainContent := container.NewVBox(
		topText,
		topSeparator,
		firstBooksRow,
		secondBooksRow,
		bottomSeparator,
		bottomText,
	)

	//Иконка главного окна
	mainIcon, err := fyne.LoadResourceFromPath("Изображения\\appIcon.png")
	if err != nil {
		log.Println(err)
	}

	//Настройки главного окна
	mainW := a.NewWindow("SoundApp")
	mainW.Resize(fyne.NewSize(700, 500))
	mainW.SetIcon(mainIcon)
	mainW.SetContent(
		mainContent,
	)

	//Горутина, которая каждую секунду возращает размер окна к стандартному
	end := make(chan int)
	count := 0
	go func(w fyne.Window) {
		for {
			time.Sleep(time.Second)
			select {
			case <-end:
				log.Println("Gorutine is closed.")
				return
			default:
				w.Resize(fyne.NewSize(700, 500))
				count++
			}
		}
	}(mainW)

	//Запуск программы
	mainW.Show()
	a.Run()
	gorutineLogClosed(count, end)
}

// Выводит информацию о горутине после ее закрытия
func gorutineLogClosed(i int, end chan int) {
	log.Println("Число итераций:", i)
	end <- 1
}
