package galaxy_store_scraper

import (
	"errors"
	"fmt"
)

func TryWithErrorValue(callback func() error) (errorValue any, ok bool) {
	ok = true

	defer func() {
		if r := recover(); r != nil {
			ok = false
			errorValue = r
		}
	}()

	err := callback()
	if err != nil {
		ok = false
		errorValue = err
	}

	return
}

func Try(callback func() error) error {
	if r, ok := TryWithErrorValue(callback); !ok {

		var err error
		if err, ok = r.(error); !ok {
			err = errors.New(fmt.Sprintf("%v", r))
		}

		return err
	}

	return nil
}
