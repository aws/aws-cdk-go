package mixinsawscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDistributionTenantPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionTenantMixinProps := &CfnDistributionTenantMixinProps{
//   	ConnectionGroupId: jsii.String("connectionGroupId"),
//   	Customizations: &CustomizationsProperty{
//   		Certificate: &CertificateProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		GeoRestrictions: &GeoRestrictionCustomizationProperty{
//   			Locations: []*string{
//   				jsii.String("locations"),
//   			},
//   			RestrictionType: jsii.String("restrictionType"),
//   		},
//   		WebAcl: &WebAclCustomizationProperty{
//   			Action: jsii.String("action"),
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   	DistributionId: jsii.String("distributionId"),
//   	Domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   	ManagedCertificateRequest: &ManagedCertificateRequestProperty{
//   		CertificateTransparencyLoggingPreference: jsii.String("certificateTransparencyLoggingPreference"),
//   		PrimaryDomainName: jsii.String("primaryDomainName"),
//   		ValidationTokenHost: jsii.String("validationTokenHost"),
//   	},
//   	Name: jsii.String("name"),
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html
//
type CfnDistributionTenantMixinProps struct {
	// The ID of the connection group for the distribution tenant.
	//
	// If you don't specify a connection group, CloudFront uses the default connection group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-connectiongroupid
	//
	ConnectionGroupId *string `field:"optional" json:"connectionGroupId" yaml:"connectionGroupId"`
	// Customizations for the distribution tenant.
	//
	// For each distribution tenant, you can specify the geographic restrictions, and the Amazon Resource Names (ARNs) for the ACM certificate and AWS WAF web ACL. These are specific values that you can override or disable from the multi-tenant distribution that was used to create the distribution tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-customizations
	//
	Customizations interface{} `field:"optional" json:"customizations" yaml:"customizations"`
	// The ID of the multi-tenant distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-distributionid
	//
	DistributionId *string `field:"optional" json:"distributionId" yaml:"distributionId"`
	// The domains associated with the distribution tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-domains
	//
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Indicates whether the distribution tenant is in an enabled state.
	//
	// If disabled, the distribution tenant won't serve traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// An object that represents the request for the Amazon CloudFront managed ACM certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-managedcertificaterequest
	//
	ManagedCertificateRequest interface{} `field:"optional" json:"managedCertificateRequest" yaml:"managedCertificateRequest"`
	// The name of the distribution tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of parameter values to add to the resource.
	//
	// A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

