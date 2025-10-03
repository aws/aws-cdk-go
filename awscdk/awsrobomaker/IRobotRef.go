package awsrobomaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Robot.
// Experimental.
type IRobotRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRobotRef
type jsiiProxy_IRobotRef struct {
	internal.Type__constructsIConstruct
}

