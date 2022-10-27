package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The abstract interface of a CodeBuild build output.
//
// Implemented by {@link Artifacts}.
type IArtifacts interface {
	// Callback when an Artifacts class is used in a CodeBuild Project.
	Bind(scope constructs.Construct, project IProject) *ArtifactsConfig
	// The artifact identifier.
	//
	// This property is required on secondary artifacts.
	Identifier() *string
	// The CodeBuild type of this artifact.
	Type() *string
}

// The jsii proxy for IArtifacts
type jsiiProxy_IArtifacts struct {
	_ byte // padding
}

func (i *jsiiProxy_IArtifacts) Bind(scope constructs.Construct, project IProject) *ArtifactsConfig {
	if err := i.validateBindParameters(scope, project); err != nil {
		panic(err)
	}
	var returns *ArtifactsConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IArtifacts) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IArtifacts) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

