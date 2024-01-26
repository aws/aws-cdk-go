package awsecs


// An object representing the networking details for a task or service.
//
// For example `awsvpcConfiguration={subnets=["subnet-12344321"],securityGroups=["sg-12344321"]}`.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-awsvpcconfiguration.html
//
type CfnTaskSet_AwsVpcConfigurationProperty struct {
	// The IDs of the subnets associated with the task or service.
	//
	// There's a limit of 16 subnets that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified subnets must be from the same VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-awsvpcconfiguration.html#cfn-ecs-taskset-awsvpcconfiguration-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Whether the task's elastic network interface receives a public IP address.
	//
	// The default value is `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-awsvpcconfiguration.html#cfn-ecs-taskset-awsvpcconfiguration-assignpublicip
	//
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The IDs of the security groups associated with the task or service.
	//
	// If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified security groups must be from the same VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-awsvpcconfiguration.html#cfn-ecs-taskset-awsvpcconfiguration-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

