package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterSubnetGroup.
// Experimental.
type IClusterSubnetGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ClusterSubnetGroup resource.
	// Experimental.
	ClusterSubnetGroupRef() *ClusterSubnetGroupReference
}

// The jsii proxy for IClusterSubnetGroupRef
type jsiiProxy_IClusterSubnetGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IClusterSubnetGroupRef) ClusterSubnetGroupRef() *ClusterSubnetGroupReference {
	var returns *ClusterSubnetGroupReference
	_jsii_.Get(
		j,
		"clusterSubnetGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterSubnetGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterSubnetGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

