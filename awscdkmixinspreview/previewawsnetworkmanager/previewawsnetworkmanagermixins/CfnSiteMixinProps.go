package previewawsnetworkmanagermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSitePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSiteMixinProps := &CfnSiteMixinProps{
//   	Description: jsii.String("description"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	Location: &LocationProperty{
//   		Address: jsii.String("address"),
//   		Latitude: jsii.String("latitude"),
//   		Longitude: jsii.String("longitude"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-site.html
//
type CfnSiteMixinProps struct {
	// A description of your site.
	//
	// Constraints: Maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-site.html#cfn-networkmanager-site-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the global network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-site.html#cfn-networkmanager-site-globalnetworkid
	//
	GlobalNetworkId *string `field:"optional" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The site location.
	//
	// This information is used for visualization in the Network Manager console. If you specify the address, the latitude and longitude are automatically calculated.
	//
	// - `Address` : The physical address of the site.
	// - `Latitude` : The latitude of the site.
	// - `Longitude` : The longitude of the site.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-site.html#cfn-networkmanager-site-location
	//
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// The tags for the site.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-site.html#cfn-networkmanager-site-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

