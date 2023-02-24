// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the service source from ECR Public.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   var secret secret
//
//   ecrPublicSource := apprunner_alpha.NewEcrPublicSource(&EcrPublicProps{
//   	ImageIdentifier: jsii.String("imageIdentifier"),
//
//   	// the properties below are optional
//   	ImageConfiguration: &ImageConfiguration{
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		EnvironmentSecrets: map[string]*secret{
//   			"environmentSecretsKey": secret,
//   		},
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		Port: jsii.Number(123),
//   		StartCommand: jsii.String("startCommand"),
//   	},
//   })
//
// Experimental.
type EcrPublicSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for EcrPublicSource
type jsiiProxy_EcrPublicSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewEcrPublicSource(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	if err := validateNewEcrPublicSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrPublicSource{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrPublicSource_Override(e EcrPublicSource, props *EcrPublicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		[]interface{}{props},
		e,
	)
}

// Source from local assets.
// Experimental.
func EcrPublicSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	if err := validateEcrPublicSource_FromAssetParameters(props); err != nil {
		panic(err)
	}
	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func EcrPublicSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	if err := validateEcrPublicSource_FromEcrParameters(props); err != nil {
		panic(err)
	}
	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func EcrPublicSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	if err := validateEcrPublicSource_FromEcrPublicParameters(props); err != nil {
		panic(err)
	}
	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func EcrPublicSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	if err := validateEcrPublicSource_FromGitHubParameters(props); err != nil {
		panic(err)
	}
	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrPublicSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPublicSource) Bind(_scope constructs.Construct) *SourceConfig {
	if err := e.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

