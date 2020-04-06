package main

import (
	"fmt"

	"github.com/Anggabaonks/liblog"
	"github.com/Anggabaonks/libutils"
	"github.com/Anggabaonks/libredsocks"
	"github.com/Anggabaonks/ss-go/src/libshadowsocks"
)

const (
	appName = "Brainfuck Tunnel"
	appVersionName = "Shadowsocks"
	appVersionCode = "1.3.200210"

	copyrightYear = "2020"
	copyrightAuthor = "Angga baonks"
)

var (
	InterruptHandler = new(libutils.InterruptHandler)
	Redsocks = new(libredsocks.Redsocks)
)

func init() {
	InterruptHandler.Handle = func() {
		libshadowsocks.Stop()
		libredsocks.Stop(Redsocks)
		liblog.LogKeyboardInterrupt()
	}
	InterruptHandler.Start()
}

func main() {
	liblog.Header(
		[]string{
			fmt.Sprintf("%s [%s Version. %s]", appName, appVersionName, appVersionCode),
			fmt.Sprintf("(c) %s %s.", copyrightYear, copyrightAuthor),
		},
		liblog.Colors["G1"],
	)

	config := new(libshadowsocks.Config)
	defaultConfig := libshadowsocks.DefaultConfig

	libutils.JsonReadWrite(libutils.RealPath("config.json"), config, defaultConfig)

	Redsocks.Config = libredsocks.DefaultConfig
	Redsocks.Start()

	Shadowsocks := new(libshadowsocks.Shadowsocks)
	Shadowsocks.Redsocks = Redsocks
	Shadowsocks.Config = config
	Shadowsocks.Start()

	InterruptHandler.Wait()
}
