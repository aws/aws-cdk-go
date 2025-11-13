package interfacesawsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SignalCatalog.
// Experimental.
type ISignalCatalogRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SignalCatalog resource.
	// Experimental.
	SignalCatalogRef() *SignalCatalogReference
}

// The jsii proxy for ISignalCatalogRef
type jsiiProxy_ISignalCatalogRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISignalCatalogRef) SignalCatalogRef() *SignalCatalogReference {
	var returns *SignalCatalogReference
	_jsii_.Get(
		j,
		"signalCatalogRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISignalCatalogRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISignalCatalogRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

