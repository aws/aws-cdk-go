package awskinesis


// Specifies the capacity mode to apply to this stream.
// Experimental.
type StreamMode string

const (
	// Specify the provisioned capacity mode.
	//
	// The stream will have `shardCount` shards unless
	// modified and will be billed according to the provisioned capacity.
	// Experimental.
	StreamMode_PROVISIONED StreamMode = "PROVISIONED"
	// Specify the on-demand capacity mode.
	//
	// The stream will autoscale and be billed according to the
	// volume of data ingested and retrieved.
	// Experimental.
	StreamMode_ON_DEMAND StreamMode = "ON_DEMAND"
)

