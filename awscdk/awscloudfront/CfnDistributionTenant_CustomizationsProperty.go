package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customizationsProperty := &CustomizationsProperty{
//   	Certificate: &CertificateProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	GeoRestrictions: &GeoRestrictionCustomizationProperty{
//   		Locations: []*string{
//   			jsii.String("locations"),
//   		},
//   		RestrictionType: jsii.String("restrictionType"),
//   	},
//   	WebAcl: &WebAclCustomizationProperty{
//   		Action: jsii.String("action"),
//   		Arn: jsii.String("arn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-customizations.html
//
type CfnDistributionTenant_CustomizationsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-customizations.html#cfn-cloudfront-distributiontenant-customizations-certificate
	//
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-customizations.html#cfn-cloudfront-distributiontenant-customizations-georestrictions
	//
	GeoRestrictions interface{} `field:"optional" json:"geoRestrictions" yaml:"geoRestrictions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-customizations.html#cfn-cloudfront-distributiontenant-customizations-webacl
	//
	WebAcl interface{} `field:"optional" json:"webAcl" yaml:"webAcl"`
}

