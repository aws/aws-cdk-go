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
// Example:
//   var listener applicationListener
//
//
//   listener.addAction(jsii.String("Fixed"), &addApplicationActionProps{
//   	priority: jsii.Number(10),
//   	conditions: []listenerCondition{
//   		elbv2.*listenerCondition.pathPatterns([]*string{
//   			jsii.String("/ok"),
//   		}),
//   	},
//   	action: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
//   		contentType: elbv2.contentType_TEXT_PLAIN,
//   		messageBody: jsii.String("OK"),
//   	}),
//   })
//
type AddApplicationActionProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Action to perform.
	Action ListenerAction `field:"required" json:"action" yaml:"action"`
}

// Properties for adding a new target group to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationTargetGroup applicationTargetGroup
//   var listenerCondition listenerCondition
//
//   addApplicationTargetGroupsProps := &addApplicationTargetGroupsProps{
//   	targetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//
//   	// the properties below are optional
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	priority: jsii.Number(123),
//   }
//
type AddApplicationTargetGroupsProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Target groups to forward requests to.
	TargetGroups *[]IApplicationTargetGroup `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

// Properties for adding new targets to a listener.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var asg autoScalingGroup
//
//   var vpc vpc
//
//
//   // Create the load balancer in a VPC. 'internetFacing' is 'false'
//   // by default, which creates an internal load balancer.
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   // Add a listener and open up the load balancer's security group
//   // to the world.
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//
//   	// 'open: true' is the default, you can leave it out if you want. Set it
//   	// to 'false' and use `listener.connections` if you want to be selective
//   	// about who can access the load balancer.
//   	open: jsii.Boolean(true),
//   })
//
//   // Create an AutoScaling group and add it as a load balancing
//   // target to the listener.
//   listener.addTargets(jsii.String("ApplicationFleet"), &addApplicationTargetsProps{
//   	port: jsii.Number(8080),
//   	targets: []iApplicationLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type AddApplicationTargetsProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The load balancing algorithm to select targets for routing requests.
	LoadBalancingAlgorithmType TargetGroupLoadBalancingAlgorithmType `field:"optional" json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// The port on which the listener listens for requests.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol version to use.
	ProtocolVersion ApplicationProtocolVersion `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The time period during which the load balancer sends a newly registered target a linearly increasing share of the traffic to the target group.
	//
	// The range is 30-900 seconds (15 minutes).
	SlowStart awscdk.Duration `field:"optional" json:"slowStart" yaml:"slowStart"`
	// The stickiness cookie expiration period.
	//
	// Setting this value enables load balancer stickiness.
	//
	// After this period, the cookie is considered stale. The minimum value is
	// 1 second and the maximum value is 7 days (604800 seconds).
	StickinessCookieDuration awscdk.Duration `field:"optional" json:"stickinessCookieDuration" yaml:"stickinessCookieDuration"`
	// The name of an application-based stickiness cookie.
	//
	// Names that start with the following prefixes are not allowed: AWSALB, AWSALBAPP,
	// and AWSALBTG; they're reserved for use by the load balancer.
	//
	// Note: `stickinessCookieName` parameter depends on the presence of `stickinessCookieDuration` parameter.
	// If `stickinessCookieDuration` is not set, `stickinessCookieName` will be omitted.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/sticky-sessions.html
	//
	StickinessCookieName *string `field:"optional" json:"stickinessCookieName" yaml:"stickinessCookieName"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. All target must be of the same type.
	Targets *[]IApplicationLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

// Properties for adding a new action to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkListenerAction networkListenerAction
//
//   addNetworkActionProps := &addNetworkActionProps{
//   	action: networkListenerAction,
//   }
//
type AddNetworkActionProps struct {
	// Action to perform.
	Action NetworkListenerAction `field:"required" json:"action" yaml:"action"`
}

// Properties for adding new network targets to a listener.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addNetworkTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type AddNetworkTargetsProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Indicates whether client IP preservation is enabled.
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Protocol for target group, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether Proxy Protocol version 2 is enabled.
	ProxyProtocolV2 *bool `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]INetworkLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

// Properties for adding a conditional load balancing rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerCondition listenerCondition
//
//   addRuleProps := &addRuleProps{
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	priority: jsii.Number(123),
//   }
//
type AddRuleProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

// Application-Layer Protocol Negotiation Policies for network load balancers.
//
// Which protocols should be used over a secure connection.
type AlpnPolicy string

const (
	// Negotiate only HTTP/1.*. The ALPN preference list is http/1.1, http/1.0.
	AlpnPolicy_HTTP1_ONLY AlpnPolicy = "HTTP1_ONLY"
	// Negotiate only HTTP/2.
	//
	// The ALPN preference list is h2.
	AlpnPolicy_HTTP2_ONLY AlpnPolicy = "HTTP2_ONLY"
	// Prefer HTTP/1.* over HTTP/2 (which can be useful for HTTP/2 testing). The ALPN preference list is http/1.1, http/1.0, h2.
	AlpnPolicy_HTTP2_OPTIONAL AlpnPolicy = "HTTP2_OPTIONAL"
	// Prefer HTTP/2 over HTTP/1.*. The ALPN preference list is h2, http/1.1, http/1.0.
	AlpnPolicy_HTTP2_PREFERRED AlpnPolicy = "HTTP2_PREFERRED"
	// Do not negotiate ALPN.
	AlpnPolicy_NONE AlpnPolicy = "NONE"
)

// Define an ApplicationListener.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var alb applicationLoadBalancer
//
//   listener := alb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   targetGroup := listener.addTargets(jsii.String("Fleet"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &serverDeploymentGroupProps{
//   	loadBalancer: codedeploy.loadBalancer.application(targetGroup),
//   })
//
type ApplicationListener interface {
	BaseListener
	IApplicationListener
	// Manage connections to this ApplicationListener.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	// Load balancer this listener is associated with.
	LoadBalancer() IApplicationLoadBalancer
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Perform the given default action on incoming requests.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses. See
	// the `ListenerAction` class for all options.
	//
	// It's possible to add routing conditions to the Action added in this way.
	// At least one Action must be added without conditions (which becomes the
	// default Action).
	AddAction(id *string, props *AddApplicationActionProps)
	// Add one or more certificates to this listener.
	//
	// After the first certificate, this creates ApplicationListenerCertificates
	// resources since cloudformation requires the certificates array on the
	// listener resource to have a length of 1.
	AddCertificates(id *string, certificates *[]IListenerCertificate)
	// Load balance incoming requests to the given target groups.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use `addAction()`.
	//
	// It's possible to add routing conditions to the TargetGroups added in this
	// way. At least one TargetGroup must be added without conditions (which will
	// become the default Action for this listener).
	AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps)
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
	// Returns: The newly created target group.
	AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Register that a connectable that has been added to this load balancer.
	//
	// Don't call this directly. It is called by ApplicationTargetGroup.
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
	// Returns a string representation of this construct.
	ToString() *string
	// Validate this listener.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

// Returns true if the construct was created by CDK, and false otherwise.
func ApplicationListener_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"isOwnedResource",
		[]interface{}{construct},
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

func (a *jsiiProxy_ApplicationListener) AddAction(id *string, props *AddApplicationActionProps) {
	_jsii_.InvokeVoid(
		a,
		"addAction",
		[]interface{}{id, props},
	)
}

func (a *jsiiProxy_ApplicationListener) AddCertificates(id *string, certificates *[]IListenerCertificate) {
	_jsii_.InvokeVoid(
		a,
		"addCertificates",
		[]interface{}{id, certificates},
	)
}

func (a *jsiiProxy_ApplicationListener) AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps) {
	_jsii_.InvokeVoid(
		a,
		"addTargetGroups",
		[]interface{}{id, props},
	)
}

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

func (a *jsiiProxy_ApplicationListener) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	_jsii_.InvokeVoid(
		a,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   applicationListenerAttributes := &applicationListenerAttributes{
//   	listenerArn: jsii.String("listenerArn"),
//
//   	// the properties below are optional
//   	defaultPort: jsii.Number(123),
//   	securityGroup: securityGroup,
//   }
//
type ApplicationListenerAttributes struct {
	// ARN of the listener.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// The default port on which this listener is listening.
	DefaultPort *float64 `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// Security group of the load balancer this listener is associated with.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

// Add certificates to a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationListener applicationListener
//   var listenerCertificate listenerCertificate
//
//   applicationListenerCertificate := awscdk.Aws_elasticloadbalancingv2.NewApplicationListenerCertificate(this, jsii.String("MyApplicationListenerCertificate"), &applicationListenerCertificateProps{
//   	listener: applicationListener,
//
//   	// the properties below are optional
//   	certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   })
//
type ApplicationListenerCertificate interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationListener applicationListener
//   var listenerCertificate listenerCertificate
//
//   applicationListenerCertificateProps := &applicationListenerCertificateProps{
//   	listener: applicationListener,
//
//   	// the properties below are optional
//   	certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   }
//
type ApplicationListenerCertificateProps struct {
	// The listener to attach the rule to.
	Listener IApplicationListener `field:"required" json:"listener" yaml:"listener"`
	// Certificates to attach.
	//
	// Duplicates are not allowed.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
}

// Options for ApplicationListener lookup.
//
// Example:
//   listener := elbv2.applicationListener.fromLookup(this, jsii.String("ALBListener"), &applicationListenerLookupOptions{
//   	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-east-2:123456789012:loadbalancer/app/my-load-balancer/1234567890123456"),
//   	listenerProtocol: elbv2.applicationProtocol_HTTPS,
//   	listenerPort: jsii.Number(443),
//   })
//
type ApplicationListenerLookupOptions struct {
	// Filter listeners by listener port.
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// ARN of the listener to look up.
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener protocol.
	ListenerProtocol ApplicationProtocol `field:"optional" json:"listenerProtocol" yaml:"listenerProtocol"`
}

// Properties for defining a standalone ApplicationListener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationLoadBalancer applicationLoadBalancer
//   var applicationTargetGroup applicationTargetGroup
//   var listenerAction listenerAction
//   var listenerCertificate listenerCertificate
//
//   applicationListenerProps := &applicationListenerProps{
//   	loadBalancer: applicationLoadBalancer,
//
//   	// the properties below are optional
//   	certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   	defaultAction: listenerAction,
//   	defaultTargetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   	open: jsii.Boolean(false),
//   	port: jsii.Number(123),
//   	protocol: awscdk.Aws_elasticloadbalancingv2.applicationProtocol_HTTP,
//   	sslPolicy: awscdk.*Aws_elasticloadbalancingv2.sslPolicy_RECOMMENDED_TLS,
//   }
//
type ApplicationListenerProps struct {
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses.
	//
	// See the `ListenerAction` class for all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction ListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]IApplicationTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Allow anyone to connect to the load balancer on the listener port.
	//
	// If this is specified, the load balancer will be opened up to anyone who can reach it.
	// For internal load balancers this is anyone in the same VPC. For public load
	// balancers, this is anyone on the internet.
	//
	// If you want to be more selective about who can access this load
	// balancer, set this to `false` and use the listener's `connections`
	// object to selectively grant access to the load balancer on the listener port.
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// The port on which the listener listens for requests.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The security policy that defines which ciphers and protocols are supported.
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// The load balancer to attach this listener to.
	LoadBalancer IApplicationLoadBalancer `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
}

// Define a new listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationListener applicationListener
//   var applicationTargetGroup applicationTargetGroup
//   var listenerAction listenerAction
//   var listenerCondition listenerCondition
//
//   applicationListenerRule := awscdk.Aws_elasticloadbalancingv2.NewApplicationListenerRule(this, jsii.String("MyApplicationListenerRule"), &applicationListenerRuleProps{
//   	listener: applicationListener,
//   	priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	action: listenerAction,
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	targetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   })
//
type ApplicationListenerRule interface {
	constructs.Construct
	// The ARN of this rule.
	ListenerRuleArn() *string
	// The tree node.
	Node() constructs.Node
	// Add a non-standard condition to this rule.
	AddCondition(condition ListenerCondition)
	// Configure the action to perform for this rule.
	ConfigureAction(action ListenerAction)
	// Returns a string representation of this construct.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (a *jsiiProxy_ApplicationListenerRule) AddCondition(condition ListenerCondition) {
	_jsii_.InvokeVoid(
		a,
		"addCondition",
		[]interface{}{condition},
	)
}

func (a *jsiiProxy_ApplicationListenerRule) ConfigureAction(action ListenerAction) {
	_jsii_.InvokeVoid(
		a,
		"configureAction",
		[]interface{}{action},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationListener applicationListener
//   var applicationTargetGroup applicationTargetGroup
//   var listenerAction listenerAction
//   var listenerCondition listenerCondition
//
//   applicationListenerRuleProps := &applicationListenerRuleProps{
//   	listener: applicationListener,
//   	priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	action: listenerAction,
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	targetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   }
//
type ApplicationListenerRuleProps struct {
	// Priority of the rule.
	//
	// The rule with the lowest priority will be used for every request.
	//
	// Priorities must be unique.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Action to perform when requests are received.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	Action ListenerAction `field:"optional" json:"action" yaml:"action"`
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Target groups to forward requests to.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	//
	// Implies a `forward` action.
	TargetGroups *[]IApplicationTargetGroup `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// The listener to attach the rule to.
	Listener IApplicationListener `field:"required" json:"listener" yaml:"listener"`
}

// Define an Application Load Balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("lb"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type ApplicationLoadBalancer interface {
	BaseLoadBalancer
	IApplicationLoadBalancer
	// The network connections associated with this resource.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The IP Address Type for this load balancer.
	IpAddressType() IpAddressType
	// A list of listeners that have been added to the load balancer.
	//
	// This list is only valid for owned constructs.
	Listeners() *[]ApplicationListener
	// The ARN of this load balancer.
	//
	// Example value: `arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-internal-load-balancer/50dc6c495c0c9188`.
	LoadBalancerArn() *string
	// The canonical hosted zone ID of this load balancer.
	//
	// Example value: `Z2P70J7EXAMPLE`.
	LoadBalancerCanonicalHostedZoneId() *string
	// The DNS name of this load balancer.
	//
	// Example value: `my-load-balancer-424835706.us-west-2.elb.amazonaws.com`
	LoadBalancerDnsName() *string
	// The full name of this load balancer.
	//
	// Example value: `app/my-load-balancer/50dc6c495c0c9188`.
	LoadBalancerFullName() *string
	// The name of this load balancer.
	//
	// Example value: `my-load-balancer`.
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The VPC this load balancer has been created in.
	//
	// This property is always defined (not `null` or `undefined`) for sub-classes of `BaseLoadBalancer`.
	Vpc() awsec2.IVpc
	// Add a new listener to this load balancer.
	AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener
	// Add a redirection listener to this load balancer.
	AddRedirect(props *ApplicationLoadBalancerRedirectConfig) ApplicationListener
	// Add a security group to this load balancer.
	AddSecurityGroup(securityGroup awsec2.ISecurityGroup)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Enable access logging for this load balancer.
	//
	// A region must be specified on the stack containing the load balancer; you cannot enable logging on
	// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	// Return the given named metric for this Application Load Balancer.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of concurrent TCP connections active from clients to the load balancer and from the load balancer to targets.
	MetricActiveConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of TLS connections initiated by the client that did not establish a session with the load balancer.
	//
	// Possible causes include a
	// mismatch of ciphers or protocols.
	MetricClientTlsNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of load balancer capacity units (LCU) used by your load balancer.
	MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of user authentications that could not be completed.
	//
	// Because an authenticate action was misconfigured, the load balancer
	// couldn't establish a connection with the IdP, or the load balancer
	// couldn't complete the authentication flow due to an internal error.
	MetricElbAuthError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of user authentications that could not be completed because the IdP denied access to the user or an authorization code was used more than once.
	MetricElbAuthFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time elapsed, in milliseconds, to query the IdP for the ID token and user info.
	//
	// If one or more of these operations fail, this is the time to failure.
	MetricElbAuthLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of authenticate actions that were successful.
	//
	// This metric is incremented at the end of the authentication workflow,
	// after the load balancer has retrieved the user claims from the IdP.
	MetricElbAuthSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of HTTP 3xx/4xx/5xx codes that originate from the load balancer.
	//
	// This does not include any response codes generated by the targets.
	MetricHttpCodeElb(code HttpCodeElb, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of HTTP 2xx/3xx/4xx/5xx response codes generated by all targets in the load balancer.
	//
	// This does not include any response codes generated by the load balancer.
	MetricHttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of fixed-response actions that were successful.
	MetricHttpFixedResponseCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of redirect actions that were successful.
	MetricHttpRedirectCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of redirect actions that couldn't be completed because the URL in the response location header is larger than 8K.
	MetricHttpRedirectUrlLimitExceededCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of bytes processed by the load balancer over IPv6.
	MetricIpv6ProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of IPv6 requests received by the load balancer.
	MetricIpv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of new TCP connections established from clients to the load balancer and from the load balancer to targets.
	MetricNewConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of bytes processed by the load balancer over IPv4 and IPv6.
	MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of connections that were rejected because the load balancer had reached its maximum number of connections.
	MetricRejectedConnectionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of requests processed over IPv4 and IPv6.
	//
	// This count includes only the requests with a response generated by a target of the load balancer.
	MetricRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of rules processed by the load balancer given a request rate averaged over an hour.
	MetricRuleEvaluations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of connections that were not successfully established between the load balancer and target.
	MetricTargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time elapsed, in seconds, after the request leaves the load balancer until a response from the target is received.
	MetricTargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of TLS connections initiated by the load balancer that did not establish a session with the target.
	//
	// Possible causes include a mismatch of ciphers or protocols.
	MetricTargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Remove an attribute from the load balancer.
	RemoveAttribute(key *string)
	// Set a non-standard attribute on the load balancer.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
	ValidateLoadBalancer() *[]*string
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

// Returns true if the construct was created by CDK, and false otherwise.
func ApplicationLoadBalancer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer",
		"isOwnedResource",
		[]interface{}{construct},
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

func (a *jsiiProxy_ApplicationLoadBalancer) AddSecurityGroup(securityGroup awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		a,
		"addSecurityGroup",
		[]interface{}{securityGroup},
	)
}

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

func (a *jsiiProxy_ApplicationLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	_jsii_.InvokeVoid(
		a,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

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

func (a *jsiiProxy_ApplicationLoadBalancer) RemoveAttribute(key *string) {
	_jsii_.InvokeVoid(
		a,
		"removeAttribute",
		[]interface{}{key},
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		a,
		"setAttribute",
		[]interface{}{key, value},
	)
}

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

func (a *jsiiProxy_ApplicationLoadBalancer) ValidateLoadBalancer() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validateLoadBalancer",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to reference an existing load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   applicationLoadBalancerAttributes := &applicationLoadBalancerAttributes{
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	securityGroupId: jsii.String("securityGroupId"),
//
//   	// the properties below are optional
//   	loadBalancerCanonicalHostedZoneId: jsii.String("loadBalancerCanonicalHostedZoneId"),
//   	loadBalancerDnsName: jsii.String("loadBalancerDnsName"),
//   	securityGroupAllowsAllOutbound: jsii.Boolean(false),
//   	vpc: vpc,
//   }
//
type ApplicationLoadBalancerAttributes struct {
	// ARN of the load balancer.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// ID of the load balancer's security group.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// The canonical hosted zone ID of this load balancer.
	LoadBalancerCanonicalHostedZoneId *string `field:"optional" json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// The DNS name of this load balancer.
	LoadBalancerDnsName *string `field:"optional" json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// Whether the security group allows all outbound traffic or not.
	//
	// Unless set to `false`, no egress rules will be added to the security group.
	SecurityGroupAllowsAllOutbound *bool `field:"optional" json:"securityGroupAllowsAllOutbound" yaml:"securityGroupAllowsAllOutbound"`
	// The VPC this load balancer has been created in, if available.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// Options for looking up an ApplicationLoadBalancer.
//
// Example:
//   loadBalancer := elbv2.applicationLoadBalancer.fromLookup(this, jsii.String("ALB"), &applicationLoadBalancerLookupOptions{
//   	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-east-2:123456789012:loadbalancer/app/my-load-balancer/1234567890123456"),
//   })
//
type ApplicationLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Properties for defining an Application Load Balancer.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
type ApplicationLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether the load balancer has an internet-routable address.
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// Which subnets place the load balancer in.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Indicates whether HTTP/2 is enabled.
	Http2Enabled *bool `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// The load balancer idle timeout, in seconds.
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// The type of IP addresses to use.
	//
	// Only applies to application load balancers.
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Security group to associate with this load balancer.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

// Properties for a redirection config.
//
// Example:
//   var lb applicationLoadBalancer
//
//
//   lb.addRedirect(&applicationLoadBalancerRedirectConfig{
//   	sourceProtocol: elbv2.applicationProtocol_HTTPS,
//   	sourcePort: jsii.Number(8443),
//   	targetProtocol: elbv2.*applicationProtocol_HTTP,
//   	targetPort: jsii.Number(8080),
//   })
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
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// The port number to listen to.
	SourcePort *float64 `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// The protocol of the listener being created.
	SourceProtocol ApplicationProtocol `field:"optional" json:"sourceProtocol" yaml:"sourceProtocol"`
	// The port number to redirect to.
	TargetPort *float64 `field:"optional" json:"targetPort" yaml:"targetPort"`
	// The protocol of the redirection target.
	TargetProtocol ApplicationProtocol `field:"optional" json:"targetProtocol" yaml:"targetProtocol"`
}

// Load balancing protocol for application load balancers.
//
// Example:
//   var vpc vpc
//
//
//   tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &applicationTargetGroupProps{
//   	targetType: elbv2.targetType_IP,
//   	port: jsii.Number(50051),
//   	protocol: elbv2.applicationProtocol_HTTP,
//   	protocolVersion: elbv2.applicationProtocolVersion_GRPC,
//   	healthCheck: &healthCheck{
//   		enabled: jsii.Boolean(true),
//   		healthyGrpcCodes: jsii.String("0-99"),
//   	},
//   	vpc: vpc,
//   })
//
type ApplicationProtocol string

const (
	// HTTP.
	ApplicationProtocol_HTTP ApplicationProtocol = "HTTP"
	// HTTPS.
	ApplicationProtocol_HTTPS ApplicationProtocol = "HTTPS"
)

// Load balancing protocol version for application load balancers.
//
// Example:
//   var vpc vpc
//
//
//   tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &applicationTargetGroupProps{
//   	targetType: elbv2.targetType_IP,
//   	port: jsii.Number(50051),
//   	protocol: elbv2.applicationProtocol_HTTP,
//   	protocolVersion: elbv2.applicationProtocolVersion_GRPC,
//   	healthCheck: &healthCheck{
//   		enabled: jsii.Boolean(true),
//   		healthyGrpcCodes: jsii.String("0-99"),
//   	},
//   	vpc: vpc,
//   })
//
type ApplicationProtocolVersion string

const (
	// GRPC.
	ApplicationProtocolVersion_GRPC ApplicationProtocolVersion = "GRPC"
	// HTTP1.
	ApplicationProtocolVersion_HTTP1 ApplicationProtocolVersion = "HTTP1"
	// HTTP2.
	ApplicationProtocolVersion_HTTP2 ApplicationProtocolVersion = "HTTP2"
)

// Define an Application Target Group.
//
// Example:
//   var vpc vpc
//
//
//   // Target group with duration-based stickiness with load-balancer generated cookie
//   tg1 := elbv2.NewApplicationTargetGroup(this, jsii.String("TG1"), &applicationTargetGroupProps{
//   	targetType: elbv2.targetType_INSTANCE,
//   	port: jsii.Number(80),
//   	stickinessCookieDuration: awscdk.Duration.minutes(jsii.Number(5)),
//   	vpc: vpc,
//   })
//
//   // Target group with application-based stickiness
//   tg2 := elbv2.NewApplicationTargetGroup(this, jsii.String("TG2"), &applicationTargetGroupProps{
//   	targetType: elbv2.*targetType_INSTANCE,
//   	port: jsii.Number(80),
//   	stickinessCookieDuration: awscdk.Duration.minutes(jsii.Number(5)),
//   	stickinessCookieName: jsii.String("MyDeliciousCookie"),
//   	vpc: vpc,
//   })
//
type ApplicationTargetGroup interface {
	TargetGroupBase
	IApplicationTargetGroup
	// Default port configured for members of this target group.
	DefaultPort() *float64
	// Full name of first load balancer.
	FirstLoadBalancerFullName() *string
	HealthCheck() *HealthCheck
	SetHealthCheck(val *HealthCheck)
	// A token representing a list of ARNs of the load balancers that route traffic to this target group.
	LoadBalancerArns() *string
	// List of constructs that need to be depended on to ensure the TargetGroup is associated to a load balancer.
	LoadBalancerAttached() constructs.IDependable
	// Configurable dependable with all resources that lead to load balancer attachment.
	LoadBalancerAttachedDependencies() constructs.DependencyGroup
	// The tree node.
	Node() constructs.Node
	// The ARN of the target group.
	TargetGroupArn() *string
	// The full name of the target group.
	TargetGroupFullName() *string
	// ARNs of load balancers load balancing to this TargetGroup.
	TargetGroupLoadBalancerArns() *[]*string
	// The name of the target group.
	TargetGroupName() *string
	// The types of the directly registered members of this target group.
	TargetType() TargetType
	SetTargetType(val TargetType)
	// Register the given load balancing target as part of this group.
	AddLoadBalancerTarget(props *LoadBalancerTargetProps)
	// Add a load balancing target to this target group.
	AddTarget(targets ...IApplicationLoadBalancerTarget)
	// Set/replace the target group's health check.
	ConfigureHealthCheck(healthCheck *HealthCheck)
	// Enable sticky routing via a cookie to members of this target group.
	//
	// Note: If the `cookieName` parameter is set, application-based stickiness will be applied,
	// otherwise it defaults to duration-based stickiness attributes (`lb_cookie`).
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/sticky-sessions.html
	//
	EnableCookieStickiness(duration awscdk.Duration, cookieName *string)
	// Return the given named metric for this Application Load Balancer Target Group.
	//
	// Returns the metric for this target group from the point of view of the first
	// load balancer load balancing to it. If you have multiple load balancers load
	// sending traffic to the same target group, you will have to override the dimensions
	// on this metric.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of healthy hosts in the target group.
	MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of HTTP 2xx/3xx/4xx/5xx response codes generated by all targets in this target group.
	//
	// This does not include any response codes generated by the load balancer.
	MetricHttpCodeTarget(code HttpCodeTarget, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of IPv6 requests received by the target group.
	MetricIpv6RequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of requests processed over IPv4 and IPv6.
	//
	// This count includes only the requests with a response generated by a target of the load balancer.
	MetricRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of requests received by each target in a target group.
	//
	// The only valid statistic is Sum. Note that this represents the average not the sum.
	MetricRequestCountPerTarget(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of connections that were not successfully established between the load balancer and target.
	MetricTargetConnectionErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time elapsed, in seconds, after the request leaves the load balancer until a response from the target is received.
	MetricTargetResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of TLS connections initiated by the load balancer that did not establish a session with the target.
	//
	// Possible causes include a mismatch of ciphers or protocols.
	MetricTargetTLSNegotiationErrorCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of unhealthy hosts in the target group.
	MetricUnhealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Register a connectable as a member of this target group.
	//
	// Don't call this directly. It will be called by load balancing targets.
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	RegisterListener(listener IApplicationListener, associatingConstruct constructs.IConstruct)
	// Set a non-standard attribute on the target group.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (a *jsiiProxy_ApplicationTargetGroup) AddLoadBalancerTarget(props *LoadBalancerTargetProps) {
	_jsii_.InvokeVoid(
		a,
		"addLoadBalancerTarget",
		[]interface{}{props},
	)
}

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

func (a *jsiiProxy_ApplicationTargetGroup) ConfigureHealthCheck(healthCheck *HealthCheck) {
	_jsii_.InvokeVoid(
		a,
		"configureHealthCheck",
		[]interface{}{healthCheck},
	)
}

func (a *jsiiProxy_ApplicationTargetGroup) EnableCookieStickiness(duration awscdk.Duration, cookieName *string) {
	_jsii_.InvokeVoid(
		a,
		"enableCookieStickiness",
		[]interface{}{duration, cookieName},
	)
}

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

func (a *jsiiProxy_ApplicationTargetGroup) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	_jsii_.InvokeVoid(
		a,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

func (a *jsiiProxy_ApplicationTargetGroup) RegisterListener(listener IApplicationListener, associatingConstruct constructs.IConstruct) {
	_jsii_.InvokeVoid(
		a,
		"registerListener",
		[]interface{}{listener, associatingConstruct},
	)
}

func (a *jsiiProxy_ApplicationTargetGroup) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		a,
		"setAttribute",
		[]interface{}{key, value},
	)
}

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
// Example:
//   var vpc vpc
//
//
//   // Target group with duration-based stickiness with load-balancer generated cookie
//   tg1 := elbv2.NewApplicationTargetGroup(this, jsii.String("TG1"), &applicationTargetGroupProps{
//   	targetType: elbv2.targetType_INSTANCE,
//   	port: jsii.Number(80),
//   	stickinessCookieDuration: awscdk.Duration.minutes(jsii.Number(5)),
//   	vpc: vpc,
//   })
//
//   // Target group with application-based stickiness
//   tg2 := elbv2.NewApplicationTargetGroup(this, jsii.String("TG2"), &applicationTargetGroupProps{
//   	targetType: elbv2.*targetType_INSTANCE,
//   	port: jsii.Number(80),
//   	stickinessCookieDuration: awscdk.Duration.minutes(jsii.Number(5)),
//   	stickinessCookieName: jsii.String("MyDeliciousCookie"),
//   	vpc: vpc,
//   })
//
type ApplicationTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	TargetType TargetType `field:"optional" json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The load balancing algorithm to select targets for routing requests.
	LoadBalancingAlgorithmType TargetGroupLoadBalancingAlgorithmType `field:"optional" json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// The port on which the listener listens for requests.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol version to use.
	ProtocolVersion ApplicationProtocolVersion `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The time period during which the load balancer sends a newly registered target a linearly increasing share of the traffic to the target group.
	//
	// The range is 30-900 seconds (15 minutes).
	SlowStart awscdk.Duration `field:"optional" json:"slowStart" yaml:"slowStart"`
	// The stickiness cookie expiration period.
	//
	// Setting this value enables load balancer stickiness.
	//
	// After this period, the cookie is considered stale. The minimum value is
	// 1 second and the maximum value is 7 days (604800 seconds).
	StickinessCookieDuration awscdk.Duration `field:"optional" json:"stickinessCookieDuration" yaml:"stickinessCookieDuration"`
	// The name of an application-based stickiness cookie.
	//
	// Names that start with the following prefixes are not allowed: AWSALB, AWSALBAPP,
	// and AWSALBTG; they're reserved for use by the load balancer.
	//
	// Note: `stickinessCookieName` parameter depends on the presence of `stickinessCookieDuration` parameter.
	// If `stickinessCookieDuration` is not set, `stickinessCookieName` will be omitted.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/sticky-sessions.html
	//
	StickinessCookieName *string `field:"optional" json:"stickinessCookieName" yaml:"stickinessCookieName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]IApplicationLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

// Options for `ListenerAction.authenciateOidc()`.
//
// Example:
//   var listener applicationListener
//   var myTargetGroup applicationTargetGroup
//
//
//   listener.addAction(jsii.String("DefaultAction"), &addApplicationActionProps{
//   	action: elbv2.listenerAction.authenticateOidc(&authenticateOidcOptions{
//   		authorizationEndpoint: jsii.String("https://example.com/openid"),
//   		// Other OIDC properties here
//   		clientId: jsii.String("..."),
//   		clientSecret: awscdk.SecretValue.secretsManager(jsii.String("...")),
//   		issuer: jsii.String("..."),
//   		tokenEndpoint: jsii.String("..."),
//   		userInfoEndpoint: jsii.String("..."),
//
//   		// Next
//   		next: elbv2.*listenerAction.forward([]iApplicationTargetGroup{
//   			myTargetGroup,
//   		}),
//   	}),
//   })
//
type AuthenticateOidcOptions struct {
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret.
	ClientSecret awscdk.SecretValue `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// What action to execute next.
	Next ListenerAction `field:"required" json:"next" yaml:"next"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams *map[string]*string `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated.
	OnUnauthenticatedRequest UnauthenticatedAction `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP.
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session.
	SessionTimeout awscdk.Duration `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

// Basic properties for an ApplicationListener.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var lb applicationLoadBalancer
//
//   listener := lb.addListener(jsii.String("listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &httpAlbIntegrationProps{
//   		parameterMapping: apigwv2.NewParameterMapping().appendHeader(jsii.String("header2"), apigwv2.mappingValue.requestHeader(jsii.String("header1"))).removeHeader(jsii.String("header1")),
//   	}),
//   })
//
type BaseApplicationListenerProps struct {
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses.
	//
	// See the `ListenerAction` class for all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction ListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]IApplicationTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Allow anyone to connect to the load balancer on the listener port.
	//
	// If this is specified, the load balancer will be opened up to anyone who can reach it.
	// For internal load balancers this is anyone in the same VPC. For public load
	// balancers, this is anyone on the internet.
	//
	// If you want to be more selective about who can access this load
	// balancer, set this to `false` and use the listener's `connections`
	// object to selectively grant access to the load balancer on the listener port.
	Open *bool `field:"optional" json:"open" yaml:"open"`
	// The port on which the listener listens for requests.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use.
	Protocol ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The security policy that defines which ciphers and protocols are supported.
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
}

// Basic properties for defining a rule on a listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationTargetGroup applicationTargetGroup
//   var listenerAction listenerAction
//   var listenerCondition listenerCondition
//
//   baseApplicationListenerRuleProps := &baseApplicationListenerRuleProps{
//   	priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	action: listenerAction,
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	targetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   }
//
type BaseApplicationListenerRuleProps struct {
	// Priority of the rule.
	//
	// The rule with the lowest priority will be used for every request.
	//
	// Priorities must be unique.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Action to perform when requests are received.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	Action ListenerAction `field:"optional" json:"action" yaml:"action"`
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Target groups to forward requests to.
	//
	// Only one of `action`, `fixedResponse`, `redirectResponse` or `targetGroups` can be specified.
	//
	// Implies a `forward` action.
	TargetGroups *[]IApplicationTargetGroup `field:"optional" json:"targetGroups" yaml:"targetGroups"`
}

// Base class for listeners.
type BaseListener interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
	// Validate this listener.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

// Returns true if the construct was created by CDK, and false otherwise.
func BaseListener_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseListener",
		"isOwnedResource",
		[]interface{}{construct},
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseListenerLookupOptions := &baseListenerLookupOptions{
//   	listenerPort: jsii.Number(123),
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: map[string]*string{
//   		"loadBalancerTagsKey": jsii.String("loadBalancerTags"),
//   	},
//   }
//
type BaseListenerLookupOptions struct {
	// Filter listeners by listener port.
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Base class for both Application and Network Load Balancers.
type BaseLoadBalancer interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN of this load balancer.
	//
	// Example value: `arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-internal-load-balancer/50dc6c495c0c9188`.
	LoadBalancerArn() *string
	// The canonical hosted zone ID of this load balancer.
	//
	// Example value: `Z2P70J7EXAMPLE`.
	LoadBalancerCanonicalHostedZoneId() *string
	// The DNS name of this load balancer.
	//
	// Example value: `my-load-balancer-424835706.us-west-2.elb.amazonaws.com`
	LoadBalancerDnsName() *string
	// The full name of this load balancer.
	//
	// Example value: `app/my-load-balancer/50dc6c495c0c9188`.
	LoadBalancerFullName() *string
	// The name of this load balancer.
	//
	// Example value: `my-load-balancer`.
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The VPC this load balancer has been created in.
	//
	// This property is always defined (not `null` or `undefined`) for sub-classes of `BaseLoadBalancer`.
	Vpc() awsec2.IVpc
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Enable access logging for this load balancer.
	//
	// A region must be specified on the stack containing the load balancer; you cannot enable logging on
	// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	// Remove an attribute from the load balancer.
	RemoveAttribute(key *string)
	// Set a non-standard attribute on the load balancer.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
	ValidateLoadBalancer() *[]*string
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

// Returns true if the construct was created by CDK, and false otherwise.
func BaseLoadBalancer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		"isOwnedResource",
		[]interface{}{construct},
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

func (b *jsiiProxy_BaseLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	_jsii_.InvokeVoid(
		b,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

func (b *jsiiProxy_BaseLoadBalancer) RemoveAttribute(key *string) {
	_jsii_.InvokeVoid(
		b,
		"removeAttribute",
		[]interface{}{key},
	)
}

func (b *jsiiProxy_BaseLoadBalancer) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		b,
		"setAttribute",
		[]interface{}{key, value},
	)
}

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

func (b *jsiiProxy_BaseLoadBalancer) ValidateLoadBalancer() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validateLoadBalancer",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for looking up load balancers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseLoadBalancerLookupOptions := &baseLoadBalancerLookupOptions{
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: map[string]*string{
//   		"loadBalancerTagsKey": jsii.String("loadBalancerTags"),
//   	},
//   }
//
type BaseLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Shared properties of both Application and Network Load Balancers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   baseLoadBalancerProps := &baseLoadBalancerProps{
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	deletionProtection: jsii.Boolean(false),
//   	internetFacing: jsii.Boolean(false),
//   	loadBalancerName: jsii.String("loadBalancerName"),
//   	vpcSubnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type BaseLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether the load balancer has an internet-routable address.
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// Which subnets place the load balancer in.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// Basic properties for a Network Listener.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addNetworkTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type BaseNetworkListenerProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Application-Layer Protocol Negotiation (ALPN) is a TLS extension that is sent on the initial TLS handshake hello messages.
	//
	// ALPN enables the application layer to negotiate which protocols should be used over a secure connection, such as HTTP/1 and HTTP/2.
	//
	// Can only be specified together with Protocol TLS.
	AlpnPolicy AlpnPolicy `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default Action of the load balancer,
	// including weighted forwarding. See the `NetworkListenerAction` class for
	// all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction NetworkListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]INetworkTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Protocol for listener, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// SSL Policy.
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
}

// Basic properties of both Application and Network Target Groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   baseTargetGroupProps := &baseTargetGroupProps{
//   	deregistrationDelay: cdk.duration.minutes(jsii.Number(30)),
//   	healthCheck: &healthCheck{
//   		enabled: jsii.Boolean(false),
//   		healthyGrpcCodes: jsii.String("healthyGrpcCodes"),
//   		healthyHttpCodes: jsii.String("healthyHttpCodes"),
//   		healthyThresholdCount: jsii.Number(123),
//   		interval: cdk.*duration.minutes(jsii.Number(30)),
//   		path: jsii.String("path"),
//   		port: jsii.String("port"),
//   		protocol: awscdk.Aws_elasticloadbalancingv2.protocol_HTTP,
//   		timeout: cdk.*duration.minutes(jsii.Number(30)),
//   		unhealthyThresholdCount: jsii.Number(123),
//   	},
//   	targetGroupName: jsii.String("targetGroupName"),
//   	targetType: awscdk.*Aws_elasticloadbalancingv2.targetType_INSTANCE,
//   	vpc: vpc,
//   }
//
type BaseTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	TargetType TargetType `field:"optional" json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::Listener`.
//
// Specifies a listener for an Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListener := awscdk.Aws_elasticloadbalancingv2.NewCfnListener(this, jsii.String("MyCfnListener"), &cfnListenerProps{
//   	defaultActions: []interface{}{
//   		&actionProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			authenticateCognitoConfig: &authenticateCognitoConfigProperty{
//   				userPoolArn: jsii.String("userPoolArn"),
//   				userPoolClientId: jsii.String("userPoolClientId"),
//   				userPoolDomain: jsii.String("userPoolDomain"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.String("sessionTimeout"),
//   			},
//   			authenticateOidcConfig: &authenticateOidcConfigProperty{
//   				authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				clientId: jsii.String("clientId"),
//   				issuer: jsii.String("issuer"),
//   				tokenEndpoint: jsii.String("tokenEndpoint"),
//   				userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				clientSecret: jsii.String("clientSecret"),
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.String("sessionTimeout"),
//   				useExistingClientSecret: jsii.Boolean(false),
//   			},
//   			fixedResponseConfig: &fixedResponseConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentType: jsii.String("contentType"),
//   				messageBody: jsii.String("messageBody"),
//   			},
//   			forwardConfig: &forwardConfigProperty{
//   				targetGroups: []interface{}{
//   					&targetGroupTupleProperty{
//   						targetGroupArn: jsii.String("targetGroupArn"),
//   						weight: jsii.Number(123),
//   					},
//   				},
//   				targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   					durationSeconds: jsii.Number(123),
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   			order: jsii.Number(123),
//   			redirectConfig: &redirectConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				host: jsii.String("host"),
//   				path: jsii.String("path"),
//   				port: jsii.String("port"),
//   				protocol: jsii.String("protocol"),
//   				query: jsii.String("query"),
//   			},
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//
//   	// the properties below are optional
//   	alpnPolicy: []*string{
//   		jsii.String("alpnPolicy"),
//   	},
//   	certificates: []interface{}{
//   		&certificateProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   		},
//   	},
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	sslPolicy: jsii.String("sslPolicy"),
//   })
//
type CfnListener interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
	AlpnPolicy() *[]*string
	SetAlpnPolicy(val *[]*string)
	// The Amazon Resource Name (ARN) of the listener.
	AttrListenerArn() *string
	// The default SSL server certificate for a secure listener.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	//
	// To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html) .
	Certificates() interface{}
	SetCertificates(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The actions for the default rule. You cannot define a condition for a default rule.
	//
	// To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html) .
	DefaultActions() interface{}
	SetDefaultActions(val interface{})
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn() *string
	SetLoadBalancerArn(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The port on which the load balancer is listening.
	//
	// You cannot specify a port for a Gateway Load Balancer.
	Port() *float64
	SetPort(val *float64)
	// The protocol for connections from clients to the load balancer.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
	Protocol() *string
	SetProtocol(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
	//
	// For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide* .
	SslPolicy() *string
	SetSslPolicy(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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

func (j *jsiiProxy_CfnListener) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (c *jsiiProxy_CfnListener) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnListener) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnListener) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnListener) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnListener) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnListener) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnListener) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	authenticateCognitoConfig: &authenticateCognitoConfigProperty{
//   		userPoolArn: jsii.String("userPoolArn"),
//   		userPoolClientId: jsii.String("userPoolClientId"),
//   		userPoolDomain: jsii.String("userPoolDomain"),
//
//   		// the properties below are optional
//   		authenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		scope: jsii.String("scope"),
//   		sessionCookieName: jsii.String("sessionCookieName"),
//   		sessionTimeout: jsii.String("sessionTimeout"),
//   	},
//   	authenticateOidcConfig: &authenticateOidcConfigProperty{
//   		authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		clientId: jsii.String("clientId"),
//   		issuer: jsii.String("issuer"),
//   		tokenEndpoint: jsii.String("tokenEndpoint"),
//   		userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   		// the properties below are optional
//   		authenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		clientSecret: jsii.String("clientSecret"),
//   		onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		scope: jsii.String("scope"),
//   		sessionCookieName: jsii.String("sessionCookieName"),
//   		sessionTimeout: jsii.String("sessionTimeout"),
//   		useExistingClientSecret: jsii.Boolean(false),
//   	},
//   	fixedResponseConfig: &fixedResponseConfigProperty{
//   		statusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		contentType: jsii.String("contentType"),
//   		messageBody: jsii.String("messageBody"),
//   	},
//   	forwardConfig: &forwardConfigProperty{
//   		targetGroups: []interface{}{
//   			&targetGroupTupleProperty{
//   				targetGroupArn: jsii.String("targetGroupArn"),
//   				weight: jsii.Number(123),
//   			},
//   		},
//   		targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   			durationSeconds: jsii.Number(123),
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	order: jsii.Number(123),
//   	redirectConfig: &redirectConfigProperty{
//   		statusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		host: jsii.String("host"),
//   		path: jsii.String("path"),
//   		port: jsii.String("port"),
//   		protocol: jsii.String("protocol"),
//   		query: jsii.String("query"),
//   	},
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   }
//
type CfnListener_ActionProperty struct {
	// The type of action.
	Type *string `field:"required" json:"type" yaml:"type"`
	// [HTTPS listeners] Information for using Amazon Cognito to authenticate users.
	//
	// Specify only when `Type` is `authenticate-cognito` .
	AuthenticateCognitoConfig interface{} `field:"optional" json:"authenticateCognitoConfig" yaml:"authenticateCognitoConfig"`
	// [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC).
	//
	// Specify only when `Type` is `authenticate-oidc` .
	AuthenticateOidcConfig interface{} `field:"optional" json:"authenticateOidcConfig" yaml:"authenticateOidcConfig"`
	// [Application Load Balancer] Information for creating an action that returns a custom HTTP response.
	//
	// Specify only when `Type` is `fixed-response` .
	FixedResponseConfig interface{} `field:"optional" json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Information for creating an action that distributes requests among one or more target groups.
	//
	// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
	ForwardConfig interface{} `field:"optional" json:"forwardConfig" yaml:"forwardConfig"`
	// The order for the action.
	//
	// This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// [Application Load Balancer] Information for creating a redirect action.
	//
	// Specify only when `Type` is `redirect` .
	RedirectConfig interface{} `field:"optional" json:"redirectConfig" yaml:"redirectConfig"`
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Specify only when `Type` is `forward` and you want to route to a single target group. To route to one or more target groups, use `ForwardConfig` instead.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

// Specifies information required when integrating with Amazon Cognito to authenticate users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticateCognitoConfigProperty := &authenticateCognitoConfigProperty{
//   	userPoolArn: jsii.String("userPoolArn"),
//   	userPoolClientId: jsii.String("userPoolClientId"),
//   	userPoolDomain: jsii.String("userPoolDomain"),
//
//   	// the properties below are optional
//   	authenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   	scope: jsii.String("scope"),
//   	sessionCookieName: jsii.String("sessionCookieName"),
//   	sessionTimeout: jsii.String("sessionTimeout"),
//   }
//
type CfnListener_AuthenticateCognitoConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// The ID of the Amazon Cognito user pool client.
	UserPoolClientId *string `field:"required" json:"userPoolClientId" yaml:"userPoolClientId"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	UserPoolDomain *string `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *string `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

// Specifies information required using an identity provide (IdP) that is compliant with OpenID Connect (OIDC) to authenticate users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticateOidcConfigProperty := &authenticateOidcConfigProperty{
//   	authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	clientId: jsii.String("clientId"),
//   	issuer: jsii.String("issuer"),
//   	tokenEndpoint: jsii.String("tokenEndpoint"),
//   	userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   	// the properties below are optional
//   	authenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	clientSecret: jsii.String("clientSecret"),
//   	onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   	scope: jsii.String("scope"),
//   	sessionCookieName: jsii.String("sessionCookieName"),
//   	sessionTimeout: jsii.String("sessionTimeout"),
//   	useExistingClientSecret: jsii.Boolean(false),
//   }
//
type CfnListener_AuthenticateOidcConfigProperty struct {
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set `UseExistingClientSecret` to true.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *string `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// `CfnListener.AuthenticateOidcConfigProperty.UseExistingClientSecret`.
	UseExistingClientSecret interface{} `field:"optional" json:"useExistingClientSecret" yaml:"useExistingClientSecret"`
}

// Specifies an SSL server certificate to use as the default certificate for a secure listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateProperty := &certificateProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   }
//
type CfnListener_CertificateProperty struct {
	// The Amazon Resource Name (ARN) of the certificate.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
}

// Specifies information required when returning a custom HTTP response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fixedResponseConfigProperty := &fixedResponseConfigProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	messageBody: jsii.String("messageBody"),
//   }
//
type CfnListener_FixedResponseConfigProperty struct {
	// The HTTP response code (2XX, 4XX, or 5XX).
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The content type.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The message.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

// Information for creating an action that distributes requests among one or more target groups.
//
// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forwardConfigProperty := &forwardConfigProperty{
//   	targetGroups: []interface{}{
//   		&targetGroupTupleProperty{
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   		durationSeconds: jsii.Number(123),
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnListener_ForwardConfigProperty struct {
	// Information about how traffic will be distributed between multiple target groups in a forward rule.
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// Information about the target group stickiness for a rule.
	TargetGroupStickinessConfig interface{} `field:"optional" json:"targetGroupStickinessConfig" yaml:"targetGroupStickinessConfig"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectConfigProperty := &redirectConfigProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	host: jsii.String("host"),
//   	path: jsii.String("path"),
//   	port: jsii.String("port"),
//   	protocol: jsii.String("protocol"),
//   	query: jsii.String("query"),
//   }
//
type CfnListener_RedirectConfigProperty struct {
	// The HTTP redirect code.
	//
	// The redirect is either permanent (HTTP 301) or temporary (HTTP 302).
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added. You can specify any of the reserved keywords.
	Query *string `field:"optional" json:"query" yaml:"query"`
}

// Information about the target group stickiness for a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupStickinessConfigProperty := &targetGroupStickinessConfigProperty{
//   	durationSeconds: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnListener_TargetGroupStickinessConfigProperty struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target group.
	//
	// The range is 1-604800 seconds (7 days).
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether target group stickiness is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// Information about how traffic will be distributed between multiple target groups in a forward rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupTupleProperty := &targetGroupTupleProperty{
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   	weight: jsii.Number(123),
//   }
//
type CfnListener_TargetGroupTupleProperty struct {
	// The Amazon Resource Name (ARN) of the target group.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
	// The weight.
	//
	// The range is 0 to 999.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::ListenerCertificate`.
//
// Specifies an SSL server certificate to add to the certificate list for an HTTPS or TLS listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerCertificate := awscdk.Aws_elasticloadbalancingv2.NewCfnListenerCertificate(this, jsii.String("MyCfnListenerCertificate"), &cfnListenerCertificateProps{
//   	certificates: []interface{}{
//   		&certificateProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   		},
//   	},
//   	listenerArn: jsii.String("listenerArn"),
//   })
//
type CfnListenerCertificate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The certificate.
	//
	// You can specify one certificate per resource.
	Certificates() interface{}
	SetCertificates(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn() *string
	SetListenerArn(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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

func (j *jsiiProxy_CfnListenerCertificate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (c *jsiiProxy_CfnListenerCertificate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnListenerCertificate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnListenerCertificate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnListenerCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnListenerCertificate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnListenerCertificate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnListenerCertificate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnListenerCertificate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateProperty := &certificateProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   }
//
type CfnListenerCertificate_CertificateProperty struct {
	// The Amazon Resource Name (ARN) of the certificate.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
}

// Properties for defining a `CfnListenerCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerCertificateProps := &cfnListenerCertificateProps{
//   	certificates: []interface{}{
//   		&certificateProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   		},
//   	},
//   	listenerArn: jsii.String("listenerArn"),
//   }
//
type CfnListenerCertificateProps struct {
	// The certificate.
	//
	// You can specify one certificate per resource.
	Certificates interface{} `field:"required" json:"certificates" yaml:"certificates"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
}

// Properties for defining a `CfnListener`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerProps := &cfnListenerProps{
//   	defaultActions: []interface{}{
//   		&actionProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			authenticateCognitoConfig: &authenticateCognitoConfigProperty{
//   				userPoolArn: jsii.String("userPoolArn"),
//   				userPoolClientId: jsii.String("userPoolClientId"),
//   				userPoolDomain: jsii.String("userPoolDomain"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.String("sessionTimeout"),
//   			},
//   			authenticateOidcConfig: &authenticateOidcConfigProperty{
//   				authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				clientId: jsii.String("clientId"),
//   				issuer: jsii.String("issuer"),
//   				tokenEndpoint: jsii.String("tokenEndpoint"),
//   				userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				clientSecret: jsii.String("clientSecret"),
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.String("sessionTimeout"),
//   				useExistingClientSecret: jsii.Boolean(false),
//   			},
//   			fixedResponseConfig: &fixedResponseConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentType: jsii.String("contentType"),
//   				messageBody: jsii.String("messageBody"),
//   			},
//   			forwardConfig: &forwardConfigProperty{
//   				targetGroups: []interface{}{
//   					&targetGroupTupleProperty{
//   						targetGroupArn: jsii.String("targetGroupArn"),
//   						weight: jsii.Number(123),
//   					},
//   				},
//   				targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   					durationSeconds: jsii.Number(123),
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   			order: jsii.Number(123),
//   			redirectConfig: &redirectConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				host: jsii.String("host"),
//   				path: jsii.String("path"),
//   				port: jsii.String("port"),
//   				protocol: jsii.String("protocol"),
//   				query: jsii.String("query"),
//   			},
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//
//   	// the properties below are optional
//   	alpnPolicy: []*string{
//   		jsii.String("alpnPolicy"),
//   	},
//   	certificates: []interface{}{
//   		&certificateProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   		},
//   	},
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	sslPolicy: jsii.String("sslPolicy"),
//   }
//
type CfnListenerProps struct {
	// The actions for the default rule. You cannot define a condition for a default rule.
	//
	// To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html) .
	DefaultActions interface{} `field:"required" json:"defaultActions" yaml:"defaultActions"`
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
	AlpnPolicy *[]*string `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// The default SSL server certificate for a secure listener.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	//
	// To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html) .
	Certificates interface{} `field:"optional" json:"certificates" yaml:"certificates"`
	// The port on which the load balancer is listening.
	//
	// You cannot specify a port for a Gateway Load Balancer.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol for connections from clients to the load balancer.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
	//
	// For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide* .
	SslPolicy *string `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::ListenerRule`.
//
// Specifies a listener rule. The listener must be associated with an Application Load Balancer. Each rule consists of a priority, one or more actions, and one or more conditions.
//
// For more information, see [Quotas for your Application Load Balancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-limits.html) in the *User Guide for Application Load Balancers* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerRule := awscdk.Aws_elasticloadbalancingv2.NewCfnListenerRule(this, jsii.String("MyCfnListenerRule"), &cfnListenerRuleProps{
//   	actions: []interface{}{
//   		&actionProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			authenticateCognitoConfig: &authenticateCognitoConfigProperty{
//   				userPoolArn: jsii.String("userPoolArn"),
//   				userPoolClientId: jsii.String("userPoolClientId"),
//   				userPoolDomain: jsii.String("userPoolDomain"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.Number(123),
//   			},
//   			authenticateOidcConfig: &authenticateOidcConfigProperty{
//   				authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				clientId: jsii.String("clientId"),
//   				issuer: jsii.String("issuer"),
//   				tokenEndpoint: jsii.String("tokenEndpoint"),
//   				userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				clientSecret: jsii.String("clientSecret"),
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.Number(123),
//   				useExistingClientSecret: jsii.Boolean(false),
//   			},
//   			fixedResponseConfig: &fixedResponseConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentType: jsii.String("contentType"),
//   				messageBody: jsii.String("messageBody"),
//   			},
//   			forwardConfig: &forwardConfigProperty{
//   				targetGroups: []interface{}{
//   					&targetGroupTupleProperty{
//   						targetGroupArn: jsii.String("targetGroupArn"),
//   						weight: jsii.Number(123),
//   					},
//   				},
//   				targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   					durationSeconds: jsii.Number(123),
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   			order: jsii.Number(123),
//   			redirectConfig: &redirectConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				host: jsii.String("host"),
//   				path: jsii.String("path"),
//   				port: jsii.String("port"),
//   				protocol: jsii.String("protocol"),
//   				query: jsii.String("query"),
//   			},
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	conditions: []interface{}{
//   		&ruleConditionProperty{
//   			field: jsii.String("field"),
//   			hostHeaderConfig: &hostHeaderConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			httpHeaderConfig: &httpHeaderConfigProperty{
//   				httpHeaderName: jsii.String("httpHeaderName"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			httpRequestMethodConfig: &httpRequestMethodConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			pathPatternConfig: &pathPatternConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			queryStringConfig: &queryStringConfigProperty{
//   				values: []interface{}{
//   					&queryStringKeyValueProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			sourceIpConfig: &sourceIpConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	listenerArn: jsii.String("listenerArn"),
//   	priority: jsii.Number(123),
//   })
//
type CfnListenerRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The actions.
	//
	// The rule must include exactly one of the following types of actions: `forward` , `fixed-response` , or `redirect` , and it must be the last action to be performed. If the rule is for an HTTPS listener, it can also optionally include an authentication action.
	Actions() interface{}
	SetActions(val interface{})
	// Indicates whether this is the default rule.
	AttrIsDefault() awscdk.IResolvable
	// The Amazon Resource Name (ARN) of the rule.
	AttrRuleArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The conditions.
	//
	// The rule can optionally include up to one of each of the following conditions: `http-request-method` , `host-header` , `path-pattern` , and `source-ip` . A rule can also optionally include one or more of each of the following conditions: `http-header` and `query-string` .
	Conditions() interface{}
	SetConditions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn() *string
	SetListenerArn(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The rule priority. A listener can't have multiple rules with the same priority.
	//
	// If you try to reorder rules by updating their priorities, do not specify a new priority if an existing rule already uses this priority, as this can cause an error. If you need to reuse a priority with a different rule, you must remove it as a priority first, and then specify it in a subsequent update.
	Priority() *float64
	SetPriority(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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

func (j *jsiiProxy_CfnListenerRule) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (c *jsiiProxy_CfnListenerRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnListenerRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnListenerRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnListenerRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnListenerRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnListenerRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnListenerRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnListenerRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	authenticateCognitoConfig: &authenticateCognitoConfigProperty{
//   		userPoolArn: jsii.String("userPoolArn"),
//   		userPoolClientId: jsii.String("userPoolClientId"),
//   		userPoolDomain: jsii.String("userPoolDomain"),
//
//   		// the properties below are optional
//   		authenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		scope: jsii.String("scope"),
//   		sessionCookieName: jsii.String("sessionCookieName"),
//   		sessionTimeout: jsii.Number(123),
//   	},
//   	authenticateOidcConfig: &authenticateOidcConfigProperty{
//   		authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		clientId: jsii.String("clientId"),
//   		issuer: jsii.String("issuer"),
//   		tokenEndpoint: jsii.String("tokenEndpoint"),
//   		userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   		// the properties below are optional
//   		authenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		clientSecret: jsii.String("clientSecret"),
//   		onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   		scope: jsii.String("scope"),
//   		sessionCookieName: jsii.String("sessionCookieName"),
//   		sessionTimeout: jsii.Number(123),
//   		useExistingClientSecret: jsii.Boolean(false),
//   	},
//   	fixedResponseConfig: &fixedResponseConfigProperty{
//   		statusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		contentType: jsii.String("contentType"),
//   		messageBody: jsii.String("messageBody"),
//   	},
//   	forwardConfig: &forwardConfigProperty{
//   		targetGroups: []interface{}{
//   			&targetGroupTupleProperty{
//   				targetGroupArn: jsii.String("targetGroupArn"),
//   				weight: jsii.Number(123),
//   			},
//   		},
//   		targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   			durationSeconds: jsii.Number(123),
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	order: jsii.Number(123),
//   	redirectConfig: &redirectConfigProperty{
//   		statusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		host: jsii.String("host"),
//   		path: jsii.String("path"),
//   		port: jsii.String("port"),
//   		protocol: jsii.String("protocol"),
//   		query: jsii.String("query"),
//   	},
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   }
//
type CfnListenerRule_ActionProperty struct {
	// The type of action.
	Type *string `field:"required" json:"type" yaml:"type"`
	// [HTTPS listeners] Information for using Amazon Cognito to authenticate users.
	//
	// Specify only when `Type` is `authenticate-cognito` .
	AuthenticateCognitoConfig interface{} `field:"optional" json:"authenticateCognitoConfig" yaml:"authenticateCognitoConfig"`
	// [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC).
	//
	// Specify only when `Type` is `authenticate-oidc` .
	AuthenticateOidcConfig interface{} `field:"optional" json:"authenticateOidcConfig" yaml:"authenticateOidcConfig"`
	// [Application Load Balancer] Information for creating an action that returns a custom HTTP response.
	//
	// Specify only when `Type` is `fixed-response` .
	FixedResponseConfig interface{} `field:"optional" json:"fixedResponseConfig" yaml:"fixedResponseConfig"`
	// Information for creating an action that distributes requests among one or more target groups.
	//
	// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
	ForwardConfig interface{} `field:"optional" json:"forwardConfig" yaml:"forwardConfig"`
	// The order for the action.
	//
	// This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// [Application Load Balancer] Information for creating a redirect action.
	//
	// Specify only when `Type` is `redirect` .
	RedirectConfig interface{} `field:"optional" json:"redirectConfig" yaml:"redirectConfig"`
	// The Amazon Resource Name (ARN) of the target group.
	//
	// Specify only when `Type` is `forward` and you want to route to a single target group. To route to one or more target groups, use `ForwardConfig` instead.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

// Specifies information required when integrating with Amazon Cognito to authenticate users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticateCognitoConfigProperty := &authenticateCognitoConfigProperty{
//   	userPoolArn: jsii.String("userPoolArn"),
//   	userPoolClientId: jsii.String("userPoolClientId"),
//   	userPoolDomain: jsii.String("userPoolDomain"),
//
//   	// the properties below are optional
//   	authenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   	scope: jsii.String("scope"),
//   	sessionCookieName: jsii.String("sessionCookieName"),
//   	sessionTimeout: jsii.Number(123),
//   }
//
type CfnListenerRule_AuthenticateCognitoConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// The ID of the Amazon Cognito user pool client.
	UserPoolClientId *string `field:"required" json:"userPoolClientId" yaml:"userPoolClientId"`
	// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
	UserPoolDomain *string `field:"required" json:"userPoolDomain" yaml:"userPoolDomain"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
}

// Specifies information required using an identity provide (IdP) that is compliant with OpenID Connect (OIDC) to authenticate users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authenticateOidcConfigProperty := &authenticateOidcConfigProperty{
//   	authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	clientId: jsii.String("clientId"),
//   	issuer: jsii.String("issuer"),
//   	tokenEndpoint: jsii.String("tokenEndpoint"),
//   	userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   	// the properties below are optional
//   	authenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	clientSecret: jsii.String("clientSecret"),
//   	onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   	scope: jsii.String("scope"),
//   	sessionCookieName: jsii.String("sessionCookieName"),
//   	sessionTimeout: jsii.Number(123),
//   	useExistingClientSecret: jsii.Boolean(false),
//   }
//
type CfnListenerRule_AuthenticateOidcConfigProperty struct {
	// The authorization endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The OIDC issuer identifier of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The token endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	//
	// This must be a full URL, including the HTTPS protocol, the domain, and the path.
	UserInfoEndpoint *string `field:"required" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
	// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set `UseExistingClientSecret` to true.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The behavior if the user is not authenticated. The following are possible values:.
	//
	// - deny `` - Return an HTTP 401 Unauthorized error.
	// - allow `` - Allow the request to be forwarded to the target.
	// - authenticate `` - Redirect the request to the IdP authorization endpoint. This is the default value.
	OnUnauthenticatedRequest *string `field:"optional" json:"onUnauthenticatedRequest" yaml:"onUnauthenticatedRequest"`
	// The set of user claims to be requested from the IdP. The default is `openid` .
	//
	// To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The name of the cookie used to maintain session information.
	//
	// The default is AWSELBAuthSessionCookie.
	SessionCookieName *string `field:"optional" json:"sessionCookieName" yaml:"sessionCookieName"`
	// The maximum duration of the authentication session, in seconds.
	//
	// The default is 604800 seconds (7 days).
	SessionTimeout *float64 `field:"optional" json:"sessionTimeout" yaml:"sessionTimeout"`
	// Indicates whether to use the existing client secret when modifying a rule.
	//
	// If you are creating a rule, you can omit this parameter or set it to false.
	UseExistingClientSecret interface{} `field:"optional" json:"useExistingClientSecret" yaml:"useExistingClientSecret"`
}

// Specifies information required when returning a custom HTTP response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fixedResponseConfigProperty := &fixedResponseConfigProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	contentType: jsii.String("contentType"),
//   	messageBody: jsii.String("messageBody"),
//   }
//
type CfnListenerRule_FixedResponseConfigProperty struct {
	// The HTTP response code (2XX, 4XX, or 5XX).
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The content type.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The message.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

// Information for creating an action that distributes requests among one or more target groups.
//
// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forwardConfigProperty := &forwardConfigProperty{
//   	targetGroups: []interface{}{
//   		&targetGroupTupleProperty{
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   		durationSeconds: jsii.Number(123),
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnListenerRule_ForwardConfigProperty struct {
	// Information about how traffic will be distributed between multiple target groups in a forward rule.
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// Information about the target group stickiness for a rule.
	TargetGroupStickinessConfig interface{} `field:"optional" json:"targetGroupStickinessConfig" yaml:"targetGroupStickinessConfig"`
}

// Information about a host header condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostHeaderConfigProperty := &hostHeaderConfigProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_HostHeaderConfigProperty struct {
	// One or more host names.
	//
	// The maximum size of each name is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the host name.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Information about an HTTP header condition.
//
// There is a set of standard HTTP header fields. You can also define custom HTTP header fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpHeaderConfigProperty := &httpHeaderConfigProperty{
//   	httpHeaderName: jsii.String("httpHeaderName"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_HttpHeaderConfigProperty struct {
	// The name of the HTTP header field.
	//
	// The maximum size is 40 characters. The header name is case insensitive. The allowed characters are specified by RFC 7230. Wildcards are not supported.
	HttpHeaderName *string `field:"optional" json:"httpHeaderName" yaml:"httpHeaderName"`
	// One or more strings to compare against the value of the HTTP header.
	//
	// The maximum size of each string is 128 characters. The comparison strings are case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If the same header appears multiple times in the request, we search them in order until a match is found.
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the value of the HTTP header. To require that all of the strings are a match, create one condition per string.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Information about an HTTP method condition.
//
// HTTP defines a set of request methods, also referred to as HTTP verbs. For more information, see the [HTTP Method Registry](https://docs.aws.amazon.com/https://www.iana.org/assignments/http-methods/http-methods.xhtml) . You can also define custom HTTP methods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRequestMethodConfigProperty := &httpRequestMethodConfigProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_HttpRequestMethodConfigProperty struct {
	// The name of the request method.
	//
	// The maximum size is 40 characters. The allowed characters are A-Z, hyphen (-), and underscore (_). The comparison is case sensitive. Wildcards are not supported; therefore, the method name must be an exact match.
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the HTTP request method. We recommend that you route GET and HEAD requests in the same way, because the response to a HEAD request may be cached.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Information about a path pattern condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathPatternConfigProperty := &pathPatternConfigProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_PathPatternConfigProperty struct {
	// The path patterns to compare against the request URL.
	//
	// The maximum size of each string is 128 characters. The comparison is case sensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If you specify multiple strings, the condition is satisfied if one of them matches the request URL. The path pattern is compared only to the path of the URL, not to its query string.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Information about a query string condition.
//
// The query string component of a URI starts after the first '?' character and is terminated by either a '#' character or the end of the URI. A typical query string contains key/value pairs separated by '&' characters. The allowed characters are specified by RFC 3986. Any character can be percentage encoded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringConfigProperty := &queryStringConfigProperty{
//   	values: []interface{}{
//   		&queryStringKeyValueProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnListenerRule_QueryStringConfigProperty struct {
	// One or more key/value pairs or values to find in the query string.
	//
	// The maximum size of each string is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, you must escape these characters in `Values` using a '\' character.
	//
	// If you specify multiple key/value pairs or values, the condition is satisfied if one of them is found in the query string.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

// Information about a key/value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringKeyValueProperty := &queryStringKeyValueProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnListenerRule_QueryStringKeyValueProperty struct {
	// The key.
	//
	// You can omit the key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value.
	Value *string `field:"optional" json:"value" yaml:"value"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectConfigProperty := &redirectConfigProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	host: jsii.String("host"),
//   	path: jsii.String("path"),
//   	port: jsii.String("port"),
//   	protocol: jsii.String("protocol"),
//   	query: jsii.String("query"),
//   }
//
type CfnListenerRule_RedirectConfigProperty struct {
	// The HTTP redirect code.
	//
	// The redirect is either permanent (HTTP 301) or temporary (HTTP 302).
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added. You can specify any of the reserved keywords.
	Query *string `field:"optional" json:"query" yaml:"query"`
}

// Specifies a condition for a listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleConditionProperty := &ruleConditionProperty{
//   	field: jsii.String("field"),
//   	hostHeaderConfig: &hostHeaderConfigProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	httpHeaderConfig: &httpHeaderConfigProperty{
//   		httpHeaderName: jsii.String("httpHeaderName"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	httpRequestMethodConfig: &httpRequestMethodConfigProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	pathPatternConfig: &pathPatternConfigProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	queryStringConfig: &queryStringConfigProperty{
//   		values: []interface{}{
//   			&queryStringKeyValueProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	sourceIpConfig: &sourceIpConfigProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_RuleConditionProperty struct {
	// The field in the HTTP request. The following are the possible values:.
	//
	// - `http-header`
	// - `http-request-method`
	// - `host-header`
	// - `path-pattern`
	// - `query-string`
	// - `source-ip`.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Information for a host header condition.
	//
	// Specify only when `Field` is `host-header` .
	HostHeaderConfig interface{} `field:"optional" json:"hostHeaderConfig" yaml:"hostHeaderConfig"`
	// Information for an HTTP header condition.
	//
	// Specify only when `Field` is `http-header` .
	HttpHeaderConfig interface{} `field:"optional" json:"httpHeaderConfig" yaml:"httpHeaderConfig"`
	// Information for an HTTP method condition.
	//
	// Specify only when `Field` is `http-request-method` .
	HttpRequestMethodConfig interface{} `field:"optional" json:"httpRequestMethodConfig" yaml:"httpRequestMethodConfig"`
	// Information for a path pattern condition.
	//
	// Specify only when `Field` is `path-pattern` .
	PathPatternConfig interface{} `field:"optional" json:"pathPatternConfig" yaml:"pathPatternConfig"`
	// Information for a query string condition.
	//
	// Specify only when `Field` is `query-string` .
	QueryStringConfig interface{} `field:"optional" json:"queryStringConfig" yaml:"queryStringConfig"`
	// Information for a source IP condition.
	//
	// Specify only when `Field` is `source-ip` .
	SourceIpConfig interface{} `field:"optional" json:"sourceIpConfig" yaml:"sourceIpConfig"`
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
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Information about a source IP condition.
//
// You can use this condition to route based on the IP address of the source that connects to the load balancer. If a client is behind a proxy, this is the IP address of the proxy not the IP address of the client.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceIpConfigProperty := &sourceIpConfigProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_SourceIpConfigProperty struct {
	// The source IP addresses, in CIDR format. You can use both IPv4 and IPv6 addresses. Wildcards are not supported.
	//
	// If you specify multiple addresses, the condition is satisfied if the source IP address of the request matches one of the CIDR blocks. This condition is not satisfied by the addresses in the X-Forwarded-For header.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Information about the target group stickiness for a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupStickinessConfigProperty := &targetGroupStickinessConfigProperty{
//   	durationSeconds: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnListenerRule_TargetGroupStickinessConfigProperty struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target group.
	//
	// The range is 1-604800 seconds (7 days).
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether target group stickiness is enabled.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

// Information about how traffic will be distributed between multiple target groups in a forward rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupTupleProperty := &targetGroupTupleProperty{
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   	weight: jsii.Number(123),
//   }
//
type CfnListenerRule_TargetGroupTupleProperty struct {
	// The Amazon Resource Name (ARN) of the target group.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
	// The weight.
	//
	// The range is 0 to 999.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// Properties for defining a `CfnListenerRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListenerRuleProps := &cfnListenerRuleProps{
//   	actions: []interface{}{
//   		&actionProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			authenticateCognitoConfig: &authenticateCognitoConfigProperty{
//   				userPoolArn: jsii.String("userPoolArn"),
//   				userPoolClientId: jsii.String("userPoolClientId"),
//   				userPoolDomain: jsii.String("userPoolDomain"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.Number(123),
//   			},
//   			authenticateOidcConfig: &authenticateOidcConfigProperty{
//   				authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   				clientId: jsii.String("clientId"),
//   				issuer: jsii.String("issuer"),
//   				tokenEndpoint: jsii.String("tokenEndpoint"),
//   				userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   				// the properties below are optional
//   				authenticationRequestExtraParams: map[string]*string{
//   					"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   				},
//   				clientSecret: jsii.String("clientSecret"),
//   				onUnauthenticatedRequest: jsii.String("onUnauthenticatedRequest"),
//   				scope: jsii.String("scope"),
//   				sessionCookieName: jsii.String("sessionCookieName"),
//   				sessionTimeout: jsii.Number(123),
//   				useExistingClientSecret: jsii.Boolean(false),
//   			},
//   			fixedResponseConfig: &fixedResponseConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentType: jsii.String("contentType"),
//   				messageBody: jsii.String("messageBody"),
//   			},
//   			forwardConfig: &forwardConfigProperty{
//   				targetGroups: []interface{}{
//   					&targetGroupTupleProperty{
//   						targetGroupArn: jsii.String("targetGroupArn"),
//   						weight: jsii.Number(123),
//   					},
//   				},
//   				targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   					durationSeconds: jsii.Number(123),
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   			order: jsii.Number(123),
//   			redirectConfig: &redirectConfigProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				host: jsii.String("host"),
//   				path: jsii.String("path"),
//   				port: jsii.String("port"),
//   				protocol: jsii.String("protocol"),
//   				query: jsii.String("query"),
//   			},
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	conditions: []interface{}{
//   		&ruleConditionProperty{
//   			field: jsii.String("field"),
//   			hostHeaderConfig: &hostHeaderConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			httpHeaderConfig: &httpHeaderConfigProperty{
//   				httpHeaderName: jsii.String("httpHeaderName"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			httpRequestMethodConfig: &httpRequestMethodConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			pathPatternConfig: &pathPatternConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			queryStringConfig: &queryStringConfigProperty{
//   				values: []interface{}{
//   					&queryStringKeyValueProperty{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			sourceIpConfig: &sourceIpConfigProperty{
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	listenerArn: jsii.String("listenerArn"),
//   	priority: jsii.Number(123),
//   }
//
type CfnListenerRuleProps struct {
	// The actions.
	//
	// The rule must include exactly one of the following types of actions: `forward` , `fixed-response` , or `redirect` , and it must be the last action to be performed. If the rule is for an HTTPS listener, it can also optionally include an authentication action.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The conditions.
	//
	// The rule can optionally include up to one of each of the following conditions: `http-request-method` , `host-header` , `path-pattern` , and `source-ip` . A rule can also optionally include one or more of each of the following conditions: `http-header` and `query-string` .
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// The rule priority. A listener can't have multiple rules with the same priority.
	//
	// If you try to reorder rules by updating their priorities, do not specify a new priority if an existing rule already uses this priority, as this can cause an error. If you need to reuse a priority with a different rule, you must remove it as a priority first, and then specify it in a subsequent update.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::LoadBalancer`.
//
// Specifies an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancer := awscdk.Aws_elasticloadbalancingv2.NewCfnLoadBalancer(this, jsii.String("MyCfnLoadBalancer"), &cfnLoadBalancerProps{
//   	ipAddressType: jsii.String("ipAddressType"),
//   	loadBalancerAttributes: []interface{}{
//   		&loadBalancerAttributeProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	scheme: jsii.String("scheme"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnetMappings: []interface{}{
//   		&subnetMappingProperty{
//   			subnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			allocationId: jsii.String("allocationId"),
//   			iPv6Address: jsii.String("iPv6Address"),
//   			privateIPv4Address: jsii.String("privateIPv4Address"),
//   		},
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   })
//
type CfnLoadBalancer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID of the Amazon Route 53 hosted zone associated with the load balancer.
	//
	// For example, `Z2P70J7EXAMPLE` .
	AttrCanonicalHostedZoneId() *string
	// The DNS name for the load balancer.
	//
	// For example, `my-load-balancer-424835706.us-west-2.elb.amazonaws.com` .
	AttrDnsName() *string
	// The full name of the load balancer.
	//
	// For example, `app/my-load-balancer/50dc6c495c0c9188` .
	AttrLoadBalancerFullName() *string
	// The name of the load balancer.
	//
	// For example, `my-load-balancer` .
	AttrLoadBalancerName() *string
	// The IDs of the security groups for the load balancer.
	AttrSecurityGroups() *[]*string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The IP address type.
	//
	// The possible values are `ipv4` (for IPv4 addresses) and `dualstack` (for IPv4 and IPv6 addresses). You can’t specify `dualstack` for a load balancer with a UDP or TCP_UDP listener.
	IpAddressType() *string
	SetIpAddressType(val *string)
	// The load balancer attributes.
	LoadBalancerAttributes() interface{}
	SetLoadBalancerAttributes(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the load balancer.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-".
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The nodes of an Internet-facing load balancer have public IP addresses.
	//
	// The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.
	//
	// The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.
	//
	// The default is an Internet-facing load balancer.
	//
	// You cannot specify a scheme for a Gateway Load Balancer.
	Scheme() *string
	SetScheme(val *string)
	// [Application Load Balancers] The IDs of the security groups for the load balancer.
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	SubnetMappings() interface{}
	SetSubnetMappings(val interface{})
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
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	// The tags to assign to the load balancer.
	Tags() awscdk.TagManager
	// The type of load balancer.
	//
	// The default is `application` .
	Type() *string
	SetType(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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

func (j *jsiiProxy_CfnLoadBalancer) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (c *jsiiProxy_CfnLoadBalancer) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnLoadBalancer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerAttributeProperty := &loadBalancerAttributeProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
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
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the attribute.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Specifies a subnet for a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetMappingProperty := &subnetMappingProperty{
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	allocationId: jsii.String("allocationId"),
//   	iPv6Address: jsii.String("iPv6Address"),
//   	privateIPv4Address: jsii.String("privateIPv4Address"),
//   }
//
type CfnLoadBalancer_SubnetMappingProperty struct {
	// The ID of the subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// [Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// [Network Load Balancers] The IPv6 address.
	IPv6Address *string `field:"optional" json:"iPv6Address" yaml:"iPv6Address"`
	// [Network Load Balancers] The private IPv4 address for an internal load balancer.
	PrivateIPv4Address *string `field:"optional" json:"privateIPv4Address" yaml:"privateIPv4Address"`
}

// Properties for defining a `CfnLoadBalancer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancerProps := &cfnLoadBalancerProps{
//   	ipAddressType: jsii.String("ipAddressType"),
//   	loadBalancerAttributes: []interface{}{
//   		&loadBalancerAttributeProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	scheme: jsii.String("scheme"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnetMappings: []interface{}{
//   		&subnetMappingProperty{
//   			subnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			allocationId: jsii.String("allocationId"),
//   			iPv6Address: jsii.String("iPv6Address"),
//   			privateIPv4Address: jsii.String("privateIPv4Address"),
//   		},
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnLoadBalancerProps struct {
	// The IP address type.
	//
	// The possible values are `ipv4` (for IPv4 addresses) and `dualstack` (for IPv4 and IPv6 addresses). You can’t specify `dualstack` for a load balancer with a UDP or TCP_UDP listener.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The load balancer attributes.
	LoadBalancerAttributes interface{} `field:"optional" json:"loadBalancerAttributes" yaml:"loadBalancerAttributes"`
	// The name of the load balancer.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-".
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The nodes of an Internet-facing load balancer have public IP addresses.
	//
	// The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.
	//
	// The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.
	//
	// The default is an Internet-facing load balancer.
	//
	// You cannot specify a scheme for a Gateway Load Balancer.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// [Application Load Balancers] The IDs of the security groups for the load balancer.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
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
	SubnetMappings interface{} `field:"optional" json:"subnetMappings" yaml:"subnetMappings"`
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
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The tags to assign to the load balancer.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of load balancer.
	//
	// The default is `application` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// A CloudFormation `AWS::ElasticLoadBalancingV2::TargetGroup`.
//
// Specifies a target group for an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.
//
// If the protocol of the target group is TCP, TLS, UDP, or TCP_UDP, you can't modify the health check protocol, interval, timeout, or success codes.
//
// Before you register a Lambda function as a target, you must create a `AWS::Lambda::Permission` resource that grants the Elastic Load Balancing service principal permission to invoke the Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTargetGroup := awscdk.Aws_elasticloadbalancingv2.NewCfnTargetGroup(this, jsii.String("MyCfnTargetGroup"), &cfnTargetGroupProps{
//   	healthCheckEnabled: jsii.Boolean(false),
//   	healthCheckIntervalSeconds: jsii.Number(123),
//   	healthCheckPath: jsii.String("healthCheckPath"),
//   	healthCheckPort: jsii.String("healthCheckPort"),
//   	healthCheckProtocol: jsii.String("healthCheckProtocol"),
//   	healthCheckTimeoutSeconds: jsii.Number(123),
//   	healthyThresholdCount: jsii.Number(123),
//   	ipAddressType: jsii.String("ipAddressType"),
//   	matcher: &matcherProperty{
//   		grpcCode: jsii.String("grpcCode"),
//   		httpCode: jsii.String("httpCode"),
//   	},
//   	name: jsii.String("name"),
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	protocolVersion: jsii.String("protocolVersion"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetGroupAttributes: []interface{}{
//   		&targetGroupAttributeProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targets: []interface{}{
//   		&targetDescriptionProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			availabilityZone: jsii.String("availabilityZone"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   	targetType: jsii.String("targetType"),
//   	unhealthyThresholdCount: jsii.Number(123),
//   	vpcId: jsii.String("vpcId"),
//   })
//
type CfnTargetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group.
	AttrLoadBalancerArns() *[]*string
	// The full name of the target group.
	//
	// For example, `targetgroup/my-target-group/cbf133c568e0d028` .
	AttrTargetGroupFullName() *string
	// The name of the target group.
	//
	// For example, `my-target-group` .
	AttrTargetGroupName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Indicates whether health checks are enabled.
	//
	// If the target type is `lambda` , health checks are disabled by default but can be enabled. If the target type is `instance` , `ip` , or `alb` , health checks are always enabled and cannot be disabled.
	HealthCheckEnabled() interface{}
	SetHealthCheckEnabled(val interface{})
	// The approximate amount of time, in seconds, between health checks of an individual target.
	//
	// If the target group protocol is HTTP or HTTPS, the default is 30 seconds. If the target group protocol is TCP, TLS, UDP, or TCP_UDP, the supported values are 10 and 30 seconds and the default is 30 seconds. If the target group protocol is GENEVE, the default is 10 seconds. If the target type is `lambda` , the default is 35 seconds.
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	//
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /.
	//
	// [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is / AWS .ALB/healthcheck.
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	// The port the load balancer uses when performing health checks on targets.
	//
	// If the protocol is HTTP, HTTPS, TCP, TLS, UDP, or TCP_UDP, the default is `traffic-port` , which is the port on which each target receives traffic from the load balancer. If the protocol is GENEVE, the default is port 80.
	HealthCheckPort() *string
	SetHealthCheckPort(val *string)
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers and Gateway Load Balancers, the default is TCP. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. The GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// For target groups with a protocol of HTTP, HTTPS, or GENEVE, the default is 5 seconds. For target groups with a protocol of TCP or TLS, this value must be 6 seconds for HTTP health checks and 10 seconds for TCP and HTTPS health checks. If the target type is `lambda` , the default is 30 seconds.
	HealthCheckTimeoutSeconds() *float64
	SetHealthCheckTimeoutSeconds(val *float64)
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	//
	// For target groups with a protocol of HTTP or HTTPS, the default is 5. For target groups with a protocol of TCP, TLS, or GENEVE, the default is 3. If the target type is `lambda` , the default is 5.
	HealthyThresholdCount() *float64
	SetHealthyThresholdCount(val *float64)
	// The type of IP address used for this target group.
	//
	// The possible values are `ipv4` and `ipv6` . This is an optional parameter. If not specified, the IP address type defaults to `ipv4` .
	IpAddressType() *string
	SetIpAddressType(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
	Matcher() interface{}
	SetMatcher(val interface{})
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The port on which the targets receive traffic.
	//
	// This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.
	Port() *float64
	SetPort(val *float64)
	// The protocol to use for routing traffic to the targets.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, or TCP_UDP. For Gateway Load Balancers, the supported protocol is GENEVE. A TCP_UDP listener must be associated with a TCP_UDP target group. If the target is a Lambda function, this parameter does not apply.
	Protocol() *string
	SetProtocol(val *string)
	// [HTTP/HTTPS protocol] The protocol version.
	//
	// The possible values are `GRPC` , `HTTP1` , and `HTTP2` .
	ProtocolVersion() *string
	SetProtocolVersion(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags.
	Tags() awscdk.TagManager
	// The attributes.
	TargetGroupAttributes() interface{}
	SetTargetGroupAttributes(val interface{})
	// The targets.
	Targets() interface{}
	SetTargets(val interface{})
	// The type of target that you must specify when registering targets with this target group.
	//
	// You can't specify targets for a target group using more than one target type.
	//
	// - `instance` - Register targets by instance ID. This is the default value.
	// - `ip` - Register targets by IP address. You can specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	// - `lambda` - Register a single Lambda function as a target.
	// - `alb` - Register a single Application Load Balancer as a target.
	TargetType() *string
	SetTargetType(val *string)
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// If the target group protocol is HTTP or HTTPS, the default is 2. If the target group protocol is TCP or TLS, this value must be the same as the healthy threshold count. If the target group protocol is GENEVE, the default is 3. If the target type is `lambda` , the default is 2.
	UnhealthyThresholdCount() *float64
	SetUnhealthyThresholdCount(val *float64)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The identifier of the virtual private cloud (VPC).
	//
	// If the target is a Lambda function, this parameter does not apply. Otherwise, this parameter is required.
	VpcId() *string
	SetVpcId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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

func (j *jsiiProxy_CfnTargetGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (c *jsiiProxy_CfnTargetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTargetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTargetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTargetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTargetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTargetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTargetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnTargetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matcherProperty := &matcherProperty{
//   	grpcCode: jsii.String("grpcCode"),
//   	httpCode: jsii.String("httpCode"),
//   }
//
type CfnTargetGroup_MatcherProperty struct {
	// You can specify values between 0 and 99.
	//
	// You can specify multiple values (for example, "0,1") or a range of values (for example, "0-5"). The default value is 12.
	GrpcCode *string `field:"optional" json:"grpcCode" yaml:"grpcCode"`
	// For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200.
	//
	// You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299").
	//
	// For Network Load Balancers and Gateway Load Balancers, this must be "200–399".
	//
	// Note that when using shorthand syntax, some values such as commas need to be escaped.
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
}

// Specifies a target to add to a target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetDescriptionProperty := &targetDescriptionProperty{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	availabilityZone: jsii.String("availabilityZone"),
//   	port: jsii.Number(123),
//   }
//
type CfnTargetGroup_TargetDescriptionProperty struct {
	// The ID of the target.
	//
	// If the target type of the target group is `instance` , specify an instance ID. If the target type is `ip` , specify an IP address. If the target type is `lambda` , specify the ARN of the Lambda function. If the target type is `alb` , specify the ARN of the Application Load Balancer target.
	Id *string `field:"required" json:"id" yaml:"id"`
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
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The port on which the target is listening.
	//
	// If the target group protocol is GENEVE, the supported port is 6081. If the target type is `alb` , the targeted Application Load Balancer must have at least one listener whose port matches the target group port. Not used if the target is a Lambda function.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// Specifies a target group attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupAttributeProperty := &targetGroupAttributeProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
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
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the attribute.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Properties for defining a `CfnTargetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTargetGroupProps := &cfnTargetGroupProps{
//   	healthCheckEnabled: jsii.Boolean(false),
//   	healthCheckIntervalSeconds: jsii.Number(123),
//   	healthCheckPath: jsii.String("healthCheckPath"),
//   	healthCheckPort: jsii.String("healthCheckPort"),
//   	healthCheckProtocol: jsii.String("healthCheckProtocol"),
//   	healthCheckTimeoutSeconds: jsii.Number(123),
//   	healthyThresholdCount: jsii.Number(123),
//   	ipAddressType: jsii.String("ipAddressType"),
//   	matcher: &matcherProperty{
//   		grpcCode: jsii.String("grpcCode"),
//   		httpCode: jsii.String("httpCode"),
//   	},
//   	name: jsii.String("name"),
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	protocolVersion: jsii.String("protocolVersion"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetGroupAttributes: []interface{}{
//   		&targetGroupAttributeProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targets: []interface{}{
//   		&targetDescriptionProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			availabilityZone: jsii.String("availabilityZone"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   	targetType: jsii.String("targetType"),
//   	unhealthyThresholdCount: jsii.Number(123),
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnTargetGroupProps struct {
	// Indicates whether health checks are enabled.
	//
	// If the target type is `lambda` , health checks are disabled by default but can be enabled. If the target type is `instance` , `ip` , or `alb` , health checks are always enabled and cannot be disabled.
	HealthCheckEnabled interface{} `field:"optional" json:"healthCheckEnabled" yaml:"healthCheckEnabled"`
	// The approximate amount of time, in seconds, between health checks of an individual target.
	//
	// If the target group protocol is HTTP or HTTPS, the default is 30 seconds. If the target group protocol is TCP, TLS, UDP, or TCP_UDP, the supported values are 10 and 30 seconds and the default is 30 seconds. If the target group protocol is GENEVE, the default is 10 seconds. If the target type is `lambda` , the default is 35 seconds.
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	//
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /.
	//
	// [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is / AWS .ALB/healthcheck.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port the load balancer uses when performing health checks on targets.
	//
	// If the protocol is HTTP, HTTPS, TCP, TLS, UDP, or TCP_UDP, the default is `traffic-port` , which is the port on which each target receives traffic from the load balancer. If the protocol is GENEVE, the default is port 80.
	HealthCheckPort *string `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers and Gateway Load Balancers, the default is TCP. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. The GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	HealthCheckProtocol *string `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// For target groups with a protocol of HTTP, HTTPS, or GENEVE, the default is 5 seconds. For target groups with a protocol of TCP or TLS, this value must be 6 seconds for HTTP health checks and 10 seconds for TCP and HTTPS health checks. If the target type is `lambda` , the default is 30 seconds.
	HealthCheckTimeoutSeconds *float64 `field:"optional" json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	//
	// For target groups with a protocol of HTTP or HTTPS, the default is 5. For target groups with a protocol of TCP, TLS, or GENEVE, the default is 3. If the target type is `lambda` , the default is 5.
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The type of IP address used for this target group.
	//
	// The possible values are `ipv4` and `ipv6` . This is an optional parameter. If not specified, the IP address type defaults to `ipv4` .
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on which the targets receive traffic.
	//
	// This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use for routing traffic to the targets.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, or TCP_UDP. For Gateway Load Balancers, the supported protocol is GENEVE. A TCP_UDP listener must be associated with a TCP_UDP target group. If the target is a Lambda function, this parameter does not apply.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// [HTTP/HTTPS protocol] The protocol version.
	//
	// The possible values are `GRPC` , `HTTP1` , and `HTTP2` .
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The attributes.
	TargetGroupAttributes interface{} `field:"optional" json:"targetGroupAttributes" yaml:"targetGroupAttributes"`
	// The targets.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The type of target that you must specify when registering targets with this target group.
	//
	// You can't specify targets for a target group using more than one target type.
	//
	// - `instance` - Register targets by instance ID. This is the default value.
	// - `ip` - Register targets by IP address. You can specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	// - `lambda` - Register a single Lambda function as a target.
	// - `alb` - Register a single Application Load Balancer as a target.
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// If the target group protocol is HTTP or HTTPS, the default is 2. If the target group protocol is TCP or TLS, this value must be the same as the healthy threshold count. If the target group protocol is GENEVE, the default is 3. If the target type is `lambda` , the default is 2.
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
	// The identifier of the virtual private cloud (VPC).
	//
	// If the target is a Lambda function, this parameter does not apply. Otherwise, this parameter is required.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

// Options for `ListenerAction.fixedResponse()`.
//
// Example:
//   var listener applicationListener
//
//
//   listener.addAction(jsii.String("Fixed"), &addApplicationActionProps{
//   	priority: jsii.Number(10),
//   	conditions: []listenerCondition{
//   		elbv2.*listenerCondition.pathPatterns([]*string{
//   			jsii.String("/ok"),
//   		}),
//   	},
//   	action: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
//   		contentType: elbv2.contentType_TEXT_PLAIN,
//   		messageBody: jsii.String("OK"),
//   	}),
//   })
//
type FixedResponseOptions struct {
	// Content Type of the response.
	//
	// Valid Values: text/plain | text/css | text/html | application/javascript | application/json.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The response body.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
}

// Options for `ListenerAction.forward()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forwardOptions := &forwardOptions{
//   	stickinessDuration: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type ForwardOptions struct {
	// For how long clients should be directed to the same target group.
	//
	// Range between 1 second and 7 days.
	StickinessDuration awscdk.Duration `field:"optional" json:"stickinessDuration" yaml:"stickinessDuration"`
}

// Properties for configuring a health check.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   loadBalancedFargateService.targetGroup.configureHealthCheck(&healthCheck{
//   	path: jsii.String("/custom-health-path"),
//   })
//
type HealthCheck struct {
	// Indicates whether health checks are enabled.
	//
	// If the target type is lambda,
	// health checks are disabled by default but can be enabled. If the target type
	// is instance or ip, health checks are always enabled and cannot be disabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// GRPC code to use when checking for a successful response from a target.
	//
	// You can specify values between 0 and 99. You can specify multiple values
	// (for example, "0,1") or a range of values (for example, "0-5").
	HealthyGrpcCodes *string `field:"optional" json:"healthyGrpcCodes" yaml:"healthyGrpcCodes"`
	// HTTP code to use when checking for a successful response from a target.
	//
	// For Application Load Balancers, you can specify values between 200 and
	// 499, and the default value is 200. You can specify multiple values (for
	// example, "200,202") or a range of values (for example, "200-299").
	HealthyHttpCodes *string `field:"optional" json:"healthyHttpCodes" yaml:"healthyHttpCodes"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	//
	// For Application Load Balancers, the default is 5. For Network Load Balancers, the default is 3.
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The approximate number of seconds between health checks for an individual target.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The ping path destination where Elastic Load Balancing sends health check requests.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port that the load balancer uses when performing health checks on the targets.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// The TCP protocol is supported for health checks only if the protocol of the target group is TCP, TLS, UDP, or TCP_UDP.
	// The TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// For Application Load Balancers, the range is 2-60 seconds and the
	// default is 5 seconds. For Network Load Balancers, this is 10 seconds for
	// TCP and HTTPS health checks and 6 seconds for HTTP health checks.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// For Application Load Balancers, the default is 2. For Network Load
	// Balancers, this value must be the same as the healthy threshold count.
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
}

// Count of HTTP status originating from the load balancer.
//
// This count does not include any response codes generated by the targets.
type HttpCodeElb string

const (
	// The number of HTTP 3XX redirection codes that originate from the load balancer.
	HttpCodeElb_ELB_3XX_COUNT HttpCodeElb = "ELB_3XX_COUNT"
	// The number of HTTP 4XX client error codes that originate from the load balancer.
	//
	// Client errors are generated when requests are malformed or incomplete.
	// These requests have not been received by the target. This count does not
	// include any response codes generated by the targets.
	HttpCodeElb_ELB_4XX_COUNT HttpCodeElb = "ELB_4XX_COUNT"
	// The number of HTTP 5XX server error codes that originate from the load balancer.
	HttpCodeElb_ELB_5XX_COUNT HttpCodeElb = "ELB_5XX_COUNT"
)

// Count of HTTP status originating from the targets.
type HttpCodeTarget string

const (
	// The number of 2xx response codes from targets.
	HttpCodeTarget_TARGET_2XX_COUNT HttpCodeTarget = "TARGET_2XX_COUNT"
	// The number of 3xx response codes from targets.
	HttpCodeTarget_TARGET_3XX_COUNT HttpCodeTarget = "TARGET_3XX_COUNT"
	// The number of 4xx response codes from targets.
	HttpCodeTarget_TARGET_4XX_COUNT HttpCodeTarget = "TARGET_4XX_COUNT"
	// The number of 5xx response codes from targets.
	HttpCodeTarget_TARGET_5XX_COUNT HttpCodeTarget = "TARGET_5XX_COUNT"
)

// Properties to reference an existing listener.
type IApplicationListener interface {
	awsec2.IConnectable
	awscdk.IResource
	// Perform the given action on incoming requests.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses. See
	// the `ListenerAction` class for all options.
	//
	// It's possible to add routing conditions to the Action added in this way.
	//
	// It is not possible to add a default action to an imported IApplicationListener.
	// In order to add actions to an imported IApplicationListener a `priority`
	// must be provided.
	AddAction(id *string, props *AddApplicationActionProps)
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
	// Returns: The newly created target group.
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

func (i *jsiiProxy_IApplicationListener) AddAction(id *string, props *AddApplicationActionProps) {
	_jsii_.InvokeVoid(
		i,
		"addAction",
		[]interface{}{id, props},
	)
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
	// Example value: `Z2P70J7EXAMPLE`.
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
	// Returns: The newly created listener.
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
	// Allocate IPv4 addresses.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// Allocate both IPv4 and IPv6 addresses.
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
// Example:
//   var listener applicationListener
//   var myTargetGroup applicationTargetGroup
//
//
//   listener.addAction(jsii.String("DefaultAction"), &addApplicationActionProps{
//   	action: elbv2.listenerAction.authenticateOidc(&authenticateOidcOptions{
//   		authorizationEndpoint: jsii.String("https://example.com/openid"),
//   		// Other OIDC properties here
//   		clientId: jsii.String("..."),
//   		clientSecret: awscdk.SecretValue.secretsManager(jsii.String("...")),
//   		issuer: jsii.String("..."),
//   		tokenEndpoint: jsii.String("..."),
//   		userInfoEndpoint: jsii.String("..."),
//
//   		// Next
//   		next: elbv2.*listenerAction.forward([]iApplicationTargetGroup{
//   			myTargetGroup,
//   		}),
//   	}),
//   })
//
type ListenerAction interface {
	IListenerAction
	Next() ListenerAction
	// Called when the action is being used in a listener.
	Bind(scope constructs.Construct, listener IApplicationListener, associatingConstruct constructs.IConstruct)
	// Render the actions in this chain.
	RenderActions() *[]*CfnListener_ActionProperty
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

func (l *jsiiProxy_ListenerAction) Bind(scope constructs.Construct, listener IApplicationListener, associatingConstruct constructs.IConstruct) {
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerCertificate := awscdk.Aws_elasticloadbalancingv2.listenerCertificate.fromArn(jsii.String("certificateArn"))
//
type ListenerCertificate interface {
	IListenerCertificate
	// The ARN of the certificate to use.
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
// Example:
//   var listener applicationListener
//   var asg autoScalingGroup
//
//
//   listener.addTargets(jsii.String("Example.Com Fleet"), &addApplicationTargetsProps{
//   	priority: jsii.Number(10),
//   	conditions: []listenerCondition{
//   		elbv2.*listenerCondition.hostHeaders([]*string{
//   			jsii.String("example.com"),
//   		}),
//   		elbv2.*listenerCondition.pathPatterns([]*string{
//   			jsii.String("/ok"),
//   			jsii.String("/path"),
//   		}),
//   	},
//   	port: jsii.Number(8080),
//   	targets: []iApplicationLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type ListenerCondition interface {
	// Render the raw Cfn listener rule condition object.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var targetJson interface{}
//
//   loadBalancerTargetProps := &loadBalancerTargetProps{
//   	targetType: awscdk.Aws_elasticloadbalancingv2.targetType_INSTANCE,
//
//   	// the properties below are optional
//   	targetJson: targetJson,
//   }
//
type LoadBalancerTargetProps struct {
	// What kind of target this is.
	TargetType TargetType `field:"required" json:"targetType" yaml:"targetType"`
	// JSON representing the target's direct addition to the TargetGroup list.
	//
	// May be omitted if the target is going to register itself later.
	TargetJson interface{} `field:"optional" json:"targetJson" yaml:"targetJson"`
}

// Options for `NetworkListenerAction.forward()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkForwardOptions := &networkForwardOptions{
//   	stickinessDuration: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type NetworkForwardOptions struct {
	// For how long clients should be directed to the same target group.
	//
	// Range between 1 second and 7 days.
	StickinessDuration awscdk.Duration `field:"optional" json:"stickinessDuration" yaml:"stickinessDuration"`
}

// Define a Network Listener.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addNetworkTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type NetworkListener interface {
	BaseListener
	INetworkListener
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	// The load balancer this listener is attached to.
	LoadBalancer() INetworkLoadBalancer
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Perform the given Action on incoming requests.
	//
	// This allows full control of the default Action of the load balancer,
	// including weighted forwarding. See the `NetworkListenerAction` class for
	// all options.
	AddAction(_id *string, props *AddNetworkActionProps)
	// Add one or more certificates to this listener.
	//
	// After the first certificate, this creates NetworkListenerCertificates
	// resources since cloudformation requires the certificates array on the
	// listener resource to have a length of 1.
	AddCertificates(id *string, certificates *[]IListenerCertificate)
	// Load balance incoming requests to the given target groups.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use `addAction()`.
	AddTargetGroups(_id *string, targetGroups ...INetworkTargetGroup)
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
	// Returns: The newly created target group.
	AddTargets(id *string, props *AddNetworkTargetsProps) NetworkTargetGroup
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
	// Validate this listener.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

// Returns true if the construct was created by CDK, and false otherwise.
func NetworkListener_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkListener",
		"isOwnedResource",
		[]interface{}{construct},
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

func (n *jsiiProxy_NetworkListener) AddAction(_id *string, props *AddNetworkActionProps) {
	_jsii_.InvokeVoid(
		n,
		"addAction",
		[]interface{}{_id, props},
	)
}

func (n *jsiiProxy_NetworkListener) AddCertificates(id *string, certificates *[]IListenerCertificate) {
	_jsii_.InvokeVoid(
		n,
		"addCertificates",
		[]interface{}{id, certificates},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkTargetGroup networkTargetGroup
//
//   networkListenerAction := awscdk.Aws_elasticloadbalancingv2.networkListenerAction.forward([]iNetworkTargetGroup{
//   	networkTargetGroup,
//   }, &networkForwardOptions{
//   	stickinessDuration: cdk.duration.minutes(jsii.Number(30)),
//   })
//
type NetworkListenerAction interface {
	IListenerAction
	Next() NetworkListenerAction
	// Called when the action is being used in a listener.
	Bind(scope constructs.Construct, listener INetworkListener)
	// Render the actions in this chain.
	RenderActions() *[]*CfnListener_ActionProperty
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

func (n *jsiiProxy_NetworkListenerAction) Bind(scope constructs.Construct, listener INetworkListener) {
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
// Example:
//   listener := elbv2.networkListener.fromLookup(this, jsii.String("ALBListener"), &networkListenerLookupOptions{
//   	loadBalancerTags: map[string]*string{
//   		"Cluster": jsii.String("MyClusterName"),
//   	},
//   	listenerProtocol: elbv2.protocol_TCP,
//   	listenerPort: jsii.Number(12345),
//   })
//
type NetworkListenerLookupOptions struct {
	// Filter listeners by listener port.
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Protocol of the listener port.
	ListenerProtocol Protocol `field:"optional" json:"listenerProtocol" yaml:"listenerProtocol"`
}

// Properties for a Network Listener attached to a Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerCertificate listenerCertificate
//   var networkListenerAction networkListenerAction
//   var networkLoadBalancer networkLoadBalancer
//   var networkTargetGroup networkTargetGroup
//
//   networkListenerProps := &networkListenerProps{
//   	loadBalancer: networkLoadBalancer,
//   	port: jsii.Number(123),
//
//   	// the properties below are optional
//   	alpnPolicy: awscdk.Aws_elasticloadbalancingv2.alpnPolicy_HTTP1_ONLY,
//   	certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   	defaultAction: networkListenerAction,
//   	defaultTargetGroups: []iNetworkTargetGroup{
//   		networkTargetGroup,
//   	},
//   	protocol: awscdk.*Aws_elasticloadbalancingv2.protocol_HTTP,
//   	sslPolicy: awscdk.*Aws_elasticloadbalancingv2.sslPolicy_RECOMMENDED_TLS,
//   }
//
type NetworkListenerProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Application-Layer Protocol Negotiation (ALPN) is a TLS extension that is sent on the initial TLS handshake hello messages.
	//
	// ALPN enables the application layer to negotiate which protocols should be used over a secure connection, such as HTTP/1 and HTTP/2.
	//
	// Can only be specified together with Protocol TLS.
	AlpnPolicy AlpnPolicy `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default Action of the load balancer,
	// including weighted forwarding. See the `NetworkListenerAction` class for
	// all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction NetworkListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]INetworkTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Protocol for listener, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// SSL Policy.
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// The load balancer to attach this listener to.
	LoadBalancer INetworkLoadBalancer `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
}

// Define a new network load balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addNetworkTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type NetworkLoadBalancer interface {
	BaseLoadBalancer
	INetworkLoadBalancer
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN of this load balancer.
	//
	// Example value: `arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-internal-load-balancer/50dc6c495c0c9188`.
	LoadBalancerArn() *string
	// The canonical hosted zone ID of this load balancer.
	//
	// Example value: `Z2P70J7EXAMPLE`.
	LoadBalancerCanonicalHostedZoneId() *string
	// The DNS name of this load balancer.
	//
	// Example value: `my-load-balancer-424835706.us-west-2.elb.amazonaws.com`
	LoadBalancerDnsName() *string
	// The full name of this load balancer.
	//
	// Example value: `app/my-load-balancer/50dc6c495c0c9188`.
	LoadBalancerFullName() *string
	// The name of this load balancer.
	//
	// Example value: `my-load-balancer`.
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The VPC this load balancer has been created in.
	//
	// This property is always defined (not `null` or `undefined`) for sub-classes of `BaseLoadBalancer`.
	Vpc() awsec2.IVpc
	// Add a listener to this load balancer.
	//
	// Returns: The newly created listener.
	AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Enable access logging for this load balancer.
	//
	// A region must be specified on the stack containing the load balancer; you cannot enable logging on
	// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	// Return the given named metric for this Network Load Balancer.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of concurrent TCP flows (or connections) from clients to targets.
	//
	// This metric includes connections in the SYN_SENT and ESTABLISHED states.
	// TCP connections are not terminated at the load balancer, so a client
	// opening a TCP connection to a target counts as a single flow.
	MetricActiveFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of load balancer capacity units (LCU) used by your load balancer.
	MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of new TCP flows (or connections) established from clients to targets in the time period.
	MetricNewFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of bytes processed by the load balancer, including TCP/IP headers.
	MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of reset (RST) packets sent from a client to a target.
	//
	// These resets are generated by the client and forwarded by the load balancer.
	MetricTcpClientResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of reset (RST) packets generated by the load balancer.
	MetricTcpElbResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of reset (RST) packets sent from a target to a client.
	//
	// These resets are generated by the target and forwarded by the load balancer.
	MetricTcpTargetResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Remove an attribute from the load balancer.
	RemoveAttribute(key *string)
	// Set a non-standard attribute on the load balancer.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
	ValidateLoadBalancer() *[]*string
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

// Returns true if the construct was created by CDK, and false otherwise.
func NetworkLoadBalancer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"isOwnedResource",
		[]interface{}{construct},
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

func (n *jsiiProxy_NetworkLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	_jsii_.InvokeVoid(
		n,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

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

func (n *jsiiProxy_NetworkLoadBalancer) RemoveAttribute(key *string) {
	_jsii_.InvokeVoid(
		n,
		"removeAttribute",
		[]interface{}{key},
	)
}

func (n *jsiiProxy_NetworkLoadBalancer) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		n,
		"setAttribute",
		[]interface{}{key, value},
	)
}

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

func (n *jsiiProxy_NetworkLoadBalancer) ValidateLoadBalancer() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validateLoadBalancer",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to reference an existing load balancer.
//
// Example:
//   // Create an Accelerator
//   accelerator := globalaccelerator.NewAccelerator(this, jsii.String("Accelerator"))
//
//   // Create a Listener
//   listener := accelerator.addListener(jsii.String("Listener"), &listenerOptions{
//   	portRanges: []portRange{
//   		&portRange{
//   			fromPort: jsii.Number(80),
//   		},
//   		&portRange{
//   			fromPort: jsii.Number(443),
//   		},
//   	},
//   })
//
//   // Import the Load Balancers
//   nlb1 := elbv2.networkLoadBalancer.fromNetworkLoadBalancerAttributes(this, jsii.String("NLB1"), &networkLoadBalancerAttributes{
//   	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-west-2:111111111111:loadbalancer/app/my-load-balancer1/e16bef66805b"),
//   })
//   nlb2 := elbv2.networkLoadBalancer.fromNetworkLoadBalancerAttributes(this, jsii.String("NLB2"), &networkLoadBalancerAttributes{
//   	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:ap-south-1:111111111111:loadbalancer/app/my-load-balancer2/5513dc2ea8a1"),
//   })
//
//   // Add one EndpointGroup for each Region we are targeting
//   listener.addEndpointGroup(jsii.String("Group1"), &endpointGroupOptions{
//   	endpoints: []iEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb1),
//   	},
//   })
//   listener.addEndpointGroup(jsii.String("Group2"), &endpointGroupOptions{
//   	// Imported load balancers automatically calculate their Region from the ARN.
//   	// If you are load balancing to other resources, you must also pass a `region`
//   	// parameter here.
//   	endpoints: []*iEndpoint{
//   		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb2),
//   	},
//   })
//
type NetworkLoadBalancerAttributes struct {
	// ARN of the load balancer.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The canonical hosted zone ID of this load balancer.
	LoadBalancerCanonicalHostedZoneId *string `field:"optional" json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// The DNS name of this load balancer.
	LoadBalancerDnsName *string `field:"optional" json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// The VPC to associate with the load balancer.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// Options for looking up an NetworkLoadBalancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkLoadBalancerLookupOptions := &networkLoadBalancerLookupOptions{
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: map[string]*string{
//   		"loadBalancerTagsKey": jsii.String("loadBalancerTags"),
//   	},
//   }
//
type NetworkLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

// Properties for a network load balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addNetworkTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type NetworkLoadBalancerProps struct {
	// The VPC network to place the load balancer in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether deletion protection is enabled.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether the load balancer has an internet-routable address.
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	// Name of the load balancer.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// Which subnets place the load balancer in.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Indicates whether cross-zone load balancing is enabled.
	CrossZoneEnabled *bool `field:"optional" json:"crossZoneEnabled" yaml:"crossZoneEnabled"`
}

// Define a Network Target Group.
//
// Example:
//   var listener networkListener
//   var asg1 autoScalingGroup
//   var asg2 autoScalingGroup
//
//
//   group := listener.addTargets(jsii.String("AppFleet"), &addNetworkTargetsProps{
//   	port: jsii.Number(443),
//   	targets: []iNetworkLoadBalancerTarget{
//   		asg1,
//   	},
//   })
//
//   group.addTarget(asg2)
//
type NetworkTargetGroup interface {
	TargetGroupBase
	INetworkTargetGroup
	// Default port configured for members of this target group.
	DefaultPort() *float64
	// Full name of first load balancer.
	FirstLoadBalancerFullName() *string
	HealthCheck() *HealthCheck
	SetHealthCheck(val *HealthCheck)
	// A token representing a list of ARNs of the load balancers that route traffic to this target group.
	LoadBalancerArns() *string
	// List of constructs that need to be depended on to ensure the TargetGroup is associated to a load balancer.
	LoadBalancerAttached() constructs.IDependable
	// Configurable dependable with all resources that lead to load balancer attachment.
	LoadBalancerAttachedDependencies() constructs.DependencyGroup
	// The tree node.
	Node() constructs.Node
	// The ARN of the target group.
	TargetGroupArn() *string
	// The full name of the target group.
	TargetGroupFullName() *string
	// ARNs of load balancers load balancing to this TargetGroup.
	TargetGroupLoadBalancerArns() *[]*string
	// The name of the target group.
	TargetGroupName() *string
	// The types of the directly registered members of this target group.
	TargetType() TargetType
	SetTargetType(val TargetType)
	// Register the given load balancing target as part of this group.
	AddLoadBalancerTarget(props *LoadBalancerTargetProps)
	// Add a load balancing target to this target group.
	AddTarget(targets ...INetworkLoadBalancerTarget)
	// Set/replace the target group's health check.
	ConfigureHealthCheck(healthCheck *HealthCheck)
	// The number of targets that are considered healthy.
	MetricHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of targets that are considered unhealthy.
	MetricUnHealthyHostCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Register a listener that is load balancing to this target group.
	//
	// Don't call this directly. It will be called by listeners.
	RegisterListener(listener INetworkListener)
	// Set a non-standard attribute on the target group.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (n *jsiiProxy_NetworkTargetGroup) AddLoadBalancerTarget(props *LoadBalancerTargetProps) {
	_jsii_.InvokeVoid(
		n,
		"addLoadBalancerTarget",
		[]interface{}{props},
	)
}

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

func (n *jsiiProxy_NetworkTargetGroup) ConfigureHealthCheck(healthCheck *HealthCheck) {
	_jsii_.InvokeVoid(
		n,
		"configureHealthCheck",
		[]interface{}{healthCheck},
	)
}

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

func (n *jsiiProxy_NetworkTargetGroup) RegisterListener(listener INetworkListener) {
	_jsii_.InvokeVoid(
		n,
		"registerListener",
		[]interface{}{listener},
	)
}

func (n *jsiiProxy_NetworkTargetGroup) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		n,
		"setAttribute",
		[]interface{}{key, value},
	)
}

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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkLoadBalancerTarget iNetworkLoadBalancerTarget
//   var vpc vpc
//
//   networkTargetGroupProps := &networkTargetGroupProps{
//   	port: jsii.Number(123),
//
//   	// the properties below are optional
//   	connectionTermination: jsii.Boolean(false),
//   	deregistrationDelay: cdk.duration.minutes(jsii.Number(30)),
//   	healthCheck: &healthCheck{
//   		enabled: jsii.Boolean(false),
//   		healthyGrpcCodes: jsii.String("healthyGrpcCodes"),
//   		healthyHttpCodes: jsii.String("healthyHttpCodes"),
//   		healthyThresholdCount: jsii.Number(123),
//   		interval: cdk.*duration.minutes(jsii.Number(30)),
//   		path: jsii.String("path"),
//   		port: jsii.String("port"),
//   		protocol: awscdk.Aws_elasticloadbalancingv2.protocol_HTTP,
//   		timeout: cdk.*duration.minutes(jsii.Number(30)),
//   		unhealthyThresholdCount: jsii.Number(123),
//   	},
//   	preserveClientIp: jsii.Boolean(false),
//   	protocol: awscdk.*Aws_elasticloadbalancingv2.*protocol_HTTP,
//   	proxyProtocolV2: jsii.Boolean(false),
//   	targetGroupName: jsii.String("targetGroupName"),
//   	targets: []*iNetworkLoadBalancerTarget{
//   		networkLoadBalancerTarget,
//   	},
//   	targetType: awscdk.*Aws_elasticloadbalancingv2.targetType_INSTANCE,
//   	vpc: vpc,
//   }
//
type NetworkTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	TargetType TargetType `field:"optional" json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The port on which the listener listens for requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Indicates whether the load balancer terminates connections at the end of the deregistration timeout.
	ConnectionTermination *bool `field:"optional" json:"connectionTermination" yaml:"connectionTermination"`
	// Indicates whether client IP preservation is enabled.
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Protocol for target group, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether Proxy Protocol version 2 is enabled.
	ProxyProtocolV2 *bool `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]INetworkLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

// A Target Group and weight combination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkTargetGroup networkTargetGroup
//
//   networkWeightedTargetGroup := &networkWeightedTargetGroup{
//   	targetGroup: networkTargetGroup,
//
//   	// the properties below are optional
//   	weight: jsii.Number(123),
//   }
//
type NetworkWeightedTargetGroup struct {
	// The target group.
	TargetGroup INetworkTargetGroup `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// The target group's weight.
	//
	// Range is [0..1000).
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

// Backend protocol for network load balancers and health checks.
//
// Example:
//   listener := elbv2.networkListener.fromLookup(this, jsii.String("ALBListener"), &networkListenerLookupOptions{
//   	loadBalancerTags: map[string]*string{
//   		"Cluster": jsii.String("MyClusterName"),
//   	},
//   	listenerProtocol: elbv2.protocol_TCP,
//   	listenerPort: jsii.Number(12345),
//   })
//
type Protocol string

const (
	// HTTP (ALB health checks and NLB health checks).
	Protocol_HTTP Protocol = "HTTP"
	// HTTPS (ALB health checks and NLB health checks).
	Protocol_HTTPS Protocol = "HTTPS"
	// TCP (NLB, NLB health checks).
	Protocol_TCP Protocol = "TCP"
	// TLS (NLB).
	Protocol_TLS Protocol = "TLS"
	// UDP (NLB).
	Protocol_UDP Protocol = "UDP"
	// Listen to both TCP and UDP on the same port (NLB).
	Protocol_TCP_UDP Protocol = "TCP_UDP"
)

// Properties for the key/value pair of the query string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queryStringCondition := &queryStringCondition{
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	key: jsii.String("key"),
//   }
//
type QueryStringCondition struct {
	// The query string value for the condition.
	Value *string `field:"required" json:"value" yaml:"value"`
	// The query string key for the condition.
	Key *string `field:"optional" json:"key" yaml:"key"`
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectOptions := &redirectOptions{
//   	host: jsii.String("host"),
//   	path: jsii.String("path"),
//   	permanent: jsii.Boolean(false),
//   	port: jsii.String("port"),
//   	protocol: jsii.String("protocol"),
//   	query: jsii.String("query"),
//   }
//
type RedirectOptions struct {
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP redirect code.
	//
	// The redirect is either permanent (HTTP 301) or temporary (HTTP 302).
	Permanent *bool `field:"optional" json:"permanent" yaml:"permanent"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added. You can specify any of the reserved keywords.
	Query *string `field:"optional" json:"query" yaml:"query"`
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
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//   var cluster cluster
//
//
//   domainZone := awscdk.HostedZone.fromLookup(this, jsii.String("Zone"), &hostedZoneProviderProps{
//   	domainName: jsii.String("example.com"),
//   })
//   certificate := awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert"), jsii.String("arn:aws:acm:us-east-1:123456:certificate/abcdefg"))
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	vpc: vpc,
//   	cluster: cluster,
//   	certificate: certificate,
//   	sslPolicy: awscdk.SslPolicy_RECOMMENDED,
//   	domainName: jsii.String("api.example.com"),
//   	domainZone: domainZone,
//   	redirectHTTP: jsii.Boolean(true),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html
//
type SslPolicy string

const (
	// The recommended security policy for TLS listeners.
	//
	// This is the default policy for listeners created using the AWS Management Console.
	SslPolicy_RECOMMENDED_TLS SslPolicy = "RECOMMENDED_TLS"
	// The recommended policy for http listeners.
	//
	// This is the default security policy for listeners created using the AWS CLI.
	SslPolicy_RECOMMENDED SslPolicy = "RECOMMENDED"
	// TLS1.2 and 1.3.
	SslPolicy_TLS13_RES SslPolicy = "TLS13_RES"
	// TLS1.2 and 1.3 and no SHA ciphers.
	SslPolicy_TLS13_EXT1 SslPolicy = "TLS13_EXT1"
	// TLS1.2 and 1.3 with all ciphers.
	SslPolicy_TLS13_EXT2 SslPolicy = "TLS13_EXT2"
	// TLS1.0 through 1.3 with all ciphers.
	SslPolicy_TLS13_10 SslPolicy = "TLS13_10"
	// TLS1.1 through 1.3 with all ciphers.
	SslPolicy_TLS13_11 SslPolicy = "TLS13_11"
	// TLS1.3 only.
	SslPolicy_TLS13_13 SslPolicy = "TLS13_13"
	// Strong foward secrecy ciphers and TLV1.2 only (2020 edition). Same as FORWARD_SECRECY_TLS12_RES, but only supports GCM versions of the TLS ciphers.
	SslPolicy_FORWARD_SECRECY_TLS12_RES_GCM SslPolicy = "FORWARD_SECRECY_TLS12_RES_GCM"
	// Strong forward secrecy ciphers and TLS1.2 only.
	SslPolicy_FORWARD_SECRECY_TLS12_RES SslPolicy = "FORWARD_SECRECY_TLS12_RES"
	// Forward secrecy ciphers and TLS1.2 only.
	SslPolicy_FORWARD_SECRECY_TLS12 SslPolicy = "FORWARD_SECRECY_TLS12"
	// Forward secrecy ciphers only with TLS1.1 and 1.2.
	SslPolicy_FORWARD_SECRECY_TLS11 SslPolicy = "FORWARD_SECRECY_TLS11"
	// Forward secrecy ciphers only.
	SslPolicy_FORWARD_SECRECY SslPolicy = "FORWARD_SECRECY"
	// TLS1.2 only and no SHA ciphers.
	SslPolicy_TLS12 SslPolicy = "TLS12"
	// TLS1.2 only with all ciphers.
	SslPolicy_TLS12_EXT SslPolicy = "TLS12_EXT"
	// TLS1.1 and 1.2 with all ciphers.
	SslPolicy_TLS11 SslPolicy = "TLS11"
	// Support for DES-CBC3-SHA.
	//
	// Do not use this security policy unless you must support a legacy client
	// that requires the DES-CBC3-SHA cipher, which is a weak cipher.
	SslPolicy_LEGACY SslPolicy = "LEGACY"
)

// Properties to reference an existing target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupAttributes := &targetGroupAttributes{
//   	targetGroupArn: jsii.String("targetGroupArn"),
//
//   	// the properties below are optional
//   	loadBalancerArns: jsii.String("loadBalancerArns"),
//   }
//
type TargetGroupAttributes struct {
	// ARN of the target group.
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
	// A Token representing the list of ARNs for the load balancer routing to this target group.
	LoadBalancerArns *string `field:"optional" json:"loadBalancerArns" yaml:"loadBalancerArns"`
}

// Define the target of a load balancer.
type TargetGroupBase interface {
	constructs.Construct
	ITargetGroup
	// Default port configured for members of this target group.
	DefaultPort() *float64
	// Full name of first load balancer.
	//
	// This identifier is emitted as a dimensions of the metrics of this target
	// group.
	//
	// Example value: `app/my-load-balancer/123456789`.
	FirstLoadBalancerFullName() *string
	HealthCheck() *HealthCheck
	SetHealthCheck(val *HealthCheck)
	// A token representing a list of ARNs of the load balancers that route traffic to this target group.
	LoadBalancerArns() *string
	// List of constructs that need to be depended on to ensure the TargetGroup is associated to a load balancer.
	LoadBalancerAttached() constructs.IDependable
	// Configurable dependable with all resources that lead to load balancer attachment.
	LoadBalancerAttachedDependencies() constructs.DependencyGroup
	// The tree node.
	Node() constructs.Node
	// The ARN of the target group.
	TargetGroupArn() *string
	// The full name of the target group.
	TargetGroupFullName() *string
	// ARNs of load balancers load balancing to this TargetGroup.
	TargetGroupLoadBalancerArns() *[]*string
	// The name of the target group.
	TargetGroupName() *string
	// The types of the directly registered members of this target group.
	TargetType() TargetType
	SetTargetType(val TargetType)
	// Register the given load balancing target as part of this group.
	AddLoadBalancerTarget(props *LoadBalancerTargetProps)
	// Set/replace the target group's health check.
	ConfigureHealthCheck(healthCheck *HealthCheck)
	// Set a non-standard attribute on the target group.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
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

func (t *jsiiProxy_TargetGroupBase) AddLoadBalancerTarget(props *LoadBalancerTargetProps) {
	_jsii_.InvokeVoid(
		t,
		"addLoadBalancerTarget",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_TargetGroupBase) ConfigureHealthCheck(healthCheck *HealthCheck) {
	_jsii_.InvokeVoid(
		t,
		"configureHealthCheck",
		[]interface{}{healthCheck},
	)
}

func (t *jsiiProxy_TargetGroupBase) SetAttribute(key *string, value *string) {
	_jsii_.InvokeVoid(
		t,
		"setAttribute",
		[]interface{}{key, value},
	)
}

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
	// round_robin.
	TargetGroupLoadBalancingAlgorithmType_ROUND_ROBIN TargetGroupLoadBalancingAlgorithmType = "ROUND_ROBIN"
	// least_outstanding_requests.
	TargetGroupLoadBalancingAlgorithmType_LEAST_OUTSTANDING_REQUESTS TargetGroupLoadBalancingAlgorithmType = "LEAST_OUTSTANDING_REQUESTS"
)

// How to interpret the load balancing target identifiers.
//
// Example:
//   var vpc vpc
//
//
//   tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &applicationTargetGroupProps{
//   	targetType: elbv2.targetType_IP,
//   	port: jsii.Number(50051),
//   	protocol: elbv2.applicationProtocol_HTTP,
//   	protocolVersion: elbv2.applicationProtocolVersion_GRPC,
//   	healthCheck: &healthCheck{
//   		enabled: jsii.Boolean(true),
//   		healthyGrpcCodes: jsii.String("0-99"),
//   	},
//   	vpc: vpc,
//   })
//
type TargetType string

const (
	// Targets identified by instance ID.
	TargetType_INSTANCE TargetType = "INSTANCE"
	// Targets identified by IP address.
	TargetType_IP TargetType = "IP"
	// Target is a single Lambda Function.
	TargetType_LAMBDA TargetType = "LAMBDA"
	// Target is a single Application Load Balancer.
	TargetType_ALB TargetType = "ALB"
)

// What to do with unauthenticated requests.
type UnauthenticatedAction string

const (
	// Return an HTTP 401 Unauthorized error.
	UnauthenticatedAction_DENY UnauthenticatedAction = "DENY"
	// Allow the request to be forwarded to the target.
	UnauthenticatedAction_ALLOW UnauthenticatedAction = "ALLOW"
	// Redirect the request to the IdP authorization endpoint.
	UnauthenticatedAction_AUTHENTICATE UnauthenticatedAction = "AUTHENTICATE"
)

// A Target Group and weight combination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationTargetGroup applicationTargetGroup
//
//   weightedTargetGroup := &weightedTargetGroup{
//   	targetGroup: applicationTargetGroup,
//
//   	// the properties below are optional
//   	weight: jsii.Number(123),
//   }
//
type WeightedTargetGroup struct {
	// The target group.
	TargetGroup IApplicationTargetGroup `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// The target group's weight.
	//
	// Range is [0..1000).
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

