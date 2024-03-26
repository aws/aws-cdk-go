package awskafkaconnect


// The S3 bucket Amazon Resource Name (ARN), file key, and object version of the plugin file stored in Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	FileKey: jsii.String("fileKey"),
//
//   	// the properties below are optional
//   	ObjectVersion: jsii.String("objectVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-s3location.html
//
type CfnCustomPlugin_S3LocationProperty struct {
	// The Amazon Resource Name (ARN) of an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-s3location.html#cfn-kafkaconnect-customplugin-s3location-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The file key for an object in an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-s3location.html#cfn-kafkaconnect-customplugin-s3location-filekey
	//
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// The version of an object in an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kafkaconnect-customplugin-s3location.html#cfn-kafkaconnect-customplugin-s3location-objectversion
	//
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

