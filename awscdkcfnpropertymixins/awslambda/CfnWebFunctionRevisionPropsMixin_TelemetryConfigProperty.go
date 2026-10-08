package awslambda


// The telemetry configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   telemetryConfigProperty := &TelemetryConfigProperty{
//   	LoggingConfig: &LoggingConfigProperty{
//   		ApplicationLogLevel: jsii.String("applicationLogLevel"),
//   		LogGroup: jsii.String("logGroup"),
//   		SystemLogLevel: jsii.String("systemLogLevel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-telemetryconfig.html
//
type CfnWebFunctionRevisionPropsMixin_TelemetryConfigProperty struct {
	// The logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-telemetryconfig.html#cfn-lambda-webfunctionrevision-telemetryconfig-loggingconfig
	//
	LoggingConfig interface{} `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}

