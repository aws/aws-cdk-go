package interfacesawswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContentAssociation.
// Experimental.
type IContentAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ContentAssociation resource.
	// Experimental.
	ContentAssociationRef() *ContentAssociationReference
}

// The jsii proxy for IContentAssociationRef
type jsiiProxy_IContentAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IContentAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IContentAssociationRef) ContentAssociationRef() *ContentAssociationReference {
	var returns *ContentAssociationReference
	_jsii_.Get(
		j,
		"contentAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContentAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContentAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

