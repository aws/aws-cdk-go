package previewawsbedrockagentcoremixins


// Network mode configuration for VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-vpcconfig.html
//
type CfnRuntimePropsMixin_VpcConfigProperty struct {
	// Security groups for VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-vpcconfig.html#cfn-bedrockagentcore-runtime-vpcconfig-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Subnets for VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-vpcconfig.html#cfn-bedrockagentcore-runtime-vpcconfig-subnets
	//
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

