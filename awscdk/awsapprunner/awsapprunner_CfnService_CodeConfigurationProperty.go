package awsapprunner


// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfigurationProperty := &codeConfigurationProperty{
//   	configurationSource: jsii.String("configurationSource"),
//
//   	// the properties below are optional
//   	codeConfigurationValues: &codeConfigurationValuesProperty{
//   		runtime: jsii.String("runtime"),
//
//   		// the properties below are optional
//   		buildCommand: jsii.String("buildCommand"),
//   		port: jsii.String("port"),
//   		runtimeEnvironmentVariables: []interface{}{
//   			&keyValuePairProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		startCommand: jsii.String("startCommand"),
//   	},
//   }
//
type CfnService_CodeConfigurationProperty struct {
	// The source of the App Runner configuration. Values are interpreted as follows:.
	//
	// - `REPOSITORY` – App Runner reads configuration values from the `apprunner.yaml` file in the source code repository and ignores `CodeConfigurationValues` .
	// - `API` – App Runner uses configuration values provided in `CodeConfigurationValues` and ignores the `apprunner.yaml` file in the source code repository.
	ConfigurationSource *string `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a `apprunner.yaml` file in the source code repository (or ignoring the file if it exists).
	CodeConfigurationValues interface{} `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

