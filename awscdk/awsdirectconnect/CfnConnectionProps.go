package awsdirectconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectionProps := &CfnConnectionProps{
//   	Bandwidth: jsii.String("bandwidth"),
//   	ConnectionName: jsii.String("connectionName"),
//   	Location: jsii.String("location"),
//
//   	// the properties below are optional
//   	LagId: jsii.String("lagId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html
//
type CfnConnectionProps struct {
	// The bandwidth of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-bandwidth
	//
	Bandwidth *string `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// The name of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-connectionname
	//
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The location of the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-location
	//
	Location *string `field:"required" json:"location" yaml:"location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-lagid
	//
	LagId interface{} `field:"optional" json:"lagId" yaml:"lagId"`
	// The name of the service provider associated with the requested connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-providername
	//
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Indicates whether you want the connection to support MAC Security (MACsec).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-requestmacsec
	//
	RequestMacSec interface{} `field:"optional" json:"requestMacSec" yaml:"requestMacSec"`
	// The tags associated with the connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-connection.html#cfn-directconnect-connection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

