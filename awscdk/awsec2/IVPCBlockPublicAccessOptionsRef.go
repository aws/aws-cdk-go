package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCBlockPublicAccessOptions.
// Experimental.
type IVPCBlockPublicAccessOptionsRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVPCBlockPublicAccessOptionsRef
type jsiiProxy_IVPCBlockPublicAccessOptionsRef struct {
	internal.Type__constructsIConstruct
}

