// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the service source from a Github repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var gitHubConnection gitHubConnection
//   var runtime runtime
//   var secret secret
//
//   githubSource := apprunner_alpha.NewGithubSource(&GithubRepositoryProps{
//   	ConfigurationSource: apprunner_alpha.ConfigurationSourceType_REPOSITORY,
//   	Connection: gitHubConnection,
//   	RepositoryUrl: jsii.String("repositoryUrl"),
//
//   	// the properties below are optional
//   	Branch: jsii.String("branch"),
//   	CodeConfigurationValues: &CodeConfigurationValues{
//   		Runtime: runtime,
//
//   		// the properties below are optional
//   		BuildCommand: jsii.String("buildCommand"),
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		EnvironmentSecrets: map[string]*secret{
//   			"environmentSecretsKey": secret,
//   		},
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		Port: jsii.String("port"),
//   		StartCommand: jsii.String("startCommand"),
//   	},
//   })
//
// Experimental.
type GithubSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for GithubSource
type jsiiProxy_GithubSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewGithubSource(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	if err := validateNewGithubSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GithubSource{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGithubSource_Override(g GithubSource, props *GithubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		[]interface{}{props},
		g,
	)
}

// Source from local assets.
// Experimental.
func GithubSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	if err := validateGithubSource_FromAssetParameters(props); err != nil {
		panic(err)
	}
	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func GithubSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	if err := validateGithubSource_FromEcrParameters(props); err != nil {
		panic(err)
	}
	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func GithubSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	if err := validateGithubSource_FromEcrPublicParameters(props); err != nil {
		panic(err)
	}
	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func GithubSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	if err := validateGithubSource_FromGitHubParameters(props); err != nil {
		panic(err)
	}
	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.GithubSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GithubSource) Bind(_scope constructs.Construct) *SourceConfig {
	if err := g.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

