// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha


// Properties of the Github repository for `Source.fromGitHub()`.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromGitHub(&GithubRepositoryProps{
//   		RepositoryUrl: jsii.String("https://github.com/aws-containers/hello-app-runner"),
//   		Branch: jsii.String("main"),
//   		ConfigurationSource: apprunner.ConfigurationSourceType_REPOSITORY,
//   		Connection: apprunner.GitHubConnection_FromConnectionArn(jsii.String("CONNECTION_ARN")),
//   	}),
//   })
//
// Experimental.
type GithubRepositoryProps struct {
	// The source of the App Runner configuration.
	// Experimental.
	ConfigurationSource ConfigurationSourceType `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// ARN of the connection to Github.
	//
	// Only required for Github source.
	// Experimental.
	Connection GitHubConnection `field:"required" json:"connection" yaml:"connection"`
	// The location of the repository that contains the source code.
	// Experimental.
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
	// The branch name that represents a specific version for the repository.
	// Experimental.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The code configuration values.
	//
	// Will be ignored if configurationSource is `REPOSITORY`.
	// Experimental.
	CodeConfigurationValues *CodeConfigurationValues `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

