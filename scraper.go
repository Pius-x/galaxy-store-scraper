package galaxy_store_scraper

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// SamsungAppLookup 查询三星商店应用信息
func SamsungAppLookup(bid string) (data SamsungApp, err error) {
	lookupUrl := fmt.Sprintf("https://galaxystore.samsung.com/detail/%s?%d", bid, time.Now().Unix())
	listenPath := fmt.Sprintf("api/detail/%s", bid)

	ctx, cancel := chromedp.NewContext(context.Background())

	done := make(chan bool)

	var detailBody []byte
	var id network.RequestID
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if event, ok := ev.(*network.EventLoadingFinished); ok {
			if id == "" {
				return
			}
			go func() {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("Panic: %v\n", r)
					}
					done <- true

					cancel()
				}()
				if event.RequestID == id {
					err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
						if detailBody, err = network.GetResponseBody(id).Do(ctx); err != nil {
							return err
						}
						return nil
					}))
				}
			}()
		}

		if event, ok := ev.(*network.EventResponseReceived); ok {
			resp := event.Response
			if strings.Contains(resp.URL, listenPath) {
				id = event.RequestID
			} else if strings.Contains(resp.URL, lookupUrl) {
				if resp.Status == 404 {
					id = "request 404"
				}
			}
		}
	})

	if err = chromedp.Run(ctx, chromedp.Navigate(lookupUrl)); err != nil {
		return data, err
	}

	select {
	case <-done: // finished
	case <-time.After(10 * time.Second): // timeout
		return data, errors.New("request timeout: samsung app lookup req")
	}

	if err != nil || len(detailBody) == 0 {
		return data, err
	}

	if err = sonic.Unmarshal(detailBody, &data); err != nil {
		return data, err
	}

	return data, err
}
