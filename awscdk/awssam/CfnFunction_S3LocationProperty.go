package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Version: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3location.html
//
type CfnFunction_S3LocationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3location.html#cfn-serverless-function-s3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3location.html#cfn-serverless-function-s3location-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3location.html#cfn-serverless-function-s3location-version
	//
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

