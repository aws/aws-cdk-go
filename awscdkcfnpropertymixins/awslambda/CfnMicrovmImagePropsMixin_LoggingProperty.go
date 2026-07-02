package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   loggingProperty := &LoggingProperty{
//   	CloudWatch: &CloudWatchLoggingProperty{
//   		LogGroup: jsii.String("logGroup"),
//   		LogStream: jsii.String("logStream"),
//   	},
//   	Disabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-logging.html
//
type CfnMicrovmImagePropsMixin_LoggingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-logging.html#cfn-lambda-microvmimage-logging-cloudwatch
	//
	CloudWatch interface{} `field:"optional" json:"cloudWatch" yaml:"cloudWatch"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-logging.html#cfn-lambda-microvmimage-logging-disabled
	//
	// Default: - false.
	//
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

