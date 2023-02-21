package awsefs


// EFS Throughput mode.
// See: https://docs.aws.amazon.com/efs/latest/ug/performance.html#throughput-modes
//
type ThroughputMode string

const (
	// This mode scales as the size of the file system in the standard storage class grows.
	ThroughputMode_BURSTING ThroughputMode = "BURSTING"
	// This mode can instantly provision the throughput of the file system (in MiB/s) independent of the amount of data stored.
	ThroughputMode_PROVISIONED ThroughputMode = "PROVISIONED"
	// This mode scales the throughput automatically regardless of file system size.
	ThroughputMode_ELASTIC ThroughputMode = "ELASTIC"
)

