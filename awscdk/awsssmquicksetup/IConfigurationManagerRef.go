package awsssmquicksetup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmquicksetup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationManager.
// Experimental.
type IConfigurationManagerRef interface {
	constructs.IConstruct
	// A reference to a ConfigurationManager resource.
	// Experimental.
	ConfigurationManagerRef() *ConfigurationManagerReference
}

// The jsii proxy for IConfigurationManagerRef
type jsiiProxy_IConfigurationManagerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConfigurationManagerRef) ConfigurationManagerRef() *ConfigurationManagerReference {
	var returns *ConfigurationManagerReference
	_jsii_.Get(
		j,
		"configurationManagerRef",
		&returns,
	)
	return returns
}

