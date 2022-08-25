package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Artifacts definition for a CodeBuild Project.
//
// Example:
//   var bucket bucket
//
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//   	artifacts: codebuild.artifacts.s3(&s3ArtifactsProps{
//   		bucket: bucket,
//   		includeBuildId: jsii.Boolean(false),
//   		packageZip: jsii.Boolean(true),
//   		path: jsii.String("another/path"),
//   		identifier: jsii.String("AddArtifact1"),
//   	}),
//   })
//
// Experimental.
type Artifacts interface {
	IArtifacts
	// The artifact identifier.
	//
	// This property is required on secondary artifacts.
	// Experimental.
	Identifier() *string
	// The CodeBuild type of this artifact.
	// Experimental.
	Type() *string
	// Callback when an Artifacts class is used in a CodeBuild Project.
	// Experimental.
	Bind(_scope awscdk.Construct, _project IProject) *ArtifactsConfig
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


// Experimental.
func NewArtifacts_Override(a Artifacts, props *ArtifactsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codebuild.Artifacts",
		[]interface{}{props},
		a,
	)
}

// Experimental.
func Artifacts_S3(props *S3ArtifactsProps) IArtifacts {
	_init_.Initialize()

	var returns IArtifacts

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.Artifacts",
		"s3",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Artifacts) Bind(_scope awscdk.Construct, _project IProject) *ArtifactsConfig {
	var returns *ArtifactsConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _project},
		&returns,
	)

	return returns
}

