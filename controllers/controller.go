package controllers

import (
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
)

type Stock struct {
	Company       string `json:"company"`
	Price         string `json:"price"`
	Change        string `json:"change"`
	ChangePercent string `json:"changePercent"`
}

func Index(c *fiber.Ctx) error {
	symbols := c.Query("symbols")
	symbol := strings.Split(symbols, ",")
	url := "https://finance.yahoo.com/quote/"
	var stocks []Stock
	results := make(chan Stock)

	cl := colly.NewCollector()

	cl.OnHTML("div#quote-header-info", func(h *colly.HTMLElement) {
		stock := Stock{}
		stock.Company = h.ChildText("h1")
		stock.Price = h.ChildText("fin-streamer[data-field='regularMarketPrice']")
		stock.Change = h.ChildText("fin-streamer[data-field='regularMarketChange']")
		stock.ChangePercent = h.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		results <- stock
	})

	go func() {
		for _, sym := range symbol {
			cl.Visit(url + sym)
		}
		close(results)
	}()

	for stock := range results {
		stocks = append(stocks, stock)
	}

	return c.JSON(stocks)
}

func HealthCheck(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
