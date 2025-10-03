package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PatchBaseline.
// Experimental.
type IPatchBaselineRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPatchBaselineRef
type jsiiProxy_IPatchBaselineRef struct {
	internal.Type__constructsIConstruct
}

