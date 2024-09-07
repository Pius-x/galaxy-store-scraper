package galaxy_store_scraper

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

func TestSamsungAppLookupOnline(t *testing.T) {
	gotData, err := SamsungAppLookup("com.igg.samsung.lordsmobile")
	if err != nil {
		fmt.Println(err)
		return
	}

	marshal, err := sonic.MarshalString(gotData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(marshal)
}

func TestSamsungAppLookupOffline(t *testing.T) {
	gotData, err := SamsungAppLookup("com.igg.samsung.lordsmobile123")
	if err != nil {
		fmt.Println(err)
		return
	}

	marshal, err := sonic.MarshalString(gotData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(marshal)
}
