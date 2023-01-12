package awsapprunner


// Describes the basic configuration needed for building and running an AWS App Runner service.
//
// This type doesn't support the full set of possible configuration options. Fur full configuration capabilities, use a `apprunner.yaml` file in the source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfigurationValuesProperty := &codeConfigurationValuesProperty{
//   	runtime: jsii.String("runtime"),
//
//   	// the properties below are optional
//   	buildCommand: jsii.String("buildCommand"),
//   	port: jsii.String("port"),
//   	runtimeEnvironmentVariables: []interface{}{
//   		&keyValuePairProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	startCommand: jsii.String("startCommand"),
//   }
//
type CfnService_CodeConfigurationValuesProperty struct {
	// A runtime environment type for building and running an App Runner service.
	//
	// It represents a programming language runtime.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// The command App Runner runs to build your application.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The port that your application listens to in the container.
	//
	// Default: `8080`.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The environment variables that are available to your running App Runner service.
	//
	// An array of key-value pairs. Keys with a prefix of `AWSAPPRUNNER` are reserved for system use and aren't valid.
	RuntimeEnvironmentVariables interface{} `field:"optional" json:"runtimeEnvironmentVariables" yaml:"runtimeEnvironmentVariables"`
	// The command App Runner runs to start your application.
	StartCommand *string `field:"optional" json:"startCommand" yaml:"startCommand"`
}

