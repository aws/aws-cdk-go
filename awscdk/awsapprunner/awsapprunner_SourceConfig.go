package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
)

// Result of binding `Source` into a `Service`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gitHubConnection gitHubConnection
//   var repository repository
//   var runtime runtime
//
//   sourceConfig := &SourceConfig{
//   	CodeRepository: &CodeRepositoryProps{
//   		CodeConfiguration: &CodeConfiguration{
//   			ConfigurationSource: awscdk.Aws_apprunner.ConfigurationSourceType_REPOSITORY,
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
//   		ImageRepositoryType: awscdk.*Aws_apprunner.ImageRepositoryType_ECR_PUBLIC,
//
//   		// the properties below are optional
//   		ImageConfiguration: &ImageConfiguration{
//   			Environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
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

