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
//   cfnSiteToSiteVpnAttachmentProps := &CfnSiteToSiteVpnAttachmentProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	VpnConnectionArn: jsii.String("vpnConnectionArn"),
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
type CfnSiteToSiteVpnAttachmentProps struct {
	// `AWS::NetworkManager::SiteToSiteVpnAttachment.CoreNetworkId`.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The ARN of the site-to-site VPN attachment.
	VpnConnectionArn *string `field:"required" json:"vpnConnectionArn" yaml:"vpnConnectionArn"`
	// `AWS::NetworkManager::SiteToSiteVpnAttachment.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

