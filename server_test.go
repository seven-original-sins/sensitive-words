package main

import (
	"fmt"
	"testing"
)

func BenchmarkGenTree(b *testing.B)  {
	s := new(Server)
	err := s.Build()

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
	err := s.Build()

	if err != nil {
		fmt.Println(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.m.MultiPatternSearch([]rune("所谓TMD，关键TMD需要如何写。 达·芬奇在不经意间这样说过，大胆和坚定的决心能够抵得上武器的精良。我希望诸位也能好好地体会这句话。 经过上述讨论， 他妈的，到底应该如何实现。 每个人都不得不面对这些问题。 在面对这种问题时， 俾斯麦在不经意间这样说过，失败是坚忍的最后考验。这不禁令我深思。 德谟克利特在不经意间这样说过，节制使快乐增加并使享受加强。这不禁令我深思。 问题的关键究竟为何？ 海贝尔曾经提到过，人生就是学校。在那里，与其说好的教师是幸福，不如说好的教师是不幸。这不禁令我深思。 一般来讲，我们都必须务必慎重的考虑考虑。 了解清楚他妈的到底是一种怎么样的存在，是解决一切问题的关键。"), false)

		//for _, item := range res {
		//	fmt.Println(item.Word)
		//}
	}
}