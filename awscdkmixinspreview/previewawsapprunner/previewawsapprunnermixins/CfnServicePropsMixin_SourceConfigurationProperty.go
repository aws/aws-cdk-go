package previewawsapprunnermixins


// Describes the source deployed to an AWS App Runner service.
//
// It can be a code or an image repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceConfigurationProperty := &SourceConfigurationProperty{
//   	AuthenticationConfiguration: &AuthenticationConfigurationProperty{
//   		AccessRoleArn: jsii.String("accessRoleArn"),
//   		ConnectionArn: jsii.String("connectionArn"),
//   	},
//   	AutoDeploymentsEnabled: jsii.Boolean(false),
//   	CodeRepository: &CodeRepositoryProperty{
//   		CodeConfiguration: &CodeConfigurationProperty{
//   			CodeConfigurationValues: &CodeConfigurationValuesProperty{
//   				BuildCommand: jsii.String("buildCommand"),
//   				Port: jsii.String("port"),
//   				Runtime: jsii.String("runtime"),
//   				RuntimeEnvironmentSecrets: []interface{}{
//   					&KeyValuePairProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				RuntimeEnvironmentVariables: []interface{}{
//   					&KeyValuePairProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				StartCommand: jsii.String("startCommand"),
//   			},
//   			ConfigurationSource: jsii.String("configurationSource"),
//   		},
//   		RepositoryUrl: jsii.String("repositoryUrl"),
//   		SourceCodeVersion: &SourceCodeVersionProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   		SourceDirectory: jsii.String("sourceDirectory"),
//   	},
//   	ImageRepository: &ImageRepositoryProperty{
//   		ImageConfiguration: &ImageConfigurationProperty{
//   			Port: jsii.String("port"),
//   			RuntimeEnvironmentSecrets: []interface{}{
//   				&KeyValuePairProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			RuntimeEnvironmentVariables: []interface{}{
//   				&KeyValuePairProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			StartCommand: jsii.String("startCommand"),
//   		},
//   		ImageIdentifier: jsii.String("imageIdentifier"),
//   		ImageRepositoryType: jsii.String("imageRepositoryType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourceconfiguration.html
//
type CfnServicePropsMixin_SourceConfigurationProperty struct {
	// Describes the resources that are needed to authenticate access to some source repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourceconfiguration.html#cfn-apprunner-service-sourceconfiguration-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// If `true` , continuous integration from the source repository is enabled for the App Runner service.
	//
	// Each repository change (including any source code commit or new image version) starts a deployment.
	//
	// Default: App Runner sets to `false` for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in. App Runner sets to `true` in all other cases (which currently include a source code repository or a source image using a same-account ECR repository).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourceconfiguration.html#cfn-apprunner-service-sourceconfiguration-autodeploymentsenabled
	//
	AutoDeploymentsEnabled interface{} `field:"optional" json:"autoDeploymentsEnabled" yaml:"autoDeploymentsEnabled"`
	// The description of a source code repository.
	//
	// You must provide either this member or `ImageRepository` (but not both).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourceconfiguration.html#cfn-apprunner-service-sourceconfiguration-coderepository
	//
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// The description of a source image repository.
	//
	// You must provide either this member or `CodeRepository` (but not both).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-sourceconfiguration.html#cfn-apprunner-service-sourceconfiguration-imagerepository
	//
	ImageRepository interface{} `field:"optional" json:"imageRepository" yaml:"imageRepository"`
}

