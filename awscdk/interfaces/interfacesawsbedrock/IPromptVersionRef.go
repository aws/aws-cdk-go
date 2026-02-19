package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PromptVersion.
// Experimental.
type IPromptVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PromptVersion resource.
	// Experimental.
	PromptVersionRef() *PromptVersionReference
}

// The jsii proxy for IPromptVersionRef
type jsiiProxy_IPromptVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPromptVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPromptVersionRef) PromptVersionRef() *PromptVersionReference {
	var returns *PromptVersionReference
	_jsii_.Get(
		j,
		"promptVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

