package awsroute53


// A complex type that contains information about a geographic location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoLocationProperty := &GeoLocationProperty{
//   	ContinentCode: jsii.String("continentCode"),
//   	CountryCode: jsii.String("countryCode"),
//   	SubdivisionCode: jsii.String("subdivisionCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
//
type CfnRecordSet_GeoLocationProperty struct {
	// For geolocation resource record sets, a two-letter abbreviation that identifies a continent. Route 53 supports the following continent codes:.
	//
	// - *AF* : Africa
	// - *AN* : Antarctica
	// - *AS* : Asia
	// - *EU* : Europe
	// - *OC* : Oceania
	// - *NA* : North America
	// - *SA* : South America
	//
	// Constraint: Specifying `ContinentCode` with either `CountryCode` or `SubdivisionCode` returns an `InvalidInput` error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-continentcode
	//
	ContinentCode *string `field:"optional" json:"continentCode" yaml:"continentCode"`
	// For geolocation resource record sets, the two-letter code for a country.
	//
	// Route 53 uses the two-letter country codes that are specified in [ISO standard 3166-1 alpha-2](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-countrycode
	//
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// For geolocation resource record sets, the two-letter code for a state of the United States.
	//
	// Route 53 doesn't support any other values for `SubdivisionCode` . For a list of state abbreviations, see [Appendix B: Two–Letter State and Possession Abbreviations](https://docs.aws.amazon.com/https://pe.usps.com/text/pub28/28apb.htm) on the United States Postal Service website.
	//
	// If you specify `subdivisioncode` , you must also specify `US` for `CountryCode` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html#cfn-route53-recordset-geolocation-subdivisioncode
	//
	SubdivisionCode *string `field:"optional" json:"subdivisionCode" yaml:"subdivisionCode"`
}

