package previewawss3mixins


// Describes the versioning state of an Amazon S3 bucket.
//
// For more information, see [PUT Bucket versioning](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTVersioningStatus.html) in the *Amazon S3 API Reference* .
//
// Keep the following timing in mind when enabling, suspending, or transitioning between versioning states:
//
// - *Enabling versioning* - Changes may take up to 15 minutes to propagate across all AWS regions for full consistency.
// - *Suspending versioning* - Takes effect immediately with no propagation delay.
// - *Transitioning between states* - Any change from Suspended to Enabled has a 15-minute delay.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   versioningConfigurationProperty := &VersioningConfigurationProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfiguration.html
//
type CfnBucketPropsMixin_VersioningConfigurationProperty struct {
	// The versioning state of the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfiguration.html#cfn-s3-bucket-versioningconfiguration-status
	//
	// Default: - "Suspended".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

