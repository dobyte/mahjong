package example

import "github.com/dobyte/mahjong"

type Player struct {
	seat *mahjong.Seat
}

// 推送消息
func (p *Player) push(route int32, message ...interface{}) {
	// TODO: 推送消息
}

// 开始倒计时
func (p *Player) startCountdown() {
	// TODO: 开始倒计时
}

func (p *Player) stopCountdown() {
	// TODO: 停止倒计时
}
