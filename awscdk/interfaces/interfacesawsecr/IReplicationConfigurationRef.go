package interfacesawsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationConfiguration.
// Experimental.
type IReplicationConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReplicationConfiguration resource.
	// Experimental.
	ReplicationConfigurationRef() *ReplicationConfigurationReference
}

// The jsii proxy for IReplicationConfigurationRef
type jsiiProxy_IReplicationConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IReplicationConfigurationRef) ReplicationConfigurationRef() *ReplicationConfigurationReference {
	var returns *ReplicationConfigurationReference
	_jsii_.Get(
		j,
		"replicationConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

