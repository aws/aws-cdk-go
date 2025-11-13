package interfacesawssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OptOutList.
// Experimental.
type IOptOutListRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OptOutList resource.
	// Experimental.
	OptOutListRef() *OptOutListReference
}

// The jsii proxy for IOptOutListRef
type jsiiProxy_IOptOutListRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IOptOutListRef) OptOutListRef() *OptOutListReference {
	var returns *OptOutListReference
	_jsii_.Get(
		j,
		"optOutListRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptOutListRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptOutListRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

