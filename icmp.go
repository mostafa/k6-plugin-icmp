package main

import (
	"context"
	"time"

	"github.com/loadimpact/k6/stats"
	goping "github.com/sparrc/go-ping"
)

type icmp struct{}

func New() *icmp {
	return &icmp{}
}

func (*icmp) Ping(ctx context.Context, hostname string, count int) error {
	state, err := GetState(ctx)

	if err == nil {
		pinger, err := goping.NewPinger(hostname)
		if err != nil {
			return err
		}

		if count == 0 {
			count = 1
		}

		pinger.SetPrivileged(true)
		pinger.Count = count
		now := time.Now()
		pinger.Run()

		currentStats := pinger.Statistics()
		tags := make(map[string]string)
		tags["address"] = currentStats.Addr

		stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
			Time:   now,
			Metric: PacketsSent,
			Tags:   stats.IntoSampleTags(&tags),
			Value:  float64(currentStats.PacketsSent),
		})

		stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
			Time:   now,
			Metric: PacketsRecv,
			Tags:   stats.IntoSampleTags(&tags),
			Value:  float64(currentStats.PacketsRecv),
		})

		stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
			Time:   now,
			Metric: PacketLoss,
			Tags:   stats.IntoSampleTags(&tags),
			Value:  float64(currentStats.PacketLoss),
		})

		stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
			Metric: MinRtt,
			Tags:   stats.IntoSampleTags(&tags),
			Value:  float64(currentStats.MinRtt.Milliseconds()),
		})

		stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
			Metric: MaxRtt,
			Tags:   stats.IntoSampleTags(&tags),
			Value:  float64(currentStats.MaxRtt.Milliseconds()),
		})

		stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
			Metric: AvgRtt,
			Tags:   stats.IntoSampleTags(&tags),
			Value:  float64(currentStats.AvgRtt.Milliseconds()),
		})

		stats.PushIfNotDone(ctx, state.Samples, stats.Sample{
			Metric: StdDevRtt,
			Tags:   stats.IntoSampleTags(&tags),
			Value:  float64(currentStats.StdDevRtt.Milliseconds()),
		})

		return nil
	}

	return err
}
