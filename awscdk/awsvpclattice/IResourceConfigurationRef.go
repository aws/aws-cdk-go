package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceConfiguration.
// Experimental.
type IResourceConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResourceConfiguration resource.
	// Experimental.
	ResourceConfigurationRef() *ResourceConfigurationReference
}

// The jsii proxy for IResourceConfigurationRef
type jsiiProxy_IResourceConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IResourceConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

