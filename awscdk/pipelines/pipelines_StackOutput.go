package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
)

// A single output of a Stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifactPath artifactPath
//
//   stackOutput := awscdk.Pipelines.NewStackOutput(artifactPath, jsii.String("outputName"))
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type StackOutput interface {
	// The artifact and file the output is stored in.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ArtifactFile() awscodepipeline.ArtifactPath
	// The name of the output in the JSON object in the file.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OutputName() *string
}

// The jsii proxy struct for StackOutput
type jsiiProxy_StackOutput struct {
	_ byte // padding
}

func (j *jsiiProxy_StackOutput) ArtifactFile() awscodepipeline.ArtifactPath {
	var returns awscodepipeline.ArtifactPath
	_jsii_.Get(
		j,
		"artifactFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackOutput) OutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputName",
		&returns,
	)
	return returns
}


// Build a StackOutput from a known artifact and an output name.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewStackOutput(artifactFile awscodepipeline.ArtifactPath, outputName *string) StackOutput {
	_init_.Initialize()

	if err := validateNewStackOutputParameters(artifactFile, outputName); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackOutput{}

	_jsii_.Create(
		"monocdk.pipelines.StackOutput",
		[]interface{}{artifactFile, outputName},
		&j,
	)

	return &j
}

// Build a StackOutput from a known artifact and an output name.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewStackOutput_Override(s StackOutput, artifactFile awscodepipeline.ArtifactPath, outputName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.StackOutput",
		[]interface{}{artifactFile, outputName},
		s,
	)
}

