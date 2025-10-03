package awsshield

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Protection.
// Experimental.
type IProtectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProtectionRef
type jsiiProxy_IProtectionRef struct {
	internal.Type__constructsIConstruct
}

