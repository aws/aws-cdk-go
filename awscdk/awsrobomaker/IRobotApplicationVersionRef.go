package awsrobomaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RobotApplicationVersion.
// Experimental.
type IRobotApplicationVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRobotApplicationVersionRef
type jsiiProxy_IRobotApplicationVersionRef struct {
	internal.Type__constructsIConstruct
}

