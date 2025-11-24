package mixinsawscloudfront


// The details about the domain result.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   domainResultProperty := &DomainResultProperty{
//   	Domain: jsii.String("domain"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-domainresult.html
//
type CfnDistributionTenantPropsMixin_DomainResultProperty struct {
	// The specified domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-domainresult.html#cfn-cloudfront-distributiontenant-domainresult-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Whether the domain is active or inactive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-domainresult.html#cfn-cloudfront-distributiontenant-domainresult-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

