package awsshield

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProactiveEngagement.
// Experimental.
type IProactiveEngagementRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProactiveEngagementRef
type jsiiProxy_IProactiveEngagementRef struct {
	internal.Type__constructsIConstruct
}

