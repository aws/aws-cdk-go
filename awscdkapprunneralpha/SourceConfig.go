// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
)

// Result of binding `Source` into a `Service`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gitHubConnection gitHubConnection
//   var repository repository
//   var runtime runtime
//   var secret secret
//
//   sourceConfig := &SourceConfig{
//   	CodeRepository: &CodeRepositoryProps{
//   		CodeConfiguration: &CodeConfiguration{
//   			ConfigurationSource: apprunner_alpha.ConfigurationSourceType_REPOSITORY,
//
//   			// the properties below are optional
//   			ConfigurationValues: &CodeConfigurationValues{
//   				Runtime: runtime,
//
//   				// the properties below are optional
//   				BuildCommand: jsii.String("buildCommand"),
//   				Environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				EnvironmentSecrets: map[string]*secret{
//   					"environmentSecretsKey": secret,
//   				},
//   				EnvironmentVariables: map[string]*string{
//   					"environmentVariablesKey": jsii.String("environmentVariables"),
//   				},
//   				Port: jsii.String("port"),
//   				StartCommand: jsii.String("startCommand"),
//   			},
//   		},
//   		Connection: gitHubConnection,
//   		RepositoryUrl: jsii.String("repositoryUrl"),
//   		SourceCodeVersion: &SourceCodeVersion{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	EcrRepository: repository,
//   	ImageRepository: &ImageRepository{
//   		ImageIdentifier: jsii.String("imageIdentifier"),
//   		ImageRepositoryType: apprunner_alpha.ImageRepositoryType_ECR_PUBLIC,
//
//   		// the properties below are optional
//   		ImageConfiguration: &ImageConfiguration{
//   			Environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			EnvironmentSecrets: map[string]*secret{
//   				"environmentSecretsKey": secret,
//   			},
//   			EnvironmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
//   			},
//   			Port: jsii.Number(123),
//   			StartCommand: jsii.String("startCommand"),
//   		},
//   	},
//   }
//
// Experimental.
type SourceConfig struct {
	// The code repository configuration (mutually exclusive  with `imageRepository`).
	// Experimental.
	CodeRepository *CodeRepositoryProps `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// The ECR repository (required to grant the pull privileges for the iam role).
	// Experimental.
	EcrRepository awsecr.IRepository `field:"optional" json:"ecrRepository" yaml:"ecrRepository"`
	// The image repository configuration (mutually exclusive  with `codeRepository`).
	// Experimental.
	ImageRepository *ImageRepository `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

