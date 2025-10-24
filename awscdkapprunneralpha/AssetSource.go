package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the source from local assets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImageAsset DockerImageAsset
//   var secret Secret
//
//   assetSource := apprunner_alpha.NewAssetSource(&AssetProps{
//   	Asset: dockerImageAsset,
//
//   	// the properties below are optional
//   	ImageConfiguration: &ImageConfiguration{
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		EnvironmentSecrets: map[string]Secret{
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
type AssetSource interface {
	Source
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(_scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for AssetSource
type jsiiProxy_AssetSource struct {
	jsiiProxy_Source
}

// Experimental.
func NewAssetSource(props *AssetProps) AssetSource {
	_init_.Initialize()

	if err := validateNewAssetSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetSource{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetSource_Override(a AssetSource, props *AssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		[]interface{}{props},
		a,
	)
}

// Source from local assets.
// Experimental.
func AssetSource_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	if err := validateAssetSource_FromAssetParameters(props); err != nil {
		panic(err)
	}
	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func AssetSource_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	if err := validateAssetSource_FromEcrParameters(props); err != nil {
		panic(err)
	}
	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func AssetSource_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	if err := validateAssetSource_FromEcrPublicParameters(props); err != nil {
		panic(err)
	}
	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func AssetSource_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	if err := validateAssetSource_FromGitHubParameters(props); err != nil {
		panic(err)
	}
	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.AssetSource",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetSource) Bind(_scope constructs.Construct) *SourceConfig {
	if err := a.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

