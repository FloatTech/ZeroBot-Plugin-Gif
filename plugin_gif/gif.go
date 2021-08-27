package plugin_gif

import (
	"image"

	"github.com/tdf1939/img"
	// "bot/img" // 基础词库
)

//摸
func (cc Paths) Mo() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	mo := []*image.NRGBA{
		img.ImDc(DLSc(`mo/0.png`), 0, 0).DstOver(tou, 80, 80, 32, 32).Im,
		img.ImDc(DLSc(`mo/1.png`), 0, 0).DstOver(tou, 70, 90, 42, 22).Im,
		img.ImDc(DLSc(`mo/2.png`), 0, 0).DstOver(tou, 75, 85, 37, 27).Im,
		img.ImDc(DLSc(`mo/3.png`), 0, 0).DstOver(tou, 85, 75, 27, 37).Im,
		img.ImDc(DLSc(`mo/4.png`), 0, 0).DstOver(tou, 90, 70, 22, 42).Im,
	}
	img.SaveGif(img.AndGif(1, mo), cc.User+`摸.gif`)
	return img.SGpic(cc.User + `摸.gif`)
}

//摸
func (cc Paths) Cuo() string {
	tou := img.ImDc(cc.Pngs[0], 110, 110).Circle(0).Im
	m1 := img.Rotate(tou, 72, 0, 0)
	m2 := img.Rotate(tou, 144, 0, 0)
	m3 := img.Rotate(tou, 216, 0, 0)
	m4 := img.Rotate(tou, 288, 0, 0)
	cuo := []*image.NRGBA{
		img.ImDc(DLSc(`cuo/0.png`), 0, 0).DstOverC(tou, 0, 0, 75, 130).Im,
		img.ImDc(DLSc(`cuo/1.png`), 0, 0).DstOverC(m1.Im, 0, 0, 75, 130).Im,
		img.ImDc(DLSc(`cuo/2.png`), 0, 0).DstOverC(m2.Im, 0, 0, 75, 130).Im,
		img.ImDc(DLSc(`cuo/3.png`), 0, 0).DstOverC(m3.Im, 0, 0, 75, 130).Im,
		img.ImDc(DLSc(`cuo/4.png`), 0, 0).DstOverC(m4.Im, 0, 0, 75, 130).Im,
	}
	img.SaveGif(img.AndGif(5, cuo), cc.User+`搓.gif`)
	return img.SGpic(cc.User + `搓.gif`)
}

//敲
func (cc Paths) Qiao() string {
	tou := img.ImDc(cc.Pngs[0], 40, 40).Circle(0).Im
	qiao := []*image.NRGBA{
		img.ImDc(DLSc(`qiao/0.png`), 0, 0).Over(tou, 40, 33, 57, 52).Im,
		img.ImDc(DLSc(`qiao/1.png`), 0, 0).Over(tou, 38, 36, 58, 50).Im,
	}
	img.SaveGif(img.AndGif(1, qiao), cc.User+`敲.gif`)
	return img.SGpic(cc.User + `敲.gif`)

}

//吃
func (cc Paths) Chi() string {
	tou := img.ImDc(cc.Pngs[0], 32, 32).Im
	chi := []*image.NRGBA{
		img.ImDc(DLSc(`chi/0.png`), 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
		img.ImDc(DLSc(`chi/1.png`), 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
		img.ImDc(DLSc(`chi/2.png`), 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
	}
	img.SaveGif(img.AndGif(1, chi), cc.User+`吃.gif`)
	return img.SGpic(cc.User + `吃.gif`)

}

//蹭
func (cc Paths) Ceng() string {
	tou := img.ImDc(cc.Pngs[0], 100, 100).Circle(0).Im
	tou2 := img.ImDc(cc.Pngs[1], 100, 100).Circle(0).Im
	ceng := []*image.NRGBA{
		img.ImDc(DLSc(`ceng/0.png`), 0, 0).Over(tou, 75, 77, 40, 88).Over(tou2, 77, 103, 102, 81).Im,
		img.ImDc(DLSc(`ceng/1.png`), 0, 0).Over(tou, 75, 77, 46, 100).Over(img.Rotate(tou2, 10, 62, 127).Im, 0, 0, 92, 40).Im,
		img.ImDc(DLSc(`ceng/2.png`), 0, 0).Over(tou, 75, 77, 67, 99).Over(tou2, 76, 117, 90, 8).Im,
		img.ImDc(DLSc(`ceng/3.png`), 0, 0).Over(tou, 75, 77, 52, 83).Over(img.Rotate(tou2, -40, 94, 94).Im, 0, 0, 53, -20).Im,
		img.ImDc(DLSc(`ceng/4.png`), 0, 0).Over(tou, 75, 77, 56, 110).Over(img.Rotate(tou2, -66, 132, 80).Im, 0, 0, 78, 40).Im,
		img.ImDc(DLSc(`ceng/5.png`), 0, 0).Over(tou, 75, 77, 62, 102).Over(tou2, 71, 100, 110, 94).Im,
	}
	img.SaveGif(img.AndGif(8, ceng), cc.User+`蹭.gif`)
	// return "file:///" + cc.User + `蹭.gif`
	return img.SGpic(cc.User + `蹭.gif`)
}


//啃
func (cc Paths) Ken() string {
	tou := img.ImDc(cc.Pngs[0], 100, 100).Im
	ken := []*image.NRGBA{
		img.ImDc(DLSc(`ken/0.png`), 0, 0).DstOver(tou, 90, 90, 105, 150).Im,
		img.ImDc(DLSc(`ken/1.png`), 0, 0).DstOver(tou, 90, 83, 96, 172).Im,
		img.ImDc(DLSc(`ken/2.png`), 0, 0).DstOver(tou, 90, 90, 106, 148).Im,
		img.ImDc(DLSc(`ken/3.png`), 0, 0).DstOver(tou, 88, 88, 97, 167).Im,
		img.ImDc(DLSc(`ken/4.png`), 0, 0).DstOver(tou, 90, 85, 89, 179).Im,
		img.ImDc(DLSc(`ken/5.png`), 0, 0).DstOver(tou, 90, 90, 106, 151).Im,
		img.ImDc(DLSc(`ken/6.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/7.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/8.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/9.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/10.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/11.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/12.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/13.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/14.png`), 0, 0).Im,
		img.ImDc(DLSc(`ken/15.png`), 0, 0).Im,
	}
	img.SaveGif(img.AndGif(7, ken), cc.User+`啃.gif`)
	return img.SGpic(cc.User + `啃.gif`)
}

//拍
func (cc Paths) Pai() string {
	tou := img.ImDc(cc.Pngs[0], 30, 30).Circle(0).Im
	pai := []*image.NRGBA{
		img.ImDc(DLSc(`pai/0.png`), 0, 0).Over(tou, 0, 0, 1, 47).Im,
		img.ImDc(DLSc(`pai/1.png`), 0, 0).Over(tou, 0, 0, 1, 67).Im,
	}
	img.SaveGif(img.AndGif(1, pai), cc.User+`拍.gif`)
	return img.SGpic(cc.User + `拍.gif`)

}

//冲
func (cc Paths) Chong() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	chong := []*image.NRGBA{
		img.ImDc(DLSc(`xqe/0.png`), 0, 0).Over(tou, 30, 30, 15, 53).Im,
		img.ImDc(DLSc(`xqe/1.png`), 0, 0).Over(tou, 30, 30, 40, 53).Im,
	}
	img.SaveGif(img.AndGif(1, chong), cc.User+`冲.gif`)
	return img.SGpic(cc.User + `冲.gif`)

}

//丢
func (cc Paths) Diu() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	diu := []*image.NRGBA{
		img.ImDc(DLSc(`diu/0.png`), 0, 0).Over(tou, 32, 32, 108, 36).Im,
		img.ImDc(DLSc(`diu/1.png`), 0, 0).Over(tou, 32, 32, 122, 36).Im,
		img.ImDc(DLSc(`diu/2.png`), 0, 0).Im,
		img.ImDc(DLSc(`diu/3.png`), 0, 0).Over(tou, 123, 123, 19, 129).Im,
		img.ImDc(DLSc(`diu/4.png`), 0, 0).Over(tou, 185, 185, -50, 200).Over(tou, 33, 33, 289, 70).Im,
		img.ImDc(DLSc(`diu/5.png`), 0, 0).Over(tou, 32, 32, 280, 73).Im,
		img.ImDc(DLSc(`diu/6.png`), 0, 0).Over(tou, 35, 35, 259, 31).Im,
		img.ImDc(DLSc(`diu/7.png`), 0, 0).Over(tou, 175, 175, -50, 220).Im,
	}
	img.SaveGif(img.AndGif(7, diu), cc.User+`丢.gif`)
	return img.SGpic(cc.User + `丢.gif`)
}
