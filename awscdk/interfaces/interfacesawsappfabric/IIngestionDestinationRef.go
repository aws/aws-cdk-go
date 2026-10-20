package interfacesawsappfabric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IngestionDestination.
// Experimental.
type IIngestionDestinationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IngestionDestination resource.
	// Experimental.
	IngestionDestinationRef() *IngestionDestinationReference
}

// The jsii proxy for IIngestionDestinationRef
type jsiiProxy_IIngestionDestinationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIngestionDestinationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIngestionDestinationRef) IngestionDestinationRef() *IngestionDestinationReference {
	var returns *IngestionDestinationReference
	_jsii_.Get(
		j,
		"ingestionDestinationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIngestionDestinationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIngestionDestinationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

