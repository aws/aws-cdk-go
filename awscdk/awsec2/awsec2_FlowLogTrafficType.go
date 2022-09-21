package awsec2


// The type of VPC traffic to log.
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
type FlowLogTrafficType string

const (
	// Only log accepts.
	FlowLogTrafficType_ACCEPT FlowLogTrafficType = "ACCEPT"
	// Log all requests.
	FlowLogTrafficType_ALL FlowLogTrafficType = "ALL"
	// Only log rejects.
	FlowLogTrafficType_REJECT FlowLogTrafficType = "REJECT"
)

