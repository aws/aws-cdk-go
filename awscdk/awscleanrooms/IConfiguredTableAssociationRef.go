package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfiguredTableAssociation.
// Experimental.
type IConfiguredTableAssociationRef interface {
	constructs.IConstruct
	// A reference to a ConfiguredTableAssociation resource.
	// Experimental.
	ConfiguredTableAssociationRef() *ConfiguredTableAssociationReference
}

// The jsii proxy for IConfiguredTableAssociationRef
type jsiiProxy_IConfiguredTableAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfiguredTableAssociationRef) ConfiguredTableAssociationRef() *ConfiguredTableAssociationReference {
	var returns *ConfiguredTableAssociationReference
	_jsii_.Get(
		j,
		"configuredTableAssociationRef",
		&returns,
	)
	return returns
}

