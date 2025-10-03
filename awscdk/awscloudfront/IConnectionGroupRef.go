package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectionGroup.
// Experimental.
type IConnectionGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConnectionGroupRef
type jsiiProxy_IConnectionGroupRef struct {
	internal.Type__constructsIConstruct
}

