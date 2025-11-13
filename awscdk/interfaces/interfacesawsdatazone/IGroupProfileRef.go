package interfacesawsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupProfile.
// Experimental.
type IGroupProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GroupProfile resource.
	// Experimental.
	GroupProfileRef() *GroupProfileReference
}

// The jsii proxy for IGroupProfileRef
type jsiiProxy_IGroupProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IGroupProfileRef) GroupProfileRef() *GroupProfileReference {
	var returns *GroupProfileReference
	_jsii_.Get(
		j,
		"groupProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

