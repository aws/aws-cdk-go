package awss3


// Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket.
//
// If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlTranslationProperty := &AccessControlTranslationProperty{
//   	Owner: jsii.String("owner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accesscontroltranslation.html
//
type CfnBucket_AccessControlTranslationProperty struct {
	// Specifies the replica ownership.
	//
	// For default and valid values, see [PUT bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) in the *Amazon S3 API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-accesscontroltranslation.html#cfn-s3-bucket-accesscontroltranslation-owner
	//
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

