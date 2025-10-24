package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// A FileSet created from a CodePipeline artifact.
//
// You only need to use this if you want to add CDK Pipeline stages
// add the end of an existing CodePipeline, which should be very rare.
//
// Example:
//   var codePipeline Pipeline
//
//
//   sourceArtifact := codepipeline.NewArtifact(jsii.String("MySourceArtifact"))
//
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
//   	CodePipeline: codePipeline,
//   	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
//   		Input: pipelines.CodePipelineFileSet_FromArtifact(sourceArtifact),
//   		Commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//   })
//
type CodePipelineFileSet interface {
	FileSet
	// Human-readable descriptor for this file set (does not need to be unique).
	Id() *string
	// The primary output of a file set producer.
	//
	// The primary output of a FileSet is itself.
	PrimaryOutput() FileSet
	// The Step that produces this FileSet.
	Producer() Step
	// Mark the given Step as the producer for this FileSet.
	//
	// This method can only be called once.
	ProducedBy(producer Step)
	// Return a string representation of this FileSet.
	ToString() *string
}

// The jsii proxy struct for CodePipelineFileSet
type jsiiProxy_CodePipelineFileSet struct {
	jsiiProxy_FileSet
}

func (j *jsiiProxy_CodePipelineFileSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineFileSet) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineFileSet) Producer() Step {
	var returns Step
	_jsii_.Get(
		j,
		"producer",
		&returns,
	)
	return returns
}


// Turn a CodePipeline Artifact into a FileSet.
func CodePipelineFileSet_FromArtifact(artifact awscodepipeline.Artifact) CodePipelineFileSet {
	_init_.Initialize()

	if err := validateCodePipelineFileSet_FromArtifactParameters(artifact); err != nil {
		panic(err)
	}
	var returns CodePipelineFileSet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipelineFileSet",
		"fromArtifact",
		[]interface{}{artifact},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineFileSet) ProducedBy(producer Step) {
	_jsii_.InvokeVoid(
		c,
		"producedBy",
		[]interface{}{producer},
	)
}

func (c *jsiiProxy_CodePipelineFileSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

