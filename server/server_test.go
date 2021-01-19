package server

import (
	"fmt"
	"testing"
)

func BenchmarkGenTree(b *testing.B)  {
	s := new(Server)
	err := s.Build(nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := s.LoadWords()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func BenchmarkMatch(b *testing.B) {
	s := new(Server)
	err := s.Build(nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.m.MultiPatternSearch([]rune("经过上述讨论， 既然如此， 问题的关键究竟为何？ 一般来说， 就我个人来说，卡宾枪对我的意义，不能不说非常重大。 从这个角度来看， 那么， 就我个人来说，卡宾枪对我的意义，不能不说非常重大。 对我个人而言，卡宾枪不仅仅是一个重大的事件，还可能会改变我的人生。 池田大作曾经说过，不要回避苦恼和困难，挺起身来向它挑战，进而克服它。这启发了我， 对我个人而言，卡宾枪不仅仅是一个重大的事件，还可能会改变我的人生。 要想清楚，卡宾枪，到底是一种怎么样的存在。 \n　　希腊说过一句富有哲理的话，最困难的事情就是认识自己。这不禁令我深思。 就我个人来说，卡宾枪对我的意义，不能不说非常重大。 卡宾枪因何而发生？ 裴斯泰洛齐曾经说过，今天应做的事没有做，明天再早也是耽误了。这句话语虽然很短，但令我浮想联翩。 美华纳曾经提到过，勿问成功的秘诀为何，且尽全力做你应该做的事吧。这句话语虽然很短，但令我浮想联翩。 这种事实对本人来说意义重大，相信对这个世界也是有一定意义的。 那么， 可是，即使是这样，卡宾枪的出现仍然代表了一定的意义。 了解清楚卡宾枪到底是一种怎么样的存在，是解决一切问题的关键。 从这个角度来看， 我们不得不面对一个非常尴尬的事实，那就是， 从这个角度来看， 莎士比亚在不经意间这样说过，本来无望的事，大胆尝试，往往能成功。带着这句话，我们还要更加慎重的审视这个问题： 那么， 卡宾枪因何而发生？ 鲁巴金曾经说过，读书是在别人思想的帮助下，建立起自己的思想。这启发了我， 经过上述讨论， 这样看来， 而这些并不是完全重要，更加重要的问题是， 卡宾枪的发生，到底需要如何做到，不卡宾枪的发生，又会如何产生。 带着这些问题，我们来审视一下卡宾枪。 既然如何， 就我个人来说，卡宾枪对我的意义，不能不说非常重大。 生活中，若卡宾枪出现了，我们就不得不考虑它出现了的事实。 总结的来说， 我们不得不面对一个非常尴尬的事实，那就是， 米歇潘在不经意间这样说过，生命是一条艰险的峡谷，只有勇敢的人才能通过。这不禁令我深思。 我们一般认为，抓住了问题的关键，其他一切则会迎刃而解。 可是，即使是这样，卡宾枪的出现仍然代表了一定的意义。 在这种困难的抉择下，本人思来想去，寝食难安。 我们不得不面对一个非常尴尬的事实，那就是， 德谟克利特在不经意间这样说过，节制使快乐增加并使享受加强。我希望诸位也能好好地体会这句话。"), false)

		//for _, item := range res {
		//	fmt.Println(item.Word)
		//}
	}
}