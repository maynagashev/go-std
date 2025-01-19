package main

import "testing"

func BenchmarkLoadBalancer(b *testing.B) {
	const connsN = 100
	const triesN = 1000

	conns := make([]*Connection, 0, connsN)
	for i := 0; i < connsN; i++ {
		conns = append(conns, &Connection{})
	}

	lbChan := NewLoadBalancerChan(conns)
	lbChan.Init()
	defer lbChan.Close()
	lbAtomic := NewLoadBalancerAtomic(conns)
	lbMutex := NewLoadBalancerMutex(conns)

	b.ResetTimer()

	b.Run("chan", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbChan.NextConn()
		}
	})

	b.Run("atomic", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbAtomic.NextConn()
		}
	})

	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbMutex.NextConn()
		}
	})
}
