package all

import (
	_ "github.com/fyannk/telegraf/plugins/outputs/amon"
	_ "github.com/fyannk/telegraf/plugins/outputs/amqp"
	_ "github.com/fyannk/telegraf/plugins/outputs/cloudwatch"
	_ "github.com/fyannk/telegraf/plugins/outputs/datadog"
	_ "github.com/fyannk/telegraf/plugins/outputs/file"
	_ "github.com/fyannk/telegraf/plugins/outputs/graphite"
	_ "github.com/fyannk/telegraf/plugins/outputs/graylog"
	_ "github.com/fyannk/telegraf/plugins/outputs/influxdb"
	_ "github.com/fyannk/telegraf/plugins/outputs/instrumental"
	_ "github.com/fyannk/telegraf/plugins/outputs/kafka"
	_ "github.com/fyannk/telegraf/plugins/outputs/kinesis"
	_ "github.com/fyannk/telegraf/plugins/outputs/librato"
	_ "github.com/fyannk/telegraf/plugins/outputs/mqtt"
	_ "github.com/fyannk/telegraf/plugins/outputs/nsq"
	_ "github.com/fyannk/telegraf/plugins/outputs/opentsdb"
	_ "github.com/fyannk/telegraf/plugins/outputs/prometheus_client"
	_ "github.com/fyannk/telegraf/plugins/outputs/riemann"
)
