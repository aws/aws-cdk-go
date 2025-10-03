package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TLSInspectionConfiguration.
// Experimental.
type ITLSInspectionConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITLSInspectionConfigurationRef
type jsiiProxy_ITLSInspectionConfigurationRef struct {
	internal.Type__constructsIConstruct
}

