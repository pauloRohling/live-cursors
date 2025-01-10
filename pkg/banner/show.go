package banner

import (
	"fmt"
	"os"
)

const bannerFileName = "banner.txt"

func Show() {
	bannerPath := fmt.Sprintf("./%s", bannerFileName)
	file, err := os.ReadFile(bannerPath)
	if err != nil {
		return
	}
	fmt.Println(string(file))
}
