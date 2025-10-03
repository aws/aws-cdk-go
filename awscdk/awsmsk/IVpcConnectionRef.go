package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcConnection.
// Experimental.
type IVpcConnectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVpcConnectionRef
type jsiiProxy_IVpcConnectionRef struct {
	internal.Type__constructsIConstruct
}

