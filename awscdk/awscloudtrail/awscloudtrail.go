package awscloudtrail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudtrail/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/constructs-go/constructs/v3"
)

// Options for adding an event selector.
//
// Example:
//   import cloudtrail "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceBucket bucket
//
//   sourceOutput := codepipeline.NewArtifact()
//   key := "some/key.zip"
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
//   trail.addS3EventSelector([]s3EventSelector{
//   	&s3EventSelector{
//   		bucket: sourceBucket,
//   		objectPrefix: key,
//   	},
//   }, &addEventSelectorOptions{
//   	readWriteType: cloudtrail.readWriteType_WRITE_ONLY,
//   })
//   sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
//   	actionName: jsii.String("S3Source"),
//   	bucketKey: key,
//   	bucket: sourceBucket,
//   	output: sourceOutput,
//   	trigger: codepipeline_actions.s3Trigger_EVENTS,
//   })
//
// Experimental.
type AddEventSelectorOptions struct {
	// An optional list of service event sources from which you do not want management events to be logged on your trail.
	// Experimental.
	ExcludeManagementEventSources *[]ManagementEventSources `field:"optional" json:"excludeManagementEventSources" yaml:"excludeManagementEventSources"`
	// Specifies whether the event selector includes management events for the trail.
	// Experimental.
	IncludeManagementEvents *bool `field:"optional" json:"includeManagementEvents" yaml:"includeManagementEvents"`
	// Specifies whether to log read-only events, write-only events, or all events.
	// Experimental.
	ReadWriteType ReadWriteType `field:"optional" json:"readWriteType" yaml:"readWriteType"`
}

// A CloudFormation `AWS::CloudTrail::EventDataStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventDataStore := awscdk.Aws_cloudtrail.NewCfnEventDataStore(this, jsii.String("MyCfnEventDataStore"), &cfnEventDataStoreProps{
//   	advancedEventSelectors: []interface{}{
//   		&advancedEventSelectorProperty{
//   			fieldSelectors: []interface{}{
//   				&advancedFieldSelectorProperty{
//   					field: jsii.String("field"),
//
//   					// the properties below are optional
//   					endsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					equalTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					notEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					notEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					notStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					startsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   		},
//   	},
//   	multiRegionEnabled: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	organizationEnabled: jsii.Boolean(false),
//   	retentionPeriod: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	terminationProtectionEnabled: jsii.Boolean(false),
//   })
//
type CfnEventDataStore interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::CloudTrail::EventDataStore.AdvancedEventSelectors`.
	AdvancedEventSelectors() interface{}
	SetAdvancedEventSelectors(val interface{})
	AttrCreatedTimestamp() *string
	AttrEventDataStoreArn() *string
	AttrStatus() *string
	AttrUpdatedTimestamp() *string
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
	// `AWS::CloudTrail::EventDataStore.MultiRegionEnabled`.
	MultiRegionEnabled() interface{}
	SetMultiRegionEnabled(val interface{})
	// `AWS::CloudTrail::EventDataStore.Name`.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// `AWS::CloudTrail::EventDataStore.OrganizationEnabled`.
	OrganizationEnabled() interface{}
	SetOrganizationEnabled(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::CloudTrail::EventDataStore.RetentionPeriod`.
	RetentionPeriod() *float64
	SetRetentionPeriod(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::CloudTrail::EventDataStore.Tags`.
	Tags() awscdk.TagManager
	// `AWS::CloudTrail::EventDataStore.TerminationProtectionEnabled`.
	TerminationProtectionEnabled() interface{}
	SetTerminationProtectionEnabled(val interface{})
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

// The jsii proxy struct for CfnEventDataStore
type jsiiProxy_CfnEventDataStore struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventDataStore) AdvancedEventSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedEventSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) AttrCreatedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) AttrEventDataStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEventDataStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) AttrUpdatedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdatedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) MultiRegionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiRegionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) OrganizationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) RetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) TerminationProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStore) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudTrail::EventDataStore`.
func NewCfnEventDataStore(scope awscdk.Construct, id *string, props *CfnEventDataStoreProps) CfnEventDataStore {
	_init_.Initialize()

	j := jsiiProxy_CfnEventDataStore{}

	_jsii_.Create(
		"monocdk.aws_cloudtrail.CfnEventDataStore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudTrail::EventDataStore`.
func NewCfnEventDataStore_Override(c CfnEventDataStore, scope awscdk.Construct, id *string, props *CfnEventDataStoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudtrail.CfnEventDataStore",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventDataStore) SetAdvancedEventSelectors(val interface{}) {
	_jsii_.Set(
		j,
		"advancedEventSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore) SetMultiRegionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"multiRegionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore) SetOrganizationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"organizationEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore) SetRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"retentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnEventDataStore) SetTerminationProtectionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"terminationProtectionEnabled",
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
func CfnEventDataStore_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnEventDataStore",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEventDataStore_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnEventDataStore",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEventDataStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnEventDataStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventDataStore_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cloudtrail.CfnEventDataStore",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventDataStore) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEventDataStore) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEventDataStore) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEventDataStore) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEventDataStore) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEventDataStore) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEventDataStore) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventDataStore) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEventDataStore) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnEventDataStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventDataStore) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedEventSelectorProperty := &advancedEventSelectorProperty{
//   	fieldSelectors: []interface{}{
//   		&advancedFieldSelectorProperty{
//   			field: jsii.String("field"),
//
//   			// the properties below are optional
//   			endsWith: []*string{
//   				jsii.String("endsWith"),
//   			},
//   			equalTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			notEndsWith: []*string{
//   				jsii.String("notEndsWith"),
//   			},
//   			notEquals: []*string{
//   				jsii.String("notEquals"),
//   			},
//   			notStartsWith: []*string{
//   				jsii.String("notStartsWith"),
//   			},
//   			startsWith: []*string{
//   				jsii.String("startsWith"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnEventDataStore_AdvancedEventSelectorProperty struct {
	// `CfnEventDataStore.AdvancedEventSelectorProperty.FieldSelectors`.
	FieldSelectors interface{} `field:"required" json:"fieldSelectors" yaml:"fieldSelectors"`
	// `CfnEventDataStore.AdvancedEventSelectorProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedFieldSelectorProperty := &advancedFieldSelectorProperty{
//   	field: jsii.String("field"),
//
//   	// the properties below are optional
//   	endsWith: []*string{
//   		jsii.String("endsWith"),
//   	},
//   	equalTo: []*string{
//   		jsii.String("equalTo"),
//   	},
//   	notEndsWith: []*string{
//   		jsii.String("notEndsWith"),
//   	},
//   	notEquals: []*string{
//   		jsii.String("notEquals"),
//   	},
//   	notStartsWith: []*string{
//   		jsii.String("notStartsWith"),
//   	},
//   	startsWith: []*string{
//   		jsii.String("startsWith"),
//   	},
//   }
//
type CfnEventDataStore_AdvancedFieldSelectorProperty struct {
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.Field`.
	Field *string `field:"required" json:"field" yaml:"field"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.EndsWith`.
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.Equals`.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.NotEndsWith`.
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.NotEquals`.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.NotStartsWith`.
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.StartsWith`.
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

// Properties for defining a `CfnEventDataStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventDataStoreProps := &cfnEventDataStoreProps{
//   	advancedEventSelectors: []interface{}{
//   		&advancedEventSelectorProperty{
//   			fieldSelectors: []interface{}{
//   				&advancedFieldSelectorProperty{
//   					field: jsii.String("field"),
//
//   					// the properties below are optional
//   					endsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					equalTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					notEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					notEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					notStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					startsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   		},
//   	},
//   	multiRegionEnabled: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	organizationEnabled: jsii.Boolean(false),
//   	retentionPeriod: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	terminationProtectionEnabled: jsii.Boolean(false),
//   }
//
type CfnEventDataStoreProps struct {
	// `AWS::CloudTrail::EventDataStore.AdvancedEventSelectors`.
	AdvancedEventSelectors interface{} `field:"optional" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
	// `AWS::CloudTrail::EventDataStore.MultiRegionEnabled`.
	MultiRegionEnabled interface{} `field:"optional" json:"multiRegionEnabled" yaml:"multiRegionEnabled"`
	// `AWS::CloudTrail::EventDataStore.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::CloudTrail::EventDataStore.OrganizationEnabled`.
	OrganizationEnabled interface{} `field:"optional" json:"organizationEnabled" yaml:"organizationEnabled"`
	// `AWS::CloudTrail::EventDataStore.RetentionPeriod`.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// `AWS::CloudTrail::EventDataStore.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::CloudTrail::EventDataStore.TerminationProtectionEnabled`.
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
}

// A CloudFormation `AWS::CloudTrail::Trail`.
//
// Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrail := awscdk.Aws_cloudtrail.NewCfnTrail(this, jsii.String("MyCfnTrail"), &cfnTrailProps{
//   	isLogging: jsii.Boolean(false),
//   	s3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	cloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   	cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	enableLogFileValidation: jsii.Boolean(false),
//   	eventSelectors: []interface{}{
//   		&eventSelectorProperty{
//   			dataResources: []interface{}{
//   				&dataResourceProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			excludeManagementEventSources: []*string{
//   				jsii.String("excludeManagementEventSources"),
//   			},
//   			includeManagementEvents: jsii.Boolean(false),
//   			readWriteType: jsii.String("readWriteType"),
//   		},
//   	},
//   	includeGlobalServiceEvents: jsii.Boolean(false),
//   	insightSelectors: []interface{}{
//   		&insightSelectorProperty{
//   			insightType: jsii.String("insightType"),
//   		},
//   	},
//   	isMultiRegionTrail: jsii.Boolean(false),
//   	isOrganizationTrail: jsii.Boolean(false),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	snsTopicName: jsii.String("snsTopicName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trailName: jsii.String("trailName"),
//   })
//
type CfnTrail interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `Ref` returns the ARN of the CloudTrail trail, such as `arn:aws:cloudtrail:us-east-2:123456789012:trail/myCloudTrail` .
	AttrArn() *string
	// `Ref` returns the ARN of the Amazon SNS topic that's associated with the CloudTrail trail, such as `arn:aws:sns:us-east-2:123456789012:mySNSTopic` .
	AttrSnsTopicArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs are delivered.
	//
	// Not required unless you specify `CloudWatchLogsRoleArn` .
	CloudWatchLogsLogGroupArn() *string
	SetCloudWatchLogsLogGroupArn(val *string)
	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
	CloudWatchLogsRoleArn() *string
	SetCloudWatchLogsRoleArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Specifies whether log file validation is enabled. The default is false.
	//
	// > When you disable log file integrity validation, the chain of digest files is broken after one hour. CloudTrail does not create digest files for log files that were delivered during a period in which log file integrity validation was disabled. For example, if you enable log file integrity validation at noon on January 1, disable it at noon on January 2, and re-enable it at noon on January 10, digest files will not be created for the log files delivered from noon on January 2 to noon on January 10. The same applies whenever you stop CloudTrail logging or delete a trail.
	EnableLogFileValidation() interface{}
	SetEnableLogFileValidation(val interface{})
	// Use event selectors to further specify the management and data event settings for your trail.
	//
	// By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event.
	//
	// You can configure up to five event selectors for a trail.
	//
	// You cannot apply both event selectors and advanced event selectors to a trail.
	EventSelectors() interface{}
	SetEventSelectors(val interface{})
	// Specifies whether the trail is publishing events from global services such as IAM to the log files.
	IncludeGlobalServiceEvents() interface{}
	SetIncludeGlobalServiceEvents(val interface{})
	// `AWS::CloudTrail::Trail.InsightSelectors`.
	InsightSelectors() interface{}
	SetInsightSelectors(val interface{})
	// Whether the CloudTrail trail is currently logging AWS API calls.
	IsLogging() interface{}
	SetIsLogging(val interface{})
	// Specifies whether the trail applies only to the current region or to all regions.
	//
	// The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
	IsMultiRegionTrail() interface{}
	SetIsMultiRegionTrail(val interface{})
	// Specifies whether the trail is applied to all accounts in an organization in AWS Organizations , or only for the current AWS account .
	//
	// The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the management account for an organization in AWS Organizations . If the trail is not an organization trail and this is set to `true` , the trail will be created in all AWS accounts that belong to the organization. If the trail is an organization trail and this is set to `false` , the trail will remain in the current AWS account but be deleted from all member accounts in the organization.
	IsOrganizationTrail() interface{}
	SetIsOrganizationTrail(val interface{})
	// Specifies the AWS KMS key ID to use to encrypt the logs delivered by CloudTrail.
	//
	// The value can be an alias name prefixed by "alias/", a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// CloudTrail also supports AWS KMS multi-Region keys. For more information about multi-Region keys, see [Using multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
	//
	// Examples:
	//
	// - alias/MyAliasName
	// - arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	// - arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	// - 12345678-1234-1234-1234-123456789012.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// Specifies the name of the Amazon S3 bucket designated for publishing log files.
	//
	// See [Amazon S3 Bucket Naming Requirements](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html) .
	S3BucketName() *string
	SetS3BucketName(val *string)
	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.
	//
	// For more information, see [Finding Your CloudTrail Log Files](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html) . The maximum length is 200 characters.
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	// Specifies the name of the Amazon SNS topic defined for notification of log file delivery.
	//
	// The maximum length is 256 characters.
	SnsTopicName() *string
	SetSnsTopicName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A custom set of tags (key-value pairs) for this trail.
	Tags() awscdk.TagManager
	// Specifies the name of the trail. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	// - Start with a letter or number, and end with a letter or number
	// - Be between 3 and 128 characters
	// - Have no adjacent periods, underscores or dashes. Names like `my-_namespace` and `my--namespace` are not valid.
	// - Not be in IP address format (for example, 192.168.5.4)
	TrailName() *string
	SetTrailName(val *string)
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

// The jsii proxy struct for CfnTrail
type jsiiProxy_CfnTrail struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTrail) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) AttrSnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CloudWatchLogsLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CloudWatchLogsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) EnableLogFileValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) EventSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IncludeGlobalServiceEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) InsightSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsMultiRegionTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsOrganizationTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) SnsTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) TrailName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudTrail::Trail`.
func NewCfnTrail(scope awscdk.Construct, id *string, props *CfnTrailProps) CfnTrail {
	_init_.Initialize()

	j := jsiiProxy_CfnTrail{}

	_jsii_.Create(
		"monocdk.aws_cloudtrail.CfnTrail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudTrail::Trail`.
func NewCfnTrail_Override(c CfnTrail, scope awscdk.Construct, id *string, props *CfnTrailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudtrail.CfnTrail",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTrail) SetCloudWatchLogsLogGroupArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetCloudWatchLogsRoleArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetEnableLogFileValidation(val interface{}) {
	_jsii_.Set(
		j,
		"enableLogFileValidation",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetEventSelectors(val interface{}) {
	_jsii_.Set(
		j,
		"eventSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetIncludeGlobalServiceEvents(val interface{}) {
	_jsii_.Set(
		j,
		"includeGlobalServiceEvents",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetInsightSelectors(val interface{}) {
	_jsii_.Set(
		j,
		"insightSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetIsLogging(val interface{}) {
	_jsii_.Set(
		j,
		"isLogging",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetIsMultiRegionTrail(val interface{}) {
	_jsii_.Set(
		j,
		"isMultiRegionTrail",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetIsOrganizationTrail(val interface{}) {
	_jsii_.Set(
		j,
		"isOrganizationTrail",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetS3BucketName(val *string) {
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetSnsTopicName(val *string) {
	_jsii_.Set(
		j,
		"snsTopicName",
		val,
	)
}

func (j *jsiiProxy_CfnTrail) SetTrailName(val *string) {
	_jsii_.Set(
		j,
		"trailName",
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
func CfnTrail_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnTrail",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTrail_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnTrail",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnTrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrail_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cloudtrail.CfnTrail",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrail) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTrail) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTrail) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTrail) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTrail) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTrail) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTrail) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTrail) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTrail) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTrail) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTrail) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTrail) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTrail) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The Amazon S3 buckets, AWS Lambda functions, or Amazon DynamoDB tables that you specify in event selectors in your AWS CloudFormation template for your trail to log data events.
//
// Data events provide information about the resource operations performed on or within a resource itself. These are also known as data plane operations. You can specify up to 250 data resources for a trail. Currently, advanced event selectors for data events are not supported in AWS CloudFormation templates.
//
// > The total number of allowed data resources is 250. This number can be distributed between 1 and 5 event selectors, but the total cannot exceed 250 across all selectors.
// >
// > If you are using advanced event selectors, the maximum total number of values for all conditions, across all advanced event selectors for the trail, is 500.
//
// The following example demonstrates how logging works when you configure logging of all data events for an S3 bucket named `bucket-1` . In this example, the CloudTrail user specified an empty prefix, and the option to log both `Read` and `Write` data events.
//
// - A user uploads an image file to `bucket-1` .
// - The `PutObject` API operation is an Amazon S3 object-level API. It is recorded as a data event in CloudTrail. Because the CloudTrail user specified an S3 bucket with an empty prefix, events that occur on any object in that bucket are logged. The trail processes and logs the event.
// - A user uploads an object to an Amazon S3 bucket named `arn:aws:s3:::bucket-2` .
// - The `PutObject` API operation occurred for an object in an S3 bucket that the CloudTrail user didn't specify for the trail. The trail doesn’t log the event.
//
// The following example demonstrates how logging works when you configure logging of AWS Lambda data events for a Lambda function named *MyLambdaFunction* , but not for all Lambda functions.
//
// - A user runs a script that includes a call to the *MyLambdaFunction* function and the *MyOtherLambdaFunction* function.
// - The `Invoke` API operation on *MyLambdaFunction* is an Lambda API. It is recorded as a data event in CloudTrail. Because the CloudTrail user specified logging data events for *MyLambdaFunction* , any invocations of that function are logged. The trail processes and logs the event.
// - The `Invoke` API operation on *MyOtherLambdaFunction* is an Lambda API. Because the CloudTrail user did not specify logging data events for all Lambda functions, the `Invoke` operation for *MyOtherLambdaFunction* does not match the function specified for the trail. The trail doesn’t log the event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataResourceProperty := &dataResourceProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnTrail_DataResourceProperty struct {
	// The resource type in which you want to log data events.
	//
	// You can specify the following *basic* event selector resource types:
	//
	// - `AWS::S3::Object`
	// - `AWS::Lambda::Function`
	// - `AWS::DynamoDB::Table`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.
	//
	// - To log data events for all objects in all S3 buckets in your AWS account , specify the prefix as `arn:aws:s3` .
	//
	// > This also enables logging of data event activity performed by any user or role in your AWS account , even if that activity is performed on a bucket that belongs to another AWS account .
	// - To log data events for all objects in an S3 bucket, specify the bucket and an empty object prefix such as `arn:aws:s3:::bucket-1/` . The trail logs data events for all objects in this S3 bucket.
	// - To log data events for specific objects, specify the S3 bucket and object prefix such as `arn:aws:s3:::bucket-1/example-images` . The trail logs data events for objects in this S3 bucket that match the prefix.
	// - To log data events for all Lambda functions in your AWS account , specify the prefix as `arn:aws:lambda` .
	//
	// > This also enables logging of `Invoke` activity performed by any user or role in your AWS account , even if that activity is performed on a function that belongs to another AWS account .
	// - To log data events for a specific Lambda function, specify the function ARN.
	//
	// > Lambda function ARNs are exact. For example, if you specify a function ARN *arn:aws:lambda:us-west-2:111111111111:function:helloworld* , data events will only be logged for *arn:aws:lambda:us-west-2:111111111111:function:helloworld* . They will not be logged for *arn:aws:lambda:us-west-2:111111111111:function:helloworld2* .
	// - To log data events for all DynamoDB tables in your AWS account , specify the prefix as `arn:aws:dynamodb` .
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Use event selectors to further specify the management and data event settings for your trail.
//
// By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event.
//
// You can configure up to five event selectors for a trail.
//
// You cannot apply both event selectors and advanced event selectors to a trail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSelectorProperty := &eventSelectorProperty{
//   	dataResources: []interface{}{
//   		&dataResourceProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	excludeManagementEventSources: []*string{
//   		jsii.String("excludeManagementEventSources"),
//   	},
//   	includeManagementEvents: jsii.Boolean(false),
//   	readWriteType: jsii.String("readWriteType"),
//   }
//
type CfnTrail_EventSelectorProperty struct {
	// In AWS CloudFormation , CloudTrail supports data event logging for Amazon S3 objects, Amazon DynamoDB tables, and AWS Lambda functions.
	//
	// Currently, advanced event selectors for data events are not supported in AWS CloudFormation templates. You can specify up to 250 resources for an individual event selector, but the total number of data resources cannot exceed 250 across all event selectors in a trail. This limit does not apply if you configure resource logging for all data events.
	//
	// For more information, see [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-and-data-events-with-cloudtrail.html#logging-data-events) and [Limits in AWS CloudTrail](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) in the *AWS CloudTrail User Guide* .
	DataResources interface{} `field:"optional" json:"dataResources" yaml:"dataResources"`
	// An optional list of service event sources from which you do not want management events to be logged on your trail.
	//
	// In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service or Amazon RDS Data API events by containing `kms.amazonaws.com` or `rdsdata.amazonaws.com` . By default, `ExcludeManagementEventSources` is empty, and AWS KMS and Amazon RDS Data API events are logged to your trail. You can exclude management event sources only in regions that support the event source.
	ExcludeManagementEventSources *[]*string `field:"optional" json:"excludeManagementEventSources" yaml:"excludeManagementEventSources"`
	// Specify if you want your event selector to include management events for your trail.
	//
	// For more information, see [Management Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-and-data-events-with-cloudtrail.html#logging-management-events) in the *AWS CloudTrail User Guide* .
	//
	// By default, the value is `true` .
	//
	// The first copy of management events is free. You are charged for additional copies of management events that you are logging on any subsequent trail in the same region. For more information about CloudTrail pricing, see [AWS CloudTrail Pricing](https://docs.aws.amazon.com/cloudtrail/pricing/) .
	IncludeManagementEvents interface{} `field:"optional" json:"includeManagementEvents" yaml:"includeManagementEvents"`
	// Specify if you want your trail to log read-only events, write-only events, or all.
	//
	// For example, the EC2 `GetConsoleOutput` is a read-only API operation and `RunInstances` is a write-only API operation.
	//
	// By default, the value is `All` .
	ReadWriteType *string `field:"optional" json:"readWriteType" yaml:"readWriteType"`
}

// A JSON string that contains a list of insight types that are logged on a trail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   insightSelectorProperty := &insightSelectorProperty{
//   	insightType: jsii.String("insightType"),
//   }
//
type CfnTrail_InsightSelectorProperty struct {
	// The type of insights to log on a trail.
	//
	// `ApiCallRateInsight` and `ApiErrorRateInsight` are valid insight types.
	InsightType *string `field:"optional" json:"insightType" yaml:"insightType"`
}

// Properties for defining a `CfnTrail`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrailProps := &cfnTrailProps{
//   	isLogging: jsii.Boolean(false),
//   	s3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	cloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   	cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	enableLogFileValidation: jsii.Boolean(false),
//   	eventSelectors: []interface{}{
//   		&eventSelectorProperty{
//   			dataResources: []interface{}{
//   				&dataResourceProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			excludeManagementEventSources: []*string{
//   				jsii.String("excludeManagementEventSources"),
//   			},
//   			includeManagementEvents: jsii.Boolean(false),
//   			readWriteType: jsii.String("readWriteType"),
//   		},
//   	},
//   	includeGlobalServiceEvents: jsii.Boolean(false),
//   	insightSelectors: []interface{}{
//   		&insightSelectorProperty{
//   			insightType: jsii.String("insightType"),
//   		},
//   	},
//   	isMultiRegionTrail: jsii.Boolean(false),
//   	isOrganizationTrail: jsii.Boolean(false),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	snsTopicName: jsii.String("snsTopicName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trailName: jsii.String("trailName"),
//   }
//
type CfnTrailProps struct {
	// Whether the CloudTrail trail is currently logging AWS API calls.
	IsLogging interface{} `field:"required" json:"isLogging" yaml:"isLogging"`
	// Specifies the name of the Amazon S3 bucket designated for publishing log files.
	//
	// See [Amazon S3 Bucket Naming Requirements](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html) .
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs are delivered.
	//
	// Not required unless you specify `CloudWatchLogsRoleArn` .
	CloudWatchLogsLogGroupArn *string `field:"optional" json:"cloudWatchLogsLogGroupArn" yaml:"cloudWatchLogsLogGroupArn"`
	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
	CloudWatchLogsRoleArn *string `field:"optional" json:"cloudWatchLogsRoleArn" yaml:"cloudWatchLogsRoleArn"`
	// Specifies whether log file validation is enabled. The default is false.
	//
	// > When you disable log file integrity validation, the chain of digest files is broken after one hour. CloudTrail does not create digest files for log files that were delivered during a period in which log file integrity validation was disabled. For example, if you enable log file integrity validation at noon on January 1, disable it at noon on January 2, and re-enable it at noon on January 10, digest files will not be created for the log files delivered from noon on January 2 to noon on January 10. The same applies whenever you stop CloudTrail logging or delete a trail.
	EnableLogFileValidation interface{} `field:"optional" json:"enableLogFileValidation" yaml:"enableLogFileValidation"`
	// Use event selectors to further specify the management and data event settings for your trail.
	//
	// By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event.
	//
	// You can configure up to five event selectors for a trail.
	//
	// You cannot apply both event selectors and advanced event selectors to a trail.
	EventSelectors interface{} `field:"optional" json:"eventSelectors" yaml:"eventSelectors"`
	// Specifies whether the trail is publishing events from global services such as IAM to the log files.
	IncludeGlobalServiceEvents interface{} `field:"optional" json:"includeGlobalServiceEvents" yaml:"includeGlobalServiceEvents"`
	// `AWS::CloudTrail::Trail.InsightSelectors`.
	InsightSelectors interface{} `field:"optional" json:"insightSelectors" yaml:"insightSelectors"`
	// Specifies whether the trail applies only to the current region or to all regions.
	//
	// The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
	IsMultiRegionTrail interface{} `field:"optional" json:"isMultiRegionTrail" yaml:"isMultiRegionTrail"`
	// Specifies whether the trail is applied to all accounts in an organization in AWS Organizations , or only for the current AWS account .
	//
	// The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the management account for an organization in AWS Organizations . If the trail is not an organization trail and this is set to `true` , the trail will be created in all AWS accounts that belong to the organization. If the trail is an organization trail and this is set to `false` , the trail will remain in the current AWS account but be deleted from all member accounts in the organization.
	IsOrganizationTrail interface{} `field:"optional" json:"isOrganizationTrail" yaml:"isOrganizationTrail"`
	// Specifies the AWS KMS key ID to use to encrypt the logs delivered by CloudTrail.
	//
	// The value can be an alias name prefixed by "alias/", a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// CloudTrail also supports AWS KMS multi-Region keys. For more information about multi-Region keys, see [Using multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
	//
	// Examples:
	//
	// - alias/MyAliasName
	// - arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	// - arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	// - 12345678-1234-1234-1234-123456789012.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.
	//
	// For more information, see [Finding Your CloudTrail Log Files](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html) . The maximum length is 200 characters.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic defined for notification of log file delivery.
	//
	// The maximum length is 256 characters.
	SnsTopicName *string `field:"optional" json:"snsTopicName" yaml:"snsTopicName"`
	// A custom set of tags (key-value pairs) for this trail.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the name of the trail. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	// - Start with a letter or number, and end with a letter or number
	// - Be between 3 and 128 characters
	// - Have no adjacent periods, underscores or dashes. Names like `my-_namespace` and `my--namespace` are not valid.
	// - Not be in IP address format (for example, 192.168.5.4)
	TrailName *string `field:"optional" json:"trailName" yaml:"trailName"`
}

// Resource type for a data event.
// Experimental.
type DataResourceType string

const (
	// Data resource type for Lambda function.
	// Experimental.
	DataResourceType_LAMBDA_FUNCTION DataResourceType = "LAMBDA_FUNCTION"
	// Data resource type for S3 objects.
	// Experimental.
	DataResourceType_S3_OBJECT DataResourceType = "S3_OBJECT"
)

// Types of management event sources that can be excluded.
// Experimental.
type ManagementEventSources string

const (
	// AWS Key Management Service (AWS KMS) events.
	// Experimental.
	ManagementEventSources_KMS ManagementEventSources = "KMS"
	// Data API events.
	// Experimental.
	ManagementEventSources_RDS_DATA_API ManagementEventSources = "RDS_DATA_API"
)

// Types of events that CloudTrail can log.
//
// Example:
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &trailProps{
//   	// ...
//   	managementEvents: cloudtrail.readWriteType_READ_ONLY,
//   })
//
// Experimental.
type ReadWriteType string

const (
	// Read-only events include API operations that read your resources, but don't make changes.
	//
	// For example, read-only events include the Amazon EC2 DescribeSecurityGroups
	// and DescribeSubnets API operations.
	// Experimental.
	ReadWriteType_READ_ONLY ReadWriteType = "READ_ONLY"
	// Write-only events include API operations that modify (or might modify) your resources.
	//
	// For example, the Amazon EC2 RunInstances and TerminateInstances API
	// operations modify your instances.
	// Experimental.
	ReadWriteType_WRITE_ONLY ReadWriteType = "WRITE_ONLY"
	// All events.
	// Experimental.
	ReadWriteType_ALL ReadWriteType = "ALL"
	// No events.
	// Experimental.
	ReadWriteType_NONE ReadWriteType = "NONE"
)

// Selecting an S3 bucket and an optional prefix to be logged for data events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3EventSelector := &s3EventSelector{
//   	bucket: bucket,
//
//   	// the properties below are optional
//   	objectPrefix: jsii.String("objectPrefix"),
//   }
//
// Experimental.
type S3EventSelector struct {
	// S3 bucket.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Data events for objects whose key matches this prefix will be logged.
	// Experimental.
	ObjectPrefix *string `field:"optional" json:"objectPrefix" yaml:"objectPrefix"`
}

// Cloud trail allows you to log events that happen in your AWS account For example:.
//
// import { CloudTrail } from '@aws-cdk/aws-cloudtrail'
//
// const cloudTrail = new CloudTrail(this, 'MyTrail');
//
// NOTE the above example creates an UNENCRYPTED bucket by default,
// If you are required to use an Encrypted bucket you can supply a preconfigured bucket
// via TrailProps.
//
// Example:
//   import cloudtrail "github.com/aws/aws-cdk-go/awscdk"
//
//
//   myKeyAlias := kms.alias.fromAliasName(this, jsii.String("myKey"), jsii.String("alias/aws/s3"))
//   trail := cloudtrail.NewTrail(this, jsii.String("myCloudTrail"), &trailProps{
//   	sendToCloudWatchLogs: jsii.Boolean(true),
//   	kmsKey: myKeyAlias,
//   })
//
// Experimental.
type Trail interface {
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
	// The CloudWatch log group to which CloudTrail events are sent.
	//
	// `undefined` if `sendToCloudWatchLogs` property is false.
	// Experimental.
	LogGroup() awslogs.ILogGroup
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
	// ARN of the CloudTrail trail i.e. arn:aws:cloudtrail:us-east-2:123456789012:trail/myCloudTrail.
	// Experimental.
	TrailArn() *string
	// ARN of the Amazon SNS topic that's associated with the CloudTrail trail, i.e. arn:aws:sns:us-east-2:123456789012:mySNSTopic.
	// Experimental.
	TrailSnsTopicArn() *string
	// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
	//
	// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
	//
	// This method adds an Event Selector for filtering events that match either S3 or Lambda function operations.
	//
	// Data events: These events provide insight into the resource operations performed on or within a resource.
	// These are also known as data plane operations.
	// Experimental.
	AddEventSelector(dataResourceType DataResourceType, dataResourceValues *[]*string, options *AddEventSelectorOptions)
	// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
	//
	// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
	//
	// This method adds a Lambda Data Event Selector for filtering events that match Lambda function operations.
	//
	// Data events: These events provide insight into the resource operations performed on or within a resource.
	// These are also known as data plane operations.
	// Experimental.
	AddLambdaEventSelector(handlers *[]awslambda.IFunction, options *AddEventSelectorOptions)
	// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
	//
	// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
	//
	// This method adds an S3 Data Event Selector for filtering events that match S3 operations.
	//
	// Data events: These events provide insight into the resource operations performed on or within a resource.
	// These are also known as data plane operations.
	// Experimental.
	AddS3EventSelector(s3Selector *[]*S3EventSelector, options *AddEventSelectorOptions)
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
	// Log all Lamda data events for all lambda functions the account.
	// See: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
	//
	// Experimental.
	LogAllLambdaDataEvents(options *AddEventSelectorOptions)
	// Log all S3 data events for all objects for all buckets in the account.
	// See: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
	//
	// Experimental.
	LogAllS3DataEvents(options *AddEventSelectorOptions)
	// Create an event rule for when an event is recorded by any Trail in the account.
	//
	// Note that the event doesn't necessarily have to come from this Trail, it can
	// be captured from any one.
	//
	// Be sure to filter the event further down using an event pattern.
	// Deprecated: - use Trail.onEvent()
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
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

// The jsii proxy struct for Trail
type jsiiProxy_Trail struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Trail) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) TrailArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Trail) TrailSnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailSnsTopicArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewTrail(scope constructs.Construct, id *string, props *TrailProps) Trail {
	_init_.Initialize()

	j := jsiiProxy_Trail{}

	_jsii_.Create(
		"monocdk.aws_cloudtrail.Trail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTrail_Override(t Trail, scope constructs.Construct, id *string, props *TrailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudtrail.Trail",
		[]interface{}{scope, id, props},
		t,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Trail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.Trail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Trail_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.Trail",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Create an event rule for when an event is recorded by any Trail in the account.
//
// Note that the event doesn't necessarily have to come from this Trail, it can
// be captured from any one.
//
// Be sure to filter the event further down using an event pattern.
// Experimental.
func Trail_OnEvent(scope constructs.Construct, id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	_init_.Initialize()

	var returns awsevents.Rule

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.Trail",
		"onEvent",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trail) AddEventSelector(dataResourceType DataResourceType, dataResourceValues *[]*string, options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"addEventSelector",
		[]interface{}{dataResourceType, dataResourceValues, options},
	)
}

func (t *jsiiProxy_Trail) AddLambdaEventSelector(handlers *[]awslambda.IFunction, options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"addLambdaEventSelector",
		[]interface{}{handlers, options},
	)
}

func (t *jsiiProxy_Trail) AddS3EventSelector(s3Selector *[]*S3EventSelector, options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"addS3EventSelector",
		[]interface{}{s3Selector, options},
	)
}

func (t *jsiiProxy_Trail) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_Trail) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trail) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trail) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trail) LogAllLambdaDataEvents(options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"logAllLambdaDataEvents",
		[]interface{}{options},
	)
}

func (t *jsiiProxy_Trail) LogAllS3DataEvents(options *AddEventSelectorOptions) {
	_jsii_.InvokeVoid(
		t,
		"logAllS3DataEvents",
		[]interface{}{options},
	)
}

func (t *jsiiProxy_Trail) OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		t,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trail) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Trail) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_Trail) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trail) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Trail) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_Trail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trail) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an AWS CloudTrail trail.
//
// Example:
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &trailProps{
//   	// ...
//   	managementEvents: cloudtrail.readWriteType_READ_ONLY,
//   })
//
// Experimental.
type TrailProps struct {
	// The Amazon S3 bucket.
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Log Group to which CloudTrail to push logs to.
	//
	// Ignored if sendToCloudWatchLogs is set to false.
	// Experimental.
	CloudWatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudWatchLogGroup" yaml:"cloudWatchLogGroup"`
	// How long to retain logs in CloudWatchLogs.
	//
	// Ignored if sendToCloudWatchLogs is false or if cloudWatchLogGroup is set.
	// Experimental.
	CloudWatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudWatchLogsRetention" yaml:"cloudWatchLogsRetention"`
	// To determine whether a log file was modified, deleted, or unchanged after CloudTrail delivered it, you can use CloudTrail log file integrity validation.
	//
	// This feature is built using industry standard algorithms: SHA-256 for hashing and SHA-256 with RSA for digital signing.
	// This makes it computationally infeasible to modify, delete or forge CloudTrail log files without detection.
	// You can use the AWS CLI to validate the files in the location where CloudTrail delivered them.
	// Experimental.
	EnableFileValidation *bool `field:"optional" json:"enableFileValidation" yaml:"enableFileValidation"`
	// The AWS Key Management Service (AWS KMS) key ID that you want to use to encrypt CloudTrail logs.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// For most services, events are recorded in the region where the action occurred.
	//
	// For global services such as AWS Identity and Access Management (IAM), AWS STS, Amazon CloudFront, and Route 53,
	// events are delivered to any trail that includes global services, and are logged as occurring in US East (N. Virginia) Region.
	// Experimental.
	IncludeGlobalServiceEvents *bool `field:"optional" json:"includeGlobalServiceEvents" yaml:"includeGlobalServiceEvents"`
	// Whether or not this trail delivers log files from multiple regions to a single S3 bucket for a single account.
	// Experimental.
	IsMultiRegionTrail *bool `field:"optional" json:"isMultiRegionTrail" yaml:"isMultiRegionTrail"`
	// The AWS Key Management Service (AWS KMS) key ID that you want to use to encrypt CloudTrail logs.
	// Deprecated: - use encryptionKey instead.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// When an event occurs in your account, CloudTrail evaluates whether the event matches the settings for your trails.
	//
	// Only events that match your trail settings are delivered to your Amazon S3 bucket and Amazon CloudWatch Logs log group.
	//
	// This method sets the management configuration for this trail.
	//
	// Management events provide insight into management operations that are performed on resources in your AWS account.
	// These are also known as control plane operations.
	// Management events can also include non-API events that occur in your account.
	// For example, when a user logs in to your account, CloudTrail logs the ConsoleLogin event.
	// Experimental.
	ManagementEvents ReadWriteType `field:"optional" json:"managementEvents" yaml:"managementEvents"`
	// An Amazon S3 object key prefix that precedes the name of all log files.
	// Experimental.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// If CloudTrail pushes logs to CloudWatch Logs in addition to S3.
	//
	// Disabled for cost out of the box.
	// Experimental.
	SendToCloudWatchLogs *bool `field:"optional" json:"sendToCloudWatchLogs" yaml:"sendToCloudWatchLogs"`
	// SNS topic that is notified when new log files are published.
	// Experimental.
	SnsTopic awssns.ITopic `field:"optional" json:"snsTopic" yaml:"snsTopic"`
	// The name of the trail.
	//
	// We recommend customers do not set an explicit name.
	// Experimental.
	TrailName *string `field:"optional" json:"trailName" yaml:"trailName"`
}

