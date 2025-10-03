package awsgroundstation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MissionProfile.
// Experimental.
type IMissionProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMissionProfileRef
type jsiiProxy_IMissionProfileRef struct {
	internal.Type__constructsIConstruct
}

