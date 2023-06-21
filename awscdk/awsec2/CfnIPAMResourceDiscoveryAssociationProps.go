package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIPAMResourceDiscoveryAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMResourceDiscoveryAssociationProps := &CfnIPAMResourceDiscoveryAssociationProps{
//   	IpamId: jsii.String("ipamId"),
//   	IpamResourceDiscoveryId: jsii.String("ipamResourceDiscoveryId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIPAMResourceDiscoveryAssociationProps struct {
	// The IPAM ID.
	IpamId *string `field:"required" json:"ipamId" yaml:"ipamId"`
	// The resource discovery ID.
	IpamResourceDiscoveryId *string `field:"required" json:"ipamResourceDiscoveryId" yaml:"ipamResourceDiscoveryId"`
	// A tag is a label that you assign to an AWS resource.
	//
	// Each tag consists of a key and an optional value. You can use tags to search and filter your resources or track your AWS costs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

