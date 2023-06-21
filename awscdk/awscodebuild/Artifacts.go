package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Artifacts definition for a CodeBuild Project.
//
// Example:
//   var bucket bucket
//
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//   	Artifacts: codebuild.Artifacts_S3(&S3ArtifactsProps{
//   		Bucket: *Bucket,
//   		IncludeBuildId: jsii.Boolean(false),
//   		PackageZip: jsii.Boolean(true),
//   		Path: jsii.String("another/path"),
//   		Identifier: jsii.String("AddArtifact1"),
//   	}),
//   })
//
type Artifacts interface {
	IArtifacts
	// The artifact identifier.
	//
	// This property is required on secondary artifacts.
	Identifier() *string
	// The CodeBuild type of this artifact.
	Type() *string
	// Callback when an Artifacts class is used in a CodeBuild Project.
	Bind(_scope constructs.Construct, _project IProject) *ArtifactsConfig
}

// The jsii proxy struct for Artifacts
type jsiiProxy_Artifacts struct {
	jsiiProxy_IArtifacts
}

func (j *jsiiProxy_Artifacts) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifacts) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewArtifacts_Override(a Artifacts, props *ArtifactsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Artifacts",
		[]interface{}{props},
		a,
	)
}

func Artifacts_S3(props *S3ArtifactsProps) IArtifacts {
	_init_.Initialize()

	if err := validateArtifacts_S3Parameters(props); err != nil {
		panic(err)
	}
	var returns IArtifacts

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Artifacts",
		"s3",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Artifacts) Bind(_scope constructs.Construct, _project IProject) *ArtifactsConfig {
	if err := a.validateBindParameters(_scope, _project); err != nil {
		panic(err)
	}
	var returns *ArtifactsConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _project},
		&returns,
	)

	return returns
}

