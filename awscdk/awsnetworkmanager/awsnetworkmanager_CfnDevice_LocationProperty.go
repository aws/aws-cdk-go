package awsnetworkmanager


// Describes a location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &LocationProperty{
//   	Address: jsii.String("address"),
//   	Latitude: jsii.String("latitude"),
//   	Longitude: jsii.String("longitude"),
//   }
//
type CfnDevice_LocationProperty struct {
	// The physical address.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The latitude.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// The longitude.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}

