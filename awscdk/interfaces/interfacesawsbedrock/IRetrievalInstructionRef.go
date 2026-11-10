package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RetrievalInstruction.
// Experimental.
type IRetrievalInstructionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RetrievalInstruction resource.
	// Experimental.
	RetrievalInstructionRef() *RetrievalInstructionReference
}

// The jsii proxy for IRetrievalInstructionRef
type jsiiProxy_IRetrievalInstructionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRetrievalInstructionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRetrievalInstructionRef) RetrievalInstructionRef() *RetrievalInstructionReference {
	var returns *RetrievalInstructionReference
	_jsii_.Get(
		j,
		"retrievalInstructionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRetrievalInstructionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRetrievalInstructionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

