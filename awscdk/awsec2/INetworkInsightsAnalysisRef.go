package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInsightsAnalysis.
// Experimental.
type INetworkInsightsAnalysisRef interface {
	constructs.IConstruct
	// A reference to a NetworkInsightsAnalysis resource.
	// Experimental.
	NetworkInsightsAnalysisRef() *NetworkInsightsAnalysisReference
}

// The jsii proxy for INetworkInsightsAnalysisRef
type jsiiProxy_INetworkInsightsAnalysisRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkInsightsAnalysisRef) NetworkInsightsAnalysisRef() *NetworkInsightsAnalysisReference {
	var returns *NetworkInsightsAnalysisReference
	_jsii_.Get(
		j,
		"networkInsightsAnalysisRef",
		&returns,
	)
	return returns
}

