package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for classes that can be the destination of a log Subscription.
type ILogSubscriptionDestination interface {
	// Return the properties required to send subscription events to this destination.
	//
	// If necessary, the destination can use the properties of the SubscriptionFilter
	// object itself to configure its permissions to allow the subscription to write
	// to it.
	//
	// The destination may reconfigure its own permissions in response to this
	// function call.
	Bind(scope constructs.Construct, sourceLogGroup ILogGroup) *LogSubscriptionDestinationConfig
}

// The jsii proxy for ILogSubscriptionDestination
type jsiiProxy_ILogSubscriptionDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_ILogSubscriptionDestination) Bind(scope constructs.Construct, sourceLogGroup ILogGroup) *LogSubscriptionDestinationConfig {
	if err := i.validateBindParameters(scope, sourceLogGroup); err != nil {
		panic(err)
	}
	var returns *LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, sourceLogGroup},
		&returns,
	)

	return returns
}

