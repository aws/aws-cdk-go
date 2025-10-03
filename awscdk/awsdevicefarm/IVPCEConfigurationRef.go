package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCEConfiguration.
// Experimental.
type IVPCEConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVPCEConfigurationRef
type jsiiProxy_IVPCEConfigurationRef struct {
	internal.Type__constructsIConstruct
}

