package awsoutposts


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addressProperty := &AddressProperty{
//   	AddressLine1: jsii.String("addressLine1"),
//   	City: jsii.String("city"),
//   	ContactName: jsii.String("contactName"),
//   	ContactPhoneNumber: jsii.String("contactPhoneNumber"),
//   	CountryCode: jsii.String("countryCode"),
//   	PostalCode: jsii.String("postalCode"),
//   	StateOrRegion: jsii.String("stateOrRegion"),
//
//   	// the properties below are optional
//   	AddressLine2: jsii.String("addressLine2"),
//   	AddressLine3: jsii.String("addressLine3"),
//   	DistrictOrCounty: jsii.String("districtOrCounty"),
//   	Municipality: jsii.String("municipality"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html
//
type CfnSite_AddressProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline1
	//
	AddressLine1 *string `field:"required" json:"addressLine1" yaml:"addressLine1"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-city
	//
	City *string `field:"required" json:"city" yaml:"city"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-contactname
	//
	ContactName *string `field:"required" json:"contactName" yaml:"contactName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-contactphonenumber
	//
	ContactPhoneNumber *string `field:"required" json:"contactPhoneNumber" yaml:"contactPhoneNumber"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-countrycode
	//
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-postalcode
	//
	PostalCode *string `field:"required" json:"postalCode" yaml:"postalCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-stateorregion
	//
	StateOrRegion *string `field:"required" json:"stateOrRegion" yaml:"stateOrRegion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline2
	//
	AddressLine2 *string `field:"optional" json:"addressLine2" yaml:"addressLine2"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-addressline3
	//
	AddressLine3 *string `field:"optional" json:"addressLine3" yaml:"addressLine3"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-districtorcounty
	//
	DistrictOrCounty *string `field:"optional" json:"districtOrCounty" yaml:"districtOrCounty"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-outposts-site-address.html#cfn-outposts-site-address-municipality
	//
	Municipality *string `field:"optional" json:"municipality" yaml:"municipality"`
}

