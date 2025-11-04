package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInsightsAccessScopeAnalysis.
// Experimental.
type INetworkInsightsAccessScopeAnalysisRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NetworkInsightsAccessScopeAnalysis resource.
	// Experimental.
	NetworkInsightsAccessScopeAnalysisRef() *NetworkInsightsAccessScopeAnalysisReference
}

// The jsii proxy for INetworkInsightsAccessScopeAnalysisRef
type jsiiProxy_INetworkInsightsAccessScopeAnalysisRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INetworkInsightsAccessScopeAnalysisRef) NetworkInsightsAccessScopeAnalysisRef() *NetworkInsightsAccessScopeAnalysisReference {
	var returns *NetworkInsightsAccessScopeAnalysisReference
	_jsii_.Get(
		j,
		"networkInsightsAccessScopeAnalysisRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInsightsAccessScopeAnalysisRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkInsightsAccessScopeAnalysisRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

