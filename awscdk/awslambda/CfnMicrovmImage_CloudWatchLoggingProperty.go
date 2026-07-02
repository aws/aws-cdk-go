package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLoggingProperty := &CloudWatchLoggingProperty{
//   	LogGroup: jsii.String("logGroup"),
//   	LogStream: jsii.String("logStream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-cloudwatchlogging.html
//
type CfnMicrovmImage_CloudWatchLoggingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-cloudwatchlogging.html#cfn-lambda-microvmimage-cloudwatchlogging-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-cloudwatchlogging.html#cfn-lambda-microvmimage-cloudwatchlogging-logstream
	//
	LogStream *string `field:"optional" json:"logStream" yaml:"logStream"`
}

