package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Evaluator.
// Experimental.
type IEvaluatorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Evaluator resource.
	// Experimental.
	EvaluatorRef() *EvaluatorReference
}

// The jsii proxy for IEvaluatorRef
type jsiiProxy_IEvaluatorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEvaluatorRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEvaluatorRef) EvaluatorRef() *EvaluatorReference {
	var returns *EvaluatorReference
	_jsii_.Get(
		j,
		"evaluatorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluatorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEvaluatorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

