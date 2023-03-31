package awscloudfront


// A complex type that contains information about the Amazon S3 bucket from which you want CloudFront to get your media files for distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OriginProperty := &s3OriginProperty{
//   	domainName: jsii.String("domainName"),
//   	originAccessIdentity: jsii.String("originAccessIdentity"),
//   }
//
type CfnStreamingDistribution_S3OriginProperty struct {
	// The DNS name of the Amazon S3 origin.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The CloudFront origin access identity to associate with the distribution.
	//
	// Use an origin access identity to configure the distribution so that end users can only access objects in an Amazon S3 bucket through CloudFront.
	//
	// If you want end users to be able to access objects using either the CloudFront URL or the Amazon S3 URL, specify an empty `OriginAccessIdentity` element.
	//
	// To delete the origin access identity from an existing distribution, update the distribution configuration and include an empty `OriginAccessIdentity` element.
	//
	// To replace the origin access identity, update the distribution configuration and specify the new origin access identity.
	//
	// For more information, see [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide* .
	OriginAccessIdentity *string `field:"required" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

