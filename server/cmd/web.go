package cmd

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"bbs-go/controller"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"

	"bbs-go/util/config"

	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/sirupsen/logrus"
)

func Web() {

	//engine := gin.Default()
	//controller.Setup(engine, cors)
	//return engine.Run(":" + viper.GetString("base.port"))

	app := gin.Default()
	//app.Logger().SetLevel("warn")
	//	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(limits.RequestSizeLimiter(10 * 1024 * 1024))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodHead, http.MethodDelete, http.MethodPut},
		AllowHeaders:     []string{"*"},
	}))

	controller.Router(app)

	server := &http.Server{Addr: ":" + config.Instance.Port}
	handleSignal(server)
	err := app.Run(":" + config.Instance.Port)
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logrus.Infof("got signal [%s], exiting now", s)
		if err := server.Close(); nil != err {
			logrus.Errorf("server close failed: " + err.Error())
		}

		simple.CloseDB()

		logrus.Infof("Exited")
		os.Exit(0)
	}()
}
