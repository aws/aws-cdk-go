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
// (Called `ListenerAction` instead of the more strictly correct
// `ListenerAction` because this is the class most users interact
// with, and we want to make it not too visually overwhelming).
//
// Example:
//   var listener applicationListener
//   var myTargetGroup applicationTargetGroup
//
//
//   listener.AddAction(jsii.String("DefaultAction"), &AddApplicationActionProps{
//   	Action: elbv2.ListenerAction_AuthenticateOidc(&AuthenticateOidcOptions{
//   		AuthorizationEndpoint: jsii.String("https://example.com/openid"),
//   		// Other OIDC properties here
//   		ClientId: jsii.String("..."),
//   		ClientSecret: awscdk.SecretValue_SecretsManager(jsii.String("...")),
//   		Issuer: jsii.String("..."),
//   		TokenEndpoint: jsii.String("..."),
//   		UserInfoEndpoint: jsii.String("..."),
//
//   		// Next
//   		Next: elbv2.ListenerAction_Forward([]iApplicationTargetGroup{
//   			myTargetGroup,
//   		}),
//   	}),
//   })
//
type ListenerAction interface {
	IListenerAction
	Next() ListenerAction
	// Sets the Action for the `ListenerRule`.
	//
	// This method is required to set a dedicated Action to a `ListenerRule`
	// when the Action for the `CfnListener` and the Action for the `CfnListenerRule`
	// have different structures. (e.g. `AuthenticateOidcConfig`)
	AddRuleAction(actionJson *CfnListenerRule_ActionProperty)
	// Called when the action is being used in a listener.
	Bind(scope constructs.Construct, listener IApplicationListener, associatingConstruct constructs.IConstruct)
	// Render the listener default actions in this chain.
	RenderActions() *[]*CfnListener_ActionProperty
	// Render the listener rule actions in this chain.
	RenderRuleActions() *[]*CfnListenerRule_ActionProperty
	// Renumber the "order" fields in the actions array.
	//
	// We don't number for 0 or 1 elements, but otherwise number them 1...#actions
	// so ELB knows about the right order.
	//
	// Do this in `ListenerAction` instead of in `Listener` so that we give
	// users the opportunity to override by subclassing and overriding `renderActions`.
	Renumber(actions *[]*CfnListener_ActionProperty) *[]*CfnListener_ActionProperty
}

// The jsii proxy struct for ListenerAction
type jsiiProxy_ListenerAction struct {
	jsiiProxy_IListenerAction
}

func (j *jsiiProxy_ListenerAction) Next() ListenerAction {
	var returns ListenerAction
	_jsii_.Get(
		j,
		"next",
		&returns,
	)
	return returns
}


// Create an instance of ListenerAction.
//
// The default class should be good enough for most cases and
// should be created by using one of the static factory functions,
// but allow overriding to make sure we allow flexibility for the future.
func NewListenerAction(defaultActionJson *CfnListener_ActionProperty, next ListenerAction) ListenerAction {
	_init_.Initialize()

	if err := validateNewListenerActionParameters(defaultActionJson); err != nil {
		panic(err)
	}
	j := jsiiProxy_ListenerAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		[]interface{}{defaultActionJson, next},
		&j,
	)

	return &j
}

// Create an instance of ListenerAction.
//
// The default class should be good enough for most cases and
// should be created by using one of the static factory functions,
// but allow overriding to make sure we allow flexibility for the future.
func NewListenerAction_Override(l ListenerAction, defaultActionJson *CfnListener_ActionProperty, next ListenerAction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		[]interface{}{defaultActionJson, next},
		l,
	)
}

// Authenticate using an identity provider (IdP) that is compliant with OpenID Connect (OIDC).
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-authenticate-users.html#oidc-requirements
//
func ListenerAction_AuthenticateOidc(options *AuthenticateOidcOptions) ListenerAction {
	_init_.Initialize()

	if err := validateListenerAction_AuthenticateOidcParameters(options); err != nil {
		panic(err)
	}
	var returns ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		"authenticateOidc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Return a fixed response.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#fixed-response-actions
//
func ListenerAction_FixedResponse(statusCode *float64, options *FixedResponseOptions) ListenerAction {
	_init_.Initialize()

	if err := validateListenerAction_FixedResponseParameters(statusCode, options); err != nil {
		panic(err)
	}
	var returns ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		"fixedResponse",
		[]interface{}{statusCode, options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#forward-actions
//
func ListenerAction_Forward(targetGroups *[]IApplicationTargetGroup, options *ForwardOptions) ListenerAction {
	_init_.Initialize()

	if err := validateListenerAction_ForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		"forward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

// Redirect to a different URI.
//
// A URI consists of the following components:
// protocol://hostname:port/path?query. You must modify at least one of the
// following components to avoid a redirect loop: protocol, hostname, port, or
// path. Any components that you do not modify retain their original values.
//
// You can reuse URI components using the following reserved keywords:
//
// - `#{protocol}`
// - `#{host}`
// - `#{port}`
// - `#{path}` (the leading "/" is removed)
// - `#{query}`
//
// For example, you can change the path to "/new/#{path}", the hostname to
// "example.#{host}", or the query to "#{query}&value=xyz".
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#redirect-actions
//
func ListenerAction_Redirect(options *RedirectOptions) ListenerAction {
	_init_.Initialize()

	if err := validateListenerAction_RedirectParameters(options); err != nil {
		panic(err)
	}
	var returns ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		"redirect",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups which are weighted differently.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#forward-actions
//
func ListenerAction_WeightedForward(targetGroups *[]*WeightedTargetGroup, options *ForwardOptions) ListenerAction {
	_init_.Initialize()

	if err := validateListenerAction_WeightedForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		"weightedForward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListenerAction) AddRuleAction(actionJson *CfnListenerRule_ActionProperty) {
	if err := l.validateAddRuleActionParameters(actionJson); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addRuleAction",
		[]interface{}{actionJson},
	)
}

func (l *jsiiProxy_ListenerAction) Bind(scope constructs.Construct, listener IApplicationListener, associatingConstruct constructs.IConstruct) {
	if err := l.validateBindParameters(scope, listener); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"bind",
		[]interface{}{scope, listener, associatingConstruct},
	)
}

func (l *jsiiProxy_ListenerAction) RenderActions() *[]*CfnListener_ActionProperty {
	var returns *[]*CfnListener_ActionProperty

	_jsii_.Invoke(
		l,
		"renderActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListenerAction) RenderRuleActions() *[]*CfnListenerRule_ActionProperty {
	var returns *[]*CfnListenerRule_ActionProperty

	_jsii_.Invoke(
		l,
		"renderRuleActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_ListenerAction) Renumber(actions *[]*CfnListener_ActionProperty) *[]*CfnListener_ActionProperty {
	if err := l.validateRenumberParameters(actions); err != nil {
		panic(err)
	}
	var returns *[]*CfnListener_ActionProperty

	_jsii_.Invoke(
		l,
		"renumber",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

