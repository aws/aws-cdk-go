package awscloudfront


// Attributes used to import a Distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudFrontWebDistributionAttributes := &cloudFrontWebDistributionAttributes{
//   	distributionId: jsii.String("distributionId"),
//   	domainName: jsii.String("domainName"),
//   }
//
type CloudFrontWebDistributionAttributes struct {
	// The distribution ID for this distribution.
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// The generated domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

