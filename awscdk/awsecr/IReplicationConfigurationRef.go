package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationConfiguration.
// Experimental.
type IReplicationConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ReplicationConfiguration resource.
	// Experimental.
	ReplicationConfigurationRef() *ReplicationConfigurationReference
}

// The jsii proxy for IReplicationConfigurationRef
type jsiiProxy_IReplicationConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IReplicationConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

