package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GuardHook.
// Experimental.
type IGuardHookRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGuardHookRef
type jsiiProxy_IGuardHookRef struct {
	internal.Type__constructsIConstruct
}

