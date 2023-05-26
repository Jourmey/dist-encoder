package main

import (
	"context"
	"flag"
	"time"

	"dist-encoder/app/manager/managerclient"
	"dist-encoder/app/woker/internal/config"
	"dist-encoder/pb/manager"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/sysx"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/woker.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	cli := zrpc.MustNewClient(c.RpcClientConf)
	client := managerclient.NewManager(cli)

	for {
		time.Sleep(1 * time.Second)

		res, err := client.GetVideoJob(context.Background(), &manager.Worker{
			Host: sysx.Hostname(),
			Ip:   "",
		})

		if err != nil {
			logx.Error("client.GetVideoJob failed. err:", err)
			continue
		}

		if res == nil || res.GetInPut() == nil || res.GetOutPut() == nil {
			continue
		}

		var (
			inputFilename, inputKwargs   = getFfmpegPut(res.GetInPut())
			outputFilename, outputKwargs = getFfmpegPut(res.GetInPut())
		)

		err = ffmpeg.Input(inputFilename, inputKwargs).Output(outputFilename, outputKwargs).OverWriteOutput().ErrorToStdOut().Run()
		if err != nil {
			logx.Error("ffmpeg failed. err:", err)
		} else {
			logx.Error("ffmpeg success. err:", err)
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
