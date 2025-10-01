package awsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOdbPeeringConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOdbPeeringConnectionProps := &CfnOdbPeeringConnectionProps{
//   	DisplayName: jsii.String("displayName"),
//   	OdbNetworkId: jsii.String("odbNetworkId"),
//   	PeerNetworkId: jsii.String("peerNetworkId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbpeeringconnection.html
//
type CfnOdbPeeringConnectionProps struct {
	// The display name of the ODB peering connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbpeeringconnection.html#cfn-odb-odbpeeringconnection-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The unique identifier of the ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbpeeringconnection.html#cfn-odb-odbpeeringconnection-odbnetworkid
	//
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// The unique identifier of the peer network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbpeeringconnection.html#cfn-odb-odbpeeringconnection-peernetworkid
	//
	PeerNetworkId *string `field:"optional" json:"peerNetworkId" yaml:"peerNetworkId"`
	// Tags to assign to the Odb peering connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbpeeringconnection.html#cfn-odb-odbpeeringconnection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

