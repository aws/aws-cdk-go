package awsroute53resolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Route53Resolver::FirewallDomainList`.
//
// High-level information about a list of firewall domains for use in a [AWS::Route53Resolver::FirewallRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-rule.html) . This is returned by [GetFirewallDomainList](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetFirewallDomainList.html) .
//
// To retrieve the domains that are defined for this domain list, call [ListFirewallDomains](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListFirewallDomains.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallDomainList := awscdk.Aws_route53resolver.NewCfnFirewallDomainList(this, jsii.String("MyCfnFirewallDomainList"), &cfnFirewallDomainListProps{
//   	domainFileUrl: jsii.String("domainFileUrl"),
//   	domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFirewallDomainList interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the firewall domain list.
	AttrArn() *string
	// The date and time that the domain list was created, in Unix time format and Coordinated Universal Time (UTC).
	AttrCreationTime() *string
	// A unique string defined by you to identify the request.
	//
	// This allows you to retry failed requests without the risk of running the operation twice. This can be any unique string, for example, a timestamp.
	AttrCreatorRequestId() *string
	// The number of domain names that are specified in the domain list.
	AttrDomainCount() *float64
	// The ID of the domain list.
	AttrId() *string
	// The owner of the list, used only for lists that are not managed by you.
	//
	// For example, the managed domain list `AWSManagedDomainsMalwareDomainList` has the managed owner name `Route 53 Resolver DNS Firewall` .
	AttrManagedOwnerName() *string
	// The date and time that the domain list was last modified, in Unix time format and Coordinated Universal Time (UTC).
	AttrModificationTime() *string
	// The status of the domain list.
	AttrStatus() *string
	// Additional information about the status of the list, if available.
	AttrStatusMessage() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The fully qualified URL or URI of the file stored in Amazon Simple Storage Service (Amazon S3) that contains the list of domains to import.
	//
	// The file must be in an S3 bucket that's in the same Region as your DNS Firewall. The file must be a text file and must contain a single domain per line.
	DomainFileUrl() *string
	SetDomainFileUrl(val *string)
	// A list of the domain lists that you have defined.
	Domains() *[]*string
	SetDomains(val *[]*string)
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
	// The name of the domain list.
	Name() *string
	SetName(val *string)
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
	// A list of the tag keys and values that you want to associate with the domain list.
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

// The jsii proxy struct for CfnFirewallDomainList
type jsiiProxy_CfnFirewallDomainList struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrDomainCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrDomainCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrManagedOwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrManagedOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) AttrStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) DomainFileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) Domains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallDomainList) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::FirewallDomainList`.
func NewCfnFirewallDomainList(scope constructs.Construct, id *string, props *CfnFirewallDomainListProps) CfnFirewallDomainList {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallDomainList{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallDomainList",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::FirewallDomainList`.
func NewCfnFirewallDomainList_Override(c CfnFirewallDomainList, scope constructs.Construct, id *string, props *CfnFirewallDomainListProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallDomainList",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFirewallDomainList) SetDomainFileUrl(val *string) {
	_jsii_.Set(
		j,
		"domainFileUrl",
		val,
	)
}

func (j *jsiiProxy_CfnFirewallDomainList) SetDomains(val *[]*string) {
	_jsii_.Set(
		j,
		"domains",
		val,
	)
}

func (j *jsiiProxy_CfnFirewallDomainList) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFirewallDomainList_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallDomainList",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFirewallDomainList_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallDomainList",
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
func CfnFirewallDomainList_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallDomainList",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewallDomainList_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallDomainList",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewallDomainList) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallDomainList) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallDomainList) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallDomainList) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallDomainList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallDomainList) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnFirewallDomainList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallDomainListProps := &cfnFirewallDomainListProps{
//   	domainFileUrl: jsii.String("domainFileUrl"),
//   	domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFirewallDomainListProps struct {
	// The fully qualified URL or URI of the file stored in Amazon Simple Storage Service (Amazon S3) that contains the list of domains to import.
	//
	// The file must be in an S3 bucket that's in the same Region as your DNS Firewall. The file must be a text file and must contain a single domain per line.
	DomainFileUrl *string `field:"optional" json:"domainFileUrl" yaml:"domainFileUrl"`
	// A list of the domain lists that you have defined.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// The name of the domain list.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the domain list.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Route53Resolver::FirewallRuleGroup`.
//
// High-level information for a firewall rule group. A firewall rule group is a collection of rules that DNS Firewall uses to filter DNS network traffic for a VPC. To retrieve the rules for the rule group, call [ListFirewallRules](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListFirewallRules.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallRuleGroup := awscdk.Aws_route53resolver.NewCfnFirewallRuleGroup(this, jsii.String("MyCfnFirewallRuleGroup"), &cfnFirewallRuleGroupProps{
//   	firewallRules: []interface{}{
//   		&firewallRuleProperty{
//   			action: jsii.String("action"),
//   			firewallDomainListId: jsii.String("firewallDomainListId"),
//   			priority: jsii.Number(123),
//
//   			// the properties below are optional
//   			blockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   			blockOverrideDomain: jsii.String("blockOverrideDomain"),
//   			blockOverrideTtl: jsii.Number(123),
//   			blockResponse: jsii.String("blockResponse"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFirewallRuleGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN (Amazon Resource Name) of the rule group.
	AttrArn() *string
	// The date and time that the rule group was created, in Unix time format and Coordinated Universal Time (UTC).
	AttrCreationTime() *string
	// A unique string defined by you to identify the request.
	//
	// This allows you to retry failed requests without the risk of running the operation twice. This can be any unique string, for example, a timestamp.
	AttrCreatorRequestId() *string
	// The ID of the rule group.
	AttrId() *string
	// The date and time that the rule group was last modified, in Unix time format and Coordinated Universal Time (UTC).
	AttrModificationTime() *string
	// The AWS account ID for the account that created the rule group.
	//
	// When a rule group is shared with your account, this is the account that has shared the rule group with you.
	AttrOwnerId() *string
	// The number of rules in the rule group.
	AttrRuleCount() *float64
	// Whether the rule group is shared with other AWS accounts , or was shared with the current account by another AWS account .
	//
	// Sharing is configured through AWS Resource Access Manager ( AWS RAM ).
	AttrShareStatus() *string
	// The status of the domain list.
	AttrStatus() *string
	// Additional information about the status of the rule group, if available.
	AttrStatusMessage() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A list of the rules that you have defined.
	FirewallRules() interface{}
	SetFirewallRules(val interface{})
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
	// The name of the rule group.
	Name() *string
	SetName(val *string)
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
	// A list of the tag keys and values that you want to associate with the rule group.
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

// The jsii proxy struct for CfnFirewallRuleGroup
type jsiiProxy_CfnFirewallRuleGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrRuleCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrRuleCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrShareStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrShareStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) AttrStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) FirewallRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::FirewallRuleGroup`.
func NewCfnFirewallRuleGroup(scope constructs.Construct, id *string, props *CfnFirewallRuleGroupProps) CfnFirewallRuleGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallRuleGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::FirewallRuleGroup`.
func NewCfnFirewallRuleGroup_Override(c CfnFirewallRuleGroup, scope constructs.Construct, id *string, props *CfnFirewallRuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFirewallRuleGroup) SetFirewallRules(val interface{}) {
	_jsii_.Set(
		j,
		"firewallRules",
		val,
	)
}

func (j *jsiiProxy_CfnFirewallRuleGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFirewallRuleGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFirewallRuleGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroup",
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
func CfnFirewallRuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewallRuleGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A single firewall rule in a rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallRuleProperty := &firewallRuleProperty{
//   	action: jsii.String("action"),
//   	firewallDomainListId: jsii.String("firewallDomainListId"),
//   	priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	blockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   	blockOverrideDomain: jsii.String("blockOverrideDomain"),
//   	blockOverrideTtl: jsii.Number(123),
//   	blockResponse: jsii.String("blockResponse"),
//   }
//
type CfnFirewallRuleGroup_FirewallRuleProperty struct {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list:  - `ALLOW` - Permit the request to go through.
	//
	// - `ALERT` - Permit the request to go through but send an alert to the logs.
	// - `BLOCK` - Disallow the request. If this is specified,then `BlockResponse` must also be specified.
	//
	// if `BlockResponse` is `OVERRIDE` , then all of the following `OVERRIDE` attributes must be specified:
	//
	// - `BlockOverrideDnsType`
	// - `BlockOverrideDomain`
	// - `BlockOverrideTtl`.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The ID of the domain list that's used in the rule.
	FirewallDomainListId *string `field:"required" json:"firewallDomainListId" yaml:"firewallDomainListId"`
	// The priority of the rule in the rule group.
	//
	// This value must be unique within the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The DNS record's type.
	//
	// This determines the format of the record value that you provided in `BlockOverrideDomain` . Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	BlockOverrideDnsType *string `field:"optional" json:"blockOverrideDnsType" yaml:"blockOverrideDnsType"`
	// The custom DNS record to send back in response to the query.
	//
	// Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	BlockOverrideDomain *string `field:"optional" json:"blockOverrideDomain" yaml:"blockOverrideDomain"`
	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.
	//
	// Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	BlockOverrideTtl *float64 `field:"optional" json:"blockOverrideTtl" yaml:"blockOverrideTtl"`
	// The way that you want DNS Firewall to block the request. Used for the rule action setting `BLOCK` .
	//
	// - `NODATA` - Respond indicating that the query was successful, but no response is available for it.
	// - `NXDOMAIN` - Respond indicating that the domain name that's in the query doesn't exist.
	// - `OVERRIDE` - Provide a custom override in the response. This option requires custom handling details in the rule's `BlockOverride*` settings.
	BlockResponse *string `field:"optional" json:"blockResponse" yaml:"blockResponse"`
}

// A CloudFormation `AWS::Route53Resolver::FirewallRuleGroupAssociation`.
//
// An association between a firewall rule group and a VPC, which enables DNS filtering for the VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallRuleGroupAssociation := awscdk.Aws_route53resolver.NewCfnFirewallRuleGroupAssociation(this, jsii.String("MyCfnFirewallRuleGroupAssociation"), &cfnFirewallRuleGroupAssociationProps{
//   	firewallRuleGroupId: jsii.String("firewallRuleGroupId"),
//   	priority: jsii.Number(123),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	mutationProtection: jsii.String("mutationProtection"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFirewallRuleGroupAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the firewall rule group association.
	AttrArn() *string
	// The date and time that the association was created, in Unix time format and Coordinated Universal Time (UTC).
	AttrCreationTime() *string
	// A unique string defined by you to identify the request.
	//
	// This allows you to retry failed requests without the risk of running the operation twice. This can be any unique string, for example, a timestamp.
	AttrCreatorRequestId() *string
	// The identifier for the association.
	AttrId() *string
	// The owner of the association, used only for associations that are not managed by you.
	//
	// If you use AWS Firewall Manager to manage your firewallls from DNS Firewall, then this reports Firewall Manager as the managed owner.
	AttrManagedOwnerName() *string
	// The date and time that the association was last modified, in Unix time format and Coordinated Universal Time (UTC).
	AttrModificationTime() *string
	// The current status of the association.
	AttrStatus() *string
	// Additional information about the status of the response, if available.
	AttrStatusMessage() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The unique identifier of the firewall rule group.
	FirewallRuleGroupId() *string
	SetFirewallRuleGroupId(val *string)
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
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	MutationProtection() *string
	SetMutationProtection(val *string)
	// The name of the association.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC traffic starting from rule group with the lowest numeric priority setting.
	//
	// You must specify a unique priority for each rule group that you associate with a single VPC. To make it easier to insert rule groups later, leave space between the numbers, for example, use 101, 200, and so on. You can change the priority setting for a rule group association after you create it.
	//
	// The allowed values for `Priority` are between 100 and 9900 (excluding 100 and 9900).
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
	// A list of the tag keys and values that you want to associate with the rule group.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The unique identifier of the VPC that is associated with the rule group.
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

// The jsii proxy struct for CfnFirewallRuleGroupAssociation
type jsiiProxy_CfnFirewallRuleGroupAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) AttrCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) AttrManagedOwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrManagedOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) AttrModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) AttrStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) MutationProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mutationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::FirewallRuleGroupAssociation`.
func NewCfnFirewallRuleGroupAssociation(scope constructs.Construct, id *string, props *CfnFirewallRuleGroupAssociationProps) CfnFirewallRuleGroupAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallRuleGroupAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroupAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::FirewallRuleGroupAssociation`.
func NewCfnFirewallRuleGroupAssociation_Override(c CfnFirewallRuleGroupAssociation, scope constructs.Construct, id *string, props *CfnFirewallRuleGroupAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroupAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) SetFirewallRuleGroupId(val *string) {
	_jsii_.Set(
		j,
		"firewallRuleGroupId",
		val,
	)
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) SetMutationProtection(val *string) {
	_jsii_.Set(
		j,
		"mutationProtection",
		val,
	)
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) SetVpcId(val *string) {
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
func CfnFirewallRuleGroupAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroupAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFirewallRuleGroupAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroupAssociation",
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
func CfnFirewallRuleGroupAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroupAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewallRuleGroupAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnFirewallRuleGroupAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnFirewallRuleGroupAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallRuleGroupAssociationProps := &cfnFirewallRuleGroupAssociationProps{
//   	firewallRuleGroupId: jsii.String("firewallRuleGroupId"),
//   	priority: jsii.Number(123),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	mutationProtection: jsii.String("mutationProtection"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFirewallRuleGroupAssociationProps struct {
	// The unique identifier of the firewall rule group.
	FirewallRuleGroupId *string `field:"required" json:"firewallRuleGroupId" yaml:"firewallRuleGroupId"`
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC traffic starting from rule group with the lowest numeric priority setting.
	//
	// You must specify a unique priority for each rule group that you associate with a single VPC. To make it easier to insert rule groups later, leave space between the numbers, for example, use 101, 200, and so on. You can change the priority setting for a rule group association after you create it.
	//
	// The allowed values for `Priority` are between 100 and 9900 (excluding 100 and 9900).
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The unique identifier of the VPC that is associated with the rule group.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	MutationProtection *string `field:"optional" json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the rule group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnFirewallRuleGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallRuleGroupProps := &cfnFirewallRuleGroupProps{
//   	firewallRules: []interface{}{
//   		&firewallRuleProperty{
//   			action: jsii.String("action"),
//   			firewallDomainListId: jsii.String("firewallDomainListId"),
//   			priority: jsii.Number(123),
//
//   			// the properties below are optional
//   			blockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   			blockOverrideDomain: jsii.String("blockOverrideDomain"),
//   			blockOverrideTtl: jsii.Number(123),
//   			blockResponse: jsii.String("blockResponse"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFirewallRuleGroupProps struct {
	// A list of the rules that you have defined.
	FirewallRules interface{} `field:"optional" json:"firewallRules" yaml:"firewallRules"`
	// The name of the rule group.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the rule group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverConfig`.
//
// A complex type that contains information about a Resolver configuration for a VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverConfig := awscdk.Aws_route53resolver.NewCfnResolverConfig(this, jsii.String("MyCfnResolverConfig"), &cfnResolverConfigProps{
//   	autodefinedReverseFlag: jsii.String("autodefinedReverseFlag"),
//   	resourceId: jsii.String("resourceId"),
//   })
//
type CfnResolverConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The status of whether or not the Route 53 Resolver will create autodefined rules for reverse DNS lookups.
	//
	// This is enabled by default.
	AttrAutodefinedReverse() *string
	// ID for the Route 53 Resolver configuration.
	AttrId() *string
	// The owner account ID of the Amazon Virtual Private Cloud VPC.
	AttrOwnerId() *string
	// Represents the desired status of `AutodefinedReverse` .
	//
	// The only supported value on creation is `DISABLE` . Deletion of this resource will return `AutodefinedReverse` to its default value of `ENABLED` .
	AutodefinedReverseFlag() *string
	SetAutodefinedReverseFlag(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the Amazon Virtual Private Cloud VPC that you're configuring Resolver for.
	ResourceId() *string
	SetResourceId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnResolverConfig
type jsiiProxy_CfnResolverConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolverConfig) AttrAutodefinedReverse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAutodefinedReverse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) AttrOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) AutodefinedReverseFlag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autodefinedReverseFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::ResolverConfig`.
func NewCfnResolverConfig(scope constructs.Construct, id *string, props *CfnResolverConfigProps) CfnResolverConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverConfig`.
func NewCfnResolverConfig_Override(c CfnResolverConfig, scope constructs.Construct, id *string, props *CfnResolverConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolverConfig) SetAutodefinedReverseFlag(val *string) {
	_jsii_.Set(
		j,
		"autodefinedReverseFlag",
		val,
	)
}

func (j *jsiiProxy_CfnResolverConfig) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResolverConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResolverConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverConfig",
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
func CfnResolverConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnResolverConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResolverConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResolverConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResolverConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResolverConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResolverConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResolverConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResolverConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResolverConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResolverConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverConfigProps := &cfnResolverConfigProps{
//   	autodefinedReverseFlag: jsii.String("autodefinedReverseFlag"),
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverConfigProps struct {
	// Represents the desired status of `AutodefinedReverse` .
	//
	// The only supported value on creation is `DISABLE` . Deletion of this resource will return `AutodefinedReverse` to its default value of `ENABLED` .
	AutodefinedReverseFlag *string `field:"required" json:"autodefinedReverseFlag" yaml:"autodefinedReverseFlag"`
	// The ID of the Amazon Virtual Private Cloud VPC that you're configuring Resolver for.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverDNSSECConfig`.
//
// The `AWS::Route53Resolver::ResolverDNSSECConfig` resource is a complex type that contains information about a configuration for DNSSEC validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverDNSSECConfig := awscdk.Aws_route53resolver.NewCfnResolverDNSSECConfig(this, jsii.String("MyCfnResolverDNSSECConfig"), &cfnResolverDNSSECConfigProps{
//   	resourceId: jsii.String("resourceId"),
//   })
//
type CfnResolverDNSSECConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The primary identifier of this `ResolverDNSSECConfig` resource.
	//
	// For example: `rdsc-689d45d1ae623bf3` .
	AttrId() *string
	// The AWS account of the owner.
	//
	// For example: `111122223333` .
	AttrOwnerId() *string
	// The current status of this `ResolverDNSSECConfig` resource.
	//
	// For example: `Enabled` .
	AttrValidationStatus() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the virtual private cloud (VPC) that you're configuring the DNSSEC validation status for.
	ResourceId() *string
	SetResourceId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnResolverDNSSECConfig
type jsiiProxy_CfnResolverDNSSECConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) AttrOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) AttrValidationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrValidationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::ResolverDNSSECConfig`.
func NewCfnResolverDNSSECConfig(scope constructs.Construct, id *string, props *CfnResolverDNSSECConfigProps) CfnResolverDNSSECConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverDNSSECConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverDNSSECConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverDNSSECConfig`.
func NewCfnResolverDNSSECConfig_Override(c CfnResolverDNSSECConfig, scope constructs.Construct, id *string, props *CfnResolverDNSSECConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverDNSSECConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolverDNSSECConfig) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResolverDNSSECConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverDNSSECConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResolverDNSSECConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverDNSSECConfig",
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
func CfnResolverDNSSECConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverDNSSECConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverDNSSECConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnResolverDNSSECConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResolverDNSSECConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverDNSSECConfigProps := &cfnResolverDNSSECConfigProps{
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverDNSSECConfigProps struct {
	// The ID of the virtual private cloud (VPC) that you're configuring the DNSSEC validation status for.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverEndpoint`.
//
// Creates a Resolver endpoint. There are two types of Resolver endpoints, inbound and outbound:
//
// - An *inbound Resolver endpoint* forwards DNS queries to the DNS service for a VPC from your network.
// - An *outbound Resolver endpoint* forwards DNS queries from the DNS service for a VPC to your network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverEndpoint := awscdk.Aws_route53resolver.NewCfnResolverEndpoint(this, jsii.String("MyCfnResolverEndpoint"), &cfnResolverEndpointProps{
//   	direction: jsii.String("direction"),
//   	ipAddresses: []interface{}{
//   		&ipAddressRequestProperty{
//   			subnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			ip: jsii.String("ip"),
//   		},
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnResolverEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the resolver endpoint, such as `arn:aws:route53resolver:us-east-1:123456789012:resolver-endpoint/resolver-endpoint-a1bzhi` .
	AttrArn() *string
	// Indicates whether the resolver endpoint allows inbound or outbound DNS queries.
	AttrDirection() *string
	// The ID of the VPC that you want to create the resolver endpoint in.
	AttrHostVpcId() *string
	// The number of IP addresses that the resolver endpoint can use for DNS queries.
	AttrIpAddressCount() *string
	// The name that you assigned to the resolver endpoint when you created the endpoint.
	AttrName() *string
	// The ID of the resolver endpoint.
	AttrResolverEndpointId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Indicates whether the Resolver endpoint allows inbound or outbound DNS queries:.
	//
	// - `INBOUND` : allows DNS queries to your VPC from your network
	// - `OUTBOUND` : allows DNS queries from your VPC to your network.
	Direction() *string
	SetDirection(val *string)
	// The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
	//
	// The subnet ID uniquely identifies a VPC.
	IpAddresses() interface{}
	SetIpAddresses(val interface{})
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
	// A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of one or more security groups that control access to this VPC.
	//
	// The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
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

// The jsii proxy struct for CfnResolverEndpoint
type jsiiProxy_CfnResolverEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolverEndpoint) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) AttrDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) AttrHostVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHostVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) AttrIpAddressCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIpAddressCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) AttrResolverEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResolverEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) IpAddresses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::ResolverEndpoint`.
func NewCfnResolverEndpoint(scope constructs.Construct, id *string, props *CfnResolverEndpointProps) CfnResolverEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverEndpoint`.
func NewCfnResolverEndpoint_Override(c CfnResolverEndpoint, scope constructs.Construct, id *string, props *CfnResolverEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolverEndpoint) SetDirection(val *string) {
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_CfnResolverEndpoint) SetIpAddresses(val interface{}) {
	_jsii_.Set(
		j,
		"ipAddresses",
		val,
	)
}

func (j *jsiiProxy_CfnResolverEndpoint) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnResolverEndpoint) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResolverEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResolverEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverEndpoint",
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
func CfnResolverEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnResolverEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverEndpoint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverEndpoint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverEndpoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// In a [CreateResolverEndpoint](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html) request, the IP address that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints). `IpAddressRequest` also includes the ID of the subnet that contains the IP address.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipAddressRequestProperty := &ipAddressRequestProperty{
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	ip: jsii.String("ip"),
//   }
//
type CfnResolverEndpoint_IpAddressRequestProperty struct {
	// The ID of the subnet that contains the IP address.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The IP address that you want to use for DNS queries.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
}

// Properties for defining a `CfnResolverEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverEndpointProps := &cfnResolverEndpointProps{
//   	direction: jsii.String("direction"),
//   	ipAddresses: []interface{}{
//   		&ipAddressRequestProperty{
//   			subnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			ip: jsii.String("ip"),
//   		},
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResolverEndpointProps struct {
	// Indicates whether the Resolver endpoint allows inbound or outbound DNS queries:.
	//
	// - `INBOUND` : allows DNS queries to your VPC from your network
	// - `OUTBOUND` : allows DNS queries from your VPC to your network.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
	//
	// The subnet ID uniquely identifies a VPC.
	IpAddresses interface{} `field:"required" json:"ipAddresses" yaml:"ipAddresses"`
	// The ID of one or more security groups that control access to this VPC.
	//
	// The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverQueryLoggingConfig`.
//
// The AWS::Route53Resolver::ResolverQueryLoggingConfig resource is a complex type that contains settings for one query logging configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverQueryLoggingConfig := awscdk.Aws_route53resolver.NewCfnResolverQueryLoggingConfig(this, jsii.String("MyCfnResolverQueryLoggingConfig"), &cfnResolverQueryLoggingConfigProps{
//   	destinationArn: jsii.String("destinationArn"),
//   	name: jsii.String("name"),
//   })
//
type CfnResolverQueryLoggingConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) for the query logging configuration.
	AttrArn() *string
	// The number of VPCs that are associated with the query logging configuration.
	AttrAssociationCount() *float64
	// The date and time that the query logging configuration was created, in Unix time format and Coordinated Universal Time (UTC).
	AttrCreationTime() *string
	// A unique string that identifies the request that created the query logging configuration.
	//
	// The `CreatorRequestId` allows failed requests to be retried without the risk of running the operation twice.
	AttrCreatorRequestId() *string
	// The ID for the query logging configuration.
	AttrId() *string
	// The AWS account ID for the account that created the query logging configuration.
	AttrOwnerId() *string
	// An indication of whether the query logging configuration is shared with other AWS account s, or was shared with the current account by another AWS account .
	//
	// Sharing is configured through AWS Resource Access Manager ( AWS RAM ).
	AttrShareStatus() *string
	// The status of the specified query logging configuration. Valid values include the following:.
	//
	// - `CREATING` : Resolver is creating the query logging configuration.
	// - `CREATED` : The query logging configuration was successfully created. Resolver is logging queries that originate in the specified VPC.
	// - `DELETING` : Resolver is deleting this query logging configuration.
	// - `FAILED` : Resolver can't deliver logs to the location that is specified in the query logging configuration. Here are two common causes:
	//
	// - The specified destination (for example, an Amazon S3 bucket) was deleted.
	// - Permissions don't allow sending logs to the destination.
	AttrStatus() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ARN of the resource that you want Resolver to send query logs: an Amazon S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn() *string
	SetDestinationArn(val *string)
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
	// The name of the query logging configuration.
	Name() *string
	SetName(val *string)
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

// The jsii proxy struct for CfnResolverQueryLoggingConfig
type jsiiProxy_CfnResolverQueryLoggingConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) AttrAssociationCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrAssociationCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) AttrCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) AttrOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) AttrShareStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrShareStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) DestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::ResolverQueryLoggingConfig`.
func NewCfnResolverQueryLoggingConfig(scope constructs.Construct, id *string, props *CfnResolverQueryLoggingConfigProps) CfnResolverQueryLoggingConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverQueryLoggingConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverQueryLoggingConfig`.
func NewCfnResolverQueryLoggingConfig_Override(c CfnResolverQueryLoggingConfig, scope constructs.Construct, id *string, props *CfnResolverQueryLoggingConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) SetDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"destinationArn",
		val,
	)
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResolverQueryLoggingConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResolverQueryLoggingConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfig",
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
func CfnResolverQueryLoggingConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverQueryLoggingConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation`.
//
// The AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation resource is a configuration for DNS query logging. After you create a query logging configuration, Amazon Route 53 begins to publish log data to an Amazon CloudWatch Logs log group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverQueryLoggingConfigAssociation := awscdk.Aws_route53resolver.NewCfnResolverQueryLoggingConfigAssociation(this, jsii.String("MyCfnResolverQueryLoggingConfigAssociation"), &cfnResolverQueryLoggingConfigAssociationProps{
//   	resolverQueryLogConfigId: jsii.String("resolverQueryLogConfigId"),
//   	resourceId: jsii.String("resourceId"),
//   })
//
type CfnResolverQueryLoggingConfigAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The date and time that the VPC was associated with the query logging configuration, in Unix time format and Coordinated Universal Time (UTC).
	AttrCreationTime() *string
	// If the value of `Status` is `FAILED` , the value of `Error` indicates the cause:.
	//
	// - `DESTINATION_NOT_FOUND` : The specified destination (for example, an Amazon S3 bucket) was deleted.
	// - `ACCESS_DENIED` : Permissions don't allow sending logs to the destination.
	//
	// If the value of `Status` is a value other than `FAILED` , `Error` is null.
	AttrError() *string
	// Contains additional information about the error.
	//
	// If the value or `Error` is null, the value of `ErrorMessage` is also null.
	AttrErrorMessage() *string
	// The ID of the query logging association.
	AttrId() *string
	// The status of the specified query logging association. Valid values include the following:.
	//
	// - `CREATING` : Resolver is creating an association between an Amazon Virtual Private Cloud (Amazon VPC) and a query logging configuration.
	// - `CREATED` : The association between an Amazon VPC and a query logging configuration was successfully created. Resolver is logging queries that originate in the specified VPC.
	// - `DELETING` : Resolver is deleting this query logging association.
	// - `FAILED` : Resolver either couldn't create or couldn't delete the query logging association.
	AttrStatus() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the query logging configuration that a VPC is associated with.
	ResolverQueryLogConfigId() *string
	SetResolverQueryLogConfigId(val *string)
	// The ID of the Amazon VPC that is associated with the query logging configuration.
	ResourceId() *string
	SetResourceId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnResolverQueryLoggingConfigAssociation
type jsiiProxy_CfnResolverQueryLoggingConfigAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AttrError() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AttrErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrErrorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) ResolverQueryLogConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverQueryLogConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation`.
func NewCfnResolverQueryLoggingConfigAssociation(scope constructs.Construct, id *string, props *CfnResolverQueryLoggingConfigAssociationProps) CfnResolverQueryLoggingConfigAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverQueryLoggingConfigAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation`.
func NewCfnResolverQueryLoggingConfigAssociation_Override(c CfnResolverQueryLoggingConfigAssociation, scope constructs.Construct, id *string, props *CfnResolverQueryLoggingConfigAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) SetResolverQueryLogConfigId(val *string) {
	_jsii_.Set(
		j,
		"resolverQueryLogConfigId",
		val,
	)
}

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResolverQueryLoggingConfigAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResolverQueryLoggingConfigAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
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
func CfnResolverQueryLoggingConfigAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverQueryLoggingConfigAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResolverQueryLoggingConfigAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverQueryLoggingConfigAssociationProps := &cfnResolverQueryLoggingConfigAssociationProps{
//   	resolverQueryLogConfigId: jsii.String("resolverQueryLogConfigId"),
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverQueryLoggingConfigAssociationProps struct {
	// The ID of the query logging configuration that a VPC is associated with.
	ResolverQueryLogConfigId *string `field:"optional" json:"resolverQueryLogConfigId" yaml:"resolverQueryLogConfigId"`
	// The ID of the Amazon VPC that is associated with the query logging configuration.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

// Properties for defining a `CfnResolverQueryLoggingConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverQueryLoggingConfigProps := &cfnResolverQueryLoggingConfigProps{
//   	destinationArn: jsii.String("destinationArn"),
//   	name: jsii.String("name"),
//   }
//
type CfnResolverQueryLoggingConfigProps struct {
	// The ARN of the resource that you want Resolver to send query logs: an Amazon S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// The name of the query logging configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverRule`.
//
// For DNS queries that originate in your VPCs, specifies which Resolver endpoint the queries pass through, one domain name that you want to forward to your network, and the IP addresses of the DNS resolvers in your network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverRule := awscdk.Aws_route53resolver.NewCfnResolverRule(this, jsii.String("MyCfnResolverRule"), &cfnResolverRuleProps{
//   	domainName: jsii.String("domainName"),
//   	ruleType: jsii.String("ruleType"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	resolverEndpointId: jsii.String("resolverEndpointId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetIps: []interface{}{
//   		&targetAddressProperty{
//   			ip: jsii.String("ip"),
//
//   			// the properties below are optional
//   			port: jsii.String("port"),
//   		},
//   	},
//   })
//
type CfnResolverRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the resolver rule, such as `arn:aws:route53resolver:us-east-1:123456789012:resolver-rule/resolver-rule-a1bzhi` .
	AttrArn() *string
	// DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps.
	//
	// If a query matches multiple resolver rules (example.com and www.example.com), the query is routed using the resolver rule that contains the most specific domain name (www.example.com).
	AttrDomainName() *string
	// A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
	AttrName() *string
	// The ID of the outbound endpoint that the rule is associated with, such as `rslvr-out-fdc049932dexample` .
	AttrResolverEndpointId() *string
	// When the value of `RuleType` is `FORWARD` , the ID that Resolver assigned to the resolver rule when you created it, such as `rslvr-rr-5328a0899aexample` .
	//
	// This value isn't applicable when `RuleType` is `SYSTEM` .
	AttrResolverRuleId() *string
	// When the value of `RuleType` is `FORWARD` , the IP addresses that the outbound endpoint forwards DNS queries to, typically the IP addresses for DNS resolvers on your network.
	//
	// This value isn't applicable when `RuleType` is `SYSTEM` .
	AttrTargetIps() awscdk.IResolvable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// DNS queries for this domain name are forwarded to the IP addresses that are specified in `TargetIps` .
	//
	// If a query matches multiple Resolver rules (example.com and www.example.com), the query is routed using the Resolver rule that contains the most specific domain name (www.example.com).
	DomainName() *string
	SetDomainName(val *string)
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
	// The name for the Resolver rule, which you specified when you created the Resolver rule.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId() *string
	SetResolverEndpointId(val *string)
	// When you want to forward DNS queries for specified domain name to resolvers on your network, specify `FORWARD` .
	//
	// When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify `SYSTEM` .
	//
	// For example, to forward DNS queries for example.com to resolvers on your network, you create a rule and specify `FORWARD` for `RuleType` . To then have Resolver process queries for apex.example.com, you create a rule and specify `SYSTEM` for `RuleType` .
	//
	// Currently, only Resolver can create rules that have a value of `RECURSIVE` for `RuleType` .
	RuleType() *string
	SetRuleType(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
	Tags() awscdk.TagManager
	// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to.
	//
	// Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
	TargetIps() interface{}
	SetTargetIps(val interface{})
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

// The jsii proxy struct for CfnResolverRule
type jsiiProxy_CfnResolverRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolverRule) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) AttrResolverEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResolverEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) AttrResolverRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResolverRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) AttrTargetIps() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrTargetIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) ResolverEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) RuleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) TargetIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::ResolverRule`.
func NewCfnResolverRule(scope constructs.Construct, id *string, props *CfnResolverRuleProps) CfnResolverRule {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverRule`.
func NewCfnResolverRule_Override(c CfnResolverRule, scope constructs.Construct, id *string, props *CfnResolverRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolverRule) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnResolverRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnResolverRule) SetResolverEndpointId(val *string) {
	_jsii_.Set(
		j,
		"resolverEndpointId",
		val,
	)
}

func (j *jsiiProxy_CfnResolverRule) SetRuleType(val *string) {
	_jsii_.Set(
		j,
		"ruleType",
		val,
	)
}

func (j *jsiiProxy_CfnResolverRule) SetTargetIps(val interface{}) {
	_jsii_.Set(
		j,
		"targetIps",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResolverRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResolverRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRule",
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
func CfnResolverRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverRule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResolverRule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResolverRule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResolverRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResolverRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResolverRule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResolverRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResolverRule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResolverRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// In a [CreateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html) request, an array of the IPs that you want to forward DNS queries to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetAddressProperty := &targetAddressProperty{
//   	ip: jsii.String("ip"),
//
//   	// the properties below are optional
//   	port: jsii.String("port"),
//   }
//
type CfnResolverRule_TargetAddressProperty struct {
	// One IP address that you want to forward DNS queries to.
	//
	// You can specify only IPv4 addresses.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The port at `Ip` that you want to forward DNS queries to.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverRuleAssociation`.
//
// In the response to an [AssociateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverRule.html) , [DisassociateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverRule.html) , or [ListResolverRuleAssociations](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRuleAssociations.html) request, provides information about an association between a resolver rule and a VPC. The association determines which DNS queries that originate in the VPC are forwarded to your network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverRuleAssociation := awscdk.Aws_route53resolver.NewCfnResolverRuleAssociation(this, jsii.String("MyCfnResolverRuleAssociation"), &cfnResolverRuleAssociationProps{
//   	resolverRuleId: jsii.String("resolverRuleId"),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   })
//
type CfnResolverRuleAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of an association between a resolver rule and a VPC, such as `test.example.com in beta VPC` .
	AttrName() *string
	// The ID of the resolver rule association that you want to get information about, such as `rslvr-rrassoc-97242eaf88example` .
	AttrResolverRuleAssociationId() *string
	// The ID of the resolver rule that you associated with the VPC that is specified by `VPCId` , such as `rslvr-rr-5328a0899example` .
	AttrResolverRuleId() *string
	// The ID of the VPC that you associated the resolver rule with, such as `vpc-03cf94c75cexample` .
	AttrVpcId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// The name of an association between a Resolver rule and a VPC.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the Resolver rule that you associated with the VPC that is specified by `VPCId` .
	ResolverRuleId() *string
	SetResolverRuleId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The ID of the VPC that you associated the Resolver rule with.
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

// The jsii proxy struct for CfnResolverRuleAssociation
type jsiiProxy_CfnResolverRuleAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolverRuleAssociation) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) AttrResolverRuleAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResolverRuleAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) AttrResolverRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResolverRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) AttrVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) ResolverRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolverRuleAssociation) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new `AWS::Route53Resolver::ResolverRuleAssociation`.
func NewCfnResolverRuleAssociation(scope constructs.Construct, id *string, props *CfnResolverRuleAssociationProps) CfnResolverRuleAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverRuleAssociation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRuleAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverRuleAssociation`.
func NewCfnResolverRuleAssociation_Override(c CfnResolverRuleAssociation, scope constructs.Construct, id *string, props *CfnResolverRuleAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRuleAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolverRuleAssociation) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnResolverRuleAssociation) SetResolverRuleId(val *string) {
	_jsii_.Set(
		j,
		"resolverRuleId",
		val,
	)
}

func (j *jsiiProxy_CfnResolverRuleAssociation) SetVpcId(val *string) {
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
func CfnResolverRuleAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRuleAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResolverRuleAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRuleAssociation",
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
func CfnResolverRuleAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRuleAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolverRuleAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_route53resolver.CfnResolverRuleAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResolverRuleAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRuleAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRuleAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRuleAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRuleAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRuleAssociation) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResolverRuleAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverRuleAssociationProps := &cfnResolverRuleAssociationProps{
//   	resolverRuleId: jsii.String("resolverRuleId"),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnResolverRuleAssociationProps struct {
	// The ID of the Resolver rule that you associated with the VPC that is specified by `VPCId` .
	ResolverRuleId *string `field:"required" json:"resolverRuleId" yaml:"resolverRuleId"`
	// The ID of the VPC that you associated the Resolver rule with.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The name of an association between a Resolver rule and a VPC.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Properties for defining a `CfnResolverRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverRuleProps := &cfnResolverRuleProps{
//   	domainName: jsii.String("domainName"),
//   	ruleType: jsii.String("ruleType"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	resolverEndpointId: jsii.String("resolverEndpointId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetIps: []interface{}{
//   		&targetAddressProperty{
//   			ip: jsii.String("ip"),
//
//   			// the properties below are optional
//   			port: jsii.String("port"),
//   		},
//   	},
//   }
//
type CfnResolverRuleProps struct {
	// DNS queries for this domain name are forwarded to the IP addresses that are specified in `TargetIps` .
	//
	// If a query matches multiple Resolver rules (example.com and www.example.com), the query is routed using the Resolver rule that contains the most specific domain name (www.example.com).
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// When you want to forward DNS queries for specified domain name to resolvers on your network, specify `FORWARD` .
	//
	// When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify `SYSTEM` .
	//
	// For example, to forward DNS queries for example.com to resolvers on your network, you create a rule and specify `FORWARD` for `RuleType` . To then have Resolver process queries for apex.example.com, you create a rule and specify `SYSTEM` for `RuleType` .
	//
	// Currently, only Resolver can create rules that have a value of `RECURSIVE` for `RuleType` .
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// The name for the Resolver rule, which you specified when you created the Resolver rule.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId *string `field:"optional" json:"resolverEndpointId" yaml:"resolverEndpointId"`
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to.
	//
	// Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
	TargetIps interface{} `field:"optional" json:"targetIps" yaml:"targetIps"`
}

