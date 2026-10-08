package awslambda


// The logging configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigProperty := &LoggingConfigProperty{
//   	ApplicationLogLevel: jsii.String("applicationLogLevel"),
//   	LogGroup: jsii.String("logGroup"),
//   	SystemLogLevel: jsii.String("systemLogLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-loggingconfig.html
//
type CfnWebFunctionRevision_LoggingConfigProperty struct {
	// The application log level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-loggingconfig.html#cfn-lambda-webfunctionrevision-loggingconfig-applicationloglevel
	//
	ApplicationLogLevel *string `field:"optional" json:"applicationLogLevel" yaml:"applicationLogLevel"`
	// The CloudWatch log group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-loggingconfig.html#cfn-lambda-webfunctionrevision-loggingconfig-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The system log level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-loggingconfig.html#cfn-lambda-webfunctionrevision-loggingconfig-systemloglevel
	//
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

