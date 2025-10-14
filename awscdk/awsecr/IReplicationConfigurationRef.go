package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationConfiguration.
// Experimental.
type IReplicationConfigurationRef interface {
	constructs.IConstruct
	// A reference to a ReplicationConfiguration resource.
	// Experimental.
	ReplicationConfigurationRef() *ReplicationConfigurationReference
}

// The jsii proxy for IReplicationConfigurationRef
type jsiiProxy_IReplicationConfigurationRef struct {
	internal.Type__constructsIConstruct
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

