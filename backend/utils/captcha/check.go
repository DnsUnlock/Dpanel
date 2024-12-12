package captcha

import (
	"encoding/json"
	"fmt"
	"github.com/wenlng/go-captcha/v2/slide"
	"strconv"
	"strings"
)

func CheckCaptcha(point string, cacheDataByte []byte) (bool, error) {
	src := strings.Split(point, ",")
	if len(src) != 2 {
		return false, fmt.Errorf("invalid point format")
	}

	var dct *slide.Block
	if err := json.Unmarshal(cacheDataByte, &dct); err != nil {
		return false, err
	}

	sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[0]), 64)
	sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[1]), 64)

	return slide.CheckPoint(int64(sx), int64(sy), int64(dct.X), int64(dct.Y), 4), nil
}
