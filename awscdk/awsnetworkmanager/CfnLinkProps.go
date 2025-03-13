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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-link.html
//
type CfnLinkProps struct {
	// The bandwidth for the link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-link.html#cfn-networkmanager-link-bandwidth
	//
	Bandwidth interface{} `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// The ID of the global network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-link.html#cfn-networkmanager-link-globalnetworkid
	//
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the site.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-link.html#cfn-networkmanager-link-siteid
	//
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// A description of the link.
	//
	// Constraints: Maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-link.html#cfn-networkmanager-link-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The provider of the link.
	//
	// Constraints: Maximum length of 128 characters. Cannot include the following characters: | \ ^
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-link.html#cfn-networkmanager-link-provider
	//
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// The tags for the link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-link.html#cfn-networkmanager-link-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the link.
	//
	// Constraints: Maximum length of 128 characters. Cannot include the following characters: | \ ^
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-link.html#cfn-networkmanager-link-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

