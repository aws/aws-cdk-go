package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the service source from ECR.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var repository repository
//   var secret secret
//
//   ecrSource := apprunner_alpha.NewEcrSource(&EcrProps{
//   	Repository: repository,
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
//   	Tag: jsii.String("tag"),
//   	TagOrDigest: jsii.String("tagOrDigest"),
//   })
//
// Experimental.
type EcrSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for EcrSource
type jsiiProxy_EcrSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewEcrSource(props *EcrProps) EcrSource {
	_init_.Initialize()

	if err := validateNewEcrSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrSource{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcrSource_Override(e EcrSource, props *EcrProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		[]interface{}{props},
		e,
	)
}

// Source from local assets.
// Experimental.
func EcrSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	if err := validateEcrSource_FromAssetParameters(props); err != nil {
		panic(err)
	}
	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func EcrSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	if err := validateEcrSource_FromEcrParameters(props); err != nil {
		panic(err)
	}
	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func EcrSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	if err := validateEcrSource_FromEcrPublicParameters(props); err != nil {
		panic(err)
	}
	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func EcrSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	if err := validateEcrSource_FromGitHubParameters(props); err != nil {
		panic(err)
	}
	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.EcrSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrSource) Bind(_scope constructs.Construct) *SourceConfig {
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

