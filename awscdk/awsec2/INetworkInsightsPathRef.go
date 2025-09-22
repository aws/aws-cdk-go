package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkInsightsPath.
// Experimental.
type INetworkInsightsPathRef interface {
	constructs.IConstruct
	// A reference to a NetworkInsightsPath resource.
	// Experimental.
	NetworkInsightsPathRef() *NetworkInsightsPathReference
}

// The jsii proxy for INetworkInsightsPathRef
type jsiiProxy_INetworkInsightsPathRef struct {
	internal.Type__constructsIConstruct
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

