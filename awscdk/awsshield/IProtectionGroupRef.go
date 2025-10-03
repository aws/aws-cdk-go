package awsshield

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProtectionGroup.
// Experimental.
type IProtectionGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProtectionGroupRef
type jsiiProxy_IProtectionGroupRef struct {
	internal.Type__constructsIConstruct
}

