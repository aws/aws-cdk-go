// The CDK Construct Library for AWS::MSK
package awscdkmskalpha


// The level of monitoring for the MSK cluster.
// See: https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html#metrics-details
//
// Experimental.
type ClusterMonitoringLevel string

const (
	// Default metrics are the essential metrics to monitor.
	// Experimental.
	ClusterMonitoringLevel_DEFAULT ClusterMonitoringLevel = "DEFAULT"
	// Per Broker metrics give you metrics at the broker level.
	// Experimental.
	ClusterMonitoringLevel_PER_BROKER ClusterMonitoringLevel = "PER_BROKER"
	// Per Topic Per Broker metrics help you understand volume at the topic level.
	// Experimental.
	ClusterMonitoringLevel_PER_TOPIC_PER_BROKER ClusterMonitoringLevel = "PER_TOPIC_PER_BROKER"
	// Per Topic Per Partition metrics help you understand consumer group lag at the topic partition level.
	// Experimental.
	ClusterMonitoringLevel_PER_TOPIC_PER_PARTITION ClusterMonitoringLevel = "PER_TOPIC_PER_PARTITION"
)

