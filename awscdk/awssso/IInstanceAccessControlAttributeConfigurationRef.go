package awssso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceAccessControlAttributeConfiguration.
// Experimental.
type IInstanceAccessControlAttributeConfigurationRef interface {
	constructs.IConstruct
	// A reference to a InstanceAccessControlAttributeConfiguration resource.
	// Experimental.
	InstanceAccessControlAttributeConfigurationRef() *InstanceAccessControlAttributeConfigurationReference
}

// The jsii proxy for IInstanceAccessControlAttributeConfigurationRef
type jsiiProxy_IInstanceAccessControlAttributeConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInstanceAccessControlAttributeConfigurationRef) InstanceAccessControlAttributeConfigurationRef() *InstanceAccessControlAttributeConfigurationReference {
	var returns *InstanceAccessControlAttributeConfigurationReference
	_jsii_.Get(
		j,
		"instanceAccessControlAttributeConfigurationRef",
		&returns,
	)
	return returns
}

