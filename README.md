## 全能麻将类库

### 一、功能介绍

- 支持任意类型的服务器框架
- 支持普通和血战到底两种玩法
- 支持一炮一响和一炮多响两种玩法
- 支持2~4人玩法
- 支持牌库配牌
- 支持一炮一响玩法
- 支持一炮多响玩法
- 支持一胡结算玩法
- 支持血战到底玩法
- 支持吃牌玩法的定制
- 支持打张吃张玩法的定制
- 支持吃张打张玩法的定制
- 支持打张碰张玩法的定制
- 支持碰张打张玩法的定制
- 支持过碰不碰玩法的定制
- 支持过胡不胡玩法的定制
- 支持杠牌尾部摸牌玩法的定制
- 支持牌库任意花色的定制
- 支持流局轮庄模式定制
- 支持自动补花

### 一、基础牌型

#### 1.万牌（36张）

```text
mahjong.CharactersOne        mahjong.Card = 1  // 一万
mahjong.CharactersTwo        mahjong.Card = 2  // 二万
mahjong.CharactersThree      mahjong.Card = 3  // 三万
mahjong.CharactersFour       mahjong.Card = 4  // 四万
mahjong.CharactersFive       mahjong.Card = 5  // 五万
mahjong.CharactersSix        mahjong.Card = 6  // 六万
mahjong.CharactersSeven      mahjong.Card = 7  // 七万
mahjong.CharactersEight      mahjong.Card = 8  // 八万
mahjong.CharactersNine       mahjong.Card = 9  // 九万
```

#### 2.条牌（36张）

```text
mahjong.BamboosOne           mahjong.Card = 11 // 一条
mahjong.BamboosTwo           mahjong.Card = 12 // 二条
mahjong.BamboosThree         mahjong.Card = 13 // 三条
mahjong.BamboosFour          mahjong.Card = 14 // 四条
mahjong.BamboosFive          mahjong.Card = 15 // 五条
mahjong.BamboosSix           mahjong.Card = 16 // 六条
mahjong.BamboosSeven         mahjong.Card = 17 // 七条
mahjong.BamboosEight         mahjong.Card = 18 // 八条
mahjong.BamboosNine          mahjong.Card = 19 // 九条
```

#### 3.筒牌（36张）

```text
mahjong.DotsOne              mahjong.Card = 21 // 一筒
mahjong.DotsTwo              mahjong.Card = 22 // 二筒
mahjong.DotsThree            mahjong.Card = 23 // 三筒
mahjong.DotsFour             mahjong.Card = 24 // 四筒
mahjong.DotsFive             mahjong.Card = 25 // 五筒
mahjong.DotsSix              mahjong.Card = 26 // 六筒
mahjong.DotsSeven            mahjong.Card = 27 // 七筒
mahjong.DotsEight            mahjong.Card = 28 // 八筒
mahjong.DotsNine             mahjong.Card = 29 // 九筒
```

#### 4.风牌（16张）

```text
mahjong.WindsEast            mahjong.Card = 31 // 东
mahjong.WindsSouth           mahjong.Card = 32 // 南
mahjong.WindsWest            mahjong.Card = 33 // 西
mahjong.WindsNorth           mahjong.Card = 34 // 北
```

#### 5.箭牌（12张）

```text
mahjong.DragonsRed           mahjong.Card = 35 // 红中
mahjong.DragonsGreen         mahjong.Card = 36 // 发财
mahjong.DragonsWhite         mahjong.Card = 37 // 白板
```

#### 6.花牌（8张）

```text
mahjong.FlowersSpring        mahjong.Card = 41 // 春
mahjong.FlowersSummer        mahjong.Card = 42 // 夏
mahjong.FlowersAutumn        mahjong.Card = 43 // 秋
mahjong.FlowersWinter        mahjong.Card = 44 // 冬
mahjong.FlowersPlum          mahjong.Card = 45 // 梅
mahjong.FlowersOrchid        mahjong.Card = 46 // 兰
mahjong.FlowersChrysanthemum mahjong.Card = 47 // 菊
mahjong.FlowersBamboo        mahjong.Card = 48 // 竹
```

### 二、特殊术语

#### 1.序牌（108张）

```text
mahjong.CharactersOne        mahjong.Card = 1  // 一万
mahjong.CharactersTwo        mahjong.Card = 2  // 二万
mahjong.CharactersThree      mahjong.Card = 3  // 三万
mahjong.CharactersFour       mahjong.Card = 4  // 四万
mahjong.CharactersFive       mahjong.Card = 5  // 五万
mahjong.CharactersSix        mahjong.Card = 6  // 六万
mahjong.CharactersSeven      mahjong.Card = 7  // 七万
mahjong.CharactersEight      mahjong.Card = 8  // 八万
mahjong.CharactersNine       mahjong.Card = 9  // 九万

mahjong.BamboosOne           mahjong.Card = 11 // 一条
mahjong.BamboosTwo           mahjong.Card = 12 // 二条
mahjong.BamboosThree         mahjong.Card = 13 // 三条
mahjong.BamboosFour          mahjong.Card = 14 // 四条
mahjong.BamboosFive          mahjong.Card = 15 // 五条
mahjong.BamboosSix           mahjong.Card = 16 // 六条
mahjong.BamboosSeven         mahjong.Card = 17 // 七条
mahjong.BamboosEight         mahjong.Card = 18 // 八条
mahjong.BamboosNine          mahjong.Card = 19 // 九条

mahjong.DotsOne              mahjong.Card = 21 // 一筒
mahjong.DotsTwo              mahjong.Card = 22 // 二筒
mahjong.DotsThree            mahjong.Card = 23 // 三筒
mahjong.DotsFour             mahjong.Card = 24 // 四筒
mahjong.DotsFive             mahjong.Card = 25 // 五筒
mahjong.DotsSix              mahjong.Card = 26 // 六筒
mahjong.DotsSeven            mahjong.Card = 27 // 七筒
mahjong.DotsEight            mahjong.Card = 28 // 八筒
mahjong.DotsNine             mahjong.Card = 29 // 九筒
```

#### 2.字牌（28张）

```text
mahjong.WindsEast            mahjong.Card = 31 // 东
mahjong.WindsSouth           mahjong.Card = 32 // 南
mahjong.WindsWest            mahjong.Card = 33 // 西
mahjong.WindsNorth           mahjong.Card = 34 // 北

mahjong.DragonsRed           mahjong.Card = 35 // 红中
mahjong.DragonsGreen         mahjong.Card = 36 // 发财
mahjong.DragonsWhite         mahjong.Card = 37 // 白板
```

#### 3.大幺牌（52张）

```text
mahjong.CharactersOne        mahjong.Card = 1  // 一万
mahjong.CharactersNine       mahjong.Card = 9  // 九万

mahjong.BamboosOne           mahjong.Card = 11 // 一条
mahjong.BamboosNine          mahjong.Card = 19 // 九条

mahjong.DotsOne              mahjong.Card = 21 // 一筒
mahjong.DotsNine             mahjong.Card = 29 // 九筒

mahjong.WindsEast            mahjong.Card = 31 // 东
mahjong.WindsSouth           mahjong.Card = 32 // 南
mahjong.WindsWest            mahjong.Card = 33 // 西
mahjong.WindsNorth           mahjong.Card = 34 // 北

mahjong.DragonsRed           mahjong.Card = 35 // 红中
mahjong.DragonsGreen         mahjong.Card = 36 // 发财
mahjong.DragonsWhite         mahjong.Card = 37 // 白板
```

#### 4.小幺牌（84张）

```text
mahjong.CharactersTwo        mahjong.Card = 2  // 二万
mahjong.CharactersThree      mahjong.Card = 3  // 三万
mahjong.CharactersFour       mahjong.Card = 4  // 四万
mahjong.CharactersFive       mahjong.Card = 5  // 五万
mahjong.CharactersSix        mahjong.Card = 6  // 六万
mahjong.CharactersSeven      mahjong.Card = 7  // 七万
mahjong.CharactersEight      mahjong.Card = 8  // 八万

mahjong.BamboosTwo           mahjong.Card = 12 // 二条
mahjong.BamboosThree         mahjong.Card = 13 // 三条
mahjong.BamboosFour          mahjong.Card = 14 // 四条
mahjong.BamboosFive          mahjong.Card = 15 // 五条
mahjong.BamboosSix           mahjong.Card = 16 // 六条
mahjong.BamboosSeven         mahjong.Card = 17 // 七条
mahjong.BamboosEight         mahjong.Card = 18 // 八条

mahjong.DotsTwo              mahjong.Card = 22 // 二筒
mahjong.DotsThree            mahjong.Card = 23 // 三筒
mahjong.DotsFour             mahjong.Card = 24 // 四筒
mahjong.DotsFive             mahjong.Card = 25 // 五筒
mahjong.DotsSix              mahjong.Card = 26 // 六筒
mahjong.DotsSeven            mahjong.Card = 27 // 七筒
mahjong.DotsEight            mahjong.Card = 28 // 八筒
```

#### 5.中张牌（36张）

```text
mahjong.CharactersFour       mahjong.Card = 4  // 四万
mahjong.CharactersFive       mahjong.Card = 5  // 五万
mahjong.CharactersSix        mahjong.Card = 6  // 六万

mahjong.BamboosFour          mahjong.Card = 14 // 四条
mahjong.BamboosFive          mahjong.Card = 15 // 五条
mahjong.BamboosSix           mahjong.Card = 16 // 六条

mahjong.DotsFour             mahjong.Card = 24 // 四筒
mahjong.DotsFive             mahjong.Card = 25 // 五筒
mahjong.DotsSix              mahjong.Card = 26 // 六筒
```

#### 6.圈风

```text
mahjong.CircleWindEast  mahjong.CircleWind = 1 // 东风圈
mahjong.CircleWindSouth mahjong.CircleWind = 2 // 南风圈
mahjong.CircleWindWest  mahjong.CircleWind = 3 // 西风圈
mahjong.CircleWindNorth mahjong.CircleWind = 4 // 北风圈
```

### 三、示例

详细实例：[example](./example)

### 四、API

#### 1.牌

- 文件：card.go
- 类库：Card

```text
// Int 获取牌的数字编号
func (c Card) Int() int

// Int8 获取牌的数字编号
func (c Card) Int8() int8

// Int16 获取牌的数字编号
func (c Card) Int16() int16

// Int32 获取牌的数字编号
func (c Card) Int32() int32

// Int64 获取牌的数字编号
func (c Card) Int64() int64

// String 获取牌的字符串代号
func (c Card) String() string

// Suit 获取牌的花色
func (c Card) Suit() Suit

// Prev 当前花色的上N张牌
func (c Card) Prev(n ...int) Card

// Next 当前花色的下N张牌
func (c Card) Next(n ...int) Card

// Offset 偏移N张牌
func (c Card) Offset(n int) Card

// IsValid 检测牌是否有效
func (c Card) IsValid() bool

// IsCharacters 是否是万字
func (c Card) IsCharacters() bool

// IsBamboos 是否是条子
func (c Card) IsBamboos() bool

// IsDots 是否是筒子
func (c Card) IsDots() bool

// IsWinds 是否是风牌
func (c Card) IsWinds() bool

// IsDragons 是否是箭牌
func (c Card) IsDragons() bool

// IsHonors 是否是字牌
func (c Card) IsHonors() bool

// IsRanks 是否是序数牌
func (c Card) IsRanks() bool

// IsFlowers 是否是花牌
func (c Card) IsFlowers() bool

// IsBigAce 是否是大幺牌（所有数牌中序号为1、9的牌，以及字牌，统称为幺九牌）
func (c Card) IsBigAce() bool

// IsSmallAce 是否是小幺牌
func (c Card) IsSmallAce() bool

// IsMiddleTile 是否是中张（所有数牌的4、5、6称为中张）
func (c Card) IsMiddleTile() bool
```

#### 2.花色

- 文件：suit.go
- 类库：Suit

```text
// Int 获取花色的数字编号
func (s Suit) Int() int

// Int8 获取花色的数字编号
func (s Suit) Int8() int8

// Int16 获取花色的数字编号
func (s Suit) Int16() int16

// Int32 获取花色的数字编号
func (s Suit) Int32() int32

// Int64 获取花色的数字编号
func (s Suit) Int64() int64

// String 获取花色的字符串编号
func (s Suit) String() string

// IsValid 是否有效花色
func (s Suit) IsValid() bool

// IsDots 是否是筒牌
func (s Suit) IsDots() bool

// IsBamboos 是否是条牌
func (s Suit) IsBamboos() bool

// IsCharacters 是否是万牌
func (s Suit) IsCharacters() bool

// IsWinds 是否是风牌
func (s Suit) IsWinds() bool

// IsDragons 是否是箭牌
func (s Suit) IsDragons() bool

// IsFlowers 是否是花牌
func (s Suit) IsFlowers() bool

// IsHonors 是否是字牌
func (s Suit) IsHonors() bool

// IsRanks 是否是序数牌
func (s Suit) IsRanks() bool
```

#### 3.牌池

- 文件：cards.go
- 类库：Cards

```text
// NewCards 新建牌池
func NewCards(excludes ...Suit) *Cards

// Shuffle 洗牌（仅打乱牌序，不恢复牌张数）
func (c *Cards) Shuffle()

// Reshuffle 重新洗牌（恢复初始牌张数，并打乱牌序）
func (c *Cards) Reshuffle()

// Count 牌张数
func (c *Cards) Count() int

// Cards 获取所有牌
func (c *Cards) Cards() []Card

// Deal 发牌
func (c *Cards) Deal(num int, isTail ...bool) ([]Card, bool)

// Remove 移除指定牌
func (c *Cards) Remove(cards ...Card)

// Tails 统计尾部发牌数
func (c *Cards) Tails() int

// Load 装载自定义牌库（配牌使用）
func (c *Cards) Load(cards []Card)
```

#### 4.骰子

- 文件：dices.go
- 类库：Dices

```text
// NewDices 新建骰子
func NewDices() *Dices

// Shake 摇动骰子
func (d *Dices) Shake() DiceResult

// Result 获取骰子结果
func (d *Dices) Result() DiceResult
```

#### 4.骰子结果

- 文件：dices.go
- 类库：DiceResult

```text
// TakeCardsSeatID 获取取牌位置编号（位置编号从1开始计算）（dealer = 1 2 3 4）
func (r DiceResult) TakeCardsSeatID(dealer int) int

// TakeCardsSeatIndex 获取取牌位置索引（位置索引从0开始计算）（dealer = 0 1 2 3）
func (r DiceResult) TakeCardsSeatIndex(dealer int) int

// TakeCardsPierID 获取取牌的牌墩编号（牌墩编号从1开始计算）
func (r DiceResult) TakeCardsPierID() int

// TakeCardsPierIndex 获取取牌的牌墩索引（牌墩索引从0开始计算）
func (r DiceResult) TakeCardsPierIndex() int

// Sum 获取骰子结果总和
func (r DiceResult) Sum() (sum int)

// Min 获取骰子结果最小值
func (r DiceResult) Min() int

// Max 获取骰子结果最大值
func (r DiceResult) Max() int

// Val 将骰子结果合并成一个值 x*10+y
func (r DiceResult) Val() int
```

#### 5.操作

- 文件：action.go
- 类库：Action

```text
const (
	ActionNone        Action = 0       // 无效
	ActionDraw        Action = 1 << 0  // 抓牌
	ActionPlay        Action = 1 << 1  // 出牌
	ActionPass        Action = 1 << 2  // 过牌
	ActionLack        Action = 1 << 3  // 定缺
	ActionChowLeft    Action = 1 << 4  // 左吃
	ActionChowMiddle  Action = 1 << 5  // 中吃
	ActionChowRight   Action = 1 << 6  // 右吃
	ActionPong        Action = 1 << 7  // 碰牌
	ActionKongShow    Action = 1 << 8  // 明杠
	ActionKongHide    Action = 1 << 9  // 暗杠
	ActionKongTurn    Action = 1 << 10 // 转弯杠
	ActionTing        Action = 1 << 11 // 听牌
	ActionWinShooting Action = 1 << 12 // 炮胡
	ActionWinDrawn    Action = 1 << 13 // 自摸
	ActionWinRobbing  Action = 1 << 14 // 抢杠胡
)

// IsPlay 检测是否是出牌操作
func (a Action) IsPlay() bool

// IsPass 检测是否是过牌操作
func (a Action) IsPass() bool

// IsChow 检测是否是吃牌操作
func (a Action) IsChow() bool

// IsPong 检测是否是碰牌操作
func (a Action) IsPong() bool

// IsKong 检测是否是杠牌操作
func (a Action) IsKong() bool

// IsKongHide 检测是否是暗杠操作
func (a Action) IsKongHide() bool

// IsKongShow 检测是否是明杠操作
func (a Action) IsKongShow() bool

// IsKongTurn 检测是否是转弯杠操作
func (a Action) IsKongTurn() bool

// IsWin 检测是否是胡牌操作
func (a Action) IsWin() bool

// IsWinDrawn 检测是否自摸胡
func (a Action) IsWinDrawn() bool

// IsWinShooting 检测是否放炮胡
func (a Action) IsWinShooting() bool

// IsWinRobbing 检测是否抢杠胡
func (a Action) IsWinRobbing() bool

// Priority 操作优先级（胡 > 杠 = 碰 > 吃 > 过 > 无）
func (a Action) Priority() int

// String 获取操作的字符串名
func (a Action) String()
```

#### 6.圈风

- 文件：circle.go
- 类库：Circle

```text
// String 获取圈风的字符串名
func (w CircleWind) String() string

// IsEast 是否东风圈
func (w CircleWind) IsEast() bool

// IsSouth 是否南风圈
func (w CircleWind) IsSouth() bool

// IsWest 是否西风圈
func (w CircleWind) IsWest() bool

// IsNorth 是否北风圈
func (w CircleWind) IsNorth() bool
```

#### 7.牌桌

- 文件：table.go
- 类库：Table

```text
// NewTable 创建桌子
func NewTable(opts ...Option) *Table

// Cards 获取牌池
func (t *Table) Cards() *Cards

// Dices 获取骰子
func (t *Table) Dices() *Dices

// Round 获取局数
func (t *Table) Round() int

// Circle 当前圈数
func (t *Table) Circle() int

// CircleWind 当前圈风
func (t *Table) CircleWind() CircleWind

// Assign 分配座位
func (t *Table) Assign(id int) *Seat

// Seats 获取座位
func (t *Table) Seats() []*Seat

// TotalSeats 位置总数
func (t *Table) TotalSeats() int

// TotalPlayers 统计玩家总数
func (t *Table) TotalPlayers() int

// Deal 开始发牌
func (t *Table) Deal(fn ...func(table *Table))

// Dealer 庄家位置
func (t *Table) Dealer() *Seat

// Next 下一局
// 1.自动累计局数
// 2.自动洗牌
// 3.自动摇骰
// 4.自动换圈
// 5.自动定庄
func (t *Table) Next() int

// CurrPlaySeat 当前出牌位置
func (t *Table) CurrPlaySeat() *Seat

// CurrPlayTile 当前所出的牌
func (t *Table) CurrPlayTile() Card

// NextPlaySeat 下个出牌位置
func (t *Table) NextPlaySeat() *Seat

// SeatsRange 座位迭代
func (t *Table) SeatsRange(from int, fn func(seat *Seat) bool)

// IsDraw 是否流局
func (t *Table) IsDraw() bool

// IsFirstRound 是否是第一局
func (t *Table) IsFirstRound() bool

// IsFirstCircle 是否是第一圈
func (t *Table) IsFirstCircle() bool

// WinResults 胡牌结果
func (t *Table) WinResults() []*WinResult

// FirstWinner 第一个胡牌位置
func (t *Table) FirstWinner() *Seat

// FirstShooter 第一个放炮位置
func (t *Table) FirstShooter() *Seat
```

#### 8.座位

- 文件：seat.go
- 类库：Seat

```text
// ID 座位ID
func (s *Seat) ID() int

// Table 牌桌
func (s *Seat) Table() *Table

// Wind 获取风位号
func (s *Seat) Wind() int

// DealerTimes 连庄次数
func (s *Seat) DealerTimes() int

// HasLackSuit 检测是否还有缺一门花色
func (s *Seat) HasLackSuit() bool

// IsLackSuit 是否时定缺花色
func (s *Seat) IsLackSuit(card Card) bool

// IsDealer 是否是庄家
func (s *Seat) IsDealer() bool

// IsWin 是否已经胡牌
func (s *Seat) IsWin() bool

// IsWinDrawn 是否自摸胡牌
func (s *Seat) IsWinDrawn() bool

// IsWinShooting 是否是炮胡
func (s *Seat) IsWinShooting() bool

// IsWinRobbing 是否抢杠胡
func (s *Seat) IsWinRobbing() bool

// IsTurnPlay 检测是否轮到出牌
func (s *Seat) IsTurnPlay() bool

// SetPlayTileChowTileDisable 设置打张吃张禁用状态
func (s *Seat) SetPlayTileChowTileDisable(disable bool)

// GetPlayTileChowTileDisable 获取打张吃张禁用状态
func (s *Seat) GetPlayTileChowTileDisable() bool

// SetChowTilePlayTileDisable 设置吃张打张禁用状态
func (s *Seat) SetChowTilePlayTileDisable(disable bool)

// GetChowTilePlayTileDisable 吃张打张禁用状态
func (s *Seat) GetChowTilePlayTileDisable() bool

// SetPlayTilePongTileDisable 设置打张碰张禁用状态
func (s *Seat) SetPlayTilePongTileDisable(disable bool)

// GetPlayTilePongTileDisable 获取打张碰张禁用状态
func (s *Seat) GetPlayTilePongTileDisable() bool

// SetPongTilePlayTileDisable 设置碰张打张禁用状态
func (s *Seat) SetPongTilePlayTileDisable(disable bool)

// GetPongTilePlayTileDisable 获取碰张打张禁用状态
func (s *Seat) GetPongTilePlayTileDisable() bool

// SetPassPongAgainPongDisable 设置过碰再碰禁用状态
func (s *Seat) SetPassPongAgainPongDisable(disable bool)

// GetPassPongAgainPongDisable 获取过碰再碰禁用状态
func (s *Seat) GetPassPongAgainPongDisable() bool

// SetPassWinAgainWinDisable 设置过胡再胡禁用状态
func (s *Seat) SetPassWinAgainWinDisable(disable bool)

// GetPassWinAgainWinDisable 获取过胡再胡禁用状态
func (s *Seat) GetPassWinAgainWinDisable() bool

// WinKind 胡牌类型
func (s *Seat) WinKind() Action

// WinTile 胡牌牌张
func (s *Seat) WinTile() Card

// WinProvider 胡牌提供者位置
func (s *Seat) WinProvider() *Seat

// HideCards 暗牌
func (s *Seat) HideCards() *HideCards

// ShowCards 显牌
func (s *Seat) ShowCards() *ShowCards

// PlayCards 打出的牌
func (s *Seat) PlayCards() *PlayCards

// DisablePlayTiles 禁止出的牌
func (s *Seat) DisablePlayTiles() Tiles

// TotalRemainingCards 剩余牌数量（排除自己的手牌、所有玩家打出的牌、所有玩家摆牌）
func (s *Seat) TotalRemainingCards(card Card) int

// AnalyzeTingList 分析听牌列表
func (s *Seat) AnalyzeTingList() []*TingResult

// AnalyzePlayTingList 分析打牌听牌列表（打x牌听m,n,y牌）
func (s *Seat) AnalyzePlayTingList() []*PlayTingResult

// CanChow 检测是否可以吃牌
func (s *Seat) CanChow(card Card, action ...Action) (bool, []*ChowResult)

// CanChowLeft 检测是否可以吃左边牌
func (s *Seat) CanChowLeft(card Card) (bool, []Card)

// CanChowMiddle 检测是否可以吃中间牌
func (s *Seat) CanChowMiddle(card Card) (bool, []Card)

// CanChowRight 检测是否可以吃右边牌
func (s *Seat) CanChowRight(card Card) (bool, []Card)

// CanPong 检测该牌是否可以碰
func (s *Seat) CanPong(card Card) bool

// CanKong 检测是否可以杠
func (s *Seat) CanKong(card Card, action ...Action) (bool, Action)

// CanKongHide 检测是否允许暗杠（被杠的牌一定要是自己的牌）
func (s *Seat) CanKongHide(card Card) bool

// CanKongShow 检测是否允许明杠
func (s *Seat) CanKongShow(card Card) bool

// CanKongTurn 检测是否允许转弯杠
func (s *Seat) CanKongTurn(card Card) bool

// CanNaturalKongHide 是否可以直接暗杠（用于手牌中已存在暗杠的情况）
func (s *Seat) CanNaturalKongHide() (bool, []Card)

// CanNaturalKongTurn 是否可以直接转弯杠（用于手牌中已存在可以转弯杠的情况）
func (s *Seat) CanNaturalKongTurn() (bool, []Card)

// CanWinDrawn 检测是否可以自摸胡
func (s *Seat) CanWinDrawn(card Card) bool

// CanWinShooting 检测是否可以炮胡
func (s *Seat) CanWinShooting(card Card) bool

// CanWinRobbing 检测是否可以抢杠胡
func (s *Seat) CanWinRobbing(card Card) bool

// CanNaturalWinDrawn 检测是否可以直接自摸胡（用于手牌张数已经满足胡牌条件）
func (s *Seat) CanNaturalWinDrawn() bool

// CanNaturalWinSevenPairs 检测是否可以直接胡七对
func (s *Seat) CanNaturalWinSevenPairs(filters ...Card) bool

// IsAllowedAction 检测是否允许操作
func (s *Seat) IsAllowedAction(action Action) bool

// Operation 获取可操作项
func (s *Seat) Operation() *Operation

// Action 操作牌
func (s *Seat) Action(action Action, card Card)

// Pass 过牌
func (s *Seat) Pass()

// Play 出牌
func (s *Seat) Play(card Card)

// Chow 吃牌
func (s *Seat) Chow(action Action)

// Pong 碰牌
func (s *Seat) Pong()

// Kong 杠牌
func (s *Seat) Kong(action Action, card Card)

// Win 胡牌
func (s *Seat) Win(action Action)

// NextSeat 下n个座位（座位可能为空）
func (s *Seat) NextSeat(n ...int) *Seat

// PrevSeat 上n个座位（座位可能为空）
func (s *Seat) PrevSeat(n ...int) *Seat

// NextNonEmptySeat 下n个非空座位
func (s *Seat) NextNonEmptySeat(n ...int) *Seat

// PrevNonEmptySeat 上n个非空座位
func (s *Seat) PrevNonEmptySeat(n ...int) *Seat

// NextNonWinSeat 下n个非胡座位
func (s *Seat) NextNonWinSeat(n ...int) *Seat

// PrevNonWinSeat 上n个非胡座位
func (s *Seat) PrevNonWinSeat(n ...int) *Seat
```