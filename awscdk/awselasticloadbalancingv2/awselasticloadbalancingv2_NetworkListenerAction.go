package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// What to do when a client makes a request to a listener.
//
// Some actions can be combined with other ones (specifically,
// you can perform authentication before serving the request).
//
// Multiple actions form a linked chain; the chain must always terminate in a
// *(weighted)forward*, *fixedResponse* or *redirect* action.
//
// If an action supports chaining, the next action can be indicated
// by passing it in the `next` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var networkTargetGroup networkTargetGroup
//
//   networkListenerAction := awscdk.Aws_elasticloadbalancingv2.networkListenerAction.forward([]iNetworkTargetGroup{
//   	networkTargetGroup,
//   }, &networkForwardOptions{
//   	stickinessDuration: duration,
//   })
//
// Experimental.
type NetworkListenerAction interface {
	IListenerAction
	// Experimental.
	Next() NetworkListenerAction
	// Called when the action is being used in a listener.
	// Experimental.
	Bind(scope awscdk.Construct, listener INetworkListener)
	// Render the actions in this chain.
	// Experimental.
	RenderActions() *[]*CfnListener_ActionProperty
	// Renumber the "order" fields in the actions array.
	//
	// We don't number for 0 or 1 elements, but otherwise number them 1...#actions
	// so ELB knows about the right order.
	//
	// Do this in `NetworkListenerAction` instead of in `Listener` so that we give
	// users the opportunity to override by subclassing and overriding `renderActions`.
	// Experimental.
	Renumber(actions *[]*CfnListener_ActionProperty) *[]*CfnListener_ActionProperty
}

// The jsii proxy struct for NetworkListenerAction
type jsiiProxy_NetworkListenerAction struct {
	jsiiProxy_IListenerAction
}

func (j *jsiiProxy_NetworkListenerAction) Next() NetworkListenerAction {
	var returns NetworkListenerAction
	_jsii_.Get(
		j,
		"next",
		&returns,
	)
	return returns
}


// Create an instance of NetworkListenerAction.
//
// The default class should be good enough for most cases and
// should be created by using one of the static factory functions,
// but allow overriding to make sure we allow flexibility for the future.
// Experimental.
func NewNetworkListenerAction(actionJson *CfnListener_ActionProperty, next NetworkListenerAction) NetworkListenerAction {
	_init_.Initialize()

	if err := validateNewNetworkListenerActionParameters(actionJson); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkListenerAction{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.NetworkListenerAction",
		[]interface{}{actionJson, next},
		&j,
	)

	return &j
}

// Create an instance of NetworkListenerAction.
//
// The default class should be good enough for most cases and
// should be created by using one of the static factory functions,
// but allow overriding to make sure we allow flexibility for the future.
// Experimental.
func NewNetworkListenerAction_Override(n NetworkListenerAction, actionJson *CfnListener_ActionProperty, next NetworkListenerAction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.NetworkListenerAction",
		[]interface{}{actionJson, next},
		n,
	)
}

// Forward to one or more Target Groups.
// Experimental.
func NetworkListenerAction_Forward(targetGroups *[]INetworkTargetGroup, options *NetworkForwardOptions) NetworkListenerAction {
	_init_.Initialize()

	if err := validateNetworkListenerAction_ForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns NetworkListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.NetworkListenerAction",
		"forward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups which are weighted differently.
// Experimental.
func NetworkListenerAction_WeightedForward(targetGroups *[]*NetworkWeightedTargetGroup, options *NetworkForwardOptions) NetworkListenerAction {
	_init_.Initialize()

	if err := validateNetworkListenerAction_WeightedForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns NetworkListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.NetworkListenerAction",
		"weightedForward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkListenerAction) Bind(scope awscdk.Construct, listener INetworkListener) {
	if err := n.validateBindParameters(scope, listener); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"bind",
		[]interface{}{scope, listener},
	)
}

func (n *jsiiProxy_NetworkListenerAction) RenderActions() *[]*CfnListener_ActionProperty {
	var returns *[]*CfnListener_ActionProperty

	_jsii_.Invoke(
		n,
		"renderActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkListenerAction) Renumber(actions *[]*CfnListener_ActionProperty) *[]*CfnListener_ActionProperty {
	if err := n.validateRenumberParameters(actions); err != nil {
		panic(err)
	}
	var returns *[]*CfnListener_ActionProperty

	_jsii_.Invoke(
		n,
		"renumber",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

