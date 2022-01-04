package awselasticloadbalancing

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::ElasticLoadBalancing::LoadBalancer`.
//
// Specifies a Classic Load Balancer.
//
// You can specify the `AvailabilityZones` or `Subnets` property, but not both.
//
// If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the VPC-gateway attachment.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AccessLoggingPolicy() interface{}
	SetAccessLoggingPolicy(val interface{})
	AppCookieStickinessPolicy() interface{}
	SetAppCookieStickinessPolicy(val interface{})
	AttrCanonicalHostedZoneName() *string
	AttrCanonicalHostedZoneNameId() *string
	AttrDnsName() *string
	AttrSourceSecurityGroupGroupName() *string
	AttrSourceSecurityGroupOwnerAlias() *string
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConnectionDrainingPolicy() interface{}
	SetConnectionDrainingPolicy(val interface{})
	ConnectionSettings() interface{}
	SetConnectionSettings(val interface{})
	CreationStack() *[]*string
	CrossZone() interface{}
	SetCrossZone(val interface{})
	HealthCheck() interface{}
	SetHealthCheck(val interface{})
	Instances() *[]*string
	SetInstances(val *[]*string)
	LbCookieStickinessPolicy() interface{}
	SetLbCookieStickinessPolicy(val interface{})
	Listeners() interface{}
	SetListeners(val interface{})
	LoadBalancerName() *string
	SetLoadBalancerName(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	Policies() interface{}
	SetPolicies(val interface{})
	Ref() *string
	Scheme() *string
	SetScheme(val *string)
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	Stack() awscdk.Stack
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	Tags() awscdk.TagManager
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnLoadBalancer) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnLoadBalancer(scope awscdk.Construct, id *string, props *CfnLoadBalancerProps) CfnLoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancer{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancing.CfnLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ElasticLoadBalancing::LoadBalancer`.
func NewCfnLoadBalancer_Override(c CfnLoadBalancer, scope awscdk.Construct, id *string, props *CfnLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancing.CfnLoadBalancer",
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
// Experimental.
func CfnLoadBalancer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancing.CfnLoadBalancer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLoadBalancer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancing.CfnLoadBalancer",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancing.CfnLoadBalancer",
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
		"monocdk.aws_elasticloadbalancing.CfnLoadBalancer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnLoadBalancer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies where and how access logs are stored for your Classic Load Balancer.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_AccessLoggingPolicyProperty struct {
	// Specifies whether access logs are enabled for the load balancer.
	Enabled interface{} `json:"enabled"`
	// The name of the Amazon S3 bucket where the access logs are stored.
	S3BucketName *string `json:"s3BucketName"`
	// The interval for publishing the access logs. You can specify an interval of either 5 minutes or 60 minutes.
	//
	// Default: 60 minutes
	EmitInterval *float64 `json:"emitInterval"`
	// The logical hierarchy you created for your Amazon S3 bucket, for example `my-bucket-prefix/prod` .
	//
	// If the prefix is not provided, the log is placed at the root level of the bucket.
	S3BucketPrefix *string `json:"s3BucketPrefix"`
}

// Specifies a policy for application-controlled session stickiness for your Classic Load Balancer.
//
// To associate a policy with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_AppCookieStickinessPolicyProperty struct {
	// The name of the application cookie used for stickiness.
	CookieName *string `json:"cookieName"`
	// The mnemonic name for the policy being created.
	//
	// The name must be unique within a set of policies for this load balancer.
	PolicyName *string `json:"policyName"`
}

// Specifies the connection draining settings for your Classic Load Balancer.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_ConnectionDrainingPolicyProperty struct {
	// Specifies whether connection draining is enabled for the load balancer.
	Enabled interface{} `json:"enabled"`
	// The maximum time, in seconds, to keep the existing connections open before deregistering the instances.
	Timeout *float64 `json:"timeout"`
}

// Specifies the idle timeout value for your Classic Load Balancer.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_ConnectionSettingsProperty struct {
	// The time, in seconds, that the connection is allowed to be idle (no data has been sent over the connection) before it is closed by the load balancer.
	IdleTimeout *float64 `json:"idleTimeout"`
}

// Specifies health check settings for your Classic Load Balancer.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_HealthCheckProperty struct {
	// The number of consecutive health checks successes required before moving the instance to the `Healthy` state.
	HealthyThreshold *string `json:"healthyThreshold"`
	// The approximate interval, in seconds, between health checks of an individual instance.
	Interval *string `json:"interval"`
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
	Target *string `json:"target"`
	// The amount of time, in seconds, during which no response means a failed health check.
	//
	// This value must be less than the `Interval` value.
	Timeout *string `json:"timeout"`
	// The number of consecutive health check failures required before moving the instance to the `Unhealthy` state.
	UnhealthyThreshold *string `json:"unhealthyThreshold"`
}

// Specifies a policy for duration-based session stickiness for your Classic Load Balancer.
//
// To associate a policy with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_LBCookieStickinessPolicyProperty struct {
	// The time period, in seconds, after which the cookie should be considered stale.
	//
	// If this parameter is not specified, the stickiness session lasts for the duration of the browser session.
	CookieExpirationPeriod *string `json:"cookieExpirationPeriod"`
	// The name of the policy.
	//
	// This name must be unique within the set of policies for this load balancer.
	PolicyName *string `json:"policyName"`
}

// Specifies a listener for your Classic Load Balancer.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_ListenersProperty struct {
	// The port on which the instance is listening.
	InstancePort *string `json:"instancePort"`
	// The port on which the load balancer is listening.
	//
	// On EC2-VPC, you can specify any port from the range 1-65535. On EC2-Classic, you can specify any port from the following list: 25, 80, 443, 465, 587, 1024-65535.
	LoadBalancerPort *string `json:"loadBalancerPort"`
	// The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.
	Protocol *string `json:"protocol"`
	// The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.
	//
	// If the front-end protocol is TCP or SSL, the back-end protocol must be TCP or SSL. If the front-end protocol is HTTP or HTTPS, the back-end protocol must be HTTP or HTTPS.
	//
	// If there is another listener with the same `InstancePort` whose `InstanceProtocol` is secure, (HTTPS or SSL), the listener's `InstanceProtocol` must also be secure.
	//
	// If there is another listener with the same `InstancePort` whose `InstanceProtocol` is HTTP or TCP, the listener's `InstanceProtocol` must be HTTP or TCP.
	InstanceProtocol *string `json:"instanceProtocol"`
	// The names of the policies to associate with the listener.
	PolicyNames *[]*string `json:"policyNames"`
	// The Amazon Resource Name (ARN) of the server certificate.
	SslCertificateId *string `json:"sslCertificateId"`
}

// Specifies policies for your Classic Load Balancer.
//
// To associate policies with a listener, use the [PolicyNames](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames) property for the listener.
//
// TODO: EXAMPLE
//
type CfnLoadBalancer_PoliciesProperty struct {
	// The policy attributes.
	Attributes interface{} `json:"attributes"`
	// The name of the policy.
	PolicyName *string `json:"policyName"`
	// The name of the policy type.
	PolicyType *string `json:"policyType"`
	// The instance ports for the policy.
	//
	// Required only for some policy types.
	InstancePorts *[]*string `json:"instancePorts"`
	// The load balancer ports for the policy.
	//
	// Required only for some policy types.
	LoadBalancerPorts *[]*string `json:"loadBalancerPorts"`
}

// Properties for defining a `CfnLoadBalancer`.
//
// TODO: EXAMPLE
//
type CfnLoadBalancerProps struct {
	// The listeners for the load balancer. You can specify at most one listener per port.
	//
	// If you update the properties for a listener, AWS CloudFormation deletes the existing listener and creates a new one with the specified properties. While the new listener is being created, clients cannot connect to the load balancer.
	Listeners interface{} `json:"listeners"`
	// Information about where and how access logs are stored for the load balancer.
	AccessLoggingPolicy interface{} `json:"accessLoggingPolicy"`
	// Information about a policy for application-controlled session stickiness.
	AppCookieStickinessPolicy interface{} `json:"appCookieStickinessPolicy"`
	// The Availability Zones for the load balancer. For load balancers in a VPC, specify `Subnets` instead.
	//
	// Update requires replacement if you did not previously specify an Availability Zone or if you are removing all Availability Zones. Otherwise, update requires no interruption.
	AvailabilityZones *[]*string `json:"availabilityZones"`
	// If enabled, the load balancer allows existing requests to complete before the load balancer shifts traffic away from a deregistered or unhealthy instance.
	//
	// For more information, see [Configure Connection Draining](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html) in the *Classic Load Balancers Guide* .
	ConnectionDrainingPolicy interface{} `json:"connectionDrainingPolicy"`
	// If enabled, the load balancer allows the connections to remain idle (no data is sent over the connection) for the specified duration.
	//
	// By default, Elastic Load Balancing maintains a 60-second idle connection timeout for both front-end and back-end connections of your load balancer. For more information, see [Configure Idle Connection Timeout](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html) in the *Classic Load Balancers Guide* .
	ConnectionSettings interface{} `json:"connectionSettings"`
	// If enabled, the load balancer routes the request traffic evenly across all instances regardless of the Availability Zones.
	//
	// For more information, see [Configure Cross-Zone Load Balancing](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html) in the *Classic Load Balancers Guide* .
	CrossZone interface{} `json:"crossZone"`
	// The health check settings to use when evaluating the health of your EC2 instances.
	//
	// Update requires replacement if you did not previously specify health check settings or if you are removing the health check settings. Otherwise, update requires no interruption.
	HealthCheck interface{} `json:"healthCheck"`
	// The IDs of the instances for the load balancer.
	Instances *[]*string `json:"instances"`
	// Information about a policy for duration-based session stickiness.
	LbCookieStickinessPolicy interface{} `json:"lbCookieStickinessPolicy"`
	// The name of the load balancer.
	//
	// This name must be unique within your set of load balancers for the region.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) . If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	LoadBalancerName *string `json:"loadBalancerName"`
	// The policies defined for your Classic Load Balancer.
	//
	// Specify only back-end server policies.
	Policies interface{} `json:"policies"`
	// The type of load balancer. Valid only for load balancers in a VPC.
	//
	// If `Scheme` is `internet-facing` , the load balancer has a public DNS name that resolves to a public IP address.
	//
	// If `Scheme` is `internal` , the load balancer has a public DNS name that resolves to a private IP address.
	Scheme *string `json:"scheme"`
	// The security groups for the load balancer.
	//
	// Valid only for load balancers in a VPC.
	SecurityGroups *[]*string `json:"securityGroups"`
	// The IDs of the subnets for the load balancer. You can specify at most one subnet per Availability Zone.
	//
	// Update requires replacement if you did not previously specify a subnet or if you are removing all subnets. Otherwise, update requires no interruption. To update to a different subnet in the current Availability Zone, you must first update to a subnet in a different Availability Zone, then update to the new subnet in the original Availability Zone.
	Subnets *[]*string `json:"subnets"`
	// The tags associated with a load balancer.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// Describe the health check to a load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type HealthCheck struct {
	// What port number to health check on.
	// Experimental.
	Port *float64 `json:"port"`
	// After how many successful checks is an instance considered healthy.
	// Experimental.
	HealthyThreshold *float64 `json:"healthyThreshold"`
	// Number of seconds between health checks.
	// Experimental.
	Interval awscdk.Duration `json:"interval"`
	// What path to use for HTTP or HTTPS health check (must return 200).
	//
	// For SSL and TCP health checks, accepting connections is enough to be considered
	// healthy.
	// Experimental.
	Path *string `json:"path"`
	// What protocol to use for health checking.
	//
	// The protocol is automatically determined from the port if it's not supplied.
	// Experimental.
	Protocol LoadBalancingProtocol `json:"protocol"`
	// Health check timeout.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// After how many unsuccessful checks is an instance considered unhealthy.
	// Experimental.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold"`
}

// Interface that is going to be implemented by constructs that you can load balance to.
// Experimental.
type ILoadBalancerTarget interface {
	awsec2.IConnectable
	// Attach load-balanced target to a classic ELB.
	// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type ListenerPort interface {
	awsec2.IConnectable
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


// Experimental.
func NewListenerPort(securityGroup awsec2.ISecurityGroup, defaultPort awsec2.Port) ListenerPort {
	_init_.Initialize()

	j := jsiiProxy_ListenerPort{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancing.ListenerPort",
		[]interface{}{securityGroup, defaultPort},
		&j,
	)

	return &j
}

// Experimental.
func NewListenerPort_Override(l ListenerPort, securityGroup awsec2.ISecurityGroup, defaultPort awsec2.Port) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancing.ListenerPort",
		[]interface{}{securityGroup, defaultPort},
		l,
	)
}

// A load balancer with a single listener.
//
// Routes to a fleet of of instances in a VPC.
//
// TODO: EXAMPLE
//
// Experimental.
type LoadBalancer interface {
	awscdk.Resource
	awsec2.IConnectable
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	ListenerPorts() *[]ListenerPort
	LoadBalancerCanonicalHostedZoneName() *string
	LoadBalancerCanonicalHostedZoneNameId() *string
	LoadBalancerDnsName() *string
	LoadBalancerName() *string
	LoadBalancerSourceSecurityGroupGroupName() *string
	LoadBalancerSourceSecurityGroupOwnerAlias() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddListener(listener *LoadBalancerListener) ListenerPort
	AddTarget(target ILoadBalancerTarget)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_LoadBalancer) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewLoadBalancer(scope constructs.Construct, id *string, props *LoadBalancerProps) LoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_LoadBalancer{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancing.LoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLoadBalancer_Override(l LoadBalancer, scope constructs.Construct, id *string, props *LoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancing.LoadBalancer",
		[]interface{}{scope, id, props},
		l,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func LoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancing.LoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func LoadBalancer_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancing.LoadBalancer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a backend to the load balancer.
//
// Returns: A ListenerPort object that controls connections to the listener port
// Experimental.
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

// Experimental.
func (l *jsiiProxy_LoadBalancer) AddTarget(target ILoadBalancerTarget) {
	_jsii_.InvokeVoid(
		l,
		"addTarget",
		[]interface{}{target},
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
// Experimental.
func (l *jsiiProxy_LoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LoadBalancer) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LoadBalancer) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_LoadBalancer) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LoadBalancer) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LoadBalancer) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_LoadBalancer) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Add a backend to the load balancer.
//
// TODO: EXAMPLE
//
// Experimental.
type LoadBalancerListener struct {
	// External listening port.
	// Experimental.
	ExternalPort *float64 `json:"externalPort"`
	// Allow connections to the load balancer from the given set of connection peers.
	//
	// By default, connections will be allowed from anywhere. Set this to an empty list
	// to deny connections, or supply a custom list of peers to allow connections from
	// (IP ranges or security groups).
	// Experimental.
	AllowConnectionsFrom *[]awsec2.IConnectable `json:"allowConnectionsFrom"`
	// What public protocol to use for load balancing.
	//
	// Either 'tcp', 'ssl', 'http' or 'https'.
	//
	// May be omitted if the external port is either 80 or 443.
	// Experimental.
	ExternalProtocol LoadBalancingProtocol `json:"externalProtocol"`
	// Instance listening port.
	//
	// Same as the externalPort if not specified.
	// Experimental.
	InternalPort *float64 `json:"internalPort"`
	// What public protocol to use for load balancing.
	//
	// Either 'tcp', 'ssl', 'http' or 'https'.
	//
	// May be omitted if the internal port is either 80 or 443.
	//
	// The instance protocol is 'tcp' if the front-end protocol
	// is 'tcp' or 'ssl', the instance protocol is 'http' if the
	// front-end protocol is 'https'.
	// Experimental.
	InternalProtocol LoadBalancingProtocol `json:"internalProtocol"`
	// SSL policy names.
	// Experimental.
	PolicyNames *[]*string `json:"policyNames"`
	// the ARN of the SSL certificate.
	// Experimental.
	SslCertificateArn *string `json:"sslCertificateArn"`
	// the ARN of the SSL certificate.
	// Deprecated: - use sslCertificateArn instead
	SslCertificateId *string `json:"sslCertificateId"`
}

// Construction properties for a LoadBalancer.
//
// TODO: EXAMPLE
//
// Experimental.
type LoadBalancerProps struct {
	// VPC network of the fleet instances.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Enable Loadbalancer access logs Can be used to avoid manual work as aws console Required S3 bucket name , enabled flag Can add interval for pushing log Can set bucket prefix in order to provide folder name inside bucket.
	// Experimental.
	AccessLoggingPolicy *CfnLoadBalancer_AccessLoggingPolicyProperty `json:"accessLoggingPolicy"`
	// Whether cross zone load balancing is enabled.
	//
	// This controls whether the load balancer evenly distributes requests
	// across each availability zone
	// Experimental.
	CrossZone *bool `json:"crossZone"`
	// Health check settings for the load balancing targets.
	//
	// Not required but recommended.
	// Experimental.
	HealthCheck *HealthCheck `json:"healthCheck"`
	// Whether this is an internet-facing Load Balancer.
	//
	// This controls whether the LB has a public IP address assigned. It does
	// not open up the Load Balancer's security groups to public internet access.
	// Experimental.
	InternetFacing *bool `json:"internetFacing"`
	// What listeners to set up for the load balancer.
	//
	// Can also be added by .addListener()
	// Experimental.
	Listeners *[]*LoadBalancerListener `json:"listeners"`
	// Which subnets to deploy the load balancer.
	//
	// Can be used to define a specific set of subnets to deploy the load balancer to.
	// Useful multiple public or private subnets are covering the same availability zone.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// What targets to load balance to.
	//
	// Can also be added by .addTarget()
	// Experimental.
	Targets *[]ILoadBalancerTarget `json:"targets"`
}

// Experimental.
type LoadBalancingProtocol string

const (
	LoadBalancingProtocol_TCP LoadBalancingProtocol = "TCP"
	LoadBalancingProtocol_SSL LoadBalancingProtocol = "SSL"
	LoadBalancingProtocol_HTTP LoadBalancingProtocol = "HTTP"
	LoadBalancingProtocol_HTTPS LoadBalancingProtocol = "HTTPS"
)

