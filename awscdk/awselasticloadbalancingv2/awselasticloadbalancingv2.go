package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties for adding a new action to a listener.
//
// TODO: EXAMPLE
//
type AddApplicationActionProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `json:"priority" yaml:"priority"`
	// Action to perform.
	Action ListenerAction `json:"action" yaml:"action"`
}

// Properties for adding a new target group to a listener.
//
// TODO: EXAMPLE
//
type AddApplicationTargetGroupsProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `json:"priority" yaml:"priority"`
	// Target groups to forward requests to.
	TargetGroups *[]IApplicationTargetGroup `json:"targetGroups" yaml:"targetGroups"`
}

// Properties for adding new targets to a listener.
//
// TODO: EXAMPLE
//
type AddApplicationTargetsProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	HealthCheck *HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// The load balancing algorithm to select targets for routing requests.
	LoadBalancingAlgorithmType TargetGroupLoadBalancingAlgorithmType `json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// The port on which the listener listens for requests.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `json:"protocol" yaml:"protocol"`
	// The protocol version to use.
	ProtocolVersion ApplicationProtocolVersion `json:"protocolVersion" yaml:"protocolVersion"`
	// The time period during which the load balancer sends a newly registered target a linearly increasing share of the traffic to the target group.
	//
	// The range is 30-900 seconds (15 minutes).
	SlowStart awscdk.Duration `json:"slowStart" yaml:"slowStart"`
	// The stickiness cookie expiration period.
	//
	// Setting this value enables load balancer stickiness.
	//
	// After this period, the cookie is considered stale. The minimum value is
	// 1 second and the maximum value is 7 days (604800 seconds).
	StickinessCookieDuration awscdk.Duration `json:"stickinessCookieDuration" yaml:"stickinessCookieDuration"`
	// The name of an application-based stickiness cookie.
	//
	// Names that start with the following prefixes are not allowed: AWSALB, AWSALBAPP,
	// and AWSALBTG; they're reserved for use by the load balancer.
	//
	// Note: `stickinessCookieName` parameter depends on the presence of `stickinessCookieDuration` parameter.
	// If `stickinessCookieDuration` is not set, `stickinessCookieName` will be omitted.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/sticky-sessions.html
	//
	StickinessCookieName *string `json:"stickinessCookieName" yaml:"stickinessCookieName"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `json:"targetGroupName" yaml:"targetGroupName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. All target must be of the same type.
	Targets *[]IApplicationLoadBalancerTarget `json:"targets" yaml:"targets"`
}

// Properties for adding a new action to a listener.
//
// TODO: EXAMPLE
//
type AddNetworkActionProps struct {
	// Action to perform.
	Action NetworkListenerAction `json:"action" yaml:"action"`
}

// Properties for adding new network targets to a listener.
//
// TODO: EXAMPLE
//
type AddNetworkTargetsProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `json:"port" yaml:"port"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	HealthCheck *HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// Indicates whether client IP preservation is enabled.
	PreserveClientIp *bool `json:"preserveClientIp" yaml:"preserveClientIp"`
	// Protocol for target group, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// Indicates whether Proxy Protocol version 2 is enabled.
	ProxyProtocolV2 *bool `json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `json:"targetGroupName" yaml:"targetGroupName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]INetworkLoadBalancerTarget `json:"targets" yaml:"targets"`
}

// Properties for adding a conditional load balancing rule.
//
// TODO: EXAMPLE
//
type AddRuleProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// Application-Layer Protocol Negotiation Policies for network load balancers.
//
// Which protocols should be used over a secure connection.
type AlpnPolicy string

const (
	AlpnPolicy_HTTP1_ONLY AlpnPolicy = "HTTP1_ONLY"
	AlpnPolicy_HTTP2_ONLY AlpnPolicy = "HTTP2_ONLY"
	AlpnPolicy_HTTP2_OPTIONAL AlpnPolicy = "HTTP2_OPTIONAL"
	AlpnPolicy_HTTP2_PREFERRED AlpnPolicy = "HTTP2_PREFERRED"
	AlpnPolicy_NONE AlpnPolicy = "NONE"
)

// Define an ApplicationListener.
//
// TODO: EXAMPLE
//
type ApplicationListener interface {
	BaseListener
	IApplicationListener
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	LoadBalancer() IApplicationLoadBalancer
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddAction(id *string, props *AddApplicationActionProps)
	AddCertificates(id *string, certificates *[]IListenerCertificate)
	AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps)
	AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
	ToString() *string
	ValidateListener() *[]*string
}

// The jsii proxy struct for ApplicationListener
type jsiiProxy_ApplicationListener struct {
	jsiiProxy_BaseListener
	jsiiProxy_IApplicationListener
}

func (j *jsiiProxy_ApplicationListener) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) LoadBalancer() IApplicationLoadBalancer {
	var returns IApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewApplicationListener(scope constructs.Construct, id *string, props *ApplicationListenerProps) ApplicationListener {
	_init_.Initialize()

	j := jsiiProxy_ApplicationListener{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationListener_Override(a ApplicationListener, scope constructs.Construct, id *string, props *ApplicationListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an existing listener.
func ApplicationListener_FromApplicationListenerAttributes(scope constructs.Construct, id *string, attrs *ApplicationListenerAttributes) IApplicationListener {
	_init_.Initialize()

	var returns IApplicationListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"fromApplicationListenerAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Look up an ApplicationListener.
func ApplicationListener_FromLookup(scope constructs.Construct, id *string, options *ApplicationListenerLookupOptions) IApplicationListener {
	_init_.Initialize()

	var returns IApplicationListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"fromLookup",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ApplicationListener_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Perform the given default action on incoming requests.
//
// This allows full control of the default action of the load balancer,
// including Action chaining, fixed responses and redirect responses. See
// the `ListenerAction` class for all options.
//
// It's possible to add routing conditions to the Action added in this way.
// At least one Action must be added without conditions (which becomes the
// default Action).
func (a *jsiiProxy_ApplicationListener) AddAction(id *string, props *AddApplicationActionProps) {
	_jsii_.InvokeVoid(
		a,
		"addAction",
		[]interface{}{id, props},
	)
}

// Add one or more certificates to this listener.
//
// After the first certificate, this creates ApplicationListenerCertificates
// resources since cloudformation requires the certificates array on the
// listener resource to have a length of 1.
func (a *jsiiProxy_ApplicationListener) AddCertificates(id *string, certificates *[]IListenerCertificate) {
	_jsii_.InvokeVoid(
		a,
		"addCertificates",
		[]interface{}{id, certificates},
	)
}

// Load balance incoming requests to the given target groups.
//
// All target groups will be load balanced to with equal weight and without
// stickiness. For a more complex configuration than that, use `addAction()`.
//
// It's possible to add routing conditions to the TargetGroups added in this
// way. At least one TargetGroup must be added without conditions (which will
// become the default Action for this listener).
func (a *jsiiProxy_ApplicationListener) AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps) {
	_jsii_.InvokeVoid(
		a,
		"addTargetGroups",
		[]interface{}{id, props},
	)
}

// Load balance incoming requests to the given load balancing targets.
//
// This method implicitly creates an ApplicationTargetGroup for the targets
// involved, and a 'forward' action to route traffic to the given TargetGroup.
//
// If you want more control over the precise setup, create the TargetGroup
// and use `addAction` yourself.
//
// It's possible to add conditions to the targets added in this way. At least
// one set of targets must be added without conditions.
//
// Returns: The newly created target group
func (a *jsiiProxy_ApplicationListener) AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup {
	var returns ApplicationTargetGroup

	_jsii_.Invoke(
		a,
		"addTargets",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (a *jsiiProxy_ApplicationListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ApplicationListener) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (a *jsiiProxy_ApplicationListener) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (a *jsiiProxy_ApplicationListener) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Register that a connectable that has been added to this load balancer.
//
// Don't call this directly. It is called by ApplicationTargetGroup.
func (a *jsiiProxy_ApplicationListener) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	_jsii_.InvokeVoid(
		a,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApplicationListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate this listener.
func (a *jsiiProxy_ApplicationListener) ValidateListener() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validateListener",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to reference an existing listener.
//
// TODO: EXAMPLE
//
type ApplicationListenerAttributes struct {
	// ARN of the listener.
	ListenerArn *string `json:"listenerArn" yaml:"listenerArn"`
	// The default port on which this listener is listening.
	DefaultPort *float64 `json:"defaultPort" yaml:"defaultPort"`
	// Security group of the load balancer this listener is associated with.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
}

// Add certificates to a listener.
//
// TODO: EXAMPLE
//
type ApplicationListenerCertificate interface {
	constructs.Construct
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for ApplicationListenerCertificate
type jsiiProxy_ApplicationListenerCertificate struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationListenerCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewApplicationListenerCertificate(scope constructs.Construct, id *string, props *ApplicationListenerCertificateProps) ApplicationListenerCertificate {
	_init_.Initialize()

	j := jsiiProxy_ApplicationListenerCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationListenerCertificate_Override(a ApplicationListenerCertificate, scope constructs.Construct, id *string, props *ApplicationListenerCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerCertificate",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationListenerCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApplicationListenerCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for adding a set of certificates to a listener.
//
// TODO: EXAMPLE
//
type ApplicationListenerCertificateProps struct {
	// The listener to attach the rule to.
	Listener IApplicationListener `json:"listener" yaml:"listener"`
	// Certificates to attach.
	//
	// Duplicates are not allowed.
	Certificates *[]IListenerCertificate `json:"certificates" yaml:"certificates"`
}

// Options for ApplicationListener lookup.
//
// TODO: EXAMPLE
//
type ApplicationListenerLookupOptions struct {
	// Filter listeners by listener port.
	ListenerPort *float64 `json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	LoadBalancerTags *map[string]*string `json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// ARN of the listener to look up.
	ListenerArn *string `json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener protocol.
	ListenerProtocol ApplicationProtocol `json:"listenerProtocol" yaml:"listenerProtocol"`
}

// Properties for defining a standalone ApplicationListener.
//
// TODO: EXAMPLE
//
type ApplicationListenerProps struct {
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses.
	//
	// See the `ListenerAction` class for all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction ListenerAction `json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]IApplicationTargetGroup `json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Allow anyone to connect to the load balancer on the listener port.
	//
	// If this is specified, the load balancer will be opened up to anyone who can reach it.
	// For internal load balancers this is anyone in the same VPC. For public load
	// balancers, this is anyone on the internet.
	//
	// If you want to be more selective about who can access this load
	// balancer, set this to `false` and use the listener's `connections`
	// object to selectively grant access to the load balancer on the listener port.
	Open *bool `json:"open" yaml:"open"`
	// The port on which the listener listens for requests.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `json:"protocol" yaml:"protocol"`
	// The security policy that defines which ciphers and protocols are supported.
	SslPolicy SslPolicy `json:"sslPolicy" yaml:"sslPolicy"`
	// The load balancer to attach this listener to.
	LoadBalancer IApplicationLoadBalancer `json:"loadBalancer" yaml:"loadBalancer"`
}

// Define a new listener rule.
//
// TODO: EXAMPLE
//
type ApplicationListenerRule interface {
	constructs.Construct
	ListenerRuleArn() *string
	Node() constructs.Node
	AddCondition(condition ListenerCondition)
	ConfigureAction(action ListenerAction)
	ToString() *string
}

// The jsii proxy struct for ApplicationListenerRule
type jsiiProxy_ApplicationListenerRule struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationListenerRule) ListenerRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListenerRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewApplicationListenerRule(scope constructs.Construct, id *string, props *ApplicationListenerRuleProps) ApplicationListenerRule {
	_init_.Initialize()

	j := jsiiProxy_ApplicationListenerRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationListenerRule_Override(a ApplicationListenerRule, scope constructs.Construct, id *string, props *ApplicationListenerRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerRule",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationListenerRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a non-standard condition to this rule.
func (a *jsiiProxy_ApplicationListenerRule) AddCondition(condition ListenerCondition) {
	_jsii_.InvokeVoid(
		a,
		"addCondition",
		[]interface{}{condition},
	)
}

// Configure the action to perform for this rule.
func (a *jsiiProxy_ApplicationListenerRule) ConfigureAction(action ListenerAction) {
	_jsii_.InvokeVoid(
		a,
		"configureAction",
		[]interface{}{action},
	)
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApplicationListenerRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for defining a listener rule.
//
// TODO: EXAMPLE
//
type ApplicationListenerRuleProps struct {
	// Priority of the rule.
	//
	// The rule with the lowest priority will be used for every request.
	//
	// Priorities must be unique.
	Priority *float64 `json:"priority" yaml:"priority"`
	// Action to perform when requests are received.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	Action ListenerAction `json:"action" yaml:"action"`
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `json:"conditions" yaml:"conditions"`
	// Target groups to forward requests to.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	//
	// Implies a `forward` action.
	TargetGroups *[]IApplicationTargetGroup `json:"targetGroups" yaml:"targetGroups"`
	// The listener to attach the rule to.
	Listener IApplicationListener `json:"listener" yaml:"listener"`
}

// Define an Application Load Balancer.
//
// TODO: EXAMPLE
//
type ApplicationLoadBalancer interface {
	BaseLoadBalancer
	IApplicationLoadBalancer
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	IpAddressType() IpAddressType
	Listeners() *[]ApplicationListener
	LoadBalancerArn() *string
	LoadBalancerCanonicalHostedZoneId() *string
	LoadBalancerDnsName() *string
	LoadBalancerFullName() *string
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener
	AddRedirect(props *ApplicationLoadBalancerRedirectConfig) ApplicationListener
	AddSecurityGroup(securityGroup awsec2.ISecurityGroup)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricActiveConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricClientTlsNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricElbAuthError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricElbAuthFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricElbAuthLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricElbAuthSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHttpCodeElb(code HttpCodeElb, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHttpFixedResponseCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHttpRedirectCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHttpRedirectUrlLimitExceededCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIpv6ProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIpv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNewConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRejectedConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRuleEvaluations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	RemoveAttribute(key *string)
	SetAttribute(key *string, value *string)
	ToString() *string
}

// The jsii proxy struct for ApplicationLoadBalancer
type jsiiProxy_ApplicationLoadBalancer struct {
	jsiiProxy_BaseLoadBalancer
	jsiiProxy_IApplicationLoadBalancer
}

func (j *jsiiProxy_ApplicationLoadBalancer) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) IpAddressType() IpAddressType {
	var returns IpAddressType
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Listeners() *[]ApplicationListener {
	var returns *[]ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) LoadBalancerSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewApplicationLoadBalancer(scope constructs.Construct, id *string, props *ApplicationLoadBalancerProps) ApplicationLoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_ApplicationLoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationLoadBalancer_Override(a ApplicationLoadBalancer, scope constructs.Construct, id *string, props *ApplicationLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an existing Application Load Balancer.
func ApplicationLoadBalancer_FromApplicationLoadBalancerAttributes(scope constructs.Construct, id *string, attrs *ApplicationLoadBalancerAttributes) IApplicationLoadBalancer {
	_init_.Initialize()

	var returns IApplicationLoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"fromApplicationLoadBalancerAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Look up an application load balancer.
func ApplicationLoadBalancer_FromLookup(scope constructs.Construct, id *string, options *ApplicationLoadBalancerLookupOptions) IApplicationLoadBalancer {
	_init_.Initialize()

	var returns IApplicationLoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"fromLookup",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ApplicationLoadBalancer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a new listener to this load balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener {
	var returns ApplicationListener

	_jsii_.Invoke(
		a,
		"addListener",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Add a redirection listener to this load balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) AddRedirect(props *ApplicationLoadBalancerRedirectConfig) ApplicationListener {
	var returns ApplicationListener

	_jsii_.Invoke(
		a,
		"addRedirect",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Add a security group to this load balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) AddSecurityGroup(securityGroup awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		a,
		"addSecurityGroup",
		[]interface{}{securityGroup},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (a *jsiiProxy_ApplicationLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (a *jsiiProxy_ApplicationLoadBalancer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (a *jsiiProxy_ApplicationLoadBalancer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Enable access logging for this load balancer.
//
// A region must be specified on the stack containing the load balancer; you cannot enable logging on
// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func (a *jsiiProxy_ApplicationLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	_jsii_.InvokeVoid(
		a,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

// Return the given named metric for this Application Load Balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The total number of concurrent TCP connections active from clients to the load balancer and from the load balancer to targets.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricActiveConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricActiveConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of TLS connections initiated by the client that did not establish a session with the load balancer.
//
// Possible causes include a
// mismatch of ciphers or protocols.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricClientTlsNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricClientTlsNegotiationErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of load balancer capacity units (LCU) used by your load balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricConsumedLCUs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of user authentications that could not be completed.
//
// Because an authenticate action was misconfigured, the load balancer
// couldn't establish a connection with the IdP, or the load balancer
// couldn't complete the authentication flow due to an internal error.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricElbAuthError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricElbAuthError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of user authentications that could not be completed because the IdP denied access to the user or an authorization code was used more than once.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricElbAuthFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricElbAuthFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The time elapsed, in milliseconds, to query the IdP for the ID token and user info.
//
// If one or more of these operations fail, this is the time to failure.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricElbAuthLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricElbAuthLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of authenticate actions that were successful.
//
// This metric is incremented at the end of the authentication workflow,
// after the load balancer has retrieved the user claims from the IdP.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricElbAuthSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricElbAuthSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of HTTP 3xx/4xx/5xx codes that originate from the load balancer.
//
// This does not include any response codes generated by the targets.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpCodeElb(code HttpCodeElb, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpCodeElb",
		[]interface{}{code, props},
		&returns,
	)

	return returns
}

// The number of HTTP 2xx/3xx/4xx/5xx response codes generated by all targets in the load balancer.
//
// This does not include any response codes generated by the load balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpCodeTarget",
		[]interface{}{code, props},
		&returns,
	)

	return returns
}

// The number of fixed-response actions that were successful.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpFixedResponseCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpFixedResponseCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of redirect actions that were successful.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpRedirectCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpRedirectCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of redirect actions that couldn't be completed because the URL in the response location header is larger than 8K.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricHttpRedirectUrlLimitExceededCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpRedirectUrlLimitExceededCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total number of bytes processed by the load balancer over IPv6.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricIpv6ProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricIpv6ProcessedBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of IPv6 requests received by the load balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricIpv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricIpv6RequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total number of new TCP connections established from clients to the load balancer and from the load balancer to targets.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricNewConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricNewConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total number of bytes processed by the load balancer over IPv4 and IPv6.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricProcessedBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of connections that were rejected because the load balancer had reached its maximum number of connections.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricRejectedConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRejectedConnectionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of requests processed over IPv4 and IPv6.
//
// This count includes only the requests with a response generated by a target of the load balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of rules processed by the load balancer given a request rate averaged over an hour.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricRuleEvaluations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRuleEvaluations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of connections that were not successfully established between the load balancer and target.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricTargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetConnectionErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The time elapsed, in seconds, after the request leaves the load balancer until a response from the target is received.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricTargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetResponseTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of TLS connections initiated by the load balancer that did not establish a session with the target.
//
// Possible causes include a mismatch of ciphers or protocols.
func (a *jsiiProxy_ApplicationLoadBalancer) MetricTargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetTLSNegotiationErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Remove an attribute from the load balancer.
func (a *jsiiProxy_ApplicationLoadBalancer) RemoveAttribute(key *string) {
	_jsii_.InvokeVoid(
		a,
		"removeAttribute",
		[]interface{}{key},
	)
}

// Set a non-standard attribute on the load balancer.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
//
func (a *jsiiProxy_ApplicationLoadBalancer) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		a,
		"setAttribute",
		[]interface{}{key, value},
	)
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApplicationLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to reference an existing load balancer.
//
// TODO: EXAMPLE
//
type ApplicationLoadBalancerAttributes struct {
	// ARN of the load balancer.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// ID of the load balancer's security group.
	SecurityGroupId *string `json:"securityGroupId" yaml:"securityGroupId"`
	// The canonical hosted zone ID of this load balancer.
	LoadBalancerCanonicalHostedZoneId *string `json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// The DNS name of this load balancer.
	LoadBalancerDnsName *string `json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// Whether the security group allows all outbound traffic or not.
	//
	// Unless set to `false`, no egress rules will be added to the security group.
	SecurityGroupAllowsAllOutbound *bool `json:"securityGroupAllowsAllOutbound" yaml:"securityGroupAllowsAllOutbound"`
	// The VPC this load balancer has been created in, if available.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Options for looking up an ApplicationLoadBalancer.
//
// TODO: EXAMPLE
//
type ApplicationLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *map[string]*string `json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Properties for defining an Application Load Balancer.
//
// TODO: EXAMPLE
//
type ApplicationLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// Whether the load balancer has an internet-routable address.
	InternetFacing *bool `json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	LoadBalancerName *string `json:"loadBalancerName" yaml:"loadBalancerName"`
	// Which subnets place the load balancer in.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// Indicates whether HTTP/2 is enabled.
	Http2Enabled *bool `json:"http2Enabled" yaml:"http2Enabled"`
	// The load balancer idle timeout, in seconds.
	IdleTimeout awscdk.Duration `json:"idleTimeout" yaml:"idleTimeout"`
	// The type of IP addresses to use.
	//
	// Only applies to application load balancers.
	IpAddressType IpAddressType `json:"ipAddressType" yaml:"ipAddressType"`
	// Security group to associate with this load balancer.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
}

// Properties for a redirection config.
//
// TODO: EXAMPLE
//
type ApplicationLoadBalancerRedirectConfig struct {
	// Allow anyone to connect to this listener.
	//
	// If this is specified, the listener will be opened up to anyone who can reach it.
	// For internal load balancers this is anyone in the same VPC. For public load
	// balancers, this is anyone on the internet.
	//
	// If you want to be more selective about who can access this load
	// balancer, set this to `false` and use the listener's `connections`
	// object to selectively grant access to the listener.
	Open *bool `json:"open" yaml:"open"`
	// The port number to listen to.
	SourcePort *float64 `json:"sourcePort" yaml:"sourcePort"`
	// The protocol of the listener being created.
	SourceProtocol ApplicationProtocol `json:"sourceProtocol" yaml:"sourceProtocol"`
	// The port number to redirect to.
	TargetPort *float64 `json:"targetPort" yaml:"targetPort"`
	// The protocol of the redirection target.
	TargetProtocol ApplicationProtocol `json:"targetProtocol" yaml:"targetProtocol"`
}

// Load balancing protocol for application load balancers.
//
// TODO: EXAMPLE
//
type ApplicationProtocol string

const (
	ApplicationProtocol_HTTP ApplicationProtocol = "HTTP"
	ApplicationProtocol_HTTPS ApplicationProtocol = "HTTPS"
)

// Load balancing protocol version for application load balancers.
//
// TODO: EXAMPLE
//
type ApplicationProtocolVersion string

const (
	ApplicationProtocolVersion_GRPC ApplicationProtocolVersion = "GRPC"
	ApplicationProtocolVersion_HTTP1 ApplicationProtocolVersion = "HTTP1"
	ApplicationProtocolVersion_HTTP2 ApplicationProtocolVersion = "HTTP2"
)

// Define an Application Target Group.
//
// TODO: EXAMPLE
//
type ApplicationTargetGroup interface {
	TargetGroupBase
	IApplicationTargetGroup
	DefaultPort() *float64
	FirstLoadBalancerFullName() *string
	HealthCheck() *HealthCheck
	SetHealthCheck(val *HealthCheck)
	LoadBalancerArns() *string
	LoadBalancerAttached() constructs.IDependable
	LoadBalancerAttachedDependencies() constructs.DependencyGroup
	Node() constructs.Node
	TargetGroupArn() *string
	TargetGroupFullName() *string
	TargetGroupLoadBalancerArns() *[]*string
	TargetGroupName() *string
	TargetType() TargetType
	SetTargetType(val TargetType)
	AddLoadBalancerTarget(props *LoadBalancerTargetProps)
	AddTarget(targets ...IApplicationLoadBalancerTarget)
	ConfigureHealthCheck(healthCheck *HealthCheck)
	EnableCookieStickiness(duration awscdk.Duration, cookieName *string)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIpv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRequestCountPerTarget(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricUnhealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
	RegisterListener(listener IApplicationListener, associatingConstruct constructs.IConstruct)
	SetAttribute(key *string, value *string)
	ToString() *string
	ValidateTargetGroup() *[]*string
}

// The jsii proxy struct for ApplicationTargetGroup
type jsiiProxy_ApplicationTargetGroup struct {
	jsiiProxy_TargetGroupBase
	jsiiProxy_IApplicationTargetGroup
}

func (j *jsiiProxy_ApplicationTargetGroup) DefaultPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) FirstLoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstLoadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) HealthCheck() *HealthCheck {
	var returns *HealthCheck
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) LoadBalancerArns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) LoadBalancerAttached() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) LoadBalancerAttachedDependencies() constructs.DependencyGroup {
	var returns constructs.DependencyGroup
	_jsii_.Get(
		j,
		"loadBalancerAttachedDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) TargetGroupFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) TargetGroupLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationTargetGroup) TargetType() TargetType {
	var returns TargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


func NewApplicationTargetGroup(scope constructs.Construct, id *string, props *ApplicationTargetGroupProps) ApplicationTargetGroup {
	_init_.Initialize()

	j := jsiiProxy_ApplicationTargetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationTargetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationTargetGroup_Override(a ApplicationTargetGroup, scope constructs.Construct, id *string, props *ApplicationTargetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationTargetGroup",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationTargetGroup) SetHealthCheck(val *HealthCheck) {
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_ApplicationTargetGroup) SetTargetType(val TargetType) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

// Import an existing target group.
func ApplicationTargetGroup_FromTargetGroupAttributes(scope constructs.Construct, id *string, attrs *TargetGroupAttributes) IApplicationTargetGroup {
	_init_.Initialize()

	var returns IApplicationTargetGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationTargetGroup",
		"fromTargetGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApplicationTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Register the given load balancing target as part of this group.
func (a *jsiiProxy_ApplicationTargetGroup) AddLoadBalancerTarget(props *LoadBalancerTargetProps) {
	_jsii_.InvokeVoid(
		a,
		"addLoadBalancerTarget",
		[]interface{}{props},
	)
}

// Add a load balancing target to this target group.
func (a *jsiiProxy_ApplicationTargetGroup) AddTarget(targets ...IApplicationLoadBalancerTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addTarget",
		args,
	)
}

// Set/replace the target group's health check.
func (a *jsiiProxy_ApplicationTargetGroup) ConfigureHealthCheck(healthCheck *HealthCheck) {
	_jsii_.InvokeVoid(
		a,
		"configureHealthCheck",
		[]interface{}{healthCheck},
	)
}

// Enable sticky routing via a cookie to members of this target group.
//
// Note: If the `cookieName` parameter is set, application-based stickiness will be applied,
// otherwise it defaults to duration-based stickiness attributes (`lb_cookie`).
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/sticky-sessions.html
//
func (a *jsiiProxy_ApplicationTargetGroup) EnableCookieStickiness(duration awscdk.Duration, cookieName *string) {
	_jsii_.InvokeVoid(
		a,
		"enableCookieStickiness",
		[]interface{}{duration, cookieName},
	)
}

// Return the given named metric for this Application Load Balancer Target Group.
//
// Returns the metric for this target group from the point of view of the first
// load balancer load balancing to it. If you have multiple load balancers load
// sending traffic to the same target group, you will have to override the dimensions
// on this metric.
func (a *jsiiProxy_ApplicationTargetGroup) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The number of healthy hosts in the target group.
func (a *jsiiProxy_ApplicationTargetGroup) MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHealthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of HTTP 2xx/3xx/4xx/5xx response codes generated by all targets in this target group.
//
// This does not include any response codes generated by the load balancer.
func (a *jsiiProxy_ApplicationTargetGroup) MetricHttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHttpCodeTarget",
		[]interface{}{code, props},
		&returns,
	)

	return returns
}

// The number of IPv6 requests received by the target group.
func (a *jsiiProxy_ApplicationTargetGroup) MetricIpv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricIpv6RequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of requests processed over IPv4 and IPv6.
//
// This count includes only the requests with a response generated by a target of the load balancer.
func (a *jsiiProxy_ApplicationTargetGroup) MetricRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The average number of requests received by each target in a target group.
//
// The only valid statistic is Sum. Note that this represents the average not the sum.
func (a *jsiiProxy_ApplicationTargetGroup) MetricRequestCountPerTarget(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRequestCountPerTarget",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of connections that were not successfully established between the load balancer and target.
func (a *jsiiProxy_ApplicationTargetGroup) MetricTargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetConnectionErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The time elapsed, in seconds, after the request leaves the load balancer until a response from the target is received.
func (a *jsiiProxy_ApplicationTargetGroup) MetricTargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetResponseTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of TLS connections initiated by the load balancer that did not establish a session with the target.
//
// Possible causes include a mismatch of ciphers or protocols.
func (a *jsiiProxy_ApplicationTargetGroup) MetricTargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTargetTLSNegotiationErrorCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of unhealthy hosts in the target group.
func (a *jsiiProxy_ApplicationTargetGroup) MetricUnhealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricUnhealthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register a connectable as a member of this target group.
//
// Don't call this directly. It will be called by load balancing targets.
func (a *jsiiProxy_ApplicationTargetGroup) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	_jsii_.InvokeVoid(
		a,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

// Register a listener that is load balancing to this target group.
//
// Don't call this directly. It will be called by listeners.
func (a *jsiiProxy_ApplicationTargetGroup) RegisterListener(listener IApplicationListener, associatingConstruct constructs.IConstruct) {
	_jsii_.InvokeVoid(
		a,
		"registerListener",
		[]interface{}{listener, associatingConstruct},
	)
}

// Set a non-standard attribute on the target group.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
//
func (a *jsiiProxy_ApplicationTargetGroup) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		a,
		"setAttribute",
		[]interface{}{key, value},
	)
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApplicationTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationTargetGroup) ValidateTargetGroup() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validateTargetGroup",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for defining an Application Target Group.
//
// TODO: EXAMPLE
//
type ApplicationTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	HealthCheck *HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	TargetType TargetType `json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The load balancing algorithm to select targets for routing requests.
	LoadBalancingAlgorithmType TargetGroupLoadBalancingAlgorithmType `json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// The port on which the listener listens for requests.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `json:"protocol" yaml:"protocol"`
	// The protocol version to use.
	ProtocolVersion ApplicationProtocolVersion `json:"protocolVersion" yaml:"protocolVersion"`
	// The time period during which the load balancer sends a newly registered target a linearly increasing share of the traffic to the target group.
	//
	// The range is 30-900 seconds (15 minutes).
	SlowStart awscdk.Duration `json:"slowStart" yaml:"slowStart"`
	// The stickiness cookie expiration period.
	//
	// Setting this value enables load balancer stickiness.
	//
	// After this period, the cookie is considered stale. The minimum value is
	// 1 second and the maximum value is 7 days (604800 seconds).
	StickinessCookieDuration awscdk.Duration `json:"stickinessCookieDuration" yaml:"stickinessCookieDuration"`
	// The name of an application-based stickiness cookie.
	//
	// Names that start with the following prefixes are not allowed: AWSALB, AWSALBAPP,
	// and AWSALBTG; they're reserved for use by the load balancer.
	//
	// Note: `stickinessCookieName` parameter depends on the presence of `stickinessCookieDuration` parameter.
	// If `stickinessCookieDuration` is not set, `stickinessCookieName` will be omitted.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/sticky-sessions.html
	//
	StickinessCookieName *string `json:"stickinessCookieName" yaml:"stickinessCookieName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]IApplicationLoadBalancerTarget `json:"targets" yaml:"targets"`
}

// Options for `ListenerAction.authenciateOidc()`.
//
// TODO: EXAMPLE
//
type AuthenticateOidcOptions struct {
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	AuthorizationEndpoint *string `json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret.
	ClientSecret awscdk.SecretValue `json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	Issuer *string `json:"issuer" yaml:"issuer"`
	// What action to execute next.
	Next ListenerAction `json:"next" yaml:"next"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	TokenEndpoint *string `json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	UserInfoEndpoint *string `json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams *map[string]*string `json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated.
	OnUnauthenticatedRequest UnauthenticatedAction `json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP.
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	SessionCookieName *string `json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session.
	SessionTimeout awscdk.Duration `json:"sessionTimeout" yaml:"sessionTimeout"`
}

// Basic properties for an ApplicationListener.
//
// TODO: EXAMPLE
//
type BaseApplicationListenerProps struct {
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses.
	//
	// See the `ListenerAction` class for all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction ListenerAction `json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]IApplicationTargetGroup `json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Allow anyone to connect to the load balancer on the listener port.
	//
	// If this is specified, the load balancer will be opened up to anyone who can reach it.
	// For internal load balancers this is anyone in the same VPC. For public load
	// balancers, this is anyone on the internet.
	//
	// If you want to be more selective about who can access this load
	// balancer, set this to `false` and use the listener's `connections`
	// object to selectively grant access to the load balancer on the listener port.
	Open *bool `json:"open" yaml:"open"`
	// The port on which the listener listens for requests.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `json:"protocol" yaml:"protocol"`
	// The security policy that defines which ciphers and protocols are supported.
	SslPolicy SslPolicy `json:"sslPolicy" yaml:"sslPolicy"`
}

// Basic properties for defining a rule on a listener.
//
// TODO: EXAMPLE
//
type BaseApplicationListenerRuleProps struct {
	// Priority of the rule.
	//
	// The rule with the lowest priority will be used for every request.
	//
	// Priorities must be unique.
	Priority *float64 `json:"priority" yaml:"priority"`
	// Action to perform when requests are received.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	Action ListenerAction `json:"action" yaml:"action"`
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `json:"conditions" yaml:"conditions"`
	// Target groups to forward requests to.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	//
	// Implies a `forward` action.
	TargetGroups *[]IApplicationTargetGroup `json:"targetGroups" yaml:"targetGroups"`
}

// Base class for listeners.
type BaseListener interface {
	awscdk.Resource
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	ValidateListener() *[]*string
}

// The jsii proxy struct for BaseListener
type jsiiProxy_BaseListener struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_BaseListener) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseListener) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBaseListener_Override(b BaseListener, scope constructs.Construct, id *string, additionalProps interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseListener",
		[]interface{}{scope, id, additionalProps},
		b,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func BaseListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BaseListener_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseListener",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (b *jsiiProxy_BaseListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BaseListener) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (b *jsiiProxy_BaseListener) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (b *jsiiProxy_BaseListener) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (b *jsiiProxy_BaseListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate this listener.
func (b *jsiiProxy_BaseListener) ValidateListener() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validateListener",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for listener lookup.
//
// TODO: EXAMPLE
//
type BaseListenerLookupOptions struct {
	// Filter listeners by listener port.
	ListenerPort *float64 `json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	LoadBalancerTags *map[string]*string `json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Base class for both Application and Network Load Balancers.
type BaseLoadBalancer interface {
	awscdk.Resource
	Env() *awscdk.ResourceEnvironment
	LoadBalancerArn() *string
	LoadBalancerCanonicalHostedZoneId() *string
	LoadBalancerDnsName() *string
	LoadBalancerFullName() *string
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	RemoveAttribute(key *string)
	SetAttribute(key *string, value *string)
	ToString() *string
}

// The jsii proxy struct for BaseLoadBalancer
type jsiiProxy_BaseLoadBalancer struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_BaseLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewBaseLoadBalancer_Override(b BaseLoadBalancer, scope constructs.Construct, id *string, baseProps *BaseLoadBalancerProps, additionalProps interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		[]interface{}{scope, id, baseProps, additionalProps},
		b,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func BaseLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BaseLoadBalancer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (b *jsiiProxy_BaseLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BaseLoadBalancer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (b *jsiiProxy_BaseLoadBalancer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (b *jsiiProxy_BaseLoadBalancer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Enable access logging for this load balancer.
//
// A region must be specified on the stack containing the load balancer; you cannot enable logging on
// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func (b *jsiiProxy_BaseLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	_jsii_.InvokeVoid(
		b,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

// Remove an attribute from the load balancer.
func (b *jsiiProxy_BaseLoadBalancer) RemoveAttribute(key *string) {
	_jsii_.InvokeVoid(
		b,
		"removeAttribute",
		[]interface{}{key},
	)
}

// Set a non-standard attribute on the load balancer.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
//
func (b *jsiiProxy_BaseLoadBalancer) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		b,
		"setAttribute",
		[]interface{}{key, value},
	)
}

// Returns a string representation of this construct.
func (b *jsiiProxy_BaseLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for looking up load balancers.
//
// TODO: EXAMPLE
//
type BaseLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *map[string]*string `json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Shared properties of both Application and Network Load Balancers.
//
// TODO: EXAMPLE
//
type BaseLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// Whether the load balancer has an internet-routable address.
	InternetFacing *bool `json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	LoadBalancerName *string `json:"loadBalancerName" yaml:"loadBalancerName"`
	// Which subnets place the load balancer in.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Basic properties for a Network Listener.
//
// TODO: EXAMPLE
//
type BaseNetworkListenerProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `json:"port" yaml:"port"`
	// Application-Layer Protocol Negotiation (ALPN) is a TLS extension that is sent on the initial TLS handshake hello messages.
	//
	// ALPN enables the application layer to negotiate which protocols should be used over a secure connection, such as HTTP/1 and HTTP/2.
	//
	// Can only be specified together with Protocol TLS.
	AlpnPolicy AlpnPolicy `json:"alpnPolicy" yaml:"alpnPolicy"`
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default Action of the load balancer,
	// including weighted forwarding. See the `NetworkListenerAction` class for
	// all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction NetworkListenerAction `json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]INetworkTargetGroup `json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Protocol for listener, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// SSL Policy.
	SslPolicy SslPolicy `json:"sslPolicy" yaml:"sslPolicy"`
}

// Basic properties of both Application and Network Target Groups.
//
// TODO: EXAMPLE
//
type BaseTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	HealthCheck *HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	TargetType TargetType `json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::Listener`.
//
// Specifies a listener for an Application Load Balancer or Network Load Balancer.
//
// TODO: EXAMPLE
//
type CfnListener interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AlpnPolicy() *[]*string
	SetAlpnPolicy(val *[]*string)
	AttrListenerArn() *string
	Certificates() interface{}
	SetCertificates(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DefaultActions() interface{}
	SetDefaultActions(val interface{})
	LoadBalancerArn() *string
	SetLoadBalancerArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	Protocol() *string
	SetProtocol(val *string)
	Ref() *string
	SslPolicy() *string
	SetSslPolicy(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnListener
type jsiiProxy_CfnListener struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnListener) AlpnPolicy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alpnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) AttrListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrListenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Certificates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) DefaultActions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) SslPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ElasticLoadBalancingV2::Listener`.
func NewCfnListener(scope constructs.Construct, id *string, props *CfnListenerProps) CfnListener {
	_init_.Initialize()

	j := jsiiProxy_CfnListener{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ElasticLoadBalancingV2::Listener`.
func NewCfnListener_Override(c CfnListener, scope constructs.Construct, id *string, props *CfnListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnListener) SetAlpnPolicy(val *[]*string) {
	_jsii_.Set(
		j,
		"alpnPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetCertificates(val interface{}) {
	_jsii_.Set(
		j,
		"certificates",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetDefaultActions(val interface{}) {
	_jsii_.Set(
		j,
		"defaultActions",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetLoadBalancerArn(val *string) {
	_jsii_.Set(
		j,
		"loadBalancerArn",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetSslPolicy(val *string) {
	_jsii_.Set(
		j,
		"sslPolicy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnListener_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnListener_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnListener_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListener",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnListener) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnListener) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnListener) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnListener) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnListener) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnListener) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnListener) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnListener) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnListener) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnListener) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnListener) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnListener) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnListener) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an action for a listener rule.
//
// TODO: EXAMPLE
//
type CfnListener_ActionProperty struct {
	// The type of action.
	Type *string `json:"type" yaml:"type"`
	// [HTTPS listeners] Information for using Amazon Cognito to authenticate users.
	//
	// Specify only when `Type` is `authenticate-cognito` .
	AuthenticateCognitoConfig interface{} `json:"authenticateCognitoConfig" yaml:"authenticateCognitoConfig"`
	// [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC).
	//
	// Specify only when `Type` is `authenticate-oidc` .
	AuthenticateOidcConfig interface{} `json:"authenticateOidcConfig" yaml:"authenticateOidcConfig"`
	// [Application Load Balancer] Information for creating an action that returns a custom HTTP response.
	//
	// Specify only when `Type` is `fixed-response` .
	FixedResponseConfig interface{} `json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Information for creating an action that distributes requests among one or more target groups.
	//
	// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
	ForwardConfig interface{} `json:"forwardConfig" yaml:"forwardConfig"`
	// The order for the action.
	//
	// This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
	Order *float64 `json:"order" yaml:"order"`
	// [Application Load Balancer] Information for creating a redirect action.
	//
	// Specify only when `Type` is `redirect` .
	RedirectConfig interface{} `json:"redirectConfig" yaml:"redirectConfig"`
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Specify only when `Type` is `forward` and you want to route to a single target group. To route to one or more target groups, use `ForwardConfig` instead.
	TargetGroupArn *string `json:"targetGroupArn" yaml:"targetGroupArn"`
}

// Specifies information required when integrating with Amazon Cognito to authenticate users.
//
// TODO: EXAMPLE
//
type CfnListener_AuthenticateCognitoConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
	UserPoolArn *string `json:"userPoolArn" yaml:"userPoolArn"`
	// The ID of the Amazon Cognito user pool client.
	UserPoolClientId *string `json:"userPoolClientId" yaml:"userPoolClientId"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	UserPoolDomain *string `json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *string `json:"sessionTimeout" yaml:"sessionTimeout"`
}

// Specifies information required using an identity provide (IdP) that is compliant with OpenID Connect (OIDC) to authenticate users.
//
// TODO: EXAMPLE
//
type CfnListener_AuthenticateOidcConfigProperty struct {
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	AuthorizationEndpoint *string `json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set `UseExistingClientSecret` to true.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	Issuer *string `json:"issuer" yaml:"issuer"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	TokenEndpoint *string `json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	UserInfoEndpoint *string `json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *string `json:"sessionTimeout" yaml:"sessionTimeout"`
}

// Specifies an SSL server certificate to use as the default certificate for a secure listener.
//
// TODO: EXAMPLE
//
type CfnListener_CertificateProperty struct {
	// The Amazon Resource Name (ARN) of the certificate.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
}

// Specifies information required when returning a custom HTTP response.
//
// TODO: EXAMPLE
//
type CfnListener_FixedResponseConfigProperty struct {
	// The HTTP response code (2XX, 4XX, or 5XX).
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
	// The content type.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json
	ContentType *string `json:"contentType" yaml:"contentType"`
	// The message.
	MessageBody *string `json:"messageBody" yaml:"messageBody"`
}

// Information for creating an action that distributes requests among one or more target groups.
//
// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
//
// TODO: EXAMPLE
//
type CfnListener_ForwardConfigProperty struct {
	// Information about how traffic will be distributed between multiple target groups in a forward rule.
	TargetGroups interface{} `json:"targetGroups" yaml:"targetGroups"`
	// Information about the target group stickiness for a rule.
	TargetGroupStickinessConfig interface{} `json:"targetGroupStickinessConfig" yaml:"targetGroupStickinessConfig"`
}

// Information about a redirect action.
//
// A URI consists of the following components: protocol://hostname:port/path?query. You must modify at least one of the following components to avoid a redirect loop: protocol, hostname, port, or path. Any components that you do not modify retain their original values.
//
// You can reuse URI components using the following reserved keywords:
//
// - #{protocol}
// - #{host}
// - #{port}
// - #{path} (the leading "/" is removed)
// - #{query}
//
// For example, you can change the path to "/new/#{path}", the hostname to "example.#{host}", or the query to "#{query}&value=xyz".
//
// TODO: EXAMPLE
//
type CfnListener_RedirectConfigProperty struct {
	// The HTTP redirect code.
	//
	// The redirect is either permanent (HTTP 301) or temporary (HTTP 302).
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	Host *string `json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.
	Path *string `json:"path" yaml:"path"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	Port *string `json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added. You can specify any of the reserved keywords.
	Query *string `json:"query" yaml:"query"`
}

// Information about the target group stickiness for a rule.
//
// TODO: EXAMPLE
//
type CfnListener_TargetGroupStickinessConfigProperty struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target group.
	//
	// The range is 1-604800 seconds (7 days).
	DurationSeconds *float64 `json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether target group stickiness is enabled.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Information about how traffic will be distributed between multiple target groups in a forward rule.
//
// TODO: EXAMPLE
//
type CfnListener_TargetGroupTupleProperty struct {
	// The Amazon Resource Name (ARN) of the target group.
	TargetGroupArn *string `json:"targetGroupArn" yaml:"targetGroupArn"`
	// The weight.
	//
	// The range is 0 to 999.
	Weight *float64 `json:"weight" yaml:"weight"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::ListenerCertificate`.
//
// Specifies an SSL server certificate to add to the certificate list for an HTTPS or TLS listener.
//
// TODO: EXAMPLE
//
type CfnListenerCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Certificates() interface{}
	SetCertificates(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	ListenerArn() *string
	SetListenerArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnListenerCertificate
type jsiiProxy_CfnListenerCertificate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnListenerCertificate) Certificates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerCertificate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ElasticLoadBalancingV2::ListenerCertificate`.
func NewCfnListenerCertificate(scope constructs.Construct, id *string, props *CfnListenerCertificateProps) CfnListenerCertificate {
	_init_.Initialize()

	j := jsiiProxy_CfnListenerCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ElasticLoadBalancingV2::ListenerCertificate`.
func NewCfnListenerCertificate_Override(c CfnListenerCertificate, scope constructs.Construct, id *string, props *CfnListenerCertificateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnListenerCertificate) SetCertificates(val interface{}) {
	_jsii_.Set(
		j,
		"certificates",
		val,
	)
}

func (j *jsiiProxy_CfnListenerCertificate) SetListenerArn(val *string) {
	_jsii_.Set(
		j,
		"listenerArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnListenerCertificate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnListenerCertificate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificate",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnListenerCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnListenerCertificate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerCertificate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnListenerCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnListenerCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnListenerCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnListenerCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnListenerCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnListenerCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnListenerCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnListenerCertificate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnListenerCertificate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnListenerCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnListenerCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnListenerCertificate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnListenerCertificate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnListenerCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnListenerCertificate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an SSL server certificate for the certificate list of a secure listener.
//
// TODO: EXAMPLE
//
type CfnListenerCertificate_CertificateProperty struct {
	// The Amazon Resource Name (ARN) of the certificate.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
}

// Properties for defining a `CfnListenerCertificate`.
//
// TODO: EXAMPLE
//
type CfnListenerCertificateProps struct {
	// The certificate.
	//
	// You can specify one certificate per resource.
	Certificates interface{} `json:"certificates" yaml:"certificates"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `json:"listenerArn" yaml:"listenerArn"`
}

// Properties for defining a `CfnListener`.
//
// TODO: EXAMPLE
//
type CfnListenerProps struct {
	// The actions for the default rule. You cannot define a condition for a default rule.
	//
	// To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html) .
	DefaultActions interface{} `json:"defaultActions" yaml:"defaultActions"`
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
	AlpnPolicy *[]*string `json:"alpnPolicy" yaml:"alpnPolicy"`
	// The default SSL server certificate for a secure listener.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	//
	// To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html) .
	Certificates interface{} `json:"certificates" yaml:"certificates"`
	// The port on which the load balancer is listening.
	//
	// You cannot specify a port for a Gateway Load Balancer.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol for connections from clients to the load balancer.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
	//
	// For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide* .
	SslPolicy *string `json:"sslPolicy" yaml:"sslPolicy"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::ListenerRule`.
//
// Specifies a listener rule. The listener must be associated with an Application Load Balancer. Each rule consists of a priority, one or more actions, and one or more conditions.
//
// TODO: EXAMPLE
//
type CfnListenerRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Actions() interface{}
	SetActions(val interface{})
	AttrIsDefault() awscdk.IResolvable
	AttrRuleArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Conditions() interface{}
	SetConditions(val interface{})
	CreationStack() *[]*string
	ListenerArn() *string
	SetListenerArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnListenerRule
type jsiiProxy_CfnListenerRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnListenerRule) Actions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) AttrIsDefault() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) AttrRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) Conditions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListenerRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ElasticLoadBalancingV2::ListenerRule`.
func NewCfnListenerRule(scope constructs.Construct, id *string, props *CfnListenerRuleProps) CfnListenerRule {
	_init_.Initialize()

	j := jsiiProxy_CfnListenerRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ElasticLoadBalancingV2::ListenerRule`.
func NewCfnListenerRule_Override(c CfnListenerRule, scope constructs.Construct, id *string, props *CfnListenerRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnListenerRule) SetActions(val interface{}) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CfnListenerRule) SetConditions(val interface{}) {
	_jsii_.Set(
		j,
		"conditions",
		val,
	)
}

func (j *jsiiProxy_CfnListenerRule) SetListenerArn(val *string) {
	_jsii_.Set(
		j,
		"listenerArn",
		val,
	)
}

func (j *jsiiProxy_CfnListenerRule) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnListenerRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnListenerRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnListenerRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnListenerRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnListenerRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnListenerRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnListenerRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnListenerRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnListenerRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnListenerRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnListenerRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnListenerRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnListenerRule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnListenerRule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnListenerRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnListenerRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnListenerRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnListenerRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnListenerRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnListenerRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an action for a listener rule.
//
// TODO: EXAMPLE
//
type CfnListenerRule_ActionProperty struct {
	// The type of action.
	Type *string `json:"type" yaml:"type"`
	// [HTTPS listeners] Information for using Amazon Cognito to authenticate users.
	//
	// Specify only when `Type` is `authenticate-cognito` .
	AuthenticateCognitoConfig interface{} `json:"authenticateCognitoConfig" yaml:"authenticateCognitoConfig"`
	// [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC).
	//
	// Specify only when `Type` is `authenticate-oidc` .
	AuthenticateOidcConfig interface{} `json:"authenticateOidcConfig" yaml:"authenticateOidcConfig"`
	// [Application Load Balancer] Information for creating an action that returns a custom HTTP response.
	//
	// Specify only when `Type` is `fixed-response` .
	FixedResponseConfig interface{} `json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Information for creating an action that distributes requests among one or more target groups.
	//
	// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
	ForwardConfig interface{} `json:"forwardConfig" yaml:"forwardConfig"`
	// The order for the action.
	//
	// This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
	Order *float64 `json:"order" yaml:"order"`
	// [Application Load Balancer] Information for creating a redirect action.
	//
	// Specify only when `Type` is `redirect` .
	RedirectConfig interface{} `json:"redirectConfig" yaml:"redirectConfig"`
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Specify only when `Type` is `forward` and you want to route to a single target group. To route to one or more target groups, use `ForwardConfig` instead.
	TargetGroupArn *string `json:"targetGroupArn" yaml:"targetGroupArn"`
}

// Specifies information required when integrating with Amazon Cognito to authenticate users.
//
// TODO: EXAMPLE
//
type CfnListenerRule_AuthenticateCognitoConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
	UserPoolArn *string `json:"userPoolArn" yaml:"userPoolArn"`
	// The ID of the Amazon Cognito user pool client.
	UserPoolClientId *string `json:"userPoolClientId" yaml:"userPoolClientId"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	UserPoolDomain *string `json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *float64 `json:"sessionTimeout" yaml:"sessionTimeout"`
}

// Specifies information required using an identity provide (IdP) that is compliant with OpenID Connect (OIDC) to authenticate users.
//
// TODO: EXAMPLE
//
type CfnListenerRule_AuthenticateOidcConfigProperty struct {
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	AuthorizationEndpoint *string `json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set `UseExistingClientSecret` to true.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	Issuer *string `json:"issuer" yaml:"issuer"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	TokenEndpoint *string `json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	UserInfoEndpoint *string `json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *float64 `json:"sessionTimeout" yaml:"sessionTimeout"`
	// Indicates whether to use the existing client secret when modifying a rule.
	//
	// If you are creating a rule, you can omit this parameter or set it to false.
	UseExistingClientSecret interface{} `json:"useExistingClientSecret" yaml:"useExistingClientSecret"`
}

// Specifies information required when returning a custom HTTP response.
//
// TODO: EXAMPLE
//
type CfnListenerRule_FixedResponseConfigProperty struct {
	// The HTTP response code (2XX, 4XX, or 5XX).
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
	// The content type.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json
	ContentType *string `json:"contentType" yaml:"contentType"`
	// The message.
	MessageBody *string `json:"messageBody" yaml:"messageBody"`
}

// Information for creating an action that distributes requests among one or more target groups.
//
// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
//
// TODO: EXAMPLE
//
type CfnListenerRule_ForwardConfigProperty struct {
	// Information about how traffic will be distributed between multiple target groups in a forward rule.
	TargetGroups interface{} `json:"targetGroups" yaml:"targetGroups"`
	// Information about the target group stickiness for a rule.
	TargetGroupStickinessConfig interface{} `json:"targetGroupStickinessConfig" yaml:"targetGroupStickinessConfig"`
}

// Information about a host header condition.
//
// TODO: EXAMPLE
//
type CfnListenerRule_HostHeaderConfigProperty struct {
	// One or more host names.
	//
	// The maximum size of each name is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the host name.
	Values *[]*string `json:"values" yaml:"values"`
}

// Information about an HTTP header condition.
//
// There is a set of standard HTTP header fields. You can also define custom HTTP header fields.
//
// TODO: EXAMPLE
//
type CfnListenerRule_HttpHeaderConfigProperty struct {
	// The name of the HTTP header field.
	//
	// The maximum size is 40 characters. The header name is case insensitive. The allowed characters are specified by RFC 7230. Wildcards are not supported.
	HttpHeaderName *string `json:"httpHeaderName" yaml:"httpHeaderName"`
	// One or more strings to compare against the value of the HTTP header.
	//
	// The maximum size of each string is 128 characters. The comparison strings are case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If the same header appears multiple times in the request, we search them in order until a match is found.
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the value of the HTTP header. To require that all of the strings are a match, create one condition per string.
	Values *[]*string `json:"values" yaml:"values"`
}

// Information about an HTTP method condition.
//
// HTTP defines a set of request methods, also referred to as HTTP verbs. For more information, see the [HTTP Method Registry](https://docs.aws.amazon.com/https://www.iana.org/assignments/http-methods/http-methods.xhtml) . You can also define custom HTTP methods.
//
// TODO: EXAMPLE
//
type CfnListenerRule_HttpRequestMethodConfigProperty struct {
	// The name of the request method.
	//
	// The maximum size is 40 characters. The allowed characters are A-Z, hyphen (-), and underscore (_). The comparison is case sensitive. Wildcards are not supported; therefore, the method name must be an exact match.
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the HTTP request method. We recommend that you route GET and HEAD requests in the same way, because the response to a HEAD request may be cached.
	Values *[]*string `json:"values" yaml:"values"`
}

// Information about a path pattern condition.
//
// TODO: EXAMPLE
//
type CfnListenerRule_PathPatternConfigProperty struct {
	// One or more path patterns to compare against the request URL.
	//
	// The maximum size of each string is 128 characters. The comparison is case sensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If you specify multiple strings, the condition is satisfied if one of them matches the request URL. The path pattern is compared only to the path of the URL, not to its query string.
	Values *[]*string `json:"values" yaml:"values"`
}

// Information about a query string condition.
//
// The query string component of a URI starts after the first '?' character and is terminated by either a '#' character or the end of the URI. A typical query string contains key/value pairs separated by '&' characters. The allowed characters are specified by RFC 3986. Any character can be percentage encoded.
//
// TODO: EXAMPLE
//
type CfnListenerRule_QueryStringConfigProperty struct {
	// One or more key/value pairs or values to find in the query string.
	//
	// The maximum size of each string is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, you must escape these characters in `Values` using a '\' character.
	//
	// If you specify multiple key/value pairs or values, the condition is satisfied if one of them is found in the query string.
	Values interface{} `json:"values" yaml:"values"`
}

// Information about a key/value pair.
//
// TODO: EXAMPLE
//
type CfnListenerRule_QueryStringKeyValueProperty struct {
	// The key.
	//
	// You can omit the key.
	Key *string `json:"key" yaml:"key"`
	// The value.
	Value *string `json:"value" yaml:"value"`
}

// Information about a redirect action.
//
// A URI consists of the following components: protocol://hostname:port/path?query. You must modify at least one of the following components to avoid a redirect loop: protocol, hostname, port, or path. Any components that you do not modify retain their original values.
//
// You can reuse URI components using the following reserved keywords:
//
// - #{protocol}
// - #{host}
// - #{port}
// - #{path} (the leading "/" is removed)
// - #{query}
//
// For example, you can change the path to "/new/#{path}", the hostname to "example.#{host}", or the query to "#{query}&value=xyz".
//
// TODO: EXAMPLE
//
type CfnListenerRule_RedirectConfigProperty struct {
	// The HTTP redirect code.
	//
	// The redirect is either permanent (HTTP 301) or temporary (HTTP 302).
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	Host *string `json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.
	Path *string `json:"path" yaml:"path"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	Port *string `json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added. You can specify any of the reserved keywords.
	Query *string `json:"query" yaml:"query"`
}

// Specifies a condition for a listener rule.
//
// TODO: EXAMPLE
//
type CfnListenerRule_RuleConditionProperty struct {
	// The field in the HTTP request. The following are the possible values:.
	//
	// - `http-header`
	// - `http-request-method`
	// - `host-header`
	// - `path-pattern`
	// - `query-string`
	// - `source-ip`
	Field *string `json:"field" yaml:"field"`
	// Information for a host header condition.
	//
	// Specify only when `Field` is `host-header` .
	HostHeaderConfig interface{} `json:"hostHeaderConfig" yaml:"hostHeaderConfig"`
	// Information for an HTTP header condition.
	//
	// Specify only when `Field` is `http-header` .
	HttpHeaderConfig interface{} `json:"httpHeaderConfig" yaml:"httpHeaderConfig"`
	// Information for an HTTP method condition.
	//
	// Specify only when `Field` is `http-request-method` .
	HttpRequestMethodConfig interface{} `json:"httpRequestMethodConfig" yaml:"httpRequestMethodConfig"`
	// Information for a path pattern condition.
	//
	// Specify only when `Field` is `path-pattern` .
	PathPatternConfig interface{} `json:"pathPatternConfig" yaml:"pathPatternConfig"`
	// Information for a query string condition.
	//
	// Specify only when `Field` is `query-string` .
	QueryStringConfig interface{} `json:"queryStringConfig" yaml:"queryStringConfig"`
	// Information for a source IP condition.
	//
	// Specify only when `Field` is `source-ip` .
	SourceIpConfig interface{} `json:"sourceIpConfig" yaml:"sourceIpConfig"`
	// The condition value.
	//
	// Specify only when `Field` is `host-header` or `path-pattern` . Alternatively, to specify multiple host names or multiple path patterns, use `HostHeaderConfig` or `PathPatternConfig` .
	//
	// If `Field` is `host-header` and you're not using `HostHeaderConfig` , you can specify a single host name (for example, my.example.com). A host name is case insensitive, can be up to 128 characters in length, and can contain any of the following characters.
	//
	// - A-Z, a-z, 0-9
	// - - .
	// - * (matches 0 or more characters)
	// - ? (matches exactly 1 character)
	//
	// If `Field` is `path-pattern` and you're not using `PathPatternConfig` , you can specify a single path pattern (for example, /img/*). A path pattern is case-sensitive, can be up to 128 characters in length, and can contain any of the following characters.
	//
	// - A-Z, a-z, 0-9
	// - _ - . $ / ~ " ' @ : +
	// - & (using &amp;)
	// - * (matches 0 or more characters)
	// - ? (matches exactly 1 character)
	Values *[]*string `json:"values" yaml:"values"`
}

// Information about a source IP condition.
//
// You can use this condition to route based on the IP address of the source that connects to the load balancer. If a client is behind a proxy, this is the IP address of the proxy not the IP address of the client.
//
// TODO: EXAMPLE
//
type CfnListenerRule_SourceIpConfigProperty struct {
	// One or more source IP addresses, in CIDR format.
	//
	// You can use both IPv4 and IPv6 addresses. Wildcards are not supported.
	//
	// If you specify multiple addresses, the condition is satisfied if the source IP address of the request matches one of the CIDR blocks. This condition is not satisfied by the addresses in the X-Forwarded-For header.
	Values *[]*string `json:"values" yaml:"values"`
}

// Information about the target group stickiness for a rule.
//
// TODO: EXAMPLE
//
type CfnListenerRule_TargetGroupStickinessConfigProperty struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target group.
	//
	// The range is 1-604800 seconds (7 days).
	DurationSeconds *float64 `json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether target group stickiness is enabled.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Information about how traffic will be distributed between multiple target groups in a forward rule.
//
// TODO: EXAMPLE
//
type CfnListenerRule_TargetGroupTupleProperty struct {
	// The Amazon Resource Name (ARN) of the target group.
	TargetGroupArn *string `json:"targetGroupArn" yaml:"targetGroupArn"`
	// The weight.
	//
	// The range is 0 to 999.
	Weight *float64 `json:"weight" yaml:"weight"`
}

// Properties for defining a `CfnListenerRule`.
//
// TODO: EXAMPLE
//
type CfnListenerRuleProps struct {
	// The actions.
	//
	// The rule must include exactly one of the following types of actions: `forward` , `fixed-response` , or `redirect` , and it must be the last action to be performed. If the rule is for an HTTPS listener, it can also optionally include an authentication action.
	Actions interface{} `json:"actions" yaml:"actions"`
	// The conditions.
	//
	// The rule can optionally include up to one of each of the following conditions: `http-request-method` , `host-header` , `path-pattern` , and `source-ip` . A rule can also optionally include one or more of each of the following conditions: `http-header` and `query-string` .
	Conditions interface{} `json:"conditions" yaml:"conditions"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `json:"listenerArn" yaml:"listenerArn"`
	// The rule priority. A listener can't have multiple rules with the same priority.
	//
	// If you try to reorder rules by updating their priorities, do not specify a new priority if an existing rule already uses this priority, as this can cause an error. If you need to reuse a priority with a different rule, you must remove it as a priority first, and then specify it in a subsequent update.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::LoadBalancer`.
//
// Specifies an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrCanonicalHostedZoneId() *string
	AttrDnsName() *string
	AttrLoadBalancerFullName() *string
	AttrLoadBalancerName() *string
	AttrSecurityGroups() *[]*string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	IpAddressType() *string
	SetIpAddressType(val *string)
	LoadBalancerAttributes() interface{}
	SetLoadBalancerAttributes(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Scheme() *string
	SetScheme(val *string)
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	Stack() awscdk.Stack
	SubnetMappings() interface{}
	SetSubnetMappings(val interface{})
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLoadBalancer
type jsiiProxy_CfnLoadBalancer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLoadBalancer) AttrCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AttrDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AttrLoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AttrLoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AttrSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) LoadBalancerAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) SubnetMappings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ElasticLoadBalancingV2::LoadBalancer`.
func NewCfnLoadBalancer(scope constructs.Construct, id *string, props *CfnLoadBalancerProps) CfnLoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ElasticLoadBalancingV2::LoadBalancer`.
func NewCfnLoadBalancer_Override(c CfnLoadBalancer, scope constructs.Construct, id *string, props *CfnLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetLoadBalancerAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"loadBalancerAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetScheme(val *string) {
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetSubnetMappings(val interface{}) {
	_jsii_.Set(
		j,
		"subnetMappings",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLoadBalancer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLoadBalancer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoadBalancer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnLoadBalancer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLoadBalancer) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnLoadBalancer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnLoadBalancer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnLoadBalancer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLoadBalancer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLoadBalancer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnLoadBalancer) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnLoadBalancer) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnLoadBalancer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLoadBalancer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnLoadBalancer) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an attribute for an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_LoadBalancerAttributeProperty struct {
	// The name of the attribute.
	//
	// The following attribute is supported by all load balancers:
	//
	// - `deletion_protection.enabled` - Indicates whether deletion protection is enabled. The value is `true` or `false` . The default is `false` .
	//
	// The following attributes are supported by both Application Load Balancers and Network Load Balancers:
	//
	// - `access_logs.s3.enabled` - Indicates whether access logs are enabled. The value is `true` or `false` . The default is `false` .
	// - `access_logs.s3.bucket` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.
	// - `access_logs.s3.prefix` - The prefix for the location in the S3 bucket for the access logs.
	// - `ipv6.deny_all_igw_traffic` - Blocks internet gateway (IGW) access to the load balancer. It is set to `false` for internet-facing load balancers and `true` for internal load balancers, preventing unintended access to your internal load balancer through an internet gateway.
	//
	// The following attributes are supported by only Application Load Balancers:
	//
	// - `idle_timeout.timeout_seconds` - The idle timeout value, in seconds. The valid range is 1-4000 seconds. The default is 60 seconds.
	// - `routing.http.desync_mitigation_mode` - Determines how the load balancer handles requests that might pose a security risk to your application. The possible values are `monitor` , `defensive` , and `strictest` . The default is `defensive` .
	// - `routing.http.drop_invalid_header_fields.enabled` - Indicates whether HTTP headers with invalid header fields are removed by the load balancer ( `true` ) or routed to targets ( `false` ). The default is `false` .
	// - `routing.http.x_amzn_tls_version_and_cipher_suite.enabled` - Indicates whether the two headers ( `x-amzn-tls-version` and `x-amzn-tls-cipher-suite` ), which contain information about the negotiated TLS version and cipher suite, are added to the client request before sending it to the target. The `x-amzn-tls-version` header has information about the TLS protocol version negotiated with the client, and the `x-amzn-tls-cipher-suite` header has information about the cipher suite negotiated with the client. Both headers are in OpenSSL format. The possible values for the attribute are `true` and `false` . The default is `false` .
	// - `routing.http.xff_client_port.enabled` - Indicates whether the `X-Forwarded-For` header should preserve the source port that the client used to connect to the load balancer. The possible values are `true` and `false` . The default is `false` .
	// - `routing.http2.enabled` - Indicates whether HTTP/2 is enabled. The possible values are `true` and `false` . The default is `true` . Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens.
	// - `waf.fail_open.enabled` - Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. The possible values are `true` and `false` . The default is `false` .
	//
	// The following attribute is supported by Network Load Balancers and Gateway Load Balancers:
	//
	// - `load_balancing.cross_zone.enabled` - Indicates whether cross-zone load balancing is enabled. The possible values are `true` and `false` . The default is `false` .
	Key *string `json:"key" yaml:"key"`
	// The value of the attribute.
	Value *string `json:"value" yaml:"value"`
}

// Specifies a subnet for a load balancer.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_SubnetMappingProperty struct {
	// The ID of the subnet.
	SubnetId *string `json:"subnetId" yaml:"subnetId"`
	// [Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.
	AllocationId *string `json:"allocationId" yaml:"allocationId"`
	// [Network Load Balancers] The IPv6 address.
	IPv6Address *string `json:"iPv6Address" yaml:"iPv6Address"`
	// [Network Load Balancers] The private IPv4 address for an internal load balancer.
	PrivateIPv4Address *string `json:"privateIPv4Address" yaml:"privateIPv4Address"`
}

// Properties for defining a `CfnLoadBalancer`.
//
// TODO: EXAMPLE
//
type CfnLoadBalancerProps struct {
	// The IP address type.
	//
	// The possible values are `ipv4` (for IPv4 addresses) and `dualstack` (for IPv4 and IPv6 addresses). You can’t specify `dualstack` for a load balancer with a UDP or TCP_UDP listener.
	IpAddressType *string `json:"ipAddressType" yaml:"ipAddressType"`
	// The load balancer attributes.
	LoadBalancerAttributes interface{} `json:"loadBalancerAttributes" yaml:"loadBalancerAttributes"`
	// The name of the load balancer.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-".
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	Name *string `json:"name" yaml:"name"`
	// The nodes of an Internet-facing load balancer have public IP addresses.
	//
	// The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.
	//
	// The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.
	//
	// The default is an Internet-facing load balancer.
	//
	// You cannot specify a scheme for a Gateway Load Balancer.
	Scheme *string `json:"scheme" yaml:"scheme"`
	// [Application Load Balancers] The IDs of the security groups for the load balancer.
	SecurityGroups *[]*string `json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the public subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
	//
	// [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.
	//
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//
	// [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
	//
	// [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.
	//
	// [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP addresses for your subnets.
	SubnetMappings interface{} `json:"subnetMappings" yaml:"subnetMappings"`
	// The IDs of the public subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
	//
	// [Application Load Balancers] You must specify subnets from at least two Availability Zones.
	//
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//
	// [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
	//
	// [Network Load Balancers] You can specify subnets from one or more Availability Zones.
	//
	// [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// The tags to assign to the load balancer.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The type of load balancer.
	//
	// The default is `application` .
	Type *string `json:"type" yaml:"type"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::TargetGroup`.
//
// Specifies a target group for a load balancer.
//
// Before you register a Lambda function as a target, you must create a `AWS::Lambda::Permission` resource that grants the Elastic Load Balancing service principal permission to invoke the Lambda function.
//
// TODO: EXAMPLE
//
type CfnTargetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrLoadBalancerArns() *[]*string
	AttrTargetGroupFullName() *string
	AttrTargetGroupName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	HealthCheckEnabled() interface{}
	SetHealthCheckEnabled(val interface{})
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPort() *string
	SetHealthCheckPort(val *string)
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	HealthCheckTimeoutSeconds() *float64
	SetHealthCheckTimeoutSeconds(val *float64)
	HealthyThresholdCount() *float64
	SetHealthyThresholdCount(val *float64)
	IpAddressType() *string
	SetIpAddressType(val *string)
	LogicalId() *string
	Matcher() interface{}
	SetMatcher(val interface{})
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	Protocol() *string
	SetProtocol(val *string)
	ProtocolVersion() *string
	SetProtocolVersion(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TargetGroupAttributes() interface{}
	SetTargetGroupAttributes(val interface{})
	Targets() interface{}
	SetTargets(val interface{})
	TargetType() *string
	SetTargetType(val *string)
	UnhealthyThresholdCount() *float64
	SetUnhealthyThresholdCount(val *float64)
	UpdatedProperites() *map[string]interface{}
	VpcId() *string
	SetVpcId(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnTargetGroup
type jsiiProxy_CfnTargetGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTargetGroup) AttrLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) AttrTargetGroupFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTargetGroupFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) AttrTargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTargetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) HealthCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) HealthCheckIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) HealthCheckPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) HealthCheckProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) HealthCheckTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) HealthyThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Matcher() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matcher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) ProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) TargetGroupAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) UnhealthyThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetGroup) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new `AWS::ElasticLoadBalancingV2::TargetGroup`.
func NewCfnTargetGroup(scope constructs.Construct, id *string, props *CfnTargetGroupProps) CfnTargetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnTargetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ElasticLoadBalancingV2::TargetGroup`.
func NewCfnTargetGroup_Override(c CfnTargetGroup, scope constructs.Construct, id *string, props *CfnTargetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetHealthCheckEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetHealthCheckIntervalSeconds(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetHealthCheckPath(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetHealthCheckPort(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPort",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetHealthCheckProtocol(val *string) {
	_jsii_.Set(
		j,
		"healthCheckProtocol",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetHealthCheckTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetHealthyThresholdCount(val *float64) {
	_jsii_.Set(
		j,
		"healthyThresholdCount",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetMatcher(val interface{}) {
	_jsii_.Set(
		j,
		"matcher",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetProtocolVersion(val *string) {
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetTargetGroupAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"targetGroupAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetTargetType(val *string) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetUnhealthyThresholdCount(val *float64) {
	_jsii_.Set(
		j,
		"unhealthyThresholdCount",
		val,
	)
}

func (j *jsiiProxy_CfnTargetGroup) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTargetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTargetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTargetGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_elasticloadbalancingv2.CfnTargetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTargetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnTargetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnTargetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnTargetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTargetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTargetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnTargetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnTargetGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnTargetGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnTargetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTargetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTargetGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnTargetGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTargetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the HTTP codes that healthy targets must use when responding to an HTTP health check.
//
// TODO: EXAMPLE
//
type CfnTargetGroup_MatcherProperty struct {
	// You can specify values between 0 and 99.
	//
	// You can specify multiple values (for example, "0,1") or a range of values (for example, "0-5"). The default value is 12.
	GrpcCode *string `json:"grpcCode" yaml:"grpcCode"`
	// For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200.
	//
	// You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299").
	//
	// For Network Load Balancers and Gateway Load Balancers, this must be "200–399".
	//
	// Note that when using shorthand syntax, some values such as commas need to be escaped.
	HttpCode *string `json:"httpCode" yaml:"httpCode"`
}

// Specifies a target to add to a target group.
//
// TODO: EXAMPLE
//
type CfnTargetGroup_TargetDescriptionProperty struct {
	// The ID of the target.
	//
	// If the target type of the target group is `instance` , specify an instance ID. If the target type is `ip` , specify an IP address. If the target type is `lambda` , specify the ARN of the Lambda function. If the target type is `alb` , specify the ARN of the Application Load Balancer target.
	Id *string `json:"id" yaml:"id"`
	// An Availability Zone or `all` .
	//
	// This determines whether the target receives traffic from the load balancer nodes in the specified Availability Zone or from all enabled Availability Zones for the load balancer.
	//
	// This parameter is not supported if the target type of the target group is `instance` or `alb` .
	//
	// If the target type is `ip` and the IP address is in a subnet of the VPC for the target group, the Availability Zone is automatically detected and this parameter is optional. If the IP address is outside the VPC, this parameter is required.
	//
	// With an Application Load Balancer, if the target type is `ip` and the IP address is outside the VPC for the target group, the only supported value is `all` .
	//
	// If the target type is `lambda` , this parameter is optional and the only supported value is `all` .
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The port on which the target is listening.
	//
	// If the target group protocol is GENEVE, the supported port is 6081. If the target type is `alb` , the targeted Application Load Balancer must have at least one listener whose port matches the target group port. Not used if the target is a Lambda function.
	Port *float64 `json:"port" yaml:"port"`
}

// Specifies a target group attribute.
//
// TODO: EXAMPLE
//
type CfnTargetGroup_TargetGroupAttributeProperty struct {
	// The name of the attribute.
	//
	// The following attribute is supported by all load balancers:
	//
	// - `deregistration_delay.timeout_seconds` - The amount of time, in seconds, for Elastic Load Balancing to wait before changing the state of a deregistering target from `draining` to `unused` . The range is 0-3600 seconds. The default value is 300 seconds. If the target is a Lambda function, this attribute is not supported.
	//
	// The following attributes are supported by both Application Load Balancers and Network Load Balancers:
	//
	// - `stickiness.enabled` - Indicates whether sticky sessions are enabled. The value is `true` or `false` . The default is `false` .
	// - `stickiness.type` - The type of sticky sessions. The possible values are `lb_cookie` and `app_cookie` for Application Load Balancers or `source_ip` for Network Load Balancers.
	//
	// The following attributes are supported only if the load balancer is an Application Load Balancer and the target is an instance or an IP address:
	//
	// - `load_balancing.algorithm.type` - The load balancing algorithm determines how the load balancer selects targets when routing requests. The value is `round_robin` or `least_outstanding_requests` . The default is `round_robin` .
	// - `slow_start.duration_seconds` - The time period, in seconds, during which a newly registered target receives an increasing share of the traffic to the target group. After this time period ends, the target receives its full share of traffic. The range is 30-900 seconds (15 minutes). The default is 0 seconds (disabled).
	// - `stickiness.app_cookie.cookie_name` - Indicates the name of the application-based cookie. Names that start with the following prefixes are not allowed: `AWSALB` , `AWSALBAPP` , and `AWSALBTG` ; they're reserved for use by the load balancer.
	// - `stickiness.app_cookie.duration_seconds` - The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the application-based cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	// - `stickiness.lb_cookie.duration_seconds` - The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	//
	// The following attribute is supported only if the load balancer is an Application Load Balancer and the target is a Lambda function:
	//
	// - `lambda.multi_value_headers.enabled` - Indicates whether the request and response headers that are exchanged between the load balancer and the Lambda function include arrays of values or strings. The value is `true` or `false` . The default is `false` . If the value is `false` and the request contains a duplicate header field name or query parameter key, the load balancer uses the last value sent by the client.
	//
	// The following attributes are supported only by Network Load Balancers:
	//
	// - `deregistration_delay.connection_termination.enabled` - Indicates whether the load balancer terminates connections at the end of the deregistration timeout. The value is `true` or `false` . The default is `false` .
	// - `preserve_client_ip.enabled` - Indicates whether client IP preservation is enabled. The value is `true` or `false` . The default is disabled if the target group type is IP address and the target group protocol is TCP or TLS. Otherwise, the default is enabled. Client IP preservation cannot be disabled for UDP and TCP_UDP target groups.
	// - `proxy_protocol_v2.enabled` - Indicates whether Proxy Protocol version 2 is enabled. The value is `true` or `false` . The default is `false` .
	Key *string `json:"key" yaml:"key"`
	// The value of the attribute.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnTargetGroup`.
//
// TODO: EXAMPLE
//
type CfnTargetGroupProps struct {
	// Indicates whether health checks are enabled.
	//
	// If the target type is `lambda` , health checks are disabled by default but can be enabled. If the target type is `instance` , `ip` , or `alb` , health checks are always enabled and cannot be disabled.
	HealthCheckEnabled interface{} `json:"healthCheckEnabled" yaml:"healthCheckEnabled"`
	// The approximate amount of time, in seconds, between health checks of an individual target.
	//
	// If the target group protocol is HTTP or HTTPS, the default is 30 seconds. If the target group protocol is TCP, TLS, UDP, or TCP_UDP, the supported values are 10 and 30 seconds and the default is 30 seconds. If the target group protocol is GENEVE, the default is 10 seconds. If the target type is `lambda` , the default is 35 seconds.
	HealthCheckIntervalSeconds *float64 `json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	//
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /.
	//
	// [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is / AWS .ALB/healthcheck.
	HealthCheckPath *string `json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port the load balancer uses when performing health checks on targets.
	//
	// If the protocol is HTTP, HTTPS, TCP, TLS, UDP, or TCP_UDP, the default is `traffic-port` , which is the port on which each target receives traffic from the load balancer. If the protocol is GENEVE, the default is port 80.
	HealthCheckPort *string `json:"healthCheckPort" yaml:"healthCheckPort"`
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers and Gateway Load Balancers, the default is TCP. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. The GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	HealthCheckProtocol *string `json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// For target groups with a protocol of HTTP, HTTPS, or GENEVE, the default is 5 seconds. For target groups with a protocol of TCP or TLS, this value must be 6 seconds for HTTP health checks and 10 seconds for TCP and HTTPS health checks. If the target type is `lambda` , the default is 30 seconds.
	HealthCheckTimeoutSeconds *float64 `json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	//
	// For target groups with a protocol of HTTP or HTTPS, the default is 5. For target groups with a protocol of TCP, TLS, or GENEVE, the default is 3. If the target type is `lambda` , the default is 5.
	HealthyThresholdCount *float64 `json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The type of IP address used for this target group.
	//
	// The possible values are `ipv4` and `ipv6` . This is an optional parameter. If not specified, the IP address type defaults to `ipv4` .
	IpAddressType *string `json:"ipAddressType" yaml:"ipAddressType"`
	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
	Matcher interface{} `json:"matcher" yaml:"matcher"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name *string `json:"name" yaml:"name"`
	// The port on which the targets receive traffic.
	//
	// This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.
	Port *float64 `json:"port" yaml:"port"`
	// The protocol to use for routing traffic to the targets.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, or TCP_UDP. For Gateway Load Balancers, the supported protocol is GENEVE. A TCP_UDP listener must be associated with a TCP_UDP target group. If the target is a Lambda function, this parameter does not apply.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// [HTTP/HTTPS protocol] The protocol version.
	//
	// The possible values are `GRPC` , `HTTP1` , and `HTTP2` .
	ProtocolVersion *string `json:"protocolVersion" yaml:"protocolVersion"`
	// The tags.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The attributes.
	TargetGroupAttributes interface{} `json:"targetGroupAttributes" yaml:"targetGroupAttributes"`
	// The targets.
	Targets interface{} `json:"targets" yaml:"targets"`
	// The type of target that you must specify when registering targets with this target group.
	//
	// You can't specify targets for a target group using more than one target type.
	//
	// - `instance` - Register targets by instance ID. This is the default value.
	// - `ip` - Register targets by IP address. You can specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	// - `lambda` - Register a single Lambda function as a target.
	// - `alb` - Register a single Application Load Balancer as a target.
	TargetType *string `json:"targetType" yaml:"targetType"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// If the target group protocol is HTTP or HTTPS, the default is 2. If the target group protocol is TCP or TLS, this value must be the same as the healthy threshold count. If the target group protocol is GENEVE, the default is 3. If the target type is `lambda` , the default is 2.
	UnhealthyThresholdCount *float64 `json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
	// The identifier of the virtual private cloud (VPC).
	//
	// If the target is a Lambda function, this parameter does not apply. Otherwise, this parameter is required.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Options for `ListenerAction.fixedResponse()`.
//
// TODO: EXAMPLE
//
type FixedResponseOptions struct {
	// Content Type of the response.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json
	ContentType *string `json:"contentType" yaml:"contentType"`
	// The response body.
	MessageBody *string `json:"messageBody" yaml:"messageBody"`
}

// Options for `ListenerAction.forward()`.
//
// TODO: EXAMPLE
//
type ForwardOptions struct {
	// For how long clients should be directed to the same target group.
	//
	// Range between 1 second and 7 days.
	StickinessDuration awscdk.Duration `json:"stickinessDuration" yaml:"stickinessDuration"`
}

// Properties for configuring a health check.
//
// TODO: EXAMPLE
//
type HealthCheck struct {
	// Indicates whether health checks are enabled.
	//
	// If the target type is lambda,
	// health checks are disabled by default but can be enabled. If the target type
	// is instance or ip, health checks are always enabled and cannot be disabled.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// GRPC code to use when checking for a successful response from a target.
	//
	// You can specify values between 0 and 99. You can specify multiple values
	// (for example, "0,1") or a range of values (for example, "0-5").
	HealthyGrpcCodes *string `json:"healthyGrpcCodes" yaml:"healthyGrpcCodes"`
	// HTTP code to use when checking for a successful response from a target.
	//
	// For Application Load Balancers, you can specify values between 200 and
	// 499, and the default value is 200. You can specify multiple values (for
	// example, "200,202") or a range of values (for example, "200-299").
	HealthyHttpCodes *string `json:"healthyHttpCodes" yaml:"healthyHttpCodes"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	//
	// For Application Load Balancers, the default is 5. For Network Load Balancers, the default is 3.
	HealthyThresholdCount *float64 `json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The approximate number of seconds between health checks for an individual target.
	Interval awscdk.Duration `json:"interval" yaml:"interval"`
	// The ping path destination where Elastic Load Balancing sends health check requests.
	Path *string `json:"path" yaml:"path"`
	// The port that the load balancer uses when performing health checks on the targets.
	Port *string `json:"port" yaml:"port"`
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// The TCP protocol is supported for health checks only if the protocol of the target group is TCP, TLS, UDP, or TCP_UDP.
	// The TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// For Application Load Balancers, the range is 2-60 seconds and the
	// default is 5 seconds. For Network Load Balancers, this is 10 seconds for
	// TCP and HTTPS health checks and 6 seconds for HTTP health checks.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// For Application Load Balancers, the default is 2. For Network Load
	// Balancers, this value must be the same as the healthy threshold count.
	UnhealthyThresholdCount *float64 `json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
}

// Count of HTTP status originating from the load balancer.
//
// This count does not include any response codes generated by the targets.
type HttpCodeElb string

const (
	HttpCodeElb_ELB_3XX_COUNT HttpCodeElb = "ELB_3XX_COUNT"
	HttpCodeElb_ELB_4XX_COUNT HttpCodeElb = "ELB_4XX_COUNT"
	HttpCodeElb_ELB_5XX_COUNT HttpCodeElb = "ELB_5XX_COUNT"
)

// Count of HTTP status originating from the targets.
type HttpCodeTarget string

const (
	HttpCodeTarget_TARGET_2XX_COUNT HttpCodeTarget = "TARGET_2XX_COUNT"
	HttpCodeTarget_TARGET_3XX_COUNT HttpCodeTarget = "TARGET_3XX_COUNT"
	HttpCodeTarget_TARGET_4XX_COUNT HttpCodeTarget = "TARGET_4XX_COUNT"
	HttpCodeTarget_TARGET_5XX_COUNT HttpCodeTarget = "TARGET_5XX_COUNT"
)

// Properties to reference an existing listener.
type IApplicationListener interface {
	awsec2.IConnectable
	awscdk.IResource
	// Add one or more certificates to this listener.
	AddCertificates(id *string, certificates *[]IListenerCertificate)
	// Load balance incoming requests to the given target groups.
	//
	// It's possible to add conditions to the TargetGroups added in this way.
	// At least one TargetGroup must be added without conditions.
	AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps)
	// Load balance incoming requests to the given load balancing targets.
	//
	// This method implicitly creates an ApplicationTargetGroup for the targets
	// involved.
	//
	// It's possible to add conditions to the targets added in this way. At least
	// one set of targets must be added without conditions.
	//
	// Returns: The newly created target group
	AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup
	// Register that a connectable that has been added to this load balancer.
	//
	// Don't call this directly. It is called by ApplicationTargetGroup.
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
	// ARN of the listener.
	ListenerArn() *string
}

// The jsii proxy for IApplicationListener
type jsiiProxy_IApplicationListener struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApplicationListener) AddCertificates(id *string, certificates *[]IListenerCertificate) {
	_jsii_.InvokeVoid(
		i,
		"addCertificates",
		[]interface{}{id, certificates},
	)
}

func (i *jsiiProxy_IApplicationListener) AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps) {
	_jsii_.InvokeVoid(
		i,
		"addTargetGroups",
		[]interface{}{id, props},
	)
}

func (i *jsiiProxy_IApplicationListener) AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup {
	var returns ApplicationTargetGroup

	_jsii_.Invoke(
		i,
		"addTargets",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationListener) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	_jsii_.InvokeVoid(
		i,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

func (i *jsiiProxy_IApplicationListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IApplicationListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationListener) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationListener) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// An application load balancer.
type IApplicationLoadBalancer interface {
	awsec2.IConnectable
	ILoadBalancerV2
	// Add a new listener to this load balancer.
	AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener
	// The IP Address Type for this load balancer.
	IpAddressType() IpAddressType
	// A list of listeners that have been added to the load balancer.
	//
	// This list is only valid for owned constructs.
	Listeners() *[]ApplicationListener
	// The ARN of this load balancer.
	LoadBalancerArn() *string
	// The VPC this load balancer has been created in (if available).
	//
	// If this interface is the result of an import call to fromApplicationLoadBalancerAttributes,
	// the vpc attribute will be undefined unless specified in the optional properties of that method.
	Vpc() awsec2.IVpc
}

// The jsii proxy for IApplicationLoadBalancer
type jsiiProxy_IApplicationLoadBalancer struct {
	internal.Type__awsec2IConnectable
	jsiiProxy_ILoadBalancerV2
}

func (i *jsiiProxy_IApplicationLoadBalancer) AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener {
	var returns ApplicationListener

	_jsii_.Invoke(
		i,
		"addListener",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IApplicationLoadBalancer) IpAddressType() IpAddressType {
	var returns IpAddressType
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Listeners() *[]ApplicationListener {
	var returns *[]ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Interface for constructs that can be targets of an application load balancer.
type IApplicationLoadBalancerTarget interface {
	// Attach load-balanced target to a TargetGroup.
	//
	// May return JSON to directly add to the [Targets] list, or return undefined
	// if the target will register itself with the load balancer.
	AttachToApplicationTargetGroup(targetGroup IApplicationTargetGroup) *LoadBalancerTargetProps
}

// The jsii proxy for IApplicationLoadBalancerTarget
type jsiiProxy_IApplicationLoadBalancerTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IApplicationLoadBalancerTarget) AttachToApplicationTargetGroup(targetGroup IApplicationTargetGroup) *LoadBalancerTargetProps {
	var returns *LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// A Target Group for Application Load Balancers.
type IApplicationTargetGroup interface {
	ITargetGroup
	// Add a load balancing target to this target group.
	AddTarget(targets ...IApplicationLoadBalancerTarget)
	// Register a connectable as a member of this target group.
	//
	// Don't call this directly. It will be called by load balancing targets.
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	RegisterListener(listener IApplicationListener, associatingConstruct constructs.IConstruct)
}

// The jsii proxy for IApplicationTargetGroup
type jsiiProxy_IApplicationTargetGroup struct {
	jsiiProxy_ITargetGroup
}

func (i *jsiiProxy_IApplicationTargetGroup) AddTarget(targets ...IApplicationLoadBalancerTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addTarget",
		args,
	)
}

func (i *jsiiProxy_IApplicationTargetGroup) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	_jsii_.InvokeVoid(
		i,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

func (i *jsiiProxy_IApplicationTargetGroup) RegisterListener(listener IApplicationListener, associatingConstruct constructs.IConstruct) {
	_jsii_.InvokeVoid(
		i,
		"registerListener",
		[]interface{}{listener, associatingConstruct},
	)
}

// Interface for listener actions.
type IListenerAction interface {
	// Render the actions in this chain.
	RenderActions() *[]*CfnListener_ActionProperty
}

// The jsii proxy for IListenerAction
type jsiiProxy_IListenerAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IListenerAction) RenderActions() *[]*CfnListener_ActionProperty {
	var returns *[]*CfnListener_ActionProperty

	_jsii_.Invoke(
		i,
		"renderActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A certificate source for an ELBv2 listener.
type IListenerCertificate interface {
	// The ARN of the certificate to use.
	CertificateArn() *string
}

// The jsii proxy for IListenerCertificate
type jsiiProxy_IListenerCertificate struct {
	_ byte // padding
}

func (j *jsiiProxy_IListenerCertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

type ILoadBalancerV2 interface {
	awscdk.IResource
	// The canonical hosted zone ID of this load balancer.
	//
	// Example value: `Z2P70J7EXAMPLE`
	LoadBalancerCanonicalHostedZoneId() *string
	// The DNS name of this load balancer.
	//
	// Example value: `my-load-balancer-424835706.us-west-2.elb.amazonaws.com`
	LoadBalancerDnsName() *string
}

// The jsii proxy for ILoadBalancerV2
type jsiiProxy_ILoadBalancerV2 struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILoadBalancerV2) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILoadBalancerV2) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

// Properties to reference an existing listener.
type INetworkListener interface {
	awscdk.IResource
	// ARN of the listener.
	ListenerArn() *string
}

// The jsii proxy for INetworkListener
type jsiiProxy_INetworkListener struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_INetworkListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

// A network load balancer.
type INetworkLoadBalancer interface {
	ILoadBalancerV2
	awsec2.IVpcEndpointServiceLoadBalancer
	// Add a listener to this load balancer.
	//
	// Returns: The newly created listener
	AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener
	// The VPC this load balancer has been created in (if available).
	Vpc() awsec2.IVpc
}

// The jsii proxy for INetworkLoadBalancer
type jsiiProxy_INetworkLoadBalancer struct {
	jsiiProxy_ILoadBalancerV2
	internal.Type__awsec2IVpcEndpointServiceLoadBalancer
}

func (i *jsiiProxy_INetworkLoadBalancer) AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener {
	var returns NetworkListener

	_jsii_.Invoke(
		i,
		"addListener",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INetworkLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_INetworkLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Interface for constructs that can be targets of an network load balancer.
type INetworkLoadBalancerTarget interface {
	// Attach load-balanced target to a TargetGroup.
	//
	// May return JSON to directly add to the [Targets] list, or return undefined
	// if the target will register itself with the load balancer.
	AttachToNetworkTargetGroup(targetGroup INetworkTargetGroup) *LoadBalancerTargetProps
}

// The jsii proxy for INetworkLoadBalancerTarget
type jsiiProxy_INetworkLoadBalancerTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_INetworkLoadBalancerTarget) AttachToNetworkTargetGroup(targetGroup INetworkTargetGroup) *LoadBalancerTargetProps {
	var returns *LoadBalancerTargetProps

	_jsii_.Invoke(
		i,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

// A network target group.
type INetworkTargetGroup interface {
	ITargetGroup
	// Add a load balancing target to this target group.
	AddTarget(targets ...INetworkLoadBalancerTarget)
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	RegisterListener(listener INetworkListener)
}

// The jsii proxy for INetworkTargetGroup
type jsiiProxy_INetworkTargetGroup struct {
	jsiiProxy_ITargetGroup
}

func (i *jsiiProxy_INetworkTargetGroup) AddTarget(targets ...INetworkLoadBalancerTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addTarget",
		args,
	)
}

func (i *jsiiProxy_INetworkTargetGroup) RegisterListener(listener INetworkListener) {
	_jsii_.InvokeVoid(
		i,
		"registerListener",
		[]interface{}{listener},
	)
}

// A target group.
type ITargetGroup interface {
	constructs.IConstruct
	// A token representing a list of ARNs of the load balancers that route traffic to this target group.
	LoadBalancerArns() *string
	// Return an object to depend on the listeners added to this target group.
	LoadBalancerAttached() constructs.IDependable
	// ARN of the target group.
	TargetGroupArn() *string
	// The name of the target group.
	TargetGroupName() *string
}

// The jsii proxy for ITargetGroup
type jsiiProxy_ITargetGroup struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITargetGroup) LoadBalancerArns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroup) LoadBalancerAttached() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroup) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetGroup) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

// What kind of addresses to allocate to the load balancer.
type IpAddressType string

const (
	IpAddressType_IPV4 IpAddressType = "IPV4"
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
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
// TODO: EXAMPLE
//
type ListenerAction interface {
	IListenerAction
	Next() ListenerAction
	Bind(scope constructs.Construct, listener IApplicationListener, associatingConstruct constructs.IConstruct)
	RenderActions() *[]*CfnListener_ActionProperty
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
func NewListenerAction(actionJson *CfnListener_ActionProperty, next ListenerAction) ListenerAction {
	_init_.Initialize()

	j := jsiiProxy_ListenerAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		[]interface{}{actionJson, next},
		&j,
	)

	return &j
}

// Create an instance of ListenerAction.
//
// The default class should be good enough for most cases and
// should be created by using one of the static factory functions,
// but allow overriding to make sure we allow flexibility for the future.
func NewListenerAction_Override(l ListenerAction, actionJson *CfnListener_ActionProperty, next ListenerAction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		[]interface{}{actionJson, next},
		l,
	)
}

// Authenticate using an identity provider (IdP) that is compliant with OpenID Connect (OIDC).
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-authenticate-users.html#oidc-requirements
//
func ListenerAction_AuthenticateOidc(options *AuthenticateOidcOptions) ListenerAction {
	_init_.Initialize()

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

	var returns ListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerAction",
		"weightedForward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

// Called when the action is being used in a listener.
func (l *jsiiProxy_ListenerAction) Bind(scope constructs.Construct, listener IApplicationListener, associatingConstruct constructs.IConstruct) {
	_jsii_.InvokeVoid(
		l,
		"bind",
		[]interface{}{scope, listener, associatingConstruct},
	)
}

// Render the actions in this chain.
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

// Renumber the "order" fields in the actions array.
//
// We don't number for 0 or 1 elements, but otherwise number them 1...#actions
// so ELB knows about the right order.
//
// Do this in `ListenerAction` instead of in `Listener` so that we give
// users the opportunity to override by subclassing and overriding `renderActions`.
func (l *jsiiProxy_ListenerAction) Renumber(actions *[]*CfnListener_ActionProperty) *[]*CfnListener_ActionProperty {
	var returns *[]*CfnListener_ActionProperty

	_jsii_.Invoke(
		l,
		"renumber",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

// A certificate source for an ELBv2 listener.
//
// TODO: EXAMPLE
//
type ListenerCertificate interface {
	IListenerCertificate
	CertificateArn() *string
}

// The jsii proxy struct for ListenerCertificate
type jsiiProxy_ListenerCertificate struct {
	jsiiProxy_IListenerCertificate
}

func (j *jsiiProxy_ListenerCertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}


func NewListenerCertificate(certificateArn *string) ListenerCertificate {
	_init_.Initialize()

	j := jsiiProxy_ListenerCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		[]interface{}{certificateArn},
		&j,
	)

	return &j
}

func NewListenerCertificate_Override(l ListenerCertificate, certificateArn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		[]interface{}{certificateArn},
		l,
	)
}

// Use any certificate, identified by its ARN, as a listener certificate.
func ListenerCertificate_FromArn(certificateArn *string) ListenerCertificate {
	_init_.Initialize()

	var returns ListenerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		"fromArn",
		[]interface{}{certificateArn},
		&returns,
	)

	return returns
}

// Use an ACM certificate as a listener certificate.
func ListenerCertificate_FromCertificateManager(acmCertificate awscertificatemanager.ICertificate) ListenerCertificate {
	_init_.Initialize()

	var returns ListenerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		"fromCertificateManager",
		[]interface{}{acmCertificate},
		&returns,
	)

	return returns
}

// ListenerCondition providers definition.
//
// TODO: EXAMPLE
//
type ListenerCondition interface {
	RenderRawCondition() interface{}
}

// The jsii proxy struct for ListenerCondition
type jsiiProxy_ListenerCondition struct {
	_ byte // padding
}

func NewListenerCondition_Override(l ListenerCondition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		nil, // no parameters
		l,
	)
}

// Create a host-header listener rule condition.
func ListenerCondition_HostHeaders(values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"hostHeaders",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a http-header listener rule condition.
func ListenerCondition_HttpHeader(name *string, values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"httpHeader",
		[]interface{}{name, values},
		&returns,
	)

	return returns
}

// Create a http-request-method listener rule condition.
func ListenerCondition_HttpRequestMethods(values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"httpRequestMethods",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a path-pattern listener rule condition.
func ListenerCondition_PathPatterns(values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"pathPatterns",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a query-string listener rule condition.
func ListenerCondition_QueryStrings(values *[]*QueryStringCondition) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"queryStrings",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Create a source-ip listener rule condition.
func ListenerCondition_SourceIps(values *[]*string) ListenerCondition {
	_init_.Initialize()

	var returns ListenerCondition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCondition",
		"sourceIps",
		[]interface{}{values},
		&returns,
	)

	return returns
}

// Render the raw Cfn listener rule condition object.
func (l *jsiiProxy_ListenerCondition) RenderRawCondition() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderRawCondition",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Result of attaching a target to load balancer.
//
// TODO: EXAMPLE
//
type LoadBalancerTargetProps struct {
	// What kind of target this is.
	TargetType TargetType `json:"targetType" yaml:"targetType"`
	// JSON representing the target's direct addition to the TargetGroup list.
	//
	// May be omitted if the target is going to register itself later.
	TargetJson interface{} `json:"targetJson" yaml:"targetJson"`
}

// Options for `NetworkListenerAction.forward()`.
//
// TODO: EXAMPLE
//
type NetworkForwardOptions struct {
	// For how long clients should be directed to the same target group.
	//
	// Range between 1 second and 7 days.
	StickinessDuration awscdk.Duration `json:"stickinessDuration" yaml:"stickinessDuration"`
}

// Define a Network Listener.
//
// TODO: EXAMPLE
//
type NetworkListener interface {
	BaseListener
	INetworkListener
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	LoadBalancer() INetworkLoadBalancer
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddAction(_id *string, props *AddNetworkActionProps)
	AddTargetGroups(_id *string, targetGroups ...INetworkTargetGroup)
	AddTargets(id *string, props *AddNetworkTargetsProps) NetworkTargetGroup
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	ValidateListener() *[]*string
}

// The jsii proxy struct for NetworkListener
type jsiiProxy_NetworkListener struct {
	jsiiProxy_BaseListener
	jsiiProxy_INetworkListener
}

func (j *jsiiProxy_NetworkListener) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkListener) LoadBalancer() INetworkLoadBalancer {
	var returns INetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkListener) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewNetworkListener(scope constructs.Construct, id *string, props *NetworkListenerProps) NetworkListener {
	_init_.Initialize()

	j := jsiiProxy_NetworkListener{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkListener_Override(n NetworkListener, scope constructs.Construct, id *string, props *NetworkListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListener",
		[]interface{}{scope, id, props},
		n,
	)
}

// Looks up a network listener.
func NetworkListener_FromLookup(scope constructs.Construct, id *string, options *NetworkListenerLookupOptions) INetworkListener {
	_init_.Initialize()

	var returns INetworkListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListener",
		"fromLookup",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Import an existing listener.
func NetworkListener_FromNetworkListenerArn(scope constructs.Construct, id *string, networkListenerArn *string) INetworkListener {
	_init_.Initialize()

	var returns INetworkListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListener",
		"fromNetworkListenerArn",
		[]interface{}{scope, id, networkListenerArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NetworkListener_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListener",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Perform the given Action on incoming requests.
//
// This allows full control of the default Action of the load balancer,
// including weighted forwarding. See the `NetworkListenerAction` class for
// all options.
func (n *jsiiProxy_NetworkListener) AddAction(_id *string, props *AddNetworkActionProps) {
	_jsii_.InvokeVoid(
		n,
		"addAction",
		[]interface{}{_id, props},
	)
}

// Load balance incoming requests to the given target groups.
//
// All target groups will be load balanced to with equal weight and without
// stickiness. For a more complex configuration than that, use `addAction()`.
func (n *jsiiProxy_NetworkListener) AddTargetGroups(_id *string, targetGroups ...INetworkTargetGroup) {
	args := []interface{}{_id}
	for _, a := range targetGroups {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addTargetGroups",
		args,
	)
}

// Load balance incoming requests to the given load balancing targets.
//
// This method implicitly creates a NetworkTargetGroup for the targets
// involved, and a 'forward' action to route traffic to the given TargetGroup.
//
// If you want more control over the precise setup, create the TargetGroup
// and use `addAction` yourself.
//
// It's possible to add conditions to the targets added in this way. At least
// one set of targets must be added without conditions.
//
// Returns: The newly created target group
func (n *jsiiProxy_NetworkListener) AddTargets(id *string, props *AddNetworkTargetsProps) NetworkTargetGroup {
	var returns NetworkTargetGroup

	_jsii_.Invoke(
		n,
		"addTargets",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (n *jsiiProxy_NetworkListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NetworkListener) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (n *jsiiProxy_NetworkListener) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (n *jsiiProxy_NetworkListener) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (n *jsiiProxy_NetworkListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate this listener.
func (n *jsiiProxy_NetworkListener) ValidateListener() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validateListener",
		nil, // no parameters
		&returns,
	)

	return returns
}

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
// TODO: EXAMPLE
//
type NetworkListenerAction interface {
	IListenerAction
	Next() NetworkListenerAction
	Bind(scope constructs.Construct, listener INetworkListener)
	RenderActions() *[]*CfnListener_ActionProperty
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
func NewNetworkListenerAction(actionJson *CfnListener_ActionProperty, next NetworkListenerAction) NetworkListenerAction {
	_init_.Initialize()

	j := jsiiProxy_NetworkListenerAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerAction",
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
func NewNetworkListenerAction_Override(n NetworkListenerAction, actionJson *CfnListener_ActionProperty, next NetworkListenerAction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerAction",
		[]interface{}{actionJson, next},
		n,
	)
}

// Forward to one or more Target Groups.
func NetworkListenerAction_Forward(targetGroups *[]INetworkTargetGroup, options *NetworkForwardOptions) NetworkListenerAction {
	_init_.Initialize()

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

	var returns NetworkListenerAction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListenerAction",
		"weightedForward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

// Called when the action is being used in a listener.
func (n *jsiiProxy_NetworkListenerAction) Bind(scope constructs.Construct, listener INetworkListener) {
	_jsii_.InvokeVoid(
		n,
		"bind",
		[]interface{}{scope, listener},
	)
}

// Render the actions in this chain.
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

// Renumber the "order" fields in the actions array.
//
// We don't number for 0 or 1 elements, but otherwise number them 1...#actions
// so ELB knows about the right order.
//
// Do this in `NetworkListenerAction` instead of in `Listener` so that we give
// users the opportunity to override by subclassing and overriding `renderActions`.
func (n *jsiiProxy_NetworkListenerAction) Renumber(actions *[]*CfnListener_ActionProperty) *[]*CfnListener_ActionProperty {
	var returns *[]*CfnListener_ActionProperty

	_jsii_.Invoke(
		n,
		"renumber",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

// Options for looking up a network listener.
//
// TODO: EXAMPLE
//
type NetworkListenerLookupOptions struct {
	// Filter listeners by listener port.
	ListenerPort *float64 `json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	LoadBalancerTags *map[string]*string `json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Protocol of the listener port.
	ListenerProtocol Protocol `json:"listenerProtocol" yaml:"listenerProtocol"`
}

// Properties for a Network Listener attached to a Load Balancer.
//
// TODO: EXAMPLE
//
type NetworkListenerProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `json:"port" yaml:"port"`
	// Application-Layer Protocol Negotiation (ALPN) is a TLS extension that is sent on the initial TLS handshake hello messages.
	//
	// ALPN enables the application layer to negotiate which protocols should be used over a secure connection, such as HTTP/1 and HTTP/2.
	//
	// Can only be specified together with Protocol TLS.
	AlpnPolicy AlpnPolicy `json:"alpnPolicy" yaml:"alpnPolicy"`
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default Action of the load balancer,
	// including weighted forwarding. See the `NetworkListenerAction` class for
	// all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction NetworkListenerAction `json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]INetworkTargetGroup `json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Protocol for listener, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// SSL Policy.
	SslPolicy SslPolicy `json:"sslPolicy" yaml:"sslPolicy"`
	// The load balancer to attach this listener to.
	LoadBalancer INetworkLoadBalancer `json:"loadBalancer" yaml:"loadBalancer"`
}

// Define a new network load balancer.
//
// TODO: EXAMPLE
//
type NetworkLoadBalancer interface {
	BaseLoadBalancer
	INetworkLoadBalancer
	Env() *awscdk.ResourceEnvironment
	LoadBalancerArn() *string
	LoadBalancerCanonicalHostedZoneId() *string
	LoadBalancerDnsName() *string
	LoadBalancerFullName() *string
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricActiveFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricNewFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTcpClientResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTcpElbResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTcpTargetResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	RemoveAttribute(key *string)
	SetAttribute(key *string, value *string)
	ToString() *string
}

// The jsii proxy struct for NetworkLoadBalancer
type jsiiProxy_NetworkLoadBalancer struct {
	jsiiProxy_BaseLoadBalancer
	jsiiProxy_INetworkLoadBalancer
}

func (j *jsiiProxy_NetworkLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewNetworkLoadBalancer(scope constructs.Construct, id *string, props *NetworkLoadBalancerProps) NetworkLoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_NetworkLoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkLoadBalancer_Override(n NetworkLoadBalancer, scope constructs.Construct, id *string, props *NetworkLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		[]interface{}{scope, id, props},
		n,
	)
}

// Looks up the network load balancer.
func NetworkLoadBalancer_FromLookup(scope constructs.Construct, id *string, options *NetworkLoadBalancerLookupOptions) INetworkLoadBalancer {
	_init_.Initialize()

	var returns INetworkLoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"fromLookup",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func NetworkLoadBalancer_FromNetworkLoadBalancerAttributes(scope constructs.Construct, id *string, attrs *NetworkLoadBalancerAttributes) INetworkLoadBalancer {
	_init_.Initialize()

	var returns INetworkLoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"fromNetworkLoadBalancerAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NetworkLoadBalancer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a listener to this load balancer.
//
// Returns: The newly created listener
func (n *jsiiProxy_NetworkLoadBalancer) AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener {
	var returns NetworkListener

	_jsii_.Invoke(
		n,
		"addListener",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (n *jsiiProxy_NetworkLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NetworkLoadBalancer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (n *jsiiProxy_NetworkLoadBalancer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (n *jsiiProxy_NetworkLoadBalancer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Enable access logging for this load balancer.
//
// A region must be specified on the stack containing the load balancer; you cannot enable logging on
// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func (n *jsiiProxy_NetworkLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	_jsii_.InvokeVoid(
		n,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

// Return the given named metric for this Network Load Balancer.
func (n *jsiiProxy_NetworkLoadBalancer) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// The total number of concurrent TCP flows (or connections) from clients to targets.
//
// This metric includes connections in the SYN_SENT and ESTABLISHED states.
// TCP connections are not terminated at the load balancer, so a client
// opening a TCP connection to a target counts as a single flow.
func (n *jsiiProxy_NetworkLoadBalancer) MetricActiveFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricActiveFlowCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of load balancer capacity units (LCU) used by your load balancer.
func (n *jsiiProxy_NetworkLoadBalancer) MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricConsumedLCUs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total number of new TCP flows (or connections) established from clients to targets in the time period.
func (n *jsiiProxy_NetworkLoadBalancer) MetricNewFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricNewFlowCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total number of bytes processed by the load balancer, including TCP/IP headers.
func (n *jsiiProxy_NetworkLoadBalancer) MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricProcessedBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total number of reset (RST) packets sent from a client to a target.
//
// These resets are generated by the client and forwarded by the load balancer.
func (n *jsiiProxy_NetworkLoadBalancer) MetricTcpClientResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricTcpClientResetCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total number of reset (RST) packets generated by the load balancer.
func (n *jsiiProxy_NetworkLoadBalancer) MetricTcpElbResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricTcpElbResetCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The total number of reset (RST) packets sent from a target to a client.
//
// These resets are generated by the target and forwarded by the load balancer.
func (n *jsiiProxy_NetworkLoadBalancer) MetricTcpTargetResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricTcpTargetResetCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Remove an attribute from the load balancer.
func (n *jsiiProxy_NetworkLoadBalancer) RemoveAttribute(key *string) {
	_jsii_.InvokeVoid(
		n,
		"removeAttribute",
		[]interface{}{key},
	)
}

// Set a non-standard attribute on the load balancer.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
//
func (n *jsiiProxy_NetworkLoadBalancer) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		n,
		"setAttribute",
		[]interface{}{key, value},
	)
}

// Returns a string representation of this construct.
func (n *jsiiProxy_NetworkLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to reference an existing load balancer.
//
// TODO: EXAMPLE
//
type NetworkLoadBalancerAttributes struct {
	// ARN of the load balancer.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The canonical hosted zone ID of this load balancer.
	LoadBalancerCanonicalHostedZoneId *string `json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// The DNS name of this load balancer.
	LoadBalancerDnsName *string `json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// The VPC to associate with the load balancer.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Options for looking up an NetworkLoadBalancer.
//
// TODO: EXAMPLE
//
type NetworkLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *map[string]*string `json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Properties for a network load balancer.
//
// TODO: EXAMPLE
//
type NetworkLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// Whether the load balancer has an internet-routable address.
	InternetFacing *bool `json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	LoadBalancerName *string `json:"loadBalancerName" yaml:"loadBalancerName"`
	// Which subnets place the load balancer in.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// Indicates whether cross-zone load balancing is enabled.
	CrossZoneEnabled *bool `json:"crossZoneEnabled" yaml:"crossZoneEnabled"`
}

// Define a Network Target Group.
//
// TODO: EXAMPLE
//
type NetworkTargetGroup interface {
	TargetGroupBase
	INetworkTargetGroup
	DefaultPort() *float64
	FirstLoadBalancerFullName() *string
	HealthCheck() *HealthCheck
	SetHealthCheck(val *HealthCheck)
	LoadBalancerArns() *string
	LoadBalancerAttached() constructs.IDependable
	LoadBalancerAttachedDependencies() constructs.DependencyGroup
	Node() constructs.Node
	TargetGroupArn() *string
	TargetGroupFullName() *string
	TargetGroupLoadBalancerArns() *[]*string
	TargetGroupName() *string
	TargetType() TargetType
	SetTargetType(val TargetType)
	AddLoadBalancerTarget(props *LoadBalancerTargetProps)
	AddTarget(targets ...INetworkLoadBalancerTarget)
	ConfigureHealthCheck(healthCheck *HealthCheck)
	MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricUnHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	RegisterListener(listener INetworkListener)
	SetAttribute(key *string, value *string)
	ToString() *string
	ValidateTargetGroup() *[]*string
}

// The jsii proxy struct for NetworkTargetGroup
type jsiiProxy_NetworkTargetGroup struct {
	jsiiProxy_TargetGroupBase
	jsiiProxy_INetworkTargetGroup
}

func (j *jsiiProxy_NetworkTargetGroup) DefaultPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) FirstLoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstLoadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) HealthCheck() *HealthCheck {
	var returns *HealthCheck
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) LoadBalancerArns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) LoadBalancerAttached() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) LoadBalancerAttachedDependencies() constructs.DependencyGroup {
	var returns constructs.DependencyGroup
	_jsii_.Get(
		j,
		"loadBalancerAttachedDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetGroupFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetGroupLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkTargetGroup) TargetType() TargetType {
	var returns TargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


func NewNetworkTargetGroup(scope constructs.Construct, id *string, props *NetworkTargetGroupProps) NetworkTargetGroup {
	_init_.Initialize()

	j := jsiiProxy_NetworkTargetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkTargetGroup_Override(n NetworkTargetGroup, scope constructs.Construct, id *string, props *NetworkTargetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NetworkTargetGroup) SetHealthCheck(val *HealthCheck) {
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_NetworkTargetGroup) SetTargetType(val TargetType) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

// Import an existing target group.
func NetworkTargetGroup_FromTargetGroupAttributes(scope constructs.Construct, id *string, attrs *TargetGroupAttributes) INetworkTargetGroup {
	_init_.Initialize()

	var returns INetworkTargetGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
		"fromTargetGroupAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NetworkTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Register the given load balancing target as part of this group.
func (n *jsiiProxy_NetworkTargetGroup) AddLoadBalancerTarget(props *LoadBalancerTargetProps) {
	_jsii_.InvokeVoid(
		n,
		"addLoadBalancerTarget",
		[]interface{}{props},
	)
}

// Add a load balancing target to this target group.
func (n *jsiiProxy_NetworkTargetGroup) AddTarget(targets ...INetworkLoadBalancerTarget) {
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addTarget",
		args,
	)
}

// Set/replace the target group's health check.
func (n *jsiiProxy_NetworkTargetGroup) ConfigureHealthCheck(healthCheck *HealthCheck) {
	_jsii_.InvokeVoid(
		n,
		"configureHealthCheck",
		[]interface{}{healthCheck},
	)
}

// The number of targets that are considered healthy.
func (n *jsiiProxy_NetworkTargetGroup) MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricHealthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The number of targets that are considered unhealthy.
func (n *jsiiProxy_NetworkTargetGroup) MetricUnHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricUnHealthyHostCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register a listener that is load balancing to this target group.
//
// Don't call this directly. It will be called by listeners.
func (n *jsiiProxy_NetworkTargetGroup) RegisterListener(listener INetworkListener) {
	_jsii_.InvokeVoid(
		n,
		"registerListener",
		[]interface{}{listener},
	)
}

// Set a non-standard attribute on the target group.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
//
func (n *jsiiProxy_NetworkTargetGroup) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		n,
		"setAttribute",
		[]interface{}{key, value},
	)
}

// Returns a string representation of this construct.
func (n *jsiiProxy_NetworkTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkTargetGroup) ValidateTargetGroup() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validateTargetGroup",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a new Network Target Group.
//
// TODO: EXAMPLE
//
type NetworkTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	HealthCheck *HealthCheck `json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	TargetType TargetType `json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The port on which the listener listens for requests.
	Port *float64 `json:"port" yaml:"port"`
	// Indicates whether client IP preservation is enabled.
	PreserveClientIp *bool `json:"preserveClientIp" yaml:"preserveClientIp"`
	// Protocol for target group, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `json:"protocol" yaml:"protocol"`
	// Indicates whether Proxy Protocol version 2 is enabled.
	ProxyProtocolV2 *bool `json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]INetworkLoadBalancerTarget `json:"targets" yaml:"targets"`
}

// A Target Group and weight combination.
//
// TODO: EXAMPLE
//
type NetworkWeightedTargetGroup struct {
	// The target group.
	TargetGroup INetworkTargetGroup `json:"targetGroup" yaml:"targetGroup"`
	// The target group's weight.
	//
	// Range is [0..1000).
	Weight *float64 `json:"weight" yaml:"weight"`
}

// Backend protocol for network load balancers and health checks.
//
// TODO: EXAMPLE
//
type Protocol string

const (
	Protocol_HTTP Protocol = "HTTP"
	Protocol_HTTPS Protocol = "HTTPS"
	Protocol_TCP Protocol = "TCP"
	Protocol_TLS Protocol = "TLS"
	Protocol_UDP Protocol = "UDP"
	Protocol_TCP_UDP Protocol = "TCP_UDP"
)

// Properties for the key/value pair of the query string.
//
// TODO: EXAMPLE
//
type QueryStringCondition struct {
	// The query string value for the condition.
	Value *string `json:"value" yaml:"value"`
	// The query string key for the condition.
	Key *string `json:"key" yaml:"key"`
}

// Options for `ListenerAction.redirect()`.
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
//
// TODO: EXAMPLE
//
type RedirectOptions struct {
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	Host *string `json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.
	Path *string `json:"path" yaml:"path"`
	// The HTTP redirect code.
	//
	// The redirect is either permanent (HTTP 301) or temporary (HTTP 302).
	Permanent *bool `json:"permanent" yaml:"permanent"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	Port *string `json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added. You can specify any of the reserved keywords.
	Query *string `json:"query" yaml:"query"`
}

// Elastic Load Balancing provides the following security policies for Application Load Balancers.
//
// We recommend the Recommended policy for general use. You can
// use the ForwardSecrecy policy if you require Forward Secrecy
// (FS).
//
// You can use one of the TLS policies to meet compliance and security
// standards that require disabling certain TLS protocol versions, or to
// support legacy clients that require deprecated ciphers.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html
//
type SslPolicy string

const (
	SslPolicy_RECOMMENDED SslPolicy = "RECOMMENDED"
	SslPolicy_FORWARD_SECRECY_TLS12_RES_GCM SslPolicy = "FORWARD_SECRECY_TLS12_RES_GCM"
	SslPolicy_FORWARD_SECRECY_TLS12_RES SslPolicy = "FORWARD_SECRECY_TLS12_RES"
	SslPolicy_FORWARD_SECRECY_TLS12 SslPolicy = "FORWARD_SECRECY_TLS12"
	SslPolicy_FORWARD_SECRECY_TLS11 SslPolicy = "FORWARD_SECRECY_TLS11"
	SslPolicy_FORWARD_SECRECY SslPolicy = "FORWARD_SECRECY"
	SslPolicy_TLS12 SslPolicy = "TLS12"
	SslPolicy_TLS12_EXT SslPolicy = "TLS12_EXT"
	SslPolicy_TLS11 SslPolicy = "TLS11"
	SslPolicy_LEGACY SslPolicy = "LEGACY"
)

// Properties to reference an existing target group.
//
// TODO: EXAMPLE
//
type TargetGroupAttributes struct {
	// ARN of the target group.
	TargetGroupArn *string `json:"targetGroupArn" yaml:"targetGroupArn"`
	// A Token representing the list of ARNs for the load balancer routing to this target group.
	LoadBalancerArns *string `json:"loadBalancerArns" yaml:"loadBalancerArns"`
}

// Define the target of a load balancer.
type TargetGroupBase interface {
	constructs.Construct
	ITargetGroup
	DefaultPort() *float64
	FirstLoadBalancerFullName() *string
	HealthCheck() *HealthCheck
	SetHealthCheck(val *HealthCheck)
	LoadBalancerArns() *string
	LoadBalancerAttached() constructs.IDependable
	LoadBalancerAttachedDependencies() constructs.DependencyGroup
	Node() constructs.Node
	TargetGroupArn() *string
	TargetGroupFullName() *string
	TargetGroupLoadBalancerArns() *[]*string
	TargetGroupName() *string
	TargetType() TargetType
	SetTargetType(val TargetType)
	AddLoadBalancerTarget(props *LoadBalancerTargetProps)
	ConfigureHealthCheck(healthCheck *HealthCheck)
	SetAttribute(key *string, value *string)
	ToString() *string
	ValidateTargetGroup() *[]*string
}

// The jsii proxy struct for TargetGroupBase
type jsiiProxy_TargetGroupBase struct {
	internal.Type__constructsConstruct
	jsiiProxy_ITargetGroup
}

func (j *jsiiProxy_TargetGroupBase) DefaultPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) FirstLoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstLoadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) HealthCheck() *HealthCheck {
	var returns *HealthCheck
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) LoadBalancerArns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) LoadBalancerAttached() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"loadBalancerAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) LoadBalancerAttachedDependencies() constructs.DependencyGroup {
	var returns constructs.DependencyGroup
	_jsii_.Get(
		j,
		"loadBalancerAttachedDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetGroupFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetGroupLoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupLoadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetGroupBase) TargetType() TargetType {
	var returns TargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}


func NewTargetGroupBase_Override(t TargetGroupBase, scope constructs.Construct, id *string, baseProps *BaseTargetGroupProps, additionalProps interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.TargetGroupBase",
		[]interface{}{scope, id, baseProps, additionalProps},
		t,
	)
}

func (j *jsiiProxy_TargetGroupBase) SetHealthCheck(val *HealthCheck) {
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_TargetGroupBase) SetTargetType(val TargetType) {
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TargetGroupBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.TargetGroupBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Register the given load balancing target as part of this group.
func (t *jsiiProxy_TargetGroupBase) AddLoadBalancerTarget(props *LoadBalancerTargetProps) {
	_jsii_.InvokeVoid(
		t,
		"addLoadBalancerTarget",
		[]interface{}{props},
	)
}

// Set/replace the target group's health check.
func (t *jsiiProxy_TargetGroupBase) ConfigureHealthCheck(healthCheck *HealthCheck) {
	_jsii_.InvokeVoid(
		t,
		"configureHealthCheck",
		[]interface{}{healthCheck},
	)
}

// Set a non-standard attribute on the target group.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
//
func (t *jsiiProxy_TargetGroupBase) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		t,
		"setAttribute",
		[]interface{}{key, value},
	)
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TargetGroupBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetGroupBase) ValidateTargetGroup() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validateTargetGroup",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Load balancing algorithmm type for target groups.
type TargetGroupLoadBalancingAlgorithmType string

const (
	TargetGroupLoadBalancingAlgorithmType_ROUND_ROBIN TargetGroupLoadBalancingAlgorithmType = "ROUND_ROBIN"
	TargetGroupLoadBalancingAlgorithmType_LEAST_OUTSTANDING_REQUESTS TargetGroupLoadBalancingAlgorithmType = "LEAST_OUTSTANDING_REQUESTS"
)

// How to interpret the load balancing target identifiers.
//
// TODO: EXAMPLE
//
type TargetType string

const (
	TargetType_INSTANCE TargetType = "INSTANCE"
	TargetType_IP TargetType = "IP"
	TargetType_LAMBDA TargetType = "LAMBDA"
	TargetType_ALB TargetType = "ALB"
)

// What to do with unauthenticated requests.
type UnauthenticatedAction string

const (
	UnauthenticatedAction_DENY UnauthenticatedAction = "DENY"
	UnauthenticatedAction_ALLOW UnauthenticatedAction = "ALLOW"
	UnauthenticatedAction_AUTHENTICATE UnauthenticatedAction = "AUTHENTICATE"
)

// A Target Group and weight combination.
//
// TODO: EXAMPLE
//
type WeightedTargetGroup struct {
	// The target group.
	TargetGroup IApplicationTargetGroup `json:"targetGroup" yaml:"targetGroup"`
	// The target group's weight.
	//
	// Range is [0..1000).
	Weight *float64 `json:"weight" yaml:"weight"`
}

