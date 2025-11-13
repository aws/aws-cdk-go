package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EvaluationForm.
// Experimental.
type IEvaluationFormRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EvaluationForm resource.
	// Experimental.
	EvaluationFormRef() *EvaluationFormReference
}

// The jsii proxy for IEvaluationFormRef
type jsiiProxy_IEvaluationFormRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEvaluationFormRef) EvaluationFormRef() *EvaluationFormReference {
	var returns *EvaluationFormReference
	_jsii_.Get(
		j,
		"evaluationFormRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluationFormRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluationFormRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

