package awssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationSet.
// Experimental.
type IConfigurationSetRef interface {
	constructs.IConstruct
	// A reference to a ConfigurationSet resource.
	// Experimental.
	ConfigurationSetRef() *ConfigurationSetReference
}

// The jsii proxy for IConfigurationSetRef
type jsiiProxy_IConfigurationSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationSetRef) ConfigurationSetRef() *ConfigurationSetReference {
	var returns *ConfigurationSetReference
	_jsii_.Get(
		j,
		"configurationSetRef",
		&returns,
	)
	return returns
}

