package awsec2


// Options to add a flow log to a VPC.
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
type FlowLogOptions struct {
	// Specifies the type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3.
	Destination FlowLogDestination `field:"optional" json:"destination" yaml:"destination"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	MaxAggregationInterval FlowLogMaxAggregationInterval `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The type of traffic to log.
	//
	// You can log traffic that the resource accepts or rejects, or all traffic.
	TrafficType FlowLogTrafficType `field:"optional" json:"trafficType" yaml:"trafficType"`
}

