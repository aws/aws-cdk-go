package awsinspectorv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspectorv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CisScanConfiguration.
// Experimental.
type ICisScanConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICisScanConfigurationRef
type jsiiProxy_ICisScanConfigurationRef struct {
	internal.Type__constructsIConstruct
}

