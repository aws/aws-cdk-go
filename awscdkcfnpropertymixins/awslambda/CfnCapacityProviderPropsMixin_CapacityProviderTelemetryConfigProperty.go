package awslambda


// Configuration that specifies the telemetry collection for the capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnCapacityProviderPropsMixin_CapacityProviderTelemetryConfigProperty struct {
	// The capacity provider's Amazon CloudWatch Logs configuration settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityprovidertelemetryconfig.html#cfn-lambda-capacityprovider-capacityprovidertelemetryconfig-loggingconfig
	//
	LoggingConfig interface{} `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}

