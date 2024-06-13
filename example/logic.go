package example

import (
	"github.com/dobyte/mahjong"
	"github.com/dobyte/mahjong/example/route"
)

const actionCountdown = 10

type Context interface {
	// Parse 解析请求
	Parse(req interface{}) error
	// Player 获取当前玩家
	Player() *Player
}

type RouteHandler func(ctx Context)

type Proxy interface {
	// AddRouteHandler 添加路由处理器
	AddRouteHandler(route int32, stateful bool, handler RouteHandler)
	// AddEventHandler 添加事件处理器
	AddEventHandler(event int32, handler RouteHandler)
}

type Logic struct {
	proxy   Proxy
	table   *mahjong.Table
	players []*Player
}

func NewLogic() *Logic {
	l := &Logic{}
	l.players = make([]*Player, 4)
	l.table = mahjong.NewTable(
		mahjong.WithDrawTurnDealer(),
		mahjong.WithSevenPairsDisable(),
		mahjong.WithWinMultipleDisable(),
		mahjong.WithOnceShootingMultipleWinDisable(),
		mahjong.WithChowTilePlayTileDisable(),
		mahjong.WithPlayTileChowTileDisable(),
		mahjong.WithPongTilePlayTileDisable(),
		mahjong.WithPlayTilePongTileDisable(),
		mahjong.WithPassWinAgainWinDisable(),
		mahjong.WithPassPongAgainPongDisable(),
		mahjong.WithDealHandler(l.dealHandler),
		mahjong.WithDrawHandler(l.drawHandler),
		mahjong.WithPlayHandler(l.playHandler),
		mahjong.WithPassHandler(l.passHandler),
		mahjong.WithChowHandler(l.chowHandler),
		mahjong.WithPongHandler(l.pongHandler),
		mahjong.WithKongHandler(l.kongHandler),
		mahjong.WithWinHandler(l.winHandler),
		mahjong.WithPatchHandler(l.patchHandler),
		mahjong.WithSettleHandler(l.settleHandler),
		mahjong.WithOperationAddHandler(l.operationAddHandler),
		mahjong.WithOperationRemoveHandler(l.operationRemoveHandler),
	)

	l.init()

	return l
}

func (l *Logic) init() {
	// 注册坐下请求路由
	l.proxy.AddRouteHandler(route.Sit, false, l.sitHandler)
	// 注册开始游戏路由
	l.proxy.AddRouteHandler(route.Start, true, l.startHandler)
	// 注册操作请求路由
	l.proxy.AddRouteHandler(route.Action, true, l.actionHandler)
}

// 坐下处理器
func (l *Logic) sitHandler(ctx Context) {
	req := &SitReq{}

	if err := ctx.Parse(req); err != nil {
		return
	}

	player := ctx.Player()
	player.seat = l.table.Assign(int(req.Seat))
	l.players[int(req.Seat)] = player

	player.push(route.Sit, nil)
}

// 开始游戏处理器
func (l *Logic) startHandler(ctx Context) {
	// 下一局
	l.table.Next()
	// 开始发牌
	l.table.Deal()
}

// 操作处理器
func (l *Logic) actionHandler(ctx Context) {
	req := &ActionReq{}

	if err := ctx.Parse(req); err != nil {
		return
	}

	ctx.Player().seat.Action(mahjong.Action(req.ActionType), mahjong.Card(req.ActionCard))
}

// 发牌成功
func (l *Logic) dealHandler(seat *mahjong.Seat, cards mahjong.Tiles) {
	notice := &StartNotice{
		Round:      int32(l.table.Round()),
		Dealer:     int32(l.table.Dealer().ID()),
		Dices:      int32(l.table.Dices().Result().Val()),
		CircleWind: int32(l.table.CircleWind()),
		Remaining:  int32(l.table.Cards().Count()),
		Seat:       int32(seat.ID()),
		Wind:       int32(seat.Wind()),
		Total:      int32(len(cards)),
	}

	for _, player := range l.players {
		if player == nil {
			continue
		}

		if player.seat == seat {
			notice.Cards = cards.ToInt32()
		} else {
			notice.Cards = nil
		}

		player.push(route.Start, notice)
	}
}

// 补花成功（发牌后的首轮补花）
func (l *Logic) patchHandler(seat *mahjong.Seat, cards mahjong.Tiles, flowers mahjong.Tiles, tail bool) {
	notice := &PatchResultNotice{
		Seat:      int32(seat.ID()),
		Cards:     cards.ToInt32(),
		Flowers:   flowers.ToInt32(),
		Remaining: int32(l.table.Cards().Count()),
		Tail:      tail,
	}

	for _, player := range l.players {
		if player == nil {
			continue
		}

		if player.seat == seat {
			notice.Cards = cards.ToInt32()
		} else {
			notice.Cards = nil
		}

		player.push(route.PatchResultNotice, notice)
	}
}

// 抓牌成功（抓牌时可能会抓到花牌，如果是花牌，服务器会自动发下一张牌，直到为普通牌为止）
func (l *Logic) drawHandler(seat *mahjong.Seat, card mahjong.Card, flowers mahjong.Tiles, tail bool) {
	for _, player := range l.players {
		if player.seat == seat {
			player.push(route.DrawResultNotice, &DrawResultNotice{
				Card:      int32(card),
				Seat:      int32(seat.ID()),
				Remaining: int32(l.table.Cards().Count()),
				Tail:      tail,
			})
		} else {
			player.push(route.DrawResultNotice, &DrawResultNotice{
				Seat:      int32(seat.ID()),
				Remaining: int32(l.table.Cards().Count()),
				Tail:      tail,
			})
		}
	}
}

// 出牌成功
func (l *Logic) playHandler(seat *mahjong.Seat, action mahjong.Action, card mahjong.Card) {
	l.broadcast(route.PlayResultNotice, &PlayResultNotice{
		Seat: int32(seat.ID()),
		Card: int32(card),
	})
}

// 过牌成功
func (l *Logic) passHandler(seat *mahjong.Seat, action mahjong.Action) {
	// ignore
}

// 吃牌成功
func (l *Logic) chowHandler(seat *mahjong.Seat, action mahjong.Action, card mahjong.Card, provider *mahjong.Seat) {
	// 广播操作结果
	l.broadcast(route.ActionResultNotice, &ActionResultNotice{
		ActionType:  int32(action),
		ActionSeat:  int32(seat.ID()),
		ActionCard:  int32(card),
		ProvideSeat: int32(provider.ID()),
	})
}

// 碰牌成功
func (l *Logic) pongHandler(seat *mahjong.Seat, action mahjong.Action, card mahjong.Card, provider *mahjong.Seat) {
	// 广播操作结果
	l.broadcast(route.ActionResultNotice, &ActionResultNotice{
		ActionType:  int32(action),
		ActionSeat:  int32(seat.ID()),
		ActionCard:  int32(card),
		ProvideSeat: int32(provider.ID()),
	})
}

// 杠牌成功
func (l *Logic) kongHandler(seat *mahjong.Seat, action mahjong.Action, card mahjong.Card, provider *mahjong.Seat) {
	// 广播操作结果
	l.broadcast(route.ActionResultNotice, &ActionResultNotice{
		ActionType:  int32(action),
		ActionSeat:  int32(seat.ID()),
		ActionCard:  int32(card),
		ProvideSeat: int32(provider.ID()),
	})
}

// 胡牌成功
func (l *Logic) winHandler(seat *mahjong.Seat, action mahjong.Action, card mahjong.Card, provider *mahjong.Seat) {
	// 广播操作结果
	l.broadcast(route.ActionResultNotice, &ActionResultNotice{
		ActionType:  int32(action),
		ActionSeat:  int32(seat.ID()),
		ActionCard:  int32(card),
		ProvideSeat: int32(provider.ID()),
	})
}

// 开始结算
func (l *Logic) settleHandler(table *mahjong.Table) {
	l.doSettle()
}

// 可选操作添加
func (l *Logic) operationAddHandler(seat *mahjong.Seat, operation *mahjong.Operation) {
	curr := l.toPlayer(seat)

	curr.startCountdown()

	curr.push(route.ActionTipsNotice, &ActionTipsNotice{
		ActionMask:   int32(operation.Mask),
		ActionCards:  operation.Cards,
		Countdown:    actionCountdown,
		PlayTingList: seat.AnalyzePlayTingList(),
		DisableCards: seat.DisablePlayTiles().ToInt32(),
	})
}

// 可选操作移除
func (l *Logic) operationRemoveHandler(seat *mahjong.Seat, operation *mahjong.Operation) {
	curr := l.toPlayer(seat)

	curr.stopCountdown()
}

// 结算
func (l *Logic) doSettle() {
	// TODO: 结算
}

// 座位转玩家
func (l *Logic) toPlayer(seat *mahjong.Seat) *Player {
	return l.players[seat.ID()]
}

// 广播消息
func (l *Logic) broadcast(route int32, message interface{}) {
	for _, player := range l.players {
		if player == nil {
			continue
		}

		player.push(route, message)
	}
}
