package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterSecurityGroupIngress.
// Experimental.
type IClusterSecurityGroupIngressRef interface {
	constructs.IConstruct
	// A reference to a ClusterSecurityGroupIngress resource.
	// Experimental.
	ClusterSecurityGroupIngressRef() *ClusterSecurityGroupIngressReference
}

// The jsii proxy for IClusterSecurityGroupIngressRef
type jsiiProxy_IClusterSecurityGroupIngressRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IClusterSecurityGroupIngressRef) ClusterSecurityGroupIngressRef() *ClusterSecurityGroupIngressReference {
	var returns *ClusterSecurityGroupIngressReference
	_jsii_.Get(
		j,
		"clusterSecurityGroupIngressRef",
		&returns,
	)
	return returns
}

