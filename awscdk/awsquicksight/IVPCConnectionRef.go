package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCConnection.
// Experimental.
type IVPCConnectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVPCConnectionRef
type jsiiProxy_IVPCConnectionRef struct {
	internal.Type__constructsIConstruct
}

