package awsapprunner


// Properties of the CodeRepository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gitHubConnection gitHubConnection
//   var runtime runtime
//
//   codeRepositoryProps := &codeRepositoryProps{
//   	codeConfiguration: &codeConfiguration{
//   		configurationSource: awscdk.Aws_apprunner.configurationSourceType_REPOSITORY,
//
//   		// the properties below are optional
//   		configurationValues: &codeConfigurationValues{
//   			runtime: runtime,
//
//   			// the properties below are optional
//   			buildCommand: jsii.String("buildCommand"),
//   			environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			port: jsii.String("port"),
//   			startCommand: jsii.String("startCommand"),
//   		},
//   	},
//   	connection: gitHubConnection,
//   	repositoryUrl: jsii.String("repositoryUrl"),
//   	sourceCodeVersion: &sourceCodeVersion{
//   		type: jsii.String("type"),
//   		value: jsii.String("value"),
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

