package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"preteristcommentedbible/common"
	"strconv"
	"text/template"

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

// TemplateRenderer is a custom html/template renderer for Echo framework
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Static("assets", "views/dist/assets")
	e.Static("favicon.ico", "views/favicon.ico")

	renderer := &Template{
		templates: template.Must(template.ParseGlob("views/dist/*.html")),
	}

	e.Renderer = renderer

	// logger
	common.NewLogger()              // new
	e.Use(common.LoggingMiddleware) // new

	e.GET("/ui/:name", func(c echo.Context) error {
		return c.Render(http.StatusOK, c.Param("name"), "")
	})

	e.GET("/:bible_id/:book_id/:chapter_id", func(c echo.Context) error {

		bibleId := c.Param("bible_id")
		bookId := c.Param("book_id")
		bookFile := fmt.Sprintf("bible/%s/%s.json", bibleId, bookId)

		configFile, err := os.Open(bookFile)

		if err == nil {
			chapterId, atoiErr := strconv.Atoi(c.Param("chapter_id"))
			chapterId = chapterId - 1

			if chapterId < 0 {
				chapterId = 0
			}

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

			result, errorMarsh := json.Marshal(loadedBook.Chapters[chapterId])

			if errorMarsh != nil {
				common.Logger.LogInfo().Msg(fmt.Sprintf("error %s", errorMarsh.Error()))
			}

			return c.String(http.StatusOK, string(result))

		} else {
			return c.String(http.StatusBadGateway, fmt.Sprintf("Error when loading bible and chapter! %s, %s", bookFile, err.Error()))
		}
	})
	e.Logger.Fatal(e.Start(":1323"))
}
