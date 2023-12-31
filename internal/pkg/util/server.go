package util

import (
	"fmt"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/go-resty/resty/v2"
	"time"
)

func CheckSingleton() {
	if !PortAvailable(constant.EmbedServerPort) {
		resp, err := resty.New().SetTimeout(500 * time.Millisecond).R().
			Get(fmt.Sprintf("http://127.0.0.1:%s/", constant.EmbedServerPort))
		if err != nil {
			flog.Error(err)
			return
		}
		if resp.String() == "ok" {
			flog.Fatal("app exists")
		}
	}
}
