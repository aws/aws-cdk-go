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
//   awsVpcConfigurationProperty := &awsVpcConfigurationProperty{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	assignPublicIp: jsii.String("assignPublicIp"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
type CfnRule_AwsVpcConfigurationProperty struct {
	// Specifies the subnets associated with the task.
	//
	// These subnets must all be in the same VPC. You can specify as many as 16 subnets.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// You can specify `ENABLED` only when `LaunchType` in `EcsParameters` is set to `FARGATE` .
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Specifies the security groups associated with the task.
	//
	// These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

