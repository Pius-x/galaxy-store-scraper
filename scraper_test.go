package galaxy_store_scraper

import (
	"fmt"
	"os"
	"testing"
)

func TestSamsungAppLookupOnline(t *testing.T) {
	gotData, err := SamsungAppLookup("com.igg.samsung.lordsmobile")
	if err != nil {
		log(err.Error())
		return
	}
	log("title:" + gotData.DetailMain.ContentName)
}

func TestSamsungAppLookupOffline(t *testing.T) {
	gotData, err := SamsungAppLookup("com.igg.samsung.lordsmobile123")
	if err != nil {
		log(err.Error())
		return
	}
	log("title:" + gotData.DetailMain.ContentName)
}

func log(str string) {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 写入内容到文件中
	_, err = file.WriteString(fmt.Sprintf("%s\n", str))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully!")
}
