package awselasticloadbalancingv2actions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2actions/internal"
)

// A Listener Action to authenticate with Cognito.
//
// TODO: EXAMPLE
//
// Experimental.
type AuthenticateCognitoAction interface {
	awselasticloadbalancingv2.ListenerAction
	Next() awselasticloadbalancingv2.ListenerAction
	Bind(scope awscdk.Construct, listener awselasticloadbalancingv2.IApplicationListener, associatingConstruct awscdk.IConstruct)
	RenderActions() *[]*awselasticloadbalancingv2.CfnListener_ActionProperty
	Renumber(actions *[]*awselasticloadbalancingv2.CfnListener_ActionProperty) *[]*awselasticloadbalancingv2.CfnListener_ActionProperty
}

// The jsii proxy struct for AuthenticateCognitoAction
type jsiiProxy_AuthenticateCognitoAction struct {
	internal.Type__awselasticloadbalancingv2ListenerAction
}

func (j *jsiiProxy_AuthenticateCognitoAction) Next() awselasticloadbalancingv2.ListenerAction {
	var returns awselasticloadbalancingv2.ListenerAction
	_jsii_.Get(
		j,
		"next",
		&returns,
	)
	return returns
}


// Authenticate using an identity provide (IdP) that is compliant with OpenID Connect (OIDC).
// Experimental.
func NewAuthenticateCognitoAction(options *AuthenticateCognitoActionProps) AuthenticateCognitoAction {
	_init_.Initialize()

	j := jsiiProxy_AuthenticateCognitoAction{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Authenticate using an identity provide (IdP) that is compliant with OpenID Connect (OIDC).
// Experimental.
func NewAuthenticateCognitoAction_Override(a AuthenticateCognitoAction, options *AuthenticateCognitoActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		[]interface{}{options},
		a,
	)
}

// Authenticate using an identity provider (IdP) that is compliant with OpenID Connect (OIDC).
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-authenticate-users.html#oidc-requirements
//
// Experimental.
func AuthenticateCognitoAction_AuthenticateOidc(options *awselasticloadbalancingv2.AuthenticateOidcOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"authenticateOidc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Return a fixed response.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#fixed-response-actions
//
// Experimental.
func AuthenticateCognitoAction_FixedResponse(statusCode *float64, options *awselasticloadbalancingv2.FixedResponseOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"fixedResponse",
		[]interface{}{statusCode, options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#forward-actions
//
// Experimental.
func AuthenticateCognitoAction_Forward(targetGroups *[]awselasticloadbalancingv2.IApplicationTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
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
// Experimental.
func AuthenticateCognitoAction_Redirect(options *awselasticloadbalancingv2.RedirectOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"redirect",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups which are weighted differently.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#forward-actions
//
// Experimental.
func AuthenticateCognitoAction_WeightedForward(targetGroups *[]*awselasticloadbalancingv2.WeightedTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"weightedForward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

// Called when the action is being used in a listener.
// Experimental.
func (a *jsiiProxy_AuthenticateCognitoAction) Bind(scope awscdk.Construct, listener awselasticloadbalancingv2.IApplicationListener, associatingConstruct awscdk.IConstruct) {
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{scope, listener, associatingConstruct},
	)
}

// Render the actions in this chain.
// Experimental.
func (a *jsiiProxy_AuthenticateCognitoAction) RenderActions() *[]*awselasticloadbalancingv2.CfnListener_ActionProperty {
	var returns *[]*awselasticloadbalancingv2.CfnListener_ActionProperty

	_jsii_.Invoke(
		a,
		"renderActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renumber the "order" fields in the actions array.
//
// We don't number for 0 or 1 elements, but otherwise number them 1...#actions
// so ELB knows about the right order.
//
// Do this in `ListenerAction` instead of in `Listener` so that we give
// users the opportunity to override by subclassing and overriding `renderActions`.
// Experimental.
func (a *jsiiProxy_AuthenticateCognitoAction) Renumber(actions *[]*awselasticloadbalancingv2.CfnListener_ActionProperty) *[]*awselasticloadbalancingv2.CfnListener_ActionProperty {
	var returns *[]*awselasticloadbalancingv2.CfnListener_ActionProperty

	_jsii_.Invoke(
		a,
		"renumber",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

// Properties for AuthenticateCognitoAction.
//
// TODO: EXAMPLE
//
// Experimental.
type AuthenticateCognitoActionProps struct {
	// What action to execute next.
	//
	// Multiple actions form a linked chain; the chain must always terminate in a
	// (weighted)forward, fixedResponse or redirect action.
	// Experimental.
	Next awselasticloadbalancingv2.ListenerAction `json:"next" yaml:"next"`
	// The Amazon Cognito user pool.
	// Experimental.
	UserPool awscognito.IUserPool `json:"userPool" yaml:"userPool"`
	// The Amazon Cognito user pool client.
	// Experimental.
	UserPoolClient awscognito.IUserPoolClient `json:"userPoolClient" yaml:"userPoolClient"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	// Experimental.
	UserPoolDomain awscognito.IUserPoolDomain `json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	// Experimental.
	AuthenticationRequestExtraParams *map[string]*string `json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated.
	// Experimental.
	OnUnauthenticatedRequest awselasticloadbalancingv2.UnauthenticatedAction `json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP.
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	// Experimental.
	Scope *string `json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	// Experimental.
	SessionCookieName *string `json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session.
	// Experimental.
	SessionTimeout awscdk.Duration `json:"sessionTimeout" yaml:"sessionTimeout"`
}

