package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSiteToSiteVpnAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSiteToSiteVpnAttachmentProps := &cfnSiteToSiteVpnAttachmentProps{
//   	coreNetworkId: jsii.String("coreNetworkId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpnConnectionArn: jsii.String("vpnConnectionArn"),
//   }
//
type CfnSiteToSiteVpnAttachmentProps struct {
	// The core network ID.
	CoreNetworkId *string `field:"optional" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The tags associated with the site-to-site VPN attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the site-to-site VPN attachment.
	VpnConnectionArn *string `field:"optional" json:"vpnConnectionArn" yaml:"vpnConnectionArn"`
}

