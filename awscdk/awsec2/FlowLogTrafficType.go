package awsec2


// The type of VPC traffic to log.
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
type FlowLogTrafficType string

const (
	// Only log accepts.
	FlowLogTrafficType_ACCEPT FlowLogTrafficType = "ACCEPT"
	// Log all requests.
	FlowLogTrafficType_ALL FlowLogTrafficType = "ALL"
	// Only log rejects.
	FlowLogTrafficType_REJECT FlowLogTrafficType = "REJECT"
)

