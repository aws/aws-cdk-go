package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Mesh.
// Experimental.
type IMeshRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMeshRef
type jsiiProxy_IMeshRef struct {
	internal.Type__constructsIConstruct
}

