package awsecs


// The `NetworkConfiguration` property specifies an object representing the network configuration for a task or service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	awsvpcConfiguration: &awsVpcConfigurationProperty{
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
type CfnService_NetworkConfigurationProperty struct {
	// The VPC subnets and security groups that are associated with a task.
	//
	// > All specified subnets and security groups must be from the same VPC.
	AwsvpcConfiguration interface{} `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
}

