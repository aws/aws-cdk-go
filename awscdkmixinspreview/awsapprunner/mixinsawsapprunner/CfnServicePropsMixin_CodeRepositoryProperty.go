package mixinsawsapprunner


// Describes a source code repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeRepositoryProperty := &CodeRepositoryProperty{
//   	CodeConfiguration: &CodeConfigurationProperty{
//   		CodeConfigurationValues: &CodeConfigurationValuesProperty{
//   			BuildCommand: jsii.String("buildCommand"),
//   			Port: jsii.String("port"),
//   			Runtime: jsii.String("runtime"),
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
//   		ConfigurationSource: jsii.String("configurationSource"),
//   	},
//   	RepositoryUrl: jsii.String("repositoryUrl"),
//   	SourceCodeVersion: &SourceCodeVersionProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.String("value"),
//   	},
//   	SourceDirectory: jsii.String("sourceDirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-coderepository.html
//
type CfnServicePropsMixin_CodeRepositoryProperty struct {
	// Configuration for building and running the service from a source code repository.
	//
	// > `CodeConfiguration` is required only for `CreateService` request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-coderepository.html#cfn-apprunner-service-coderepository-codeconfiguration
	//
	CodeConfiguration interface{} `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// The location of the repository that contains the source code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-coderepository.html#cfn-apprunner-service-coderepository-repositoryurl
	//
	RepositoryUrl *string `field:"optional" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-coderepository.html#cfn-apprunner-service-coderepository-sourcecodeversion
	//
	SourceCodeVersion interface{} `field:"optional" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
	// The path of the directory that stores source code and configuration files.
	//
	// The build and start commands also execute from here. The path is absolute from root and, if not specified, defaults to the repository root.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-coderepository.html#cfn-apprunner-service-coderepository-sourcedirectory
	//
	SourceDirectory *string `field:"optional" json:"sourceDirectory" yaml:"sourceDirectory"`
}

