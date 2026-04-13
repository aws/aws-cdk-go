package awssecurityagent


// Customer VPC configuration that the security testing environment accesses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &VpcConfigProperty{
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	SubnetArns: []*string{
//   		jsii.String("subnetArns"),
//   	},
//   	VpcArn: jsii.String("vpcArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-vpcconfig.html
//
type CfnAgentSpace_VpcConfigProperty struct {
	// List of security group ARNs in the customer VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-vpcconfig.html#cfn-securityagent-agentspace-vpcconfig-securitygrouparns
	//
	SecurityGroupArns *[]*string `field:"optional" json:"securityGroupArns" yaml:"securityGroupArns"`
	// List of subnet ARNs in the customer VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-vpcconfig.html#cfn-securityagent-agentspace-vpcconfig-subnetarns
	//
	SubnetArns *[]*string `field:"optional" json:"subnetArns" yaml:"subnetArns"`
	// ARN of the customer VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-vpcconfig.html#cfn-securityagent-agentspace-vpcconfig-vpcarn
	//
	VpcArn *string `field:"optional" json:"vpcArn" yaml:"vpcArn"`
}

