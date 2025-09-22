package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExtensionAssociation.
// Experimental.
type IExtensionAssociationRef interface {
	constructs.IConstruct
	// A reference to a ExtensionAssociation resource.
	// Experimental.
	ExtensionAssociationRef() *ExtensionAssociationReference
}

// The jsii proxy for IExtensionAssociationRef
type jsiiProxy_IExtensionAssociationRef struct {
	internal.Type__constructsIConstruct
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

