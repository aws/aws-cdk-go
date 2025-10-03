package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkAnalyzerConfiguration.
// Experimental.
type INetworkAnalyzerConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for INetworkAnalyzerConfigurationRef
type jsiiProxy_INetworkAnalyzerConfigurationRef struct {
	internal.Type__constructsIConstruct
}

