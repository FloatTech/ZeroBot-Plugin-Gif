// Package plugin_gif 制图
package plugin_gif

import (
	"os"
	"strconv"
	"strings"

	"github.com/FloatTech/ZeroBot-Plugin/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	cmds = []string{"搓", "冲", "摸", "拍", "抬", "次", "丢", "吃", "敲", "透", "啃", "亲", "蹭",
		"爬", "撕", "一直",
		"灰度", "上翻", "下翻", "左翻", "右翻", "反色", "浮雕", "打码", "负片"}
	botpath, _ = os.Getwd()
	datapath   = botpath + "/data/gif/"
)

// var a2 = []string{"旋转", "变形"}

func init() { // 插件主体
	engine := control.Register("gif", &control.Options{
		DisableOnDefault: false,
		Help:             "制图\n- " + strings.Join(cmds, "\n- "),
	})
	engine.OnRegex(`^(` + strings.Join(cmds, "|") + `)\D*?(\[CQ:(image\,file=([0-9a-zA-Z]+).*?|at.+?(\d{5,11}))\].*|(\d+))$`).
		SetBlock(true).SetPriority(20).Handle(func(ctx *zero.Ctx) {
		c := newContext(ctx.Event.UserID)
		list := ctx.State["regex_matched"].([]string)
		c.prepareLogos(list[4]+list[5]+list[6], strconv.FormatInt(ctx.Event.UserID, 10))
		var picurl string
		switch list[1] {
		case "爬":
			picurl = c.pa()
		case "摸":
			picurl = c.mo()
		case "吃":
			picurl = c.chi()
		case "啃":
			picurl = c.ken()
		case "蹭":
			picurl = c.ceng()
		case "敲":
			picurl = c.qiao()
		case "搓":
			picurl = c.cuo()
		case "拍":
			picurl = c.pai()
		case "丢":
			picurl = c.diu()
		case "撕":
			picurl = c.si()
		case "冲":
			picurl = c.chong()
		case "一直":
			picurl = c.yiZhi()
		default:
			picurl = c.other(list[1]) // "灰度", "上翻", "下翻", "左翻", "右翻", "反色", "倒放", "浮雕", "打码", "负片"
		}
		ctx.Send(message.Image(picurl))
	})
}

func (c context) prepareLogos(s ...string) {
	for i, v := range s {
		_, err := strconv.Atoi(v)
		if err != nil {
			download("https://gchat.qpic.cn/gchatpic_new//--"+strings.ToUpper(v)+"/0", c.user+"yuan"+strconv.Itoa(i)+".gif")
		} else {
			download("http://q4.qlogo.cn/g?b=qq&nk="+v+"&s=640", c.user+"yuan"+strconv.Itoa(i)+".gif")
		}
	}
}
