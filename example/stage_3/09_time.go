/*
 * @Descripttion: time 时间处理 时间戳 格式化
 * @Author: tacks321@qq.com
 * @Date: 2021-02-20 14:36:09
 * @LastEditTime: 2021-02-20 15:02:42
 */
package main

import (
	"fmt"
	"time"
)

var P = fmt.Println

func main() {
	P("=============================【 时间格式化处理 】======================================")
	now1 := time.Now()

	// 当前时间
	P(time.Now())

	// 提供年月日时分秒以及时区
	P(time.Date(2021, 02, 20, 20, 20, 20, 1, time.UTC))

	// 时间的各个部分
	P(time.Now().Year())
	P(time.Now().Month())
	P(time.Now().Day())
	P(time.Now().Weekday())
	P(time.Now().Hour())
	P(time.Now().Minute())
	P(time.Now().Second())
	P(time.Now().Nanosecond())
	P(time.Now().Location())

	now2 := time.Now()
	// 三个比较
	P(now1.Before(now2))
	P(now1.After(now2))
	P(now1.Equal(now2))
	// 时间差
	P(now2.Sub(now1))
	// 向后移动时间
	P(now2.Add(1000000 * now2.Sub(now1)))
	// 向前移动时间
	P(now2.Add(-1000000 * now2.Sub(now1)))

	P("=============================【 时间戳 】======================================")
	// 当前时间
	nowTime := time.Now()
	// 秒
	nowSecs := nowTime.Unix()
	// 纳秒
	nowNanos := nowTime.UnixNano()
	P(nowTime)
	P(nowSecs)
	P(nowNanos)
	// 毫秒 需要转化一下
	nowMillis := nowNanos / 1000000
	P(nowMillis)

	// 转化为对应的时间
	P(time.Unix(nowSecs, 0))
	P(time.Unix(0, nowNanos))
}
