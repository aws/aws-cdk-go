package interfacesawstranslate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstranslate/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Terminology.
// Experimental.
type ITerminologyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Terminology resource.
	// Experimental.
	TerminologyRef() *TerminologyReference
}

// The jsii proxy for ITerminologyRef
type jsiiProxy_ITerminologyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITerminologyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITerminologyRef) TerminologyRef() *TerminologyReference {
	var returns *TerminologyReference
	_jsii_.Get(
		j,
		"terminologyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerminologyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerminologyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

