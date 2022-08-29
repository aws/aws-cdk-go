package pipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Factory for explicit CodePipeline Actions.
//
// If you have specific types of Actions you want to add to a
// CodePipeline, write a subclass of `Step` that implements this
// interface, and add the action or actions you want in the `produce` method.
//
// There needs to be a level of indirection here, because some aspects of the
// Action creation need to be controlled by the workflow engine (name and
// runOrder). All the rest of the properties are controlled by the factory.
type ICodePipelineActionFactory interface {
	// Create the desired Action and add it to the pipeline.
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
}

// The jsii proxy for ICodePipelineActionFactory
type jsiiProxy_ICodePipelineActionFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_ICodePipelineActionFactory) ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult {
	var returns *CodePipelineActionFactoryResult

	_jsii_.Invoke(
		i,
		"produceAction",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

