package awsidentitystore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   addressesItemsProperty := &AddressesItemsProperty{
//   	Country: jsii.String("country"),
//   	Formatted: jsii.String("formatted"),
//   	Locality: jsii.String("locality"),
//   	PostalCode: jsii.String("postalCode"),
//   	Primary: jsii.Boolean(false),
//   	Region: jsii.String("region"),
//   	StreetAddress: jsii.String("streetAddress"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html
//
type CfnUserPropsMixin_AddressesItemsProperty struct {
	// The country of the address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html#cfn-identitystore-user-addressesitems-country
	//
	Country *string `field:"optional" json:"country" yaml:"country"`
	// A formatted version of the address for display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html#cfn-identitystore-user-addressesitems-formatted
	//
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// A string of the address locality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html#cfn-identitystore-user-addressesitems-locality
	//
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The postal code of the address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html#cfn-identitystore-user-addressesitems-postalcode
	//
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// Whether this is the primary address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html#cfn-identitystore-user-addressesitems-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The region of the address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html#cfn-identitystore-user-addressesitems-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The street of the address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html#cfn-identitystore-user-addressesitems-streetaddress
	//
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// The type of address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-identitystore-user-addressesitems.html#cfn-identitystore-user-addressesitems-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

