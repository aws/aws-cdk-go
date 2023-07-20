package awsapprunner


// Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeConfigurationProperty := &CodeConfigurationProperty{
//   	ConfigurationSource: jsii.String("configurationSource"),
//
//   	// the properties below are optional
//   	CodeConfigurationValues: &CodeConfigurationValuesProperty{
//   		Runtime: jsii.String("runtime"),
//
//   		// the properties below are optional
//   		BuildCommand: jsii.String("buildCommand"),
//   		Port: jsii.String("port"),
//   		RuntimeEnvironmentSecrets: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RuntimeEnvironmentVariables: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		StartCommand: jsii.String("startCommand"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfiguration.html
//
type CfnService_CodeConfigurationProperty struct {
	// The source of the App Runner configuration. Values are interpreted as follows:.
	//
	// - `REPOSITORY` – App Runner reads configuration values from the `apprunner.yaml` file in the source code repository and ignores `CodeConfigurationValues` .
	// - `API` – App Runner uses configuration values provided in `CodeConfigurationValues` and ignores the `apprunner.yaml` file in the source code repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfiguration.html#cfn-apprunner-service-codeconfiguration-configurationsource
	//
	ConfigurationSource *string `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// The basic configuration for building and running the App Runner service.
	//
	// Use it to quickly launch an App Runner service without providing a `apprunner.yaml` file in the source code repository (or ignoring the file if it exists).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-codeconfiguration.html#cfn-apprunner-service-codeconfiguration-codeconfigurationvalues
	//
	CodeConfigurationValues interface{} `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

