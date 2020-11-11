package infra

import "github.com/spf13/viper"

// Starter接口，定义Starter基本方法
type Starter interface {
	// 初始化Starter配置
	Init(conf *viper.Viper)
	// 配置Starter参数
	Setup(conf *viper.Viper)
	// 启动Starter
	Start(conf *viper.Viper)
}

// BaseStarter，Starter接口的默认空实现
var _ Starter = new(BaseStarter)

type BaseStarter struct{}

func (b *BaseStarter) Init(conf *viper.Viper)  {}
func (b *BaseStarter) Setup(conf *viper.Viper) {}
func (b *BaseStarter) Start(conf *viper.Viper) {}

// Starter的注册器,保存所有Starter实例,全局单例
type starterRegister struct {
	starters []Starter
}

var register = new(starterRegister)

// 获取starterRegister
func StarterRegister() *starterRegister {
	return register
}

// 注册Starter
func (s *starterRegister) register(starter Starter) {
	s.starters = append(s.starters, starter)
}

// 获取所有被注册的Starters
func (s *starterRegister) allStarters() []Starter {
	return s.starters
}

func Register(starter Starter) {
	register.register(starter)
}

func AllStarters() []Starter {
	return register.allStarters()
}
