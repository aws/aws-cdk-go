package interfacesawsoutposts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsoutposts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Outpost.
// Experimental.
type IOutpostRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Outpost resource.
	// Experimental.
	OutpostRef() *OutpostReference
}

// The jsii proxy for IOutpostRef
type jsiiProxy_IOutpostRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOutpostRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOutpostRef) OutpostRef() *OutpostReference {
	var returns *OutpostReference
	_jsii_.Get(
		j,
		"outpostRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOutpostRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOutpostRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

