package awswafv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awswafv2/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::WAFv2::IPSet`.
//
// > This is the latest version of *AWS WAF* , named AWS WAF V2, released in November, 2019. For information, including how to migrate your AWS WAF resources from the prior release, see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Use an `IPSet` to identify web requests that originate from specific IP addresses or ranges of IP addresses. For example, if you're receiving a lot of requests from a ranges of IP addresses, you can configure AWS WAF to block them using an IP set that lists those IP addresses.
//
// You use an IP set by providing its Amazon Resource Name (ARN) to the rule statement `IPSetReferenceStatement` , when you add a rule to a rule group or web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPSet := awscdk.Aws_wafv2.NewCfnIPSet(this, jsii.String("MyCfnIPSet"), &cfnIPSetProps{
//   	addresses: []*string{
//   		jsii.String("addresses"),
//   	},
//   	ipAddressVersion: jsii.String("ipAddressVersion"),
//   	scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnIPSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation.
	//
	// AWS WAF supports all IPv4 and IPv6 CIDR ranges except for /0.
	//
	// Example address strings:
	//
	// - To configure AWS WAF to allow, block, or count requests that originated from the IP address 192.0.2.44, specify `192.0.2.44/32` .
	// - To configure AWS WAF to allow, block, or count requests that originated from IP addresses from 192.0.2.0 to 192.0.2.255, specify `192.0.2.0/24` .
	// - To configure AWS WAF to allow, block, or count requests that originated from the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128` .
	// - To configure AWS WAF to allow, block, or count requests that originated from IP addresses 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64` .
	//
	// For more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
	//
	// Example JSON `Addresses` specifications:
	//
	// - Empty array: `"Addresses": []`
	// - Array with one address: `"Addresses": ["192.0.2.44/32"]`
	// - Array with three addresses: `"Addresses": ["192.0.2.44/32", "192.0.2.0/24", "192.0.0.0/16"]`
	// - INVALID specification: `"Addresses": [""]` INVALID.
	Addresses() *[]*string
	SetAddresses(val *[]*string)
	// The Amazon Resource Name (ARN) of the IP set.
	AttrArn() *string
	// The ID of the IP set.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A description of the IP set that helps with identification.
	Description() *string
	SetDescription(val *string)
	// The version of the IP addresses, either `IPV4` or `IPV6` .
	IpAddressVersion() *string
	SetIpAddressVersion(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the IP set.
	//
	// You cannot change the name of an `IPSet` after you create it.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope() *string
	SetScope(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnIPSet
type jsiiProxy_CfnIPSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIPSet) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) IpAddressVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::IPSet`.
func NewCfnIPSet(scope awscdk.Construct, id *string, props *CfnIPSetProps) CfnIPSet {
	_init_.Initialize()

	j := jsiiProxy_CfnIPSet{}

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnIPSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::IPSet`.
func NewCfnIPSet_Override(c CfnIPSet, scope awscdk.Construct, id *string, props *CfnIPSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnIPSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIPSet) SetAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"addresses",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetIpAddressVersion(val *string) {
	_jsii_.Set(
		j,
		"ipAddressVersion",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnIPSet) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
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
func CfnIPSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnIPSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIPSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnIPSet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnIPSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnIPSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIPSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_wafv2.CfnIPSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIPSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIPSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIPSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIPSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIPSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIPSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIPSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIPSet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPSet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIPSet) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIPSet) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIPSet) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIPSet) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnIPSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPSet) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnIPSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPSet) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIPSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnIPSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPSetProps := &cfnIPSetProps{
//   	addresses: []*string{
//   		jsii.String("addresses"),
//   	},
//   	ipAddressVersion: jsii.String("ipAddressVersion"),
//   	scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIPSetProps struct {
	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation.
	//
	// AWS WAF supports all IPv4 and IPv6 CIDR ranges except for /0.
	//
	// Example address strings:
	//
	// - To configure AWS WAF to allow, block, or count requests that originated from the IP address 192.0.2.44, specify `192.0.2.44/32` .
	// - To configure AWS WAF to allow, block, or count requests that originated from IP addresses from 192.0.2.0 to 192.0.2.255, specify `192.0.2.0/24` .
	// - To configure AWS WAF to allow, block, or count requests that originated from the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128` .
	// - To configure AWS WAF to allow, block, or count requests that originated from IP addresses 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64` .
	//
	// For more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
	//
	// Example JSON `Addresses` specifications:
	//
	// - Empty array: `"Addresses": []`
	// - Array with one address: `"Addresses": ["192.0.2.44/32"]`
	// - Array with three addresses: `"Addresses": ["192.0.2.44/32", "192.0.2.0/24", "192.0.0.0/16"]`
	// - INVALID specification: `"Addresses": [""]` INVALID.
	Addresses *[]*string `field:"required" json:"addresses" yaml:"addresses"`
	// The version of the IP addresses, either `IPV4` or `IPV6` .
	IpAddressVersion *string `field:"required" json:"ipAddressVersion" yaml:"ipAddressVersion"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// A description of the IP set that helps with identification.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the IP set.
	//
	// You cannot change the name of an `IPSet` after you create it.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::WAFv2::LoggingConfiguration`.
//
// Defines an association between logging destinations and a web ACL resource, for logging from AWS WAF . As part of the association, you can specify parts of the standard logging fields to keep out of the logs and you can specify filters so that you log only a subset of the logging records.
//
// > You can define one logging destination per web ACL.
//
// You can access information about the traffic that AWS WAF inspects using the following steps:
//
// - Create your logging destination. You can use an Amazon CloudWatch Logs log group, an Amazon Simple Storage Service (Amazon S3) bucket, or an Amazon Kinesis Data Firehose. For information about configuring logging destinations and the permissions that are required for each, see [Logging web ACL traffic information](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) in the *AWS WAF Developer Guide* .
// - Associate your logging destination to your web ACL using a `PutLoggingConfiguration` request.
//
// When you successfully enable logging using a `PutLoggingConfiguration` request, AWS WAF creates an additional role or policy that is required to write logs to the logging destination. For an Amazon CloudWatch Logs log group, AWS WAF creates a resource policy on the log group. For an Amazon S3 bucket, AWS WAF creates a bucket policy. For an Amazon Kinesis Data Firehose, AWS WAF creates a service-linked role.
//
// For additional information about web ACL logging, see [Logging web ACL traffic information](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) in the *AWS WAF Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonBody interface{}
//   var loggingFilter interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var uriPath interface{}
//
//   cfnLoggingConfiguration := awscdk.Aws_wafv2.NewCfnLoggingConfiguration(this, jsii.String("MyCfnLoggingConfiguration"), &cfnLoggingConfigurationProps{
//   	logDestinationConfigs: []*string{
//   		jsii.String("logDestinationConfigs"),
//   	},
//   	resourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	loggingFilter: loggingFilter,
//   	redactedFields: []interface{}{
//   		&fieldToMatchProperty{
//   			jsonBody: jsonBody,
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			uriPath: uriPath,
//   		},
//   	},
//   })
//
type CfnLoggingConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Indicates whether the logging configuration was created by AWS Firewall Manager , as part of an AWS WAF policy configuration.
	//
	// If true, only Firewall Manager can modify or delete the configuration.
	AttrManagedByFirewallManager() awscdk.IResolvable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logging destination configuration that you want to associate with the web ACL.
	//
	// > You can associate one logging destination to a web ACL.
	LogDestinationConfigs() *[]*string
	SetLogDestinationConfigs(val *[]*string)
	// Filtering that specifies which web requests are kept in the logs and which are dropped.
	//
	// You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
	LoggingFilter() interface{}
	SetLoggingFilter(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The parts of the request that you want to keep out of the logs.
	//
	// For example, if you redact the `SingleHeader` field, the `HEADER` field in the logs will be `xxx` .
	//
	// > You can specify only the following fields for redaction: `UriPath` , `QueryString` , `SingleHeader` , `Method` , and `JsonBody` .
	RedactedFields() interface{}
	SetRedactedFields(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `LogDestinationConfigs` .
	ResourceArn() *string
	SetResourceArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLoggingConfiguration
type jsiiProxy_CfnLoggingConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLoggingConfiguration) AttrManagedByFirewallManager() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrManagedByFirewallManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) LogDestinationConfigs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logDestinationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) LoggingFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) RedactedFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redactedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::LoggingConfiguration`.
func NewCfnLoggingConfiguration(scope awscdk.Construct, id *string, props *CfnLoggingConfigurationProps) CfnLoggingConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnLoggingConfiguration{}

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnLoggingConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::LoggingConfiguration`.
func NewCfnLoggingConfiguration_Override(c CfnLoggingConfiguration, scope awscdk.Construct, id *string, props *CfnLoggingConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnLoggingConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLoggingConfiguration) SetLogDestinationConfigs(val *[]*string) {
	_jsii_.Set(
		j,
		"logDestinationConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnLoggingConfiguration) SetLoggingFilter(val interface{}) {
	_jsii_.Set(
		j,
		"loggingFilter",
		val,
	)
}

func (j *jsiiProxy_CfnLoggingConfiguration) SetRedactedFields(val interface{}) {
	_jsii_.Set(
		j,
		"redactedFields",
		val,
	)
}

func (j *jsiiProxy_CfnLoggingConfiguration) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
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
func CfnLoggingConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnLoggingConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLoggingConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnLoggingConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnLoggingConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnLoggingConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoggingConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_wafv2.CfnLoggingConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoggingConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggingConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggingConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggingConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggingConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggingConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnLoggingConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggingConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLoggingConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The part of the web request that you want AWS WAF to inspect.
//
// Include the single `FieldToMatch` type that you want to inspect, with additional specifications as needed, according to the type. You specify a single request component in `FieldToMatch` for each rule statement that requires it. To inspect more than one component of the web request, create a separate rule statement for each component.
//
// Example JSON for a `QueryString` field to match:
//
// `"FieldToMatch": { "QueryString": {} }`
//
// Example JSON for a `Method` field to match specification:
//
// `"FieldToMatch": { "Method": { "Name": "DELETE" } }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonBody interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var uriPath interface{}
//
//   fieldToMatchProperty := &fieldToMatchProperty{
//   	jsonBody: jsonBody,
//   	method: method,
//   	queryString: queryString,
//   	singleHeader: singleHeader,
//   	uriPath: uriPath,
//   }
//
type CfnLoggingConfiguration_FieldToMatchProperty struct {
	// Inspect the request body as JSON.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// Only the first 8 KB (8192 bytes) of the request body are forwarded to AWS WAF for inspection by the underlying host service. For information about how to handle oversized request bodies, see the `JsonBody` object configuration.
	JsonBody interface{} `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// Inspect the HTTP method.
	//
	// The method indicates the type of operation that the request is asking the origin to perform.
	Method interface{} `field:"optional" json:"method" yaml:"method"`
	// Inspect the query string.
	//
	// This is the part of a URL that appears after a `?` character, if any.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// Inspect a single header.
	//
	// Provide the name of the header to inspect, for example, `User-Agent` or `Referer` . This setting isn't case sensitive.
	//
	// Example JSON: `"SingleHeader": { "Name": "haystack" }`
	//
	// Alternately, you can filter and inspect all headers with the `Headers` `FieldToMatch` setting.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Inspect the request URI path.
	//
	// This is the part of the web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

// Properties for defining a `CfnLoggingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonBody interface{}
//   var loggingFilter interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var uriPath interface{}
//
//   cfnLoggingConfigurationProps := &cfnLoggingConfigurationProps{
//   	logDestinationConfigs: []*string{
//   		jsii.String("logDestinationConfigs"),
//   	},
//   	resourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	loggingFilter: loggingFilter,
//   	redactedFields: []interface{}{
//   		&fieldToMatchProperty{
//   			jsonBody: jsonBody,
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			uriPath: uriPath,
//   		},
//   	},
//   }
//
type CfnLoggingConfigurationProps struct {
	// The logging destination configuration that you want to associate with the web ACL.
	//
	// > You can associate one logging destination to a web ACL.
	LogDestinationConfigs *[]*string `field:"required" json:"logDestinationConfigs" yaml:"logDestinationConfigs"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `LogDestinationConfigs` .
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Filtering that specifies which web requests are kept in the logs and which are dropped.
	//
	// You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
	LoggingFilter interface{} `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// The parts of the request that you want to keep out of the logs.
	//
	// For example, if you redact the `SingleHeader` field, the `HEADER` field in the logs will be `xxx` .
	//
	// > You can specify only the following fields for redaction: `UriPath` , `QueryString` , `SingleHeader` , `Method` , and `JsonBody` .
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

// A CloudFormation `AWS::WAFv2::RegexPatternSet`.
//
// > This is the latest version of *AWS WAF* , named AWS WAF V2, released in November, 2019. For information, including how to migrate your AWS WAF resources from the prior release, see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Use an `RegexPatternSet` to have AWS WAF inspect a web request component for a specific set of regular expression patterns.
//
// You use a regex pattern set by providing its Amazon Resource Name (ARN) to the rule statement `RegexPatternSetReferenceStatement` , when you add a rule to a rule group or web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegexPatternSet := awscdk.Aws_wafv2.NewCfnRegexPatternSet(this, jsii.String("MyCfnRegexPatternSet"), &cfnRegexPatternSetProps{
//   	regularExpressionList: []*string{
//   		jsii.String("regularExpressionList"),
//   	},
//   	scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnRegexPatternSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the regex pattern set.
	AttrArn() *string
	// The ID of the regex pattern set.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A description of the set that helps with identification.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the set.
	//
	// You cannot change the name after you create the set.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The regular expression patterns in the set.
	RegularExpressionList() *[]*string
	SetRegularExpressionList(val *[]*string)
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope() *string
	SetScope(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRegexPatternSet
type jsiiProxy_CfnRegexPatternSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRegexPatternSet) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) RegularExpressionList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regularExpressionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegexPatternSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::RegexPatternSet`.
func NewCfnRegexPatternSet(scope awscdk.Construct, id *string, props *CfnRegexPatternSetProps) CfnRegexPatternSet {
	_init_.Initialize()

	j := jsiiProxy_CfnRegexPatternSet{}

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnRegexPatternSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::RegexPatternSet`.
func NewCfnRegexPatternSet_Override(c CfnRegexPatternSet, scope awscdk.Construct, id *string, props *CfnRegexPatternSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnRegexPatternSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetRegularExpressionList(val *[]*string) {
	_jsii_.Set(
		j,
		"regularExpressionList",
		val,
	)
}

func (j *jsiiProxy_CfnRegexPatternSet) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
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
func CfnRegexPatternSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnRegexPatternSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRegexPatternSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnRegexPatternSet",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnRegexPatternSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnRegexPatternSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRegexPatternSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_wafv2.CfnRegexPatternSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnRegexPatternSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegexPatternSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnRegexPatternSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegexPatternSetProps := &cfnRegexPatternSetProps{
//   	regularExpressionList: []*string{
//   		jsii.String("regularExpressionList"),
//   	},
//   	scope: jsii.String("scope"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRegexPatternSetProps struct {
	// The regular expression patterns in the set.
	RegularExpressionList *[]*string `field:"required" json:"regularExpressionList" yaml:"regularExpressionList"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// A description of the set that helps with identification.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the set.
	//
	// You cannot change the name after you create the set.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::WAFv2::RuleGroup`.
//
// > This is the latest version of *AWS WAF* , named AWS WAF V2, released in November, 2019. For information, including how to migrate your AWS WAF resources from the prior release, see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Use an `RuleGroup` to define a collection of rules for inspecting and controlling web requests. You use a rule group in an `WebACL` by providing its Amazon Resource Name (ARN) to the rule statement `RuleGroupReferenceStatement` , when you add rules to the web ACL.
//
// When you create a rule group, you define an immutable capacity limit. If you update a rule group, you must stay within the capacity. This allows others to reuse the rule group with confidence in its capacity requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allow interface{}
//   var allQueryArguments interface{}
//   var block interface{}
//   var captcha interface{}
//   var count interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   cfnRuleGroup := awscdk.Aws_wafv2.NewCfnRuleGroup(this, jsii.String("MyCfnRuleGroup"), &cfnRuleGroupProps{
//   	capacity: jsii.Number(123),
//   	scope: jsii.String("scope"),
//   	visibilityConfig: &visibilityConfigProperty{
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   		metricName: jsii.String("metricName"),
//   		sampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	customResponseBodies: map[string]interface{}{
//   		"customResponseBodiesKey": &CustomResponseBodyProperty{
//   			"content": jsii.String("content"),
//   			"contentType": jsii.String("contentType"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			name: jsii.String("name"),
//   			priority: jsii.Number(123),
//   			statement: &statementProperty{
//   				andStatement: &andStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: &orStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   			visibilityConfig: &visibilityConfigProperty{
//   				cloudWatchMetricsEnabled: jsii.Boolean(false),
//   				metricName: jsii.String("metricName"),
//   				sampledRequestsEnabled: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			action: &ruleActionProperty{
//   				allow: allow,
//   				block: block,
//   				captcha: captcha,
//   				count: count,
//   			},
//   			captchaConfig: &captchaConfigProperty{
//   				immunityTimeProperty: &immunityTimePropertyProperty{
//   					immunityTime: jsii.Number(123),
//   				},
//   			},
//   			ruleLabels: []interface{}{
//   				&labelProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnRuleGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the rule group.
	AttrArn() *string
	// Labels that rules in this rule group add to matching requests.
	//
	// These labels are defined in the `RuleLabels` for a `Rule` .
	AttrAvailableLabels() awscdk.IResolvable
	// Labels that rules in this rule group match against.
	//
	// Each of these labels is defined in a `LabelMatchStatement` specification, in the rule statement.
	AttrConsumedLabels() awscdk.IResolvable
	// The ID of the rule group.
	AttrId() *string
	// The label namespace prefix for this rule group.
	//
	// All labels added by rules in this rule group have this prefix.
	//
	// The syntax for the label namespace prefix for a rule group is the following: `awswaf:<account ID>:rule group:<rule group name>:`
	//
	// When a rule with a label matches a web request, AWS WAF adds the fully qualified label to the request. A fully qualified label is made up of the label namespace from the rule group or web ACL where the rule is defined and the label from the rule, separated by a colon.
	AttrLabelNamespace() *string
	// The web ACL capacity units (WCUs) required for this rule group.
	//
	// When you create your own rule group, you define this, and you cannot change it after creation. When you add or modify the rules in a rule group, AWS WAF enforces this limit.
	//
	// AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
	Capacity() *float64
	SetCapacity(val *float64)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A map of custom response keys and content bodies.
	//
	// When you create a rule with a block action, you can send a custom response to the web request. You define these for the rule group, and then use them in the rules that you define in the rule group.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomResponseBodies() interface{}
	SetCustomResponseBodies(val interface{})
	// A description of the rule group that helps with identification.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the rule group.
	//
	// You cannot change the name of a rule group after you create it.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The rule statements used to identify the web requests that you want to allow, block, or count.
	//
	// Each rule includes one top-level statement that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
	Rules() interface{}
	SetRules(val interface{})
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope() *string
	SetScope(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig() interface{}
	SetVisibilityConfig(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRuleGroup
type jsiiProxy_CfnRuleGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRuleGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrAvailableLabels() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAvailableLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrConsumedLabels() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrConsumedLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrLabelNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLabelNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CustomResponseBodies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customResponseBodies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) VisibilityConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::RuleGroup`.
func NewCfnRuleGroup(scope awscdk.Construct, id *string, props *CfnRuleGroupProps) CfnRuleGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnRuleGroup{}

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnRuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::RuleGroup`.
func NewCfnRuleGroup_Override(c CfnRuleGroup, scope awscdk.Construct, id *string, props *CfnRuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnRuleGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetCapacity(val *float64) {
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetCustomResponseBodies(val interface{}) {
	_jsii_.Set(
		j,
		"customResponseBodies",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetVisibilityConfig(val interface{}) {
	_jsii_.Set(
		j,
		"visibilityConfig",
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
func CfnRuleGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnRuleGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRuleGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnRuleGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnRuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnRuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRuleGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_wafv2.CfnRuleGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRuleGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRuleGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRuleGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRuleGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRuleGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnRuleGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRuleGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnRuleGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnRuleGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A logical rule statement used to combine other rule statements with AND logic.
//
// You provide more than one `Statement` within the `AndStatement` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   andStatementProperty := &andStatementProperty{
//   	statements: []interface{}{
//   		&statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			notStatement: &notStatementProperty{
//   				statement: statementProperty_,
//   			},
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: &rateBasedStatementProperty{
//   				aggregateKeyType: jsii.String("aggregateKeyType"),
//   				limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   				scopeDownStatement: statementProperty_,
//   			},
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_AndStatementProperty struct {
	// The statements to combine with AND logic.
	//
	// You can use any statements that can be nested.
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

// Inspect the body of the web request. The body immediately follows the request headers.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodyProperty := &bodyProperty{
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnRuleGroup_BodyProperty struct {
	// What AWS WAF should do if the body is larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of the body of a web request when the body exceeds 8 KB (8192 bytes). Only the first 8 KB of the request body are forwarded to AWS WAF by the underlying host service.
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the body normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// You can combine the `MATCH` or `NO_MATCH` settings for oversize handling with your rule and web ACL action settings, so that you block any request whose body is over 8 KB.
	//
	// Default: `CONTINUE`.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

// A rule statement that defines a string match search for AWS WAF to apply to web requests.
//
// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is refered to as a string match statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   byteMatchStatementProperty := &byteMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	positionalConstraint: jsii.String("positionalConstraint"),
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	searchString: jsii.String("searchString"),
//   	searchStringBase64: jsii.String("searchStringBase64"),
//   }
//
type CfnRuleGroup_ByteMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The area within the portion of the web request that you want AWS WAF to search for `SearchString` .
	//
	// Valid values include the following:
	//
	// *CONTAINS*
	//
	// The specified part of the web request must include the value of `SearchString` , but the location doesn't matter.
	//
	// *CONTAINS_WORD*
	//
	// The specified part of the web request must include the value of `SearchString` , and `SearchString` must contain only alphanumeric characters or underscore (A-Z, a-z, 0-9, or _). In addition, `SearchString` must be a word, which means that both of the following are true:
	//
	// - `SearchString` is at the beginning of the specified part of the web request or is preceded by a character other than an alphanumeric character or underscore (_). Examples include the value of a header and `;BadBot` .
	// - `SearchString` is at the end of the specified part of the web request or is followed by a character other than an alphanumeric character or underscore (_), for example, `BadBot;` and `-BadBot;` .
	//
	// *EXACTLY*
	//
	// The value of the specified part of the web request must exactly match the value of `SearchString` .
	//
	// *STARTS_WITH*
	//
	// The value of `SearchString` must appear at the beginning of the specified part of the web request.
	//
	// *ENDS_WITH*
	//
	// The value of `SearchString` must appear at the end of the specified part of the web request.
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
	// A string value that you want AWS WAF to search for.
	//
	// AWS WAF searches only in the part of web requests that you designate for inspection in `FieldToMatch` . The maximum length of the value is 50 bytes. For alphabetic characters A-Z and a-z, the value is case sensitive.
	//
	// Don't encode this string. Provide the value that you want AWS WAF to search for. AWS CloudFormation automatically base64 encodes the value for you.
	//
	// For example, suppose the value of `Type` is `HEADER` and the value of `Data` is `User-Agent` . If you want to search the `User-Agent` header for the value `BadBot` , you provide the string `BadBot` in the value of `SearchString` .
	//
	// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
	// String to search for in a web request component, base64-encoded.
	//
	// If you don't want to encode the string, specify the unencoded value in `SearchString` instead.
	//
	// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
	SearchStringBase64 *string `field:"optional" json:"searchStringBase64" yaml:"searchStringBase64"`
}

// Specifies how AWS WAF should handle `CAPTCHA` evaluations.
//
// This is available at the web ACL level and in each rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captchaConfigProperty := &captchaConfigProperty{
//   	immunityTimeProperty: &immunityTimePropertyProperty{
//   		immunityTime: jsii.Number(123),
//   	},
//   }
//
type CfnRuleGroup_CaptchaConfigProperty struct {
	// Determines how long a `CAPTCHA` token remains valid after the client successfully solves a `CAPTCHA` puzzle.
	ImmunityTimeProperty interface{} `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

// The filter to use to identify the subset of cookies to inspect in a web request.
//
// You must specify exactly one setting: either `All` , `IncludedCookies` , or `ExcludedCookies` .
//
// Example JSON: `"MatchPattern": { "IncludedCookies": {"KeyToInclude1", "KeyToInclude2", "KeyToInclude3"} }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   cookieMatchPatternProperty := &cookieMatchPatternProperty{
//   	all: all,
//   	excludedCookies: []*string{
//   		jsii.String("excludedCookies"),
//   	},
//   	includedCookies: []*string{
//   		jsii.String("includedCookies"),
//   	},
//   }
//
type CfnRuleGroup_CookieMatchPatternProperty struct {
	// Inspect all cookies.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Inspect only the cookies whose keys don't match any of the strings specified here.
	ExcludedCookies *[]*string `field:"optional" json:"excludedCookies" yaml:"excludedCookies"`
	// Inspect only the cookies that have a key that matches one of the strings specified here.
	IncludedCookies *[]*string `field:"optional" json:"includedCookies" yaml:"includedCookies"`
}

// Inspect the cookies in the web request.
//
// You can specify the parts of the cookies to inspect and you can narrow the set of cookies to inspect by including or excluding specific keys.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example JSON: `"Cookies": { "MatchPattern": { "All": {} }, "MatchScope": "KEY", "OversizeHandling": "MATCH" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   cookiesProperty := &cookiesProperty{
//   	matchPattern: &cookieMatchPatternProperty{
//   		all: all,
//   		excludedCookies: []*string{
//   			jsii.String("excludedCookies"),
//   		},
//   		includedCookies: []*string{
//   			jsii.String("includedCookies"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnRuleGroup_CookiesProperty struct {
	// The filter to use to identify the subset of cookies to inspect in a web request.
	//
	// You must specify exactly one setting: either `All` , `IncludedCookies` , or `ExcludedCookies` .
	//
	// Example JSON: `"MatchPattern": { "IncludedCookies": {"KeyToInclude1", "KeyToInclude2", "KeyToInclude3"} }`.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the cookies to inspect with the rule inspection criteria.
	//
	// If you specify `All` , AWS WAF inspects both keys and values.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if the cookies of the request are larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of request cookies when they exceed 8 KB (8192 bytes) or 200 total cookies. The underlying host service forwards a maximum of 200 cookies and at most 8 KB of cookie contents to AWS WAF .
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the cookies normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

// The response body to use in a custom response to a web request.
//
// This is referenced by key from `CustomResponse` `CustomResponseBodyKey` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResponseBodyProperty := &customResponseBodyProperty{
//   	content: jsii.String("content"),
//   	contentType: jsii.String("contentType"),
//   }
//
type CfnRuleGroup_CustomResponseBodyProperty struct {
	// The payload of the custom response.
	//
	// You can use JSON escape strings in JSON content. To do this, you must specify JSON content in the `ContentType` setting.
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	Content *string `field:"required" json:"content" yaml:"content"`
	// The type of content in the payload that you are defining in the `Content` string.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
}

// The part of the web request that you want AWS WAF to inspect.
//
// Include the single `FieldToMatch` type that you want to inspect, with additional specifications as needed, according to the type. You specify a single request component in `FieldToMatch` for each rule statement that requires it. To inspect more than one component of the web request, create a separate rule statement for each component.
//
// Example JSON for a `QueryString` field to match:
//
// `"FieldToMatch": { "QueryString": {} }`
//
// Example JSON for a `Method` field to match specification:
//
// `"FieldToMatch": { "Method": { "Name": "DELETE" } }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   fieldToMatchProperty := &fieldToMatchProperty{
//   	allQueryArguments: allQueryArguments,
//   	body: &bodyProperty{
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	cookies: &cookiesProperty{
//   		matchPattern: &cookieMatchPatternProperty{
//   			all: all,
//   			excludedCookies: []*string{
//   				jsii.String("excludedCookies"),
//   			},
//   			includedCookies: []*string{
//   				jsii.String("includedCookies"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	headers: &headersProperty{
//   		matchPattern: &headerMatchPatternProperty{
//   			all: all,
//   			excludedHeaders: []*string{
//   				jsii.String("excludedHeaders"),
//   			},
//   			includedHeaders: []*string{
//   				jsii.String("includedHeaders"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	jsonBody: &jsonBodyProperty{
//   		matchPattern: &jsonMatchPatternProperty{
//   			all: all,
//   			includedPaths: []*string{
//   				jsii.String("includedPaths"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//
//   		// the properties below are optional
//   		invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	method: method,
//   	queryString: queryString,
//   	singleHeader: singleHeader,
//   	singleQueryArgument: singleQueryArgument,
//   	uriPath: uriPath,
//   }
//
type CfnRuleGroup_FieldToMatchProperty struct {
	// Inspect all query arguments.
	AllQueryArguments interface{} `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// Inspect the request body as plain text.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// Only the first 8 KB (8192 bytes) of the request body are forwarded to AWS WAF for inspection by the underlying host service. For information about how to handle oversized request bodies, see the `Body` object configuration.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// Inspect the request cookies.
	//
	// You must configure scope and pattern matching filters in the `Cookies` object, to define the set of cookies and the parts of the cookies that AWS WAF inspects.
	//
	// Only the first 8 KB (8192 bytes) of a request's cookies and only the first 200 cookies are forwarded to AWS WAF for inspection by the underlying host service. You must configure how to handle any oversize cookie content in the `Cookies` object. AWS WAF applies the pattern matching filters to the cookies that it receives from the underlying host service.
	Cookies interface{} `field:"optional" json:"cookies" yaml:"cookies"`
	// Inspect the request headers.
	//
	// You must configure scope and pattern matching filters in the `Headers` object, to define the set of headers to and the parts of the headers that AWS WAF inspects.
	//
	// Only the first 8 KB (8192 bytes) of a request's headers and only the first 200 headers are forwarded to AWS WAF for inspection by the underlying host service. You must configure how to handle any oversize header content in the `Headers` object. AWS WAF applies the pattern matching filters to the headers that it receives from the underlying host service.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Inspect the request body as JSON.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// Only the first 8 KB (8192 bytes) of the request body are forwarded to AWS WAF for inspection by the underlying host service. For information about how to handle oversized request bodies, see the `JsonBody` object configuration.
	JsonBody interface{} `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// Inspect the HTTP method.
	//
	// The method indicates the type of operation that the request is asking the origin to perform.
	Method interface{} `field:"optional" json:"method" yaml:"method"`
	// Inspect the query string.
	//
	// This is the part of a URL that appears after a `?` character, if any.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// Inspect a single header.
	//
	// Provide the name of the header to inspect, for example, `User-Agent` or `Referer` . This setting isn't case sensitive.
	//
	// Example JSON: `"SingleHeader": { "Name": "haystack" }`
	//
	// Alternately, you can filter and inspect all headers with the `Headers` `FieldToMatch` setting.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Inspect a single query argument.
	//
	// Provide the name of the query argument to inspect, such as *UserName* or *SalesRegion* . The name can be up to 30 characters long and isn't case sensitive.
	//
	// Example JSON: `"SingleQueryArgument": { "Name": "myArgument" }`.
	SingleQueryArgument interface{} `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// Inspect the request URI path.
	//
	// This is the part of the web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
//
// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
//
// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
//
// This configuration is used for `GeoMatchStatement` and `RateBasedStatement` . For `IPSetReferenceStatement` , use `IPSetForwardedIPConfig` instead.
//
// AWS WAF only evaluates the first IP address found in the specified HTTP header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forwardedIPConfigurationProperty := &forwardedIPConfigurationProperty{
//   	fallbackBehavior: jsii.String("fallbackBehavior"),
//   	headerName: jsii.String("headerName"),
//   }
//
type CfnRuleGroup_ForwardedIPConfigurationProperty struct {
	// The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// You can specify the following fallback behaviors:
	//
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// The name of the HTTP header to use for the IP address.
	//
	// For example, to use the X-Forwarded-For (XFF) header, set this to `X-Forwarded-For` .
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

// A rule statement used to identify web requests based on country of origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoMatchStatementProperty := &geoMatchStatementProperty{
//   	countryCodes: []*string{
//   		jsii.String("countryCodes"),
//   	},
//   	forwardedIpConfig: &forwardedIPConfigurationProperty{
//   		fallbackBehavior: jsii.String("fallbackBehavior"),
//   		headerName: jsii.String("headerName"),
//   	},
//   }
//
type CfnRuleGroup_GeoMatchStatementProperty struct {
	// An array of two-character country codes, for example, `[ "US", "CN" ]` , from the alpha-2 country ISO codes of the ISO 3166 international standard.
	CountryCodes *[]*string `field:"optional" json:"countryCodes" yaml:"countryCodes"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

// The filter to use to identify the subset of headers to inspect in a web request.
//
// You must specify exactly one setting: either `All` , `IncludedHeaders` , or `ExcludedHeaders` .
//
// Example JSON: `"MatchPattern": { "ExcludedHeaders": {"KeyToExclude1", "KeyToExclude2"} }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   headerMatchPatternProperty := &headerMatchPatternProperty{
//   	all: all,
//   	excludedHeaders: []*string{
//   		jsii.String("excludedHeaders"),
//   	},
//   	includedHeaders: []*string{
//   		jsii.String("includedHeaders"),
//   	},
//   }
//
type CfnRuleGroup_HeaderMatchPatternProperty struct {
	// Inspect all headers.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Inspect only the headers whose keys don't match any of the strings specified here.
	ExcludedHeaders *[]*string `field:"optional" json:"excludedHeaders" yaml:"excludedHeaders"`
	// Inspect only the headers that have a key that matches one of the strings specified here.
	IncludedHeaders *[]*string `field:"optional" json:"includedHeaders" yaml:"includedHeaders"`
}

// Inspect all headers in the web request.
//
// You can specify the parts of the headers to inspect and you can narrow the set of headers to inspect by including or excluding specific keys.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// If you want to inspect just the value of a single header, use the `SingleHeader` `FieldToMatch` setting instead.
//
// Example JSON: `"Headers": { "MatchPattern": { "All": {} }, "MatchScope": "KEY", "OversizeHandling": "MATCH" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   headersProperty := &headersProperty{
//   	matchPattern: &headerMatchPatternProperty{
//   		all: all,
//   		excludedHeaders: []*string{
//   			jsii.String("excludedHeaders"),
//   		},
//   		includedHeaders: []*string{
//   			jsii.String("includedHeaders"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnRuleGroup_HeadersProperty struct {
	// The filter to use to identify the subset of headers to inspect in a web request.
	//
	// You must specify exactly one setting: either `All` , `IncludedHeaders` , or `ExcludedHeaders` .
	//
	// Example JSON: `"MatchPattern": { "ExcludedHeaders": {"KeyToExclude1", "KeyToExclude2"} }`.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the headers to match with the rule inspection criteria.
	//
	// If you specify `All` , AWS WAF inspects both keys and values.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if the headers of the request are larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of request headers when they exceed 8 KB (8192 bytes) or 200 total headers. The underlying host service forwards a maximum of 200 headers and at most 8 KB of header contents to AWS WAF .
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the headers normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
//
// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
//
// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
//
// This configuration is used only for `IPSetReferenceStatement` . For `GeoMatchStatement` and `RateBasedStatement` , use `ForwardedIPConfig` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetForwardedIPConfigurationProperty := map[string]*string{
//   	"fallbackBehavior": jsii.String("fallbackBehavior"),
//   	"headerName": jsii.String("headerName"),
//   	"position": jsii.String("position"),
//   }
//
type CfnRuleGroup_IPSetForwardedIPConfigurationProperty struct {
	// The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// You can specify the following fallback behaviors:
	//
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// The name of the HTTP header to use for the IP address.
	//
	// For example, to use the X-Forwarded-For (XFF) header, set this to `X-Forwarded-For` .
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The position in the header to search for the IP address.
	//
	// The header can contain IP addresses of the original client and also of proxies. For example, the header value could be `10.1.1.1, 127.0.0.0, 10.10.10.10` where the first IP address identifies the original client and the rest identify proxies that the request went through.
	//
	// The options for this setting are the following:
	//
	// - FIRST - Inspect the first IP address in the list of IP addresses in the header. This is usually the client's original IP.
	// - LAST - Inspect the last IP address in the list of IP addresses in the header.
	// - ANY - Inspect all IP addresses in the header for a match. If the header contains more than 10 IP addresses, AWS WAF inspects the last 10.
	Position *string `field:"required" json:"position" yaml:"position"`
}

// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
//
// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
//
// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetReferenceStatementProperty := map[string]interface{}{
//   	"arn": jsii.String("arn"),
//
//   	// the properties below are optional
//   	"ipSetForwardedIpConfig": map[string]*string{
//   		"fallbackBehavior": jsii.String("fallbackBehavior"),
//   		"headerName": jsii.String("headerName"),
//   		"position": jsii.String("position"),
//   	},
//   }
//
type CfnRuleGroup_IPSetReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the `IPSet` that this statement references.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	IpSetForwardedIpConfig interface{} `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

// Determines how long a `CAPTCHA` token remains valid after the client successfully solves a `CAPTCHA` puzzle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   immunityTimePropertyProperty := &immunityTimePropertyProperty{
//   	immunityTime: jsii.Number(123),
//   }
//
type CfnRuleGroup_ImmunityTimePropertyProperty struct {
	// The amount of time, in seconds, that a `CAPTCHA` token is valid.
	//
	// The default setting is 300.
	ImmunityTime *float64 `field:"required" json:"immunityTime" yaml:"immunityTime"`
}

// Inspect the body of the web request as JSON. The body immediately follows the request headers.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Use the specifications in this object to indicate which parts of the JSON body to inspect using the rule's inspection criteria. AWS WAF inspects only the parts of the JSON that result from the matches that you indicate.
//
// Example JSON: `"JsonBody": { "MatchPattern": { "All": {} }, "MatchScope": "ALL" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   jsonBodyProperty := &jsonBodyProperty{
//   	matchPattern: &jsonMatchPatternProperty{
//   		all: all,
//   		includedPaths: []*string{
//   			jsii.String("includedPaths"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//
//   	// the properties below are optional
//   	invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnRuleGroup_JsonBodyProperty struct {
	// The patterns to look for in the JSON body.
	//
	// AWS WAF inspects the results of these pattern matches against the rule inspection criteria.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the JSON to match against using the `MatchPattern` .
	//
	// If you specify `All` , AWS WAF matches against keys and values.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if it fails to completely parse the JSON body. The options are the following:.
	//
	// - `EVALUATE_AS_STRING` - Inspect the body as plain text. AWS WAF applies the text transformations and inspection criteria that you defined for the JSON inspection to the body text string.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// If you don't provide this setting, AWS WAF parses and evaluates the content only up to the first parsing failure that it encounters.
	//
	// AWS WAF does its best to parse the entire JSON body, but might be forced to stop for reasons such as invalid characters, duplicate keys, truncation, and any content whose root node isn't an object or an array.
	//
	// AWS WAF parses the JSON in the following examples as two valid key, value pairs:
	//
	// - Missing comma: `{"key1":"value1""key2":"value2"}`
	// - Missing colon: `{"key1":"value1","key2""value2"}`
	// - Extra colons: `{"key1"::"value1","key2""value2"}`.
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
	// What AWS WAF should do if the body is larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of the body of a web request when the body exceeds 8 KB (8192 bytes). Only the first 8 KB of the request body are forwarded to AWS WAF by the underlying host service.
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the body normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// You can combine the `MATCH` or `NO_MATCH` settings for oversize handling with your rule and web ACL action settings, so that you block any request whose body is over 8 KB.
	//
	// Default: `CONTINUE`.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

// The patterns to look for in the JSON body.
//
// AWS WAF inspects the results of these pattern matches against the rule inspection criteria. This is used with the `FieldToMatch` option `JsonBody` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   jsonMatchPatternProperty := &jsonMatchPatternProperty{
//   	all: all,
//   	includedPaths: []*string{
//   		jsii.String("includedPaths"),
//   	},
//   }
//
type CfnRuleGroup_JsonMatchPatternProperty struct {
	// Match all of the elements. See also `MatchScope` in the `JsonBody` `FieldToMatch` specification.
	//
	// You must specify either this setting or the `IncludedPaths` setting, but not both.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Match only the specified include paths. See also `MatchScope` in the `JsonBody` `FieldToMatch` specification.
	//
	// Provide the include paths using JSON Pointer syntax. For example, `"IncludedPaths": ["/dogs/0/name", "/dogs/1/name"]` . For information about this syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// You must specify either this setting or the `All` setting, but not both.
	//
	// > Don't use this option to include all paths. Instead, use the `All` setting.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

// A rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL.
//
// The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelMatchStatementProperty := &labelMatchStatementProperty{
//   	key: jsii.String("key"),
//   	scope: jsii.String("scope"),
//   }
//
type CfnRuleGroup_LabelMatchStatementProperty struct {
	// The string to match against. The setting you provide for this depends on the match statement's `Scope` setting:.
	//
	// - If the `Scope` indicates `LABEL` , then this specification must include the name and can include any number of preceding namespace specifications and prefix up to providing the fully qualified label name.
	// - If the `Scope` indicates `NAMESPACE` , then this specification can include any number of contiguous namespace strings, and can include the entire label namespace prefix from the rule group or web ACL where the label originates.
	//
	// Labels are case sensitive and components of a label must be separated by colon, for example `NS1:NS2:name` .
	Key *string `field:"required" json:"key" yaml:"key"`
	// Specify whether you want to match using the label name or just the namespace.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

// A single label container.
//
// This is used as an element of a label array in `RuleLabels` inside a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelProperty := &labelProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnRuleGroup_LabelProperty struct {
	// The label string.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// List of labels used by one or more of the rules of a `RuleGroup` .
//
// This summary object is used for the following rule group lists:
//
// - `AvailableLabels` - Labels that rules add to matching requests. These labels are defined in the `RuleLabels` for a rule.
// - `ConsumedLabels` - Labels that rules match against. These labels are defined in a `LabelMatchStatement` specification, in the `Statement` definition of a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelSummaryProperty := &labelSummaryProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnRuleGroup_LabelSummaryProperty struct {
	// An individual label specification.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// A logical rule statement used to negate the results of another rule statement.
//
// You provide one `Statement` within the `NotStatement` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   notStatementProperty := &notStatementProperty{
//   	statement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_NotStatementProperty struct {
	// The statement to negate.
	//
	// You can use any statement that can be nested.
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
}

// A logical rule statement used to combine other rule statements with OR logic.
//
// You provide more than one `Statement` within the `OrStatement` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   orStatementProperty := &orStatementProperty{
//   	statements: []interface{}{
//   		&statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			notStatement: &notStatementProperty{
//   				statement: statementProperty_,
//   			},
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: &rateBasedStatementProperty{
//   				aggregateKeyType: jsii.String("aggregateKeyType"),
//   				limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   				scopeDownStatement: statementProperty_,
//   			},
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_OrStatementProperty struct {
	// The statements to combine with OR logic.
	//
	// You can use any statements that can be nested.
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

// A rate-based rule tracks the rate of requests for each originating IP address, and triggers the rule action when the rate exceeds a limit that you specify on the number of requests in any 5-minute time span.
//
// You can use this to put a temporary block on requests from an IP address that is sending excessive requests.
//
// AWS WAF tracks and manages web requests separately for each instance of a rate-based rule that you use. For example, if you provide the same rate-based rule settings in two web ACLs, each of the two rule statements represents a separate instance of the rate-based rule and gets its own tracking and management by AWS WAF . If you define a rate-based rule inside a rule group, and then use that rule group in multiple places, each use creates a separate instance of the rate-based rule that gets its own tracking and management by AWS WAF .
//
// When the rule action triggers, AWS WAF blocks additional requests from the IP address until the request rate falls below the limit.
//
// You can optionally nest another statement inside the rate-based statement, to narrow the scope of the rule so that it only counts requests that match the nested statement. For example, based on recent requests that you have seen from an attacker, you might create a rate-based rule with a nested AND rule statement that contains the following nested statements:
//
// - An IP match statement with an IP set that specified the address 192.0.2.44.
// - A string match statement that searches in the User-Agent header for the string BadBot.
//
// In this rate-based rule, you also define a rate limit. For this example, the rate limit is 1,000. Requests that meet both of the conditions in the statements are counted. If the count exceeds 1,000 requests per five minutes, the rule action triggers. Requests that do not meet both conditions are not counted towards the rate limit and are not affected by this rule.
//
// You cannot nest a `RateBasedStatement` inside another statement, for example inside a `NotStatement` or `OrStatement` . You can define a `RateBasedStatement` inside a web ACL and inside a rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   rateBasedStatementProperty := &rateBasedStatementProperty{
//   	aggregateKeyType: jsii.String("aggregateKeyType"),
//   	limit: jsii.Number(123),
//
//   	// the properties below are optional
//   	forwardedIpConfig: &forwardedIPConfigurationProperty{
//   		fallbackBehavior: jsii.String("fallbackBehavior"),
//   		headerName: jsii.String("headerName"),
//   	},
//   	scopeDownStatement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_RateBasedStatementProperty struct {
	// Setting that indicates how to aggregate the request counts. The options are the following:.
	//
	// - IP - Aggregate the request counts on the IP address from the web request origin.
	// - FORWARDED_IP - Aggregate the request counts on the first IP address in an HTTP header. If you use this, configure the `ForwardedIPConfig` , to specify the header to use.
	AggregateKeyType *string `field:"required" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// The limit on requests per 5-minute period for a single originating IP address.
	//
	// If the statement includes a `ScopeDownStatement` , this limit is applied only to the requests that match the statement.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// This is required if `AggregateKeyType` is set to `FORWARDED_IP` .
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the rate-based statement.
	//
	// Requests are only tracked by the rate-based statement if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

// A rule statement used to search web request components for a match against a single regular expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   regexMatchStatementProperty := &regexMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	regexString: jsii.String("regexString"),
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_RegexMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The string representing the regular expression.
	RegexString *string `field:"required" json:"regexString" yaml:"regexString"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// A rule statement used to search web request components for matches with regular expressions.
//
// To use this, create a `RegexPatternSet` that specifies the expressions that you want to detect, then use the ARN of that set in this statement. A web request matches the pattern set rule statement if the request component matches any of the patterns in the set.
//
// Each regex pattern set rule statement references a regex pattern set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   regexPatternSetReferenceStatementProperty := &regexPatternSetReferenceStatementProperty{
//   	arn: jsii.String("arn"),
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_RegexPatternSetReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the `RegexPatternSet` that this statement references.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// The action that AWS WAF should take on a web request when it matches a rule's statement.
//
// Settings at the web ACL level can override the rule action setting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allow interface{}
//   var block interface{}
//   var captcha interface{}
//   var count interface{}
//
//   ruleActionProperty := &ruleActionProperty{
//   	allow: allow,
//   	block: block,
//   	captcha: captcha,
//   	count: count,
//   }
//
type CfnRuleGroup_RuleActionProperty struct {
	// Instructs AWS WAF to allow the web request.
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Instructs AWS WAF to block the web request.
	Block interface{} `field:"optional" json:"block" yaml:"block"`
	// Specifies that AWS WAF should run a `CAPTCHA` check against the request:.
	//
	// - If the request includes a valid, unexpired `CAPTCHA` token, AWS WAF allows the web request inspection to proceed to the next rule, similar to a `CountAction` .
	// - If the request doesn't include a valid, unexpired `CAPTCHA` token, AWS WAF discontinues the web ACL evaluation of the request and blocks it from going to its intended destination.
	//
	// AWS WAF generates a response that it sends back to the client, which includes the following:
	//
	// - The header `x-amzn-waf-action` with a value of `captcha` .
	// - The HTTP status code `405 Method Not Allowed` .
	// - If the request contains an `Accept` header with a value of `text/html` , the response includes a `CAPTCHA` challenge.
	//
	// You can configure the expiration time in the `CaptchaConfig` `ImmunityTimeProperty` setting at the rule and web ACL level. The rule setting overrides the web ACL setting.
	//
	// This action option is available for rules. It isn't available for web ACL default actions.
	Captcha interface{} `field:"optional" json:"captcha" yaml:"captcha"`
	// Instructs AWS WAF to count the web request and allow it.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}

// A single rule, which you can use in a `WebACL` or `RuleGroup` to identify web requests that you want to allow, block, or count.
//
// Each rule includes one top-level `Statement` that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allow interface{}
//   var allQueryArguments interface{}
//   var block interface{}
//   var captcha interface{}
//   var count interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   ruleProperty := &ruleProperty{
//   	name: jsii.String("name"),
//   	priority: jsii.Number(123),
//   	statement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	visibilityConfig: &visibilityConfigProperty{
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   		metricName: jsii.String("metricName"),
//   		sampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	action: &ruleActionProperty{
//   		allow: allow,
//   		block: block,
//   		captcha: captcha,
//   		count: count,
//   	},
//   	captchaConfig: &captchaConfigProperty{
//   		immunityTimeProperty: &immunityTimePropertyProperty{
//   			immunityTime: jsii.Number(123),
//   		},
//   	},
//   	ruleLabels: []interface{}{
//   		&labelProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_RuleProperty struct {
	// The name of the rule.
	//
	// You can't change the name of a `Rule` after you create it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// If you define more than one `Rule` in a `WebACL` , AWS WAF evaluates each request against the `Rules` in order based on the value of `Priority` .
	//
	// AWS WAF processes rules with lower priority first. The priorities don't need to be consecutive, but they must all be different.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The AWS WAF processing statement for the rule, for example `ByteMatchStatement` or `SizeConstraintStatement` .
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig interface{} `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// The action that AWS WAF should take on a web request when it matches the rule statement.
	//
	// Settings at the web ACL level can override the rule action setting.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations.
	//
	// If you don't specify this, AWS WAF uses the `CAPTCHA` configuration that's defined for the web ACL.
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// Labels to apply to web requests that match the rule match statement.
	//
	// AWS WAF applies fully qualified labels to matching web requests. A fully qualified label is the concatenation of a label namespace and a rule label. The rule's rule group or web ACL defines the label namespace.
	//
	// Rules that run after this rule in the web ACL can match against these labels using a `LabelMatchStatement` .
	//
	// For each label, provide a case-sensitive string containing optional namespaces and a label name, according to the following guidelines:
	//
	// - Separate each component of the label with a colon.
	// - Each namespace or name can have up to 128 characters.
	// - You can specify up to 5 namespaces in a label.
	// - Don't use the following reserved words in your label specification: `aws` , `waf` , `managed` , `rulegroup` , `webacl` , `regexpatternset` , or `ipset` .
	//
	// For example, `myLabelName` or `nameSpace1:nameSpace2:myLabelName` .
	RuleLabels interface{} `field:"optional" json:"ruleLabels" yaml:"ruleLabels"`
}

// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
//
// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
//
// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the first 8192 bytes (8 KB). If the request body for your web requests never exceeds 8192 bytes, you can create a size constraint condition and block requests that have a request body greater than 8192 bytes.
//
// If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   sizeConstraintStatementProperty := &sizeConstraintStatementProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	size: jsii.Number(123),
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_SizeConstraintStatementProperty struct {
	// The operator to use to compare the request part to the size setting.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The size, in byte, to compare to the request part, after any transformations.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// Attackers sometimes insert malicious SQL code into web requests in an effort to extract data from your database.
//
// To allow or block web requests that appear to contain malicious SQL code, create one or more SQL injection match conditions. An SQL injection match condition identifies the part of web requests, such as the URI or the query string, that you want AWS WAF to inspect. Later in the process, when you create a web ACL, you specify whether to allow or block requests that appear to contain malicious SQL code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   sqliMatchStatementProperty := &sqliMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_SqliMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// The processing guidance for a rule, used by AWS WAF to determine whether a web request matches the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var andStatementProperty_ andStatementProperty
//   var method interface{}
//   var notStatementProperty_ notStatementProperty
//   var orStatementProperty_ orStatementProperty
//   var queryString interface{}
//   var rateBasedStatementProperty_ rateBasedStatementProperty
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   statementProperty := &statementProperty{
//   	andStatement: &andStatementProperty{
//   		statements: []interface{}{
//   			&statementProperty{
//   				andStatement: andStatementProperty_,
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: &orStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	byteMatchStatement: &byteMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		positionalConstraint: jsii.String("positionalConstraint"),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		searchString: jsii.String("searchString"),
//   		searchStringBase64: jsii.String("searchStringBase64"),
//   	},
//   	geoMatchStatement: &geoMatchStatementProperty{
//   		countryCodes: []*string{
//   			jsii.String("countryCodes"),
//   		},
//   		forwardedIpConfig: &forwardedIPConfigurationProperty{
//   			fallbackBehavior: jsii.String("fallbackBehavior"),
//   			headerName: jsii.String("headerName"),
//   		},
//   	},
//   	ipSetReferenceStatement: map[string]interface{}{
//   		"arn": jsii.String("arn"),
//
//   		// the properties below are optional
//   		"ipSetForwardedIpConfig": map[string]*string{
//   			"fallbackBehavior": jsii.String("fallbackBehavior"),
//   			"headerName": jsii.String("headerName"),
//   			"position": jsii.String("position"),
//   		},
//   	},
//   	labelMatchStatement: &labelMatchStatementProperty{
//   		key: jsii.String("key"),
//   		scope: jsii.String("scope"),
//   	},
//   	notStatement: &notStatementProperty{
//   		statement: &statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			notStatement: notStatementProperty_,
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: &rateBasedStatementProperty{
//   				aggregateKeyType: jsii.String("aggregateKeyType"),
//   				limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   				scopeDownStatement: statementProperty_,
//   			},
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	orStatement: &orStatementProperty{
//   		statements: []interface{}{
//   			&statementProperty{
//   				andStatement: &andStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: orStatementProperty_,
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	rateBasedStatement: &rateBasedStatementProperty{
//   		aggregateKeyType: jsii.String("aggregateKeyType"),
//   		limit: jsii.Number(123),
//
//   		// the properties below are optional
//   		forwardedIpConfig: &forwardedIPConfigurationProperty{
//   			fallbackBehavior: jsii.String("fallbackBehavior"),
//   			headerName: jsii.String("headerName"),
//   		},
//   		scopeDownStatement: &statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			notStatement: &notStatementProperty{
//   				statement: statementProperty_,
//   			},
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: rateBasedStatementProperty_,
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	regexMatchStatement: &regexMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		regexString: jsii.String("regexString"),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   		arn: jsii.String("arn"),
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	sizeConstraintStatement: &sizeConstraintStatementProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		size: jsii.Number(123),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	sqliMatchStatement: &sqliMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	xssMatchStatement: &xssMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_StatementProperty struct {
	// A logical rule statement used to combine other rule statements with AND logic.
	//
	// You provide more than one `Statement` within the `AndStatement` .
	AndStatement interface{} `field:"optional" json:"andStatement" yaml:"andStatement"`
	// A rule statement that defines a string match search for AWS WAF to apply to web requests.
	//
	// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is refered to as a string match statement.
	ByteMatchStatement interface{} `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// A rule statement used to identify web requests based on country of origin.
	GeoMatchStatement interface{} `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
	//
	// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
	//
	// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	IpSetReferenceStatement interface{} `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// A rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL.
	//
	// The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.
	LabelMatchStatement interface{} `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// A logical rule statement used to negate the results of another rule statement.
	//
	// You provide one `Statement` within the `NotStatement` .
	NotStatement interface{} `field:"optional" json:"notStatement" yaml:"notStatement"`
	// A logical rule statement used to combine other rule statements with OR logic.
	//
	// You provide more than one `Statement` within the `OrStatement` .
	OrStatement interface{} `field:"optional" json:"orStatement" yaml:"orStatement"`
	// A rate-based rule tracks the rate of requests for each originating IP address, and triggers the rule action when the rate exceeds a limit that you specify on the number of requests in any 5-minute time span.
	//
	// You can use this to put a temporary block on requests from an IP address that is sending excessive requests.
	//
	// AWS WAF tracks and manages web requests separately for each instance of a rate-based rule that you use. For example, if you provide the same rate-based rule settings in two web ACLs, each of the two rule statements represents a separate instance of the rate-based rule and gets its own tracking and management by AWS WAF . If you define a rate-based rule inside a rule group, and then use that rule group in multiple places, each use creates a separate instance of the rate-based rule that gets its own tracking and management by AWS WAF .
	//
	// When the rule action triggers, AWS WAF blocks additional requests from the IP address until the request rate falls below the limit.
	//
	// You can optionally nest another statement inside the rate-based statement, to narrow the scope of the rule so that it only counts requests that match the nested statement. For example, based on recent requests that you have seen from an attacker, you might create a rate-based rule with a nested AND rule statement that contains the following nested statements:
	//
	// - An IP match statement with an IP set that specified the address 192.0.2.44.
	// - A string match statement that searches in the User-Agent header for the string BadBot.
	//
	// In this rate-based rule, you also define a rate limit. For this example, the rate limit is 1,000. Requests that meet both of the conditions in the statements are counted. If the count exceeds 1,000 requests per five minutes, the rule action triggers. Requests that do not meet both conditions are not counted towards the rate limit and are not affected by this rule.
	//
	// You cannot nest a `RateBasedStatement` inside another statement, for example inside a `NotStatement` or `OrStatement` . You can define a `RateBasedStatement` inside a web ACL and inside a rule group.
	RateBasedStatement interface{} `field:"optional" json:"rateBasedStatement" yaml:"rateBasedStatement"`
	// A rule statement used to search web request components for a match against a single regular expression.
	RegexMatchStatement interface{} `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// A rule statement used to search web request components for matches with regular expressions.
	//
	// To use this, create a `RegexPatternSet` that specifies the expressions that you want to detect, then use the ARN of that set in this statement. A web request matches the pattern set rule statement if the request component matches any of the patterns in the set.
	//
	// Each regex pattern set rule statement references a regex pattern set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	RegexPatternSetReferenceStatement interface{} `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
	//
	// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
	//
	// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the first 8192 bytes (8 KB). If the request body for your web requests never exceeds 8192 bytes, you can create a size constraint condition and block requests that have a request body greater than 8192 bytes.
	//
	// If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.
	SizeConstraintStatement interface{} `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// Attackers sometimes insert malicious SQL code into web requests in an effort to extract data from your database.
	//
	// To allow or block web requests that appear to contain malicious SQL code, create one or more SQL injection match conditions. An SQL injection match condition identifies the part of web requests, such as the URI or the query string, that you want AWS WAF to inspect. Later in the process, when you create a web ACL, you specify whether to allow or block requests that appear to contain malicious SQL code.
	SqliMatchStatement interface{} `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// A rule statement that defines a cross-site scripting (XSS) match search for AWS WAF to apply to web requests.
	//
	// XSS attacks are those where the attacker uses vulnerabilities in a benign website as a vehicle to inject malicious client-site scripts into other legitimate web browsers. The XSS match statement provides the location in requests that you want AWS WAF to search and text transformations to use on the search area before AWS WAF searches for character sequences that are likely to be malicious strings.
	XssMatchStatement interface{} `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textTransformationProperty := &textTransformationProperty{
//   	priority: jsii.Number(123),
//   	type: jsii.String("type"),
//   }
//
type CfnRuleGroup_TextTransformationProperty struct {
	// Sets the relative processing order for multiple transformations that are defined for a rule statement.
	//
	// AWS WAF processes all transformations, from lowest priority to highest, before inspecting the transformed content. The priorities don't need to be consecutive, but they must all be different.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// You can specify the following transformation types:.
	//
	// *BASE64_DECODE* - Decode a `Base64` -encoded string.
	//
	// *BASE64_DECODE_EXT* - Decode a `Base64` -encoded string, but use a forgiving implementation that ignores characters that aren't valid.
	//
	// *CMD_LINE* - Command-line transformations. These are helpful in reducing effectiveness of attackers who inject an operating system command-line command and use unusual formatting to disguise some or all of the command.
	//
	// - Delete the following characters: `\ " ' ^`
	// - Delete spaces before the following characters: `/ (`
	// - Replace the following characters with a space: `, ;`
	// - Replace multiple spaces with one space
	// - Convert uppercase letters (A-Z) to lowercase (a-z)
	//
	// *COMPRESS_WHITE_SPACE* - Replace these characters with a space character (decimal 32):
	//
	// - `\f` , formfeed, decimal 12
	// - `\t` , tab, decimal 9
	// - `\n` , newline, decimal 10
	// - `\r` , carriage return, decimal 13
	// - `\v` , vertical tab, decimal 11
	// - Non-breaking space, decimal 160
	//
	// `COMPRESS_WHITE_SPACE` also replaces multiple spaces with one space.
	//
	// *CSS_DECODE* - Decode characters that were encoded using CSS 2.x escape rules `syndata.html#characters` . This function uses up to two bytes in the decoding process, so it can help to uncover ASCII characters that were encoded using CSS encoding that wouldn’t typically be encoded. It's also useful in countering evasion, which is a combination of a backslash and non-hexadecimal characters. For example, `ja\vascript` for javascript.
	//
	// *ESCAPE_SEQ_DECODE* - Decode the following ANSI C escape sequences: `\a` , `\b` , `\f` , `\n` , `\r` , `\t` , `\v` , `\\` , `\?` , `\'` , `\"` , `\xHH` (hexadecimal), `\0OOO` (octal). Encodings that aren't valid remain in the output.
	//
	// *HEX_DECODE* - Decode a string of hexadecimal characters into a binary.
	//
	// *HTML_ENTITY_DECODE* - Replace HTML-encoded characters with unencoded characters. `HTML_ENTITY_DECODE` performs these operations:
	//
	// - Replaces `(ampersand)quot;` with `"`
	// - Replaces `(ampersand)nbsp;` with a non-breaking space, decimal 160
	// - Replaces `(ampersand)lt;` with a "less than" symbol
	// - Replaces `(ampersand)gt;` with `>`
	// - Replaces characters that are represented in hexadecimal format, `(ampersand)#xhhhh;` , with the corresponding characters
	// - Replaces characters that are represented in decimal format, `(ampersand)#nnnn;` , with the corresponding characters
	//
	// *JS_DECODE* - Decode JavaScript escape sequences. If a `\` `u` `HHHH` code is in the full-width ASCII code range of `FF01-FF5E` , then the higher byte is used to detect and adjust the lower byte. If not, only the lower byte is used and the higher byte is zeroed, causing a possible loss of information.
	//
	// *LOWERCASE* - Convert uppercase letters (A-Z) to lowercase (a-z).
	//
	// *MD5* - Calculate an MD5 hash from the data in the input. The computed hash is in a raw binary form.
	//
	// *NONE* - Specify `NONE` if you don't want any text transformations.
	//
	// *NORMALIZE_PATH* - Remove multiple slashes, directory self-references, and directory back-references that are not at the beginning of the input from an input string.
	//
	// *NORMALIZE_PATH_WIN* - This is the same as `NORMALIZE_PATH` , but first converts backslash characters to forward slashes.
	//
	// *REMOVE_NULLS* - Remove all `NULL` bytes from the input.
	//
	// *REPLACE_COMMENTS* - Replace each occurrence of a C-style comment ( `/* ... * /` ) with a single space. Multiple consecutive occurrences are not compressed. Unterminated comments are also replaced with a space (ASCII 0x20). However, a standalone termination of a comment ( `* /` ) is not acted upon.
	//
	// *REPLACE_NULLS* - Replace NULL bytes in the input with space characters (ASCII `0x20` ).
	//
	// *SQL_HEX_DECODE* - Decode SQL hex data. Example ( `0x414243` ) will be decoded to ( `ABC` ).
	//
	// *URL_DECODE* - Decode a URL-encoded value.
	//
	// *URL_DECODE_UNI* - Like `URL_DECODE` , but with support for Microsoft-specific `%u` encoding. If the code is in the full-width ASCII code range of `FF01-FF5E` , the higher byte is used to detect and adjust the lower byte. Otherwise, only the lower byte is used and the higher byte is zeroed.
	//
	// *UTF8_TO_UNICODE* - Convert all UTF-8 character sequences to Unicode. This helps input normalization, and minimizing false-positives and false-negatives for non-English languages.
	Type *string `field:"required" json:"type" yaml:"type"`
}

// Defines and enables Amazon CloudWatch metrics and web request sample collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visibilityConfigProperty := &visibilityConfigProperty{
//   	cloudWatchMetricsEnabled: jsii.Boolean(false),
//   	metricName: jsii.String("metricName"),
//   	sampledRequestsEnabled: jsii.Boolean(false),
//   }
//
type CfnRuleGroup_VisibilityConfigProperty struct {
	// A boolean indicating whether the associated resource sends metrics to Amazon CloudWatch.
	//
	// For the list of available metrics, see [AWS WAF Metrics](https://docs.aws.amazon.com/waf/latest/developerguide/monitoring-cloudwatch.html#waf-metrics) .
	CloudWatchMetricsEnabled interface{} `field:"required" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// A name of the Amazon CloudWatch metric.
	//
	// The name can contain only the characters: A-Z, a-z, 0-9, - (hyphen), and _ (underscore). The name can be from one to 128 characters long. It can't contain whitespace or metric names reserved for AWS WAF , for example "All" and "Default_Action."
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules.
	//
	// You can view the sampled requests through the AWS WAF console.
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}

// A rule statement that defines a cross-site scripting (XSS) match search for AWS WAF to apply to web requests.
//
// XSS attacks are those where the attacker uses vulnerabilities in a benign website as a vehicle to inject malicious client-site scripts into other legitimate web browsers. The XSS match statement provides the location in requests that you want AWS WAF to search and text transformations to use on the search area before AWS WAF searches for character sequences that are likely to be malicious strings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   xssMatchStatementProperty := &xssMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnRuleGroup_XssMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// Properties for defining a `CfnRuleGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allow interface{}
//   var allQueryArguments interface{}
//   var block interface{}
//   var captcha interface{}
//   var count interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   cfnRuleGroupProps := &cfnRuleGroupProps{
//   	capacity: jsii.Number(123),
//   	scope: jsii.String("scope"),
//   	visibilityConfig: &visibilityConfigProperty{
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   		metricName: jsii.String("metricName"),
//   		sampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	customResponseBodies: map[string]interface{}{
//   		"customResponseBodiesKey": &CustomResponseBodyProperty{
//   			"content": jsii.String("content"),
//   			"contentType": jsii.String("contentType"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			name: jsii.String("name"),
//   			priority: jsii.Number(123),
//   			statement: &statementProperty{
//   				andStatement: &andStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: &orStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   			visibilityConfig: &visibilityConfigProperty{
//   				cloudWatchMetricsEnabled: jsii.Boolean(false),
//   				metricName: jsii.String("metricName"),
//   				sampledRequestsEnabled: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			action: &ruleActionProperty{
//   				allow: allow,
//   				block: block,
//   				captcha: captcha,
//   				count: count,
//   			},
//   			captchaConfig: &captchaConfigProperty{
//   				immunityTimeProperty: &immunityTimePropertyProperty{
//   					immunityTime: jsii.Number(123),
//   				},
//   			},
//   			ruleLabels: []interface{}{
//   				&labelProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleGroupProps struct {
	// The web ACL capacity units (WCUs) required for this rule group.
	//
	// When you create your own rule group, you define this, and you cannot change it after creation. When you add or modify the rules in a rule group, AWS WAF enforces this limit.
	//
	// AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig interface{} `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// A map of custom response keys and content bodies.
	//
	// When you create a rule with a block action, you can send a custom response to the web request. You define these for the rule group, and then use them in the rules that you define in the rule group.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomResponseBodies interface{} `field:"optional" json:"customResponseBodies" yaml:"customResponseBodies"`
	// A description of the rule group that helps with identification.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the rule group.
	//
	// You cannot change the name of a rule group after you create it.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rule statements used to identify the web requests that you want to allow, block, or count.
	//
	// Each rule includes one top-level statement that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::WAFv2::WebACL`.
//
// > This is the latest version of *AWS WAF* , named AWS WAF V2, released in November, 2019. For information, including how to migrate your AWS WAF resources from the prior release, see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Use an `WebACL` to define a collection of rules to use to inspect and control web requests. Each rule has an action defined (allow, block, or count) for requests that match the statement of the rule. In the web ACL, you assign a default action to take (allow, block) for any request that does not match any of the rules. The rules in a web ACL can contain rule statements that you define explicitly and rule statements that reference rule groups and managed rule groups. You can associate a web ACL with one or more AWS resources to protect. The resources can be an Amazon CloudFront distribution, an Amazon API Gateway REST API, an Application Load Balancer , or an AWS AppSync GraphQL API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var count interface{}
//   var method interface{}
//   var none interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   cfnWebACL := awscdk.Aws_wafv2.NewCfnWebACL(this, jsii.String("MyCfnWebACL"), &cfnWebACLProps{
//   	defaultAction: &defaultActionProperty{
//   		allow: &allowActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		block: &blockActionProperty{
//   			customResponse: &customResponseProperty{
//   				responseCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   				responseHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	scope: jsii.String("scope"),
//   	visibilityConfig: &visibilityConfigProperty{
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   		metricName: jsii.String("metricName"),
//   		sampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	captchaConfig: &captchaConfigProperty{
//   		immunityTimeProperty: &immunityTimePropertyProperty{
//   			immunityTime: jsii.Number(123),
//   		},
//   	},
//   	customResponseBodies: map[string]interface{}{
//   		"customResponseBodiesKey": &CustomResponseBodyProperty{
//   			"content": jsii.String("content"),
//   			"contentType": jsii.String("contentType"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			name: jsii.String("name"),
//   			priority: jsii.Number(123),
//   			statement: &statementProperty{
//   				andStatement: &andStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   					name: jsii.String("name"),
//   					vendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					managedRuleGroupConfigs: []interface{}{
//   						&managedRuleGroupConfigProperty{
//   							loginPath: jsii.String("loginPath"),
//   							passwordField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   							payloadType: jsii.String("payloadType"),
//   							usernameField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					scopeDownStatement: statementProperty_,
//   					version: jsii.String("version"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: &orStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   			visibilityConfig: &visibilityConfigProperty{
//   				cloudWatchMetricsEnabled: jsii.Boolean(false),
//   				metricName: jsii.String("metricName"),
//   				sampledRequestsEnabled: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			action: &ruleActionProperty{
//   				allow: &allowActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				block: &blockActionProperty{
//   					customResponse: &customResponseProperty{
//   						responseCode: jsii.Number(123),
//
//   						// the properties below are optional
//   						customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   						responseHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				captcha: &captchaActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				count: &countActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			captchaConfig: &captchaConfigProperty{
//   				immunityTimeProperty: &immunityTimePropertyProperty{
//   					immunityTime: jsii.Number(123),
//   				},
//   			},
//   			overrideAction: &overrideActionProperty{
//   				count: count,
//   				none: none,
//   			},
//   			ruleLabels: []interface{}{
//   				&labelProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnWebACL interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the web ACL.
	AttrArn() *string
	// The web ACL capacity units (WCUs) currently being used by this web ACL.
	//
	// AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
	AttrCapacity() *float64
	// The ID of the web ACL.
	AttrId() *string
	// The label namespace prefix for this web ACL.
	//
	// All labels added by rules in this web ACL have this prefix.
	//
	// The syntax for the label namespace prefix for a web ACL is the following: `awswaf:<account ID>:webacl:<web ACL name>:`
	//
	// When a rule with a label matches a web request, AWS WAF adds the fully qualified label to the request. A fully qualified label is made up of the label namespace from the rule group or web ACL where the rule is defined and the label from the rule, separated by a colon.
	AttrLabelNamespace() *string
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings.
	//
	// If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
	CaptchaConfig() interface{}
	SetCaptchaConfig(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A map of custom response keys and content bodies.
	//
	// When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomResponseBodies() interface{}
	SetCustomResponseBodies(val interface{})
	// The action to perform if none of the `Rules` contained in the `WebACL` match.
	DefaultAction() interface{}
	SetDefaultAction(val interface{})
	// A description of the web ACL that helps with identification.
	Description() *string
	SetDescription(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the web ACL.
	//
	// You cannot change the name of a web ACL after you create it.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The rule statements used to identify the web requests that you want to allow, block, or count.
	//
	// Each rule includes one top-level statement that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
	Rules() interface{}
	SetRules(val interface{})
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	//
	// For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
	Scope() *string
	SetScope(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig() interface{}
	SetVisibilityConfig(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnWebACL
type jsiiProxy_CfnWebACL struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWebACL) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) AttrCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) AttrLabelNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLabelNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CaptchaConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captchaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) CustomResponseBodies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customResponseBodies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) DefaultAction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACL) VisibilityConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::WebACL`.
func NewCfnWebACL(scope awscdk.Construct, id *string, props *CfnWebACLProps) CfnWebACL {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACL{}

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnWebACL",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::WebACL`.
func NewCfnWebACL_Override(c CfnWebACL, scope awscdk.Construct, id *string, props *CfnWebACLProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnWebACL",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWebACL) SetCaptchaConfig(val interface{}) {
	_jsii_.Set(
		j,
		"captchaConfig",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetCustomResponseBodies(val interface{}) {
	_jsii_.Set(
		j,
		"customResponseBodies",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetDefaultAction(val interface{}) {
	_jsii_.Set(
		j,
		"defaultAction",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnWebACL) SetVisibilityConfig(val interface{}) {
	_jsii_.Set(
		j,
		"visibilityConfig",
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
func CfnWebACL_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnWebACL",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWebACL_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnWebACL",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnWebACL_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnWebACL",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebACL_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_wafv2.CfnWebACL",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWebACL) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWebACL) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWebACL) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWebACL) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWebACL) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWebACL) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWebACL) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWebACL) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACL) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACL) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWebACL) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWebACL) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWebACL) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACL) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWebACL) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWebACL) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACL) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACL) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWebACL) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACL) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACL) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies that AWS WAF should allow the request and optionally defines additional custom handling for the request.
//
// This is used in the context of other settings, for example to specify values for a rule action or a web ACL default action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowActionProperty := &allowActionProperty{
//   	customRequestHandling: &customRequestHandlingProperty{
//   		insertHeaders: []interface{}{
//   			&customHTTPHeaderProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_AllowActionProperty struct {
	// Defines custom handling for the web request.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

// A logical rule statement used to combine other rule statements with AND logic.
//
// You provide more than one `Statement` within the `AndStatement` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   andStatementProperty := &andStatementProperty{
//   	statements: []interface{}{
//   		&statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   				name: jsii.String("name"),
//   				vendorName: jsii.String("vendorName"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				managedRuleGroupConfigs: []interface{}{
//   					&managedRuleGroupConfigProperty{
//   						loginPath: jsii.String("loginPath"),
//   						passwordField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   						payloadType: jsii.String("payloadType"),
//   						usernameField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   					},
//   				},
//   				scopeDownStatement: statementProperty_,
//   				version: jsii.String("version"),
//   			},
//   			notStatement: &notStatementProperty{
//   				statement: statementProperty_,
//   			},
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: &rateBasedStatementProperty{
//   				aggregateKeyType: jsii.String("aggregateKeyType"),
//   				limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   				scopeDownStatement: statementProperty_,
//   			},
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_AndStatementProperty struct {
	// The statements to combine with AND logic.
	//
	// You can use any statements that can be nested.
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

// Specifies that AWS WAF should block the request and optionally defines additional custom handling for the response to the web request.
//
// This is used in the context of other settings, for example to specify values for a rule action or a web ACL default action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockActionProperty := &blockActionProperty{
//   	customResponse: &customResponseProperty{
//   		responseCode: jsii.Number(123),
//
//   		// the properties below are optional
//   		customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   		responseHeaders: []interface{}{
//   			&customHTTPHeaderProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_BlockActionProperty struct {
	// Defines a custom response for the web request.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomResponse interface{} `field:"optional" json:"customResponse" yaml:"customResponse"`
}

// Inspect the body of the web request. The body immediately follows the request headers.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodyProperty := &bodyProperty{
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnWebACL_BodyProperty struct {
	// What AWS WAF should do if the body is larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of the body of a web request when the body exceeds 8 KB (8192 bytes). Only the first 8 KB of the request body are forwarded to AWS WAF by the underlying host service.
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the body normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// You can combine the `MATCH` or `NO_MATCH` settings for oversize handling with your rule and web ACL action settings, so that you block any request whose body is over 8 KB.
	//
	// Default: `CONTINUE`.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

// A rule statement that defines a string match search for AWS WAF to apply to web requests.
//
// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is refered to as a string match statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   byteMatchStatementProperty := &byteMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	positionalConstraint: jsii.String("positionalConstraint"),
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	searchString: jsii.String("searchString"),
//   	searchStringBase64: jsii.String("searchStringBase64"),
//   }
//
type CfnWebACL_ByteMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The area within the portion of the web request that you want AWS WAF to search for `SearchString` .
	//
	// Valid values include the following:
	//
	// *CONTAINS*
	//
	// The specified part of the web request must include the value of `SearchString` , but the location doesn't matter.
	//
	// *CONTAINS_WORD*
	//
	// The specified part of the web request must include the value of `SearchString` , and `SearchString` must contain only alphanumeric characters or underscore (A-Z, a-z, 0-9, or _). In addition, `SearchString` must be a word, which means that both of the following are true:
	//
	// - `SearchString` is at the beginning of the specified part of the web request or is preceded by a character other than an alphanumeric character or underscore (_). Examples include the value of a header and `;BadBot` .
	// - `SearchString` is at the end of the specified part of the web request or is followed by a character other than an alphanumeric character or underscore (_), for example, `BadBot;` and `-BadBot;` .
	//
	// *EXACTLY*
	//
	// The value of the specified part of the web request must exactly match the value of `SearchString` .
	//
	// *STARTS_WITH*
	//
	// The value of `SearchString` must appear at the beginning of the specified part of the web request.
	//
	// *ENDS_WITH*
	//
	// The value of `SearchString` must appear at the end of the specified part of the web request.
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
	// A string value that you want AWS WAF to search for.
	//
	// AWS WAF searches only in the part of web requests that you designate for inspection in `FieldToMatch` . The maximum length of the value is 50 bytes. For alphabetic characters A-Z and a-z, the value is case sensitive.
	//
	// Don't encode this string. Provide the value that you want AWS WAF to search for. AWS CloudFormation automatically base64 encodes the value for you.
	//
	// For example, suppose the value of `Type` is `HEADER` and the value of `Data` is `User-Agent` . If you want to search the `User-Agent` header for the value `BadBot` , you provide the string `BadBot` in the value of `SearchString` .
	//
	// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
	SearchString *string `field:"optional" json:"searchString" yaml:"searchString"`
	// String to search for in a web request component, base64-encoded.
	//
	// If you don't want to encode the string, specify the unencoded value in `SearchString` instead.
	//
	// You must specify either `SearchString` or `SearchStringBase64` in a `ByteMatchStatement` .
	SearchStringBase64 *string `field:"optional" json:"searchStringBase64" yaml:"searchStringBase64"`
}

// Specifies that AWS WAF should run a `CAPTCHA` check against the request:.
//
// - If the request includes a valid, unexpired `CAPTCHA` token, AWS WAF allows the web request inspection to proceed to the next rule, similar to a `CountAction` .
// - If the request doesn't include a valid, unexpired `CAPTCHA` token, AWS WAF discontinues the web ACL evaluation of the request and blocks it from going to its intended destination.
//
// AWS WAF generates a response that it sends back to the client, which includes the following:
//
// - The header `x-amzn-waf-action` with a value of `captcha` .
// - The HTTP status code `405 Method Not Allowed` .
// - If the request contains an `Accept` header with a value of `text/html` , the response includes a `CAPTCHA` challenge.
//
// You can configure the expiration time in the `CaptchaConfig` `ImmunityTimeProperty` setting at the rule and web ACL level. The rule setting overrides the web ACL setting.
//
// This action option is available for rules. It isn't available for web ACL default actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captchaActionProperty := &captchaActionProperty{
//   	customRequestHandling: &customRequestHandlingProperty{
//   		insertHeaders: []interface{}{
//   			&customHTTPHeaderProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_CaptchaActionProperty struct {
	// Defines custom handling for the web request.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings.
//
// If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captchaConfigProperty := &captchaConfigProperty{
//   	immunityTimeProperty: &immunityTimePropertyProperty{
//   		immunityTime: jsii.Number(123),
//   	},
//   }
//
type CfnWebACL_CaptchaConfigProperty struct {
	// Determines how long a `CAPTCHA` token remains valid after the client successfully solves a `CAPTCHA` puzzle.
	ImmunityTimeProperty interface{} `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

// The filter to use to identify the subset of cookies to inspect in a web request.
//
// You must specify exactly one setting: either `All` , `IncludedCookies` , or `ExcludedCookies` .
//
// Example JSON: `"MatchPattern": { "IncludedCookies": {"KeyToInclude1", "KeyToInclude2", "KeyToInclude3"} }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   cookieMatchPatternProperty := &cookieMatchPatternProperty{
//   	all: all,
//   	excludedCookies: []*string{
//   		jsii.String("excludedCookies"),
//   	},
//   	includedCookies: []*string{
//   		jsii.String("includedCookies"),
//   	},
//   }
//
type CfnWebACL_CookieMatchPatternProperty struct {
	// Inspect all cookies.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Inspect only the cookies whose keys don't match any of the strings specified here.
	ExcludedCookies *[]*string `field:"optional" json:"excludedCookies" yaml:"excludedCookies"`
	// Inspect only the cookies that have a key that matches one of the strings specified here.
	IncludedCookies *[]*string `field:"optional" json:"includedCookies" yaml:"includedCookies"`
}

// Inspect the cookies in the web request.
//
// You can specify the parts of the cookies to inspect and you can narrow the set of cookies to inspect by including or excluding specific keys.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Example JSON: `"Cookies": { "MatchPattern": { "All": {} }, "MatchScope": "KEY", "OversizeHandling": "MATCH" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   cookiesProperty := &cookiesProperty{
//   	matchPattern: &cookieMatchPatternProperty{
//   		all: all,
//   		excludedCookies: []*string{
//   			jsii.String("excludedCookies"),
//   		},
//   		includedCookies: []*string{
//   			jsii.String("includedCookies"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnWebACL_CookiesProperty struct {
	// The filter to use to identify the subset of cookies to inspect in a web request.
	//
	// You must specify exactly one setting: either `All` , `IncludedCookies` , or `ExcludedCookies` .
	//
	// Example JSON: `"MatchPattern": { "IncludedCookies": {"KeyToInclude1", "KeyToInclude2", "KeyToInclude3"} }`.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the cookies to inspect with the rule inspection criteria.
	//
	// If you specify `All` , AWS WAF inspects both keys and values.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if the cookies of the request are larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of request cookies when they exceed 8 KB (8192 bytes) or 200 total cookies. The underlying host service forwards a maximum of 200 cookies and at most 8 KB of cookie contents to AWS WAF .
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the cookies normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

// Specifies that AWS WAF should count the request. Optionally defines additional custom handling for the request.
//
// This is used in the context of other settings, for example to specify values for a rule action or a web ACL default action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   countActionProperty := &countActionProperty{
//   	customRequestHandling: &customRequestHandlingProperty{
//   		insertHeaders: []interface{}{
//   			&customHTTPHeaderProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_CountActionProperty struct {
	// Defines custom handling for the web request.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

// A custom header for custom request and response handling.
//
// This is used in `CustomResponse` and `CustomRequestHandling` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customHTTPHeaderProperty := &customHTTPHeaderProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnWebACL_CustomHTTPHeaderProperty struct {
	// The name of the custom header.
	//
	// For custom request header insertion, when AWS WAF inserts the header into the request, it prefixes this name `x-amzn-waf-` , to avoid confusion with the headers that are already in the request. For example, for the header name `sample` , AWS WAF inserts the header `x-amzn-waf-sample` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the custom header.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Custom request handling behavior that inserts custom headers into a web request.
//
// You can add custom request handling for the rule actions allow and count.
//
// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customRequestHandlingProperty := &customRequestHandlingProperty{
//   	insertHeaders: []interface{}{
//   		&customHTTPHeaderProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWebACL_CustomRequestHandlingProperty struct {
	// The HTTP headers to insert into the request. Duplicate header names are not allowed.
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	InsertHeaders interface{} `field:"required" json:"insertHeaders" yaml:"insertHeaders"`
}

// The response body to use in a custom response to a web request.
//
// This is referenced by key from `CustomResponse` `CustomResponseBodyKey` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResponseBodyProperty := &customResponseBodyProperty{
//   	content: jsii.String("content"),
//   	contentType: jsii.String("contentType"),
//   }
//
type CfnWebACL_CustomResponseBodyProperty struct {
	// The payload of the custom response.
	//
	// You can use JSON escape strings in JSON content. To do this, you must specify JSON content in the `ContentType` setting.
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	Content *string `field:"required" json:"content" yaml:"content"`
	// The type of content in the payload that you are defining in the `Content` string.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
}

// A custom response to send to the client.
//
// You can define a custom response for rule actions and default web ACL actions that are set to the block action.
//
// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customResponseProperty := &customResponseProperty{
//   	responseCode: jsii.Number(123),
//
//   	// the properties below are optional
//   	customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   	responseHeaders: []interface{}{
//   		&customHTTPHeaderProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWebACL_CustomResponseProperty struct {
	// The HTTP status code to return to the client.
	//
	// For a list of status codes that you can use in your custom reqponses, see [Supported status codes for custom response](https://docs.aws.amazon.com/waf/latest/developerguide/customizing-the-response-status-codes.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
	// References the response body that you want AWS WAF to return to the web request client.
	//
	// You can define a custom response for a rule action or a default web ACL action that is set to block. To do this, you first define the response body key and value in the `CustomResponseBodies` setting for the `WebACL` or `RuleGroup` where you want to use it. Then, in the rule action or web ACL default action `BlockAction` setting, you reference the response body using this key.
	CustomResponseBodyKey *string `field:"optional" json:"customResponseBodyKey" yaml:"customResponseBodyKey"`
	// The HTTP headers to use in the response. Duplicate header names are not allowed.
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	ResponseHeaders interface{} `field:"optional" json:"responseHeaders" yaml:"responseHeaders"`
}

// In a `WebACL` , this is the action that you want AWS WAF to perform when a web request doesn't match any of the rules in the `WebACL` .
//
// The default action must be a terminating action, so you can't use count.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultActionProperty := &defaultActionProperty{
//   	allow: &allowActionProperty{
//   		customRequestHandling: &customRequestHandlingProperty{
//   			insertHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	block: &blockActionProperty{
//   		customResponse: &customResponseProperty{
//   			responseCode: jsii.Number(123),
//
//   			// the properties below are optional
//   			customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   			responseHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_DefaultActionProperty struct {
	// Specifies that AWS WAF should allow requests by default.
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Specifies that AWS WAF should block requests by default.
	Block interface{} `field:"optional" json:"block" yaml:"block"`
}

// Specifies a single rule in a rule group whose action you want to override to `Count` .
//
// When you exclude a rule, AWS WAF evaluates it exactly as it would if the rule action setting were `Count` . This is a useful option for testing the rules in a rule group without modifying how they handle your web traffic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   excludedRuleProperty := &excludedRuleProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnWebACL_ExcludedRuleProperty struct {
	// The name of the rule whose action you want to override to `Count` .
	Name *string `field:"required" json:"name" yaml:"name"`
}

// The identifier of the username or password field, used in the `ManagedRuleGroupConfig` settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldIdentifierProperty := &fieldIdentifierProperty{
//   	identifier: jsii.String("identifier"),
//   }
//
type CfnWebACL_FieldIdentifierProperty struct {
	// The name of the username or password field, used in the `ManagedRuleGroupConfig` settings.
	//
	// When the `PayloadType` is `JSON` , the identifier must be in JSON pointer syntax. For example `/form/username` . For information about the JSON Pointer syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// When the `PayloadType` is `FORM_ENCODED` , use the HTML form names. For example, `username` .
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

// The part of the web request that you want AWS WAF to inspect.
//
// Include the single `FieldToMatch` type that you want to inspect, with additional specifications as needed, according to the type. You specify a single request component in `FieldToMatch` for each rule statement that requires it. To inspect more than one component of the web request, create a separate rule statement for each component.
//
// Example JSON for a `QueryString` field to match:
//
// `"FieldToMatch": { "QueryString": {} }`
//
// Example JSON for a `Method` field to match specification:
//
// `"FieldToMatch": { "Method": { "Name": "DELETE" } }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   fieldToMatchProperty := &fieldToMatchProperty{
//   	allQueryArguments: allQueryArguments,
//   	body: &bodyProperty{
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	cookies: &cookiesProperty{
//   		matchPattern: &cookieMatchPatternProperty{
//   			all: all,
//   			excludedCookies: []*string{
//   				jsii.String("excludedCookies"),
//   			},
//   			includedCookies: []*string{
//   				jsii.String("includedCookies"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	headers: &headersProperty{
//   		matchPattern: &headerMatchPatternProperty{
//   			all: all,
//   			excludedHeaders: []*string{
//   				jsii.String("excludedHeaders"),
//   			},
//   			includedHeaders: []*string{
//   				jsii.String("includedHeaders"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	jsonBody: &jsonBodyProperty{
//   		matchPattern: &jsonMatchPatternProperty{
//   			all: all,
//   			includedPaths: []*string{
//   				jsii.String("includedPaths"),
//   			},
//   		},
//   		matchScope: jsii.String("matchScope"),
//
//   		// the properties below are optional
//   		invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   		oversizeHandling: jsii.String("oversizeHandling"),
//   	},
//   	method: method,
//   	queryString: queryString,
//   	singleHeader: singleHeader,
//   	singleQueryArgument: singleQueryArgument,
//   	uriPath: uriPath,
//   }
//
type CfnWebACL_FieldToMatchProperty struct {
	// Inspect all query arguments.
	AllQueryArguments interface{} `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// Inspect the request body as plain text.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// Only the first 8 KB (8192 bytes) of the request body are forwarded to AWS WAF for inspection by the underlying host service. For information about how to handle oversized request bodies, see the `Body` object configuration.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// Inspect the request cookies.
	//
	// You must configure scope and pattern matching filters in the `Cookies` object, to define the set of cookies and the parts of the cookies that AWS WAF inspects.
	//
	// Only the first 8 KB (8192 bytes) of a request's cookies and only the first 200 cookies are forwarded to AWS WAF for inspection by the underlying host service. You must configure how to handle any oversize cookie content in the `Cookies` object. AWS WAF applies the pattern matching filters to the cookies that it receives from the underlying host service.
	Cookies interface{} `field:"optional" json:"cookies" yaml:"cookies"`
	// Inspect the request headers.
	//
	// You must configure scope and pattern matching filters in the `Headers` object, to define the set of headers to and the parts of the headers that AWS WAF inspects.
	//
	// Only the first 8 KB (8192 bytes) of a request's headers and only the first 200 headers are forwarded to AWS WAF for inspection by the underlying host service. You must configure how to handle any oversize header content in the `Headers` object. AWS WAF applies the pattern matching filters to the headers that it receives from the underlying host service.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Inspect the request body as JSON.
	//
	// The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form.
	//
	// Only the first 8 KB (8192 bytes) of the request body are forwarded to AWS WAF for inspection by the underlying host service. For information about how to handle oversized request bodies, see the `JsonBody` object configuration.
	JsonBody interface{} `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// Inspect the HTTP method.
	//
	// The method indicates the type of operation that the request is asking the origin to perform.
	Method interface{} `field:"optional" json:"method" yaml:"method"`
	// Inspect the query string.
	//
	// This is the part of a URL that appears after a `?` character, if any.
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// Inspect a single header.
	//
	// Provide the name of the header to inspect, for example, `User-Agent` or `Referer` . This setting isn't case sensitive.
	//
	// Example JSON: `"SingleHeader": { "Name": "haystack" }`
	//
	// Alternately, you can filter and inspect all headers with the `Headers` `FieldToMatch` setting.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Inspect a single query argument.
	//
	// Provide the name of the query argument to inspect, such as *UserName* or *SalesRegion* . The name can be up to 30 characters long and isn't case sensitive.
	//
	// Example JSON: `"SingleQueryArgument": { "Name": "myArgument" }`.
	SingleQueryArgument interface{} `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// Inspect the request URI path.
	//
	// This is the part of the web request that identifies a resource, for example, `/images/daily-ad.jpg` .
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
//
// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
//
// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
//
// This configuration is used for `GeoMatchStatement` and `RateBasedStatement` . For `IPSetReferenceStatement` , use `IPSetForwardedIPConfig` instead.
//
// AWS WAF only evaluates the first IP address found in the specified HTTP header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forwardedIPConfigurationProperty := &forwardedIPConfigurationProperty{
//   	fallbackBehavior: jsii.String("fallbackBehavior"),
//   	headerName: jsii.String("headerName"),
//   }
//
type CfnWebACL_ForwardedIPConfigurationProperty struct {
	// The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// You can specify the following fallback behaviors:
	//
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// The name of the HTTP header to use for the IP address.
	//
	// For example, to use the X-Forwarded-For (XFF) header, set this to `X-Forwarded-For` .
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

// A rule statement used to identify web requests based on country of origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoMatchStatementProperty := &geoMatchStatementProperty{
//   	countryCodes: []*string{
//   		jsii.String("countryCodes"),
//   	},
//   	forwardedIpConfig: &forwardedIPConfigurationProperty{
//   		fallbackBehavior: jsii.String("fallbackBehavior"),
//   		headerName: jsii.String("headerName"),
//   	},
//   }
//
type CfnWebACL_GeoMatchStatementProperty struct {
	// An array of two-character country codes, for example, `[ "US", "CN" ]` , from the alpha-2 country ISO codes of the ISO 3166 international standard.
	CountryCodes *[]*string `field:"optional" json:"countryCodes" yaml:"countryCodes"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

// The filter to use to identify the subset of headers to inspect in a web request.
//
// You must specify exactly one setting: either `All` , `IncludedHeaders` , or `ExcludedHeaders` .
//
// Example JSON: `"MatchPattern": { "ExcludedHeaders": {"KeyToExclude1", "KeyToExclude2"} }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   headerMatchPatternProperty := &headerMatchPatternProperty{
//   	all: all,
//   	excludedHeaders: []*string{
//   		jsii.String("excludedHeaders"),
//   	},
//   	includedHeaders: []*string{
//   		jsii.String("includedHeaders"),
//   	},
//   }
//
type CfnWebACL_HeaderMatchPatternProperty struct {
	// Inspect all headers.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Inspect only the headers whose keys don't match any of the strings specified here.
	ExcludedHeaders *[]*string `field:"optional" json:"excludedHeaders" yaml:"excludedHeaders"`
	// Inspect only the headers that have a key that matches one of the strings specified here.
	IncludedHeaders *[]*string `field:"optional" json:"includedHeaders" yaml:"includedHeaders"`
}

// Inspect all headers in the web request.
//
// You can specify the parts of the headers to inspect and you can narrow the set of headers to inspect by including or excluding specific keys.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// If you want to inspect just the value of a single header, use the `SingleHeader` `FieldToMatch` setting instead.
//
// Example JSON: `"Headers": { "MatchPattern": { "All": {} }, "MatchScope": "KEY", "OversizeHandling": "MATCH" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   headersProperty := &headersProperty{
//   	matchPattern: &headerMatchPatternProperty{
//   		all: all,
//   		excludedHeaders: []*string{
//   			jsii.String("excludedHeaders"),
//   		},
//   		includedHeaders: []*string{
//   			jsii.String("includedHeaders"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnWebACL_HeadersProperty struct {
	// The filter to use to identify the subset of headers to inspect in a web request.
	//
	// You must specify exactly one setting: either `All` , `IncludedHeaders` , or `ExcludedHeaders` .
	//
	// Example JSON: `"MatchPattern": { "ExcludedHeaders": {"KeyToExclude1", "KeyToExclude2"} }`.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the headers to match with the rule inspection criteria.
	//
	// If you specify `All` , AWS WAF inspects both keys and values.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if the headers of the request are larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of request headers when they exceed 8 KB (8192 bytes) or 200 total headers. The underlying host service forwards a maximum of 200 headers and at most 8 KB of header contents to AWS WAF .
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the headers normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
//
// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
//
// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
//
// This configuration is used only for `IPSetReferenceStatement` . For `GeoMatchStatement` and `RateBasedStatement` , use `ForwardedIPConfig` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetForwardedIPConfigurationProperty := map[string]*string{
//   	"fallbackBehavior": jsii.String("fallbackBehavior"),
//   	"headerName": jsii.String("headerName"),
//   	"position": jsii.String("position"),
//   }
//
type CfnWebACL_IPSetForwardedIPConfigurationProperty struct {
	// The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// You can specify the following fallback behaviors:
	//
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// The name of the HTTP header to use for the IP address.
	//
	// For example, to use the X-Forwarded-For (XFF) header, set this to `X-Forwarded-For` .
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The position in the header to search for the IP address.
	//
	// The header can contain IP addresses of the original client and also of proxies. For example, the header value could be `10.1.1.1, 127.0.0.0, 10.10.10.10` where the first IP address identifies the original client and the rest identify proxies that the request went through.
	//
	// The options for this setting are the following:
	//
	// - FIRST - Inspect the first IP address in the list of IP addresses in the header. This is usually the client's original IP.
	// - LAST - Inspect the last IP address in the list of IP addresses in the header.
	// - ANY - Inspect all IP addresses in the header for a match. If the header contains more than 10 IP addresses, AWS WAF inspects the last 10.
	Position *string `field:"required" json:"position" yaml:"position"`
}

// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
//
// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
//
// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPSetReferenceStatementProperty := map[string]interface{}{
//   	"arn": jsii.String("arn"),
//
//   	// the properties below are optional
//   	"ipSetForwardedIpConfig": map[string]*string{
//   		"fallbackBehavior": jsii.String("fallbackBehavior"),
//   		"headerName": jsii.String("headerName"),
//   		"position": jsii.String("position"),
//   	},
//   }
//
type CfnWebACL_IPSetReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the `IPSet` that this statement references.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	IpSetForwardedIpConfig interface{} `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

// Determines how long a `CAPTCHA` token remains valid after the client successfully solves a `CAPTCHA` puzzle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   immunityTimePropertyProperty := &immunityTimePropertyProperty{
//   	immunityTime: jsii.Number(123),
//   }
//
type CfnWebACL_ImmunityTimePropertyProperty struct {
	// The amount of time, in seconds, that a `CAPTCHA` token is valid.
	//
	// The default setting is 300.
	ImmunityTime *float64 `field:"required" json:"immunityTime" yaml:"immunityTime"`
}

// Inspect the body of the web request as JSON. The body immediately follows the request headers.
//
// This is used to indicate the web request component to inspect, in the `FieldToMatch` specification.
//
// Use the specifications in this object to indicate which parts of the JSON body to inspect using the rule's inspection criteria. AWS WAF inspects only the parts of the JSON that result from the matches that you indicate.
//
// Example JSON: `"JsonBody": { "MatchPattern": { "All": {} }, "MatchScope": "ALL" }`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   jsonBodyProperty := &jsonBodyProperty{
//   	matchPattern: &jsonMatchPatternProperty{
//   		all: all,
//   		includedPaths: []*string{
//   			jsii.String("includedPaths"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//
//   	// the properties below are optional
//   	invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   	oversizeHandling: jsii.String("oversizeHandling"),
//   }
//
type CfnWebACL_JsonBodyProperty struct {
	// The patterns to look for in the JSON body.
	//
	// AWS WAF inspects the results of these pattern matches against the rule inspection criteria.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the JSON to match against using the `MatchPattern` .
	//
	// If you specify `All` , AWS WAF matches against keys and values.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if it fails to completely parse the JSON body. The options are the following:.
	//
	// - `EVALUATE_AS_STRING` - Inspect the body as plain text. AWS WAF applies the text transformations and inspection criteria that you defined for the JSON inspection to the body text string.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// If you don't provide this setting, AWS WAF parses and evaluates the content only up to the first parsing failure that it encounters.
	//
	// AWS WAF does its best to parse the entire JSON body, but might be forced to stop for reasons such as invalid characters, duplicate keys, truncation, and any content whose root node isn't an object or an array.
	//
	// AWS WAF parses the JSON in the following examples as two valid key, value pairs:
	//
	// - Missing comma: `{"key1":"value1""key2":"value2"}`
	// - Missing colon: `{"key1":"value1","key2""value2"}`
	// - Extra colons: `{"key1"::"value1","key2""value2"}`.
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
	// What AWS WAF should do if the body is larger than AWS WAF can inspect.
	//
	// AWS WAF does not support inspecting the entire contents of the body of a web request when the body exceeds 8 KB (8192 bytes). Only the first 8 KB of the request body are forwarded to AWS WAF by the underlying host service.
	//
	// The options for oversize handling are the following:
	//
	// - `CONTINUE` - Inspect the body normally, according to the rule inspection criteria.
	// - `MATCH` - Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.
	// - `NO_MATCH` - Treat the web request as not matching the rule statement.
	//
	// You can combine the `MATCH` or `NO_MATCH` settings for oversize handling with your rule and web ACL action settings, so that you block any request whose body is over 8 KB.
	//
	// Default: `CONTINUE`.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

// The patterns to look for in the JSON body.
//
// AWS WAF inspects the results of these pattern matches against the rule inspection criteria. This is used with the `FieldToMatch` option `JsonBody` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   jsonMatchPatternProperty := &jsonMatchPatternProperty{
//   	all: all,
//   	includedPaths: []*string{
//   		jsii.String("includedPaths"),
//   	},
//   }
//
type CfnWebACL_JsonMatchPatternProperty struct {
	// Match all of the elements. See also `MatchScope` in the `JsonBody` `FieldToMatch` specification.
	//
	// You must specify either this setting or the `IncludedPaths` setting, but not both.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Match only the specified include paths. See also `MatchScope` in the `JsonBody` `FieldToMatch` specification.
	//
	// Provide the include paths using JSON Pointer syntax. For example, `"IncludedPaths": ["/dogs/0/name", "/dogs/1/name"]` . For information about this syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// You must specify either this setting or the `All` setting, but not both.
	//
	// > Don't use this option to include all paths. Instead, use the `All` setting.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

// A rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL.
//
// The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelMatchStatementProperty := &labelMatchStatementProperty{
//   	key: jsii.String("key"),
//   	scope: jsii.String("scope"),
//   }
//
type CfnWebACL_LabelMatchStatementProperty struct {
	// The string to match against. The setting you provide for this depends on the match statement's `Scope` setting:.
	//
	// - If the `Scope` indicates `LABEL` , then this specification must include the name and can include any number of preceding namespace specifications and prefix up to providing the fully qualified label name.
	// - If the `Scope` indicates `NAMESPACE` , then this specification can include any number of contiguous namespace strings, and can include the entire label namespace prefix from the rule group or web ACL where the label originates.
	//
	// Labels are case sensitive and components of a label must be separated by colon, for example `NS1:NS2:name` .
	Key *string `field:"required" json:"key" yaml:"key"`
	// Specify whether you want to match using the label name or just the namespace.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

// A single label container.
//
// This is used as an element of a label array in `RuleLabels` inside a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelProperty := &labelProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnWebACL_LabelProperty struct {
	// The label string.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Additional information that's used by a managed rule group. Most managed rule groups don't require this.
//
// Use this for the account takeover prevention managed rule group `AWSManagedRulesATPRuleSet` , to provide information about the sign-in page of your application.
//
// You can provide multiple individual `ManagedRuleGroupConfig` objects for any rule group configuration, for example `UsernameField` and `PasswordField` . The configuration that you provide depends on the needs of the managed rule group. For the ATP managed rule group, you provide the following individual configuration objects: `LoginPath` , `PasswordField` , `PayloadType` and `UsernameField` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedRuleGroupConfigProperty := &managedRuleGroupConfigProperty{
//   	loginPath: jsii.String("loginPath"),
//   	passwordField: &fieldIdentifierProperty{
//   		identifier: jsii.String("identifier"),
//   	},
//   	payloadType: jsii.String("payloadType"),
//   	usernameField: &fieldIdentifierProperty{
//   		identifier: jsii.String("identifier"),
//   	},
//   }
//
type CfnWebACL_ManagedRuleGroupConfigProperty struct {
	// The path of the login endpoint for your application.
	//
	// For example, for the URL `https://example.com/web/login` , you would provide the path `/web/login` .
	LoginPath *string `field:"optional" json:"loginPath" yaml:"loginPath"`
	// Details about your login page password field.
	PasswordField interface{} `field:"optional" json:"passwordField" yaml:"passwordField"`
	// The payload type for your login endpoint, either JSON or form encoded.
	PayloadType *string `field:"optional" json:"payloadType" yaml:"payloadType"`
	// Details about your login page username field.
	UsernameField interface{} `field:"optional" json:"usernameField" yaml:"usernameField"`
}

// A rule statement used to run the rules that are defined in a managed rule group.
//
// To use this, provide the vendor name and the name of the rule group in this statement.
//
// You cannot nest a `ManagedRuleGroupStatement` , for example for use inside a `NotStatement` or `OrStatement` . It can only be referenced as a top-level statement within a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   managedRuleGroupStatementProperty := &managedRuleGroupStatementProperty{
//   	name: jsii.String("name"),
//   	vendorName: jsii.String("vendorName"),
//
//   	// the properties below are optional
//   	excludedRules: []interface{}{
//   		&excludedRuleProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	managedRuleGroupConfigs: []interface{}{
//   		&managedRuleGroupConfigProperty{
//   			loginPath: jsii.String("loginPath"),
//   			passwordField: &fieldIdentifierProperty{
//   				identifier: jsii.String("identifier"),
//   			},
//   			payloadType: jsii.String("payloadType"),
//   			usernameField: &fieldIdentifierProperty{
//   				identifier: jsii.String("identifier"),
//   			},
//   		},
//   	},
//   	scopeDownStatement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   			name: jsii.String("name"),
//   			vendorName: jsii.String("vendorName"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			managedRuleGroupConfigs: []interface{}{
//   				&managedRuleGroupConfigProperty{
//   					loginPath: jsii.String("loginPath"),
//   					passwordField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   					payloadType: jsii.String("payloadType"),
//   					usernameField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   				},
//   			},
//   			scopeDownStatement: statementProperty_,
//   			version: jsii.String("version"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	version: jsii.String("version"),
//   }
//
type CfnWebACL_ManagedRuleGroupStatementProperty struct {
	// The name of the managed rule group.
	//
	// You use this, along with the vendor name, to identify the rule group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the managed rule group vendor.
	//
	// You use this, along with the rule group name, to identify the rule group.
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// The rules in the referenced rule group whose actions are set to `Count` .
	//
	// When you exclude a rule, AWS WAF evaluates it exactly as it would if the rule action setting were `Count` . This is a useful option for testing the rules in a rule group without modifying how they handle your web traffic.
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
	// Additional information that's used by a managed rule group. Most managed rule groups don't require this.
	//
	// Use this for the account takeover prevention managed rule group `AWSManagedRulesATPRuleSet` , to provide information about the sign-in page of your application.
	//
	// You can provide multiple individual `ManagedRuleGroupConfig` objects for any rule group configuration, for example `UsernameField` and `PasswordField` . The configuration that you provide depends on the needs of the managed rule group. For the ATP managed rule group, you provide the following individual configuration objects: `LoginPath` , `PasswordField` , `PayloadType` and `UsernameField` .
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the managed rule group.
	//
	// Requests are only evaluated by the rule group if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// The version of the managed rule group to use.
	//
	// If you specify this, the version setting is fixed until you change it. If you don't specify this, AWS WAF uses the vendor's default version, and then keeps the version at the vendor's default when the vendor updates the managed rule group settings.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// A logical rule statement used to negate the results of another rule statement.
//
// You provide one `Statement` within the `NotStatement` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   notStatementProperty := &notStatementProperty{
//   	statement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   			name: jsii.String("name"),
//   			vendorName: jsii.String("vendorName"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			managedRuleGroupConfigs: []interface{}{
//   				&managedRuleGroupConfigProperty{
//   					loginPath: jsii.String("loginPath"),
//   					passwordField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   					payloadType: jsii.String("payloadType"),
//   					usernameField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   				},
//   			},
//   			scopeDownStatement: statementProperty_,
//   			version: jsii.String("version"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_NotStatementProperty struct {
	// The statement to negate.
	//
	// You can use any statement that can be nested.
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
}

// A logical rule statement used to combine other rule statements with OR logic.
//
// You provide more than one `Statement` within the `OrStatement` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   orStatementProperty := &orStatementProperty{
//   	statements: []interface{}{
//   		&statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   				name: jsii.String("name"),
//   				vendorName: jsii.String("vendorName"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				managedRuleGroupConfigs: []interface{}{
//   					&managedRuleGroupConfigProperty{
//   						loginPath: jsii.String("loginPath"),
//   						passwordField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   						payloadType: jsii.String("payloadType"),
//   						usernameField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   					},
//   				},
//   				scopeDownStatement: statementProperty_,
//   				version: jsii.String("version"),
//   			},
//   			notStatement: &notStatementProperty{
//   				statement: statementProperty_,
//   			},
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: &rateBasedStatementProperty{
//   				aggregateKeyType: jsii.String("aggregateKeyType"),
//   				limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   				scopeDownStatement: statementProperty_,
//   			},
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_OrStatementProperty struct {
	// The statements to combine with OR logic.
	//
	// You can use any statements that can be nested.
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

// The action to use in the place of the action that results from the rule group evaluation.
//
// Set the override action to none to leave the result of the rule group alone. Set it to count to override the result to count only.
//
// You can only use this for rule statements that reference a rule group, like `RuleGroupReferenceStatement` and `ManagedRuleGroupStatement` .
//
// > This option is usually set to none. It does not affect how the rules in the rule group are evaluated. If you want the rules in the rule group to only count matches, do not use this and instead exclude those rules in your rule group reference statement settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var count interface{}
//   var none interface{}
//
//   overrideActionProperty := &overrideActionProperty{
//   	count: count,
//   	none: none,
//   }
//
type CfnWebACL_OverrideActionProperty struct {
	// Override the rule group evaluation result to count only.
	//
	// > This option is usually set to none. It does not affect how the rules in the rule group are evaluated. If you want the rules in the rule group to only count matches, do not use this and instead exclude those rules in your rule group reference statement settings.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Don't override the rule group evaluation result.
	//
	// This is the most common setting.
	None interface{} `field:"optional" json:"none" yaml:"none"`
}

// A rate-based rule tracks the rate of requests for each originating IP address, and triggers the rule action when the rate exceeds a limit that you specify on the number of requests in any 5-minute time span.
//
// You can use this to put a temporary block on requests from an IP address that is sending excessive requests.
//
// AWS WAF tracks and manages web requests separately for each instance of a rate-based rule that you use. For example, if you provide the same rate-based rule settings in two web ACLs, each of the two rule statements represents a separate instance of the rate-based rule and gets its own tracking and management by AWS WAF . If you define a rate-based rule inside a rule group, and then use that rule group in multiple places, each use creates a separate instance of the rate-based rule that gets its own tracking and management by AWS WAF .
//
// When the rule action triggers, AWS WAF blocks additional requests from the IP address until the request rate falls below the limit.
//
// You can optionally nest another statement inside the rate-based statement, to narrow the scope of the rule so that it only counts requests that match the nested statement. For example, based on recent requests that you have seen from an attacker, you might create a rate-based rule with a nested AND rule statement that contains the following nested statements:
//
// - An IP match statement with an IP set that specified the address 192.0.2.44.
// - A string match statement that searches in the User-Agent header for the string BadBot.
//
// In this rate-based rule, you also define a rate limit. For this example, the rate limit is 1,000. Requests that meet both of the conditions in the statements are counted. If the count exceeds 1,000 requests per five minutes, the rule action triggers. Requests that do not meet both conditions are not counted towards the rate limit and are not affected by this rule.
//
// You cannot nest a `RateBasedStatement` inside another statement, for example inside a `NotStatement` or `OrStatement` . You can define a `RateBasedStatement` inside a web ACL and inside a rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   rateBasedStatementProperty := &rateBasedStatementProperty{
//   	aggregateKeyType: jsii.String("aggregateKeyType"),
//   	limit: jsii.Number(123),
//
//   	// the properties below are optional
//   	forwardedIpConfig: &forwardedIPConfigurationProperty{
//   		fallbackBehavior: jsii.String("fallbackBehavior"),
//   		headerName: jsii.String("headerName"),
//   	},
//   	scopeDownStatement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   			name: jsii.String("name"),
//   			vendorName: jsii.String("vendorName"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			managedRuleGroupConfigs: []interface{}{
//   				&managedRuleGroupConfigProperty{
//   					loginPath: jsii.String("loginPath"),
//   					passwordField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   					payloadType: jsii.String("payloadType"),
//   					usernameField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   				},
//   			},
//   			scopeDownStatement: statementProperty_,
//   			version: jsii.String("version"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_RateBasedStatementProperty struct {
	// Setting that indicates how to aggregate the request counts. The options are the following:.
	//
	// - IP - Aggregate the request counts on the IP address from the web request origin.
	// - FORWARDED_IP - Aggregate the request counts on the first IP address in an HTTP header. If you use this, configure the `ForwardedIPConfig` , to specify the header to use.
	AggregateKeyType *string `field:"required" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// The limit on requests per 5-minute period for a single originating IP address.
	//
	// If the statement includes a `ScopeDownStatement` , this limit is applied only to the requests that match the statement.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// This is required if `AggregateKeyType` is set to `FORWARDED_IP` .
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the rate-based statement.
	//
	// Requests are only tracked by the rate-based statement if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

// A rule statement used to search web request components for a match against a single regular expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   regexMatchStatementProperty := &regexMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	regexString: jsii.String("regexString"),
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnWebACL_RegexMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The string representing the regular expression.
	RegexString *string `field:"required" json:"regexString" yaml:"regexString"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// A rule statement used to search web request components for matches with regular expressions.
//
// To use this, create a `RegexPatternSet` that specifies the expressions that you want to detect, then use that set in this statement. A web request matches the pattern set rule statement if the request component matches any of the patterns in the set.
//
// Each regex pattern set rule statement references a regex pattern set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   regexPatternSetReferenceStatementProperty := &regexPatternSetReferenceStatementProperty{
//   	arn: jsii.String("arn"),
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnWebACL_RegexPatternSetReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the `RegexPatternSet` that this statement references.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// The action that AWS WAF should take on a web request when it matches a rule's statement.
//
// Settings at the web ACL level can override the rule action setting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleActionProperty := &ruleActionProperty{
//   	allow: &allowActionProperty{
//   		customRequestHandling: &customRequestHandlingProperty{
//   			insertHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	block: &blockActionProperty{
//   		customResponse: &customResponseProperty{
//   			responseCode: jsii.Number(123),
//
//   			// the properties below are optional
//   			customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   			responseHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	captcha: &captchaActionProperty{
//   		customRequestHandling: &customRequestHandlingProperty{
//   			insertHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	count: &countActionProperty{
//   		customRequestHandling: &customRequestHandlingProperty{
//   			insertHeaders: []interface{}{
//   				&customHTTPHeaderProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_RuleActionProperty struct {
	// Instructs AWS WAF to allow the web request.
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Instructs AWS WAF to block the web request.
	Block interface{} `field:"optional" json:"block" yaml:"block"`
	// Specifies that AWS WAF should run a `CAPTCHA` check against the request:.
	//
	// - If the request includes a valid, unexpired `CAPTCHA` token, AWS WAF allows the web request inspection to proceed to the next rule, similar to a `CountAction` .
	// - If the request doesn't include a valid, unexpired `CAPTCHA` token, AWS WAF discontinues the web ACL evaluation of the request and blocks it from going to its intended destination.
	//
	// AWS WAF generates a response that it sends back to the client, which includes the following:
	//
	// - The header `x-amzn-waf-action` with a value of `captcha` .
	// - The HTTP status code `405 Method Not Allowed` .
	// - If the request contains an `Accept` header with a value of `text/html` , the response includes a `CAPTCHA` challenge.
	//
	// You can configure the expiration time in the `CaptchaConfig` `ImmunityTimeProperty` setting at the rule and web ACL level. The rule setting overrides the web ACL setting.
	//
	// This action option is available for rules. It isn't available for web ACL default actions.
	Captcha interface{} `field:"optional" json:"captcha" yaml:"captcha"`
	// Instructs AWS WAF to count the web request and allow it.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}

// A rule statement used to run the rules that are defined in a `RuleGroup` .
//
// To use this, create a rule group with your rules, then provide the ARN of the rule group in this statement.
//
// You cannot nest a `RuleGroupReferenceStatement` , for example for use inside a `NotStatement` or `OrStatement` . You can only use a rule group reference statement at the top level inside a web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupReferenceStatementProperty := &ruleGroupReferenceStatementProperty{
//   	arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	excludedRules: []interface{}{
//   		&excludedRuleProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnWebACL_RuleGroupReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the entity.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The rules in the referenced rule group whose actions are set to `Count` .
	//
	// When you exclude a rule, AWS WAF evaluates it exactly as it would if the rule action setting were `Count` . This is a useful option for testing the rules in a rule group without modifying how they handle your web traffic.
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
}

// A single rule, which you can use in a `WebACL` or `RuleGroup` to identify web requests that you want to allow, block, or count.
//
// Each rule includes one top-level `Statement` that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var count interface{}
//   var method interface{}
//   var none interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   ruleProperty := &ruleProperty{
//   	name: jsii.String("name"),
//   	priority: jsii.Number(123),
//   	statement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   			name: jsii.String("name"),
//   			vendorName: jsii.String("vendorName"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			managedRuleGroupConfigs: []interface{}{
//   				&managedRuleGroupConfigProperty{
//   					loginPath: jsii.String("loginPath"),
//   					passwordField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   					payloadType: jsii.String("payloadType"),
//   					usernameField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   				},
//   			},
//   			scopeDownStatement: statementProperty_,
//   			version: jsii.String("version"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: singleHeader,
//   				singleQueryArgument: singleQueryArgument,
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	visibilityConfig: &visibilityConfigProperty{
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   		metricName: jsii.String("metricName"),
//   		sampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	action: &ruleActionProperty{
//   		allow: &allowActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		block: &blockActionProperty{
//   			customResponse: &customResponseProperty{
//   				responseCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   				responseHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		captcha: &captchaActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		count: &countActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	captchaConfig: &captchaConfigProperty{
//   		immunityTimeProperty: &immunityTimePropertyProperty{
//   			immunityTime: jsii.Number(123),
//   		},
//   	},
//   	overrideAction: &overrideActionProperty{
//   		count: count,
//   		none: none,
//   	},
//   	ruleLabels: []interface{}{
//   		&labelProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnWebACL_RuleProperty struct {
	// The name of the rule.
	//
	// You can't change the name of a `Rule` after you create it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// If you define more than one `Rule` in a `WebACL` , AWS WAF evaluates each request against the `Rules` in order based on the value of `Priority` .
	//
	// AWS WAF processes rules with lower priority first. The priorities don't need to be consecutive, but they must all be different.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The AWS WAF processing statement for the rule, for example `ByteMatchStatement` or `SizeConstraintStatement` .
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig interface{} `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// The action that AWS WAF should take on a web request when it matches the rule's statement.
	//
	// Settings at the web ACL level can override the rule action setting.
	//
	// This is used only for rules whose statements don't reference a rule group. Rule statements that reference a rule group are `RuleGroupReferenceStatement` and `ManagedRuleGroupStatement` .
	//
	// You must set either this `Action` setting or the rule's `OverrideAction` , but not both:
	//
	// - If the rule statement doesn't reference a rule group, you must set this rule action setting and you must not set the rule's override action setting.
	// - If the rule statement references a rule group, you must not set this action setting, because the actions are already set on the rules inside the rule group. You must set the rule's override action setting to indicate specifically whether to override the actions that are set on the rules in the rule group.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations.
	//
	// If you don't specify this, AWS WAF uses the `CAPTCHA` configuration that's defined for the web ACL.
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// The override action to apply to the rules in a rule group, instead of the individual rule action settings.
	//
	// This is used only for rules whose statements reference a rule group. Rule statements that reference a rule group are `RuleGroupReferenceStatement` and `ManagedRuleGroupStatement` .
	//
	// Set the override action to none to leave the rule group rule actions in effect. Set it to count to only count matches, regardless of the rule action settings.
	//
	// You must set either this `OverrideAction` setting or the `Action` setting, but not both:
	//
	// - If the rule statement references a rule group, you must set this override action setting and you must not set the rule's action setting.
	// - If the rule statement doesn't reference a rule group, you must set the rule action setting and you must not set the rule's override action setting.
	OverrideAction interface{} `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Labels to apply to web requests that match the rule match statement.
	//
	// AWS WAF applies fully qualified labels to matching web requests. A fully qualified label is the concatenation of a label namespace and a rule label. The rule's rule group or web ACL defines the label namespace.
	//
	// Rules that run after this rule in the web ACL can match against these labels using a `LabelMatchStatement` .
	//
	// For each label, provide a case-sensitive string containing optional namespaces and a label name, according to the following guidelines:
	//
	// - Separate each component of the label with a colon.
	// - Each namespace or name can have up to 128 characters.
	// - You can specify up to 5 namespaces in a label.
	// - Don't use the following reserved words in your label specification: `aws` , `waf` , `managed` , `rulegroup` , `webacl` , `regexpatternset` , or `ipset` .
	//
	// For example, `myLabelName` or `nameSpace1:nameSpace2:myLabelName` .
	RuleLabels interface{} `field:"optional" json:"ruleLabels" yaml:"ruleLabels"`
}

// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
//
// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
//
// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the first 8192 bytes (8 KB). If the request body for your web requests never exceeds 8192 bytes, you can create a size constraint condition and block requests that have a request body greater than 8192 bytes.
//
// If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   sizeConstraintStatementProperty := &sizeConstraintStatementProperty{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	size: jsii.Number(123),
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnWebACL_SizeConstraintStatementProperty struct {
	// The operator to use to compare the request part to the size setting.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// The size, in byte, to compare to the request part, after any transformations.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// Attackers sometimes insert malicious SQL code into web requests in an effort to extract data from your database.
//
// To allow or block web requests that appear to contain malicious SQL code, create one or more SQL injection match conditions. An SQL injection match condition identifies the part of web requests, such as the URI or the query string, that you want AWS WAF to inspect. Later in the process, when you create a web ACL, you specify whether to allow or block requests that appear to contain malicious SQL code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   sqliMatchStatementProperty := &sqliMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnWebACL_SqliMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// The processing guidance for a rule, used by AWS WAF to determine whether a web request matches the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var andStatementProperty_ andStatementProperty
//   var managedRuleGroupStatementProperty_ managedRuleGroupStatementProperty
//   var method interface{}
//   var notStatementProperty_ notStatementProperty
//   var orStatementProperty_ orStatementProperty
//   var queryString interface{}
//   var rateBasedStatementProperty_ rateBasedStatementProperty
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   statementProperty := &statementProperty{
//   	andStatement: &andStatementProperty{
//   		statements: []interface{}{
//   			&statementProperty{
//   				andStatement: andStatementProperty_,
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   					name: jsii.String("name"),
//   					vendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					managedRuleGroupConfigs: []interface{}{
//   						&managedRuleGroupConfigProperty{
//   							loginPath: jsii.String("loginPath"),
//   							passwordField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   							payloadType: jsii.String("payloadType"),
//   							usernameField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					scopeDownStatement: statementProperty_,
//   					version: jsii.String("version"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: &orStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	byteMatchStatement: &byteMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		positionalConstraint: jsii.String("positionalConstraint"),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		searchString: jsii.String("searchString"),
//   		searchStringBase64: jsii.String("searchStringBase64"),
//   	},
//   	geoMatchStatement: &geoMatchStatementProperty{
//   		countryCodes: []*string{
//   			jsii.String("countryCodes"),
//   		},
//   		forwardedIpConfig: &forwardedIPConfigurationProperty{
//   			fallbackBehavior: jsii.String("fallbackBehavior"),
//   			headerName: jsii.String("headerName"),
//   		},
//   	},
//   	ipSetReferenceStatement: map[string]interface{}{
//   		"arn": jsii.String("arn"),
//
//   		// the properties below are optional
//   		"ipSetForwardedIpConfig": map[string]*string{
//   			"fallbackBehavior": jsii.String("fallbackBehavior"),
//   			"headerName": jsii.String("headerName"),
//   			"position": jsii.String("position"),
//   		},
//   	},
//   	labelMatchStatement: &labelMatchStatementProperty{
//   		key: jsii.String("key"),
//   		scope: jsii.String("scope"),
//   	},
//   	managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   		name: jsii.String("name"),
//   		vendorName: jsii.String("vendorName"),
//
//   		// the properties below are optional
//   		excludedRules: []interface{}{
//   			&excludedRuleProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		managedRuleGroupConfigs: []interface{}{
//   			&managedRuleGroupConfigProperty{
//   				loginPath: jsii.String("loginPath"),
//   				passwordField: &fieldIdentifierProperty{
//   					identifier: jsii.String("identifier"),
//   				},
//   				payloadType: jsii.String("payloadType"),
//   				usernameField: &fieldIdentifierProperty{
//   					identifier: jsii.String("identifier"),
//   				},
//   			},
//   		},
//   		scopeDownStatement: &statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			managedRuleGroupStatement: managedRuleGroupStatementProperty_,
//   			notStatement: &notStatementProperty{
//   				statement: statementProperty_,
//   			},
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: &rateBasedStatementProperty{
//   				aggregateKeyType: jsii.String("aggregateKeyType"),
//   				limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   				scopeDownStatement: statementProperty_,
//   			},
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   		version: jsii.String("version"),
//   	},
//   	notStatement: &notStatementProperty{
//   		statement: &statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   				name: jsii.String("name"),
//   				vendorName: jsii.String("vendorName"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				managedRuleGroupConfigs: []interface{}{
//   					&managedRuleGroupConfigProperty{
//   						loginPath: jsii.String("loginPath"),
//   						passwordField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   						payloadType: jsii.String("payloadType"),
//   						usernameField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   					},
//   				},
//   				scopeDownStatement: statementProperty_,
//   				version: jsii.String("version"),
//   			},
//   			notStatement: notStatementProperty_,
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: &rateBasedStatementProperty{
//   				aggregateKeyType: jsii.String("aggregateKeyType"),
//   				limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   				scopeDownStatement: statementProperty_,
//   			},
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	orStatement: &orStatementProperty{
//   		statements: []interface{}{
//   			&statementProperty{
//   				andStatement: &andStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   					name: jsii.String("name"),
//   					vendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					managedRuleGroupConfigs: []interface{}{
//   						&managedRuleGroupConfigProperty{
//   							loginPath: jsii.String("loginPath"),
//   							passwordField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   							payloadType: jsii.String("payloadType"),
//   							usernameField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					scopeDownStatement: statementProperty_,
//   					version: jsii.String("version"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: orStatementProperty_,
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	rateBasedStatement: &rateBasedStatementProperty{
//   		aggregateKeyType: jsii.String("aggregateKeyType"),
//   		limit: jsii.Number(123),
//
//   		// the properties below are optional
//   		forwardedIpConfig: &forwardedIPConfigurationProperty{
//   			fallbackBehavior: jsii.String("fallbackBehavior"),
//   			headerName: jsii.String("headerName"),
//   		},
//   		scopeDownStatement: &statementProperty{
//   			andStatement: &andStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			byteMatchStatement: &byteMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				positionalConstraint: jsii.String("positionalConstraint"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				searchString: jsii.String("searchString"),
//   				searchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			geoMatchStatement: &geoMatchStatementProperty{
//   				countryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				forwardedIpConfig: &forwardedIPConfigurationProperty{
//   					fallbackBehavior: jsii.String("fallbackBehavior"),
//   					headerName: jsii.String("headerName"),
//   				},
//   			},
//   			ipSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			labelMatchStatement: &labelMatchStatementProperty{
//   				key: jsii.String("key"),
//   				scope: jsii.String("scope"),
//   			},
//   			managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   				name: jsii.String("name"),
//   				vendorName: jsii.String("vendorName"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				managedRuleGroupConfigs: []interface{}{
//   					&managedRuleGroupConfigProperty{
//   						loginPath: jsii.String("loginPath"),
//   						passwordField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   						payloadType: jsii.String("payloadType"),
//   						usernameField: &fieldIdentifierProperty{
//   							identifier: jsii.String("identifier"),
//   						},
//   					},
//   				},
//   				scopeDownStatement: statementProperty_,
//   				version: jsii.String("version"),
//   			},
//   			notStatement: &notStatementProperty{
//   				statement: statementProperty_,
//   			},
//   			orStatement: &orStatementProperty{
//   				statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			rateBasedStatement: rateBasedStatementProperty_,
//   			regexMatchStatement: &regexMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				regexString: jsii.String("regexString"),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   				arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				excludedRules: []interface{}{
//   					&excludedRuleProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			sizeConstraintStatement: &sizeConstraintStatementProperty{
//   				comparisonOperator: jsii.String("comparisonOperator"),
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				size: jsii.Number(123),
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			sqliMatchStatement: &sqliMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			xssMatchStatement: &xssMatchStatementProperty{
//   				fieldToMatch: &fieldToMatchProperty{
//   					allQueryArguments: allQueryArguments,
//   					body: &bodyProperty{
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					cookies: &cookiesProperty{
//   						matchPattern: &cookieMatchPatternProperty{
//   							all: all,
//   							excludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							includedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					headers: &headersProperty{
//   						matchPattern: &headerMatchPatternProperty{
//   							all: all,
//   							excludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							includedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					jsonBody: &jsonBodyProperty{
//   						matchPattern: &jsonMatchPatternProperty{
//   							all: all,
//   							includedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						matchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						oversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					method: method,
//   					queryString: queryString,
//   					singleHeader: singleHeader,
//   					singleQueryArgument: singleQueryArgument,
//   					uriPath: uriPath,
//   				},
//   				textTransformations: []interface{}{
//   					&textTransformationProperty{
//   						priority: jsii.Number(123),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	regexMatchStatement: &regexMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		regexString: jsii.String("regexString"),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   		arn: jsii.String("arn"),
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   		arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		excludedRules: []interface{}{
//   			&excludedRuleProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	sizeConstraintStatement: &sizeConstraintStatementProperty{
//   		comparisonOperator: jsii.String("comparisonOperator"),
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		size: jsii.Number(123),
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	sqliMatchStatement: &sqliMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	xssMatchStatement: &xssMatchStatementProperty{
//   		fieldToMatch: &fieldToMatchProperty{
//   			allQueryArguments: allQueryArguments,
//   			body: &bodyProperty{
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			cookies: &cookiesProperty{
//   				matchPattern: &cookieMatchPatternProperty{
//   					all: all,
//   					excludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					includedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			headers: &headersProperty{
//   				matchPattern: &headerMatchPatternProperty{
//   					all: all,
//   					excludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					includedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &jsonMatchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				oversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: singleHeader,
//   			singleQueryArgument: singleQueryArgument,
//   			uriPath: uriPath,
//   		},
//   		textTransformations: []interface{}{
//   			&textTransformationProperty{
//   				priority: jsii.Number(123),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_StatementProperty struct {
	// A logical rule statement used to combine other rule statements with AND logic.
	//
	// You provide more than one `Statement` within the `AndStatement` .
	AndStatement interface{} `field:"optional" json:"andStatement" yaml:"andStatement"`
	// A rule statement that defines a string match search for AWS WAF to apply to web requests.
	//
	// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is refered to as a string match statement.
	ByteMatchStatement interface{} `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// A rule statement used to identify web requests based on country of origin.
	GeoMatchStatement interface{} `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
	//
	// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
	//
	// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	IpSetReferenceStatement interface{} `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// A rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL.
	//
	// The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.
	LabelMatchStatement interface{} `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// A rule statement used to run the rules that are defined in a managed rule group.
	//
	// To use this, provide the vendor name and the name of the rule group in this statement.
	//
	// You cannot nest a `ManagedRuleGroupStatement` , for example for use inside a `NotStatement` or `OrStatement` . It can only be referenced as a top-level statement within a rule.
	ManagedRuleGroupStatement interface{} `field:"optional" json:"managedRuleGroupStatement" yaml:"managedRuleGroupStatement"`
	// A logical rule statement used to negate the results of another rule statement.
	//
	// You provide one `Statement` within the `NotStatement` .
	NotStatement interface{} `field:"optional" json:"notStatement" yaml:"notStatement"`
	// A logical rule statement used to combine other rule statements with OR logic.
	//
	// You provide more than one `Statement` within the `OrStatement` .
	OrStatement interface{} `field:"optional" json:"orStatement" yaml:"orStatement"`
	// A rate-based rule tracks the rate of requests for each originating IP address, and triggers the rule action when the rate exceeds a limit that you specify on the number of requests in any 5-minute time span.
	//
	// You can use this to put a temporary block on requests from an IP address that is sending excessive requests.
	//
	// AWS WAF tracks and manages web requests separately for each instance of a rate-based rule that you use. For example, if you provide the same rate-based rule settings in two web ACLs, each of the two rule statements represents a separate instance of the rate-based rule and gets its own tracking and management by AWS WAF . If you define a rate-based rule inside a rule group, and then use that rule group in multiple places, each use creates a separate instance of the rate-based rule that gets its own tracking and management by AWS WAF .
	//
	// When the rule action triggers, AWS WAF blocks additional requests from the IP address until the request rate falls below the limit.
	//
	// You can optionally nest another statement inside the rate-based statement, to narrow the scope of the rule so that it only counts requests that match the nested statement. For example, based on recent requests that you have seen from an attacker, you might create a rate-based rule with a nested AND rule statement that contains the following nested statements:
	//
	// - An IP match statement with an IP set that specified the address 192.0.2.44.
	// - A string match statement that searches in the User-Agent header for the string BadBot.
	//
	// In this rate-based rule, you also define a rate limit. For this example, the rate limit is 1,000. Requests that meet both of the conditions in the statements are counted. If the count exceeds 1,000 requests per five minutes, the rule action triggers. Requests that do not meet both conditions are not counted towards the rate limit and are not affected by this rule.
	//
	// You cannot nest a `RateBasedStatement` inside another statement, for example inside a `NotStatement` or `OrStatement` . You can define a `RateBasedStatement` inside a web ACL and inside a rule group.
	RateBasedStatement interface{} `field:"optional" json:"rateBasedStatement" yaml:"rateBasedStatement"`
	// A rule statement used to search web request components for a match against a single regular expression.
	RegexMatchStatement interface{} `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// A rule statement used to search web request components for matches with regular expressions.
	//
	// To use this, create a `RegexPatternSet` that specifies the expressions that you want to detect, then use the ARN of that set in this statement. A web request matches the pattern set rule statement if the request component matches any of the patterns in the set.
	//
	// Each regex pattern set rule statement references a regex pattern set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	RegexPatternSetReferenceStatement interface{} `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// A rule statement used to run the rules that are defined in a `RuleGroup` .
	//
	// To use this, create a rule group with your rules, then provide the ARN of the rule group in this statement.
	//
	// You cannot nest a `RuleGroupReferenceStatement` , for example for use inside a `NotStatement` or `OrStatement` . You can only use a rule group reference statement at the top level inside a web ACL.
	RuleGroupReferenceStatement interface{} `field:"optional" json:"ruleGroupReferenceStatement" yaml:"ruleGroupReferenceStatement"`
	// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
	//
	// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
	//
	// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the first 8192 bytes (8 KB). If the request body for your web requests never exceeds 8192 bytes, you can create a size constraint condition and block requests that have a request body greater than 8192 bytes.
	//
	// If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.
	SizeConstraintStatement interface{} `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// Attackers sometimes insert malicious SQL code into web requests in an effort to extract data from your database.
	//
	// To allow or block web requests that appear to contain malicious SQL code, create one or more SQL injection match conditions. An SQL injection match condition identifies the part of web requests, such as the URI or the query string, that you want AWS WAF to inspect. Later in the process, when you create a web ACL, you specify whether to allow or block requests that appear to contain malicious SQL code.
	SqliMatchStatement interface{} `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// A rule statement that defines a cross-site scripting (XSS) match search for AWS WAF to apply to web requests.
	//
	// XSS attacks are those where the attacker uses vulnerabilities in a benign website as a vehicle to inject malicious client-site scripts into other legitimate web browsers. The XSS match statement provides the location in requests that you want AWS WAF to search and text transformations to use on the search area before AWS WAF searches for character sequences that are likely to be malicious strings.
	XssMatchStatement interface{} `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textTransformationProperty := &textTransformationProperty{
//   	priority: jsii.Number(123),
//   	type: jsii.String("type"),
//   }
//
type CfnWebACL_TextTransformationProperty struct {
	// Sets the relative processing order for multiple transformations that are defined for a rule statement.
	//
	// AWS WAF processes all transformations, from lowest priority to highest, before inspecting the transformed content. The priorities don't need to be consecutive, but they must all be different.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// You can specify the following transformation types:.
	//
	// *BASE64_DECODE* - Decode a `Base64` -encoded string.
	//
	// *BASE64_DECODE_EXT* - Decode a `Base64` -encoded string, but use a forgiving implementation that ignores characters that aren't valid.
	//
	// *CMD_LINE* - Command-line transformations. These are helpful in reducing effectiveness of attackers who inject an operating system command-line command and use unusual formatting to disguise some or all of the command.
	//
	// - Delete the following characters: `\ " ' ^`
	// - Delete spaces before the following characters: `/ (`
	// - Replace the following characters with a space: `, ;`
	// - Replace multiple spaces with one space
	// - Convert uppercase letters (A-Z) to lowercase (a-z)
	//
	// *COMPRESS_WHITE_SPACE* - Replace these characters with a space character (decimal 32):
	//
	// - `\f` , formfeed, decimal 12
	// - `\t` , tab, decimal 9
	// - `\n` , newline, decimal 10
	// - `\r` , carriage return, decimal 13
	// - `\v` , vertical tab, decimal 11
	// - Non-breaking space, decimal 160
	//
	// `COMPRESS_WHITE_SPACE` also replaces multiple spaces with one space.
	//
	// *CSS_DECODE* - Decode characters that were encoded using CSS 2.x escape rules `syndata.html#characters` . This function uses up to two bytes in the decoding process, so it can help to uncover ASCII characters that were encoded using CSS encoding that wouldn’t typically be encoded. It's also useful in countering evasion, which is a combination of a backslash and non-hexadecimal characters. For example, `ja\vascript` for javascript.
	//
	// *ESCAPE_SEQ_DECODE* - Decode the following ANSI C escape sequences: `\a` , `\b` , `\f` , `\n` , `\r` , `\t` , `\v` , `\\` , `\?` , `\'` , `\"` , `\xHH` (hexadecimal), `\0OOO` (octal). Encodings that aren't valid remain in the output.
	//
	// *HEX_DECODE* - Decode a string of hexadecimal characters into a binary.
	//
	// *HTML_ENTITY_DECODE* - Replace HTML-encoded characters with unencoded characters. `HTML_ENTITY_DECODE` performs these operations:
	//
	// - Replaces `(ampersand)quot;` with `"`
	// - Replaces `(ampersand)nbsp;` with a non-breaking space, decimal 160
	// - Replaces `(ampersand)lt;` with a "less than" symbol
	// - Replaces `(ampersand)gt;` with `>`
	// - Replaces characters that are represented in hexadecimal format, `(ampersand)#xhhhh;` , with the corresponding characters
	// - Replaces characters that are represented in decimal format, `(ampersand)#nnnn;` , with the corresponding characters
	//
	// *JS_DECODE* - Decode JavaScript escape sequences. If a `\` `u` `HHHH` code is in the full-width ASCII code range of `FF01-FF5E` , then the higher byte is used to detect and adjust the lower byte. If not, only the lower byte is used and the higher byte is zeroed, causing a possible loss of information.
	//
	// *LOWERCASE* - Convert uppercase letters (A-Z) to lowercase (a-z).
	//
	// *MD5* - Calculate an MD5 hash from the data in the input. The computed hash is in a raw binary form.
	//
	// *NONE* - Specify `NONE` if you don't want any text transformations.
	//
	// *NORMALIZE_PATH* - Remove multiple slashes, directory self-references, and directory back-references that are not at the beginning of the input from an input string.
	//
	// *NORMALIZE_PATH_WIN* - This is the same as `NORMALIZE_PATH` , but first converts backslash characters to forward slashes.
	//
	// *REMOVE_NULLS* - Remove all `NULL` bytes from the input.
	//
	// *REPLACE_COMMENTS* - Replace each occurrence of a C-style comment ( `/* ... * /` ) with a single space. Multiple consecutive occurrences are not compressed. Unterminated comments are also replaced with a space (ASCII 0x20). However, a standalone termination of a comment ( `* /` ) is not acted upon.
	//
	// *REPLACE_NULLS* - Replace NULL bytes in the input with space characters (ASCII `0x20` ).
	//
	// *SQL_HEX_DECODE* - Decode SQL hex data. Example ( `0x414243` ) will be decoded to ( `ABC` ).
	//
	// *URL_DECODE* - Decode a URL-encoded value.
	//
	// *URL_DECODE_UNI* - Like `URL_DECODE` , but with support for Microsoft-specific `%u` encoding. If the code is in the full-width ASCII code range of `FF01-FF5E` , the higher byte is used to detect and adjust the lower byte. Otherwise, only the lower byte is used and the higher byte is zeroed.
	//
	// *UTF8_TO_UNICODE* - Convert all UTF-8 character sequences to Unicode. This helps input normalization, and minimizing false-positives and false-negatives for non-English languages.
	Type *string `field:"required" json:"type" yaml:"type"`
}

// Defines and enables Amazon CloudWatch metrics and web request sample collection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visibilityConfigProperty := &visibilityConfigProperty{
//   	cloudWatchMetricsEnabled: jsii.Boolean(false),
//   	metricName: jsii.String("metricName"),
//   	sampledRequestsEnabled: jsii.Boolean(false),
//   }
//
type CfnWebACL_VisibilityConfigProperty struct {
	// A boolean indicating whether the associated resource sends metrics to Amazon CloudWatch.
	//
	// For the list of available metrics, see [AWS WAF Metrics](https://docs.aws.amazon.com/waf/latest/developerguide/monitoring-cloudwatch.html#waf-metrics) .
	CloudWatchMetricsEnabled interface{} `field:"required" json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// A name of the Amazon CloudWatch metric.
	//
	// The name can contain only the characters: A-Z, a-z, 0-9, - (hyphen), and _ (underscore). The name can be from one to 128 characters long. It can't contain whitespace or metric names reserved for AWS WAF , for example "All" and "Default_Action."
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules.
	//
	// You can view the sampled requests through the AWS WAF console.
	SampledRequestsEnabled interface{} `field:"required" json:"sampledRequestsEnabled" yaml:"sampledRequestsEnabled"`
}

// A rule statement that defines a cross-site scripting (XSS) match search for AWS WAF to apply to web requests.
//
// XSS attacks are those where the attacker uses vulnerabilities in a benign website as a vehicle to inject malicious client-site scripts into other legitimate web browsers. The XSS match statement provides the location in requests that you want AWS WAF to search and text transformations to use on the search area before AWS WAF searches for character sequences that are likely to be malicious strings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   xssMatchStatementProperty := &xssMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnWebACL_XssMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
}

// A CloudFormation `AWS::WAFv2::WebACLAssociation`.
//
// > This is the latest version of *AWS WAF* , named AWS WAF V2, released in November, 2019. For information, including how to migrate your AWS WAF resources from the prior release, see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Use a web ACL association to define an association between a web ACL and a regional application resource, to protect the resource. A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API.
//
// For Amazon CloudFront , don't use this resource. Instead, use your CloudFront distribution configuration. To associate a web ACL with a distribution, provide the Amazon Resource Name (ARN) of the `WebACL` to your CloudFront distribution configuration. To disassociate a web ACL, provide an empty ARN. For information, see [AWS::CloudFront::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-distribution.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebACLAssociation := awscdk.Aws_wafv2.NewCfnWebACLAssociation(this, jsii.String("MyCfnWebACLAssociation"), &cfnWebACLAssociationProps{
//   	resourceArn: jsii.String("resourceArn"),
//   	webAclArn: jsii.String("webAclArn"),
//   })
//
type CfnWebACLAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL.
	//
	// The ARN must be in one of the following formats:
	//
	// - For an Application Load Balancer: `arn:aws:elasticloadbalancing: *region* : *account-id* :loadbalancer/app/ *load-balancer-name* / *load-balancer-id*`
	// - For an Amazon API Gateway REST API: `arn:aws:apigateway: *region* ::/restapis/ *api-id* /stages/ *stage-name*`
	// - For an AWS AppSync GraphQL API: `arn:aws:appsync: *region* : *account-id* :apis/ *GraphQLApiId*`.
	ResourceArn() *string
	SetResourceArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with the resource.
	WebAclArn() *string
	SetWebAclArn(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnWebACLAssociation
type jsiiProxy_CfnWebACLAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWebACLAssociation) WebAclArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclArn",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::WebACLAssociation`.
func NewCfnWebACLAssociation(scope awscdk.Construct, id *string, props *CfnWebACLAssociationProps) CfnWebACLAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACLAssociation{}

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnWebACLAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::WebACLAssociation`.
func NewCfnWebACLAssociation_Override(c CfnWebACLAssociation, scope awscdk.Construct, id *string, props *CfnWebACLAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_wafv2.CfnWebACLAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWebACLAssociation) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnWebACLAssociation) SetWebAclArn(val *string) {
	_jsii_.Set(
		j,
		"webAclArn",
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
func CfnWebACLAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnWebACLAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWebACLAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnWebACLAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnWebACLAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_wafv2.CfnWebACLAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWebACLAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_wafv2.CfnWebACLAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWebACLAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWebACLAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnWebACLAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebACLAssociationProps := &cfnWebACLAssociationProps{
//   	resourceArn: jsii.String("resourceArn"),
//   	webAclArn: jsii.String("webAclArn"),
//   }
//
type CfnWebACLAssociationProps struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL.
	//
	// The ARN must be in one of the following formats:
	//
	// - For an Application Load Balancer: `arn:aws:elasticloadbalancing: *region* : *account-id* :loadbalancer/app/ *load-balancer-name* / *load-balancer-id*`
	// - For an Amazon API Gateway REST API: `arn:aws:apigateway: *region* ::/restapis/ *api-id* /stages/ *stage-name*`
	// - For an AWS AppSync GraphQL API: `arn:aws:appsync: *region* : *account-id* :apis/ *GraphQLApiId*`.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with the resource.
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
}

// Properties for defining a `CfnWebACL`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var count interface{}
//   var method interface{}
//   var none interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   cfnWebACLProps := &cfnWebACLProps{
//   	defaultAction: &defaultActionProperty{
//   		allow: &allowActionProperty{
//   			customRequestHandling: &customRequestHandlingProperty{
//   				insertHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		block: &blockActionProperty{
//   			customResponse: &customResponseProperty{
//   				responseCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   				responseHeaders: []interface{}{
//   					&customHTTPHeaderProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	scope: jsii.String("scope"),
//   	visibilityConfig: &visibilityConfigProperty{
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   		metricName: jsii.String("metricName"),
//   		sampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	captchaConfig: &captchaConfigProperty{
//   		immunityTimeProperty: &immunityTimePropertyProperty{
//   			immunityTime: jsii.Number(123),
//   		},
//   	},
//   	customResponseBodies: map[string]interface{}{
//   		"customResponseBodiesKey": &CustomResponseBodyProperty{
//   			"content": jsii.String("content"),
//   			"contentType": jsii.String("contentType"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			name: jsii.String("name"),
//   			priority: jsii.Number(123),
//   			statement: &statementProperty{
//   				andStatement: &andStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   					name: jsii.String("name"),
//   					vendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					managedRuleGroupConfigs: []interface{}{
//   						&managedRuleGroupConfigProperty{
//   							loginPath: jsii.String("loginPath"),
//   							passwordField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   							payloadType: jsii.String("payloadType"),
//   							usernameField: &fieldIdentifierProperty{
//   								identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					scopeDownStatement: statementProperty_,
//   					version: jsii.String("version"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: &orStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					excludedRules: []interface{}{
//   						&excludedRuleProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   			visibilityConfig: &visibilityConfigProperty{
//   				cloudWatchMetricsEnabled: jsii.Boolean(false),
//   				metricName: jsii.String("metricName"),
//   				sampledRequestsEnabled: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			action: &ruleActionProperty{
//   				allow: &allowActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				block: &blockActionProperty{
//   					customResponse: &customResponseProperty{
//   						responseCode: jsii.Number(123),
//
//   						// the properties below are optional
//   						customResponseBodyKey: jsii.String("customResponseBodyKey"),
//   						responseHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				captcha: &captchaActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				count: &countActionProperty{
//   					customRequestHandling: &customRequestHandlingProperty{
//   						insertHeaders: []interface{}{
//   							&customHTTPHeaderProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			captchaConfig: &captchaConfigProperty{
//   				immunityTimeProperty: &immunityTimePropertyProperty{
//   					immunityTime: jsii.Number(123),
//   				},
//   			},
//   			overrideAction: &overrideActionProperty{
//   				count: count,
//   				none: none,
//   			},
//   			ruleLabels: []interface{}{
//   				&labelProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWebACLProps struct {
	// The action to perform if none of the `Rules` contained in the `WebACL` match.
	DefaultAction interface{} `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	//
	// For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig interface{} `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings.
	//
	// If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
	CaptchaConfig interface{} `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// A map of custom response keys and content bodies.
	//
	// When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomResponseBodies interface{} `field:"optional" json:"customResponseBodies" yaml:"customResponseBodies"`
	// A description of the web ACL that helps with identification.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the web ACL.
	//
	// You cannot change the name of a web ACL after you create it.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rule statements used to identify the web requests that you want to allow, block, or count.
	//
	// Each rule includes one top-level statement that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

