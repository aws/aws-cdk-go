package awscloudfront


// The customizations that you specified for the distribution tenant for geographic restrictions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoRestrictionCustomizationProperty := &GeoRestrictionCustomizationProperty{
//   	Locations: []*string{
//   		jsii.String("locations"),
//   	},
//   	RestrictionType: jsii.String("restrictionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-georestrictioncustomization.html
//
type CfnDistributionTenant_GeoRestrictionCustomizationProperty struct {
	// The locations for geographic restrictions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-georestrictioncustomization.html#cfn-cloudfront-distributiontenant-georestrictioncustomization-locations
	//
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// The method that you want to use to restrict distribution of your content by country:.
	//
	// - `none` : No geographic restriction is enabled, meaning access to content is not restricted by client geo location.
	// - `blacklist` : The `Location` elements specify the countries in which you don't want CloudFront to distribute your content.
	// - `whitelist` : The `Location` elements specify the countries in which you want CloudFront to distribute your content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-georestrictioncustomization.html#cfn-cloudfront-distributiontenant-georestrictioncustomization-restrictiontype
	//
	RestrictionType *string `field:"optional" json:"restrictionType" yaml:"restrictionType"`
}

