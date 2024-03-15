package awsecs


// The network configuration for a task or service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   		AssignPublicIp: jsii.String("assignPublicIp"),
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-networkconfiguration.html
//
type CfnService_NetworkConfigurationProperty struct {
	// The VPC subnets and security groups that are associated with a task.
	//
	// > All specified subnets and security groups must be from the same VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-networkconfiguration.html#cfn-ecs-service-networkconfiguration-awsvpcconfiguration
	//
	AwsvpcConfiguration interface{} `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
}

