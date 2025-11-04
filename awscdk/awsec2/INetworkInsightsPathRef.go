package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInsightsPath.
// Experimental.
type INetworkInsightsPathRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NetworkInsightsPath resource.
	// Experimental.
	NetworkInsightsPathRef() *NetworkInsightsPathReference
}

// The jsii proxy for INetworkInsightsPathRef
type jsiiProxy_INetworkInsightsPathRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INetworkInsightsPathRef) NetworkInsightsPathRef() *NetworkInsightsPathReference {
	var returns *NetworkInsightsPathReference
	_jsii_.Get(
		j,
		"networkInsightsPathRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInsightsPathRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInsightsPathRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

