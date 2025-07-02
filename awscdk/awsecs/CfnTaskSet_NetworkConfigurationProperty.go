package awsecs


// The network configuration for a task or service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	AwsVpcConfiguration: &AwsVpcConfigurationProperty{
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//
//   		// the properties below are optional
//   		AssignPublicIp: jsii.String("assignPublicIp"),
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-networkconfiguration.html
//
type CfnTaskSet_NetworkConfigurationProperty struct {
	// The VPC subnets and security groups that are associated with a task.
	//
	// > All specified subnets and security groups must be from the same VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskset-networkconfiguration.html#cfn-ecs-taskset-networkconfiguration-awsvpcconfiguration
	//
	AwsVpcConfiguration interface{} `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
}

