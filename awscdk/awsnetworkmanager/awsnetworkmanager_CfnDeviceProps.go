package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceProps := &cfnDeviceProps{
//   	globalNetworkId: jsii.String("globalNetworkId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	location: &locationProperty{
//   		address: jsii.String("address"),
//   		latitude: jsii.String("latitude"),
//   		longitude: jsii.String("longitude"),
//   	},
//   	model: jsii.String("model"),
//   	serialNumber: jsii.String("serialNumber"),
//   	siteId: jsii.String("siteId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   	vendor: jsii.String("vendor"),
//   }
//
type CfnDeviceProps struct {
	// The ID of the global network.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// A description of the device.
	//
	// Constraints: Maximum length of 256 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The site location.
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// The model of the device.
	//
	// Constraints: Maximum length of 128 characters.
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The serial number of the device.
	//
	// Constraints: Maximum length of 128 characters.
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// The site ID.
	SiteId *string `field:"optional" json:"siteId" yaml:"siteId"`
	// The tags for the device.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The device type.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The vendor of the device.
	//
	// Constraints: Maximum length of 128 characters.
	Vendor *string `field:"optional" json:"vendor" yaml:"vendor"`
}

