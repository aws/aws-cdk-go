package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceProps := &CfnDeviceProps{
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//
//   	// the properties below are optional
//   	AwsLocation: &AWSLocationProperty{
//   		SubnetArn: jsii.String("subnetArn"),
//   		Zone: jsii.String("zone"),
//   	},
//   	Description: jsii.String("description"),
//   	Location: &LocationProperty{
//   		Address: jsii.String("address"),
//   		Latitude: jsii.String("latitude"),
//   		Longitude: jsii.String("longitude"),
//   	},
//   	Model: jsii.String("model"),
//   	SerialNumber: jsii.String("serialNumber"),
//   	SiteId: jsii.String("siteId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	Vendor: jsii.String("vendor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html
//
type CfnDeviceProps struct {
	// The ID of the global network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-globalnetworkid
	//
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The AWS location of the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-awslocation
	//
	AwsLocation interface{} `field:"optional" json:"awsLocation" yaml:"awsLocation"`
	// A description of the device.
	//
	// Constraints: Maximum length of 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The site location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-location
	//
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// The model of the device.
	//
	// Constraints: Maximum length of 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The serial number of the device.
	//
	// Constraints: Maximum length of 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-serialnumber
	//
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// The site ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-siteid
	//
	SiteId *string `field:"optional" json:"siteId" yaml:"siteId"`
	// The tags for the device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The device type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The vendor of the device.
	//
	// Constraints: Maximum length of 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-vendor
	//
	Vendor *string `field:"optional" json:"vendor" yaml:"vendor"`
}

