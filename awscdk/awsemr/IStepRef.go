package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Step.
// Experimental.
type IStepRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Step resource.
	// Experimental.
	StepRef() *StepReference
}

// The jsii proxy for IStepRef
type jsiiProxy_IStepRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStepRef) StepRef() *StepReference {
	var returns *StepReference
	_jsii_.Get(
		j,
		"stepRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStepRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStepRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

