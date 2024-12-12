package captcha

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/slide"
	"log"
)

var slideBasicCapt slide.Captcha

func init() {
	builder := slide.NewBuilder(
		//slide.WithGenGraphNumber(2),
		slide.WithEnableGraphVerticalRandom(true),
	)

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	graphs, err := tiles.GetTiles()
	if err != nil {
		log.Fatalln(err)
	}

	var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
	for i := 0; i < len(graphs); i++ {
		graph := graphs[i]
		newGraphs = append(newGraphs, &slide.GraphImage{
			OverlayImage: graph.OverlayImage,
			MaskImage:    graph.MaskImage,
			ShadowImage:  graph.ShadowImage,
		})
	}

	// set resources
	builder.SetResources(
		slide.WithGraphImages(newGraphs),
		slide.WithBackgrounds(imgs),
	)

	slideBasicCapt = builder.Make()
}

type Response struct {
	CaptchaKey  string `json:"captcha_key,omitempty"`
	ImageBase64 string `json:"image_base64,omitempty"`
	TileBase64  string `json:"tile_base64,omitempty"`
	TileWidth   int    `json:"tile_width,omitempty"`
	TileHeight  int    `json:"tile_height,omitempty"`
	TileX       int    `json:"tile_x,omitempty"`
	TileY       int    `json:"tile_y,omitempty"`
}

type CheckData struct {
	CaptchaByte []byte `json:"captcha_byte"`
	CaptchaKey  string `json:"captcha_key"`
}

// GenerateCaptcha generates the captcha data.
func GenerateCaptcha() (*Response, *CheckData, error) {
	captData, err := slideBasicCapt.Generate()
	if err != nil {
		return nil, nil, err
	}

	blockData := captData.GetData()
	if blockData == nil {
		return nil, nil, errors.New("gen captcha data failed")
	}

	masterImageBase64 := captData.GetMasterImage().ToBase64()
	tileImageBase64 := captData.GetTileImage().ToBase64()
	dotsByte, _ := json.Marshal(blockData)
	key := StringToMD5(string(dotsByte))

	return &Response{
			CaptchaKey:  key,
			ImageBase64: masterImageBase64,
			TileBase64:  tileImageBase64,
			TileWidth:   blockData.Width,
			TileHeight:  blockData.Height,
			TileX:       blockData.TileX,
			TileY:       blockData.TileY,
		}, &CheckData{
			CaptchaByte: dotsByte,
			CaptchaKey:  key,
		}, nil
}

// StringToMD5 MD5
func StringToMD5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
