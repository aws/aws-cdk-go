package awscloudfront


// A complex type that controls the countries in which your content is distributed.
//
// CloudFront determines the location of your users using `MaxMind` GeoIP databases. To disable geo restriction, remove the [Restrictions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-restrictions) property from your stack template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoRestrictionProperty := &geoRestrictionProperty{
//   	restrictionType: jsii.String("restrictionType"),
//
//   	// the properties below are optional
//   	locations: []*string{
//   		jsii.String("locations"),
//   	},
//   }
//
type CfnDistribution_GeoRestrictionProperty struct {
	// The method that you want to use to restrict distribution of your content by country:.
	//
	// - `none` : No geo restriction is enabled, meaning access to content is not restricted by client geo location.
	// - `blacklist` : The `Location` elements specify the countries in which you don't want CloudFront to distribute your content.
	// - `whitelist` : The `Location` elements specify the countries in which you want CloudFront to distribute your content.
	RestrictionType *string `field:"required" json:"restrictionType" yaml:"restrictionType"`
	// A complex type that contains a `Location` element for each country in which you want CloudFront either to distribute your content ( `whitelist` ) or not distribute your content ( `blacklist` ).
	//
	// The `Location` element is a two-letter, uppercase country code for a country that you want to include in your `blacklist` or `whitelist` . Include one `Location` element for each country.
	//
	// CloudFront and `MaxMind` both use `ISO 3166` country codes. For the current list of countries and the corresponding codes, see `ISO 3166-1-alpha-2` code on the *International Organization for Standardization* website. You can also refer to the country list on the CloudFront console, which includes both country names and codes.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
}

