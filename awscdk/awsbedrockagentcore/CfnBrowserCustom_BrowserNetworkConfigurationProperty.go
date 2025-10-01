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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browsernetworkconfiguration.html
//
type CfnBrowserCustom_BrowserNetworkConfigurationProperty struct {
	// The network mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-browsercustom-browsernetworkconfiguration.html#cfn-bedrockagentcore-browsercustom-browsernetworkconfiguration-networkmode
	//
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
}

