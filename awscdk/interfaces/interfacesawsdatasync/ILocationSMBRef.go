package interfacesawsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationSMB.
// Experimental.
type ILocationSMBRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocationSMB resource.
	// Experimental.
	LocationSmbRef() *LocationSMBReference
}

// The jsii proxy for ILocationSMBRef
type jsiiProxy_ILocationSMBRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILocationSMBRef) LocationSmbRef() *LocationSMBReference {
	var returns *LocationSMBReference
	_jsii_.Get(
		j,
		"locationSmbRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationSMBRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationSMBRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

