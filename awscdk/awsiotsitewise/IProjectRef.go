package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Project.
// Experimental.
type IProjectRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProjectRef
type jsiiProxy_IProjectRef struct {
	internal.Type__constructsIConstruct
}

