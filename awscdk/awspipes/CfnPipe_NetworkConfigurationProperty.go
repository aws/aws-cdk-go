package awspipes


// This structure specifies the network configuration for an Amazon ECS task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	AwsvpcConfiguration: &AwsVpcConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-networkconfiguration.html
//
type CfnPipe_NetworkConfigurationProperty struct {
	// Use this structure to specify the VPC subnets and security groups for the task, and whether a public IP address is to be used.
	//
	// This structure is relevant only for ECS tasks that use the `awsvpc` network mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-networkconfiguration.html#cfn-pipes-pipe-networkconfiguration-awsvpcconfiguration
	//
	AwsvpcConfiguration interface{} `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
}

