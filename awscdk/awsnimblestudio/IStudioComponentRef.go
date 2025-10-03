package awsnimblestudio

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StudioComponent.
// Experimental.
type IStudioComponentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStudioComponentRef
type jsiiProxy_IStudioComponentRef struct {
	internal.Type__constructsIConstruct
}

