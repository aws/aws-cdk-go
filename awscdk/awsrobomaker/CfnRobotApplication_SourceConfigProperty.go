package awsrobomaker


// Information about a source configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfigProperty := &SourceConfigProperty{
//   	Architecture: jsii.String("architecture"),
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Key: jsii.String("s3Key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-robotapplication-sourceconfig.html
//
type CfnRobotApplication_SourceConfigProperty struct {
	// The target processor architecture for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-robotapplication-sourceconfig.html#cfn-robomaker-robotapplication-sourceconfig-architecture
	//
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The Amazon S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-robotapplication-sourceconfig.html#cfn-robomaker-robotapplication-sourceconfig-s3bucket
	//
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The s3 object key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-robotapplication-sourceconfig.html#cfn-robomaker-robotapplication-sourceconfig-s3key
	//
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

