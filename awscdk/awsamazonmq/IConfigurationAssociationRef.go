package awsamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsamazonmq/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationAssociation.
// Experimental.
type IConfigurationAssociationRef interface {
	constructs.IConstruct
	// A reference to a ConfigurationAssociation resource.
	// Experimental.
	ConfigurationAssociationRef() *ConfigurationAssociationReference
}

// The jsii proxy for IConfigurationAssociationRef
type jsiiProxy_IConfigurationAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationAssociationRef) ConfigurationAssociationRef() *ConfigurationAssociationReference {
	var returns *ConfigurationAssociationReference
	_jsii_.Get(
		j,
		"configurationAssociationRef",
		&returns,
	)
	return returns
}

