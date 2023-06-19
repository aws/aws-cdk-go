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
//   vpc.addFlowLog(jsii.String("FlowLogCloudWatch"), &FlowLogOptions{
//   	TrafficType: ec2.FlowLogTrafficType_REJECT,
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

