package main

import "log"

type (
	// 订单创建者，相应能响应的事件
	IOrderCreator interface {
		OnOrderDone(*Order)
	}

	// 订单
	Order struct {
		id       string // SendOrder返回的id
		strategy IOrderCreator
	}

	// 数据库中存着的品种详情，可通过symbol索引到
	OrderManager struct {
		pending map[string]*Order // 未成交或部分成交订单，可cancle
	}
)

func NewOrderManager() *OrderManager {
	return &OrderManager{pending: make(map[string]*Order)}
}

func (om *OrderManager) AppendOrder(i string, s IOrderCreator) {
	(om.pending)[i] = &Order{id: i, strategy: s}
}

func (om *OrderManager) OrderDone() {
	for k := range om.pending {
		o := om.pending[k]
		o.strategy.OnOrderDone(o)
	}
}

type Strategy struct {
	name        string
	realCreator IOrderCreator
}

// IOrderCreator interface: 订单创建者能响应的事件
func (s *Strategy) AddOrder(om *OrderManager) {
	om.AppendOrder(s.name, s.realCreator)
}

// IOrderCreator interface: 订单创建者能响应的事件
func (s *Strategy) OnOrderDone(o *Order) {
	log.Print(o.id)
}

type MyStrategy struct {
	Strategy
	myname string
}

func (s *MyStrategy) OnOrderDone(o *Order) {
	log.Print("my")
}

func main() {
	mys := MyStrategy{Strategy: Strategy{name: "parent"}, myname: "my"}
	mys.Strategy.realCreator = &mys
	om := NewOrderManager()

	mys.AddOrder(om)
	om.OrderDone()
}
