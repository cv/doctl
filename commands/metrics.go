package commands

func RunMetricsChart(c *CmdConfig) error {
	chart := &chart{
		Link: "http://godo.internal.digitalocean.com/xxx.png",
	}

	return c.Display(chart)
}

func RunMetricsGet(c *CmdConfig) error {
	metric := &metric{
		Name:      "cpu_total",
		HumanName: "CPU Total",
		HumanDesc: "Average CPU utilization",
		Unit:      "%",
		Value:     0.843,
	}

	return c.Display(metric)
}

func RunMetricsList(c *CmdConfig) error {
	metrics := &metricList{
		&metric{
			Name:      "cpu_total",
			HumanName: "CPU Total",
			HumanDesc: "Average CPU utilization",
			Unit:      "%",
		},
		&metric{
			Name:      "disk_read",
			HumanName: "Disk Read",
			HumanDesc: "Disk - Read",
			Unit:      "MB/s",
		},
		&metric{
			Name:      "disk_utilization",
			HumanName: "Disk utilization",
			HumanDesc: "Average disk utilization on /",
			Unit:      "%",
		},
	}

	return c.Display(metrics)
}
