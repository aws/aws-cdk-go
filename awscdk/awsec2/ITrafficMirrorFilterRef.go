package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorFilter.
// Experimental.
type ITrafficMirrorFilterRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrafficMirrorFilterRef
type jsiiProxy_ITrafficMirrorFilterRef struct {
	internal.Type__constructsIConstruct
}

