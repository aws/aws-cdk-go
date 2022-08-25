package awsapprunner


// Describes a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeRepositoryProperty := &codeRepositoryProperty{
//   	repositoryUrl: jsii.String("repositoryUrl"),
//   	sourceCodeVersion: &sourceCodeVersionProperty{
//   		type: jsii.String("type"),
//   		value: jsii.String("value"),
//   	},
//
//   	// the properties below are optional
//   	codeConfiguration: &codeConfigurationProperty{
//   		configurationSource: jsii.String("configurationSource"),
//
//   		// the properties below are optional
//   		codeConfigurationValues: &codeConfigurationValuesProperty{
//   			runtime: jsii.String("runtime"),
//
//   			// the properties below are optional
//   			buildCommand: jsii.String("buildCommand"),
//   			port: jsii.String("port"),
//   			runtimeEnvironmentVariables: []interface{}{
//   				&keyValuePairProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   }
//
type CfnService_CodeRepositoryProperty struct {
	// The location of the repository that contains the source code.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	SourceCodeVersion interface{} `field:"required" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// Configuration for building and running the service from a source code repository.
	//
	// > `CodeConfiguration` is required only for `CreateService` request.
	CodeConfiguration interface{} `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
}

