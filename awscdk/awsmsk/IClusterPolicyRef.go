package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterPolicy.
// Experimental.
type IClusterPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IClusterPolicyRef
type jsiiProxy_IClusterPolicyRef struct {
	internal.Type__constructsIConstruct
}

