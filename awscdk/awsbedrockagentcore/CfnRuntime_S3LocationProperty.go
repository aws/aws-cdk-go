package awsbedrockagentcore


// S3 Location Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Prefix: jsii.String("prefix"),
//
//   	// the properties below are optional
//   	VersionId: jsii.String("versionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-s3location.html
//
type CfnRuntime_S3LocationProperty struct {
	// S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-s3location.html#cfn-bedrockagentcore-runtime-s3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// S3 object key prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-s3location.html#cfn-bedrockagentcore-runtime-s3location-prefix
	//
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// S3 object version ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-s3location.html#cfn-bedrockagentcore-runtime-s3location-versionid
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

