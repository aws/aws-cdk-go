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
//   sourceConfig := &sourceConfig{
//   	codeRepository: &codeRepositoryProps{
//   		codeConfiguration: &codeConfiguration{
//   			configurationSource: awscdk.Aws_apprunner.configurationSourceType_REPOSITORY,
//
//   			// the properties below are optional
//   			configurationValues: &codeConfigurationValues{
//   				runtime: runtime,
//
//   				// the properties below are optional
//   				buildCommand: jsii.String("buildCommand"),
//   				environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				port: jsii.String("port"),
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   		connection: gitHubConnection,
//   		repositoryUrl: jsii.String("repositoryUrl"),
//   		sourceCodeVersion: &sourceCodeVersion{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrRepository: repository,
//   	imageRepository: &imageRepository{
//   		imageIdentifier: jsii.String("imageIdentifier"),
//   		imageRepositoryType: awscdk.*Aws_apprunner.imageRepositoryType_ECR_PUBLIC,
//
//   		// the properties below are optional
//   		imageConfiguration: &imageConfiguration{
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			port: jsii.Number(123),
//   			startCommand: jsii.String("startCommand"),
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

