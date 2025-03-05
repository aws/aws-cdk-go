package awslambda


// The function's Amazon CloudWatch Logs configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigProperty := &LoggingConfigProperty{
//   	ApplicationLogLevel: jsii.String("applicationLogLevel"),
//   	LogFormat: jsii.String("logFormat"),
//   	LogGroup: jsii.String("logGroup"),
//   	SystemLogLevel: jsii.String("systemLogLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-loggingconfig.html
//
type CfnFunction_LoggingConfigProperty struct {
	// Set this property to filter the application logs for your function that Lambda sends to CloudWatch.
	//
	// Lambda only sends application logs at the selected level of detail and lower, where `TRACE` is the highest level and `FATAL` is the lowest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-loggingconfig.html#cfn-lambda-function-loggingconfig-applicationloglevel
	//
	ApplicationLogLevel *string `field:"optional" json:"applicationLogLevel" yaml:"applicationLogLevel"`
	// The format in which Lambda sends your function's application and system logs to CloudWatch.
	//
	// Select between plain text and structured JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-loggingconfig.html#cfn-lambda-function-loggingconfig-logformat
	//
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The name of the Amazon CloudWatch log group the function sends logs to.
	//
	// By default, Lambda functions send logs to a default log group named `/aws/lambda/<function name>` . To use a different log group, enter an existing log group or enter a new log group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-loggingconfig.html#cfn-lambda-function-loggingconfig-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Set this property to filter the system logs for your function that Lambda sends to CloudWatch.
	//
	// Lambda only sends system logs at the selected level of detail and lower, where `DEBUG` is the highest level and `WARN` is the lowest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-loggingconfig.html#cfn-lambda-function-loggingconfig-systemloglevel
	//
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

