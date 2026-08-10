package awssigner


// Information about the Amazon S3 bucket where unsigned code is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceProperty := &S3SourceProperty{
//   	BucketName: jsii.String("bucketName"),
//   	Key: jsii.String("key"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-s3source.html
//
type CfnSigningJob_S3SourceProperty struct {
	// Name of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-s3source.html#cfn-signer-signingjob-s3source-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Key name of the bucket object that contains unsigned code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-s3source.html#cfn-signer-signingjob-s3source-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// Version of the source image in the version-enabled S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-s3source.html#cfn-signer-signingjob-s3source-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
}

