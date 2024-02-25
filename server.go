package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"preteristcommentedbible/common"
	"strconv"

	"github.com/labstack/echo/v4"
)

type verse struct {
	Verse string `json:"verse"`
	Text  string `json:"text"`
}

type chapter struct {
	Chapter string  `json:"chapter"`
	Verses  []verse `json:"verses"`
}

type book struct {
	Book     string    `json:"book"`
	Chapters []chapter `json:"chapters"`
}

func main() {
	e := echo.New()

	// logger
	common.NewLogger()              // new
	e.Use(common.LoggingMiddleware) // new

	e.GET("/:bible_id/:book_id/:chapter_id", func(c echo.Context) error {

		bibleId := c.Param("bible_id")
		bookId := c.Param("book_id")
		bookFile := fmt.Sprintf("bible/%s/%s.json", bibleId, bookId)

		configFile, err := os.Open(bookFile)

		if err == nil {
			chapterId, atoiErr := strconv.Atoi(c.Param("chapter_id"))

			if atoiErr != nil {
				return c.String(http.StatusBadGateway, "Error when loading bible and chapter! "+atoiErr.Error())
			}

			var loadedBook book

			jsonParser := json.NewDecoder(configFile)
			errDecode := jsonParser.Decode(&loadedBook)

			if errDecode != nil {
				return c.String(http.StatusBadGateway, "Error when loading bible and chapter! "+errDecode.Error())
			}

			common.Logger.LogInfo().Msg(fmt.Sprintf("size %d", len(loadedBook.Chapters)))
			//
			return c.String(http.StatusOK, loadedBook.Chapters[chapterId].Verses[1].Text)

		} else {
			return c.String(http.StatusBadGateway, fmt.Sprintf("Error when loading bible and chapter! %s, %s", bookFile, err.Error()))
		}
	})
	e.Logger.Fatal(e.Start(":1323"))
}
