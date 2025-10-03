package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplify/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Branch.
// Experimental.
type IBranchRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBranchRef
type jsiiProxy_IBranchRef struct {
	internal.Type__constructsIConstruct
}

