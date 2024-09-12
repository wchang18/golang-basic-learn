package lesson2

import (
	"fmt"
	"testing"
)

// Fruit 接口定义
type Fruit interface {
	Eat()
}

// Apple 结构体
type Apple struct{}

// Eat 方法实现了吃苹果的行为
func (a Apple) Eat() {
	fmt.Println("Eating an apple")
}

// Mango 结构体
type Mango struct{}

// Eat 方法实现了吃芒果的行为
func (m Mango) Eat() {
	fmt.Println("Eating a mango")
}

// Orange 结构体
type Orange struct{}

// Eat 方法实现了吃橙子的行为
func (o Orange) Eat() {
	fmt.Println("Eating an orange")
}

// FruitFactory 类
type FruitFactory struct{}

// CreateFruit 工厂方法，根据输入的字符串创建对应的水果实例
func (f *FruitFactory) CreateFruit(fruitType string) Fruit {
	switch fruitType {
	case "apple":
		return Apple{}
	case "mango":
		return Mango{}
	case "orange":
		return Orange{}
	default:
		fmt.Println("Unknown fruit type:", fruitType)
		return nil
	}
}

func TestEat(t *testing.T) {
	factory := FruitFactory{}

	// 创建并食用苹果
	apple := factory.CreateFruit("apple")
	if apple != nil {
		apple.Eat()
	}

	// 创建并食用芒果
	mango := factory.CreateFruit("mango")
	if mango != nil {
		mango.Eat()
	}

	// 创建并食用橙子
	orange := factory.CreateFruit("orange")
	if orange != nil {
		orange.Eat()
	}
}
