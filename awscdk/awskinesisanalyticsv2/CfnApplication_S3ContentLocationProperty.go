package awskinesisanalyticsv2


// The location of an application or a custom artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ContentLocationProperty := &S3ContentLocationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	FileKey: jsii.String("fileKey"),
//
//   	// the properties below are optional
//   	ObjectVersion: jsii.String("objectVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html
//
type CfnApplication_S3ContentLocationProperty struct {
	// The Amazon Resource Name (ARN) for the S3 bucket containing the application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The file key for the object containing the application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-filekey
	//
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// The version of the object containing the application code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-objectversion
	//
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

