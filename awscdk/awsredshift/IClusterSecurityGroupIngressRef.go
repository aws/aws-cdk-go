package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterSecurityGroupIngress.
// Experimental.
type IClusterSecurityGroupIngressRef interface {
	constructs.IConstruct
}

// The jsii proxy for IClusterSecurityGroupIngressRef
type jsiiProxy_IClusterSecurityGroupIngressRef struct {
	internal.Type__constructsIConstruct
}

