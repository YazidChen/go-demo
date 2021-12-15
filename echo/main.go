package main

import (
	"echo/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	e := echo.New()
	// 记录访问日志
	e.Use(middleware.Logger())
	// 恢复panics并打印信息，输出到HTTPErrorHandler
	//e.Use(middleware.Recover())
	// 通过 X-Forwared-For 获取IP地址
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	// 拦截器
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Info().Str("ip", c.RealIP()).Msg("IP地址")
			if _, err := c.Cookie("username"); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "不存在cookie：username")
			}
			return next(c)
		}
	})
	// Body转储中间件,可取到传入body和响应body
	e.Use(middleware.BodyDump(func(c echo.Context, req []byte, res []byte) {
		log.Info().RawJSON("req", req).Msg("")
		log.Info().RawJSON("res", res).Msg("")
	}))
	// Body Content-Length限制
	e.Use(middleware.BodyLimit("2M"))
	// 防跨站点请求伪造
	e.Use(middleware.CSRF())
	// 解压中间件，解压gzip
	e.Use(middleware.Decompress())
	// gzip中间件，压缩响应
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	// 请求ID
	e.Use(middleware.RequestID())
	// 安全中间件
	e.Use(middleware.Secure())
	// 超时中间件，不能和流响应、websocket共用
	/*	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      nil,
		ErrorMessage: "",
		Timeout:      30 * time.Second,
	}))*/
	controller.NewUserController(e)

	log.Info().Interface("routers", e.Routes()).Msg("所有路由")

	log.Fatal().Err(e.Start(":6060"))
}
