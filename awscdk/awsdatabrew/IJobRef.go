package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Job.
// Experimental.
type IJobRef interface {
	constructs.IConstruct
}

// The jsii proxy for IJobRef
type jsiiProxy_IJobRef struct {
	internal.Type__constructsIConstruct
}

