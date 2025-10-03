package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterSecurityGroup.
// Experimental.
type IClusterSecurityGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IClusterSecurityGroupRef
type jsiiProxy_IClusterSecurityGroupRef struct {
	internal.Type__constructsIConstruct
}

