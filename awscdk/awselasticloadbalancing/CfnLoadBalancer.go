package awselasticloadbalancing

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a Classic Load Balancer.
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
//   cfnLoadBalancer := awscdk.Aws_elasticloadbalancing.NewCfnLoadBalancer(this, jsii.String("MyCfnLoadBalancer"), &CfnLoadBalancerProps{
//   	Listeners: []interface{}{
//   		&ListenersProperty{
//   			InstancePort: jsii.String("instancePort"),
//   			LoadBalancerPort: jsii.String("loadBalancerPort"),
//   			Protocol: jsii.String("protocol"),
//
//   			// the properties below are optional
//   			InstanceProtocol: jsii.String("instanceProtocol"),
//   			PolicyNames: []*string{
//   				jsii.String("policyNames"),
//   			},
//   			SslCertificateId: jsii.String("sslCertificateId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AccessLoggingPolicy: &AccessLoggingPolicyProperty{
//   		Enabled: jsii.Boolean(false),
//   		S3BucketName: jsii.String("s3BucketName"),
//
//   		// the properties below are optional
//   		EmitInterval: jsii.Number(123),
//   		S3BucketPrefix: jsii.String("s3BucketPrefix"),
//   	},
//   	AppCookieStickinessPolicy: []interface{}{
//   		&AppCookieStickinessPolicyProperty{
//   			CookieName: jsii.String("cookieName"),
//   			PolicyName: jsii.String("policyName"),
//   		},
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	ConnectionDrainingPolicy: &ConnectionDrainingPolicyProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Timeout: jsii.Number(123),
//   	},
//   	ConnectionSettings: &ConnectionSettingsProperty{
//   		IdleTimeout: jsii.Number(123),
//   	},
//   	CrossZone: jsii.Boolean(false),
//   	HealthCheck: &HealthCheckProperty{
//   		HealthyThreshold: jsii.String("healthyThreshold"),
//   		Interval: jsii.String("interval"),
//   		Target: jsii.String("target"),
//   		Timeout: jsii.String("timeout"),
//   		UnhealthyThreshold: jsii.String("unhealthyThreshold"),
//   	},
//   	Instances: []*string{
//   		jsii.String("instances"),
//   	},
//   	LbCookieStickinessPolicy: []interface{}{
//   		&LBCookieStickinessPolicyProperty{
//   			CookieExpirationPeriod: jsii.String("cookieExpirationPeriod"),
//   			PolicyName: jsii.String("policyName"),
//   		},
//   	},
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//   	Policies: []interface{}{
//   		&PoliciesProperty{
//   			Attributes: []interface{}{
//   				attributes,
//   			},
//   			PolicyName: jsii.String("policyName"),
//   			PolicyType: jsii.String("policyType"),
//
//   			// the properties below are optional
//   			InstancePorts: []*string{
//   				jsii.String("instancePorts"),
//   			},
//   			LoadBalancerPorts: []*string{
//   				jsii.String("loadBalancerPorts"),
//   			},
//   		},
//   	},
//   	Scheme: jsii.String("scheme"),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancing-loadbalancer.html
//
type CfnLoadBalancer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
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
	AttrId() *string
	// The name of the security group that you can use as part of your inbound rules for your load balancer's back-end instances.
	AttrSourceSecurityGroupGroupName() *string
	// The owner of the source security group.
	AttrSourceSecurityGroupOwnerAlias() *string
	// The Availability Zones for a load balancer in a default VPC.
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance.
	ConnectionDrainingPolicy() interface{}
	SetConnectionDrainingPolicy(val interface{})
	// If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration.
	ConnectionSettings() interface{}
	SetConnectionSettings(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones.
	CrossZone() interface{}
	SetCrossZone(val interface{})
	// The health check settings to use when evaluating the health of your EC2 instances.
	HealthCheck() interface{}
	SetHealthCheck(val interface{})
	// The IDs of the instances for the load balancer.
	Instances() *[]*string
	SetInstances(val *[]*string)
	// Information about a policy for duration-based session stickiness.
	LbCookieStickinessPolicy() interface{}
	SetLbCookieStickinessPolicy(val interface{})
	// The listeners for the load balancer.
	//
	// You can specify at most one listener per port.
	Listeners() interface{}
	SetListeners(val interface{})
	// The name of the load balancer.
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
	Policies() interface{}
	SetPolicies(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The type of load balancer.
	//
	// Valid only for load balancers in a VPC.
	Scheme() *string
	SetScheme(val *string)
	// The security groups for the load balancer.
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The IDs of the subnets for the load balancer.
	//
	// You can specify at most one subnet per Availability Zone.
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags associated with a load balancer.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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
	internal.Type__awscdkITaggable
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

func (j *jsiiProxy_CfnLoadBalancer) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
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

func (j *jsiiProxy_CfnLoadBalancer) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
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


func NewCfnLoadBalancer(scope constructs.Construct, id *string, props *CfnLoadBalancerProps) CfnLoadBalancer {
	_init_.Initialize()

	if err := validateNewCfnLoadBalancerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnLoadBalancer_Override(c CfnLoadBalancer, scope constructs.Construct, id *string, props *CfnLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetAccessLoggingPolicy(val interface{}) {
	if err := j.validateSetAccessLoggingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLoggingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetAppCookieStickinessPolicy(val interface{}) {
	if err := j.validateSetAppCookieStickinessPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appCookieStickinessPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetAvailabilityZones(val *[]*string) {
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetConnectionDrainingPolicy(val interface{}) {
	if err := j.validateSetConnectionDrainingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionDrainingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetConnectionSettings(val interface{}) {
	if err := j.validateSetConnectionSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionSettings",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetCrossZone(val interface{}) {
	if err := j.validateSetCrossZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossZone",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetHealthCheck(val interface{}) {
	if err := j.validateSetHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetInstances(val *[]*string) {
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetLbCookieStickinessPolicy(val interface{}) {
	if err := j.validateSetLbCookieStickinessPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lbCookieStickinessPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetListeners(val interface{}) {
	if err := j.validateSetListenersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listeners",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetLoadBalancerName(val *string) {
	_jsii_.Set(
		j,
		"loadBalancerName",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetPolicies(val interface{}) {
	if err := j.validateSetPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetScheme(val *string) {
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_CfnLoadBalancer)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
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

	if err := validateCfnLoadBalancer_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnLoadBalancer_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLoadBalancer_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancing.CfnLoadBalancer",
		"isCfnResource",
		[]interface{}{x},
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

	if err := validateCfnLoadBalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
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
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoadBalancer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoadBalancer) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

