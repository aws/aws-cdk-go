package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Collaboration.
// Experimental.
type ICollaborationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICollaborationRef
type jsiiProxy_ICollaborationRef struct {
	internal.Type__constructsIConstruct
}

