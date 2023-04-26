// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha


// The source of the App Runner configuration.
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
type ConfigurationSourceType string

const (
	// App Runner reads configuration values from `the apprunner.yaml` file in the source code repository and ignores `configurationValues`.
	// Experimental.
	ConfigurationSourceType_REPOSITORY ConfigurationSourceType = "REPOSITORY"
	// App Runner uses configuration values provided in `configurationValues` and ignores the `apprunner.yaml` file in the source code repository.
	// Experimental.
	ConfigurationSourceType_API ConfigurationSourceType = "API"
)

