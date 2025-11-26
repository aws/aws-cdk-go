package previewawsnetworkmanagermixins


// Describes a location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   locationProperty := &LocationProperty{
//   	Address: jsii.String("address"),
//   	Latitude: jsii.String("latitude"),
//   	Longitude: jsii.String("longitude"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-location.html
//
type CfnDevicePropsMixin_LocationProperty struct {
	// The physical address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-location.html#cfn-networkmanager-device-location-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The latitude.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-location.html#cfn-networkmanager-device-location-latitude
	//
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// The longitude.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-device-location.html#cfn-networkmanager-device-location-longitude
	//
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}

