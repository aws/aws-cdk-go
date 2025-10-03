package awsshield

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DRTAccess.
// Experimental.
type IDRTAccessRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDRTAccessRef
type jsiiProxy_IDRTAccessRef struct {
	internal.Type__constructsIConstruct
}

