package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Translate stack outputs to Codepipline variable references.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipelineBase pipelineBase
//
//   stackOutputsMap := awscdk.Pipelines.NewStackOutputsMap(pipelineBase)
//
type StackOutputsMap interface {
	// Return the matching variable reference string for a StackOutputReference.
	ToCodePipeline(x StackOutputReference) *string
}

// The jsii proxy struct for StackOutputsMap
type jsiiProxy_StackOutputsMap struct {
	_ byte // padding
}

func NewStackOutputsMap(pipeline PipelineBase) StackOutputsMap {
	_init_.Initialize()

	if err := validateNewStackOutputsMapParameters(pipeline); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackOutputsMap{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.StackOutputsMap",
		[]interface{}{pipeline},
		&j,
	)

	return &j
}

func NewStackOutputsMap_Override(s StackOutputsMap, pipeline PipelineBase) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.StackOutputsMap",
		[]interface{}{pipeline},
		s,
	)
}

func (s *jsiiProxy_StackOutputsMap) ToCodePipeline(x StackOutputReference) *string {
	if err := s.validateToCodePipelineParameters(x); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"toCodePipeline",
		[]interface{}{x},
		&returns,
	)

	return returns
}

