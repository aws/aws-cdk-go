package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkGroup.
// Experimental.
type IWorkGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkGroupRef
type jsiiProxy_IWorkGroupRef struct {
	internal.Type__constructsIConstruct
}

