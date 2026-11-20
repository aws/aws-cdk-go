package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContextStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContextStoreProps := &CfnContextStoreProps{
//   	Description: jsii.String("description"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-contextstore.html
//
type CfnContextStoreProps struct {
	// The description of the context store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-contextstore.html#cfn-networkmanager-contextstore-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the global network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-contextstore.html#cfn-networkmanager-contextstore-globalnetworkid
	//
	GlobalNetworkId *string `field:"optional" json:"globalNetworkId" yaml:"globalNetworkId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-contextstore.html#cfn-networkmanager-contextstore-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

