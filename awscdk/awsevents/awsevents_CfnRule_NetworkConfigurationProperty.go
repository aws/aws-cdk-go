package awsevents


// This structure specifies the network configuration for an ECS task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	awsVpcConfiguration: &awsVpcConfigurationProperty{
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//
//   		// the properties below are optional
//   		assignPublicIp: jsii.String("assignPublicIp"),
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   }
//
type CfnRule_NetworkConfigurationProperty struct {
	// Use this structure to specify the VPC subnets and security groups for the task, and whether a public IP address is to be used.
	//
	// This structure is relevant only for ECS tasks that use the `awsvpc` network mode.
	AwsVpcConfiguration interface{} `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
}

