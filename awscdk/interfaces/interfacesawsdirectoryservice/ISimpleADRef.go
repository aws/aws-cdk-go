package interfacesawsdirectoryservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdirectoryservice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimpleAD.
// Experimental.
type ISimpleADRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SimpleAD resource.
	// Experimental.
	SimpleAdRef() *SimpleADReference
}

// The jsii proxy for ISimpleADRef
type jsiiProxy_ISimpleADRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISimpleADRef) SimpleAdRef() *SimpleADReference {
	var returns *SimpleADReference
	_jsii_.Get(
		j,
		"simpleAdRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISimpleADRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISimpleADRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

