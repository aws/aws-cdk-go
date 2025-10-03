package awsnimblestudio

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Studio.
// Experimental.
type IStudioRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStudioRef
type jsiiProxy_IStudioRef struct {
	internal.Type__constructsIConstruct
}

