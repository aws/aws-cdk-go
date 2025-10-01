package awssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProtectConfiguration.
// Experimental.
type IProtectConfigurationRef interface {
	constructs.IConstruct
	// A reference to a ProtectConfiguration resource.
	// Experimental.
	ProtectConfigurationRef() *ProtectConfigurationReference
}

// The jsii proxy for IProtectConfigurationRef
type jsiiProxy_IProtectConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProtectConfigurationRef) ProtectConfigurationRef() *ProtectConfigurationReference {
	var returns *ProtectConfigurationReference
	_jsii_.Get(
		j,
		"protectConfigurationRef",
		&returns,
	)
	return returns
}

