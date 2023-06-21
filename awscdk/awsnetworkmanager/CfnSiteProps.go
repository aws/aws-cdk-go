package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSite`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSiteProps := &CfnSiteProps{
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Location: &LocationProperty{
//   		Address: jsii.String("address"),
//   		Latitude: jsii.String("latitude"),
//   		Longitude: jsii.String("longitude"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSiteProps struct {
	// The ID of the global network.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// A description of your site.
	//
	// Constraints: Maximum length of 256 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The site location.
	//
	// This information is used for visualization in the Network Manager console. If you specify the address, the latitude and longitude are automatically calculated.
	//
	// - `Address` : The physical address of the site.
	// - `Latitude` : The latitude of the site.
	// - `Longitude` : The longitude of the site.
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// The tags for the site.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

