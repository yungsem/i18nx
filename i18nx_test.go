package i18nx

import (
	"fmt"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	i18nx, err := New("bundle")
	if err != nil {
		log.Fatal(err)
	}

	translate := i18nx.Translate("SUCCESS", "en_US")
	fmt.Println(translate)

	translate = i18nx.Translate("SUCCESS", "zh_CN")
	fmt.Println(translate)
}
