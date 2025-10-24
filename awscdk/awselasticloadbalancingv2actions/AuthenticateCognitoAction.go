package awselasticloadbalancingv2actions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2actions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Listener Action to authenticate with Cognito.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc Vpc
//   var certificate Certificate
//
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   userPool := awscdk.Aws_cognito.NewUserPool(this, jsii.String("UserPool"))
//   userPoolClient := awscdk.Aws_cognito.NewUserPoolClient(this, jsii.String("Client"), &UserPoolClientProps{
//   	UserPool: UserPool,
//
//   	// Required minimal configuration for use with an ELB
//   	GenerateSecret: jsii.Boolean(true),
//   	AuthFlows: &AuthFlow{
//   		UserPassword: jsii.Boolean(true),
//   	},
//   	OAuth: &OAuthSettings{
//   		Flows: &OAuthFlows{
//   			AuthorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		Scopes: []OAuthScope{
//   			awscdk.*Aws_cognito.OAuthScope_EMAIL(),
//   		},
//   		CallbackUrls: []*string{
//   			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.LoadBalancerDnsName),
//   		},
//   	},
//   })
//   cfnClient := userPoolClient.Node.defaultChild.(CfnUserPoolClient)
//   cfnClient.AddPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
//   cfnClient.AddPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
//   	jsii.String("COGNITO"),
//   })
//
//   userPoolDomain := awscdk.Aws_cognito.NewUserPoolDomain(this, jsii.String("Domain"), &UserPoolDomainProps{
//   	UserPool: UserPool,
//   	CognitoDomain: &CognitoDomainOptions{
//   		DomainPrefix: jsii.String("test-cdk-prefix"),
//   	},
//   })
//
//   lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(443),
//   	Certificates: []IListenerCertificate{
//   		certificate,
//   	},
//   	DefaultAction: actions.NewAuthenticateCognitoAction(&AuthenticateCognitoActionProps{
//   		UserPool: *UserPool,
//   		UserPoolClient: *UserPoolClient,
//   		UserPoolDomain: *UserPoolDomain,
//   		Next: elbv2.ListenerAction_FixedResponse(jsii.Number(200), &FixedResponseOptions{
//   			ContentType: jsii.String("text/plain"),
//   			MessageBody: jsii.String("Authenticated"),
//   		}),
//   	}),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("DNS"), &CfnOutputProps{
//   	Value: lb.*LoadBalancerDnsName,
//   })
//
type AuthenticateCognitoAction interface {
	awselasticloadbalancingv2.ListenerAction
	Next() awselasticloadbalancingv2.ListenerAction
	// Sets the Action for the `ListenerRule`.
	//
	// This method is required to set a dedicated Action to a `ListenerRule`
	// when the Action for the `CfnListener` and the Action for the `CfnListenerRule`
	// have different structures. (e.g. `AuthenticateOidcConfig`)
	AddRuleAction(actionJson *awselasticloadbalancingv2.CfnListenerRule_ActionProperty)
	// Called when the action is being used in a listener.
	Bind(scope constructs.Construct, listener awselasticloadbalancingv2.IApplicationListener, associatingConstruct constructs.IConstruct)
	// Render the listener default actions in this chain.
	RenderActions() *[]*awselasticloadbalancingv2.CfnListener_ActionProperty
	// Render the listener rule actions in this chain.
	RenderRuleActions() *[]*awselasticloadbalancingv2.CfnListenerRule_ActionProperty
	// Renumber the "order" fields in the actions array.
	//
	// We don't number for 0 or 1 elements, but otherwise number them 1...#actions
	// so ELB knows about the right order.
	//
	// Do this in `ListenerAction` instead of in `Listener` so that we give
	// users the opportunity to override by subclassing and overriding `renderActions`.
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
func NewAuthenticateCognitoAction(options *AuthenticateCognitoActionProps) AuthenticateCognitoAction {
	_init_.Initialize()

	if err := validateNewAuthenticateCognitoActionParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthenticateCognitoAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Authenticate using an identity provide (IdP) that is compliant with OpenID Connect (OIDC).
func NewAuthenticateCognitoAction_Override(a AuthenticateCognitoAction, options *AuthenticateCognitoActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		[]interface{}{options},
		a,
	)
}

// Authenticate using an identity provider (IdP) that is compliant with OpenID Connect (OIDC).
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-authenticate-users.html#oidc-requirements
//
func AuthenticateCognitoAction_AuthenticateOidc(options *awselasticloadbalancingv2.AuthenticateOidcOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_AuthenticateOidcParameters(options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"authenticateOidc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Return a fixed response.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#fixed-response-actions
//
func AuthenticateCognitoAction_FixedResponse(statusCode *float64, options *awselasticloadbalancingv2.FixedResponseOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_FixedResponseParameters(statusCode, options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"fixedResponse",
		[]interface{}{statusCode, options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#forward-actions
//
func AuthenticateCognitoAction_Forward(targetGroups *[]awselasticloadbalancingv2.IApplicationTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_ForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
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
func AuthenticateCognitoAction_Redirect(options *awselasticloadbalancingv2.RedirectOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_RedirectParameters(options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"redirect",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups which are weighted differently.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#forward-actions
//
func AuthenticateCognitoAction_WeightedForward(targetGroups *[]*awselasticloadbalancingv2.WeightedTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_WeightedForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"weightedForward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticateCognitoAction) AddRuleAction(actionJson *awselasticloadbalancingv2.CfnListenerRule_ActionProperty) {
	if err := a.validateAddRuleActionParameters(actionJson); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addRuleAction",
		[]interface{}{actionJson},
	)
}

func (a *jsiiProxy_AuthenticateCognitoAction) Bind(scope constructs.Construct, listener awselasticloadbalancingv2.IApplicationListener, associatingConstruct constructs.IConstruct) {
	if err := a.validateBindParameters(scope, listener); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{scope, listener, associatingConstruct},
	)
}

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

func (a *jsiiProxy_AuthenticateCognitoAction) RenderRuleActions() *[]*awselasticloadbalancingv2.CfnListenerRule_ActionProperty {
	var returns *[]*awselasticloadbalancingv2.CfnListenerRule_ActionProperty

	_jsii_.Invoke(
		a,
		"renderRuleActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticateCognitoAction) Renumber(actions *[]*awselasticloadbalancingv2.CfnListener_ActionProperty) *[]*awselasticloadbalancingv2.CfnListener_ActionProperty {
	if err := a.validateRenumberParameters(actions); err != nil {
		panic(err)
	}
	var returns *[]*awselasticloadbalancingv2.CfnListener_ActionProperty

	_jsii_.Invoke(
		a,
		"renumber",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

