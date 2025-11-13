package interfacesawsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReferenceStore.
// Experimental.
type IReferenceStoreRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReferenceStore resource.
	// Experimental.
	ReferenceStoreRef() *ReferenceStoreReference
}

// The jsii proxy for IReferenceStoreRef
type jsiiProxy_IReferenceStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IReferenceStoreRef) ReferenceStoreRef() *ReferenceStoreReference {
	var returns *ReferenceStoreReference
	_jsii_.Get(
		j,
		"referenceStoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReferenceStoreRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReferenceStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

