package interfacesawsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MeteredProduct.
// Experimental.
type IMeteredProductRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MeteredProduct resource.
	// Experimental.
	MeteredProductRef() *MeteredProductReference
}

// The jsii proxy for IMeteredProductRef
type jsiiProxy_IMeteredProductRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMeteredProductRef) MeteredProductRef() *MeteredProductReference {
	var returns *MeteredProductReference
	_jsii_.Get(
		j,
		"meteredProductRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMeteredProductRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMeteredProductRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

