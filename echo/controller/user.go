package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewUserController(e *echo.Echo) {
	ctrl := userController{}
	g := e.Group("/user")
	g.GET("/:id", ctrl.getUser)
	g.GET("/show", ctrl.show)
	g.GET("/queryParamsBinder", ctrl.queryParamsBinder)
	g.POST("/body", ctrl.body)
	g.POST("/writeCookie", ctrl.writeCookie)
	g.GET("/readCookie", ctrl.readCookie)
	g.GET("/readAllCookie", ctrl.readAllCookie)
	g.GET("/streamJson", ctrl.streamJSON)
	g.GET("/attachment", ctrl.attachment)
	g.GET("/file", ctrl.file)
	g.GET("/inlinePng", ctrl.inlinePng)
	g.GET("/inline", ctrl.inline)
	g.GET("/stream", ctrl.stream)
	g.GET("/noContent", ctrl.noContent)
	g.GET("/redirect", ctrl.redirect)
	g.GET("/hooksBefore", ctrl.hooksBefore)
	g.GET("/hooksAfter", ctrl.hooksAfter)
	g.GET("/streamingResponse", ctrl.streamingResponse)
	g.GET("/ws", ctrl.ws)
}

type userController struct {
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type (
	Geolocation struct {
		Altitude  float64
		Latitude  float64
		Longitude float64
	}
)

var (
	locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{1899, 39.096849, -120.032351},
		{2619, 37.865101, -119.538329},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}
	upgrader = websocket.Upgrader{}
)

// 路径参数示例
func (ctrl *userController) getUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	log.Info().Int64("id", id).Msg("路径参数示例")
	return c.String(http.StatusOK, idStr)
}

// 请求参数示例
func (ctrl *userController) show(c echo.Context) error {
	idStr := c.QueryParam("id")
	name := c.QueryParam("name")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	u := new(User)
	u.ID = id
	u.Name = name
	log.Info().Dict("user", zerolog.Dict().Int64("id", u.ID).Str("name", u.Name)).Msg("请求参数示例")
	return c.JSON(http.StatusOK, u)
}

// 请求参数绑定示例
func (ctrl *userController) queryParamsBinder(c echo.Context) (err error) {
	u := new(User)
	err = echo.QueryParamsBinder(c).Int64("id", &u.ID).String("name", &u.Name).BindError()
	if err != nil {
		return err
	}
	log.Info().Dict("user", zerolog.Dict().Int64("id", u.ID).Str("name", u.Name)).Msg("请求参数绑定示例")
	return c.JSON(http.StatusOK, u)
}

// 请求body示例
func (ctrl *userController) body(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	log.Info().Dict("user", zerolog.Dict().Int64("id", u.ID).Str("name", u.Name)).Msg("请求body参数示例")
	return c.JSON(http.StatusOK, u)
}

// 写cookie
func (ctrl *userController) writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "Yazid"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

// 读取单个cookie
func (ctrl *userController) readCookie(c echo.Context) (err error) {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	log.Info().Str(cookie.Name, cookie.Value).Msg("读取单个cookie")
	return c.JSON(http.StatusOK, cookie)
}

// 读取所有cookie
func (ctrl *userController) readAllCookie(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		log.Info().Str(cookie.Name, cookie.Value).Msg("读取所有cookie")
	}
	return c.JSON(http.StatusOK, c.Cookies())
}

// 直接使用json流进行响应
func (ctrl *userController) streamJSON(c echo.Context) error {
	u := &User{
		ID:   111,
		Name: "Yazid",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(u)
}

// 下载文件，服务端命名，在浏览器中下载
func (ctrl *userController) attachment(c echo.Context) error {
	return c.Attachment("/Users/yazid/Downloads/大数据知识图谱.png", "大数据知识图谱PNG.png")
}

// 下载文件，客户端命名，默认名称为接口名称file
func (ctrl *userController) file(c echo.Context) error {
	return c.File("/Users/yazid/Downloads/组织结构.xlsx")
}

// 内联文件，在浏览器中打开
func (ctrl *userController) inlinePng(c echo.Context) error {
	return c.Inline("/Users/yazid/Downloads/大数据知识图谱.png", "大数据知识图谱PNG.png")
}

// 内联文件，无法在浏览器中打开的直接下载
func (ctrl *userController) inline(c echo.Context) error {
	return c.Inline("/Users/yazid/Downloads/组织结构.xlsx", "组织结构XLSX.xlsx")
}

// 响应流文件
func (ctrl *userController) stream(c echo.Context) error {
	f, err := os.Open("/Users/yazid/Downloads/大数据知识图谱.png")
	if err != nil {
		return err
	}
	return c.Stream(http.StatusOK, "image/png", f)
}

// 响应无正文
func (ctrl *userController) noContent(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// 重定向
// 如果不带协议，当前的写法只能在当前配置的group下做重定向
// 带了协议则直接重定向
func (ctrl *userController) redirect(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
}

// Hooks Response Before
func (ctrl *userController) hooksBefore(c echo.Context) error {
	c.Response().Before(func() {
		log.Info().Msg("Before")
	})
	return c.String(http.StatusOK, "Hooks Response Before")
}

// Hooks Response After
func (ctrl *userController) hooksAfter(c echo.Context) error {
	c.Response().After(func() {
		log.Info().Msg("After")
	})
	return c.String(http.StatusOK, "Hooks Response After")
}

// 流响应，不能和超时中间件共用
func (ctrl *userController) streamingResponse(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())
	for _, l := range locations {
		if err := enc.Encode(l); err != nil {
			return err
		}
		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}
	return nil
}

// WebSocket
func (ctrl *userController) ws(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	}(ws)

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			log.Error().Err(err).Msg("")
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Error().Err(err).Msg("")
		}
		log.Info().Str("msg", string(msg)).Msg("WS receive msg")
	}
}
