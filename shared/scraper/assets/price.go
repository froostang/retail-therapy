package assets

import (
	"fmt"
	"regexp"
)

func GetPrice(body []byte) (float64, error) {

	re := regexp.MustCompile(`"current_retail\\":(\d+\.\d+)`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return 0, fmt.Errorf("price not found")
	}

	priceStr := matches[1]
	var price float64
	_, err := fmt.Sscanf(priceStr, "%f", &price)
	if err != nil {
		return 0, err
	}

	return price, nil
}
