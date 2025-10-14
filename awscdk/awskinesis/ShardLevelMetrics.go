package awskinesis


// Enhanced shard-level metrics.
//
// Example:
//   stream := kinesis.NewStream(this, jsii.String("MyStream"), &StreamProps{
//   	ShardLevelMetrics: []shardLevelMetrics{
//   		kinesis.*shardLevelMetrics_ALL,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html#kinesis-metrics-shard
//
type ShardLevelMetrics string

const (
	// The number of bytes successfully put to the shard over the specified time period.
	ShardLevelMetrics_INCOMING_BYTES ShardLevelMetrics = "INCOMING_BYTES"
	// The number of records successfully put to the shard over the specified time period.
	ShardLevelMetrics_INCOMING_RECORDS ShardLevelMetrics = "INCOMING_RECORDS"
	// The age of the last record in all GetRecords calls made against a shard, measured over the specified time period.
	ShardLevelMetrics_ITERATOR_AGE_MILLISECONDS ShardLevelMetrics = "ITERATOR_AGE_MILLISECONDS"
	// The number of bytes retrieved from the shard, measured over the specified time period.
	ShardLevelMetrics_OUTGOING_BYTES ShardLevelMetrics = "OUTGOING_BYTES"
	// The number of records retrieved from the shard, measured over the specified time period.
	ShardLevelMetrics_OUTGOING_RECORDS ShardLevelMetrics = "OUTGOING_RECORDS"
	// The number of GetRecords calls throttled for the shard over the specified time period.
	ShardLevelMetrics_READ_PROVISIONED_THROUGHPUT_EXCEEDED ShardLevelMetrics = "READ_PROVISIONED_THROUGHPUT_EXCEEDED"
	// The number of records rejected due to throttling for the shard over the specified time period.
	ShardLevelMetrics_WRITE_PROVISIONED_THROUGHPUT_EXCEEDED ShardLevelMetrics = "WRITE_PROVISIONED_THROUGHPUT_EXCEEDED"
	// All metrics.
	ShardLevelMetrics_ALL ShardLevelMetrics = "ALL"
)

