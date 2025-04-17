package awsec2


// Options to add a flow log to a VPC.
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
type FlowLogOptions struct {
	// Specifies the type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3.
	// Default: FlowLogDestinationType.toCloudWatchLogs()
	//
	Destination FlowLogDestination `field:"optional" json:"destination" yaml:"destination"`
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// If multiple fields are specified, they will be separated by spaces. For full control over the literal log format
	// string, pass a single field constructed with `LogFormat.custom()`.
	//
	// See https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records
	// Default: - default log format is used.
	//
	LogFormat *[]LogFormat `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	//
	// When creating flow logs for a Transit Gateway or Transit Gateway Attachment,
	// this property must be ONE_MINUTES.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-maxaggregationinterval
	//
	// Default: - FlowLogMaxAggregationInterval.ONE_MINUTES if creating flow logs for Transit Gateway, otherwise FlowLogMaxAggregationInterval.TEN_MINUTES.
	//
	MaxAggregationInterval FlowLogMaxAggregationInterval `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The type of traffic to log.
	//
	// You can log traffic that the resource accepts or rejects, or all traffic.
	// When the target is either `TransitGateway` or `TransitGatewayAttachment`, setting the traffic type is not possible.
	// See: https://docs.aws.amazon.com/vpc/latest/tgw/working-with-flow-logs.html
	//
	// Default: ALL.
	//
	TrafficType FlowLogTrafficType `field:"optional" json:"trafficType" yaml:"trafficType"`
}

