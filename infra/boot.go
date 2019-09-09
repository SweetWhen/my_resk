package infra

import "github.com/tietang/props/kvs"

type BootApplication struct {
	conf kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{
		conf:conf,
		starterContext:StarterContext{}}

	b.starterContext[KeyProps] = conf
	return b
}


func (b *BootApplication)Start()  {
	//1. 初始化starter
	b.init()
	//2.安装starter
	b.setup()
	//3. 安装starter
	b.start()
}

func (b* BootApplication)init()  {

	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}

}

func (b* BootApplication)setup()  {

	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}

}

func (b* BootApplication)start()  {
	for i, starter := range StarterRegister.AllStarters() {
		if starter.StartBlocking() {
			if i+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)

			} else {
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)

		}
	}
}