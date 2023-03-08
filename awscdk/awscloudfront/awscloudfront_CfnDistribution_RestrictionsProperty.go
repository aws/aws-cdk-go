package awscloudfront


// A complex type that identifies ways in which you want to restrict distribution of your content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restrictionsProperty := &restrictionsProperty{
//   	geoRestriction: &geoRestrictionProperty{
//   		restrictionType: jsii.String("restrictionType"),
//
//   		// the properties below are optional
//   		locations: []*string{
//   			jsii.String("locations"),
//   		},
//   	},
//   }
//
type CfnDistribution_RestrictionsProperty struct {
	// A complex type that controls the countries in which your content is distributed.
	//
	// CloudFront determines the location of your users using `MaxMind` GeoIP databases. To disable geo restriction, remove the [Restrictions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-restrictions) property from your stack template.
	GeoRestriction interface{} `field:"required" json:"geoRestriction" yaml:"geoRestriction"`
}

