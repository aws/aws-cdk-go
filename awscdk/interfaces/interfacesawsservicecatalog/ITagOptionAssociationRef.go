package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagOptionAssociation.
// Experimental.
type ITagOptionAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TagOptionAssociation resource.
	// Experimental.
	TagOptionAssociationRef() *TagOptionAssociationReference
}

// The jsii proxy for ITagOptionAssociationRef
type jsiiProxy_ITagOptionAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITagOptionAssociationRef) TagOptionAssociationRef() *TagOptionAssociationReference {
	var returns *TagOptionAssociationReference
	_jsii_.Get(
		j,
		"tagOptionAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagOptionAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITagOptionAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

