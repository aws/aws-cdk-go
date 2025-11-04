package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdNamespaceAssociation.
// Experimental.
type IIdNamespaceAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IdNamespaceAssociation resource.
	// Experimental.
	IdNamespaceAssociationRef() *IdNamespaceAssociationReference
}

// The jsii proxy for IIdNamespaceAssociationRef
type jsiiProxy_IIdNamespaceAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIdNamespaceAssociationRef) IdNamespaceAssociationRef() *IdNamespaceAssociationReference {
	var returns *IdNamespaceAssociationReference
	_jsii_.Get(
		j,
		"idNamespaceAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdNamespaceAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdNamespaceAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

