package awskinesisanalytics


// The base location of the Amazon Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ContentBaseLocationProperty := &S3ContentBaseLocationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//
//   	// the properties below are optional
//   	BasePath: jsii.String("basePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentbaselocation.html
//
type CfnApplicationV2_S3ContentBaseLocationProperty struct {
	// The Amazon Resource Name (ARN) of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentbaselocation.html#cfn-kinesisanalyticsv2-application-s3contentbaselocation-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The base path for the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentbaselocation.html#cfn-kinesisanalyticsv2-application-s3contentbaselocation-basepath
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
}

