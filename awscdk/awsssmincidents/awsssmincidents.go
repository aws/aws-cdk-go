package awsssmincidents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmincidents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::SSMIncidents::ReplicationSet`.
//
// The `AWS::SSMIncidents::ReplicationSet` resource specifies a set of Regions that Incident Manager data is replicated to and the KMS key used to encrypt the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationSet := awscdk.Aws_ssmincidents.NewCfnReplicationSet(this, jsii.String("MyCfnReplicationSet"), &cfnReplicationSetProps{
//   	regions: []interface{}{
//   		&replicationRegionProperty{
//   			regionConfiguration: &regionConfigurationProperty{
//   				sseKmsKeyId: jsii.String("sseKmsKeyId"),
//   			},
//   			regionName: jsii.String("regionName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	deletionProtected: jsii.Boolean(false),
//   })
//
type CfnReplicationSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Determines if the replication set deletion protection is enabled or not.
	//
	// If deletion protection is enabled, you can't delete the last Region in the replication set.
	DeletionProtected() interface{}
	SetDeletionProtected(val interface{})
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
	// Specifies the Regions of the replication set.
	Regions() interface{}
	SetRegions(val interface{})
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

// The jsii proxy struct for CfnReplicationSet
type jsiiProxy_CfnReplicationSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReplicationSet) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) DeletionProtected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) Regions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReplicationSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSMIncidents::ReplicationSet`.
func NewCfnReplicationSet(scope constructs.Construct, id *string, props *CfnReplicationSetProps) CfnReplicationSet {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicationSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ssmincidents.CfnReplicationSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSMIncidents::ReplicationSet`.
func NewCfnReplicationSet_Override(c CfnReplicationSet, scope constructs.Construct, id *string, props *CfnReplicationSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ssmincidents.CfnReplicationSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReplicationSet) SetDeletionProtected(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtected",
		val,
	)
}

func (j *jsiiProxy_CfnReplicationSet) SetRegions(val interface{}) {
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnReplicationSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssmincidents.CfnReplicationSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReplicationSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssmincidents.CfnReplicationSet",
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
func CfnReplicationSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssmincidents.CfnReplicationSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReplicationSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ssmincidents.CfnReplicationSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReplicationSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnReplicationSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnReplicationSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnReplicationSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnReplicationSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnReplicationSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnReplicationSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnReplicationSet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationSet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnReplicationSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReplicationSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReplicationSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `RegionConfiguration` property specifies the Region and KMS key to add to the replication set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regionConfigurationProperty := &regionConfigurationProperty{
//   	sseKmsKeyId: jsii.String("sseKmsKeyId"),
//   }
//
type CfnReplicationSet_RegionConfigurationProperty struct {
	// The KMS key ID to use to encrypt your replication set.
	SseKmsKeyId *string `field:"required" json:"sseKmsKeyId" yaml:"sseKmsKeyId"`
}

// The `ReplicationRegion` property type specifies the Region and KMS key to add to the replication set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRegionProperty := &replicationRegionProperty{
//   	regionConfiguration: &regionConfigurationProperty{
//   		sseKmsKeyId: jsii.String("sseKmsKeyId"),
//   	},
//   	regionName: jsii.String("regionName"),
//   }
//
type CfnReplicationSet_ReplicationRegionProperty struct {
	// Specifies the Region configuration.
	RegionConfiguration interface{} `field:"optional" json:"regionConfiguration" yaml:"regionConfiguration"`
	// Specifies the region name to add to the replication set.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

// Properties for defining a `CfnReplicationSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationSetProps := &cfnReplicationSetProps{
//   	regions: []interface{}{
//   		&replicationRegionProperty{
//   			regionConfiguration: &regionConfigurationProperty{
//   				sseKmsKeyId: jsii.String("sseKmsKeyId"),
//   			},
//   			regionName: jsii.String("regionName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	deletionProtected: jsii.Boolean(false),
//   }
//
type CfnReplicationSetProps struct {
	// Specifies the Regions of the replication set.
	Regions interface{} `field:"required" json:"regions" yaml:"regions"`
	// Determines if the replication set deletion protection is enabled or not.
	//
	// If deletion protection is enabled, you can't delete the last Region in the replication set.
	DeletionProtected interface{} `field:"optional" json:"deletionProtected" yaml:"deletionProtected"`
}

// A CloudFormation `AWS::SSMIncidents::ResponsePlan`.
//
// The `AWS::SSMIncidents::ResponsePlan` resource specifies the details of the response plan that are used when creating an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResponsePlan := awscdk.Aws_ssmincidents.NewCfnResponsePlan(this, jsii.String("MyCfnResponsePlan"), &cfnResponsePlanProps{
//   	incidentTemplate: &incidentTemplateProperty{
//   		impact: jsii.Number(123),
//   		title: jsii.String("title"),
//
//   		// the properties below are optional
//   		dedupeString: jsii.String("dedupeString"),
//   		notificationTargets: []interface{}{
//   			&notificationTargetItemProperty{
//   				snsTopicArn: jsii.String("snsTopicArn"),
//   			},
//   		},
//   		summary: jsii.String("summary"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	actions: []interface{}{
//   		&actionProperty{
//   			ssmAutomation: &ssmAutomationProperty{
//   				documentName: jsii.String("documentName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				documentVersion: jsii.String("documentVersion"),
//   				dynamicParameters: []interface{}{
//   					&dynamicSsmParameterProperty{
//   						key: jsii.String("key"),
//   						value: &dynamicSsmParameterValueProperty{
//   							variable: jsii.String("variable"),
//   						},
//   					},
//   				},
//   				parameters: []interface{}{
//   					&ssmParameterProperty{
//   						key: jsii.String("key"),
//   						values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				targetAccount: jsii.String("targetAccount"),
//   			},
//   		},
//   	},
//   	chatChannel: &chatChannelProperty{
//   		chatbotSns: []*string{
//   			jsii.String("chatbotSns"),
//   		},
//   	},
//   	displayName: jsii.String("displayName"),
//   	engagements: []*string{
//   		jsii.String("engagements"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnResponsePlan interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The actions that the response plan starts at the beginning of an incident.
	Actions() interface{}
	SetActions(val interface{})
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The AWS Chatbot chat channel used for collaboration during an incident.
	ChatChannel() interface{}
	SetChatChannel(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The human readable name of the response plan.
	DisplayName() *string
	SetDisplayName(val *string)
	// The contacts and escalation plans that the response plan engages during an incident.
	Engagements() *[]*string
	SetEngagements(val *[]*string)
	// Details used to create an incident when using this response plan.
	IncidentTemplate() interface{}
	SetIncidentTemplate(val interface{})
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
	// The name of the response plan.
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
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
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

// The jsii proxy struct for CfnResponsePlan
type jsiiProxy_CfnResponsePlan struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResponsePlan) Actions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) ChatChannel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chatChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) Engagements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"engagements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) IncidentTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incidentTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlan) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSMIncidents::ResponsePlan`.
func NewCfnResponsePlan(scope constructs.Construct, id *string, props *CfnResponsePlanProps) CfnResponsePlan {
	_init_.Initialize()

	j := jsiiProxy_CfnResponsePlan{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ssmincidents.CfnResponsePlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSMIncidents::ResponsePlan`.
func NewCfnResponsePlan_Override(c CfnResponsePlan, scope constructs.Construct, id *string, props *CfnResponsePlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ssmincidents.CfnResponsePlan",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResponsePlan) SetActions(val interface{}) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CfnResponsePlan) SetChatChannel(val interface{}) {
	_jsii_.Set(
		j,
		"chatChannel",
		val,
	)
}

func (j *jsiiProxy_CfnResponsePlan) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnResponsePlan) SetEngagements(val *[]*string) {
	_jsii_.Set(
		j,
		"engagements",
		val,
	)
}

func (j *jsiiProxy_CfnResponsePlan) SetIncidentTemplate(val interface{}) {
	_jsii_.Set(
		j,
		"incidentTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnResponsePlan) SetName(val *string) {
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
func CfnResponsePlan_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssmincidents.CfnResponsePlan",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResponsePlan_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssmincidents.CfnResponsePlan",
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
func CfnResponsePlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssmincidents.CfnResponsePlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResponsePlan_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ssmincidents.CfnResponsePlan",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResponsePlan) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResponsePlan) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResponsePlan) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResponsePlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResponsePlan) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResponsePlan) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResponsePlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResponsePlan) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResponsePlan) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResponsePlan) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResponsePlan) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResponsePlan) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResponsePlan) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResponsePlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResponsePlan) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `Action` property type specifies the configuration to launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	ssmAutomation: &ssmAutomationProperty{
//   		documentName: jsii.String("documentName"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		documentVersion: jsii.String("documentVersion"),
//   		dynamicParameters: []interface{}{
//   			&dynamicSsmParameterProperty{
//   				key: jsii.String("key"),
//   				value: &dynamicSsmParameterValueProperty{
//   					variable: jsii.String("variable"),
//   				},
//   			},
//   		},
//   		parameters: []interface{}{
//   			&ssmParameterProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		targetAccount: jsii.String("targetAccount"),
//   	},
//   }
//
type CfnResponsePlan_ActionProperty struct {
	// Details about the Systems Manager automation document that will be used as a runbook during an incident.
	SsmAutomation interface{} `field:"optional" json:"ssmAutomation" yaml:"ssmAutomation"`
}

// The AWS Chatbot chat channel used for collaboration during an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chatChannelProperty := &chatChannelProperty{
//   	chatbotSns: []*string{
//   		jsii.String("chatbotSns"),
//   	},
//   }
//
type CfnResponsePlan_ChatChannelProperty struct {
	// The SNS targets that AWS Chatbot uses to notify the chat channel of updates to an incident.
	//
	// You can also make updates to the incident through the chat channel by using the SNS topics.
	ChatbotSns *[]*string `field:"optional" json:"chatbotSns" yaml:"chatbotSns"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamicSsmParameterProperty := &dynamicSsmParameterProperty{
//   	key: jsii.String("key"),
//   	value: &dynamicSsmParameterValueProperty{
//   		variable: jsii.String("variable"),
//   	},
//   }
//
type CfnResponsePlan_DynamicSsmParameterProperty struct {
	// `CfnResponsePlan.DynamicSsmParameterProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnResponsePlan.DynamicSsmParameterProperty.Value`.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamicSsmParameterValueProperty := &dynamicSsmParameterValueProperty{
//   	variable: jsii.String("variable"),
//   }
//
type CfnResponsePlan_DynamicSsmParameterValueProperty struct {
	// `CfnResponsePlan.DynamicSsmParameterValueProperty.Variable`.
	Variable *string `field:"optional" json:"variable" yaml:"variable"`
}

// The `IncidentTemplate` property type specifies details used to create an incident when using this response plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   incidentTemplateProperty := &incidentTemplateProperty{
//   	impact: jsii.Number(123),
//   	title: jsii.String("title"),
//
//   	// the properties below are optional
//   	dedupeString: jsii.String("dedupeString"),
//   	notificationTargets: []interface{}{
//   		&notificationTargetItemProperty{
//   			snsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	summary: jsii.String("summary"),
//   }
//
type CfnResponsePlan_IncidentTemplateProperty struct {
	// Defines the impact to the customers. Providing an impact overwrites the impact provided by a response plan.
	//
	// **Possible impacts:** - `1` - Critical impact, this typically relates to full application failure that impacts many to all customers.
	// - `2` - High impact, partial application failure with impact to many customers.
	// - `3` - Medium impact, the application is providing reduced service to customers.
	// - `4` - Low impact, customer might aren't impacted by the problem yet.
	// - `5` - No impact, customers aren't currently impacted but urgent action is needed to avoid impact.
	Impact *float64 `field:"required" json:"impact" yaml:"impact"`
	// The title of the incident is a brief and easily recognizable.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Used to create only one incident record for an incident.
	DedupeString *string `field:"optional" json:"dedupeString" yaml:"dedupeString"`
	// The SNS targets that AWS Chatbot uses to notify the chat channel of updates to an incident.
	//
	// You can also make updates to the incident through the chat channel using the SNS topics.
	NotificationTargets interface{} `field:"optional" json:"notificationTargets" yaml:"notificationTargets"`
	// The summary describes what has happened during the incident.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}

// The SNS topic that's used by AWS Chatbot to notify the incidents chat channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationTargetItemProperty := &notificationTargetItemProperty{
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   }
//
type CfnResponsePlan_NotificationTargetItemProperty struct {
	// The Amazon Resource Name (ARN) of the SNS topic.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

// The `SsmAutomation` property type specifies details about the Systems Manager automation document that will be used as a runbook during an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmAutomationProperty := &ssmAutomationProperty{
//   	documentName: jsii.String("documentName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	documentVersion: jsii.String("documentVersion"),
//   	dynamicParameters: []interface{}{
//   		&dynamicSsmParameterProperty{
//   			key: jsii.String("key"),
//   			value: &dynamicSsmParameterValueProperty{
//   				variable: jsii.String("variable"),
//   			},
//   		},
//   	},
//   	parameters: []interface{}{
//   		&ssmParameterProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	targetAccount: jsii.String("targetAccount"),
//   }
//
type CfnResponsePlan_SsmAutomationProperty struct {
	// The automation document's name.
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
	// The Amazon Resource Name (ARN) of the role that the automation document will assume when running commands.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The automation document's version to use when running.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// `CfnResponsePlan.SsmAutomationProperty.DynamicParameters`.
	DynamicParameters interface{} `field:"optional" json:"dynamicParameters" yaml:"dynamicParameters"`
	// The key-value pair parameters to use when running the automation document.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The account that the automation document will be run in.
	//
	// This can be in either the management account or an application account.
	TargetAccount *string `field:"optional" json:"targetAccount" yaml:"targetAccount"`
}

// The key-value pair parameters to use when running the automation document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmParameterProperty := &ssmParameterProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnResponsePlan_SsmParameterProperty struct {
	// The key parameter to use when running the automation document.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value parameter to use when running the automation document.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

// Properties for defining a `CfnResponsePlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResponsePlanProps := &cfnResponsePlanProps{
//   	incidentTemplate: &incidentTemplateProperty{
//   		impact: jsii.Number(123),
//   		title: jsii.String("title"),
//
//   		// the properties below are optional
//   		dedupeString: jsii.String("dedupeString"),
//   		notificationTargets: []interface{}{
//   			&notificationTargetItemProperty{
//   				snsTopicArn: jsii.String("snsTopicArn"),
//   			},
//   		},
//   		summary: jsii.String("summary"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	actions: []interface{}{
//   		&actionProperty{
//   			ssmAutomation: &ssmAutomationProperty{
//   				documentName: jsii.String("documentName"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				documentVersion: jsii.String("documentVersion"),
//   				dynamicParameters: []interface{}{
//   					&dynamicSsmParameterProperty{
//   						key: jsii.String("key"),
//   						value: &dynamicSsmParameterValueProperty{
//   							variable: jsii.String("variable"),
//   						},
//   					},
//   				},
//   				parameters: []interface{}{
//   					&ssmParameterProperty{
//   						key: jsii.String("key"),
//   						values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				targetAccount: jsii.String("targetAccount"),
//   			},
//   		},
//   	},
//   	chatChannel: &chatChannelProperty{
//   		chatbotSns: []*string{
//   			jsii.String("chatbotSns"),
//   		},
//   	},
//   	displayName: jsii.String("displayName"),
//   	engagements: []*string{
//   		jsii.String("engagements"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResponsePlanProps struct {
	// Details used to create an incident when using this response plan.
	IncidentTemplate interface{} `field:"required" json:"incidentTemplate" yaml:"incidentTemplate"`
	// The name of the response plan.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The actions that the response plan starts at the beginning of an incident.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The AWS Chatbot chat channel used for collaboration during an incident.
	ChatChannel interface{} `field:"optional" json:"chatChannel" yaml:"chatChannel"`
	// The human readable name of the response plan.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The contacts and escalation plans that the response plan engages during an incident.
	Engagements *[]*string `field:"optional" json:"engagements" yaml:"engagements"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

