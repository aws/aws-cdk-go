package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrefixList.
// Experimental.
type IPrefixListRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PrefixList resource.
	// Experimental.
	PrefixListRef() *PrefixListReference
}

// The jsii proxy for IPrefixListRef
type jsiiProxy_IPrefixListRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPrefixListRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPrefixListRef) PrefixListRef() *PrefixListReference {
	var returns *PrefixListReference
	_jsii_.Get(
		j,
		"prefixListRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixListRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrefixListRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

