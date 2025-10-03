package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Stack.
// Experimental.
type IStackRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStackRef
type jsiiProxy_IStackRef struct {
	internal.Type__constructsIConstruct
}

