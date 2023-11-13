package wb

import (
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/setting"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"time"
)

var sessionStore = NewSessionStore(idleSessionTimeout + 15*time.Second)

func Init() {
	u := url.URL{
		Scheme: "ws",
		Host:   setting.DefaultConfig().ServerHost,
		Path:   "/session",
	}
	flog.Info("connecting to %s", u.String())

	header := http.Header{}
	header.Set("X-AccessToken", setting.DefaultConfig().AccessToken)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		flog.Error(err)
		return
	}

	sess, count := sessionStore.NewSession(conn, "")
	flog.Info("ws: session started %s %d", sess.sid, count)

	// Do work in goroutines to return from serveWebSocket() to release file pointers.
	// Otherwise, "too many open files" will happen.
	go sess.writeLoop()
	go sess.readLoop()
}
