package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
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
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkTargetGroup networkTargetGroup
//
//   networkListenerAction := awscdk.Aws_elasticloadbalancingv2.NetworkListenerAction_Forward([]iNetworkTargetGroup{
//   	networkTargetGroup,
//   }, &NetworkForwardOptions{
//   	StickinessDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   })
//
type NetworkListenerAction interface {
	IListenerAction
	Next() NetworkListenerAction
	// Called when the action is being used in a listener.
	Bind(scope constructs.Construct, listener INetworkListener)
	// Render the listener default actions in this chain.
	RenderActions() *[]*CfnListener_ActionProperty
	// Render the listener rule actions in this chain.
	RenderRuleActions() *[]*CfnListenerRule_ActionProperty
	// Renumber the "order" fields in the actions array.
	//
	// We don't number for 0 or 1 elements, but otherwise number them 1...#actions
	// so ELB knows about the right order.
	//
	// Do this in `NetworkListenerAction` instead of in `Listener` so that we give
	// users the opportunity to override by subclassing and overriding `renderActions`.
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
func NewNetworkListenerAction(defaultActionJson *CfnListener_ActionProperty, next NetworkListenerAction) NetworkListenerAction {
	_init_.Initialize()

	if err := validateNewNetworkListenerActionParameters(defaultActionJson); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkListenerAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerAction",
		[]interface{}{defaultActionJson, next},
		&j,
	)

	return &j
}

// Create an instance of NetworkListenerAction.
//
// The default class should be good enough for most cases and
// should be created by using one of the static factory functions,
// but allow overriding to make sure we allow flexibility for the future.
func NewNetworkListenerAction_Override(n NetworkListenerAction, defaultActionJson *CfnListener_ActionProperty, next NetworkListenerAction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerAction",
		[]interface{}{defaultActionJson, next},
		n,
	)
}

// Forward to one or more Target Groups.
func NetworkListenerAction_Forward(targetGroups *[]INetworkTargetGroup, options *NetworkForwardOptions) NetworkListenerAction {
	_init_.Initialize()

	if err := validateNetworkListenerAction_ForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns NetworkListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerAction",
		"forward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups which are weighted differently.
func NetworkListenerAction_WeightedForward(targetGroups *[]*NetworkWeightedTargetGroup, options *NetworkForwardOptions) NetworkListenerAction {
	_init_.Initialize()

	if err := validateNetworkListenerAction_WeightedForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns NetworkListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerAction",
		"weightedForward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkListenerAction) Bind(scope constructs.Construct, listener INetworkListener) {
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

func (n *jsiiProxy_NetworkListenerAction) RenderRuleActions() *[]*CfnListenerRule_ActionProperty {
	var returns *[]*CfnListenerRule_ActionProperty

	_jsii_.Invoke(
		n,
		"renderRuleActions",
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

