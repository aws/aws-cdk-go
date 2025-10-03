package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MountTarget.
// Experimental.
type IMountTargetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMountTargetRef
type jsiiProxy_IMountTargetRef struct {
	internal.Type__constructsIConstruct
}

