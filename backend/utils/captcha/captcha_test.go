package captcha

import "testing"

func TestCaptcha(t *testing.T) {
	res, check, err := GenerateCaptcha()
	if err != nil {
		t.Error(err)
	}
	t.Log(res, check)
}
