package awsec2


// The type of VPC traffic to log.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//
//   vpc.addFlowLog(jsii.String("FlowLogS3"), &flowLogOptions{
//   	destination: ec2.flowLogDestination.toS3(),
//   })
//
//   vpc.addFlowLog(jsii.String("FlowLogCloudWatch"), &flowLogOptions{
//   	trafficType: ec2.flowLogTrafficType_REJECT,
//   })
//
// Experimental.
type FlowLogTrafficType string

const (
	// Only log accepts.
	// Experimental.
	FlowLogTrafficType_ACCEPT FlowLogTrafficType = "ACCEPT"
	// Log all requests.
	// Experimental.
	FlowLogTrafficType_ALL FlowLogTrafficType = "ALL"
	// Only log rejects.
	// Experimental.
	FlowLogTrafficType_REJECT FlowLogTrafficType = "REJECT"
)

