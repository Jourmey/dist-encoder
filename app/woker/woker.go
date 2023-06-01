package main

import (
	"flag"
	"time"

	"dist-encoder/app/woker/internal/config"
	"dist-encoder/pb/manager"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/woker.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	for {
		time.Sleep(1 * time.Second)

		for _, job := range c.Jobs {

			err := ffmpeg.Input(job.Input).Output(job.Output, c.Args).OverWriteOutput().Run()
			if err != nil {
				logx.Error("ffmpeg failed. err:", err)
			} else {
				logx.Error("ffmpeg success. err:", err)
			}
		}

	}

}

func getFfmpegPut(put *manager.PutInfo) (filename string, kwargs ffmpeg.KwArgs) {
	if put == nil {
		return "", nil
	}

	filename = put.GetPath()
	args := put.GetKwArgs()

	kwargs = make(ffmpeg.KwArgs, len(args))
	for i := 0; i < len(args); i++ {
		kwargs[args[i].GetKey()] = args[i].GetValue()
	}

	return
}
