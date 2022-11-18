package awsec2


// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//
//   vpc.addFlowLog(jsii.String("FlowLogS3"), &flowLogOptions{
//   	destination: ec2.flowLogDestination.toS3(),
//   })
//
//   // Only reject traffic and interval every minute.
//   vpc.addFlowLog(jsii.String("FlowLogCloudWatch"), &flowLogOptions{
//   	trafficType: ec2.flowLogTrafficType_REJECT,
//   	maxAggregationInterval: flowLogMaxAggregationInterval_ONE_MINUTE,
//   })
//
type FlowLogMaxAggregationInterval string

const (
	// 1 minute (60 seconds).
	FlowLogMaxAggregationInterval_ONE_MINUTE FlowLogMaxAggregationInterval = "ONE_MINUTE"
	// 10 minutes (600 seconds).
	FlowLogMaxAggregationInterval_TEN_MINUTES FlowLogMaxAggregationInterval = "TEN_MINUTES"
)

