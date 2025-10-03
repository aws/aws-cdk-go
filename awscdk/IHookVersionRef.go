package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookVersion.
// Experimental.
type IHookVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHookVersionRef
type jsiiProxy_IHookVersionRef struct {
	internal.Type__constructsIConstruct
}

