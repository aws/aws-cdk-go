package awspcs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputeNodeGroup.
// Experimental.
type IComputeNodeGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IComputeNodeGroupRef
type jsiiProxy_IComputeNodeGroupRef struct {
	internal.Type__constructsIConstruct
}

