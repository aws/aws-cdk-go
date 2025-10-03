package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationConfiguration.
// Experimental.
type IReplicationConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReplicationConfigurationRef
type jsiiProxy_IReplicationConfigurationRef struct {
	internal.Type__constructsIConstruct
}

