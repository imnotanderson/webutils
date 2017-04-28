package webutils

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

type Utils struct {
	client *http.Client
}

func New() *Utils {
	jar, err := cookiejar.New(nil)
	checkErr(err)
	utils := &Utils{
		client: &http.Client{
			Jar: jar,
		},
	}
	return utils
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (u *Utils) Get(url string) {
	r, err := u.client.Get(url)
	checkErr(err)
	data, err := ioutil.ReadAll(r.Body)
	checkErr(err)
	println(string(data))
}

func (u *Utils) GetBody(url string) io.ReadCloser {
	r, err := u.client.Get(url)
	checkErr(err)
	return r.Body
}

func (u *Utils) Find(body io.ReadCloser) string {
	doc, err := goquery.NewDocumentFromReader(body)
	checkErr(err)
	return doc.Find(".item_title").Text()
}
