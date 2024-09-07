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

const PageNotFound = "404 Page Not Found"

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
				}()
				if event.RequestID == id {
					err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
						detailBody, err = network.GetResponseBody(id).Do(ctx)
						if err != nil {
							return err
						}
						if len(detailBody) == 0 {
							return fmt.Errorf("empty body")
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
					id = PageNotFound
				}
			}
		}
	})

	if err = chromedp.Run(ctx, chromedp.Navigate(lookupUrl)); err != nil {
		cancel()
		return data, err
	}

	select {
	case <-done: // finished
		cancel()
	case <-time.After(10 * time.Second): // timeout
		cancel()
		return data, errors.New("request timeout: samsung app lookup req")
	}

	if err != nil {
		return data, err
	}

	if len(detailBody) == 0 {
		if id == PageNotFound {
			return data, err
		}
		return data, fmt.Errorf("empty detail body")
	}

	if err = sonic.Unmarshal(detailBody, &data); err != nil {
		return data, err
	}

	data.DetailUrl = lookupUrl

	return data, err
}
