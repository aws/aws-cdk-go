package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &VpcConfigProperty{
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-vpcconfig.html
//
type CfnHarness_VpcConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-vpcconfig.html#cfn-bedrockagentcore-harness-vpcconfig-securitygroups
	//
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-vpcconfig.html#cfn-bedrockagentcore-harness-vpcconfig-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

