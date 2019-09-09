package algo

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSimpleRand(t *testing.T) {
	count, amount := int64(10), int64(10000)
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := SimpleRand(count-i, remain)
		remain -= x
		sum += x
		fmt.Printf(" %d ", x)
	}
	fmt.Printf("\nsum:%d\n", sum)
	Convey("简单随机算法", t, func() {
		So(sum, ShouldEqual, amount)
	})
}
