package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutomatedReasoningPolicyVersion.
// Experimental.
type IAutomatedReasoningPolicyVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AutomatedReasoningPolicyVersion resource.
	// Experimental.
	AutomatedReasoningPolicyVersionRef() *AutomatedReasoningPolicyVersionReference
}

// The jsii proxy for IAutomatedReasoningPolicyVersionRef
type jsiiProxy_IAutomatedReasoningPolicyVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAutomatedReasoningPolicyVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAutomatedReasoningPolicyVersionRef) AutomatedReasoningPolicyVersionRef() *AutomatedReasoningPolicyVersionReference {
	var returns *AutomatedReasoningPolicyVersionReference
	_jsii_.Get(
		j,
		"automatedReasoningPolicyVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutomatedReasoningPolicyVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutomatedReasoningPolicyVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

