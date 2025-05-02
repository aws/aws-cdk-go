package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDistributionTenant`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDistributionTenantProps := &CfnDistributionTenantProps{
//   	DistributionId: jsii.String("distributionId"),
//   	Domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
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
//   	Enabled: jsii.Boolean(false),
//   	ManagedCertificateRequest: &ManagedCertificateRequestProperty{
//   		CertificateTransparencyLoggingPreference: jsii.String("certificateTransparencyLoggingPreference"),
//   		PrimaryDomainName: jsii.String("primaryDomainName"),
//   		ValidationTokenHost: jsii.String("validationTokenHost"),
//   	},
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html
//
type CfnDistributionTenantProps struct {
	// The distribution's identifier.
	//
	// For example: `E1U5RQF7T870K0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-distributionid
	//
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-domains
	//
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-connectiongroupid
	//
	ConnectionGroupId *string `field:"optional" json:"connectionGroupId" yaml:"connectionGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-customizations
	//
	Customizations interface{} `field:"optional" json:"customizations" yaml:"customizations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-managedcertificaterequest
	//
	ManagedCertificateRequest interface{} `field:"optional" json:"managedCertificateRequest" yaml:"managedCertificateRequest"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A complex type that contains zero or more `Tag` elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distributiontenant.html#cfn-cloudfront-distributiontenant-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

