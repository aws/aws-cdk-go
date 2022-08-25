package awsapprunner


// Describes the source deployed to an AWS App Runner service.
//
// It can be a code or an image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceConfigurationProperty := &sourceConfigurationProperty{
//   	authenticationConfiguration: &authenticationConfigurationProperty{
//   		accessRoleArn: jsii.String("accessRoleArn"),
//   		connectionArn: jsii.String("connectionArn"),
//   	},
//   	autoDeploymentsEnabled: jsii.Boolean(false),
//   	codeRepository: &codeRepositoryProperty{
//   		repositoryUrl: jsii.String("repositoryUrl"),
//   		sourceCodeVersion: &sourceCodeVersionProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//
//   		// the properties below are optional
//   		codeConfiguration: &codeConfigurationProperty{
//   			configurationSource: jsii.String("configurationSource"),
//
//   			// the properties below are optional
//   			codeConfigurationValues: &codeConfigurationValuesProperty{
//   				runtime: jsii.String("runtime"),
//
//   				// the properties below are optional
//   				buildCommand: jsii.String("buildCommand"),
//   				port: jsii.String("port"),
//   				runtimeEnvironmentVariables: []interface{}{
//   					&keyValuePairProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				startCommand: jsii.String("startCommand"),
//   			},
//   		},
//   	},
//   	imageRepository: &imageRepositoryProperty{
//   		imageIdentifier: jsii.String("imageIdentifier"),
//   		imageRepositoryType: jsii.String("imageRepositoryType"),
//
//   		// the properties below are optional
//   		imageConfiguration: &imageConfigurationProperty{
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
type CfnService_SourceConfigurationProperty struct {
	// Describes the resources that are needed to authenticate access to some source repositories.
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// If `true` , continuous integration from the source repository is enabled for the App Runner service.
	//
	// Each repository change (including any source code commit or new image version) starts a deployment.
	//
	// Default: App Runner sets to `false` for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in. App Runner sets to `true` in all other cases (which currently include a source code repository or a source image using a same-account ECR repository).
	AutoDeploymentsEnabled interface{} `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// The description of a source code repository.
	//
	// You must provide either this member or `ImageRepository` (but not both).
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// The description of a source image repository.
	//
	// You must provide either this member or `CodeRepository` (but not both).
	ImageRepository interface{} `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

