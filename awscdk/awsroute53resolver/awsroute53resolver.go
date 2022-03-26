package awsroute53resolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53resolver/internal"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Route53Resolver::FirewallDomainList`.
//
// High-level information about a list of firewall domains for use in a [AWS::Route53Resolver::FirewallRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-rule.html) . This is returned by [GetFirewallDomainList](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetFirewallDomainList.html) .
//
// To retrieve the domains that are defined for this domain list, call [ListFirewallDomains](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListFirewallDomains.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnFirewallDomainList := route53resolver.NewCfnFirewallDomainList(this, jsii.String("MyCfnFirewallDomainList"), &cfnFirewallDomainListProps{
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
	// Experimental.
	LogicalId() *string
	// The name of the domain list.
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of the tag keys and values that you want to associate with the domain list.
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

func (j *jsiiProxy_CfnFirewallDomainList) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnFirewallDomainList(scope awscdk.Construct, id *string, props *CfnFirewallDomainListProps) CfnFirewallDomainList {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallDomainList{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnFirewallDomainList",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::FirewallDomainList`.
func NewCfnFirewallDomainList_Override(c CfnFirewallDomainList, scope awscdk.Construct, id *string, props *CfnFirewallDomainListProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnFirewallDomainList",
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
// Experimental.
func CfnFirewallDomainList_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallDomainList",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFirewallDomainList_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallDomainList",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFirewallDomainList_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallDomainList",
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
		"monocdk.aws_route53resolver.CfnFirewallDomainList",
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

func (c *jsiiProxy_CfnFirewallDomainList) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallDomainList) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFirewallDomainList) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnFirewallDomainList) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnFirewallDomainList) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
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
	DomainFileUrl *string `json:"domainFileUrl" yaml:"domainFileUrl"`
	// A list of the domain lists that you have defined.
	Domains *[]*string `json:"domains" yaml:"domains"`
	// The name of the domain list.
	Name *string `json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the domain list.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Route53Resolver::FirewallRuleGroup`.
//
// High-level information for a firewall rule group. A firewall rule group is a collection of rules that DNS Firewall uses to filter DNS network traffic for a VPC. To retrieve the rules for the rule group, call [ListFirewallRules](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListFirewallRules.html) .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnFirewallRuleGroup := route53resolver.NewCfnFirewallRuleGroup(this, jsii.String("MyCfnFirewallRuleGroup"), &cfnFirewallRuleGroupProps{
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
	// Experimental.
	LogicalId() *string
	// The name of the rule group.
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of the tag keys and values that you want to associate with the rule group.
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

func (j *jsiiProxy_CfnFirewallRuleGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnFirewallRuleGroup(scope awscdk.Construct, id *string, props *CfnFirewallRuleGroupProps) CfnFirewallRuleGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallRuleGroup{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::FirewallRuleGroup`.
func NewCfnFirewallRuleGroup_Override(c CfnFirewallRuleGroup, scope awscdk.Construct, id *string, props *CfnFirewallRuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroup",
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
// Experimental.
func CfnFirewallRuleGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFirewallRuleGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFirewallRuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroup",
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
		"monocdk.aws_route53resolver.CfnFirewallRuleGroup",
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

func (c *jsiiProxy_CfnFirewallRuleGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnFirewallRuleGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnFirewallRuleGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
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
	Action *string `json:"action" yaml:"action"`
	// The ID of the domain list that's used in the rule.
	FirewallDomainListId *string `json:"firewallDomainListId" yaml:"firewallDomainListId"`
	// The priority of the rule in the rule group.
	//
	// This value must be unique within the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The DNS record's type.
	//
	// This determines the format of the record value that you provided in `BlockOverrideDomain` . Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	BlockOverrideDnsType *string `json:"blockOverrideDnsType" yaml:"blockOverrideDnsType"`
	// The custom DNS record to send back in response to the query.
	//
	// Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	BlockOverrideDomain *string `json:"blockOverrideDomain" yaml:"blockOverrideDomain"`
	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.
	//
	// Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	BlockOverrideTtl *float64 `json:"blockOverrideTtl" yaml:"blockOverrideTtl"`
	// The way that you want DNS Firewall to block the request. Used for the rule action setting `BLOCK` .
	//
	// - `NODATA` - Respond indicating that the query was successful, but no response is available for it.
	// - `NXDOMAIN` - Respond indicating that the domain name that's in the query doesn't exist.
	// - `OVERRIDE` - Provide a custom override in the response. This option requires custom handling details in the rule's `BlockOverride*` settings.
	BlockResponse *string `json:"blockResponse" yaml:"blockResponse"`
}

// A CloudFormation `AWS::Route53Resolver::FirewallRuleGroupAssociation`.
//
// An association between a firewall rule group and a VPC, which enables DNS filtering for the VPC.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnFirewallRuleGroupAssociation := route53resolver.NewCfnFirewallRuleGroupAssociation(this, jsii.String("MyCfnFirewallRuleGroupAssociation"), &cfnFirewallRuleGroupAssociationProps{
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
	// Experimental.
	LogicalId() *string
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	MutationProtection() *string
	SetMutationProtection(val *string)
	// The name of the association.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of the tag keys and values that you want to associate with the rule group.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The unique identifier of the VPC that is associated with the rule group.
	VpcId() *string
	SetVpcId(val *string)
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

func (j *jsiiProxy_CfnFirewallRuleGroupAssociation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnFirewallRuleGroupAssociation(scope awscdk.Construct, id *string, props *CfnFirewallRuleGroupAssociationProps) CfnFirewallRuleGroupAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallRuleGroupAssociation{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroupAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::FirewallRuleGroupAssociation`.
func NewCfnFirewallRuleGroupAssociation_Override(c CfnFirewallRuleGroupAssociation, scope awscdk.Construct, id *string, props *CfnFirewallRuleGroupAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroupAssociation",
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
// Experimental.
func CfnFirewallRuleGroupAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroupAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFirewallRuleGroupAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroupAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFirewallRuleGroupAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnFirewallRuleGroupAssociation",
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
		"monocdk.aws_route53resolver.CfnFirewallRuleGroupAssociation",
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

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnFirewallRuleGroupAssociation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
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
	FirewallRuleGroupId *string `json:"firewallRuleGroupId" yaml:"firewallRuleGroupId"`
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC traffic starting from rule group with the lowest numeric priority setting.
	//
	// You must specify a unique priority for each rule group that you associate with a single VPC. To make it easier to insert rule groups later, leave space between the numbers, for example, use 101, 200, and so on. You can change the priority setting for a rule group association after you create it.
	//
	// The allowed values for `Priority` are between 100 and 9900 (excluding 100 and 9900).
	Priority *float64 `json:"priority" yaml:"priority"`
	// The unique identifier of the VPC that is associated with the rule group.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	MutationProtection *string `json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	Name *string `json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the rule group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnFirewallRuleGroup`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
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
	FirewallRules interface{} `json:"firewallRules" yaml:"firewallRules"`
	// The name of the rule group.
	Name *string `json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the rule group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverConfig`.
//
// A complex type that contains information about a Resolver configuration for a VPC.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverConfig := route53resolver.NewCfnResolverConfig(this, jsii.String("MyCfnResolverConfig"), &cfnResolverConfigProps{
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
	// The ID of the Amazon Virtual Private Cloud VPC that you're configuring Resolver for.
	ResourceId() *string
	SetResourceId(val *string)
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

func (j *jsiiProxy_CfnResolverConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnResolverConfig(scope awscdk.Construct, id *string, props *CfnResolverConfigProps) CfnResolverConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverConfig{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverConfig`.
func NewCfnResolverConfig_Override(c CfnResolverConfig, scope awscdk.Construct, id *string, props *CfnResolverConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverConfig",
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
// Experimental.
func CfnResolverConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResolverConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResolverConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverConfig",
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
		"monocdk.aws_route53resolver.CfnResolverConfig",
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

func (c *jsiiProxy_CfnResolverConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResolverConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResolverConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnResolverConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnResolverConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverConfigProps := &cfnResolverConfigProps{
//   	autodefinedReverseFlag: jsii.String("autodefinedReverseFlag"),
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverConfigProps struct {
	// Represents the desired status of `AutodefinedReverse` .
	//
	// The only supported value on creation is `DISABLE` . Deletion of this resource will return `AutodefinedReverse` to its default value of `ENABLED` .
	AutodefinedReverseFlag *string `json:"autodefinedReverseFlag" yaml:"autodefinedReverseFlag"`
	// The ID of the Amazon Virtual Private Cloud VPC that you're configuring Resolver for.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverDNSSECConfig`.
//
// The `AWS::Route53Resolver::ResolverDNSSECConfig` resource is a complex type that contains information about a configuration for DNSSEC validation.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverDNSSECConfig := route53resolver.NewCfnResolverDNSSECConfig(this, jsii.String("MyCfnResolverDNSSECConfig"), &cfnResolverDNSSECConfigProps{
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
	// The ID of the virtual private cloud (VPC) that you're configuring the DNSSEC validation status for.
	ResourceId() *string
	SetResourceId(val *string)
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

func (j *jsiiProxy_CfnResolverDNSSECConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnResolverDNSSECConfig(scope awscdk.Construct, id *string, props *CfnResolverDNSSECConfigProps) CfnResolverDNSSECConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverDNSSECConfig{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverDNSSECConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverDNSSECConfig`.
func NewCfnResolverDNSSECConfig_Override(c CfnResolverDNSSECConfig, scope awscdk.Construct, id *string, props *CfnResolverDNSSECConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverDNSSECConfig",
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
// Experimental.
func CfnResolverDNSSECConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverDNSSECConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResolverDNSSECConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverDNSSECConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResolverDNSSECConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverDNSSECConfig",
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
		"monocdk.aws_route53resolver.CfnResolverDNSSECConfig",
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

func (c *jsiiProxy_CfnResolverDNSSECConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverDNSSECConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnResolverDNSSECConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnResolverDNSSECConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverDNSSECConfigProps := &cfnResolverDNSSECConfigProps{
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverDNSSECConfigProps struct {
	// The ID of the virtual private cloud (VPC) that you're configuring the DNSSEC validation status for.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverEndpoint`.
//
// Creates a Resolver endpoint. There are two types of Resolver endpoints, inbound and outbound:
//
// - An *inbound Resolver endpoint* forwards DNS queries to the DNS service for a VPC from your network.
// - An *outbound Resolver endpoint* forwards DNS queries from the DNS service for a VPC to your network.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverEndpoint := route53resolver.NewCfnResolverEndpoint(this, jsii.String("MyCfnResolverEndpoint"), &cfnResolverEndpointProps{
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
	// Experimental.
	LogicalId() *string
	// A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.
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
	// The ID of one or more security groups that control access to this VPC.
	//
	// The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
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

func (j *jsiiProxy_CfnResolverEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnResolverEndpoint(scope awscdk.Construct, id *string, props *CfnResolverEndpointProps) CfnResolverEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverEndpoint{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverEndpoint`.
func NewCfnResolverEndpoint_Override(c CfnResolverEndpoint, scope awscdk.Construct, id *string, props *CfnResolverEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverEndpoint",
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
// Experimental.
func CfnResolverEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResolverEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResolverEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverEndpoint",
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
		"monocdk.aws_route53resolver.CfnResolverEndpoint",
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

func (c *jsiiProxy_CfnResolverEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnResolverEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnResolverEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   ipAddressRequestProperty := &ipAddressRequestProperty{
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	ip: jsii.String("ip"),
//   }
//
type CfnResolverEndpoint_IpAddressRequestProperty struct {
	// The ID of the subnet that contains the IP address.
	SubnetId *string `json:"subnetId" yaml:"subnetId"`
	// The IP address that you want to use for DNS queries.
	Ip *string `json:"ip" yaml:"ip"`
}

// Properties for defining a `CfnResolverEndpoint`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
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
	Direction *string `json:"direction" yaml:"direction"`
	// The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
	//
	// The subnet ID uniquely identifies a VPC.
	IpAddresses interface{} `json:"ipAddresses" yaml:"ipAddresses"`
	// The ID of one or more security groups that control access to this VPC.
	//
	// The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.
	Name *string `json:"name" yaml:"name"`
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverQueryLoggingConfig`.
//
// The AWS::Route53Resolver::ResolverQueryLoggingConfig resource is a complex type that contains settings for one query logging configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverQueryLoggingConfig := route53resolver.NewCfnResolverQueryLoggingConfig(this, jsii.String("MyCfnResolverQueryLoggingConfig"), &cfnResolverQueryLoggingConfigProps{
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
	// Experimental.
	LogicalId() *string
	// The name of the query logging configuration.
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

func (j *jsiiProxy_CfnResolverQueryLoggingConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnResolverQueryLoggingConfig(scope awscdk.Construct, id *string, props *CfnResolverQueryLoggingConfigProps) CfnResolverQueryLoggingConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverQueryLoggingConfig{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverQueryLoggingConfig`.
func NewCfnResolverQueryLoggingConfig_Override(c CfnResolverQueryLoggingConfig, scope awscdk.Construct, id *string, props *CfnResolverQueryLoggingConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfig",
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
// Experimental.
func CfnResolverQueryLoggingConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResolverQueryLoggingConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResolverQueryLoggingConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfig",
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
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfig",
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

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnResolverQueryLoggingConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverQueryLoggingConfigAssociation := route53resolver.NewCfnResolverQueryLoggingConfigAssociation(this, jsii.String("MyCfnResolverQueryLoggingConfigAssociation"), &cfnResolverQueryLoggingConfigAssociationProps{
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
	// The ID of the query logging configuration that a VPC is associated with.
	ResolverQueryLogConfigId() *string
	SetResolverQueryLogConfigId(val *string)
	// The ID of the Amazon VPC that is associated with the query logging configuration.
	ResourceId() *string
	SetResourceId(val *string)
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

func (j *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnResolverQueryLoggingConfigAssociation(scope awscdk.Construct, id *string, props *CfnResolverQueryLoggingConfigAssociationProps) CfnResolverQueryLoggingConfigAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverQueryLoggingConfigAssociation{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation`.
func NewCfnResolverQueryLoggingConfigAssociation_Override(c CfnResolverQueryLoggingConfigAssociation, scope awscdk.Construct, id *string, props *CfnResolverQueryLoggingConfigAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
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
// Experimental.
func CfnResolverQueryLoggingConfigAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResolverQueryLoggingConfigAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResolverQueryLoggingConfigAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
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
		"monocdk.aws_route53resolver.CfnResolverQueryLoggingConfigAssociation",
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

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnResolverQueryLoggingConfigAssociation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverQueryLoggingConfigAssociationProps := &cfnResolverQueryLoggingConfigAssociationProps{
//   	resolverQueryLogConfigId: jsii.String("resolverQueryLogConfigId"),
//   	resourceId: jsii.String("resourceId"),
//   }
//
type CfnResolverQueryLoggingConfigAssociationProps struct {
	// The ID of the query logging configuration that a VPC is associated with.
	ResolverQueryLogConfigId *string `json:"resolverQueryLogConfigId" yaml:"resolverQueryLogConfigId"`
	// The ID of the Amazon VPC that is associated with the query logging configuration.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
}

// Properties for defining a `CfnResolverQueryLoggingConfig`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverQueryLoggingConfigProps := &cfnResolverQueryLoggingConfigProps{
//   	destinationArn: jsii.String("destinationArn"),
//   	name: jsii.String("name"),
//   }
//
type CfnResolverQueryLoggingConfigProps struct {
	// The ARN of the resource that you want Resolver to send query logs: an Amazon S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
	// The name of the query logging configuration.
	Name *string `json:"name" yaml:"name"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverRule`.
//
// For DNS queries that originate in your VPCs, specifies which Resolver endpoint the queries pass through, one domain name that you want to forward to your network, and the IP addresses of the DNS resolvers in your network.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverRule := route53resolver.NewCfnResolverRule(this, jsii.String("MyCfnResolverRule"), &cfnResolverRuleProps{
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
	// Experimental.
	LogicalId() *string
	// The name for the Resolver rule, which you specified when you created the Resolver rule.
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
	// Experimental.
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

func (j *jsiiProxy_CfnResolverRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnResolverRule(scope awscdk.Construct, id *string, props *CfnResolverRuleProps) CfnResolverRule {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverRule{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverRule`.
func NewCfnResolverRule_Override(c CfnResolverRule, scope awscdk.Construct, id *string, props *CfnResolverRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverRule",
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
// Experimental.
func CfnResolverRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResolverRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResolverRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverRule",
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
		"monocdk.aws_route53resolver.CfnResolverRule",
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

func (c *jsiiProxy_CfnResolverRule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResolverRule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResolverRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverRule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnResolverRule) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnResolverRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
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
	Ip *string `json:"ip" yaml:"ip"`
	// The port at `Ip` that you want to forward DNS queries to.
	Port *string `json:"port" yaml:"port"`
}

// A CloudFormation `AWS::Route53Resolver::ResolverRuleAssociation`.
//
// In the response to an [AssociateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverRule.html) , [DisassociateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverRule.html) , or [ListResolverRuleAssociations](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRuleAssociations.html) request, provides information about an association between a resolver rule and a VPC. The association determines which DNS queries that originate in the VPC are forwarded to your network.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   cfnResolverRuleAssociation := route53resolver.NewCfnResolverRuleAssociation(this, jsii.String("MyCfnResolverRuleAssociation"), &cfnResolverRuleAssociationProps{
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
	// The name of an association between a Resolver rule and a VPC.
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
	// The ID of the Resolver rule that you associated with the VPC that is specified by `VPCId` .
	ResolverRuleId() *string
	SetResolverRuleId(val *string)
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
	// The ID of the VPC that you associated the Resolver rule with.
	VpcId() *string
	SetVpcId(val *string)
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

func (j *jsiiProxy_CfnResolverRuleAssociation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnResolverRuleAssociation(scope awscdk.Construct, id *string, props *CfnResolverRuleAssociationProps) CfnResolverRuleAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnResolverRuleAssociation{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverRuleAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Route53Resolver::ResolverRuleAssociation`.
func NewCfnResolverRuleAssociation_Override(c CfnResolverRuleAssociation, scope awscdk.Construct, id *string, props *CfnResolverRuleAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.CfnResolverRuleAssociation",
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
// Experimental.
func CfnResolverRuleAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverRuleAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResolverRuleAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverRuleAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResolverRuleAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.CfnResolverRuleAssociation",
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
		"monocdk.aws_route53resolver.CfnResolverRuleAssociation",
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

func (c *jsiiProxy_CfnResolverRuleAssociation) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResolverRuleAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResolverRuleAssociation) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnResolverRuleAssociation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnResolverRuleAssociation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
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
	ResolverRuleId *string `json:"resolverRuleId" yaml:"resolverRuleId"`
	// The ID of the VPC that you associated the Resolver rule with.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
	// The name of an association between a Resolver rule and a VPC.
	Name *string `json:"name" yaml:"name"`
}

// Properties for defining a `CfnResolverRule`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
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
	DomainName *string `json:"domainName" yaml:"domainName"`
	// When you want to forward DNS queries for specified domain name to resolvers on your network, specify `FORWARD` .
	//
	// When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify `SYSTEM` .
	//
	// For example, to forward DNS queries for example.com to resolvers on your network, you create a rule and specify `FORWARD` for `RuleType` . To then have Resolver process queries for apex.example.com, you create a rule and specify `SYSTEM` for `RuleType` .
	//
	// Currently, only Resolver can create rules that have a value of `RECURSIVE` for `RuleType` .
	RuleType *string `json:"ruleType" yaml:"ruleType"`
	// The name for the Resolver rule, which you specified when you created the Resolver rule.
	Name *string `json:"name" yaml:"name"`
	// The ID of the endpoint that the rule is associated with.
	ResolverEndpointId *string `json:"resolverEndpointId" yaml:"resolverEndpointId"`
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to.
	//
	// Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.
	TargetIps interface{} `json:"targetIps" yaml:"targetIps"`
}

// The way that you want DNS Firewall to block the request.
//
// Example:
//   var myBlockList firewallDomainList
//   var ruleGroup firewallRuleGroup
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(10),
//   	firewallDomainList: myBlockList,
//   	// block and reply with NXDOMAIN
//   	action: route53resolver.firewallRuleAction.block(route53resolver.dnsBlockResponse.nxDomain()),
//   })
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(20),
//   	firewallDomainList: myBlockList,
//   	// block and override DNS response with a custom domain
//   	action: route53resolver.*firewallRuleAction.block(route53resolver.*dnsBlockResponse.override(jsii.String("amazon.com"))),
//   })
//
// Experimental.
type DnsBlockResponse interface {
	// The DNS record's type.
	// Experimental.
	BlockOverrideDnsType() *string
	// The custom DNS record to send back in response to the query.
	// Experimental.
	BlockOverrideDomain() *string
	// The recommended amount of time for the DNS resolver or web browser to cache the provided override record.
	// Experimental.
	BlockOverrideTtl() awscdk.Duration
	// The way that you want DNS Firewall to block the request.
	// Experimental.
	BlockResponse() *string
}

// The jsii proxy struct for DnsBlockResponse
type jsiiProxy_DnsBlockResponse struct {
	_ byte // padding
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideDnsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideTtl() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"blockOverrideTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}


// Experimental.
func NewDnsBlockResponse_Override(d DnsBlockResponse) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.DnsBlockResponse",
		nil, // no parameters
		d,
	)
}

// Respond indicating that the query was successful, but no response is available for it.
// Experimental.
func DnsBlockResponse_NoData() DnsBlockResponse {
	_init_.Initialize()

	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.DnsBlockResponse",
		"noData",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Respond indicating that the domain name that's in the query doesn't exist.
// Experimental.
func DnsBlockResponse_NxDomain() DnsBlockResponse {
	_init_.Initialize()

	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.DnsBlockResponse",
		"nxDomain",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Provides a custom override response to the query.
// Experimental.
func DnsBlockResponse_Override(domain *string, ttl awscdk.Duration) DnsBlockResponse {
	_init_.Initialize()

	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.DnsBlockResponse",
		"override",
		[]interface{}{domain, ttl},
		&returns,
	)

	return returns
}

// Domains configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//   domainsConfig := &domainsConfig{
//   	domainFileUrl: jsii.String("domainFileUrl"),
//   	domains: []*string{
//   		jsii.String("domains"),
//   	},
//   }
//
// Experimental.
type DomainsConfig struct {
	// The fully qualified URL or URI of the file stored in Amazon S3 that contains the list of domains to import.
	//
	// The file must be a text file and must contain
	// a single domain per line. The content type of the S3 object must be `plain/text`.
	// Experimental.
	DomainFileUrl *string `json:"domainFileUrl" yaml:"domainFileUrl"`
	// A list of domains.
	// Experimental.
	Domains *[]*string `json:"domains" yaml:"domains"`
}

// A Firewall Domain List.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
//   	domains: route53resolver.firewallDomains.fromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
//   })
//
// Experimental.
type FirewallDomainList interface {
	awscdk.Resource
	IFirewallDomainList
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN (Amazon Resource Name) of the domain list.
	// Experimental.
	FirewallDomainListArn() *string
	// The date and time that the domain list was created.
	// Experimental.
	FirewallDomainListCreationTime() *string
	// The creator request ID.
	// Experimental.
	FirewallDomainListCreatorRequestId() *string
	// The number of domains in the list.
	// Experimental.
	FirewallDomainListDomainCount() *float64
	// The ID of the domain list.
	// Experimental.
	FirewallDomainListId() *string
	// The owner of the list, used only for lists that are not managed by you.
	//
	// For example, the managed domain list `AWSManagedDomainsMalwareDomainList`
	// has the managed owner name `Route 53 Resolver DNS Firewall`.
	// Experimental.
	FirewallDomainListManagedOwnerName() *string
	// The date and time that the domain list was last modified.
	// Experimental.
	FirewallDomainListModificationTime() *string
	// The status of the domain list.
	// Experimental.
	FirewallDomainListStatus() *string
	// Additional information about the status of the rule group.
	// Experimental.
	FirewallDomainListStatusMessage() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
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
}

// The jsii proxy struct for FirewallDomainList
type jsiiProxy_FirewallDomainList struct {
	internal.Type__awscdkResource
	jsiiProxy_IFirewallDomainList
}

func (j *jsiiProxy_FirewallDomainList) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListDomainCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firewallDomainListDomainCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListManagedOwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListManagedOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallDomainList(scope constructs.Construct, id *string, props *FirewallDomainListProps) FirewallDomainList {
	_init_.Initialize()

	j := jsiiProxy_FirewallDomainList{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallDomainList",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallDomainList_Override(f FirewallDomainList, scope constructs.Construct, id *string, props *FirewallDomainListProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallDomainList",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing Firewall Rule Group.
// Experimental.
func FirewallDomainList_FromFirewallDomainListId(scope constructs.Construct, id *string, firewallDomainListId *string) IFirewallDomainList {
	_init_.Initialize()

	var returns IFirewallDomainList

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomainList",
		"fromFirewallDomainListId",
		[]interface{}{scope, id, firewallDomainListId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func FirewallDomainList_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomainList",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FirewallDomainList_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomainList",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FirewallDomainList) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallDomainList) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FirewallDomainList) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallDomainList) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FirewallDomainList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a Firewall Domain List.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
//   	domains: route53resolver.firewallDomains.fromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
//   })
//
// Experimental.
type FirewallDomainListProps struct {
	// A list of domains.
	// Experimental.
	Domains FirewallDomains `json:"domains" yaml:"domains"`
	// A name for the domain list.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

// A list of domains.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
//   	domains: route53resolver.firewallDomains.fromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
//   })
//
// Experimental.
type FirewallDomains interface {
	// Binds the domains to a domain list.
	// Experimental.
	Bind(scope constructs.Construct) *DomainsConfig
}

// The jsii proxy struct for FirewallDomains
type jsiiProxy_FirewallDomains struct {
	_ byte // padding
}

// Experimental.
func NewFirewallDomains_Override(f FirewallDomains) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallDomains",
		nil, // no parameters
		f,
	)
}

// Firewall domains created from a local disk path to a text file.
//
// The file must be a text file (`.txt` extension) and must contain a single
// domain per line. It will be uploaded to S3.
// Experimental.
func FirewallDomains_FromAsset(assetPath *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomains",
		"fromAsset",
		[]interface{}{assetPath},
		&returns,
	)

	return returns
}

// Firewall domains created from a list of domains.
// Experimental.
func FirewallDomains_FromList(list *[]*string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomains",
		"fromList",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// Firewall domains created from a file stored in Amazon S3.
//
// The file must be a text file and must contain a single domain per line.
// The content type of the S3 object must be `plain/text`.
// Experimental.
func FirewallDomains_FromS3(bucket awss3.IBucket, key *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomains",
		"fromS3",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Firewall domains created from the URL of a file stored in Amazon S3.
//
// The file must be a text file and must contain a single domain per line.
// The content type of the S3 object must be `plain/text`.
// Experimental.
func FirewallDomains_FromS3Url(url *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomains",
		"fromS3Url",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomains) Bind(scope constructs.Construct) *DomainsConfig {
	var returns *DomainsConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// A Firewall Rule.
//
// Example:
//   var myBlockList firewallDomainList
//   var ruleGroup firewallRuleGroup
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(10),
//   	firewallDomainList: myBlockList,
//   	// block and reply with NXDOMAIN
//   	action: route53resolver.firewallRuleAction.block(route53resolver.dnsBlockResponse.nxDomain()),
//   })
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(20),
//   	firewallDomainList: myBlockList,
//   	// block and override DNS response with a custom domain
//   	action: route53resolver.*firewallRuleAction.block(route53resolver.*dnsBlockResponse.override(jsii.String("amazon.com"))),
//   })
//
// Experimental.
type FirewallRule struct {
	// The action for this rule.
	// Experimental.
	Action FirewallRuleAction `json:"action" yaml:"action"`
	// The domain list for this rule.
	// Experimental.
	FirewallDomainList IFirewallDomainList `json:"firewallDomainList" yaml:"firewallDomainList"`
	// The priority of the rule in the rule group.
	//
	// This value must be unique within
	// the rule group.
	// Experimental.
	Priority *float64 `json:"priority" yaml:"priority"`
}

// A Firewall Rule.
//
// Example:
//   var myBlockList firewallDomainList
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
//   	rules: []firewallRule{
//   		&firewallRule{
//   			priority: jsii.Number(10),
//   			firewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			action: route53resolver.firewallRuleAction.block(),
//   		},
//   	},
//   })
//
// Experimental.
type FirewallRuleAction interface {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list.
	// Experimental.
	Action() *string
	// The way that you want DNS Firewall to block the request.
	// Experimental.
	BlockResponse() DnsBlockResponse
}

// The jsii proxy struct for FirewallRuleAction
type jsiiProxy_FirewallRuleAction struct {
	_ byte // padding
}

func (j *jsiiProxy_FirewallRuleAction) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleAction) BlockResponse() DnsBlockResponse {
	var returns DnsBlockResponse
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallRuleAction_Override(f FirewallRuleAction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallRuleAction",
		nil, // no parameters
		f,
	)
}

// Permit the request to go through but send an alert to the logs.
// Experimental.
func FirewallRuleAction_Alert() FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallRuleAction",
		"alert",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Permit the request to go through.
// Experimental.
func FirewallRuleAction_Allow() FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallRuleAction",
		"allow",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Disallow the request.
// Experimental.
func FirewallRuleAction_Block(response DnsBlockResponse) FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallRuleAction",
		"block",
		[]interface{}{response},
		&returns,
	)

	return returns
}

// A Firewall Rule Group.
//
// Example:
//   var myBlockList firewallDomainList
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
//   	rules: []firewallRule{
//   		&firewallRule{
//   			priority: jsii.Number(10),
//   			firewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			action: route53resolver.firewallRuleAction.block(),
//   		},
//   	},
//   })
//
// Experimental.
type FirewallRuleGroup interface {
	awscdk.Resource
	IFirewallRuleGroup
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN (Amazon Resource Name) of the rule group.
	// Experimental.
	FirewallRuleGroupArn() *string
	// The date and time that the rule group was created.
	// Experimental.
	FirewallRuleGroupCreationTime() *string
	// The creator request ID.
	// Experimental.
	FirewallRuleGroupCreatorRequestId() *string
	// The ID of the rule group.
	// Experimental.
	FirewallRuleGroupId() *string
	// The date and time that the rule group was last modified.
	// Experimental.
	FirewallRuleGroupModificationTime() *string
	// The AWS account ID for the account that created the rule group.
	// Experimental.
	FirewallRuleGroupOwnerId() *string
	// The number of rules in the rule group.
	// Experimental.
	FirewallRuleGroupRuleCount() *float64
	// Whether the rule group is shared with other AWS accounts, or was shared with the current account by another AWS account.
	// Experimental.
	FirewallRuleGroupShareStatus() *string
	// The status of the rule group.
	// Experimental.
	FirewallRuleGroupStatus() *string
	// Additional information about the status of the rule group.
	// Experimental.
	FirewallRuleGroupStatusMessage() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds a rule to this group.
	// Experimental.
	AddRule(rule *FirewallRule) FirewallRuleGroup
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Associates this Firewall Rule Group with a VPC.
	// Experimental.
	Associate(id *string, props *FirewallRuleGroupAssociationOptions) FirewallRuleGroupAssociation
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
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
}

// The jsii proxy struct for FirewallRuleGroup
type jsiiProxy_FirewallRuleGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IFirewallRuleGroup
}

func (j *jsiiProxy_FirewallRuleGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupRuleCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firewallRuleGroupRuleCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupShareStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupShareStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallRuleGroup(scope constructs.Construct, id *string, props *FirewallRuleGroupProps) FirewallRuleGroup {
	_init_.Initialize()

	j := jsiiProxy_FirewallRuleGroup{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallRuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallRuleGroup_Override(f FirewallRuleGroup, scope constructs.Construct, id *string, props *FirewallRuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallRuleGroup",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing Firewall Rule Group.
// Experimental.
func FirewallRuleGroup_FromFirewallRuleGroupId(scope constructs.Construct, id *string, firewallRuleGroupId *string) IFirewallRuleGroup {
	_init_.Initialize()

	var returns IFirewallRuleGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallRuleGroup",
		"fromFirewallRuleGroupId",
		[]interface{}{scope, id, firewallRuleGroupId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func FirewallRuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallRuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FirewallRuleGroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallRuleGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) AddRule(rule *FirewallRule) FirewallRuleGroup {
	var returns FirewallRuleGroup

	_jsii_.Invoke(
		f,
		"addRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FirewallRuleGroup) Associate(id *string, props *FirewallRuleGroupAssociationOptions) FirewallRuleGroupAssociation {
	var returns FirewallRuleGroupAssociation

	_jsii_.Invoke(
		f,
		"associate",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRuleGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FirewallRuleGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRuleGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FirewallRuleGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A Firewall Rule Group Association.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//
//   var firewallRuleGroup firewallRuleGroup
//   var vpc vpc
//   firewallRuleGroupAssociation := route53resolver.NewFirewallRuleGroupAssociation(this, jsii.String("MyFirewallRuleGroupAssociation"), &firewallRuleGroupAssociationProps{
//   	firewallRuleGroup: firewallRuleGroup,
//   	priority: jsii.Number(123),
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	mutationProtection: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   })
//
// Experimental.
type FirewallRuleGroupAssociation interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN (Amazon Resource Name) of the association.
	// Experimental.
	FirewallRuleGroupAssociationArn() *string
	// The date and time that the association was created.
	// Experimental.
	FirewallRuleGroupAssociationCreationTime() *string
	// The creator request ID.
	// Experimental.
	FirewallRuleGroupAssociationCreatorRequestId() *string
	// The ID of the association.
	// Experimental.
	FirewallRuleGroupAssociationId() *string
	// The owner of the association, used only for lists that are not managed by you.
	//
	// If you use AWS Firewall Manager to manage your firewallls from DNS Firewall,
	// then this reports Firewall Manager as the managed owner.
	// Experimental.
	FirewallRuleGroupAssociationManagedOwnerName() *string
	// The date and time that the association was last modified.
	// Experimental.
	FirewallRuleGroupAssociationModificationTime() *string
	// The status of the association.
	// Experimental.
	FirewallRuleGroupAssociationStatus() *string
	// Additional information about the status of the association.
	// Experimental.
	FirewallRuleGroupAssociationStatusMessage() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
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
}

// The jsii proxy struct for FirewallRuleGroupAssociation
type jsiiProxy_FirewallRuleGroupAssociation struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationManagedOwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationManagedOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) FirewallRuleGroupAssociationStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroupAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallRuleGroupAssociation(scope constructs.Construct, id *string, props *FirewallRuleGroupAssociationProps) FirewallRuleGroupAssociation {
	_init_.Initialize()

	j := jsiiProxy_FirewallRuleGroupAssociation{}

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallRuleGroupAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallRuleGroupAssociation_Override(f FirewallRuleGroupAssociation, scope constructs.Construct, id *string, props *FirewallRuleGroupAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallRuleGroupAssociation",
		[]interface{}{scope, id, props},
		f,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func FirewallRuleGroupAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallRuleGroupAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FirewallRuleGroupAssociation_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallRuleGroupAssociation",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroupAssociation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for a Firewall Rule Group Association.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var ruleGroup firewallRuleGroup
//   var myVpc vpc
//
//   ruleGroup.associate(jsii.String("Association"), &firewallRuleGroupAssociationOptions{
//   	priority: jsii.Number(101),
//   	vpc: myVpc,
//   })
//
// Experimental.
type FirewallRuleGroupAssociationOptions struct {
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC
	// traffic starting from rule group with the lowest numeric priority setting.
	//
	// This value must be greater than 100 and less than 9,000.
	// Experimental.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The VPC that to associate with the rule group.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	// Experimental.
	MutationProtection *bool `json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
}

// Properties for a Firewall Rule Group Association.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import route53resolver "github.com/aws/aws-cdk-go/awscdk/aws_route53resolver"
//
//   var firewallRuleGroup firewallRuleGroup
//   var vpc vpc
//   firewallRuleGroupAssociationProps := &firewallRuleGroupAssociationProps{
//   	firewallRuleGroup: firewallRuleGroup,
//   	priority: jsii.Number(123),
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	mutationProtection: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   }
//
// Experimental.
type FirewallRuleGroupAssociationProps struct {
	// The setting that determines the processing order of the rule group among the rule groups that are associated with a single VPC.
	//
	// DNS Firewall filters VPC
	// traffic starting from rule group with the lowest numeric priority setting.
	//
	// This value must be greater than 100 and less than 9,000.
	// Experimental.
	Priority *float64 `json:"priority" yaml:"priority"`
	// The VPC that to associate with the rule group.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.
	// Experimental.
	MutationProtection *bool `json:"mutationProtection" yaml:"mutationProtection"`
	// The name of the association.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// The firewall rule group which must be associated.
	// Experimental.
	FirewallRuleGroup IFirewallRuleGroup `json:"firewallRuleGroup" yaml:"firewallRuleGroup"`
}

// Properties for a Firewall Rule Group.
//
// Example:
//   var myBlockList firewallDomainList
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
//   	rules: []firewallRule{
//   		&firewallRule{
//   			priority: jsii.Number(10),
//   			firewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			action: route53resolver.firewallRuleAction.block(),
//   		},
//   	},
//   })
//
// Experimental.
type FirewallRuleGroupProps struct {
	// The name of the rule group.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// A list of rules for this group.
	// Experimental.
	Rules *[]*FirewallRule `json:"rules" yaml:"rules"`
}

// A Firewall Domain List.
// Experimental.
type IFirewallDomainList interface {
	awscdk.IResource
	// The ID of the domain list.
	// Experimental.
	FirewallDomainListId() *string
}

// The jsii proxy for IFirewallDomainList
type jsiiProxy_IFirewallDomainList struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFirewallDomainList) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}

// A Firewall Rule Group.
// Experimental.
type IFirewallRuleGroup interface {
	awscdk.IResource
	// The ID of the rule group.
	// Experimental.
	FirewallRuleGroupId() *string
}

// The jsii proxy for IFirewallRuleGroup
type jsiiProxy_IFirewallRuleGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFirewallRuleGroup) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

