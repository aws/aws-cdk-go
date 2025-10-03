package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrolesanywhere/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Profile.
// Experimental.
type IProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProfileRef
type jsiiProxy_IProfileRef struct {
	internal.Type__constructsIConstruct
}

