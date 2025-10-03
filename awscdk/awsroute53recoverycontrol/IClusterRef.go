package awsroute53recoverycontrol

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Cluster.
// Experimental.
type IClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IClusterRef
type jsiiProxy_IClusterRef struct {
	internal.Type__constructsIConstruct
}

