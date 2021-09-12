package plugin_gif

import (
	"image"
	"math/rand"
	"strconv"
	"time"

	"github.com/tdf1939/img"
)

// 爬
func (cc context) pa() string {
	tou := img.ImDc(cc.imgs[0], 0, 0).Circle(0).Im
	// 随机爬图序号
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(60) + 1
	dc := img.ImDc(downloadBy(`pa/`+strconv.Itoa(rand)+`.png`), 0, 0).
		DstOver(tou, 100, 100, 0, 400).Im
	img.SavePng(dc, cc.user+`爬.png`)
	return img.SGpic(cc.user + `爬.png`)

}

// 撕
func (cc context) si() string {
	tou := img.ImDc(cc.imgs[0], 0, 0).Im
	im1 := img.Rotate(tou, 20, 380, 380)
	im2 := img.Rotate(tou, -12, 380, 380)
	dc := img.ImDc(downloadBy(`si/0.png`), 0, 0).
		DstOver(im1.Im, im1.W, im1.H, -3, 370).
		DstOver(im2.Im, im2.W, im2.H, 653, 310).Im
	img.SavePng(dc, cc.user+`撕.png`)
	return img.SGpic(cc.user + `撕.png`)
}

// 一直
func (cc context) yiZhi() string {
	tou := img.ImsDc(cc.imgs[0], 0, 0)
	var dc []*image.NRGBA
	for _, v := range tou {
		dc = append(dc, img.ImDc(downloadBy(`xiao/0.png`), 0, 0).
			Over(v, 249, 249, 0, 0).
			Over(v, 47, 47, 140, 250).Clone().Im,
		)
	}
	img.SaveGif(img.AndGif(1, dc), cc.user+`一直.gif`)
	return img.SGpic(cc.user + `一直.gif`)
}

// 关注 必须下载素材到指定位置才可用
func (cc context) guanZhu(txt1, txt2 string) string {
	// 加载图片
	im := img.ImDc(cc.imgs[0], 0, 0)
	//叠加图片
	dst := img.ImDc(downloadBy(`guanzhu/gg.png`), 0, 0).
		DstOver(im.Im, 155, 155, 45, 45)
	dc := dst.Text(downloadBy(`font/3.ttf`), 38, img.Black, 210, 110, txt1).
		Text(downloadBy(`font/3.ttf`), 36, img.Grey, 210, 170, txt2).Im
	img.SavePng(dc, cc.user+`关注.png`)
	return img.SGpic(cc.user + `关注.png`)
}

// 简单
func (cc context) other(value ...string) string {
	// 加载图片
	im := img.ImDc(cc.imgs[0], 0, 0)
	var a *image.NRGBA
	if value[0] == "上翻" || value[0] == "下翻" {
		a = im.FlipV().Im
	} else if value[0] == "左翻" || value[0] == "右翻" {
		a = im.FlipH().Im
	} else if value[0] == "反色" {
		a = im.Invert().Im
	} else if value[0] == "灰度" {
		a = im.Grayscale().Im
	} else if value[0] == "负片" {
		a = im.Invert().Grayscale().Im
	} else if value[0] == "浮雕" {
		a = im.Convolve3x3().Im
	} else if value[0] == "打码" {
		a = im.Blur(10).Im
	} else if value[0] == "旋转" {
		r, _ := strconv.ParseFloat(value[1], 64)
		a = img.Rotate(im.Im, r, 0, 0).Im
	} else if value[0] == "变形" {
		w, _ := strconv.Atoi(value[1])
		h, _ := strconv.Atoi(value[2])
		a = img.Size(im.Im, w, h).Im
	}
	img.SavePng(a, cc.user+value[0]+`.png`)
	return img.SGpic(cc.user + value[0] + `.png`)
}
