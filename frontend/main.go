package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"syscall/js"
	"time"

	attributes "github.com/mdigger/goldmark-attributes"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"

	"github.com/nobonobo/spago"
	"github.com/nobonobo/spago-slides/frontend/actions"
	"github.com/nobonobo/spago-slides/frontend/slide"
	"github.com/nobonobo/spago/dispatcher"
	"github.com/nobonobo/spago/jsutil"
)

var (
	document    = js.Global().Get("document")
	location    = js.Global().Get("location")
	console     = js.Global().Get("console")
	currentPage = 0
	maxPage     = 0
)

var rep = regexp.MustCompile(`\n{.*}\n\n`)

func parseMarkdown(content string) []spago.Component {
	md := goldmark.New(
		attributes.Enable,
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
			),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)
	log.Println(md)
	res := []spago.Component{}
	for id, chunk := range strings.Split(content, "\n====\n") {
		chunk = strings.Trim(chunk, "\r\n\t ")
		chunk = rep.ReplaceAllStringFunc(chunk, func(m string) string {
			parts := rep.FindStringSubmatch(m)
			fmt.Printf("%#v\n", parts)
			return m[:len(m)-1]
		})
		var output bytes.Buffer
		if err := md.Convert([]byte(chunk), &output); err != nil {
			log.Fatal(err)
		}
		log.Println(output.String())
		res = append(res, &slide.Slide{
			ID:      fmt.Sprintf("page%d", id+1),
			Content: output.String(),
		})
	}
	return res
}

func getFragments() js.Value {
	n := getCurrentPage()
	return document.Call("querySelectorAll", fmt.Sprintf("#page%d .fragment", n))
}

func prevAction(args ...interface{}) {
	currentPage--
	if currentPage < 1 {
		currentPage = 1
		return
	}
	setCurrentPage(currentPage)
}

func nextAction(args ...interface{}) {
	fragments := getFragments()
	if fragments.Length() > 0 {
		fragments.Index(0).Get("classList").Call("remove", "fragment")
		return
	}
	currentPage++
	if currentPage > maxPage {
		currentPage = maxPage
		return
	}
	setCurrentPage(currentPage)
}

func keyHandler(ev js.Value) {
	console.Call("log", ev)
	switch ev.Get("code").String() {
	case "ArrowLeft":
		prevAction()
	case "ArrowRight":
		nextAction()
	case "KeyR":
		go dispatcher.Dispatch(actions.ReLoad)
	}
}

func setCurrentPage(n int) {
	location.Set("hash", fmt.Sprintf("#page%d", n))
}

func getCurrentPage() int {
	page, err := strconv.Atoi(strings.TrimPrefix(location.Get("hash").String(), "#page"))
	if err != nil {
		page = 1
	}
	return page
}

func loadContent() []spago.Component {
	resp, err := jsutil.Fetch(time.Now().Format("/content.md?20060102-150405"), nil)
	if err != nil {
		log.Print(err)
	}
	content, err := jsutil.Await(resp.Call("text"))
	if err != nil {
		log.Print(err)
	}
	slides := parseMarkdown(content.String())
	maxPage = len(slides)
	return slides
}

func main() {
	go func() {
		top := &slide.Slides{
			Slides: loadContent(),
		}
		spago.RenderBody(top)
		dispatcher.Register(actions.PrevStep, prevAction)
		dispatcher.Register(actions.NextStep, nextAction)
		dispatcher.Register(actions.ReLoad, func(args ...interface{}) {
			log.Println("reloading...")
			top.Slides = loadContent()
			spago.RenderBody(top)
		})
		currentPage = getCurrentPage()
		location.Set("hash", "")
		setCurrentPage(currentPage)
		log.Println("currentPage:", currentPage)
		jsutil.Bind(js.Global(), "keydown", keyHandler)
		log.Println("load completed")
	}()
	select {}
}
