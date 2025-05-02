package awscloudfront


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-georestrictioncustomization.html#cfn-cloudfront-distributiontenant-georestrictioncustomization-locations
	//
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-georestrictioncustomization.html#cfn-cloudfront-distributiontenant-georestrictioncustomization-restrictiontype
	//
	RestrictionType *string `field:"optional" json:"restrictionType" yaml:"restrictionType"`
}

