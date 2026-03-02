package awsdirectconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLag`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLagProps := &CfnLagProps{
//   	ConnectionsBandwidth: jsii.String("connectionsBandwidth"),
//   	LagName: jsii.String("lagName"),
//   	Location: jsii.String("location"),
//
//   	// the properties below are optional
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
type CfnLagProps struct {
	// The bandwidth of the individual physical dedicated connections bundled by the LAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-connectionsbandwidth
	//
	ConnectionsBandwidth *string `field:"required" json:"connectionsBandwidth" yaml:"connectionsBandwidth"`
	// The name of the LAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-lagname
	//
	LagName *string `field:"required" json:"lagName" yaml:"lagName"`
	// The location for the LAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-lag.html#cfn-directconnect-lag-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
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

