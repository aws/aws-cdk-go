package awsec2


// Properties of a VPC Flow Log.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyCustomLogGroup"))
//
//   role := iam.NewRole(this, jsii.String("MyCustomRole"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("vpc-flow-logs.amazonaws.com")),
//   })
//
//   ec2.NewFlowLog(this, jsii.String("FlowLog"), &flowLogProps{
//   	resourceType: ec2.flowLogResourceType.fromVpc(vpc),
//   	destination: ec2.flowLogDestination.toCloudWatchLogs(logGroup, role),
//   })
//
type FlowLogProps struct {
	// Specifies the type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3.
	Destination FlowLogDestination `field:"optional" json:"destination" yaml:"destination"`
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// If multiple fields are specified, they will be separated by spaces. For full control over the literal log format
	// string, pass a single field constructed with `LogFormat.custom()`.
	//
	// See https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records
	LogFormat *[]LogFormat `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	MaxAggregationInterval FlowLogMaxAggregationInterval `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// The type of traffic to log.
	//
	// You can log traffic that the resource accepts or rejects, or all traffic.
	TrafficType FlowLogTrafficType `field:"optional" json:"trafficType" yaml:"trafficType"`
	// The type of resource for which to create the flow log.
	ResourceType FlowLogResourceType `field:"required" json:"resourceType" yaml:"resourceType"`
	// The name of the FlowLog.
	//
	// It is not recommended to use an explicit name.
	FlowLogName *string `field:"optional" json:"flowLogName" yaml:"flowLogName"`
}

