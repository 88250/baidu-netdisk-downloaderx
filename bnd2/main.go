package main

import (
	"encoding/json"
	"github.com/b3log/gulu"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/b3log/bnd2/command"
	"github.com/b3log/bnd2/util"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

var logger *gulu.Logger

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	gulu.Log.SetLevel("error")
	logger = gulu.Log.NewLogger(os.Stdout)

	if util.CheckUpgrade() {
		logger.Fatalf("current BND2 kernel v[%s] is outdated, exited", util.Ver)
	}

	go util.ParentExited()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/status", func(c *gin.Context) {
		result := util.NewResult()
		defer c.JSON(http.StatusOK, result)

		result.Data = map[string]interface{}{
			"version": util.Ver,
		}
	})

	m := melody.New()
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.POST("/login", func(c *gin.Context) {
		result := util.NewResult()
		defer c.JSON(http.StatusOK, result)

		arg := map[string]interface{}{}
		if err := c.BindJSON(&arg); nil != err {
			result.Code = -1
			result.Msg = "parses login request failed"

			return
		}

		util.BDUSS = arg["bd"].(string)
		logger.Debug("signed in [baidu]")

		go util.StartAria2()
	})

	m.HandleConnect(func(s *melody.Session) {
		util.SetPushChan(s)
		logger.Debug("websocket connected")
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		logger.Infof("request [%s]", msg)
		request := map[string]interface{}{}
		if err := json.Unmarshal(msg, &request); nil != err {
			result := util.NewResult()
			result.Code = -1
			result.Msg = "Bad Request"
			responseData, _ := json.Marshal(result)
			util.Push(responseData)

			return
		}

		cmdStr := request["cmd"].(string)
		cmd := command.Commands[cmdStr]
		if nil == cmd {
			result := util.NewResult()
			result.Code = -1
			result.Msg = "Invalid Command"
			responseData, _ := json.Marshal(result)
			util.Push(responseData)

			return
		}

		param := request["param"].(map[string]interface{})
		go cmd.Exec(param)
	})

	port := strconv.Itoa(util.ServerPort)
	logger.Infof("BND2 kernel (v%s) is running [:%s]", util.Ver, port)

	if err := r.Run(":" + port); nil != err {
		logger.Errorf("start BND2 kernel failed [%s]", err)
	}
}
