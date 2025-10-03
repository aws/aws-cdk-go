package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Master.
// Experimental.
type IMasterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMasterRef
type jsiiProxy_IMasterRef struct {
	internal.Type__constructsIConstruct
}

