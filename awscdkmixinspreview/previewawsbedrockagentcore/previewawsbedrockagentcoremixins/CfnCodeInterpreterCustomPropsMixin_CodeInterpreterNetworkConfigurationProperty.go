package previewawsbedrockagentcoremixins


// The network configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeInterpreterNetworkConfigurationProperty := &CodeInterpreterNetworkConfigurationProperty{
//   	NetworkMode: jsii.String("networkMode"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration.html
//
type CfnCodeInterpreterCustomPropsMixin_CodeInterpreterNetworkConfigurationProperty struct {
	// The network mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration.html#cfn-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration-networkmode
	//
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Network mode configuration for VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration.html#cfn-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

