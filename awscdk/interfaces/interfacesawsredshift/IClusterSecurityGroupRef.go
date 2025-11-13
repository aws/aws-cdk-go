package interfacesawsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClusterSecurityGroup.
// Experimental.
type IClusterSecurityGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ClusterSecurityGroup resource.
	// Experimental.
	ClusterSecurityGroupRef() *ClusterSecurityGroupReference
}

// The jsii proxy for IClusterSecurityGroupRef
type jsiiProxy_IClusterSecurityGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IClusterSecurityGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterSecurityGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

