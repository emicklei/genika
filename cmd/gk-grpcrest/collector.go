package main

import "github.com/emicklei/proto"

type Collector struct {
	proto *proto.Proto
}

func Collect(p *proto.Proto) Collector {
	return Collector{p}
}

func (c Collector) Services() (list []*proto.Service) {
	for _, each := range c.proto.Elements {
		if c, ok := each.(*proto.Service); ok {
			list = append(list, c)
		}
	}
	return
}

func (c Collector) RPCsOf(s *proto.Service) (list []*proto.RPC) {
	for _, each := range s.Elements {
		if other, ok := each.(*proto.RPC); ok {
			list = append(list, other)
		}
	}
	return
}
