package kafka_consumer

import (
	"fmt"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/fyannk/telegraf/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/fyannk/telegraf/plugins/parsers"
)

func TestReadsMetricsFromKafka(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	brokerPeers := []string{testutil.GetLocalHost() + ":9092"}
	zkPeers := []string{testutil.GetLocalHost() + ":2181"}
	testTopic := fmt.Sprintf("telegraf_test_topic_%d", time.Now().Unix())

	// Send a Kafka message to the kafka host
	msg := "cpu_load_short,direction=in,host=server01,region=us-west value=23422.0 1422568543702900257"
	producer, err := sarama.NewSyncProducer(brokerPeers, nil)
	require.NoError(t, err)
	_, _, err = producer.SendMessage(
		&sarama.ProducerMessage{
			Topic: testTopic,
			Value: sarama.StringEncoder(msg),
		})
	require.NoError(t, err)
	defer producer.Close()

	// Start the Kafka Consumer
	k := &Kafka{
		ConsumerGroup:  "telegraf_test_consumers",
		Topics:         []string{testTopic},
		ZookeeperPeers: zkPeers,
		PointBuffer:    100000,
		Offset:         "oldest",
	}
	p, _ := parsers.NewInfluxParser()
	k.SetParser(p)

	// Verify that we can now gather the sent message
	var acc testutil.Accumulator

	// Sanity check
	assert.Equal(t, 0, len(acc.Metrics), "There should not be any points")
	if err := k.Start(&acc); err != nil {
		t.Fatal(err.Error())
	} else {
		defer k.Stop()
	}

	waitForPoint(&acc, t)

	// Gather points
	err = k.Gather(&acc)
	require.NoError(t, err)
	if len(acc.Metrics) == 1 {
		point := acc.Metrics[0]
		assert.Equal(t, "cpu_load_short", point.Measurement)
		assert.Equal(t, map[string]interface{}{"value": 23422.0}, point.Fields)
		assert.Equal(t, map[string]string{
			"host":      "server01",
			"direction": "in",
			"region":    "us-west",
		}, point.Tags)
		assert.Equal(t, time.Unix(0, 1422568543702900257).Unix(), point.Time.Unix())
	} else {
		t.Errorf("No points found in accumulator, expected 1")
	}
}

// Waits for the metric that was sent to the kafka broker to arrive at the kafka
// consumer
func waitForPoint(acc *testutil.Accumulator, t *testing.T) {
	// Give the kafka container up to 2 seconds to get the point to the consumer
	ticker := time.NewTicker(5 * time.Millisecond)
	counter := 0
	for {
		select {
		case <-ticker.C:
			counter++
			if counter > 1000 {
				t.Fatal("Waited for 5s, point never arrived to consumer")
			} else if acc.NFields() == 1 {
				return
			}
		}
	}
}
