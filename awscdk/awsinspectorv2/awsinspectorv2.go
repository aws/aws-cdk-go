package awsinspectorv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspectorv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::InspectorV2::Filter`.
//
// Details about a filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFilter := awscdk.Aws_inspectorv2.NewCfnFilter(this, jsii.String("MyCfnFilter"), &cfnFilterProps{
//   	filterAction: jsii.String("filterAction"),
//   	filterCriteria: &filterCriteriaProperty{
//   		awsAccountId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		componentId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		componentType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceImageId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceSubnetId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceVpcId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageArchitecture: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageHash: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImagePushedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		ecrImageRegistry: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageRepositoryName: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageTags: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingArn: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingStatus: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		firstObservedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		inspectorScore: []interface{}{
//   			&numberFilterProperty{
//   				lowerInclusive: jsii.Number(123),
//   				upperInclusive: jsii.Number(123),
//   			},
//   		},
//   		lastObservedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		networkProtocol: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		portRange: []interface{}{
//   			&portRangeFilterProperty{
//   				beginInclusive: jsii.Number(123),
//   				endInclusive: jsii.Number(123),
//   			},
//   		},
//   		relatedVulnerabilities: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceTags: []interface{}{
//   			&mapFilterProperty{
//   				comparison: jsii.String("comparison"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		severity: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		title: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		updatedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		vendorSeverity: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerabilityId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerabilitySource: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerablePackages: []interface{}{
//   			&packageFilterProperty{
//   				architecture: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				epoch: &numberFilterProperty{
//   					lowerInclusive: jsii.Number(123),
//   					upperInclusive: jsii.Number(123),
//   				},
//   				name: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				release: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				sourceLayerHash: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				version: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   })
//
type CfnFilter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Number (ARN) associated with this filter.
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
	// A description of the filter.
	Description() *string
	SetDescription(val *string)
	// The action that is to be applied to the findings that match the filter.
	FilterAction() *string
	SetFilterAction(val *string)
	// Details on the filter criteria associated with this filter.
	FilterCriteria() interface{}
	SetFilterCriteria(val interface{})
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
	// The name of the filter.
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

// The jsii proxy struct for CfnFilter
type jsiiProxy_CfnFilter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFilter) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) FilterAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) FilterCriteria() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFilter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::InspectorV2::Filter`.
func NewCfnFilter(scope constructs.Construct, id *string, props *CfnFilterProps) CfnFilter {
	_init_.Initialize()

	j := jsiiProxy_CfnFilter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::InspectorV2::Filter`.
func NewCfnFilter_Override(c CfnFilter, scope constructs.Construct, id *string, props *CfnFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFilter) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetFilterAction(val *string) {
	_jsii_.Set(
		j,
		"filterAction",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetFilterCriteria(val interface{}) {
	_jsii_.Set(
		j,
		"filterCriteria",
		val,
	)
}

func (j *jsiiProxy_CfnFilter) SetName(val *string) {
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
func CfnFilter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFilter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
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
func CfnFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFilter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_inspectorv2.CfnFilter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFilter) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFilter) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFilter) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFilter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFilter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFilter) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFilter) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFilter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFilter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFilter) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains details on the time range used to filter findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateFilterProperty := &dateFilterProperty{
//   	endInclusive: jsii.Number(123),
//   	startInclusive: jsii.Number(123),
//   }
//
type CfnFilter_DateFilterProperty struct {
	// A timestamp representing the end of the time period filtered on.
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
	// A timestamp representing the start of the time period filtered on.
	StartInclusive *float64 `field:"optional" json:"startInclusive" yaml:"startInclusive"`
}

// Details on the criteria used to define the filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterCriteriaProperty := &filterCriteriaProperty{
//   	awsAccountId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	componentId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	componentType: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2InstanceImageId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2InstanceSubnetId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2InstanceVpcId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImageArchitecture: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImageHash: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImagePushedAt: []interface{}{
//   		&dateFilterProperty{
//   			endInclusive: jsii.Number(123),
//   			startInclusive: jsii.Number(123),
//   		},
//   	},
//   	ecrImageRegistry: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImageRepositoryName: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ecrImageTags: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	findingArn: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	findingStatus: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	findingType: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	firstObservedAt: []interface{}{
//   		&dateFilterProperty{
//   			endInclusive: jsii.Number(123),
//   			startInclusive: jsii.Number(123),
//   		},
//   	},
//   	inspectorScore: []interface{}{
//   		&numberFilterProperty{
//   			lowerInclusive: jsii.Number(123),
//   			upperInclusive: jsii.Number(123),
//   		},
//   	},
//   	lastObservedAt: []interface{}{
//   		&dateFilterProperty{
//   			endInclusive: jsii.Number(123),
//   			startInclusive: jsii.Number(123),
//   		},
//   	},
//   	networkProtocol: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	portRange: []interface{}{
//   		&portRangeFilterProperty{
//   			beginInclusive: jsii.Number(123),
//   			endInclusive: jsii.Number(123),
//   		},
//   	},
//   	relatedVulnerabilities: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	resourceId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	resourceTags: []interface{}{
//   		&mapFilterProperty{
//   			comparison: jsii.String("comparison"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	resourceType: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	severity: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	title: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	updatedAt: []interface{}{
//   		&dateFilterProperty{
//   			endInclusive: jsii.Number(123),
//   			startInclusive: jsii.Number(123),
//   		},
//   	},
//   	vendorSeverity: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vulnerabilityId: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vulnerabilitySource: []interface{}{
//   		&stringFilterProperty{
//   			comparison: jsii.String("comparison"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vulnerablePackages: []interface{}{
//   		&packageFilterProperty{
//   			architecture: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   			epoch: &numberFilterProperty{
//   				lowerInclusive: jsii.Number(123),
//   				upperInclusive: jsii.Number(123),
//   			},
//   			name: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   			release: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   			sourceLayerHash: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   			version: &stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnFilter_FilterCriteriaProperty struct {
	// Details of the AWS account IDs used to filter findings.
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Details of the component IDs used to filter findings.
	ComponentId interface{} `field:"optional" json:"componentId" yaml:"componentId"`
	// Details of the component types used to filter findings.
	ComponentType interface{} `field:"optional" json:"componentType" yaml:"componentType"`
	// Details of the Amazon EC2 instance image IDs used to filter findings.
	Ec2InstanceImageId interface{} `field:"optional" json:"ec2InstanceImageId" yaml:"ec2InstanceImageId"`
	// Details of the Amazon EC2 instance subnet IDs used to filter findings.
	Ec2InstanceSubnetId interface{} `field:"optional" json:"ec2InstanceSubnetId" yaml:"ec2InstanceSubnetId"`
	// Details of the Amazon EC2 instance VPC IDs used to filter findings.
	Ec2InstanceVpcId interface{} `field:"optional" json:"ec2InstanceVpcId" yaml:"ec2InstanceVpcId"`
	// Details of the Amazon ECR image architecture types used to filter findings.
	EcrImageArchitecture interface{} `field:"optional" json:"ecrImageArchitecture" yaml:"ecrImageArchitecture"`
	// Details of the Amazon ECR image hashes used to filter findings.
	EcrImageHash interface{} `field:"optional" json:"ecrImageHash" yaml:"ecrImageHash"`
	// Details on the Amazon ECR image push date and time used to filter findings.
	EcrImagePushedAt interface{} `field:"optional" json:"ecrImagePushedAt" yaml:"ecrImagePushedAt"`
	// Details on the Amazon ECR registry used to filter findings.
	EcrImageRegistry interface{} `field:"optional" json:"ecrImageRegistry" yaml:"ecrImageRegistry"`
	// Details on the name of the Amazon ECR repository used to filter findings.
	EcrImageRepositoryName interface{} `field:"optional" json:"ecrImageRepositoryName" yaml:"ecrImageRepositoryName"`
	// The tags attached to the Amazon ECR container image.
	EcrImageTags interface{} `field:"optional" json:"ecrImageTags" yaml:"ecrImageTags"`
	// Details on the finding ARNs used to filter findings.
	FindingArn interface{} `field:"optional" json:"findingArn" yaml:"findingArn"`
	// Details on the finding status types used to filter findings.
	FindingStatus interface{} `field:"optional" json:"findingStatus" yaml:"findingStatus"`
	// Details on the finding types used to filter findings.
	FindingType interface{} `field:"optional" json:"findingType" yaml:"findingType"`
	// Details on the date and time a finding was first seen used to filter findings.
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// The Amazon Inspector score to filter on.
	InspectorScore interface{} `field:"optional" json:"inspectorScore" yaml:"inspectorScore"`
	// Details on the date and time a finding was last seen used to filter findings.
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// Details on the ingress source addresses used to filter findings.
	NetworkProtocol interface{} `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// Details on the port ranges used to filter findings.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// Details on the related vulnerabilities used to filter findings.
	RelatedVulnerabilities interface{} `field:"optional" json:"relatedVulnerabilities" yaml:"relatedVulnerabilities"`
	// Details on the resource IDs used to filter findings.
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Details on the resource tags used to filter findings.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Details on the resource types used to filter findings.
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Details on the severity used to filter findings.
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// Details on the finding title used to filter findings.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// Details on the date and time a finding was last updated at used to filter findings.
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Details on the vendor severity used to filter findings.
	VendorSeverity interface{} `field:"optional" json:"vendorSeverity" yaml:"vendorSeverity"`
	// Details on the vulnerability ID used to filter findings.
	VulnerabilityId interface{} `field:"optional" json:"vulnerabilityId" yaml:"vulnerabilityId"`
	// Details on the vulnerability score to filter findings by.
	VulnerabilitySource interface{} `field:"optional" json:"vulnerabilitySource" yaml:"vulnerabilitySource"`
	// Details on the vulnerable packages used to filter findings.
	VulnerablePackages interface{} `field:"optional" json:"vulnerablePackages" yaml:"vulnerablePackages"`
}

// An object that describes details of a map filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mapFilterProperty := &mapFilterProperty{
//   	comparison: jsii.String("comparison"),
//
//   	// the properties below are optional
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnFilter_MapFilterProperty struct {
	// The operator to use when comparing values in the filter.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The tag key used in the filter.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value used in the filter.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// An object that describes the details of a number filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberFilterProperty := &numberFilterProperty{
//   	lowerInclusive: jsii.Number(123),
//   	upperInclusive: jsii.Number(123),
//   }
//
type CfnFilter_NumberFilterProperty struct {
	// The lowest number to be included in the filter.
	LowerInclusive *float64 `field:"optional" json:"lowerInclusive" yaml:"lowerInclusive"`
	// The highest number to be included in the filter.
	UpperInclusive *float64 `field:"optional" json:"upperInclusive" yaml:"upperInclusive"`
}

// Contains information on the details of a package filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   packageFilterProperty := &packageFilterProperty{
//   	architecture: &stringFilterProperty{
//   		comparison: jsii.String("comparison"),
//   		value: jsii.String("value"),
//   	},
//   	epoch: &numberFilterProperty{
//   		lowerInclusive: jsii.Number(123),
//   		upperInclusive: jsii.Number(123),
//   	},
//   	name: &stringFilterProperty{
//   		comparison: jsii.String("comparison"),
//   		value: jsii.String("value"),
//   	},
//   	release: &stringFilterProperty{
//   		comparison: jsii.String("comparison"),
//   		value: jsii.String("value"),
//   	},
//   	sourceLayerHash: &stringFilterProperty{
//   		comparison: jsii.String("comparison"),
//   		value: jsii.String("value"),
//   	},
//   	version: &stringFilterProperty{
//   		comparison: jsii.String("comparison"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnFilter_PackageFilterProperty struct {
	// An object that contains details on the package architecture type to filter on.
	Architecture interface{} `field:"optional" json:"architecture" yaml:"architecture"`
	// An object that contains details on the package epoch to filter on.
	Epoch interface{} `field:"optional" json:"epoch" yaml:"epoch"`
	// An object that contains details on the name of the package to filter on.
	Name interface{} `field:"optional" json:"name" yaml:"name"`
	// An object that contains details on the package release to filter on.
	Release interface{} `field:"optional" json:"release" yaml:"release"`
	// An object that contains details on the source layer hash to filter on.
	SourceLayerHash interface{} `field:"optional" json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// The package version to filter on.
	Version interface{} `field:"optional" json:"version" yaml:"version"`
}

// An object that describes the details of a port range filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRangeFilterProperty := &portRangeFilterProperty{
//   	beginInclusive: jsii.Number(123),
//   	endInclusive: jsii.Number(123),
//   }
//
type CfnFilter_PortRangeFilterProperty struct {
	// The port number the port range begins at.
	BeginInclusive *float64 `field:"optional" json:"beginInclusive" yaml:"beginInclusive"`
	// The port number the port range ends at.
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
}

// An object that describes the details of a string filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringFilterProperty := &stringFilterProperty{
//   	comparison: jsii.String("comparison"),
//   	value: jsii.String("value"),
//   }
//
type CfnFilter_StringFilterProperty struct {
	// The operator to use when comparing values in the filter.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The value to filter on.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Properties for defining a `CfnFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFilterProps := &cfnFilterProps{
//   	filterAction: jsii.String("filterAction"),
//   	filterCriteria: &filterCriteriaProperty{
//   		awsAccountId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		componentId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		componentType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceImageId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceSubnetId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ec2InstanceVpcId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageArchitecture: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageHash: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImagePushedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		ecrImageRegistry: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageRepositoryName: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		ecrImageTags: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingArn: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingStatus: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		findingType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		firstObservedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		inspectorScore: []interface{}{
//   			&numberFilterProperty{
//   				lowerInclusive: jsii.Number(123),
//   				upperInclusive: jsii.Number(123),
//   			},
//   		},
//   		lastObservedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		networkProtocol: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		portRange: []interface{}{
//   			&portRangeFilterProperty{
//   				beginInclusive: jsii.Number(123),
//   				endInclusive: jsii.Number(123),
//   			},
//   		},
//   		relatedVulnerabilities: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceTags: []interface{}{
//   			&mapFilterProperty{
//   				comparison: jsii.String("comparison"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		resourceType: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		severity: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		title: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		updatedAt: []interface{}{
//   			&dateFilterProperty{
//   				endInclusive: jsii.Number(123),
//   				startInclusive: jsii.Number(123),
//   			},
//   		},
//   		vendorSeverity: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerabilityId: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerabilitySource: []interface{}{
//   			&stringFilterProperty{
//   				comparison: jsii.String("comparison"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		vulnerablePackages: []interface{}{
//   			&packageFilterProperty{
//   				architecture: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				epoch: &numberFilterProperty{
//   					lowerInclusive: jsii.Number(123),
//   					upperInclusive: jsii.Number(123),
//   				},
//   				name: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				release: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				sourceLayerHash: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   				version: &stringFilterProperty{
//   					comparison: jsii.String("comparison"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnFilterProps struct {
	// The action that is to be applied to the findings that match the filter.
	FilterAction *string `field:"required" json:"filterAction" yaml:"filterAction"`
	// Details on the filter criteria associated with this filter.
	FilterCriteria interface{} `field:"required" json:"filterCriteria" yaml:"filterCriteria"`
	// The name of the filter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the filter.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

