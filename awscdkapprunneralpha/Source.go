package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the App Runner service source.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   service := apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   })
//
//   service.AddToRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Effect: iam.Effect_ALLOW,
//   	Actions: []*string{
//   		jsii.String("s3:GetObject"),
//   	},
//   	Resources: []*string{
//   		jsii.String("*"),
//   	},
//   }))
//
// Experimental.
type Source interface {
	// Called when the Job is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct) *SourceConfig
}

// The jsii proxy struct for Source
type jsiiProxy_Source struct {
	_ byte // padding
}

// Experimental.
func NewSource_Override(s Source) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.Source",
		nil, // no parameters
		s,
	)
}

// Source from local assets.
// Experimental.
func Source_FromAsset(props *AssetProps) AssetSource {
	_init_.Initialize()

	if err := validateSource_FromAssetParameters(props); err != nil {
		panic(err)
	}
	var returns AssetSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Source",
		"fromAsset",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR repository.
// Experimental.
func Source_FromEcr(props *EcrProps) EcrSource {
	_init_.Initialize()

	if err := validateSource_FromEcrParameters(props); err != nil {
		panic(err)
	}
	var returns EcrSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Source",
		"fromEcr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the ECR Public repository.
// Experimental.
func Source_FromEcrPublic(props *EcrPublicProps) EcrPublicSource {
	_init_.Initialize()

	if err := validateSource_FromEcrPublicParameters(props); err != nil {
		panic(err)
	}
	var returns EcrPublicSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Source",
		"fromEcrPublic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Source from the GitHub repository.
// Experimental.
func Source_FromGitHub(props *GithubRepositoryProps) GithubSource {
	_init_.Initialize()

	if err := validateSource_FromGitHubParameters(props); err != nil {
		panic(err)
	}
	var returns GithubSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Source",
		"fromGitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Source) Bind(scope constructs.Construct) *SourceConfig {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

