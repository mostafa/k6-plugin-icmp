package main

import "github.com/loadimpact/k6/stats"

var (
	PacketsRecv = stats.New("icmp.packets_sent", stats.Counter)
	PacketsSent = stats.New("icmp.packets_received", stats.Counter)
	PacketLoss  = stats.New("icmp.packets_loss", stats.Counter)

	MinRtt    = stats.New("icmp.min_rtt", stats.Trend, stats.Time)
	MaxRtt    = stats.New("icmp.max_rtt", stats.Trend, stats.Time)
	AvgRtt    = stats.New("icmp.avg_rtt", stats.Trend, stats.Time)
	StdDevRtt = stats.New("icmp.std_dev_rtt", stats.Trend, stats.Time)
)
