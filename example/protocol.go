package example

import "github.com/dobyte/mahjong"

// SitReq 坐下请求
type SitReq struct {
	Seat int32 `json:"seat"` // 座位号
}

// ActionReq 操作请求
type ActionReq struct {
	ActionType int32 `json:"actionType"` // 操作类型
	ActionCard int32 `json:"actionCard"` // 操作的牌
}

// StartNotice 游戏开始通知
type StartNotice struct {
	Round      int32   `json:"round"`      // 当前局
	Dealer     int32   `json:"dealer"`     // 庄家座位
	Dices      int32   `json:"dices"`      // 骰子结果
	CircleWind int32   `json:"circleWind"` // 圈风
	Remaining  int32   `json:"remaining"`  // 剩余牌数
	Seat       int32   `json:"seat"`       // 座位
	Wind       int32   `json:"wind"`       // 风位
	Cards      []int32 `json:"cards"`      // 手牌
	Total      int32   `json:"total"`      // 手牌总数
}

// DrawResultNotice 抓牌通知
type DrawResultNotice struct {
	Card      int32 `json:"card"`      // 所抓的牌
	Seat      int32 `json:"seat"`      // 抓牌位置
	Remaining int32 `json:"remaining"` // 剩余牌数
	Tail      bool  `json:"tail"`      // 尾部抓牌
}

// PlayResultNotice 出牌结果通知
type PlayResultNotice struct {
	Card int32 `json:"card"` // 所出的牌
	Seat int32 `json:"seat"` // 出牌位置
}

// PatchResultNotice 补花结果通知
type PatchResultNotice struct {
	Seat      int32   `json:"seat"`      // 补花位置
	Cards     []int32 `json:"cards"`     // 手牌
	Flowers   []int32 `json:"flowers"`   // 花牌
	Remaining int32   `json:"remaining"` // 剩余牌数
	Tail      bool    `json:"tail"`      // 尾部补花
}

// ActionTipsNotice 操作提示通知
type ActionTipsNotice struct {
	ActionMask   int32                     `json:"actionMask"`   // 操作项掩码
	ActionCards  []int32                   `json:"actionCards"`  // 可操作的牌
	Countdown    int32                     `json:"countdown"`    // 操作倒计时
	PlayTingList []*mahjong.PlayTingResult `json:"playTingList"` // 打听牌列表
	DisableCards []int32                   `json:"disableCards"` // 禁止出的牌
}

// ActionResultNotice 操作结果通知
type ActionResultNotice struct {
	ActionType  int32 `json:"actionType"`  // 操作类型
	ActionCard  int32 `json:"actionCard"`  // 操作的牌
	ActionSeat  int32 `json:"actionSeat"`  // 操作位置
	ProvideSeat int32 `json:"provideSeat"` // 供牌位置
}
