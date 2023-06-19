package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// A Reference to a Stack Output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnOutput cfnOutput
//
//   stackOutputReference := awscdk.Pipelines.StackOutputReference_FromCfnOutput(cfnOutput)
//
// Experimental.
type StackOutputReference interface {
	// Output name of the producing stack.
	// Experimental.
	OutputName() *string
	// A human-readable description of the producing stack.
	// Experimental.
	StackDescription() *string
	// Whether or not this stack output is being produced by the given Stack deployment.
	// Experimental.
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
// Experimental.
func StackOutputReference_FromCfnOutput(output awscdk.CfnOutput) StackOutputReference {
	_init_.Initialize()

	if err := validateStackOutputReference_FromCfnOutputParameters(output); err != nil {
		panic(err)
	}
	var returns StackOutputReference

	_jsii_.StaticInvoke(
		"monocdk.pipelines.StackOutputReference",
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

