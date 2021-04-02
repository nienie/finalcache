package stats

//Distribution 访问分布
type Distribution struct {
	Count    int64   //有多少个key访问被访问了
	Max      int64   //访问最多的次数
	Min      int64   //访问最少的次数
	Mean     float64 //平均每个key访问多少次
	StdDev   float64 //方差
	Variance float64 //协方差
	Pn       float64 //动态的分位数
	P999     float64 //99.9分位数，99.9%的比例是这个访问次数
	P995     float64 //99.5分位数，99.5%的比例是这个访问次数
	P99      float64 //99分位数，99%的比例是这个访问次数
	P95      float64 //90分位数，90%的比例是这个访问次数
	P75      float64 //75分位数，75%的比例是这个访问次数
	P50      float64 //50分位数，50%的比例是这个访问次数
	P25      float64 //25分位数，25%的比例是这个访问次数
	Sum      int64   //总访问次数
}

//NewDistribution ...
func NewDistribution(counter *Counter, percentile float64) Distribution {
	percentiles := []float64{percentile, 0.999, 0.995, 0.99, 0.95, 0.75, 0.50, 0.25}
	histogram := counter.Histogram()
	percs := histogram.Percentiles(percentiles)
	return Distribution{
		Count:    histogram.Count(),
		Max:      histogram.Max(),
		Min:      histogram.Min(),
		Mean:     histogram.Mean(),
		StdDev:   histogram.StdDev(),
		Variance: histogram.Variance(),
		Pn:       percs[0],
		P999:     percs[1],
		P995:     percs[2],
		P99:      percs[3],
		P95:      percs[4],
		P75:      percs[5],
		P50:      percs[6],
		P25:      percs[7],
		Sum:      histogram.Sum(),
	}
}
