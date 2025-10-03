package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookDefaultVersion.
// Experimental.
type IHookDefaultVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHookDefaultVersionRef
type jsiiProxy_IHookDefaultVersionRef struct {
	internal.Type__constructsIConstruct
}

