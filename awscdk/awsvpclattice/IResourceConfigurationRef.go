package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceConfiguration.
// Experimental.
type IResourceConfigurationRef interface {
	constructs.IConstruct
	// A reference to a ResourceConfiguration resource.
	// Experimental.
	ResourceConfigurationRef() *ResourceConfigurationReference
}

// The jsii proxy for IResourceConfigurationRef
type jsiiProxy_IResourceConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceConfigurationRef) ResourceConfigurationRef() *ResourceConfigurationReference {
	var returns *ResourceConfigurationReference
	_jsii_.Get(
		j,
		"resourceConfigurationRef",
		&returns,
	)
	return returns
}

