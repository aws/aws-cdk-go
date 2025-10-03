package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterParameterGroup.
// Experimental.
type IClusterParameterGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IClusterParameterGroupRef
type jsiiProxy_IClusterParameterGroupRef struct {
	internal.Type__constructsIConstruct
}

