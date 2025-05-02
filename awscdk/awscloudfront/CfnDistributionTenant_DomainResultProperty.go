package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainResultProperty := &DomainResultProperty{
//   	Domain: jsii.String("domain"),
//   	Reason: jsii.String("reason"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-domainresult.html
//
type CfnDistributionTenant_DomainResultProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-domainresult.html#cfn-cloudfront-distributiontenant-domainresult-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-domainresult.html#cfn-cloudfront-distributiontenant-domainresult-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-domainresult.html#cfn-cloudfront-distributiontenant-domainresult-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

