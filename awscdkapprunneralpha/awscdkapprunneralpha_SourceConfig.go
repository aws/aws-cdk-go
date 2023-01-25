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
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gitHubConnection gitHubConnection
//   var repository repository
//   var runtime runtime
//   var secret secret
//
//   sourceConfig := &sourceConfig{
//   	codeRepository: &codeRepositoryProps{
//   		codeConfiguration: &codeConfiguration{
//   			configurationSource: apprunner_alpha.configurationSourceType_REPOSITORY,
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
//   				environmentSecrets: map[string]*secret{
//   					"environmentSecretsKey": secret,
//   				},
//   				environmentVariables: map[string]*string{
//   					"environmentVariablesKey": jsii.String("environmentVariables"),
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
//   		imageRepositoryType: apprunner_alpha.imageRepositoryType_ECR_PUBLIC,
//
//   		// the properties below are optional
//   		imageConfiguration: &imageConfiguration{
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			environmentSecrets: map[string]*secret{
//   				"environmentSecretsKey": secret,
//   			},
//   			environmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
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

