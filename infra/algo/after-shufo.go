package algo

import "math/rand"

//后洗牌算法

func AfterShuffle(count, amount int64) []int64  {
	//计算最大可调度金额
	//洗牌
	inds := make([]int64, 0)
	max := amount - min* count
	remain := max
	for i := int64(0); i < count; i++{
		x := SimpleRand(count-i, remain)
		remain -= x
		inds = append(inds, x)
	}
	//洗牌
	rand.Shuffle(len(inds), func(i, j int) {
		inds[i], inds[j] = inds[j], inds[i]
	})

	return inds
}
