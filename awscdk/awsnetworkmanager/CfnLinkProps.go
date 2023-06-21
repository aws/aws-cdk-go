package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLinkProps := &CfnLinkProps{
//   	Bandwidth: &BandwidthProperty{
//   		DownloadSpeed: jsii.Number(123),
//   		UploadSpeed: jsii.Number(123),
//   	},
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	SiteId: jsii.String("siteId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Provider: jsii.String("provider"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
type CfnLinkProps struct {
	// The bandwidth for the link.
	Bandwidth interface{} `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// The ID of the global network.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the site.
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// A description of the link.
	//
	// Constraints: Maximum length of 256 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The provider of the link.
	//
	// Constraints: Maximum length of 128 characters. Cannot include the following characters: | \ ^
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// The tags for the link.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the link.
	//
	// Constraints: Maximum length of 128 characters. Cannot include the following characters: | \ ^
	Type *string `field:"optional" json:"type" yaml:"type"`
}

