package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SignalMap.
// Experimental.
type ISignalMapRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISignalMapRef
type jsiiProxy_ISignalMapRef struct {
	internal.Type__constructsIConstruct
}

