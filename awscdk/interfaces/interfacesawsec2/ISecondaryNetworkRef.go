package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecondaryNetwork.
// Experimental.
type ISecondaryNetworkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SecondaryNetwork resource.
	// Experimental.
	SecondaryNetworkRef() *SecondaryNetworkReference
}

// The jsii proxy for ISecondaryNetworkRef
type jsiiProxy_ISecondaryNetworkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISecondaryNetworkRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISecondaryNetworkRef) SecondaryNetworkRef() *SecondaryNetworkReference {
	var returns *SecondaryNetworkReference
	_jsii_.Get(
		j,
		"secondaryNetworkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecondaryNetworkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecondaryNetworkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

