package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A Reference to a Stack Output.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   type myLambdaStep struct {
//   	step
//   	stackOutputReference stackOutputReferencefunction lambda.Function
//   	stackOutput cfnOutput
//   }
//
//   func newMyLambdaStep( ) *myLambdaStep {
//   	this := &myLambdaStep{}
//   	return this
//   }{
//   	newStack_Override(this, jsii.String("MyLambdaStep"))
//   	this.stackOutputReference = pipelines.stackOutputReference_FromCfnOutput(stackOutput)
//   }produceAction(stage, codepipeline.iStage, options, pipelines.produceActionOptions)pipelines.codePipelineActionFactoryResult{
//
//   	stage.addAction(cpactions.NewLambdaInvokeAction(&LambdaInvokeActionProps{
//   		ActionName: options.actionName,
//   		RunOrder: options.runOrder,
//   		// Map the reference to the variable name the CDK has generated for you.
//   		UserParameters: map[string]interface{}{
//   			"stackOutput": options.stackOutputsMap.toCodePipeline(this.stackOutputReference),
//   		},
//   		Lambda: this.function,
//   	}))
//
//   	return map[string]*f64{
//   		"runOrdersConsumed": jsii.Number(1),
//   	}
//   }getconsumedStackOutputs()pipelines.stackOutputReference[]{
//   	return []interface{}{
//   		this.stackOutputReference,
//   	}
//   }
//
type StackOutputReference interface {
	// Output name of the producing stack.
	OutputName() *string
	// A human-readable description of the producing stack.
	StackDescription() *string
	// Whether or not this stack output is being produced by the given Stack deployment.
	IsProducedBy(stack StackDeployment) *bool
}

// The jsii proxy struct for StackOutputReference
type jsiiProxy_StackOutputReference struct {
	_ byte // padding
}

func (j *jsiiProxy_StackOutputReference) OutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackOutputReference) StackDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackDescription",
		&returns,
	)
	return returns
}


// Create a StackOutputReference that references the given CfnOutput.
func StackOutputReference_FromCfnOutput(output awscdk.CfnOutput) StackOutputReference {
	_init_.Initialize()

	if err := validateStackOutputReference_FromCfnOutputParameters(output); err != nil {
		panic(err)
	}
	var returns StackOutputReference

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.StackOutputReference",
		"fromCfnOutput",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackOutputReference) IsProducedBy(stack StackDeployment) *bool {
	if err := s.validateIsProducedByParameters(stack); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		s,
		"isProducedBy",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

