package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderTelemetryConfigProperty := &CapacityProviderTelemetryConfigProperty{
//   	LoggingConfig: &CapacityProviderLoggingConfigProperty{
//   		LogGroup: jsii.String("logGroup"),
//   		SystemLogLevel: jsii.String("systemLogLevel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityprovidertelemetryconfig.html
//
type CfnCapacityProvider_CapacityProviderTelemetryConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityprovidertelemetryconfig.html#cfn-lambda-capacityprovider-capacityprovidertelemetryconfig-loggingconfig
	//
	LoggingConfig interface{} `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}

