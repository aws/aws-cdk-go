package awsappstream


// Describes the S3 location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &S3LocationProperty{
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Key: jsii.String("s3Key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-s3location.html
//
type CfnFleet_S3LocationProperty struct {
	// The S3 bucket of the S3 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-s3location.html#cfn-appstream-fleet-s3location-s3bucket
	//
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The S3 key of the S3 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-s3location.html#cfn-appstream-fleet-s3location-s3key
	//
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

