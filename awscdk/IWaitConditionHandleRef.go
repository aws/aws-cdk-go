package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WaitConditionHandle.
// Experimental.
type IWaitConditionHandleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWaitConditionHandleRef
type jsiiProxy_IWaitConditionHandleRef struct {
	internal.Type__constructsIConstruct
}

