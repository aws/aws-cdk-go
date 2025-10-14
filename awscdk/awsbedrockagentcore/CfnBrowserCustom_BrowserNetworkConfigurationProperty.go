package awsbedrockagentcore


// The network configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   browserNetworkConfigurationProperty := &BrowserNetworkConfigurationProperty{
//   	NetworkMode: jsii.String("networkMode"),
//
//   	// the properties below are optional
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browsernetworkconfiguration.html
//
type CfnBrowserCustom_BrowserNetworkConfigurationProperty struct {
	// The network mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browsernetworkconfiguration.html#cfn-bedrockagentcore-browsercustom-browsernetworkconfiguration-networkmode
	//
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
	// Network mode configuration for VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browsernetworkconfiguration.html#cfn-bedrockagentcore-browsercustom-browsernetworkconfiguration-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

