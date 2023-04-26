package awsec2


// Properties of a VPC Flow Log.
//
// Example:
//   var vpc vpc
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyCustomLogGroup"))
//
//   role := iam.NewRole(this, jsii.String("MyCustomRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("vpc-flow-logs.amazonaws.com")),
//   })
//
//   ec2.NewFlowLog(this, jsii.String("FlowLog"), &FlowLogProps{
//   	ResourceType: ec2.FlowLogResourceType_FromVpc(vpc),
//   	Destination: ec2.FlowLogDestination_ToCloudWatchLogs(logGroup, role),
//   })
//
// Experimental.
type FlowLogProps struct {
	// Specifies the type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3.
	// Experimental.
	Destination FlowLogDestination `field:"optional" json:"destination" yaml:"destination"`
	// The type of traffic to log.
	//
	// You can log traffic that the resource accepts or rejects, or all traffic.
	// Experimental.
	TrafficType FlowLogTrafficType `field:"optional" json:"trafficType" yaml:"trafficType"`
	// The type of resource for which to create the flow log.
	// Experimental.
	ResourceType FlowLogResourceType `field:"required" json:"resourceType" yaml:"resourceType"`
	// The name of the FlowLog.
	//
	// It is not recommended to use an explicit name.
	// Experimental.
	FlowLogName *string `field:"optional" json:"flowLogName" yaml:"flowLogName"`
}

