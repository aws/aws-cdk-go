package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterSecurityGroup.
// Experimental.
type IClusterSecurityGroupRef interface {
	constructs.IConstruct
	// A reference to a ClusterSecurityGroup resource.
	// Experimental.
	ClusterSecurityGroupRef() *ClusterSecurityGroupReference
}

// The jsii proxy for IClusterSecurityGroupRef
type jsiiProxy_IClusterSecurityGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IClusterSecurityGroupRef) ClusterSecurityGroupRef() *ClusterSecurityGroupReference {
	var returns *ClusterSecurityGroupReference
	_jsii_.Get(
		j,
		"clusterSecurityGroupRef",
		&returns,
	)
	return returns
}

