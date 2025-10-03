package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppBlock.
// Experimental.
type IAppBlockRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAppBlockRef
type jsiiProxy_IAppBlockRef struct {
	internal.Type__constructsIConstruct
}

