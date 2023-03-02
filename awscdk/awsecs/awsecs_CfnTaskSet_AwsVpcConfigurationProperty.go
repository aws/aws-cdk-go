package awsecs


// The networking details for a task.
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
type CfnTaskSet_AwsVpcConfigurationProperty struct {
	// The IDs of the subnets associated with the task or service.
	//
	// There's a limit of 16 subnets that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified subnets must be from the same VPC.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Whether the task's elastic network interface receives a public IP address.
	//
	// The default value is `DISABLED` .
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// The IDs of the security groups associated with the task or service.
	//
	// If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified per `AwsVpcConfiguration` .
	//
	// > All specified security groups must be from the same VPC.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

