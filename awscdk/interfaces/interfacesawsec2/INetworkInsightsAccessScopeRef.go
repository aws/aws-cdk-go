package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInsightsAccessScope.
// Experimental.
type INetworkInsightsAccessScopeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NetworkInsightsAccessScope resource.
	// Experimental.
	NetworkInsightsAccessScopeRef() *NetworkInsightsAccessScopeReference
}

// The jsii proxy for INetworkInsightsAccessScopeRef
type jsiiProxy_INetworkInsightsAccessScopeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_INetworkInsightsAccessScopeRef) NetworkInsightsAccessScopeRef() *NetworkInsightsAccessScopeReference {
	var returns *NetworkInsightsAccessScopeReference
	_jsii_.Get(
		j,
		"networkInsightsAccessScopeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInsightsAccessScopeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInsightsAccessScopeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

