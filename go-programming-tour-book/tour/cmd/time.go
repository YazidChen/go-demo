package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"yazidchen.com/go-programming-tour-book/tour/internal/timer"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		n := timer.GetNowTime()
		log.Printf("输出结果： %s, %d, %s", n.Format("2006-01-02 15:04:05"), n.Unix(), n.Format(time.RFC3339))
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			// 得到输入待时间字符串中存在多少空格，用以判断是什么时间格式
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				// 支持输入 YYYY-MM-DD 格式
				layout = "2006-01-02"
			}
			if space == 1 {
				// 支持输入 YYYY-MM-DD HH:mm:ss 格式
				layout = "2006-01-02 15:04:05"
			}
			location, _ := time.LoadLocation("Asia/Shanghai")
			currentTimer, err = time.ParseInLocation(layout, calculateTime, location)
			if err != nil {
				// 支持输入时间戳格式
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("输出结果：%s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "待计算时间，格式为时间戳或格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，单位："ns","us","ms","s","m","h"`)
}
