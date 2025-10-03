package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WaitCondition.
// Experimental.
type IWaitConditionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWaitConditionRef
type jsiiProxy_IWaitConditionRef struct {
	internal.Type__constructsIConstruct
}

