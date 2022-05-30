package awselasticloadbalancing

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::ElasticLoadBalancing::LoadBalancer`.
//
// Specifies a Classic Load Balancer.
//
// You can specify the `AvailabilityZones` or `Subnets` property, but not both.
//
// If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the VPC-gateway attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//
//   cfnLoadBalancer := awscdk.Aws_elasticloadbalancing.NewCfnLoadBalancer(this, jsii.String("MyCfnLoadBalancer"), &cfnLoadBalancerProps{
//   	listeners: []interface{}{
//   		&listenersProperty{
//   			instancePort: jsii.String("instancePort"),
//   			loadBalancerPort: jsii.String("loadBalancerPort"),
//   			protocol: jsii.String("protocol"),
//
//   			// the properties below are optional
//   			instanceProtocol: jsii.String("instanceProtocol"),
//   			policyNames: []*string{
//   				jsii.String("policyNames"),
//   			},
//   			sslCertificateId: jsii.String("sslCertificateId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	accessLoggingPolicy: &accessLoggingPolicyProperty{
//   		enabled: jsii.Boolean(false),
//   		s3BucketName: jsii.String("s3BucketName"),
//
//   		// the properties below are optional
//   		emitInterval: jsii.Number(123),
//   		s3BucketPrefix: jsii.String("s3BucketPrefix"),
//   	},
//   	appCookieStickinessPolicy: []interface{}{
//   		&appCookieStickinessPolicyProperty{
//   			cookieName: jsii.String("cookieName"),
//   			policyName: jsii.String("policyName"),
//   		},
//   	},
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	connectionDrainingPolicy: &connectionDrainingPolicyProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		timeout: jsii.Number(123),
//   	},
//   	connectionSettings: &connectionSettingsProperty{
//   		idleTimeout: jsii.Number(123),
//   	},
//   	crossZone: jsii.Boolean(false),
//   	healthCheck: &healthCheckProperty{
//   		healthyThreshold: jsii.String("healthyThreshold"),
//   		interval: jsii.String("interval"),
//   		target: jsii.String("target"),
//   		timeout: jsii.String("timeout"),
//   		unhealthyThreshold: jsii.String("unhealthyThreshold"),
//   	},
//   	instances: []*string{
//   		jsii.String("instances"),
//   	},
//   	lbCookieStickinessPolicy: []interface{}{
//   		&lBCookieStickinessPolicyProperty{
//   			cookieExpirationPeriod: jsii.String("cookieExpirationPeriod"),
//   			policyName: jsii.String("policyName"),
//   		},
//   	},
//   	loadBalancerName: jsii.String("loadBalancerName"),
//   	policies: []interface{}{
//   		&policiesProperty{
//   			attributes: []interface{}{
//   				attributes,
//   			},
//   			policyName: jsii.String("policyName"),
//   			policyType: jsii.String("policyType"),
//
//   			// the properties below are optional
//   			instancePorts: []*string{
//   				jsii.String("instancePorts"),
//   			},
//   			loadBalancerPorts: []*string{
//   				jsii.String("loadBalancerPorts"),
//   			},
//   		},
//   	},
//   	scheme: jsii.String("scheme"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
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
//   })
//
type CfnLoadBalancer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Information about where and how access logs are stored for the load balancer.
	AccessLoggingPolicy() interface{}
	SetAccessLoggingPolicy(val interface{})
	// Information about a policy for application-controlled session stickiness.
	AppCookieStickinessPolicy() interface{}
	SetAppCookieStickinessPolicy(val interface{})
	// The name of the Route 53 hosted zone that is associated with the load balancer.
	//
	// Internal-facing load balancers don't use this value, use `DNSName` instead.
	AttrCanonicalHostedZoneName() *string
	// The ID of the Route 53 hosted zone name that is associated with the load balancer.
	AttrCanonicalHostedZoneNameId() *string
	// The DNS name for the load balancer.
	AttrDnsName() *string
	// The name of the security group that you can use as part of your inbound rules for your load balancer's back-end instances.
	AttrSourceSecurityGroupGroupName() *string
	// The owner of the source security group.
	AttrSourceSecurityGroupOwnerAlias() *string
	// The Availability Zones for the load balancer. For load balancers in a VPC, specify `Subnets` instead.
	//
	// Update requires replacement if you did not previously specify an Availability Zone or if you are removing all Availability Zones. Otherwise, update requires no interruption.
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance.
	//
	// For more information, see [Configure Connection Draining](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html) in the *Classic Load Balancers Guide* .
	ConnectionDrainingPolicy() interface{}
	SetConnectionDrainingPolicy(val interface{})
	// If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration.
	//
	// By default, Elastic Load Balancing maintains a 60-second idle connection timeout for both front-end and back-end connections of your load balancer. For more information, see [Configure Idle Connection Timeout](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html) in the *Classic Load Balancers Guide* .
	ConnectionSettings() interface{}
	SetConnectionSettings(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones.
	//
	// For more information, see [Configure Cross-Zone Load Balancing](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html) in the *Classic Load Balancers Guide* .
	CrossZone() interface{}
	SetCrossZone(val interface{})
	// The health check settings to use when evaluating the health of your EC2 instances.
	//
	// Update requires replacement if you did not previously specify health check settings or if you are removing the health check settings. Otherwise, update requires no interruption.
	HealthCheck() interface{}
	SetHealthCheck(val interface{})
	// The IDs of the instances for the load balancer.
	Instances() *[]*string
	SetInstances(val *[]*string)
	// Information about a policy for duration-based session stickiness.
	LbCookieStickinessPolicy() interface{}
	SetLbCookieStickinessPolicy(val interface{})
	// The listeners for the load balancer. You can specify at most one listener per port.
	//
	// If you update the properties for a listener, AWS CloudFormation deletes the existing listener and creates a new one with the specified properties. While the new listener is being created, clients cannot connect to the load balancer.
	Listeners() interface{}
	SetListeners(val interface{})
	// The name of the load balancer.
	//
	// This name must be unique within your set of load balancers for the region.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) . If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	LoadBalancerName() *string
	SetLoadBalancerName(val *string)
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
	// The policies defined for your Classic Load Balancer.
	//
	// Specify only back-end server policies.
	Policies() interface{}
	SetPolicies(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The type of load balancer. Valid only for load balancers in a VPC.
	//
	// If `Scheme` is `internet-facing` , the load balancer has a public DNS name that resolves to a public IP address.
	//
	// If `Scheme` is `internal` , the load balancer has a public DNS name that resolves to a private IP address.
	Scheme() *string
	SetScheme(val *string)
	// The security groups for the load balancer.
	//
	// Valid only for load balancers in a VPC.
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The IDs of the subnets for the load balancer. You can specify at most one subnet per Availability Zone.
	//
	// Update requires replacement if you did not previously specify a subnet or if you are removing all subnets. Otherwise, update requires no interruption. To update to a different subnet in the current Availability Zone, you must first update to a subnet in a different Availability Zone, then update to the new subnet in the original Availability Zone.
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	// The tags associated with a load balancer.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

func (j *jsiiProxy_CfnLoadBalancer) AccessLoggingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLoggingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AppCookieStickinessPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appCookieStickinessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AttrCanonicalHostedZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCanonicalHostedZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AttrCanonicalHostedZoneNameId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCanonicalHostedZoneNameId",
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

func (j *jsiiProxy_CfnLoadBalancer) AttrSourceSecurityGroupGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceSecurityGroupGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AttrSourceSecurityGroupOwnerAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSourceSecurityGroupOwnerAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
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

func (j *jsiiProxy_CfnLoadBalancer) ConnectionDrainingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionDrainingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) ConnectionSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionSettings",
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

func (j *jsiiProxy_CfnLoadBalancer) CrossZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) HealthCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Instances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) LbCookieStickinessPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbCookieStickinessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Listeners() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
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

func (j *jsiiProxy_CfnLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoadBalancer) Policies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policies",
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

func (j *jsiiProxy_CfnLoadBalancer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ElasticLoadBalancing::LoadBalancer`.
func NewCfnLoadBalancer(scope constructs.Construct, id *string, props *CfnLoadBalancerProps) CfnLoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ElasticLoadBalancing::LoadBalancer`.
func NewCfnLoadBalancer_Override(c CfnLoadBalancer, scope constructs.Construct, id *string, props *CfnLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetAccessLoggingPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"accessLoggingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetAppCookieStickinessPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"appCookieStickinessPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetAvailabilityZones(val *[]*string) {
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetConnectionDrainingPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"connectionDrainingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetConnectionSettings(val interface{}) {
	_jsii_.Set(
		j,
		"connectionSettings",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetCrossZone(val interface{}) {
	_jsii_.Set(
		j,
		"crossZone",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetHealthCheck(val interface{}) {
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetInstances(val *[]*string) {
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetLbCookieStickinessPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"lbCookieStickinessPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetListeners(val interface{}) {
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetLoadBalancerName(val *string) {
	_jsii_.Set(
		j,
		"loadBalancerName",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer) SetPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"policies",
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

func (j *jsiiProxy_CfnLoadBalancer) SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
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
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
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
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
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
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
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

// Specifies where and how access logs are stored for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLoggingPolicyProperty := &accessLoggingPolicyProperty{
//   	enabled: jsii.Boolean(false),
//   	s3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	emitInterval: jsii.Number(123),
//   	s3BucketPrefix: jsii.String("s3BucketPrefix"),
//   }
//
type CfnLoadBalancer_AccessLoggingPolicyProperty struct {
	// Specifies whether access logs are enabled for the load balancer.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the Amazon S3 bucket where the access logs are stored.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The interval for publishing the access logs. You can specify an interval of either 5 minutes or 60 minutes.
	//
	// Default: 60 minutes.
	EmitInterval *float64 `field:"optional" json:"emitInterval" yaml:"emitInterval"`
	// The logical hierarchy you created for your Amazon S3 bucket, for example `my-bucket-prefix/prod` .
	//
	// If the prefix is not provided, the log is placed at the root level of the bucket.
	S3BucketPrefix *string `field:"optional" json:"s3BucketPrefix" yaml:"s3BucketPrefix"`
}

// Specifies a policy for application-controlled session stickiness for your Classic Load Balancer.
//
// To associate a policy with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appCookieStickinessPolicyProperty := &appCookieStickinessPolicyProperty{
//   	cookieName: jsii.String("cookieName"),
//   	policyName: jsii.String("policyName"),
//   }
//
type CfnLoadBalancer_AppCookieStickinessPolicyProperty struct {
	// The name of the application cookie used for stickiness.
	CookieName *string `field:"required" json:"cookieName" yaml:"cookieName"`
	// The mnemonic name for the policy being created.
	//
	// The name must be unique within a set of policies for this load balancer.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

// Specifies the connection draining settings for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionDrainingPolicyProperty := &connectionDrainingPolicyProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	timeout: jsii.Number(123),
//   }
//
type CfnLoadBalancer_ConnectionDrainingPolicyProperty struct {
	// Specifies whether connection draining is enabled for the load balancer.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The maximum time, in seconds, to keep the existing connections open before deregistering the instances.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

// Specifies the idle timeout value for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionSettingsProperty := &connectionSettingsProperty{
//   	idleTimeout: jsii.Number(123),
//   }
//
type CfnLoadBalancer_ConnectionSettingsProperty struct {
	// The time, in seconds, that the connection is allowed to be idle (no data has been sent over the connection) before it is closed by the load balancer.
	IdleTimeout *float64 `field:"required" json:"idleTimeout" yaml:"idleTimeout"`
}

// Specifies health check settings for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckProperty := &healthCheckProperty{
//   	healthyThreshold: jsii.String("healthyThreshold"),
//   	interval: jsii.String("interval"),
//   	target: jsii.String("target"),
//   	timeout: jsii.String("timeout"),
//   	unhealthyThreshold: jsii.String("unhealthyThreshold"),
//   }
//
type CfnLoadBalancer_HealthCheckProperty struct {
	// The number of consecutive health checks successes required before moving the instance to the `Healthy` state.
	HealthyThreshold *string `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The approximate interval, in seconds, between health checks of an individual instance.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The instance being checked.
	//
	// The protocol is either TCP, HTTP, HTTPS, or SSL. The range of valid ports is one (1) through 65535.
	//
	// TCP is the default, specified as a TCP: port pair, for example "TCP:5000". In this case, a health check simply attempts to open a TCP connection to the instance on the specified port. Failure to connect within the configured timeout is considered unhealthy.
	//
	// SSL is also specified as SSL: port pair, for example, SSL:5000.
	//
	// For HTTP/HTTPS, you must include a ping path in the string. HTTP is specified as a HTTP:port;/;PathToPing; grouping, for example "HTTP:80/weather/us/wa/seattle". In this case, a HTTP GET request is issued to the instance on the given port and path. Any answer other than "200 OK" within the timeout period is considered unhealthy.
	//
	// The total length of the HTTP ping target must be 1024 16-bit Unicode characters or less.
	Target *string `field:"required" json:"target" yaml:"target"`
	// The amount of time, in seconds, during which no response means a failed health check.
	//
	// This value must be less than the `Interval` value.
	Timeout *string `field:"required" json:"timeout" yaml:"timeout"`
	// The number of consecutive health check failures required before moving the instance to the `Unhealthy` state.
	UnhealthyThreshold *string `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// Specifies a policy for duration-based session stickiness for your Classic Load Balancer.
//
// To associate a policy with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lBCookieStickinessPolicyProperty := &lBCookieStickinessPolicyProperty{
//   	cookieExpirationPeriod: jsii.String("cookieExpirationPeriod"),
//   	policyName: jsii.String("policyName"),
//   }
//
type CfnLoadBalancer_LBCookieStickinessPolicyProperty struct {
	// The time period, in seconds, after which the cookie should be considered stale.
	//
	// If this parameter is not specified, the stickiness session lasts for the duration of the browser session.
	CookieExpirationPeriod *string `field:"optional" json:"cookieExpirationPeriod" yaml:"cookieExpirationPeriod"`
	// The name of the policy.
	//
	// This name must be unique within the set of policies for this load balancer.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

// Specifies a listener for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenersProperty := &listenersProperty{
//   	instancePort: jsii.String("instancePort"),
//   	loadBalancerPort: jsii.String("loadBalancerPort"),
//   	protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	instanceProtocol: jsii.String("instanceProtocol"),
//   	policyNames: []*string{
//   		jsii.String("policyNames"),
//   	},
//   	sslCertificateId: jsii.String("sslCertificateId"),
//   }
//
type CfnLoadBalancer_ListenersProperty struct {
	// The port on which the instance is listening.
	InstancePort *string `field:"required" json:"instancePort" yaml:"instancePort"`
	// The port on which the load balancer is listening.
	//
	// On EC2-VPC, you can specify any port from the range 1-65535. On EC2-Classic, you can specify any port from the following list: 25, 80, 443, 465, 587, 1024-65535.
	LoadBalancerPort *string `field:"required" json:"loadBalancerPort" yaml:"loadBalancerPort"`
	// The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.
	//
	// If the front-end protocol is TCP or SSL, the back-end protocol must be TCP or SSL. If the front-end protocol is HTTP or HTTPS, the back-end protocol must be HTTP or HTTPS.
	//
	// If there is another listener with the same `InstancePort` whose `InstanceProtocol` is secure, (HTTPS or SSL), the listener's `InstanceProtocol` must also be secure.
	//
	// If there is another listener with the same `InstancePort` whose `InstanceProtocol` is HTTP or TCP, the listener's `InstanceProtocol` must be HTTP or TCP.
	InstanceProtocol *string `field:"optional" json:"instanceProtocol" yaml:"instanceProtocol"`
	// The names of the policies to associate with the listener.
	PolicyNames *[]*string `field:"optional" json:"policyNames" yaml:"policyNames"`
	// The Amazon Resource Name (ARN) of the server certificate.
	SslCertificateId *string `field:"optional" json:"sslCertificateId" yaml:"sslCertificateId"`
}

// Specifies policies for your Classic Load Balancer.
//
// To associate policies with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//
//   policiesProperty := &policiesProperty{
//   	attributes: []interface{}{
//   		attributes,
//   	},
//   	policyName: jsii.String("policyName"),
//   	policyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	instancePorts: []*string{
//   		jsii.String("instancePorts"),
//   	},
//   	loadBalancerPorts: []*string{
//   		jsii.String("loadBalancerPorts"),
//   	},
//   }
//
type CfnLoadBalancer_PoliciesProperty struct {
	// The policy attributes.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the policy.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The name of the policy type.
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// The instance ports for the policy.
	//
	// Required only for some policy types.
	InstancePorts *[]*string `field:"optional" json:"instancePorts" yaml:"instancePorts"`
	// The load balancer ports for the policy.
	//
	// Required only for some policy types.
	LoadBalancerPorts *[]*string `field:"optional" json:"loadBalancerPorts" yaml:"loadBalancerPorts"`
}

// Properties for defining a `CfnLoadBalancer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//
//   cfnLoadBalancerProps := &cfnLoadBalancerProps{
//   	listeners: []interface{}{
//   		&listenersProperty{
//   			instancePort: jsii.String("instancePort"),
//   			loadBalancerPort: jsii.String("loadBalancerPort"),
//   			protocol: jsii.String("protocol"),
//
//   			// the properties below are optional
//   			instanceProtocol: jsii.String("instanceProtocol"),
//   			policyNames: []*string{
//   				jsii.String("policyNames"),
//   			},
//   			sslCertificateId: jsii.String("sslCertificateId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	accessLoggingPolicy: &accessLoggingPolicyProperty{
//   		enabled: jsii.Boolean(false),
//   		s3BucketName: jsii.String("s3BucketName"),
//
//   		// the properties below are optional
//   		emitInterval: jsii.Number(123),
//   		s3BucketPrefix: jsii.String("s3BucketPrefix"),
//   	},
//   	appCookieStickinessPolicy: []interface{}{
//   		&appCookieStickinessPolicyProperty{
//   			cookieName: jsii.String("cookieName"),
//   			policyName: jsii.String("policyName"),
//   		},
//   	},
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	connectionDrainingPolicy: &connectionDrainingPolicyProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		timeout: jsii.Number(123),
//   	},
//   	connectionSettings: &connectionSettingsProperty{
//   		idleTimeout: jsii.Number(123),
//   	},
//   	crossZone: jsii.Boolean(false),
//   	healthCheck: &healthCheckProperty{
//   		healthyThreshold: jsii.String("healthyThreshold"),
//   		interval: jsii.String("interval"),
//   		target: jsii.String("target"),
//   		timeout: jsii.String("timeout"),
//   		unhealthyThreshold: jsii.String("unhealthyThreshold"),
//   	},
//   	instances: []*string{
//   		jsii.String("instances"),
//   	},
//   	lbCookieStickinessPolicy: []interface{}{
//   		&lBCookieStickinessPolicyProperty{
//   			cookieExpirationPeriod: jsii.String("cookieExpirationPeriod"),
//   			policyName: jsii.String("policyName"),
//   		},
//   	},
//   	loadBalancerName: jsii.String("loadBalancerName"),
//   	policies: []interface{}{
//   		&policiesProperty{
//   			attributes: []interface{}{
//   				attributes,
//   			},
//   			policyName: jsii.String("policyName"),
//   			policyType: jsii.String("policyType"),
//
//   			// the properties below are optional
//   			instancePorts: []*string{
//   				jsii.String("instancePorts"),
//   			},
//   			loadBalancerPorts: []*string{
//   				jsii.String("loadBalancerPorts"),
//   			},
//   		},
//   	},
//   	scheme: jsii.String("scheme"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
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
//   }
//
type CfnLoadBalancerProps struct {
	// The listeners for the load balancer. You can specify at most one listener per port.
	//
	// If you update the properties for a listener, AWS CloudFormation deletes the existing listener and creates a new one with the specified properties. While the new listener is being created, clients cannot connect to the load balancer.
	Listeners interface{} `field:"required" json:"listeners" yaml:"listeners"`
	// Information about where and how access logs are stored for the load balancer.
	AccessLoggingPolicy interface{} `field:"optional" json:"accessLoggingPolicy" yaml:"accessLoggingPolicy"`
	// Information about a policy for application-controlled session stickiness.
	AppCookieStickinessPolicy interface{} `field:"optional" json:"appCookieStickinessPolicy" yaml:"appCookieStickinessPolicy"`
	// The Availability Zones for the load balancer. For load balancers in a VPC, specify `Subnets` instead.
	//
	// Update requires replacement if you did not previously specify an Availability Zone or if you are removing all Availability Zones. Otherwise, update requires no interruption.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance.
	//
	// For more information, see [Configure Connection Draining](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html) in the *Classic Load Balancers Guide* .
	ConnectionDrainingPolicy interface{} `field:"optional" json:"connectionDrainingPolicy" yaml:"connectionDrainingPolicy"`
	// If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration.
	//
	// By default, Elastic Load Balancing maintains a 60-second idle connection timeout for both front-end and back-end connections of your load balancer. For more information, see [Configure Idle Connection Timeout](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html) in the *Classic Load Balancers Guide* .
	ConnectionSettings interface{} `field:"optional" json:"connectionSettings" yaml:"connectionSettings"`
	// If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones.
	//
	// For more information, see [Configure Cross-Zone Load Balancing](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html) in the *Classic Load Balancers Guide* .
	CrossZone interface{} `field:"optional" json:"crossZone" yaml:"crossZone"`
	// The health check settings to use when evaluating the health of your EC2 instances.
	//
	// Update requires replacement if you did not previously specify health check settings or if you are removing the health check settings. Otherwise, update requires no interruption.
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The IDs of the instances for the load balancer.
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// Information about a policy for duration-based session stickiness.
	LbCookieStickinessPolicy interface{} `field:"optional" json:"lbCookieStickinessPolicy" yaml:"lbCookieStickinessPolicy"`
	// The name of the load balancer.
	//
	// This name must be unique within your set of load balancers for the region.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) . If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	LoadBalancerName *string `field:"optional" json:"loadBalancerName" yaml:"loadBalancerName"`
	// The policies defined for your Classic Load Balancer.
	//
	// Specify only back-end server policies.
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// The type of load balancer. Valid only for load balancers in a VPC.
	//
	// If `Scheme` is `internet-facing` , the load balancer has a public DNS name that resolves to a public IP address.
	//
	// If `Scheme` is `internal` , the load balancer has a public DNS name that resolves to a private IP address.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// The security groups for the load balancer.
	//
	// Valid only for load balancers in a VPC.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the subnets for the load balancer. You can specify at most one subnet per Availability Zone.
	//
	// Update requires replacement if you did not previously specify a subnet or if you are removing all subnets. Otherwise, update requires no interruption. To update to a different subnet in the current Availability Zone, you must first update to a subnet in a different Availability Zone, then update to the new subnet in the original Availability Zone.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The tags associated with a load balancer.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Describe the health check to a load balancer.
//
// Example:
//   var vpc iVpc
//
//   var myAutoScalingGroup autoScalingGroup
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   	healthCheck: &healthCheck{
//   		port: jsii.Number(80),
//   	},
//   })
//   lb.addTarget(myAutoScalingGroup)
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//
type HealthCheck struct {
	// What port number to health check on.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// After how many successful checks is an instance considered healthy.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Number of seconds between health checks.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// What path to use for HTTP or HTTPS health check (must return 200).
	//
	// For SSL and TCP health checks, accepting connections is enough to be considered
	// healthy.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// What protocol to use for health checking.
	//
	// The protocol is automatically determined from the port if it's not supplied.
	Protocol LoadBalancingProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Health check timeout.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// After how many unsuccessful checks is an instance considered unhealthy.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

// Interface that is going to be implemented by constructs that you can load balance to.
type ILoadBalancerTarget interface {
	awsec2.IConnectable
	// Attach load-balanced target to a classic ELB.
	AttachToClassicLB(loadBalancer LoadBalancer)
}

// The jsii proxy for ILoadBalancerTarget
type jsiiProxy_ILoadBalancerTarget struct {
	internal.Type__awsec2IConnectable
}

func (i *jsiiProxy_ILoadBalancerTarget) AttachToClassicLB(loadBalancer LoadBalancer) {
	_jsii_.InvokeVoid(
		i,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

// Reference to a listener's port just created.
//
// This implements IConnectable with a default port (the port that an ELB
// listener was just created on) for a given security group so that it can be
// conveniently used just like any Connectable. E.g:
//
//     const listener = elb.addListener(...);
//
//     listener.connections.allowDefaultPortFromAnyIPv4();
//     // or
//     instance.connections.allowToDefaultPort(listener);
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var port port
//   var securityGroup securityGroup
//
//   listenerPort := awscdk.Aws_elasticloadbalancing.NewListenerPort(securityGroup, port)
//
type ListenerPort interface {
	awsec2.IConnectable
	// The network connections associated with this resource.
	Connections() awsec2.Connections
}

// The jsii proxy struct for ListenerPort
type jsiiProxy_ListenerPort struct {
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_ListenerPort) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


func NewListenerPort(securityGroup awsec2.ISecurityGroup, defaultPort awsec2.Port) ListenerPort {
	_init_.Initialize()

	j := jsiiProxy_ListenerPort{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.ListenerPort",
		[]interface{}{securityGroup, defaultPort},
		&j,
	)

	return &j
}

func NewListenerPort_Override(l ListenerPort, securityGroup awsec2.ISecurityGroup, defaultPort awsec2.Port) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.ListenerPort",
		[]interface{}{securityGroup, defaultPort},
		l,
	)
}

// A load balancer with a single listener.
//
// Routes to a fleet of of instances in a VPC.
//
// Example:
//   var vpc iVpc
//
//   var myAutoScalingGroup autoScalingGroup
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   	healthCheck: &healthCheck{
//   		port: jsii.Number(80),
//   	},
//   })
//   lb.addTarget(myAutoScalingGroup)
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//
type LoadBalancer interface {
	awscdk.Resource
	awsec2.IConnectable
	// Control all connections from and to this load balancer.
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
	// An object controlling specifically the connections for each listener added to this load balancer.
	ListenerPorts() *[]ListenerPort
	LoadBalancerCanonicalHostedZoneName() *string
	LoadBalancerCanonicalHostedZoneNameId() *string
	LoadBalancerDnsName() *string
	LoadBalancerName() *string
	LoadBalancerSourceSecurityGroupGroupName() *string
	LoadBalancerSourceSecurityGroupOwnerAlias() *string
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
	// Add a backend to the load balancer.
	//
	// Returns: A ListenerPort object that controls connections to the listener port.
	AddListener(listener *LoadBalancerListener) ListenerPort
	AddTarget(target ILoadBalancerTarget)
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
}

// The jsii proxy struct for LoadBalancer
type jsiiProxy_LoadBalancer struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_LoadBalancer) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) ListenerPorts() *[]ListenerPort {
	var returns *[]ListenerPort
	_jsii_.Get(
		j,
		"listenerPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LoadBalancerCanonicalHostedZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LoadBalancerCanonicalHostedZoneNameId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneNameId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LoadBalancerSourceSecurityGroupGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerSourceSecurityGroupGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) LoadBalancerSourceSecurityGroupOwnerAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerSourceSecurityGroupOwnerAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewLoadBalancer(scope constructs.Construct, id *string, props *LoadBalancerProps) LoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_LoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.LoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLoadBalancer_Override(l LoadBalancer, scope constructs.Construct, id *string, props *LoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.LoadBalancer",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func LoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancing.LoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func LoadBalancer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancing.LoadBalancer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) AddListener(listener *LoadBalancerListener) ListenerPort {
	var returns ListenerPort

	_jsii_.Invoke(
		l,
		"addListener",
		[]interface{}{listener},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) AddTarget(target ILoadBalancerTarget) {
	_jsii_.InvokeVoid(
		l,
		"addTarget",
		[]interface{}{target},
	)
}

func (l *jsiiProxy_LoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LoadBalancer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Add a backend to the load balancer.
//
// Example:
//   var vpc iVpc
//
//   var myAutoScalingGroup autoScalingGroup
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   	healthCheck: &healthCheck{
//   		port: jsii.Number(80),
//   	},
//   })
//   lb.addTarget(myAutoScalingGroup)
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//
type LoadBalancerListener struct {
	// External listening port.
	ExternalPort *float64 `field:"required" json:"externalPort" yaml:"externalPort"`
	// Allow connections to the load balancer from the given set of connection peers.
	//
	// By default, connections will be allowed from anywhere. Set this to an empty list
	// to deny connections, or supply a custom list of peers to allow connections from
	// (IP ranges or security groups).
	AllowConnectionsFrom *[]awsec2.IConnectable `field:"optional" json:"allowConnectionsFrom" yaml:"allowConnectionsFrom"`
	// What public protocol to use for load balancing.
	//
	// Either 'tcp', 'ssl', 'http' or 'https'.
	//
	// May be omitted if the external port is either 80 or 443.
	ExternalProtocol LoadBalancingProtocol `field:"optional" json:"externalProtocol" yaml:"externalProtocol"`
	// Instance listening port.
	//
	// Same as the externalPort if not specified.
	InternalPort *float64 `field:"optional" json:"internalPort" yaml:"internalPort"`
	// What public protocol to use for load balancing.
	//
	// Either 'tcp', 'ssl', 'http' or 'https'.
	//
	// May be omitted if the internal port is either 80 or 443.
	//
	// The instance protocol is 'tcp' if the front-end protocol
	// is 'tcp' or 'ssl', the instance protocol is 'http' if the
	// front-end protocol is 'https'.
	InternalProtocol LoadBalancingProtocol `field:"optional" json:"internalProtocol" yaml:"internalProtocol"`
	// SSL policy names.
	PolicyNames *[]*string `field:"optional" json:"policyNames" yaml:"policyNames"`
	// the ARN of the SSL certificate.
	SslCertificateArn *string `field:"optional" json:"sslCertificateArn" yaml:"sslCertificateArn"`
}

// Construction properties for a LoadBalancer.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
//   	vpc: vpc,
//   })
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//   lb.addTarget(service)
//
type LoadBalancerProps struct {
	// VPC network of the fleet instances.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Enable Loadbalancer access logs Can be used to avoid manual work as aws console Required S3 bucket name , enabled flag Can add interval for pushing log Can set bucket prefix in order to provide folder name inside bucket.
	AccessLoggingPolicy *CfnLoadBalancer_AccessLoggingPolicyProperty `field:"optional" json:"accessLoggingPolicy" yaml:"accessLoggingPolicy"`
	// Whether cross zone load balancing is enabled.
	//
	// This controls whether the load balancer evenly distributes requests
	// across each availability zone.
	CrossZone *bool `field:"optional" json:"crossZone" yaml:"crossZone"`
	// Health check settings for the load balancing targets.
	//
	// Not required but recommended.
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Whether this is an internet-facing Load Balancer.
	//
	// This controls whether the LB has a public IP address assigned. It does
	// not open up the Load Balancer's security groups to public internet access.
	InternetFacing *bool `field:"optional" json:"internetFacing" yaml:"internetFacing"`
	// What listeners to set up for the load balancer.
	//
	// Can also be added by .addListener()
	Listeners *[]*LoadBalancerListener `field:"optional" json:"listeners" yaml:"listeners"`
	// Which subnets to deploy the load balancer.
	//
	// Can be used to define a specific set of subnets to deploy the load balancer to.
	// Useful multiple public or private subnets are covering the same availability zone.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// What targets to load balance to.
	//
	// Can also be added by .addTarget()
	Targets *[]ILoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

type LoadBalancingProtocol string

const (
	LoadBalancingProtocol_TCP LoadBalancingProtocol = "TCP"
	LoadBalancingProtocol_SSL LoadBalancingProtocol = "SSL"
	LoadBalancingProtocol_HTTP LoadBalancingProtocol = "HTTP"
	LoadBalancingProtocol_HTTPS LoadBalancingProtocol = "HTTPS"
)

