package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExtensionAssociation.
// Experimental.
type IExtensionAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ExtensionAssociation resource.
	// Experimental.
	ExtensionAssociationRef() *ExtensionAssociationReference
}

// The jsii proxy for IExtensionAssociationRef
type jsiiProxy_IExtensionAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IExtensionAssociationRef) ExtensionAssociationRef() *ExtensionAssociationReference {
	var returns *ExtensionAssociationReference
	_jsii_.Get(
		j,
		"extensionAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtensionAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtensionAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

