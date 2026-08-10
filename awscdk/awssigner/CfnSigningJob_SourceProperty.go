package awssigner


// Information about the S3 bucket where unsigned code is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	S3: &S3SourceProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-source.html
//
type CfnSigningJob_SourceProperty struct {
	// Information about the Amazon S3 bucket where unsigned code is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-source.html#cfn-signer-signingjob-source-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

