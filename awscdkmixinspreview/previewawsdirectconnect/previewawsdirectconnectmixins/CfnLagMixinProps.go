package previewawsdirectconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLagPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLagMixinProps := &CfnLagMixinProps{
//   	ConnectionsBandwidth: jsii.String("connectionsBandwidth"),
//   	LagName: jsii.String("lagName"),
//   	Location: jsii.String("location"),
//   	MinimumLinks: jsii.Number(123),
//   	ProviderName: jsii.String("providerName"),
//   	RequestMacSec: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html
//
type CfnLagMixinProps struct {
	// The bandwidth of the individual physical dedicated connections bundled by the LAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-connectionsbandwidth
	//
	ConnectionsBandwidth *string `field:"optional" json:"connectionsBandwidth" yaml:"connectionsBandwidth"`
	// The name of the LAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-lagname
	//
	LagName *string `field:"optional" json:"lagName" yaml:"lagName"`
	// The location for the LAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The minimum number of physical dedicated connections that must be operational for the LAG itself to be operational.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-minimumlinks
	//
	MinimumLinks *float64 `field:"optional" json:"minimumLinks" yaml:"minimumLinks"`
	// The name of the service provider associated with the requested LAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-providername
	//
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Indicates whether you want the LAG to support MAC Security (MACsec).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-requestmacsec
	//
	RequestMacSec interface{} `field:"optional" json:"requestMacSec" yaml:"requestMacSec"`
	// The tags associated with the LAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

