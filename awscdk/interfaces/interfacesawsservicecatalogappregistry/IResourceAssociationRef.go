package interfacesawsservicecatalogappregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceAssociation.
// Experimental.
type IResourceAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ResourceAssociation resource.
	// Experimental.
	ResourceAssociationRef() *ResourceAssociationReference
}

// The jsii proxy for IResourceAssociationRef
type jsiiProxy_IResourceAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IResourceAssociationRef) ResourceAssociationRef() *ResourceAssociationReference {
	var returns *ResourceAssociationReference
	_jsii_.Get(
		j,
		"resourceAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

