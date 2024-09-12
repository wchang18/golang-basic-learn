package lesson2

import "testing"

// 手机充电器接口
type PhoneCharger interface {
	Output5V()
}

// 小米手机充电器
type MiPhoneChanger struct {
}

func (m *MiPhoneChanger) Output5V() {
	println("小米充电器输出5V电压")
}

// 苹果手机充电器
type IPhoneChanger struct {
}

func (m *IPhoneChanger) Output5V() {
	println("苹果充电器输出5V电压")
}

// 苹果笔记本充电器
type MacbookChanger struct {
}

func (m *MacbookChanger) Output28V() {
	println("Macbook充电器输出28V电压")
}

func (m *MacbookChanger) Output5V() {
	m.Output28V()
	println("macbook适配器转换成5V")
}

// 苹果笔记本适配手机
type MacbookChangerAdapter struct {
	core *MacbookChanger
}

func (m *MacbookChangerAdapter) Output5V() {
	m.core.Output28V()
	println("macbook适配器转换成5V")
}

// 手机充电器适配接口
type PhoneChargerAdapter interface {
	Charger(phoneChanger PhoneCharger)
}

// 小米手机
type MiPhone struct{}

func (m *MiPhone) Charger(phoneChanger PhoneCharger) {
	phoneChanger.Output5V()
}

func TestMiPhone(t *testing.T) {
	miPhone := &MiPhone{}
	macChanger := &MacbookChanger{}
	miPhone.Charger(macChanger)

	//iphoneChanger := &IPhoneChanger{}
	//miPhone.Charger(iphoneChanger)
}
