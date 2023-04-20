package awsec2


// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//
//   vpc.addFlowLog(jsii.String("FlowLogS3"), &FlowLogOptions{
//   	Destination: ec2.FlowLogDestination_ToS3(),
//   })
//
//   // Only reject traffic and interval every minute.
//   vpc.addFlowLog(jsii.String("FlowLogCloudWatch"), &FlowLogOptions{
//   	TrafficType: ec2.FlowLogTrafficType_REJECT,
//   	MaxAggregationInterval: ec2.FlowLogMaxAggregationInterval_ONE_MINUTE,
//   })
//
type FlowLogMaxAggregationInterval string

const (
	// 1 minute (60 seconds).
	FlowLogMaxAggregationInterval_ONE_MINUTE FlowLogMaxAggregationInterval = "ONE_MINUTE"
	// 10 minutes (600 seconds).
	FlowLogMaxAggregationInterval_TEN_MINUTES FlowLogMaxAggregationInterval = "TEN_MINUTES"
)

