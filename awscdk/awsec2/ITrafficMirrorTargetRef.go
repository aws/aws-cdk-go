package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorTarget.
// Experimental.
type ITrafficMirrorTargetRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrafficMirrorTargetRef
type jsiiProxy_ITrafficMirrorTargetRef struct {
	internal.Type__constructsIConstruct
}

