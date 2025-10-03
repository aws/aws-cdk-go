package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventDataStore.
// Experimental.
type IEventDataStoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventDataStoreRef
type jsiiProxy_IEventDataStoreRef struct {
	internal.Type__constructsIConstruct
}

