package awslambda


// The service configuration for the revision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConfigProperty := &ServiceConfigProperty{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//
//   	// the properties below are optional
//   	EgressNetworkConnectorArn: jsii.String("egressNetworkConnectorArn"),
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	MaxConcurrencyPerEnvironment: jsii.Number(123),
//   	TelemetryConfig: &TelemetryConfigProperty{
//   		LoggingConfig: &LoggingConfigProperty{
//   			ApplicationLogLevel: jsii.String("applicationLogLevel"),
//   			LogGroup: jsii.String("logGroup"),
//   			SystemLogLevel: jsii.String("systemLogLevel"),
//   		},
//   	},
//   	TimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-serviceconfig.html
//
type CfnWebFunctionRevision_ServiceConfigProperty struct {
	// The ARN of the execution role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-serviceconfig.html#cfn-lambda-webfunctionrevision-serviceconfig-executionrolearn
	//
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The versioned ARN of the network connector used for outbound network access from the revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-serviceconfig.html#cfn-lambda-webfunctionrevision-serviceconfig-egressnetworkconnectorarn
	//
	EgressNetworkConnectorArn *string `field:"optional" json:"egressNetworkConnectorArn" yaml:"egressNetworkConnectorArn"`
	// Environment variables for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-serviceconfig.html#cfn-lambda-webfunctionrevision-serviceconfig-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The maximum concurrency per environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-serviceconfig.html#cfn-lambda-webfunctionrevision-serviceconfig-maxconcurrencyperenvironment
	//
	MaxConcurrencyPerEnvironment *float64 `field:"optional" json:"maxConcurrencyPerEnvironment" yaml:"maxConcurrencyPerEnvironment"`
	// The telemetry configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-serviceconfig.html#cfn-lambda-webfunctionrevision-serviceconfig-telemetryconfig
	//
	TelemetryConfig interface{} `field:"optional" json:"telemetryConfig" yaml:"telemetryConfig"`
	// The function timeout in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionrevision-serviceconfig.html#cfn-lambda-webfunctionrevision-serviceconfig-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

