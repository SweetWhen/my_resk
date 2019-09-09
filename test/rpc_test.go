package test

import (
	"git.imooc.com/wendell1000/infra/gorpc"
	"git.imooc.com/wendell1000/infra/lb"
	"git.imooc.com/wendell1000/resk/services"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/go-eureka-client/eureka"
	"github.com/tietang/props/ini"
	"testing"
)

func TestGoRpcClient_Call(t *testing.T) {
	//创建一个eureka client
	conf := ini.NewIniFileConfigSource("ec_test.ini")
	client := eureka.NewClient(conf)
	client.Start()
	client.Applications, _ = client.GetApplications()

	//创建一个apps实例
	apps := &lb.Apps{Client: client}
	g := gorpc.GoRpcClient{Apps: apps}
	Convey("goRpc测试", t, func() {
		in := services.RedEnvelopeSendingDTO{
			Amount:       decimal.NewFromFloat(1),
			UserId:       "1MD35g7HA9aukHZN5VEg2kTNYYx",
			Username:     "测试用户",
			EnvelopeType: services.LuckyEnvelopeType,
			Quantity:     2,
			Blessing:     "",
		}
		out := &services.RedEnvelopeActivity{}
		err := g.Call("resk", "EnvelopeRpc.SendOut", in, &out)
		So(err, ShouldBeNil)
		So(out.EnvelopeNo, ShouldNotBeNil)
		So(out.Amount.String(), ShouldEqual, in.Amount.String())
		if err != nil {
			logrus.Error(err)
		}
		logrus.Infof("%+v", out)
	})

}
