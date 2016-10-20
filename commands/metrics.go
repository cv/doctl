package commands

func RunMetricsChart(c *CmdConfig) error {
	return c.Display(&chart{Link: "http://godo.internal.digitalocean.com/xxx.png"})
}

func RunMetricsGet(c *CmdConfig) error {
	return c.Display(&metric{})
}

func RunMetricsList(c *CmdConfig) error {
	return c.Display(&metric{})
}
