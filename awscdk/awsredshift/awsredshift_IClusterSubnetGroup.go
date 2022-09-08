package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsredshift/internal"
)

// Interface for a cluster subnet group.
// Experimental.
type IClusterSubnetGroup interface {
	awscdk.IResource
	// The name of the cluster subnet group.
	// Experimental.
	ClusterSubnetGroupName() *string
}

// The jsii proxy for IClusterSubnetGroup
type jsiiProxy_IClusterSubnetGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IClusterSubnetGroup) ClusterSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetGroupName",
		&returns,
	)
	return returns
}

