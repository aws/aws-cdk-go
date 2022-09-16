package awscloudfront


// Attributes used to import a Distribution.
//
// Example:
//   // Using a reference to an imported Distribution
//   distribution := cloudfront.distribution.fromDistributionAttributes(this, jsii.String("ImportedDist"), &distributionAttributes{
//   	domainName: jsii.String("d111111abcdef8.cloudfront.net"),
//   	distributionId: jsii.String("012345ABCDEF"),
//   })
//
type DistributionAttributes struct {
	// The distribution ID for this distribution.
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// The generated domain name of the Distribution, such as d111111abcdef8.cloudfront.net.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

