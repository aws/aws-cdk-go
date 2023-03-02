package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnSite`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSiteProps := &cfnSiteProps{
//   	globalNetworkId: jsii.String("globalNetworkId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	location: &locationProperty{
//   		address: jsii.String("address"),
//   		latitude: jsii.String("latitude"),
//   		longitude: jsii.String("longitude"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

