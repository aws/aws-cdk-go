package awsrobomaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RobotApplication.
// Experimental.
type IRobotApplicationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRobotApplicationRef
type jsiiProxy_IRobotApplicationRef struct {
	internal.Type__constructsIConstruct
}

