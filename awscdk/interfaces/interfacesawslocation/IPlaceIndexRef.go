package interfacesawslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaceIndex.
// Experimental.
type IPlaceIndexRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PlaceIndex resource.
	// Experimental.
	PlaceIndexRef() *PlaceIndexReference
}

// The jsii proxy for IPlaceIndexRef
type jsiiProxy_IPlaceIndexRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPlaceIndexRef) PlaceIndexRef() *PlaceIndexReference {
	var returns *PlaceIndexReference
	_jsii_.Get(
		j,
		"placeIndexRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlaceIndexRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlaceIndexRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

