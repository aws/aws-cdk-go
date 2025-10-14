package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInsightsAccessScope.
// Experimental.
type INetworkInsightsAccessScopeRef interface {
	constructs.IConstruct
	// A reference to a NetworkInsightsAccessScope resource.
	// Experimental.
	NetworkInsightsAccessScopeRef() *NetworkInsightsAccessScopeReference
}

// The jsii proxy for INetworkInsightsAccessScopeRef
type jsiiProxy_INetworkInsightsAccessScopeRef struct {
	internal.Type__constructsIConstruct
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

