package awsinspectorv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspectorv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeSecurityScanConfiguration.
// Experimental.
type ICodeSecurityScanConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICodeSecurityScanConfigurationRef
type jsiiProxy_ICodeSecurityScanConfigurationRef struct {
	internal.Type__constructsIConstruct
}

