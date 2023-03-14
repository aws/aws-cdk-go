package awscloudfront


// A complex type that contains information about the Amazon S3 origin.
//
// If the origin is a custom origin or an S3 bucket that is configured as a website endpoint, use the `CustomOriginConfig` element instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OriginConfigProperty := &s3OriginConfigProperty{
//   	originAccessIdentity: jsii.String("originAccessIdentity"),
//   }
//
type CfnDistribution_S3OriginConfigProperty struct {
	// The CloudFront origin access identity to associate with the origin.
	//
	// Use an origin access identity to configure the origin so that viewers can *only* access objects in an Amazon S3 bucket through CloudFront. The format of the value is:
	//
	// origin-access-identity/cloudfront/ *ID-of-origin-access-identity*
	//
	// where `*ID-of-origin-access-identity*` is the value that CloudFront returned in the `ID` element when you created the origin access identity.
	//
	// If you want viewers to be able to access objects using either the CloudFront URL or the Amazon S3 URL, specify an empty `OriginAccessIdentity` element.
	//
	// To delete the origin access identity from an existing distribution, update the distribution configuration and include an empty `OriginAccessIdentity` element.
	//
	// To replace the origin access identity, update the distribution configuration and specify the new origin access identity.
	//
	// For more information about the origin access identity, see [Serving Private Content through CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide* .
	OriginAccessIdentity *string `field:"optional" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

