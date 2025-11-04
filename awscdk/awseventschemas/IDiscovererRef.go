package awseventschemas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Discoverer.
// Experimental.
type IDiscovererRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Discoverer resource.
	// Experimental.
	DiscovererRef() *DiscovererReference
}

// The jsii proxy for IDiscovererRef
type jsiiProxy_IDiscovererRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDiscovererRef) DiscovererRef() *DiscovererReference {
	var returns *DiscovererReference
	_jsii_.Get(
		j,
		"discovererRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDiscovererRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDiscovererRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

