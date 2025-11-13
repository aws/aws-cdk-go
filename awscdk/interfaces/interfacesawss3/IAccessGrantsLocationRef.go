package interfacesawss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessGrantsLocation.
// Experimental.
type IAccessGrantsLocationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AccessGrantsLocation resource.
	// Experimental.
	AccessGrantsLocationRef() *AccessGrantsLocationReference
}

// The jsii proxy for IAccessGrantsLocationRef
type jsiiProxy_IAccessGrantsLocationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAccessGrantsLocationRef) AccessGrantsLocationRef() *AccessGrantsLocationReference {
	var returns *AccessGrantsLocationReference
	_jsii_.Get(
		j,
		"accessGrantsLocationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessGrantsLocationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessGrantsLocationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

