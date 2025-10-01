package awsevents


// This structure specifies the VPC subnets and security groups for the task, and whether a public IP address is to be used.
//
// This structure is relevant only for ECS tasks that use the `awsvpc` network mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsVpcConfigurationProperty := &AwsVpcConfigurationProperty{
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	AssignPublicIp: jsii.String("assignPublicIp"),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-awsvpcconfiguration.html
//
type CfnRule_AwsVpcConfigurationProperty struct {
	// Specifies the subnets associated with the task.
	//
	// These subnets must all be in the same VPC. You can specify as many as 16 subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-awsvpcconfiguration.html#cfn-events-rule-awsvpcconfiguration-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// You can specify `ENABLED` only when `LaunchType` in `EcsParameters` is set to `FARGATE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-awsvpcconfiguration.html#cfn-events-rule-awsvpcconfiguration-assignpublicip
	//
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Specifies the security groups associated with the task.
	//
	// These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-awsvpcconfiguration.html#cfn-events-rule-awsvpcconfiguration-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

