package interfacesawsram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SourceAssociation.
// Experimental.
type ISourceAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SourceAssociation resource.
	// Experimental.
	SourceAssociationRef() *SourceAssociationReference
}

// The jsii proxy for ISourceAssociationRef
type jsiiProxy_ISourceAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISourceAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISourceAssociationRef) SourceAssociationRef() *SourceAssociationReference {
	var returns *SourceAssociationReference
	_jsii_.Get(
		j,
		"sourceAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISourceAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

