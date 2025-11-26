package previewawscloudfrontmixins


// Customizations for the distribution tenant.
//
// For each distribution tenant, you can specify the geographic restrictions, and the Amazon Resource Names (ARNs) for the ACM certificate and AWS WAF web ACL. These are specific values that you can override or disable from the multi-tenant distribution that was used to create the distribution tenant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnDistributionTenantPropsMixin_CustomizationsProperty struct {
	// The Certificate Manager (ACM) certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-customizations.html#cfn-cloudfront-distributiontenant-customizations-certificate
	//
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// The geographic restrictions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-customizations.html#cfn-cloudfront-distributiontenant-customizations-georestrictions
	//
	GeoRestrictions interface{} `field:"optional" json:"geoRestrictions" yaml:"geoRestrictions"`
	// The AWS WAF web ACL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-customizations.html#cfn-cloudfront-distributiontenant-customizations-webacl
	//
	WebAcl interface{} `field:"optional" json:"webAcl" yaml:"webAcl"`
}

