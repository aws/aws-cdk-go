package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the App Runner connection that enables the App Runner service to connect to a source repository.
//
// It's required for GitHub code repositories.
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
type GitHubConnection interface {
	// The ARN of the Connection for App Runner service to connect to the repository.
	// Experimental.
	ConnectionArn() *string
}

// The jsii proxy struct for GitHubConnection
type jsiiProxy_GitHubConnection struct {
	_ byte // padding
}

func (j *jsiiProxy_GitHubConnection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHubConnection(arn *string) GitHubConnection {
	_init_.Initialize()

	if err := validateNewGitHubConnectionParameters(arn); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitHubConnection{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.GitHubConnection",
		[]interface{}{arn},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubConnection_Override(g GitHubConnection, arn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.GitHubConnection",
		[]interface{}{arn},
		g,
	)
}

// Using existing App Runner connection by specifying the connection ARN.
//
// Returns: Connection.
// Experimental.
func GitHubConnection_FromConnectionArn(arn *string) GitHubConnection {
	_init_.Initialize()

	if err := validateGitHubConnection_FromConnectionArnParameters(arn); err != nil {
		panic(err)
	}
	var returns GitHubConnection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GitHubConnection",
		"fromConnectionArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

