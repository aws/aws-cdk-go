package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// Collection of grant methods for a IEventBusRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventBusRef IEventBusRef
//
//   eventBusGrants := awscdk.Aws_events.EventBusGrants_FromEventBus(eventBusRef)
//
type EventBusGrants interface {
	Resource() interfacesawsevents.IEventBusRef
	// Permits an IAM Principal to send custom events to EventBridge so that they can be matched to rules.
	AllPutEvents(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for EventBusGrants
type jsiiProxy_EventBusGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_EventBusGrants) Resource() interfacesawsevents.IEventBusRef {
	var returns interfacesawsevents.IEventBusRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for EventBusGrants.
func EventBusGrants_FromEventBus(resource interfacesawsevents.IEventBusRef) EventBusGrants {
	_init_.Initialize()

	if err := validateEventBusGrants_FromEventBusParameters(resource); err != nil {
		panic(err)
	}
	var returns EventBusGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventBusGrants",
		"fromEventBus",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventBusGrants) AllPutEvents(grantee awsiam.IGrantable) awsiam.Grant {
	if err := e.validateAllPutEventsParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		e,
		"allPutEvents",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

