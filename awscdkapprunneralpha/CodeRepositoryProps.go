package awscdkapprunneralpha


// Properties of the CodeRepository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var gitHubConnection gitHubConnection
//   var runtime runtime
//   var secret secret
//
//   codeRepositoryProps := &CodeRepositoryProps{
//   	CodeConfiguration: &CodeConfiguration{
//   		ConfigurationSource: apprunner_alpha.ConfigurationSourceType_REPOSITORY,
//
//   		// the properties below are optional
//   		ConfigurationValues: &CodeConfigurationValues{
//   			Runtime: runtime,
//
//   			// the properties below are optional
//   			BuildCommand: jsii.String("buildCommand"),
//   			Environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			EnvironmentSecrets: map[string]*secret{
//   				"environmentSecretsKey": secret,
//   			},
//   			EnvironmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
//   			},
//   			Port: jsii.String("port"),
//   			StartCommand: jsii.String("startCommand"),
//   		},
//   	},
//   	Connection: gitHubConnection,
//   	RepositoryUrl: jsii.String("repositoryUrl"),
//   	SourceCodeVersion: &SourceCodeVersion{
//   		Type: jsii.String("type"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// Experimental.
type CodeRepositoryProps struct {
	// Configuration for building and running the service from a source code repository.
	// Experimental.
	CodeConfiguration *CodeConfiguration `field:"required" json:"codeConfiguration" yaml:"codeConfiguration"`
	// The App Runner connection for GitHub.
	// Experimental.
	Connection GitHubConnection `field:"required" json:"connection" yaml:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The version that should be used within the source code repository.
	// Experimental.
	SourceCodeVersion *SourceCodeVersion `field:"required" json:"sourceCodeVersion" yaml:"sourceCodeVersion"`
}

