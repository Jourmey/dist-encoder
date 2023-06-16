package main

import (
	"context"
	"dist-encoder/app/manager/distributeclient"
	"dist-encoder/pb/distribute"
	"flag"
	"time"

	"dist-encoder/app/woker/internal/config"

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
	client := distributeclient.NewDistribute(cli)

	for {
		time.Sleep(10 * time.Second)

		res, err := client.GetVideoJob(context.Background(), &distributeclient.GetVideoJobRequest{
			Host: sysx.Hostname(),
			Ip:   "",
		})

		if err != nil {
			logx.Error("client.GetVideoJob failed. err:", err)
			continue
		}

		if res == nil || res.Job == nil {
			continue
		}

		inputKwargs, outputKwargs := getKwargs(res.ConvertCnf)

		err = ffmpeg.Input(res.Job.InPut, inputKwargs).Output(res.Job.OutPut, outputKwargs).OverWriteOutput().ErrorToStdOut().Run()
		if err != nil {
			logx.Error("ffmpeg failed. err:", err)
		} else {
			logx.Error("ffmpeg success. err:", err)
		}
	}

}

func getKwargs(cnf *distribute.ConvertCnf) (inputKwargs ffmpeg.KwArgs, outputKwargs ffmpeg.KwArgs) {

	if cnf == nil {
		return nil, nil
	}

	if len(cnf.InKwArgs) != 0 {
		inputKwargs = make(map[string]interface{}, len(cnf.InKwArgs))

		for i := 0; i < len(cnf.InKwArgs); i++ {
			inputKwargs[cnf.InKwArgs[i].Key] = cnf.InKwArgs[i].Value
		}
	}

	if len(cnf.OutKwArgs) != 0 {
		outputKwargs = make(map[string]interface{}, len(cnf.OutKwArgs))

		for i := 0; i < len(cnf.OutKwArgs); i++ {
			outputKwargs[cnf.OutKwArgs[i].Key] = cnf.OutKwArgs[i].Value
		}
	}

	return
}
