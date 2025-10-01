package awscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventDataStore.
// Experimental.
type IEventDataStoreRef interface {
	constructs.IConstruct
	// A reference to a EventDataStore resource.
	// Experimental.
	EventDataStoreRef() *EventDataStoreReference
}

// The jsii proxy for IEventDataStoreRef
type jsiiProxy_IEventDataStoreRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEventDataStoreRef) EventDataStoreRef() *EventDataStoreReference {
	var returns *EventDataStoreReference
	_jsii_.Get(
		j,
		"eventDataStoreRef",
		&returns,
	)
	return returns
}

