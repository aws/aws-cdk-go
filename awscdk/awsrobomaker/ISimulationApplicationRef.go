package awsrobomaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimulationApplication.
// Experimental.
type ISimulationApplicationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISimulationApplicationRef
type jsiiProxy_ISimulationApplicationRef struct {
	internal.Type__constructsIConstruct
}

