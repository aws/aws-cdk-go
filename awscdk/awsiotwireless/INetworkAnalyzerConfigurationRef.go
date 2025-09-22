package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkAnalyzerConfiguration.
// Experimental.
type INetworkAnalyzerConfigurationRef interface {
	constructs.IConstruct
	// A reference to a NetworkAnalyzerConfiguration resource.
	// Experimental.
	NetworkAnalyzerConfigurationRef() *NetworkAnalyzerConfigurationReference
}

// The jsii proxy for INetworkAnalyzerConfigurationRef
type jsiiProxy_INetworkAnalyzerConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkAnalyzerConfigurationRef) NetworkAnalyzerConfigurationRef() *NetworkAnalyzerConfigurationReference {
	var returns *NetworkAnalyzerConfigurationReference
	_jsii_.Get(
		j,
		"networkAnalyzerConfigurationRef",
		&returns,
	)
	return returns
}

