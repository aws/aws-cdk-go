package awsgreengrass


// A logger represents logging settings for the AWS IoT Greengrass group, which can be stored in CloudWatch and the local file system of your core device.
//
// All log entries include a timestamp, log level, and information about the event. For more information, see [Monitoring with AWS IoT Greengrass Logs](https://docs.aws.amazon.com/greengrass/latest/developerguide/greengrass-logs-overview.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Loggers` property of the [`AWS::Greengrass::LoggerDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinitionversion.html) resource contains a list of `Logger` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggerProperty := &loggerProperty{
//   	component: jsii.String("component"),
//   	id: jsii.String("id"),
//   	level: jsii.String("level"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	space: jsii.Number(123),
//   }
//
type CfnLoggerDefinitionVersion_LoggerProperty struct {
	// The source of the log event.
	//
	// Valid values are `GreengrassSystem` or `Lambda` . When `GreengrassSystem` is used, events from Greengrass system components are logged. When `Lambda` is used, events from user-defined Lambda functions are logged.
	Component *string `field:"required" json:"component" yaml:"component"`
	// A descriptive or arbitrary ID for the logger.
	//
	// This value must be unique within the logger definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
	// The log-level threshold.
	//
	// Log events below this threshold are filtered out and aren't stored. Valid values are `DEBUG` , `INFO` (recommended), `WARN` , `ERROR` , or `FATAL` .
	Level *string `field:"required" json:"level" yaml:"level"`
	// The storage mechanism for log events.
	//
	// Valid values are `FileSystem` or `AWSCloudWatch` . When `AWSCloudWatch` is used, log events are sent to CloudWatch Logs . When `FileSystem` is used, log events are stored on the local file system.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The amount of file space (in KB) to use when writing logs to the local file system.
	//
	// This property does not apply for CloudWatch Logs .
	Space *float64 `field:"optional" json:"space" yaml:"space"`
}

