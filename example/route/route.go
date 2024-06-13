package route

const (
	Sit                int32 = 1000 // 坐下
	Start              int32 = 1001 // 开始
	Action             int32 = 1002 // 操作
	PatchResultNotice  int32 = 1003 // 补花结果通知（发牌后的首轮补花）
	DrawResultNotice   int32 = 1003 // 抓牌结果通知
	PlayResultNotice   int32 = 1004 // 出牌结果通知
	ActionResultNotice int32 = 1005 // 操作结果通知
	ActionTipsNotice   int32 = 1006 // 操作提示通知
)
