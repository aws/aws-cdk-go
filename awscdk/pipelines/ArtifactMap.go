package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Translate FileSets to CodePipeline Artifacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactMap := awscdk.Pipelines.NewArtifactMap()
//
type ArtifactMap interface {
	// Return the matching CodePipeline artifact for a FileSet.
	ToCodePipeline(x FileSet) awscodepipeline.Artifact
}

// The jsii proxy struct for ArtifactMap
type jsiiProxy_ArtifactMap struct {
	_ byte // padding
}

func NewArtifactMap() ArtifactMap {
	_init_.Initialize()

	j := jsiiProxy_ArtifactMap{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ArtifactMap",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewArtifactMap_Override(a ArtifactMap) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ArtifactMap",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_ArtifactMap) ToCodePipeline(x FileSet) awscodepipeline.Artifact {
	if err := a.validateToCodePipelineParameters(x); err != nil {
		panic(err)
	}
	var returns awscodepipeline.Artifact

	_jsii_.Invoke(
		a,
		"toCodePipeline",
		[]interface{}{x},
		&returns,
	)

	return returns
}

