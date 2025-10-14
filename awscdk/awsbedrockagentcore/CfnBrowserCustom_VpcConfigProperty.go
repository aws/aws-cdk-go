package awsbedrockagentcore


// Network mode configuration for VPC.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-vpcconfig.html
//
type CfnBrowserCustom_VpcConfigProperty struct {
	// Security groups for VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-vpcconfig.html#cfn-bedrockagentcore-browsercustom-vpcconfig-securitygroups
	//
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Subnets for VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-vpcconfig.html#cfn-bedrockagentcore-browsercustom-vpcconfig-subnets
	//
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

