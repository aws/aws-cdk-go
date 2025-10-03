package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookTypeConfig.
// Experimental.
type IHookTypeConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHookTypeConfigRef
type jsiiProxy_IHookTypeConfigRef struct {
	internal.Type__constructsIConstruct
}

