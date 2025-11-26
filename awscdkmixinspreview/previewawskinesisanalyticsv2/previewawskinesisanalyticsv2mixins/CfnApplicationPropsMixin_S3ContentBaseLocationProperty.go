package previewawskinesisanalyticsv2mixins


// The base location of the Amazon Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3ContentBaseLocationProperty := &S3ContentBaseLocationProperty{
//   	BasePath: jsii.String("basePath"),
//   	BucketArn: jsii.String("bucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentbaselocation.html
//
type CfnApplicationPropsMixin_S3ContentBaseLocationProperty struct {
	// The base path for the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentbaselocation.html#cfn-kinesisanalyticsv2-application-s3contentbaselocation-basepath
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The Amazon Resource Name (ARN) of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentbaselocation.html#cfn-kinesisanalyticsv2-application-s3contentbaselocation-bucketarn
	//
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
}

