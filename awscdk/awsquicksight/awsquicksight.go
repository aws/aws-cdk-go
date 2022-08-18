package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::QuickSight::Analysis`.
//
// Creates an analysis in Amazon QuickSight.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnalysis := awscdk.Aws_quicksight.NewCfnAnalysis(this, jsii.String("MyCfnAnalysis"), &cfnAnalysisProps{
//   	analysisId: jsii.String("analysisId"),
//   	awsAccountId: jsii.String("awsAccountId"),
//   	sourceEntity: &analysisSourceEntityProperty{
//   		sourceTemplate: &analysisSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	errors: []interface{}{
//   		&analysisErrorProperty{
//   			message: jsii.String("message"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	parameters: &parametersProperty{
//   		dateTimeParameters: []interface{}{
//   			&dateTimeParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		decimalParameters: []interface{}{
//   			&decimalParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		integerParameters: []interface{}{
//   			&integerParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		stringParameters: []interface{}{
//   			&stringParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	themeArn: jsii.String("themeArn"),
//   })
//
type CfnAnalysis interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID for the analysis that you're creating.
	//
	// This ID displays in the URL of the analysis.
	AnalysisId() *string
	SetAnalysisId(val *string)
	// The Amazon Resource Name (ARN) of the analysis.
	AttrArn() *string
	AttrCreatedTime() *string
	// The ARNs of the datasets of the analysis.
	AttrDataSetArns() *[]*string
	// The time that the analysis was last updated.
	AttrLastUpdatedTime() *string
	AttrSheets() awscdk.IResolvable
	AttrStatus() *string
	// The ID of the AWS account where you are creating an analysis.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::QuickSight::Analysis.Errors`.
	Errors() interface{}
	SetErrors(val interface{})
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
	// A descriptive name for the analysis that you're creating.
	//
	// This name displays for the analysis in the Amazon QuickSight console.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The parameter names and override values that you want to use.
	//
	// An analysis can have any parameter type, and some parameters might accept multiple values.
	Parameters() interface{}
	SetParameters(val interface{})
	// A structure that describes the principals and the resource-level permissions on an analysis.
	//
	// You can use the `Permissions` structure to grant permissions by providing a list of AWS Identity and Access Management (IAM) action information for each principal listed by Amazon Resource Name (ARN).
	//
	// To specify no permissions, omit `Permissions` .
	Permissions() interface{}
	SetPermissions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A source entity to use for the analysis that you're creating.
	//
	// This metadata structure contains details that describe a source template and one or more datasets.
	SourceEntity() interface{}
	SetSourceEntity(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the analysis.
	Tags() awscdk.TagManager
	// The ARN for the theme to apply to the analysis that you're creating.
	//
	// To see the theme in the Amazon QuickSight console, make sure that you have access to it.
	ThemeArn() *string
	SetThemeArn(val *string)
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

// The jsii proxy struct for CfnAnalysis
type jsiiProxy_CfnAnalysis struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAnalysis) AnalysisId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrDataSetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDataSetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrSheets() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrSheets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Errors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) SourceEntity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) ThemeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnalysis) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::Analysis`.
func NewCfnAnalysis(scope constructs.Construct, id *string, props *CfnAnalysisProps) CfnAnalysis {
	_init_.Initialize()

	j := jsiiProxy_CfnAnalysis{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::Analysis`.
func NewCfnAnalysis_Override(c CfnAnalysis, scope constructs.Construct, id *string, props *CfnAnalysisProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetAnalysisId(val *string) {
	_jsii_.Set(
		j,
		"analysisId",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetErrors(val interface{}) {
	_jsii_.Set(
		j,
		"errors",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetSourceEntity(val interface{}) {
	_jsii_.Set(
		j,
		"sourceEntity",
		val,
	)
}

func (j *jsiiProxy_CfnAnalysis) SetThemeArn(val *string) {
	_jsii_.Set(
		j,
		"themeArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAnalysis_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAnalysis_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
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
func CfnAnalysis_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnalysis_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnAnalysis",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnalysis) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAnalysis) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAnalysis) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAnalysis) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAnalysis) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAnalysis) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAnalysis) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAnalysis) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnalysis) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnalysis) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAnalysis) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAnalysis) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnalysis) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnalysis) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnalysis) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Analysis error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisErrorProperty := &analysisErrorProperty{
//   	message: jsii.String("message"),
//   	type: jsii.String("type"),
//   }
//
type CfnAnalysis_AnalysisErrorProperty struct {
	// The message associated with the analysis error.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The type of the analysis error.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// The source entity of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSourceEntityProperty := &analysisSourceEntityProperty{
//   	sourceTemplate: &analysisSourceTemplateProperty{
//   		arn: jsii.String("arn"),
//   		dataSetReferences: []interface{}{
//   			&dataSetReferenceProperty{
//   				dataSetArn: jsii.String("dataSetArn"),
//   				dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_AnalysisSourceEntityProperty struct {
	// The source template for the source entity of the analysis.
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

// The source template of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSourceTemplateProperty := &analysisSourceTemplateProperty{
//   	arn: jsii.String("arn"),
//   	dataSetReferences: []interface{}{
//   		&dataSetReferenceProperty{
//   			dataSetArn: jsii.String("dataSetArn"),
//   			dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   		},
//   	},
//   }
//
type CfnAnalysis_AnalysisSourceTemplateProperty struct {
	// The Amazon Resource Name (ARN) of the source template of an analysis.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The dataset references of the source template of an analysis.
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

// Dataset reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetReferenceProperty := &dataSetReferenceProperty{
//   	dataSetArn: jsii.String("dataSetArn"),
//   	dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   }
//
type CfnAnalysis_DataSetReferenceProperty struct {
	// Dataset Amazon Resource Name (ARN).
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Dataset placeholder.
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

// A date-time parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeParameterProperty := &dateTimeParameterProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnAnalysis_DateTimeParameterProperty struct {
	// A display name for the date-time parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the date-time parameter.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

// A decimal parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalParameterProperty := &decimalParameterProperty{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_DecimalParameterProperty struct {
	// A display name for the decimal parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the decimal parameter.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

// An integer parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerParameterProperty := &integerParameterProperty{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_IntegerParameterProperty struct {
	// The name of the integer parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the integer parameter.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

// A list of Amazon QuickSight parameters and the list's override values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersProperty := &parametersProperty{
//   	dateTimeParameters: []interface{}{
//   		&dateTimeParameterProperty{
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	decimalParameters: []interface{}{
//   		&decimalParameterProperty{
//   			name: jsii.String("name"),
//   			values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	integerParameters: []interface{}{
//   		&integerParameterProperty{
//   			name: jsii.String("name"),
//   			values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	stringParameters: []interface{}{
//   		&stringParameterProperty{
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_ParametersProperty struct {
	// The parameters that have a data type of date-time.
	DateTimeParameters interface{} `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// The parameters that have a data type of decimal.
	DecimalParameters interface{} `field:"optional" json:"decimalParameters" yaml:"decimalParameters"`
	// The parameters that have a data type of integer.
	IntegerParameters interface{} `field:"optional" json:"integerParameters" yaml:"integerParameters"`
	// The parameters that have a data type of string.
	StringParameters interface{} `field:"optional" json:"stringParameters" yaml:"stringParameters"`
}

// Permission for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePermissionProperty := &resourcePermissionProperty{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	principal: jsii.String("principal"),
//   }
//
type CfnAnalysis_ResourcePermissionProperty struct {
	// The IAM action to grant or revoke permissions on.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Name (ARN) of the principal. This can be one of the following:.
	//
	// - The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)
	// - The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)
	// - The ARN of an AWS account root: This is an IAM ARN rather than a Amazon QuickSight ARN. Use this option only to share resources (templates) across AWS accounts . (This is less common.)
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

// A *sheet* , which is an object that contains a set of visuals that are viewed together on one page in Amazon QuickSight.
//
// Every analysis and dashboard contains at least one sheet. Each sheet contains at least one visualization widget, for example a chart, pivot table, or narrative insight. Sheets can be associated with other components, such as controls, filters, and so on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetProperty := &sheetProperty{
//   	name: jsii.String("name"),
//   	sheetId: jsii.String("sheetId"),
//   }
//
type CfnAnalysis_SheetProperty struct {
	// The name of a sheet.
	//
	// This name is displayed on the sheet's tab in the Amazon QuickSight console.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The unique identifier associated with a sheet.
	SheetId *string `field:"optional" json:"sheetId" yaml:"sheetId"`
}

// A string parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringParameterProperty := &stringParameterProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnAnalysis_StringParameterProperty struct {
	// A display name for a string parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values of a string parameter.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

// Properties for defining a `CfnAnalysis`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnalysisProps := &cfnAnalysisProps{
//   	analysisId: jsii.String("analysisId"),
//   	awsAccountId: jsii.String("awsAccountId"),
//   	sourceEntity: &analysisSourceEntityProperty{
//   		sourceTemplate: &analysisSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	errors: []interface{}{
//   		&analysisErrorProperty{
//   			message: jsii.String("message"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	parameters: &parametersProperty{
//   		dateTimeParameters: []interface{}{
//   			&dateTimeParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		decimalParameters: []interface{}{
//   			&decimalParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		integerParameters: []interface{}{
//   			&integerParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		stringParameters: []interface{}{
//   			&stringParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	themeArn: jsii.String("themeArn"),
//   }
//
type CfnAnalysisProps struct {
	// The ID for the analysis that you're creating.
	//
	// This ID displays in the URL of the analysis.
	AnalysisId *string `field:"required" json:"analysisId" yaml:"analysisId"`
	// The ID of the AWS account where you are creating an analysis.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// A source entity to use for the analysis that you're creating.
	//
	// This metadata structure contains details that describe a source template and one or more datasets.
	SourceEntity interface{} `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// `AWS::QuickSight::Analysis.Errors`.
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// A descriptive name for the analysis that you're creating.
	//
	// This name displays for the analysis in the Amazon QuickSight console.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameter names and override values that you want to use.
	//
	// An analysis can have any parameter type, and some parameters might accept multiple values.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A structure that describes the principals and the resource-level permissions on an analysis.
	//
	// You can use the `Permissions` structure to grant permissions by providing a list of AWS Identity and Access Management (IAM) action information for each principal listed by Amazon Resource Name (ARN).
	//
	// To specify no permissions, omit `Permissions` .
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the analysis.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN for the theme to apply to the analysis that you're creating.
	//
	// To see the theme in the Amazon QuickSight console, make sure that you have access to it.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
}

// A CloudFormation `AWS::QuickSight::Dashboard`.
//
// Creates a dashboard from a template. To first create a template, see the `CreateTemplate` API operation.
//
// A dashboard is an entity in Amazon QuickSight that identifies Amazon QuickSight reports, created from analyses. You can share Amazon QuickSight dashboards. With the right permissions, you can create scheduled email reports from them. If you have the correct permissions, you can create a dashboard from a template that exists in a different AWS account .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDashboard := awscdk.Aws_quicksight.NewCfnDashboard(this, jsii.String("MyCfnDashboard"), &cfnDashboardProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	dashboardId: jsii.String("dashboardId"),
//   	sourceEntity: &dashboardSourceEntityProperty{
//   		sourceTemplate: &dashboardSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	dashboardPublishOptions: &dashboardPublishOptionsProperty{
//   		adHocFilteringOption: &adHocFilteringOptionProperty{
//   			availabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   		exportToCsvOption: &exportToCSVOptionProperty{
//   			availabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   		sheetControlsOption: &sheetControlsOptionProperty{
//   			visibilityState: jsii.String("visibilityState"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	parameters: &parametersProperty{
//   		dateTimeParameters: []interface{}{
//   			&dateTimeParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		decimalParameters: []interface{}{
//   			&decimalParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		integerParameters: []interface{}{
//   			&integerParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		stringParameters: []interface{}{
//   			&stringParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	themeArn: jsii.String("themeArn"),
//   	versionDescription: jsii.String("versionDescription"),
//   })
//
type CfnDashboard interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the dashboard.
	AttrArn() *string
	// The time this dashboard version was created.
	AttrCreatedTime() *string
	// The time that the dashboard was last published.
	AttrLastPublishedTime() *string
	// The time that the dashboard was last updated.
	AttrLastUpdatedTime() *string
	// The ID of the AWS account where you want to create the dashboard.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ID for the dashboard, also added to the IAM policy.
	DashboardId() *string
	SetDashboardId(val *string)
	// Options for publishing the dashboard when you create it:.
	//
	// - `AvailabilityStatus` for `AdHocFilteringOption` - This status can be either `ENABLED` or `DISABLED` . When this is set to `DISABLED` , Amazon QuickSight disables the left filter pane on the published dashboard, which can be used for ad hoc (one-time) filtering. This option is `ENABLED` by default.
	// - `AvailabilityStatus` for `ExportToCSVOption` - This status can be either `ENABLED` or `DISABLED` . The visual option to export data to .CSV format isn't enabled when this is set to `DISABLED` . This option is `ENABLED` by default.
	// - `VisibilityState` for `SheetControlsOption` - This visibility state can be either `COLLAPSED` or `EXPANDED` . This option is `COLLAPSED` by default.
	DashboardPublishOptions() interface{}
	SetDashboardPublishOptions(val interface{})
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
	// The display name of the dashboard.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The parameters for the creation of the dashboard, which you want to use to override the default settings.
	//
	// A dashboard can have any type of parameters, and some parameters might accept multiple values.
	Parameters() interface{}
	SetParameters(val interface{})
	// A structure that contains the permissions of the dashboard.
	//
	// You can use this structure for granting permissions by providing a list of IAM action information for each principal ARN.
	//
	// To specify no permissions, omit the permissions list.
	Permissions() interface{}
	SetPermissions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The entity that you are using as a source when you create the dashboard.
	//
	// In `SourceEntity` , you specify the type of object that you want to use. You can only create a dashboard from a template, so you use a `SourceTemplate` entity. If you need to create a dashboard from an analysis, first convert the analysis to a template by using the `CreateTemplate` API operation. For `SourceTemplate` , specify the Amazon Resource Name (ARN) of the source template. The `SourceTemplate` ARN can contain any AWS account; and any QuickSight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	SourceEntity() interface{}
	SetSourceEntity(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dashboard.
	Tags() awscdk.TagManager
	// The Amazon Resource Name (ARN) of the theme that is being used for this dashboard.
	//
	// If you add a value for this field, it overrides the value that is used in the source entity. The theme ARN must exist in the same AWS account where you create the dashboard.
	ThemeArn() *string
	SetThemeArn(val *string)
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
	// A description for the first version of the dashboard being created.
	VersionDescription() *string
	SetVersionDescription(val *string)
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

// The jsii proxy struct for CfnDashboard
type jsiiProxy_CfnDashboard struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDashboard) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AttrLastPublishedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastPublishedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) DashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) DashboardPublishOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashboardPublishOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) SourceEntity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) ThemeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboard) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::Dashboard`.
func NewCfnDashboard(scope constructs.Construct, id *string, props *CfnDashboardProps) CfnDashboard {
	_init_.Initialize()

	j := jsiiProxy_CfnDashboard{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::Dashboard`.
func NewCfnDashboard_Override(c CfnDashboard, scope constructs.Construct, id *string, props *CfnDashboardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDashboard) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetDashboardId(val *string) {
	_jsii_.Set(
		j,
		"dashboardId",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetDashboardPublishOptions(val interface{}) {
	_jsii_.Set(
		j,
		"dashboardPublishOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetSourceEntity(val interface{}) {
	_jsii_.Set(
		j,
		"sourceEntity",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetThemeArn(val *string) {
	_jsii_.Set(
		j,
		"themeArn",
		val,
	)
}

func (j *jsiiProxy_CfnDashboard) SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDashboard_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDashboard_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
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
func CfnDashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDashboard_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnDashboard",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDashboard) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDashboard) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDashboard) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDashboard) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDashboard) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDashboard) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDashboard) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDashboard) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDashboard) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDashboard) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDashboard) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Ad hoc (one-time) filtering option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adHocFilteringOptionProperty := &adHocFilteringOptionProperty{
//   	availabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_AdHocFilteringOptionProperty struct {
	// Availability status.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

// Dashboard publish options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardPublishOptionsProperty := &dashboardPublishOptionsProperty{
//   	adHocFilteringOption: &adHocFilteringOptionProperty{
//   		availabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	exportToCsvOption: &exportToCSVOptionProperty{
//   		availabilityStatus: jsii.String("availabilityStatus"),
//   	},
//   	sheetControlsOption: &sheetControlsOptionProperty{
//   		visibilityState: jsii.String("visibilityState"),
//   	},
//   }
//
type CfnDashboard_DashboardPublishOptionsProperty struct {
	// Ad hoc (one-time) filtering option.
	AdHocFilteringOption interface{} `field:"optional" json:"adHocFilteringOption" yaml:"adHocFilteringOption"`
	// Export to .csv option.
	ExportToCsvOption interface{} `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// Sheet controls option.
	SheetControlsOption interface{} `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
}

// Dashboard source entity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardSourceEntityProperty := &dashboardSourceEntityProperty{
//   	sourceTemplate: &dashboardSourceTemplateProperty{
//   		arn: jsii.String("arn"),
//   		dataSetReferences: []interface{}{
//   			&dataSetReferenceProperty{
//   				dataSetArn: jsii.String("dataSetArn"),
//   				dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_DashboardSourceEntityProperty struct {
	// Source template.
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

// Dashboard source template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardSourceTemplateProperty := &dashboardSourceTemplateProperty{
//   	arn: jsii.String("arn"),
//   	dataSetReferences: []interface{}{
//   		&dataSetReferenceProperty{
//   			dataSetArn: jsii.String("dataSetArn"),
//   			dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   		},
//   	},
//   }
//
type CfnDashboard_DashboardSourceTemplateProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Dataset references.
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

// Dataset reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetReferenceProperty := &dataSetReferenceProperty{
//   	dataSetArn: jsii.String("dataSetArn"),
//   	dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   }
//
type CfnDashboard_DataSetReferenceProperty struct {
	// Dataset Amazon Resource Name (ARN).
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Dataset placeholder.
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

// A date-time parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeParameterProperty := &dateTimeParameterProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnDashboard_DateTimeParameterProperty struct {
	// A display name for the date-time parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the date-time parameter.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

// A decimal parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalParameterProperty := &decimalParameterProperty{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnDashboard_DecimalParameterProperty struct {
	// A display name for the decimal parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the decimal parameter.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

// Export to .csv option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exportToCSVOptionProperty := &exportToCSVOptionProperty{
//   	availabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_ExportToCSVOptionProperty struct {
	// Availability status.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

// An integer parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerParameterProperty := &integerParameterProperty{
//   	name: jsii.String("name"),
//   	values: []interface{}{
//   		jsii.Number(123),
//   	},
//   }
//
type CfnDashboard_IntegerParameterProperty struct {
	// The name of the integer parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the integer parameter.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

// A list of Amazon QuickSight parameters and the list's override values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametersProperty := &parametersProperty{
//   	dateTimeParameters: []interface{}{
//   		&dateTimeParameterProperty{
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	decimalParameters: []interface{}{
//   		&decimalParameterProperty{
//   			name: jsii.String("name"),
//   			values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	integerParameters: []interface{}{
//   		&integerParameterProperty{
//   			name: jsii.String("name"),
//   			values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	stringParameters: []interface{}{
//   		&stringParameterProperty{
//   			name: jsii.String("name"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
type CfnDashboard_ParametersProperty struct {
	// The parameters that have a data type of date-time.
	DateTimeParameters interface{} `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// The parameters that have a data type of decimal.
	DecimalParameters interface{} `field:"optional" json:"decimalParameters" yaml:"decimalParameters"`
	// The parameters that have a data type of integer.
	IntegerParameters interface{} `field:"optional" json:"integerParameters" yaml:"integerParameters"`
	// The parameters that have a data type of string.
	StringParameters interface{} `field:"optional" json:"stringParameters" yaml:"stringParameters"`
}

// Permission for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePermissionProperty := &resourcePermissionProperty{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	principal: jsii.String("principal"),
//   }
//
type CfnDashboard_ResourcePermissionProperty struct {
	// The IAM action to grant or revoke permissions on.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Name (ARN) of the principal. This can be one of the following:.
	//
	// - The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)
	// - The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)
	// - The ARN of an AWS account root: This is an IAM ARN rather than a Amazon QuickSight ARN. Use this option only to share resources (templates) across AWS accounts . (This is less common.)
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

// Sheet controls option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetControlsOptionProperty := &sheetControlsOptionProperty{
//   	visibilityState: jsii.String("visibilityState"),
//   }
//
type CfnDashboard_SheetControlsOptionProperty struct {
	// Visibility state.
	VisibilityState *string `field:"optional" json:"visibilityState" yaml:"visibilityState"`
}

// A string parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringParameterProperty := &stringParameterProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnDashboard_StringParameterProperty struct {
	// A display name for a string parameter.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values of a string parameter.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

// Properties for defining a `CfnDashboard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDashboardProps := &cfnDashboardProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	dashboardId: jsii.String("dashboardId"),
//   	sourceEntity: &dashboardSourceEntityProperty{
//   		sourceTemplate: &dashboardSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	dashboardPublishOptions: &dashboardPublishOptionsProperty{
//   		adHocFilteringOption: &adHocFilteringOptionProperty{
//   			availabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   		exportToCsvOption: &exportToCSVOptionProperty{
//   			availabilityStatus: jsii.String("availabilityStatus"),
//   		},
//   		sheetControlsOption: &sheetControlsOptionProperty{
//   			visibilityState: jsii.String("visibilityState"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	parameters: &parametersProperty{
//   		dateTimeParameters: []interface{}{
//   			&dateTimeParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		decimalParameters: []interface{}{
//   			&decimalParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		integerParameters: []interface{}{
//   			&integerParameterProperty{
//   				name: jsii.String("name"),
//   				values: []interface{}{
//   					jsii.Number(123),
//   				},
//   			},
//   		},
//   		stringParameters: []interface{}{
//   			&stringParameterProperty{
//   				name: jsii.String("name"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	themeArn: jsii.String("themeArn"),
//   	versionDescription: jsii.String("versionDescription"),
//   }
//
type CfnDashboardProps struct {
	// The ID of the AWS account where you want to create the dashboard.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ID for the dashboard, also added to the IAM policy.
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
	// The entity that you are using as a source when you create the dashboard.
	//
	// In `SourceEntity` , you specify the type of object that you want to use. You can only create a dashboard from a template, so you use a `SourceTemplate` entity. If you need to create a dashboard from an analysis, first convert the analysis to a template by using the `CreateTemplate` API operation. For `SourceTemplate` , specify the Amazon Resource Name (ARN) of the source template. The `SourceTemplate` ARN can contain any AWS account; and any QuickSight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	SourceEntity interface{} `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// Options for publishing the dashboard when you create it:.
	//
	// - `AvailabilityStatus` for `AdHocFilteringOption` - This status can be either `ENABLED` or `DISABLED` . When this is set to `DISABLED` , Amazon QuickSight disables the left filter pane on the published dashboard, which can be used for ad hoc (one-time) filtering. This option is `ENABLED` by default.
	// - `AvailabilityStatus` for `ExportToCSVOption` - This status can be either `ENABLED` or `DISABLED` . The visual option to export data to .CSV format isn't enabled when this is set to `DISABLED` . This option is `ENABLED` by default.
	// - `VisibilityState` for `SheetControlsOption` - This visibility state can be either `COLLAPSED` or `EXPANDED` . This option is `COLLAPSED` by default.
	DashboardPublishOptions interface{} `field:"optional" json:"dashboardPublishOptions" yaml:"dashboardPublishOptions"`
	// The display name of the dashboard.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters for the creation of the dashboard, which you want to use to override the default settings.
	//
	// A dashboard can have any type of parameters, and some parameters might accept multiple values.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A structure that contains the permissions of the dashboard.
	//
	// You can use this structure for granting permissions by providing a list of IAM action information for each principal ARN.
	//
	// To specify no permissions, omit the permissions list.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dashboard.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the theme that is being used for this dashboard.
	//
	// If you add a value for this field, it overrides the value that is used in the source entity. The theme ARN must exist in the same AWS account where you create the dashboard.
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
	// A description for the first version of the dashboard being created.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

// A CloudFormation `AWS::QuickSight::DataSet`.
//
// Creates a dataset. This operation doesn't support datasets that include uploaded files as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSet := awscdk.Aws_quicksight.NewCfnDataSet(this, jsii.String("MyCfnDataSet"), &cfnDataSetProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	columnGroups: []interface{}{
//   		&columnGroupProperty{
//   			geoSpatialColumnGroup: &geoSpatialColumnGroupProperty{
//   				columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				countryCode: jsii.String("countryCode"),
//   			},
//   		},
//   	},
//   	columnLevelPermissionRules: []interface{}{
//   		&columnLevelPermissionRuleProperty{
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			principals: []*string{
//   				jsii.String("principals"),
//   			},
//   		},
//   	},
//   	dataSetId: jsii.String("dataSetId"),
//   	dataSetUsageConfiguration: &dataSetUsageConfigurationProperty{
//   		disableUseAsDirectQuerySource: jsii.Boolean(false),
//   		disableUseAsImportedSource: jsii.Boolean(false),
//   	},
//   	fieldFolders: map[string]interface{}{
//   		"fieldFoldersKey": &FieldFolderProperty{
//   			"columns": []*string{
//   				jsii.String("columns"),
//   			},
//   			"description": jsii.String("description"),
//   		},
//   	},
//   	importMode: jsii.String("importMode"),
//   	ingestionWaitPolicy: &ingestionWaitPolicyProperty{
//   		ingestionWaitTimeInHours: jsii.Number(123),
//   		waitForSpiceIngestion: jsii.Boolean(false),
//   	},
//   	logicalTableMap: map[string]interface{}{
//   		"logicalTableMapKey": &LogicalTableProperty{
//   			"alias": jsii.String("alias"),
//   			"source": &LogicalTableSourceProperty{
//   				"dataSetArn": jsii.String("dataSetArn"),
//   				"joinInstruction": &JoinInstructionProperty{
//   					"leftOperand": jsii.String("leftOperand"),
//   					"onClause": jsii.String("onClause"),
//   					"rightOperand": jsii.String("rightOperand"),
//   					"type": jsii.String("type"),
//
//   					// the properties below are optional
//   					"leftJoinKeyProperties": &JoinKeyPropertiesProperty{
//   						"uniqueKey": jsii.Boolean(false),
//   					},
//   					"rightJoinKeyProperties": &JoinKeyPropertiesProperty{
//   						"uniqueKey": jsii.Boolean(false),
//   					},
//   				},
//   				"physicalTableId": jsii.String("physicalTableId"),
//   			},
//
//   			// the properties below are optional
//   			"dataTransforms": []interface{}{
//   				&TransformOperationProperty{
//   					"castColumnTypeOperation": &CastColumnTypeOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnType": jsii.String("newColumnType"),
//
//   						// the properties below are optional
//   						"format": jsii.String("format"),
//   					},
//   					"createColumnsOperation": &CreateColumnsOperationProperty{
//   						"columns": []interface{}{
//   							&CalculatedColumnProperty{
//   								"columnId": jsii.String("columnId"),
//   								"columnName": jsii.String("columnName"),
//   								"expression": jsii.String("expression"),
//   							},
//   						},
//   					},
//   					"filterOperation": &FilterOperationProperty{
//   						"conditionExpression": jsii.String("conditionExpression"),
//   					},
//   					"projectOperation": &ProjectOperationProperty{
//   						"projectedColumns": []*string{
//   							jsii.String("projectedColumns"),
//   						},
//   					},
//   					"renameColumnOperation": &RenameColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnName": jsii.String("newColumnName"),
//   					},
//   					"tagColumnOperation": &TagColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"tags": []ColumnTagProperty{
//   							&ColumnTagProperty{
//   								"columnDescription": &ColumnDescriptionProperty{
//   									"text": jsii.String("text"),
//   								},
//   								"columnGeographicRole": jsii.String("columnGeographicRole"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	physicalTableMap: map[string]interface{}{
//   		"physicalTableMapKey": &PhysicalTableProperty{
//   			"customSql": &CustomSqlProperty{
//   				"columns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"name": jsii.String("name"),
//   				"sqlQuery": jsii.String("sqlQuery"),
//   			},
//   			"relationalTable": &RelationalTableProperty{
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//   				"name": jsii.String("name"),
//
//   				// the properties below are optional
//   				"catalog": jsii.String("catalog"),
//   				"schema": jsii.String("schema"),
//   			},
//   			"s3Source": &S3SourceProperty{
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				"uploadSettings": &UploadSettingsProperty{
//   					"containsHeader": jsii.Boolean(false),
//   					"delimiter": jsii.String("delimiter"),
//   					"format": jsii.String("format"),
//   					"startFromRow": jsii.Number(123),
//   					"textQualifier": jsii.String("textQualifier"),
//   				},
//   			},
//   		},
//   	},
//   	rowLevelPermissionDataSet: &rowLevelPermissionDataSetProperty{
//   		arn: jsii.String("arn"),
//   		permissionPolicy: jsii.String("permissionPolicy"),
//
//   		// the properties below are optional
//   		formatVersion: jsii.String("formatVersion"),
//   		namespace: jsii.String("namespace"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDataSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the dataset.
	AttrArn() *string
	AttrConsumedSpiceCapacityInBytes() awscdk.IResolvable
	// The time this dataset version was created.
	AttrCreatedTime() *string
	// The time this dataset version was last updated.
	AttrLastUpdatedTime() *string
	AttrOutputColumns() awscdk.IResolvable
	// The AWS account ID.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Groupings of columns that work together in certain Amazon QuickSight features.
	//
	// Currently, only geospatial hierarchy is supported.
	ColumnGroups() interface{}
	SetColumnGroups(val interface{})
	// A set of one or more definitions of a `ColumnLevelPermissionRule` .
	ColumnLevelPermissionRules() interface{}
	SetColumnLevelPermissionRules(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An ID for the dataset that you want to create.
	//
	// This ID is unique per AWS Region for each AWS account.
	DataSetId() *string
	SetDataSetId(val *string)
	// `AWS::QuickSight::DataSet.DataSetUsageConfiguration`.
	DataSetUsageConfiguration() interface{}
	SetDataSetUsageConfiguration(val interface{})
	// The folder that contains fields and nested subfolders for your dataset.
	FieldFolders() interface{}
	SetFieldFolders(val interface{})
	// Indicates whether you want to import the data into SPICE.
	ImportMode() *string
	SetImportMode(val *string)
	// The wait policy to use when creating or updating a Dataset.
	//
	// The default is to wait for SPICE ingestion to finish with timeout of 36 hours.
	IngestionWaitPolicy() interface{}
	SetIngestionWaitPolicy(val interface{})
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
	// Configures the combination and transformation of the data from the physical tables.
	LogicalTableMap() interface{}
	SetLogicalTableMap(val interface{})
	// The display name for the dataset.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A list of resource permissions on the dataset.
	Permissions() interface{}
	SetPermissions(val interface{})
	// Declares the physical tables that are available in the underlying data sources.
	PhysicalTableMap() interface{}
	SetPhysicalTableMap(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The row-level security configuration for the data that you want to create.
	RowLevelPermissionDataSet() interface{}
	SetRowLevelPermissionDataSet(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnDataSet
type jsiiProxy_CfnDataSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataSet) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AttrConsumedSpiceCapacityInBytes() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrConsumedSpiceCapacityInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AttrOutputColumns() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrOutputColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) ColumnGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) ColumnLevelPermissionRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnLevelPermissionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) DataSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) DataSetUsageConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSetUsageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) FieldFolders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) ImportMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) IngestionWaitPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingestionWaitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) LogicalTableMap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logicalTableMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) PhysicalTableMap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"physicalTableMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) RowLevelPermissionDataSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowLevelPermissionDataSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::DataSet`.
func NewCfnDataSet(scope constructs.Construct, id *string, props *CfnDataSetProps) CfnDataSet {
	_init_.Initialize()

	j := jsiiProxy_CfnDataSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::DataSet`.
func NewCfnDataSet_Override(c CfnDataSet, scope constructs.Construct, id *string, props *CfnDataSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSet) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetColumnGroups(val interface{}) {
	_jsii_.Set(
		j,
		"columnGroups",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetColumnLevelPermissionRules(val interface{}) {
	_jsii_.Set(
		j,
		"columnLevelPermissionRules",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetDataSetId(val *string) {
	_jsii_.Set(
		j,
		"dataSetId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetDataSetUsageConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"dataSetUsageConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetFieldFolders(val interface{}) {
	_jsii_.Set(
		j,
		"fieldFolders",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetImportMode(val *string) {
	_jsii_.Set(
		j,
		"importMode",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetIngestionWaitPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"ingestionWaitPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetLogicalTableMap(val interface{}) {
	_jsii_.Set(
		j,
		"logicalTableMap",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetPhysicalTableMap(val interface{}) {
	_jsii_.Set(
		j,
		"physicalTableMap",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet) SetRowLevelPermissionDataSet(val interface{}) {
	_jsii_.Set(
		j,
		"rowLevelPermissionDataSet",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
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
func CfnDataSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSet) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSet) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSet) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSet) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSet) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A calculated column for a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculatedColumnProperty := &calculatedColumnProperty{
//   	columnId: jsii.String("columnId"),
//   	columnName: jsii.String("columnName"),
//   	expression: jsii.String("expression"),
//   }
//
type CfnDataSet_CalculatedColumnProperty struct {
	// A unique ID to identify a calculated column.
	//
	// During a dataset update, if the column ID of a calculated column matches that of an existing calculated column, Amazon QuickSight preserves the existing calculated column.
	ColumnId *string `field:"required" json:"columnId" yaml:"columnId"`
	// Column name.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// An expression that defines the calculated column.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

// A transform operation that casts a column to a different type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   castColumnTypeOperationProperty := &castColumnTypeOperationProperty{
//   	columnName: jsii.String("columnName"),
//   	newColumnType: jsii.String("newColumnType"),
//
//   	// the properties below are optional
//   	format: jsii.String("format"),
//   }
//
type CfnDataSet_CastColumnTypeOperationProperty struct {
	// Column name.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// New column data type.
	NewColumnType *string `field:"required" json:"newColumnType" yaml:"newColumnType"`
	// When casting a column from string to datetime type, you can supply a string in a format supported by Amazon QuickSight to denote the source data format.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

// Metadata that contains a description for a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnDescriptionProperty := &columnDescriptionProperty{
//   	text: jsii.String("text"),
//   }
//
type CfnDataSet_ColumnDescriptionProperty struct {
	// The text of a description for a column.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

// Groupings of columns that work together in certain Amazon QuickSight features.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnGroupProperty := &columnGroupProperty{
//   	geoSpatialColumnGroup: &geoSpatialColumnGroupProperty{
//   		columns: []*string{
//   			jsii.String("columns"),
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		countryCode: jsii.String("countryCode"),
//   	},
//   }
//
type CfnDataSet_ColumnGroupProperty struct {
	// Geospatial column group that denotes a hierarchy.
	GeoSpatialColumnGroup interface{} `field:"optional" json:"geoSpatialColumnGroup" yaml:"geoSpatialColumnGroup"`
}

// A rule defined to grant access on one or more restricted columns.
//
// Each dataset can have multiple rules. To create a restricted column, you add it to one or more rules. Each rule must contain at least one column and at least one user or group. To be able to see a restricted column, a user or group needs to be added to a rule for that column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnLevelPermissionRuleProperty := &columnLevelPermissionRuleProperty{
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   }
//
type CfnDataSet_ColumnLevelPermissionRuleProperty struct {
	// An array of column names.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// An array of Amazon Resource Names (ARNs) for Amazon QuickSight users or groups.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

// A tag for a column in a `[TagColumnOperation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_TagColumnOperation.html)` structure. This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnTagProperty := &columnTagProperty{
//   	columnDescription: &columnDescriptionProperty{
//   		text: jsii.String("text"),
//   	},
//   	columnGeographicRole: jsii.String("columnGeographicRole"),
//   }
//
type CfnDataSet_ColumnTagProperty struct {
	// A description for a column.
	ColumnDescription interface{} `field:"optional" json:"columnDescription" yaml:"columnDescription"`
	// A geospatial role for a column.
	ColumnGeographicRole *string `field:"optional" json:"columnGeographicRole" yaml:"columnGeographicRole"`
}

// A transform operation that creates calculated columns.
//
// Columns created in one such operation form a lexical closure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createColumnsOperationProperty := &createColumnsOperationProperty{
//   	columns: []interface{}{
//   		&calculatedColumnProperty{
//   			columnId: jsii.String("columnId"),
//   			columnName: jsii.String("columnName"),
//   			expression: jsii.String("expression"),
//   		},
//   	},
//   }
//
type CfnDataSet_CreateColumnsOperationProperty struct {
	// Calculated columns to create.
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
}

// A physical table type built from the results of the custom SQL query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customSqlProperty := &customSqlProperty{
//   	columns: []interface{}{
//   		&inputColumnProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	dataSourceArn: jsii.String("dataSourceArn"),
//   	name: jsii.String("name"),
//   	sqlQuery: jsii.String("sqlQuery"),
//   }
//
type CfnDataSet_CustomSqlProperty struct {
	// The column schema from the SQL query result set.
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// The Amazon Resource Name (ARN) of the data source.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// A display name for the SQL query result.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The SQL query.
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetUsageConfigurationProperty := &dataSetUsageConfigurationProperty{
//   	disableUseAsDirectQuerySource: jsii.Boolean(false),
//   	disableUseAsImportedSource: jsii.Boolean(false),
//   }
//
type CfnDataSet_DataSetUsageConfigurationProperty struct {
	// `CfnDataSet.DataSetUsageConfigurationProperty.DisableUseAsDirectQuerySource`.
	DisableUseAsDirectQuerySource interface{} `field:"optional" json:"disableUseAsDirectQuerySource" yaml:"disableUseAsDirectQuerySource"`
	// `CfnDataSet.DataSetUsageConfigurationProperty.DisableUseAsImportedSource`.
	DisableUseAsImportedSource interface{} `field:"optional" json:"disableUseAsImportedSource" yaml:"disableUseAsImportedSource"`
}

// A FieldFolder element is a folder that contains fields and nested subfolders.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldFolderProperty := &fieldFolderProperty{
//   	columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	description: jsii.String("description"),
//   }
//
type CfnDataSet_FieldFolderProperty struct {
	// A folder has a list of columns.
	//
	// A column can only be in one folder.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// The description for a field folder.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// A transform operation that filters rows based on a condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterOperationProperty := &filterOperationProperty{
//   	conditionExpression: jsii.String("conditionExpression"),
//   }
//
type CfnDataSet_FilterOperationProperty struct {
	// An expression that must evaluate to a Boolean value.
	//
	// Rows for which the expression evaluates to true are kept in the dataset.
	ConditionExpression *string `field:"required" json:"conditionExpression" yaml:"conditionExpression"`
}

// Geospatial column group that denotes a hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geoSpatialColumnGroupProperty := &geoSpatialColumnGroupProperty{
//   	columns: []*string{
//   		jsii.String("columns"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	countryCode: jsii.String("countryCode"),
//   }
//
type CfnDataSet_GeoSpatialColumnGroupProperty struct {
	// Columns in this hierarchy.
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// A display name for the hierarchy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Country code.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
}

// The wait policy to use when creating or updating a Dataset.
//
// The default is to wait for SPICE ingestion to finish with timeout of 36 hours.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestionWaitPolicyProperty := &ingestionWaitPolicyProperty{
//   	ingestionWaitTimeInHours: jsii.Number(123),
//   	waitForSpiceIngestion: jsii.Boolean(false),
//   }
//
type CfnDataSet_IngestionWaitPolicyProperty struct {
	// The maximum time (in hours) to wait for Ingestion to complete.
	//
	// Default timeout is 36 hours. Applicable only when `DataSetImportMode` mode is set to SPICE and `WaitForSpiceIngestion` is set to true.
	IngestionWaitTimeInHours *float64 `field:"optional" json:"ingestionWaitTimeInHours" yaml:"ingestionWaitTimeInHours"`
	// Wait for SPICE ingestion to finish to mark dataset creation or update as successful.
	//
	// Default (true). Applicable only when `DataSetImportMode` mode is set to SPICE.
	WaitForSpiceIngestion interface{} `field:"optional" json:"waitForSpiceIngestion" yaml:"waitForSpiceIngestion"`
}

// Metadata for a column that is used as the input of a transform operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputColumnProperty := &inputColumnProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   }
//
type CfnDataSet_InputColumnProperty struct {
	// The name of this column in the underlying data source.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the column.
	Type *string `field:"required" json:"type" yaml:"type"`
}

// The instructions associated with a join.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   joinInstructionProperty := &joinInstructionProperty{
//   	leftOperand: jsii.String("leftOperand"),
//   	onClause: jsii.String("onClause"),
//   	rightOperand: jsii.String("rightOperand"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	leftJoinKeyProperties: &joinKeyPropertiesProperty{
//   		uniqueKey: jsii.Boolean(false),
//   	},
//   	rightJoinKeyProperties: &joinKeyPropertiesProperty{
//   		uniqueKey: jsii.Boolean(false),
//   	},
//   }
//
type CfnDataSet_JoinInstructionProperty struct {
	// The operand on the left side of a join.
	LeftOperand *string `field:"required" json:"leftOperand" yaml:"leftOperand"`
	// The join instructions provided in the `ON` clause of a join.
	OnClause *string `field:"required" json:"onClause" yaml:"onClause"`
	// The operand on the right side of a join.
	RightOperand *string `field:"required" json:"rightOperand" yaml:"rightOperand"`
	// The type of join that it is.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Join key properties of the left operand.
	LeftJoinKeyProperties interface{} `field:"optional" json:"leftJoinKeyProperties" yaml:"leftJoinKeyProperties"`
	// Join key properties of the right operand.
	RightJoinKeyProperties interface{} `field:"optional" json:"rightJoinKeyProperties" yaml:"rightJoinKeyProperties"`
}

// Properties associated with the columns participating in a join.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   joinKeyPropertiesProperty := &joinKeyPropertiesProperty{
//   	uniqueKey: jsii.Boolean(false),
//   }
//
type CfnDataSet_JoinKeyPropertiesProperty struct {
	// A value that indicates that a row in a table is uniquely identified by the columns in a join key.
	//
	// This is used by Amazon QuickSight to optimize query performance.
	UniqueKey interface{} `field:"optional" json:"uniqueKey" yaml:"uniqueKey"`
}

// A *logical table* is a unit that joins and that data transformations operate on.
//
// A logical table has a source, which can be either a physical table or result of a join. When a logical table points to a physical table, the logical table acts as a mutable copy of that physical table through transform operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logicalTableProperty := &logicalTableProperty{
//   	alias: jsii.String("alias"),
//   	source: &logicalTableSourceProperty{
//   		dataSetArn: jsii.String("dataSetArn"),
//   		joinInstruction: &joinInstructionProperty{
//   			leftOperand: jsii.String("leftOperand"),
//   			onClause: jsii.String("onClause"),
//   			rightOperand: jsii.String("rightOperand"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			leftJoinKeyProperties: &joinKeyPropertiesProperty{
//   				uniqueKey: jsii.Boolean(false),
//   			},
//   			rightJoinKeyProperties: &joinKeyPropertiesProperty{
//   				uniqueKey: jsii.Boolean(false),
//   			},
//   		},
//   		physicalTableId: jsii.String("physicalTableId"),
//   	},
//
//   	// the properties below are optional
//   	dataTransforms: []interface{}{
//   		&transformOperationProperty{
//   			castColumnTypeOperation: &castColumnTypeOperationProperty{
//   				columnName: jsii.String("columnName"),
//   				newColumnType: jsii.String("newColumnType"),
//
//   				// the properties below are optional
//   				format: jsii.String("format"),
//   			},
//   			createColumnsOperation: &createColumnsOperationProperty{
//   				columns: []interface{}{
//   					&calculatedColumnProperty{
//   						columnId: jsii.String("columnId"),
//   						columnName: jsii.String("columnName"),
//   						expression: jsii.String("expression"),
//   					},
//   				},
//   			},
//   			filterOperation: &filterOperationProperty{
//   				conditionExpression: jsii.String("conditionExpression"),
//   			},
//   			projectOperation: &projectOperationProperty{
//   				projectedColumns: []*string{
//   					jsii.String("projectedColumns"),
//   				},
//   			},
//   			renameColumnOperation: &renameColumnOperationProperty{
//   				columnName: jsii.String("columnName"),
//   				newColumnName: jsii.String("newColumnName"),
//   			},
//   			tagColumnOperation: &tagColumnOperationProperty{
//   				columnName: jsii.String("columnName"),
//   				tags: []columnTagProperty{
//   					&columnTagProperty{
//   						columnDescription: &columnDescriptionProperty{
//   							text: jsii.String("text"),
//   						},
//   						columnGeographicRole: jsii.String("columnGeographicRole"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDataSet_LogicalTableProperty struct {
	// A display name for the logical table.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Source of this logical table.
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Transform operations that act on this logical table.
	DataTransforms interface{} `field:"optional" json:"dataTransforms" yaml:"dataTransforms"`
}

// Information about the source of a logical table.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logicalTableSourceProperty := &logicalTableSourceProperty{
//   	dataSetArn: jsii.String("dataSetArn"),
//   	joinInstruction: &joinInstructionProperty{
//   		leftOperand: jsii.String("leftOperand"),
//   		onClause: jsii.String("onClause"),
//   		rightOperand: jsii.String("rightOperand"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		leftJoinKeyProperties: &joinKeyPropertiesProperty{
//   			uniqueKey: jsii.Boolean(false),
//   		},
//   		rightJoinKeyProperties: &joinKeyPropertiesProperty{
//   			uniqueKey: jsii.Boolean(false),
//   		},
//   	},
//   	physicalTableId: jsii.String("physicalTableId"),
//   }
//
type CfnDataSet_LogicalTableSourceProperty struct {
	// `CfnDataSet.LogicalTableSourceProperty.DataSetArn`.
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// Specifies the result of a join of two logical tables.
	JoinInstruction interface{} `field:"optional" json:"joinInstruction" yaml:"joinInstruction"`
	// Physical table ID.
	PhysicalTableId *string `field:"optional" json:"physicalTableId" yaml:"physicalTableId"`
}

// Output column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputColumnProperty := &outputColumnProperty{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//   }
//
type CfnDataSet_OutputColumnProperty struct {
	// A description for a column.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A display name for the dataset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// A view of a data source that contains information about the shape of the data in the underlying source.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   physicalTableProperty := &physicalTableProperty{
//   	customSql: &customSqlProperty{
//   		columns: []interface{}{
//   			&inputColumnProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		dataSourceArn: jsii.String("dataSourceArn"),
//   		name: jsii.String("name"),
//   		sqlQuery: jsii.String("sqlQuery"),
//   	},
//   	relationalTable: &relationalTableProperty{
//   		dataSourceArn: jsii.String("dataSourceArn"),
//   		inputColumns: []interface{}{
//   			&inputColumnProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		catalog: jsii.String("catalog"),
//   		schema: jsii.String("schema"),
//   	},
//   	s3Source: &s3SourceProperty{
//   		dataSourceArn: jsii.String("dataSourceArn"),
//   		inputColumns: []interface{}{
//   			&inputColumnProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		uploadSettings: &uploadSettingsProperty{
//   			containsHeader: jsii.Boolean(false),
//   			delimiter: jsii.String("delimiter"),
//   			format: jsii.String("format"),
//   			startFromRow: jsii.Number(123),
//   			textQualifier: jsii.String("textQualifier"),
//   		},
//   	},
//   }
//
type CfnDataSet_PhysicalTableProperty struct {
	// A physical table type built from the results of the custom SQL query.
	CustomSql interface{} `field:"optional" json:"customSql" yaml:"customSql"`
	// A physical table type for relational data sources.
	RelationalTable interface{} `field:"optional" json:"relationalTable" yaml:"relationalTable"`
	// A physical table type for as S3 data source.
	S3Source interface{} `field:"optional" json:"s3Source" yaml:"s3Source"`
}

// A transform operation that projects columns.
//
// Operations that come after a projection can only refer to projected columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectOperationProperty := &projectOperationProperty{
//   	projectedColumns: []*string{
//   		jsii.String("projectedColumns"),
//   	},
//   }
//
type CfnDataSet_ProjectOperationProperty struct {
	// Projected columns.
	ProjectedColumns *[]*string `field:"required" json:"projectedColumns" yaml:"projectedColumns"`
}

// A physical table type for relational data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationalTableProperty := &relationalTableProperty{
//   	dataSourceArn: jsii.String("dataSourceArn"),
//   	inputColumns: []interface{}{
//   		&inputColumnProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	catalog: jsii.String("catalog"),
//   	schema: jsii.String("schema"),
//   }
//
type CfnDataSet_RelationalTableProperty struct {
	// The Amazon Resource Name (ARN) for the data source.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// The column schema of the table.
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// The name of the relational table.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnDataSet.RelationalTableProperty.Catalog`.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// The schema name.
	//
	// This name applies to certain relational database engines.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

// A transform operation that renames a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renameColumnOperationProperty := &renameColumnOperationProperty{
//   	columnName: jsii.String("columnName"),
//   	newColumnName: jsii.String("newColumnName"),
//   }
//
type CfnDataSet_RenameColumnOperationProperty struct {
	// The name of the column to be renamed.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The new name for the column.
	NewColumnName *string `field:"required" json:"newColumnName" yaml:"newColumnName"`
}

// Permission for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePermissionProperty := &resourcePermissionProperty{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	principal: jsii.String("principal"),
//   }
//
type CfnDataSet_ResourcePermissionProperty struct {
	// The IAM action to grand or revoke permisions on.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Name (ARN) of the principal. This can be one of the following:.
	//
	// - The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)
	// - The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)
	// - The ARN of an AWS account root: This is an IAM ARN rather than a Amazon QuickSight ARN. Use this option only to share resources (templates) across AWS accounts . (This is less common.)
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

// Information about a dataset that contains permissions for row-level security (RLS).
//
// The permissions dataset maps fields to users or groups. For more information, see [Using Row-Level Security (RLS) to Restrict Access to a Dataset](https://docs.aws.amazon.com/quicksight/latest/user/restrict-access-to-a-data-set-using-row-level-security.html) in the *Amazon QuickSight User Guide* .
//
// The option to deny permissions by setting `PermissionPolicy` to `DENY_ACCESS` is not supported for new RLS datasets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rowLevelPermissionDataSetProperty := &rowLevelPermissionDataSetProperty{
//   	arn: jsii.String("arn"),
//   	permissionPolicy: jsii.String("permissionPolicy"),
//
//   	// the properties below are optional
//   	formatVersion: jsii.String("formatVersion"),
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnDataSet_RowLevelPermissionDataSetProperty struct {
	// The Amazon Resource Name (ARN) of the dataset that contains permissions for RLS.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The type of permissions to use when interpreting the permissions for RLS.
	//
	// `DENY_ACCESS` is included for backward compatibility only.
	PermissionPolicy *string `field:"required" json:"permissionPolicy" yaml:"permissionPolicy"`
	// The user or group rules associated with the dataset that contains permissions for RLS.
	//
	// By default, `FormatVersion` is `VERSION_1` . When `FormatVersion` is `VERSION_1` , `UserName` and `GroupName` are required. When `FormatVersion` is `VERSION_2` , `UserARN` and `GroupARN` are required, and `Namespace` must not exist.
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// The namespace associated with the dataset that contains permissions for RLS.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

// A physical table type for an S3 data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceProperty := &s3SourceProperty{
//   	dataSourceArn: jsii.String("dataSourceArn"),
//   	inputColumns: []interface{}{
//   		&inputColumnProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	uploadSettings: &uploadSettingsProperty{
//   		containsHeader: jsii.Boolean(false),
//   		delimiter: jsii.String("delimiter"),
//   		format: jsii.String("format"),
//   		startFromRow: jsii.Number(123),
//   		textQualifier: jsii.String("textQualifier"),
//   	},
//   }
//
type CfnDataSet_S3SourceProperty struct {
	// The Amazon Resource Name (ARN) for the data source.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// A physical table type for an S3 data source.
	//
	// > For files that aren't JSON, only `STRING` data types are supported in input columns.
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// Information about the format for the S3 source file or files.
	UploadSettings interface{} `field:"optional" json:"uploadSettings" yaml:"uploadSettings"`
}

// A transform operation that tags a column with additional information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagColumnOperationProperty := &tagColumnOperationProperty{
//   	columnName: jsii.String("columnName"),
//   	tags: []columnTagProperty{
//   		&columnTagProperty{
//   			columnDescription: &columnDescriptionProperty{
//   				text: jsii.String("text"),
//   			},
//   			columnGeographicRole: jsii.String("columnGeographicRole"),
//   		},
//   	},
//   }
//
type CfnDataSet_TagColumnOperationProperty struct {
	// The column that this operation acts on.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The dataset column tag, currently only used for geospatial type tagging.
	//
	// > This is not tags for the AWS tagging feature.
	Tags *[]*CfnDataSet_ColumnTagProperty `field:"required" json:"tags" yaml:"tags"`
}

// A data transformation on a logical table.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformOperationProperty := &transformOperationProperty{
//   	castColumnTypeOperation: &castColumnTypeOperationProperty{
//   		columnName: jsii.String("columnName"),
//   		newColumnType: jsii.String("newColumnType"),
//
//   		// the properties below are optional
//   		format: jsii.String("format"),
//   	},
//   	createColumnsOperation: &createColumnsOperationProperty{
//   		columns: []interface{}{
//   			&calculatedColumnProperty{
//   				columnId: jsii.String("columnId"),
//   				columnName: jsii.String("columnName"),
//   				expression: jsii.String("expression"),
//   			},
//   		},
//   	},
//   	filterOperation: &filterOperationProperty{
//   		conditionExpression: jsii.String("conditionExpression"),
//   	},
//   	projectOperation: &projectOperationProperty{
//   		projectedColumns: []*string{
//   			jsii.String("projectedColumns"),
//   		},
//   	},
//   	renameColumnOperation: &renameColumnOperationProperty{
//   		columnName: jsii.String("columnName"),
//   		newColumnName: jsii.String("newColumnName"),
//   	},
//   	tagColumnOperation: &tagColumnOperationProperty{
//   		columnName: jsii.String("columnName"),
//   		tags: []columnTagProperty{
//   			&columnTagProperty{
//   				columnDescription: &columnDescriptionProperty{
//   					text: jsii.String("text"),
//   				},
//   				columnGeographicRole: jsii.String("columnGeographicRole"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataSet_TransformOperationProperty struct {
	// A transform operation that casts a column to a different type.
	CastColumnTypeOperation interface{} `field:"optional" json:"castColumnTypeOperation" yaml:"castColumnTypeOperation"`
	// An operation that creates calculated columns.
	//
	// Columns created in one such operation form a lexical closure.
	CreateColumnsOperation interface{} `field:"optional" json:"createColumnsOperation" yaml:"createColumnsOperation"`
	// An operation that filters rows based on some condition.
	FilterOperation interface{} `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// An operation that projects columns.
	//
	// Operations that come after a projection can only refer to projected columns.
	ProjectOperation interface{} `field:"optional" json:"projectOperation" yaml:"projectOperation"`
	// An operation that renames a column.
	RenameColumnOperation interface{} `field:"optional" json:"renameColumnOperation" yaml:"renameColumnOperation"`
	// An operation that tags a column with additional information.
	TagColumnOperation interface{} `field:"optional" json:"tagColumnOperation" yaml:"tagColumnOperation"`
}

// Information about the format for a source file or files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uploadSettingsProperty := &uploadSettingsProperty{
//   	containsHeader: jsii.Boolean(false),
//   	delimiter: jsii.String("delimiter"),
//   	format: jsii.String("format"),
//   	startFromRow: jsii.Number(123),
//   	textQualifier: jsii.String("textQualifier"),
//   }
//
type CfnDataSet_UploadSettingsProperty struct {
	// Whether the file has a header row, or the files each have a header row.
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// The delimiter between values in the file.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// File format.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// A row number to start reading data from.
	StartFromRow *float64 `field:"optional" json:"startFromRow" yaml:"startFromRow"`
	// Text qualifier.
	TextQualifier *string `field:"optional" json:"textQualifier" yaml:"textQualifier"`
}

// Properties for defining a `CfnDataSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSetProps := &cfnDataSetProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	columnGroups: []interface{}{
//   		&columnGroupProperty{
//   			geoSpatialColumnGroup: &geoSpatialColumnGroupProperty{
//   				columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				countryCode: jsii.String("countryCode"),
//   			},
//   		},
//   	},
//   	columnLevelPermissionRules: []interface{}{
//   		&columnLevelPermissionRuleProperty{
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			principals: []*string{
//   				jsii.String("principals"),
//   			},
//   		},
//   	},
//   	dataSetId: jsii.String("dataSetId"),
//   	dataSetUsageConfiguration: &dataSetUsageConfigurationProperty{
//   		disableUseAsDirectQuerySource: jsii.Boolean(false),
//   		disableUseAsImportedSource: jsii.Boolean(false),
//   	},
//   	fieldFolders: map[string]interface{}{
//   		"fieldFoldersKey": &FieldFolderProperty{
//   			"columns": []*string{
//   				jsii.String("columns"),
//   			},
//   			"description": jsii.String("description"),
//   		},
//   	},
//   	importMode: jsii.String("importMode"),
//   	ingestionWaitPolicy: &ingestionWaitPolicyProperty{
//   		ingestionWaitTimeInHours: jsii.Number(123),
//   		waitForSpiceIngestion: jsii.Boolean(false),
//   	},
//   	logicalTableMap: map[string]interface{}{
//   		"logicalTableMapKey": &LogicalTableProperty{
//   			"alias": jsii.String("alias"),
//   			"source": &LogicalTableSourceProperty{
//   				"dataSetArn": jsii.String("dataSetArn"),
//   				"joinInstruction": &JoinInstructionProperty{
//   					"leftOperand": jsii.String("leftOperand"),
//   					"onClause": jsii.String("onClause"),
//   					"rightOperand": jsii.String("rightOperand"),
//   					"type": jsii.String("type"),
//
//   					// the properties below are optional
//   					"leftJoinKeyProperties": &JoinKeyPropertiesProperty{
//   						"uniqueKey": jsii.Boolean(false),
//   					},
//   					"rightJoinKeyProperties": &JoinKeyPropertiesProperty{
//   						"uniqueKey": jsii.Boolean(false),
//   					},
//   				},
//   				"physicalTableId": jsii.String("physicalTableId"),
//   			},
//
//   			// the properties below are optional
//   			"dataTransforms": []interface{}{
//   				&TransformOperationProperty{
//   					"castColumnTypeOperation": &CastColumnTypeOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnType": jsii.String("newColumnType"),
//
//   						// the properties below are optional
//   						"format": jsii.String("format"),
//   					},
//   					"createColumnsOperation": &CreateColumnsOperationProperty{
//   						"columns": []interface{}{
//   							&CalculatedColumnProperty{
//   								"columnId": jsii.String("columnId"),
//   								"columnName": jsii.String("columnName"),
//   								"expression": jsii.String("expression"),
//   							},
//   						},
//   					},
//   					"filterOperation": &FilterOperationProperty{
//   						"conditionExpression": jsii.String("conditionExpression"),
//   					},
//   					"projectOperation": &ProjectOperationProperty{
//   						"projectedColumns": []*string{
//   							jsii.String("projectedColumns"),
//   						},
//   					},
//   					"renameColumnOperation": &RenameColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnName": jsii.String("newColumnName"),
//   					},
//   					"tagColumnOperation": &TagColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"tags": []ColumnTagProperty{
//   							&ColumnTagProperty{
//   								"columnDescription": &ColumnDescriptionProperty{
//   									"text": jsii.String("text"),
//   								},
//   								"columnGeographicRole": jsii.String("columnGeographicRole"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	physicalTableMap: map[string]interface{}{
//   		"physicalTableMapKey": &PhysicalTableProperty{
//   			"customSql": &CustomSqlProperty{
//   				"columns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"name": jsii.String("name"),
//   				"sqlQuery": jsii.String("sqlQuery"),
//   			},
//   			"relationalTable": &RelationalTableProperty{
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//   				"name": jsii.String("name"),
//
//   				// the properties below are optional
//   				"catalog": jsii.String("catalog"),
//   				"schema": jsii.String("schema"),
//   			},
//   			"s3Source": &S3SourceProperty{
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				"uploadSettings": &UploadSettingsProperty{
//   					"containsHeader": jsii.Boolean(false),
//   					"delimiter": jsii.String("delimiter"),
//   					"format": jsii.String("format"),
//   					"startFromRow": jsii.Number(123),
//   					"textQualifier": jsii.String("textQualifier"),
//   				},
//   			},
//   		},
//   	},
//   	rowLevelPermissionDataSet: &rowLevelPermissionDataSetProperty{
//   		arn: jsii.String("arn"),
//   		permissionPolicy: jsii.String("permissionPolicy"),
//
//   		// the properties below are optional
//   		formatVersion: jsii.String("formatVersion"),
//   		namespace: jsii.String("namespace"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataSetProps struct {
	// The AWS account ID.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Groupings of columns that work together in certain Amazon QuickSight features.
	//
	// Currently, only geospatial hierarchy is supported.
	ColumnGroups interface{} `field:"optional" json:"columnGroups" yaml:"columnGroups"`
	// A set of one or more definitions of a `ColumnLevelPermissionRule` .
	ColumnLevelPermissionRules interface{} `field:"optional" json:"columnLevelPermissionRules" yaml:"columnLevelPermissionRules"`
	// An ID for the dataset that you want to create.
	//
	// This ID is unique per AWS Region for each AWS account.
	DataSetId *string `field:"optional" json:"dataSetId" yaml:"dataSetId"`
	// `AWS::QuickSight::DataSet.DataSetUsageConfiguration`.
	DataSetUsageConfiguration interface{} `field:"optional" json:"dataSetUsageConfiguration" yaml:"dataSetUsageConfiguration"`
	// The folder that contains fields and nested subfolders for your dataset.
	FieldFolders interface{} `field:"optional" json:"fieldFolders" yaml:"fieldFolders"`
	// Indicates whether you want to import the data into SPICE.
	ImportMode *string `field:"optional" json:"importMode" yaml:"importMode"`
	// The wait policy to use when creating or updating a Dataset.
	//
	// The default is to wait for SPICE ingestion to finish with timeout of 36 hours.
	IngestionWaitPolicy interface{} `field:"optional" json:"ingestionWaitPolicy" yaml:"ingestionWaitPolicy"`
	// Configures the combination and transformation of the data from the physical tables.
	LogicalTableMap interface{} `field:"optional" json:"logicalTableMap" yaml:"logicalTableMap"`
	// The display name for the dataset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions on the dataset.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Declares the physical tables that are available in the underlying data sources.
	PhysicalTableMap interface{} `field:"optional" json:"physicalTableMap" yaml:"physicalTableMap"`
	// The row-level security configuration for the data that you want to create.
	RowLevelPermissionDataSet interface{} `field:"optional" json:"rowLevelPermissionDataSet" yaml:"rowLevelPermissionDataSet"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::QuickSight::DataSource`.
//
// Creates a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSource := awscdk.Aws_quicksight.NewCfnDataSource(this, jsii.String("MyCfnDataSource"), &cfnDataSourceProps{
//   	alternateDataSourceParameters: []interface{}{
//   		&dataSourceParametersProperty{
//   			amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			athenaParameters: &athenaParametersProperty{
//   				workGroup: jsii.String("workGroup"),
//   			},
//   			auroraParameters: &auroraParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mariaDbParameters: &mariaDbParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mySqlParameters: &mySqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			oracleParameters: &oracleParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			postgreSqlParameters: &postgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			prestoParameters: &prestoParametersProperty{
//   				catalog: jsii.String("catalog"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			rdsParameters: &rdsParametersProperty{
//   				database: jsii.String("database"),
//   				instanceId: jsii.String("instanceId"),
//   			},
//   			redshiftParameters: &redshiftParametersProperty{
//   				database: jsii.String("database"),
//
//   				// the properties below are optional
//   				clusterId: jsii.String("clusterId"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			s3Parameters: &s3ParametersProperty{
//   				manifestFileLocation: &manifestFileLocationProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			snowflakeParameters: &snowflakeParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				warehouse: jsii.String("warehouse"),
//   			},
//   			sparkParameters: &sparkParametersProperty{
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			sqlServerParameters: &sqlServerParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			teradataParameters: &teradataParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	awsAccountId: jsii.String("awsAccountId"),
//   	credentials: &dataSourceCredentialsProperty{
//   		copySourceArn: jsii.String("copySourceArn"),
//   		credentialPair: &credentialPairProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//
//   			// the properties below are optional
//   			alternateDataSourceParameters: []interface{}{
//   				&dataSourceParametersProperty{
//   					amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   						domain: jsii.String("domain"),
//   					},
//   					amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   						domain: jsii.String("domain"),
//   					},
//   					athenaParameters: &athenaParametersProperty{
//   						workGroup: jsii.String("workGroup"),
//   					},
//   					auroraParameters: &auroraParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					mariaDbParameters: &mariaDbParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					mySqlParameters: &mySqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					oracleParameters: &oracleParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					postgreSqlParameters: &postgreSqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					prestoParameters: &prestoParametersProperty{
//   						catalog: jsii.String("catalog"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					rdsParameters: &rdsParametersProperty{
//   						database: jsii.String("database"),
//   						instanceId: jsii.String("instanceId"),
//   					},
//   					redshiftParameters: &redshiftParametersProperty{
//   						database: jsii.String("database"),
//
//   						// the properties below are optional
//   						clusterId: jsii.String("clusterId"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					s3Parameters: &s3ParametersProperty{
//   						manifestFileLocation: &manifestFileLocationProperty{
//   							bucket: jsii.String("bucket"),
//   							key: jsii.String("key"),
//   						},
//   					},
//   					snowflakeParameters: &snowflakeParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						warehouse: jsii.String("warehouse"),
//   					},
//   					sparkParameters: &sparkParametersProperty{
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					sqlServerParameters: &sqlServerParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					teradataParameters: &teradataParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	dataSourceId: jsii.String("dataSourceId"),
//   	dataSourceParameters: &dataSourceParametersProperty{
//   		amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   			domain: jsii.String("domain"),
//   		},
//   		amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   			domain: jsii.String("domain"),
//   		},
//   		athenaParameters: &athenaParametersProperty{
//   			workGroup: jsii.String("workGroup"),
//   		},
//   		auroraParameters: &auroraParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		mariaDbParameters: &mariaDbParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		mySqlParameters: &mySqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		oracleParameters: &oracleParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		postgreSqlParameters: &postgreSqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		prestoParameters: &prestoParametersProperty{
//   			catalog: jsii.String("catalog"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		rdsParameters: &rdsParametersProperty{
//   			database: jsii.String("database"),
//   			instanceId: jsii.String("instanceId"),
//   		},
//   		redshiftParameters: &redshiftParametersProperty{
//   			database: jsii.String("database"),
//
//   			// the properties below are optional
//   			clusterId: jsii.String("clusterId"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		s3Parameters: &s3ParametersProperty{
//   			manifestFileLocation: &manifestFileLocationProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//   			},
//   		},
//   		snowflakeParameters: &snowflakeParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			warehouse: jsii.String("warehouse"),
//   		},
//   		sparkParameters: &sparkParametersProperty{
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		sqlServerParameters: &sqlServerParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		teradataParameters: &teradataParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   	errorInfo: &dataSourceErrorInfoProperty{
//   		message: jsii.String("message"),
//   		type: jsii.String("type"),
//   	},
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	sslProperties: &sslPropertiesProperty{
//   		disableSsl: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   	vpcConnectionProperties: &vpcConnectionPropertiesProperty{
//   		vpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   })
//
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A set of alternate data source parameters that you want to share for the credentials stored with this data source.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the credentials from this existing data source. If the `AlternateDataSourceParameters` list is null, the `Credentials` originally used with this `DataSourceParameters` are automatically allowed.
	AlternateDataSourceParameters() interface{}
	SetAlternateDataSourceParameters(val interface{})
	// The Amazon Resource Name (ARN) of the dataset.
	AttrArn() *string
	// The time that this data source was created.
	AttrCreatedTime() *string
	// The last time that this data source was updated.
	AttrLastUpdatedTime() *string
	// The HTTP status of the request.
	AttrStatus() *string
	// The AWS account ID.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The credentials Amazon QuickSight that uses to connect to your underlying source.
	//
	// Currently, only credentials based on user name and password are supported.
	Credentials() interface{}
	SetCredentials(val interface{})
	// An ID for the data source.
	//
	// This ID is unique per AWS Region for each AWS account.
	DataSourceId() *string
	SetDataSourceId(val *string)
	// The parameters that Amazon QuickSight uses to connect to your underlying source.
	DataSourceParameters() interface{}
	SetDataSourceParameters(val interface{})
	// Error information from the last update or the creation of the data source.
	ErrorInfo() interface{}
	SetErrorInfo(val interface{})
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
	// A display name for the data source.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A list of resource permissions on the data source.
	Permissions() interface{}
	SetPermissions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source.
	SslProperties() interface{}
	SetSslProperties(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
	Tags() awscdk.TagManager
	// The type of the data source. To return a list of all data sources, use `ListDataSources` .
	//
	// Use `AMAZON_ELASTICSEARCH` for Amazon OpenSearch Service.
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
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source.
	VpcConnectionProperties() interface{}
	SetVpcConnectionProperties(val interface{})
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

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataSource) AlternateDataSourceParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alternateDataSourceParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Credentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ErrorInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) SslProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) VpcConnectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConnectionProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::DataSource`.
func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::DataSource`.
func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource) SetAlternateDataSourceParameters(val interface{}) {
	_jsii_.Set(
		j,
		"alternateDataSourceParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDataSourceId(val *string) {
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDataSourceParameters(val interface{}) {
	_jsii_.Set(
		j,
		"dataSourceParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetErrorInfo(val interface{}) {
	_jsii_.Set(
		j,
		"errorInfo",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetSslProperties(val interface{}) {
	_jsii_.Set(
		j,
		"sslProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetVpcConnectionProperties(val interface{}) {
	_jsii_.Set(
		j,
		"vpcConnectionProperties",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
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
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The parameters for OpenSearch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonElasticsearchParametersProperty := &amazonElasticsearchParametersProperty{
//   	domain: jsii.String("domain"),
//   }
//
type CfnDataSource_AmazonElasticsearchParametersProperty struct {
	// The OpenSearch domain.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

// The parameters for OpenSearch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonOpenSearchParametersProperty := &amazonOpenSearchParametersProperty{
//   	domain: jsii.String("domain"),
//   }
//
type CfnDataSource_AmazonOpenSearchParametersProperty struct {
	// The OpenSearch domain.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

// Parameters for Amazon Athena.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   athenaParametersProperty := &athenaParametersProperty{
//   	workGroup: jsii.String("workGroup"),
//   }
//
type CfnDataSource_AthenaParametersProperty struct {
	// The workgroup that Amazon Athena uses.
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

// Parameters for Amazon Aurora.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auroraParametersProperty := &auroraParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_AuroraParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// Parameters for Amazon Aurora PostgreSQL-Compatible Edition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auroraPostgreSqlParametersProperty := &auroraPostgreSqlParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_AuroraPostgreSqlParametersProperty struct {
	// The Amazon Aurora PostgreSQL database to connect to.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The Amazon Aurora PostgreSQL-Compatible host to connect to.
	Host *string `field:"required" json:"host" yaml:"host"`
	// The port that Amazon Aurora PostgreSQL is listening on.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// The combination of user name and password that are used as credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   credentialPairProperty := &credentialPairProperty{
//   	password: jsii.String("password"),
//   	username: jsii.String("username"),
//
//   	// the properties below are optional
//   	alternateDataSourceParameters: []interface{}{
//   		&dataSourceParametersProperty{
//   			amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			athenaParameters: &athenaParametersProperty{
//   				workGroup: jsii.String("workGroup"),
//   			},
//   			auroraParameters: &auroraParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mariaDbParameters: &mariaDbParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mySqlParameters: &mySqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			oracleParameters: &oracleParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			postgreSqlParameters: &postgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			prestoParameters: &prestoParametersProperty{
//   				catalog: jsii.String("catalog"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			rdsParameters: &rdsParametersProperty{
//   				database: jsii.String("database"),
//   				instanceId: jsii.String("instanceId"),
//   			},
//   			redshiftParameters: &redshiftParametersProperty{
//   				database: jsii.String("database"),
//
//   				// the properties below are optional
//   				clusterId: jsii.String("clusterId"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			s3Parameters: &s3ParametersProperty{
//   				manifestFileLocation: &manifestFileLocationProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			snowflakeParameters: &snowflakeParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				warehouse: jsii.String("warehouse"),
//   			},
//   			sparkParameters: &sparkParametersProperty{
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			sqlServerParameters: &sqlServerParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			teradataParameters: &teradataParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnDataSource_CredentialPairProperty struct {
	// Password.
	Password *string `field:"required" json:"password" yaml:"password"`
	// User name.
	Username *string `field:"required" json:"username" yaml:"username"`
	// A set of alternate data source parameters that you want to share for these credentials.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the new data source with the existing credentials. If the `AlternateDataSourceParameters` list is null, the `DataSourceParameters` originally used with these `Credentials` is automatically allowed.
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
}

// Data source credentials.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceCredentialsProperty := &dataSourceCredentialsProperty{
//   	copySourceArn: jsii.String("copySourceArn"),
//   	credentialPair: &credentialPairProperty{
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//
//   		// the properties below are optional
//   		alternateDataSourceParameters: []interface{}{
//   			&dataSourceParametersProperty{
//   				amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   					domain: jsii.String("domain"),
//   				},
//   				amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   					domain: jsii.String("domain"),
//   				},
//   				athenaParameters: &athenaParametersProperty{
//   					workGroup: jsii.String("workGroup"),
//   				},
//   				auroraParameters: &auroraParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				mariaDbParameters: &mariaDbParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				mySqlParameters: &mySqlParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				oracleParameters: &oracleParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				postgreSqlParameters: &postgreSqlParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				prestoParameters: &prestoParametersProperty{
//   					catalog: jsii.String("catalog"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				rdsParameters: &rdsParametersProperty{
//   					database: jsii.String("database"),
//   					instanceId: jsii.String("instanceId"),
//   				},
//   				redshiftParameters: &redshiftParametersProperty{
//   					database: jsii.String("database"),
//
//   					// the properties below are optional
//   					clusterId: jsii.String("clusterId"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				s3Parameters: &s3ParametersProperty{
//   					manifestFileLocation: &manifestFileLocationProperty{
//   						bucket: jsii.String("bucket"),
//   						key: jsii.String("key"),
//   					},
//   				},
//   				snowflakeParameters: &snowflakeParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					warehouse: jsii.String("warehouse"),
//   				},
//   				sparkParameters: &sparkParametersProperty{
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				sqlServerParameters: &sqlServerParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   				teradataParameters: &teradataParametersProperty{
//   					database: jsii.String("database"),
//   					host: jsii.String("host"),
//   					port: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDataSource_DataSourceCredentialsProperty struct {
	// The Amazon Resource Name (ARN) of a data source that has the credential pair that you want to use.
	//
	// When `CopySourceArn` is not null, the credential pair from the data source in the ARN is used as the credentials for the `DataSourceCredentials` structure.
	CopySourceArn *string `field:"optional" json:"copySourceArn" yaml:"copySourceArn"`
	// Credential pair.
	//
	// For more information, see `[CredentialPair](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CredentialPair.html)` .
	CredentialPair interface{} `field:"optional" json:"credentialPair" yaml:"credentialPair"`
}

// Error information for the data source creation or update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceErrorInfoProperty := &dataSourceErrorInfoProperty{
//   	message: jsii.String("message"),
//   	type: jsii.String("type"),
//   }
//
type CfnDataSource_DataSourceErrorInfoProperty struct {
	// Error message.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Error type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

// The parameters that Amazon QuickSight uses to connect to your underlying data source.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceParametersProperty := &dataSourceParametersProperty{
//   	amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   		domain: jsii.String("domain"),
//   	},
//   	amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   		domain: jsii.String("domain"),
//   	},
//   	athenaParameters: &athenaParametersProperty{
//   		workGroup: jsii.String("workGroup"),
//   	},
//   	auroraParameters: &auroraParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	mariaDbParameters: &mariaDbParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	mySqlParameters: &mySqlParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	oracleParameters: &oracleParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	postgreSqlParameters: &postgreSqlParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	prestoParameters: &prestoParametersProperty{
//   		catalog: jsii.String("catalog"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	rdsParameters: &rdsParametersProperty{
//   		database: jsii.String("database"),
//   		instanceId: jsii.String("instanceId"),
//   	},
//   	redshiftParameters: &redshiftParametersProperty{
//   		database: jsii.String("database"),
//
//   		// the properties below are optional
//   		clusterId: jsii.String("clusterId"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	s3Parameters: &s3ParametersProperty{
//   		manifestFileLocation: &manifestFileLocationProperty{
//   			bucket: jsii.String("bucket"),
//   			key: jsii.String("key"),
//   		},
//   	},
//   	snowflakeParameters: &snowflakeParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		warehouse: jsii.String("warehouse"),
//   	},
//   	sparkParameters: &sparkParametersProperty{
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	sqlServerParameters: &sqlServerParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   	teradataParameters: &teradataParametersProperty{
//   		database: jsii.String("database"),
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//   	},
//   }
//
type CfnDataSource_DataSourceParametersProperty struct {
	// The parameters for OpenSearch.
	AmazonElasticsearchParameters interface{} `field:"optional" json:"amazonElasticsearchParameters" yaml:"amazonElasticsearchParameters"`
	// The parameters for OpenSearch.
	AmazonOpenSearchParameters interface{} `field:"optional" json:"amazonOpenSearchParameters" yaml:"amazonOpenSearchParameters"`
	// The parameters for Amazon Athena.
	AthenaParameters interface{} `field:"optional" json:"athenaParameters" yaml:"athenaParameters"`
	// The parameters for Amazon Aurora MySQL.
	AuroraParameters interface{} `field:"optional" json:"auroraParameters" yaml:"auroraParameters"`
	// The parameters for Amazon Aurora.
	AuroraPostgreSqlParameters interface{} `field:"optional" json:"auroraPostgreSqlParameters" yaml:"auroraPostgreSqlParameters"`
	// The parameters for MariaDB.
	MariaDbParameters interface{} `field:"optional" json:"mariaDbParameters" yaml:"mariaDbParameters"`
	// The parameters for MySQL.
	MySqlParameters interface{} `field:"optional" json:"mySqlParameters" yaml:"mySqlParameters"`
	// Oracle parameters.
	OracleParameters interface{} `field:"optional" json:"oracleParameters" yaml:"oracleParameters"`
	// The parameters for PostgreSQL.
	PostgreSqlParameters interface{} `field:"optional" json:"postgreSqlParameters" yaml:"postgreSqlParameters"`
	// The parameters for Presto.
	PrestoParameters interface{} `field:"optional" json:"prestoParameters" yaml:"prestoParameters"`
	// The parameters for Amazon RDS.
	RdsParameters interface{} `field:"optional" json:"rdsParameters" yaml:"rdsParameters"`
	// The parameters for Amazon Redshift.
	RedshiftParameters interface{} `field:"optional" json:"redshiftParameters" yaml:"redshiftParameters"`
	// The parameters for S3.
	S3Parameters interface{} `field:"optional" json:"s3Parameters" yaml:"s3Parameters"`
	// The parameters for Snowflake.
	SnowflakeParameters interface{} `field:"optional" json:"snowflakeParameters" yaml:"snowflakeParameters"`
	// The parameters for Spark.
	SparkParameters interface{} `field:"optional" json:"sparkParameters" yaml:"sparkParameters"`
	// The parameters for SQL Server.
	SqlServerParameters interface{} `field:"optional" json:"sqlServerParameters" yaml:"sqlServerParameters"`
	// The parameters for Teradata.
	TeradataParameters interface{} `field:"optional" json:"teradataParameters" yaml:"teradataParameters"`
}

// Amazon S3 manifest file location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manifestFileLocationProperty := &manifestFileLocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   }
//
type CfnDataSource_ManifestFileLocationProperty struct {
	// Amazon S3 bucket.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Amazon S3 key that identifies an object.
	Key *string `field:"required" json:"key" yaml:"key"`
}

// The parameters for MariaDB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mariaDbParametersProperty := &mariaDbParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_MariaDbParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// The parameters for MySQL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mySqlParametersProperty := &mySqlParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_MySqlParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// Oracle parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oracleParametersProperty := &oracleParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_OracleParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// The parameters for PostgreSQL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postgreSqlParametersProperty := &postgreSqlParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_PostgreSqlParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// The parameters for Presto.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prestoParametersProperty := &prestoParametersProperty{
//   	catalog: jsii.String("catalog"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_PrestoParametersProperty struct {
	// Catalog.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// The parameters for Amazon RDS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsParametersProperty := &rdsParametersProperty{
//   	database: jsii.String("database"),
//   	instanceId: jsii.String("instanceId"),
//   }
//
type CfnDataSource_RdsParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Instance ID.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

// The parameters for Amazon Redshift.
//
// The `ClusterId` field can be blank if `Host` and `Port` are both set. The `Host` and `Port` fields can be blank if the `ClusterId` field is set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftParametersProperty := &redshiftParametersProperty{
//   	database: jsii.String("database"),
//
//   	// the properties below are optional
//   	clusterId: jsii.String("clusterId"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_RedshiftParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Cluster ID.
	//
	// This field can be blank if the `Host` and `Port` are provided.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Host.
	//
	// This field can be blank if `ClusterId` is provided.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Port.
	//
	// This field can be blank if the `ClusterId` is provided.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

// Permission for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePermissionProperty := &resourcePermissionProperty{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	principal: jsii.String("principal"),
//   }
//
type CfnDataSource_ResourcePermissionProperty struct {
	// The IAM action to grant or revoke permissions on.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Name (ARN) of the principal. This can be one of the following:.
	//
	// - The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)
	// - The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)
	// - The ARN of an AWS account root: This is an IAM ARN rather than a Amazon QuickSight ARN. Use this option only to share resources (templates) across AWS accounts . (This is less common.)
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

// The parameters for S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ParametersProperty := &s3ParametersProperty{
//   	manifestFileLocation: &manifestFileLocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataSource_S3ParametersProperty struct {
	// Location of the Amazon S3 manifest file.
	//
	// This is NULL if the manifest file was uploaded into Amazon QuickSight.
	ManifestFileLocation interface{} `field:"required" json:"manifestFileLocation" yaml:"manifestFileLocation"`
}

// The parameters for Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeParametersProperty := &snowflakeParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	warehouse: jsii.String("warehouse"),
//   }
//
type CfnDataSource_SnowflakeParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Warehouse.
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
}

// The parameters for Spark.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sparkParametersProperty := &sparkParametersProperty{
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_SparkParametersProperty struct {
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// The parameters for SQL Server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqlServerParametersProperty := &sqlServerParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_SqlServerParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sslPropertiesProperty := &sslPropertiesProperty{
//   	disableSsl: jsii.Boolean(false),
//   }
//
type CfnDataSource_SslPropertiesProperty struct {
	// A Boolean option to control whether SSL should be disabled.
	DisableSsl interface{} `field:"optional" json:"disableSsl" yaml:"disableSsl"`
}

// The parameters for Teradata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   teradataParametersProperty := &teradataParametersProperty{
//   	database: jsii.String("database"),
//   	host: jsii.String("host"),
//   	port: jsii.Number(123),
//   }
//
type CfnDataSource_TeradataParametersProperty struct {
	// Database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Port.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

// VPC connection properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectionPropertiesProperty := &vpcConnectionPropertiesProperty{
//   	vpcConnectionArn: jsii.String("vpcConnectionArn"),
//   }
//
type CfnDataSource_VpcConnectionPropertiesProperty struct {
	// The Amazon Resource Name (ARN) for the VPC connection.
	VpcConnectionArn *string `field:"required" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

// Properties for defining a `CfnDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSourceProps := &cfnDataSourceProps{
//   	alternateDataSourceParameters: []interface{}{
//   		&dataSourceParametersProperty{
//   			amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   				domain: jsii.String("domain"),
//   			},
//   			athenaParameters: &athenaParametersProperty{
//   				workGroup: jsii.String("workGroup"),
//   			},
//   			auroraParameters: &auroraParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mariaDbParameters: &mariaDbParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			mySqlParameters: &mySqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			oracleParameters: &oracleParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			postgreSqlParameters: &postgreSqlParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			prestoParameters: &prestoParametersProperty{
//   				catalog: jsii.String("catalog"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			rdsParameters: &rdsParametersProperty{
//   				database: jsii.String("database"),
//   				instanceId: jsii.String("instanceId"),
//   			},
//   			redshiftParameters: &redshiftParametersProperty{
//   				database: jsii.String("database"),
//
//   				// the properties below are optional
//   				clusterId: jsii.String("clusterId"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			s3Parameters: &s3ParametersProperty{
//   				manifestFileLocation: &manifestFileLocationProperty{
//   					bucket: jsii.String("bucket"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			snowflakeParameters: &snowflakeParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				warehouse: jsii.String("warehouse"),
//   			},
//   			sparkParameters: &sparkParametersProperty{
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			sqlServerParameters: &sqlServerParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   			teradataParameters: &teradataParametersProperty{
//   				database: jsii.String("database"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	awsAccountId: jsii.String("awsAccountId"),
//   	credentials: &dataSourceCredentialsProperty{
//   		copySourceArn: jsii.String("copySourceArn"),
//   		credentialPair: &credentialPairProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//
//   			// the properties below are optional
//   			alternateDataSourceParameters: []interface{}{
//   				&dataSourceParametersProperty{
//   					amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   						domain: jsii.String("domain"),
//   					},
//   					amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   						domain: jsii.String("domain"),
//   					},
//   					athenaParameters: &athenaParametersProperty{
//   						workGroup: jsii.String("workGroup"),
//   					},
//   					auroraParameters: &auroraParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					mariaDbParameters: &mariaDbParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					mySqlParameters: &mySqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					oracleParameters: &oracleParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					postgreSqlParameters: &postgreSqlParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					prestoParameters: &prestoParametersProperty{
//   						catalog: jsii.String("catalog"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					rdsParameters: &rdsParametersProperty{
//   						database: jsii.String("database"),
//   						instanceId: jsii.String("instanceId"),
//   					},
//   					redshiftParameters: &redshiftParametersProperty{
//   						database: jsii.String("database"),
//
//   						// the properties below are optional
//   						clusterId: jsii.String("clusterId"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					s3Parameters: &s3ParametersProperty{
//   						manifestFileLocation: &manifestFileLocationProperty{
//   							bucket: jsii.String("bucket"),
//   							key: jsii.String("key"),
//   						},
//   					},
//   					snowflakeParameters: &snowflakeParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						warehouse: jsii.String("warehouse"),
//   					},
//   					sparkParameters: &sparkParametersProperty{
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					sqlServerParameters: &sqlServerParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   					teradataParameters: &teradataParametersProperty{
//   						database: jsii.String("database"),
//   						host: jsii.String("host"),
//   						port: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	dataSourceId: jsii.String("dataSourceId"),
//   	dataSourceParameters: &dataSourceParametersProperty{
//   		amazonElasticsearchParameters: &amazonElasticsearchParametersProperty{
//   			domain: jsii.String("domain"),
//   		},
//   		amazonOpenSearchParameters: &amazonOpenSearchParametersProperty{
//   			domain: jsii.String("domain"),
//   		},
//   		athenaParameters: &athenaParametersProperty{
//   			workGroup: jsii.String("workGroup"),
//   		},
//   		auroraParameters: &auroraParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		auroraPostgreSqlParameters: &auroraPostgreSqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		mariaDbParameters: &mariaDbParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		mySqlParameters: &mySqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		oracleParameters: &oracleParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		postgreSqlParameters: &postgreSqlParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		prestoParameters: &prestoParametersProperty{
//   			catalog: jsii.String("catalog"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		rdsParameters: &rdsParametersProperty{
//   			database: jsii.String("database"),
//   			instanceId: jsii.String("instanceId"),
//   		},
//   		redshiftParameters: &redshiftParametersProperty{
//   			database: jsii.String("database"),
//
//   			// the properties below are optional
//   			clusterId: jsii.String("clusterId"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		s3Parameters: &s3ParametersProperty{
//   			manifestFileLocation: &manifestFileLocationProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//   			},
//   		},
//   		snowflakeParameters: &snowflakeParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			warehouse: jsii.String("warehouse"),
//   		},
//   		sparkParameters: &sparkParametersProperty{
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		sqlServerParameters: &sqlServerParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   		teradataParameters: &teradataParametersProperty{
//   			database: jsii.String("database"),
//   			host: jsii.String("host"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   	errorInfo: &dataSourceErrorInfoProperty{
//   		message: jsii.String("message"),
//   		type: jsii.String("type"),
//   	},
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	sslProperties: &sslPropertiesProperty{
//   		disableSsl: jsii.Boolean(false),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   	vpcConnectionProperties: &vpcConnectionPropertiesProperty{
//   		vpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	},
//   }
//
type CfnDataSourceProps struct {
	// A set of alternate data source parameters that you want to share for the credentials stored with this data source.
	//
	// The credentials are applied in tandem with the data source parameters when you copy a data source by using a create or update request. The API operation compares the `DataSourceParameters` structure that's in the request with the structures in the `AlternateDataSourceParameters` allow list. If the structures are an exact match, the request is allowed to use the credentials from this existing data source. If the `AlternateDataSourceParameters` list is null, the `Credentials` originally used with this `DataSourceParameters` are automatically allowed.
	AlternateDataSourceParameters interface{} `field:"optional" json:"alternateDataSourceParameters" yaml:"alternateDataSourceParameters"`
	// The AWS account ID.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The credentials Amazon QuickSight that uses to connect to your underlying source.
	//
	// Currently, only credentials based on user name and password are supported.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// An ID for the data source.
	//
	// This ID is unique per AWS Region for each AWS account.
	DataSourceId *string `field:"optional" json:"dataSourceId" yaml:"dataSourceId"`
	// The parameters that Amazon QuickSight uses to connect to your underlying source.
	DataSourceParameters interface{} `field:"optional" json:"dataSourceParameters" yaml:"dataSourceParameters"`
	// Error information from the last update or the creation of the data source.
	ErrorInfo interface{} `field:"optional" json:"errorInfo" yaml:"errorInfo"`
	// A display name for the data source.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions on the data source.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source.
	SslProperties interface{} `field:"optional" json:"sslProperties" yaml:"sslProperties"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the data source. To return a list of all data sources, use `ListDataSources` .
	//
	// Use `AMAZON_ELASTICSEARCH` for Amazon OpenSearch Service.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source.
	VpcConnectionProperties interface{} `field:"optional" json:"vpcConnectionProperties" yaml:"vpcConnectionProperties"`
}

// A CloudFormation `AWS::QuickSight::Template`.
//
// Creates a template from an existing Amazon QuickSight analysis or template. You can use the resulting template to create a dashboard.
//
// A *template* is an entity in Amazon QuickSight that encapsulates the metadata required to create an analysis and that you can use to create s dashboard. A template adds a layer of abstraction by using placeholders to replace the dataset associated with the analysis. You can use templates to create dashboards by replacing dataset placeholders with datasets that follow the same schema that was used to create the source analysis and template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTemplate := awscdk.Aws_quicksight.NewCfnTemplate(this, jsii.String("MyCfnTemplate"), &cfnTemplateProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	sourceEntity: &templateSourceEntityProperty{
//   		sourceAnalysis: &templateSourceAnalysisProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   		sourceTemplate: &templateSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   		},
//   	},
//   	templateId: jsii.String("templateId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	versionDescription: jsii.String("versionDescription"),
//   })
//
type CfnTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the template.
	AttrArn() *string
	// The time this template was created.
	AttrCreatedTime() *string
	// The time this template was last updated.
	AttrLastUpdatedTime() *string
	// The ID for the AWS account that the group is in.
	//
	// You use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
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
	// A display name for the template.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A list of resource permissions to be set on the template.
	Permissions() interface{}
	SetPermissions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The entity that you are using as a source when you create the template.
	//
	// In `SourceEntity` , you specify the type of object you're using as source: `SourceTemplate` for a template or `SourceAnalysis` for an analysis. Both of these require an Amazon Resource Name (ARN). For `SourceTemplate` , specify the ARN of the source template. For `SourceAnalysis` , specify the ARN of the source analysis. The `SourceTemplate` ARN can contain any AWS account and any Amazon QuickSight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` or `SourceAnalysis` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	SourceEntity() interface{}
	SetSourceEntity(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.
	Tags() awscdk.TagManager
	// An ID for the template that you want to create.
	//
	// This template is unique per AWS Region ; in each AWS account.
	TemplateId() *string
	SetTemplateId(val *string)
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
	// A description of the current template version being created.
	//
	// This API operation creates the first version of the template. Every time `UpdateTemplate` is called, a new version is created. Each version of the template maintains a description of the version in the `VersionDescription` field.
	VersionDescription() *string
	SetVersionDescription(val *string)
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

// The jsii proxy struct for CfnTemplate
type jsiiProxy_CfnTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTemplate) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) SourceEntity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceEntity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) TemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTemplate) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::Template`.
func NewCfnTemplate(scope constructs.Construct, id *string, props *CfnTemplateProps) CfnTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::Template`.
func NewCfnTemplate_Override(c CfnTemplate, scope constructs.Construct, id *string, props *CfnTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTemplate) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetSourceEntity(val interface{}) {
	_jsii_.Set(
		j,
		"sourceEntity",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetTemplateId(val *string) {
	_jsii_.Set(
		j,
		"templateId",
		val,
	)
}

func (j *jsiiProxy_CfnTemplate) SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
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
func CfnTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTemplate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTemplate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTemplate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTemplate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Dataset reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetReferenceProperty := &dataSetReferenceProperty{
//   	dataSetArn: jsii.String("dataSetArn"),
//   	dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   }
//
type CfnTemplate_DataSetReferenceProperty struct {
	// Dataset Amazon Resource Name (ARN).
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Dataset placeholder.
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

// Permission for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePermissionProperty := &resourcePermissionProperty{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	principal: jsii.String("principal"),
//   }
//
type CfnTemplate_ResourcePermissionProperty struct {
	// The IAM action to grant or revoke permissions on.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Name (ARN) of the principal. This can be one of the following:.
	//
	// - The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)
	// - The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)
	// - The ARN of an AWS account root: This is an IAM ARN rather than a Amazon QuickSight ARN. Use this option only to share resources (templates) across AWS accounts . (This is less common.)
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

// The source analysis of the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSourceAnalysisProperty := &templateSourceAnalysisProperty{
//   	arn: jsii.String("arn"),
//   	dataSetReferences: []interface{}{
//   		&dataSetReferenceProperty{
//   			dataSetArn: jsii.String("dataSetArn"),
//   			dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   		},
//   	},
//   }
//
type CfnTemplate_TemplateSourceAnalysisProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// A structure containing information about the dataset references used as placeholders in the template.
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

// The source entity of the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSourceEntityProperty := &templateSourceEntityProperty{
//   	sourceAnalysis: &templateSourceAnalysisProperty{
//   		arn: jsii.String("arn"),
//   		dataSetReferences: []interface{}{
//   			&dataSetReferenceProperty{
//   				dataSetArn: jsii.String("dataSetArn"),
//   				dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   			},
//   		},
//   	},
//   	sourceTemplate: &templateSourceTemplateProperty{
//   		arn: jsii.String("arn"),
//   	},
//   }
//
type CfnTemplate_TemplateSourceEntityProperty struct {
	// The source analysis, if it is based on an analysis.
	SourceAnalysis interface{} `field:"optional" json:"sourceAnalysis" yaml:"sourceAnalysis"`
	// The source template, if it is based on an template.
	SourceTemplate interface{} `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

// The source template of the template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateSourceTemplateProperty := &templateSourceTemplateProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnTemplate_TemplateSourceTemplateProperty struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

// Properties for defining a `CfnTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTemplateProps := &cfnTemplateProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	sourceEntity: &templateSourceEntityProperty{
//   		sourceAnalysis: &templateSourceAnalysisProperty{
//   			arn: jsii.String("arn"),
//   			dataSetReferences: []interface{}{
//   				&dataSetReferenceProperty{
//   					dataSetArn: jsii.String("dataSetArn"),
//   					dataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   				},
//   			},
//   		},
//   		sourceTemplate: &templateSourceTemplateProperty{
//   			arn: jsii.String("arn"),
//   		},
//   	},
//   	templateId: jsii.String("templateId"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	versionDescription: jsii.String("versionDescription"),
//   }
//
type CfnTemplateProps struct {
	// The ID for the AWS account that the group is in.
	//
	// You use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The entity that you are using as a source when you create the template.
	//
	// In `SourceEntity` , you specify the type of object you're using as source: `SourceTemplate` for a template or `SourceAnalysis` for an analysis. Both of these require an Amazon Resource Name (ARN). For `SourceTemplate` , specify the ARN of the source template. For `SourceAnalysis` , specify the ARN of the source analysis. The `SourceTemplate` ARN can contain any AWS account and any Amazon QuickSight-supported AWS Region .
	//
	// Use the `DataSetReferences` entity within `SourceTemplate` or `SourceAnalysis` to list the replacement datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.
	SourceEntity interface{} `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// An ID for the template that you want to create.
	//
	// This template is unique per AWS Region ; in each AWS account.
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// A display name for the template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of resource permissions to be set on the template.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A description of the current template version being created.
	//
	// This API operation creates the first version of the template. Every time `UpdateTemplate` is called, a new version is created. Each version of the template maintains a description of the version in the `VersionDescription` field.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

// A CloudFormation `AWS::QuickSight::Theme`.
//
// Creates a theme.
//
// A *theme* is set of configuration options for color and layout. Themes apply to analyses and dashboards. For more information, see [Using Themes in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTheme := awscdk.Aws_quicksight.NewCfnTheme(this, jsii.String("MyCfnTheme"), &cfnThemeProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	themeId: jsii.String("themeId"),
//
//   	// the properties below are optional
//   	baseThemeId: jsii.String("baseThemeId"),
//   	configuration: &themeConfigurationProperty{
//   		dataColorPalette: &dataColorPaletteProperty{
//   			colors: []*string{
//   				jsii.String("colors"),
//   			},
//   			emptyFillColor: jsii.String("emptyFillColor"),
//   			minMaxGradient: []*string{
//   				jsii.String("minMaxGradient"),
//   			},
//   		},
//   		sheet: &sheetStyleProperty{
//   			tile: &tileStyleProperty{
//   				border: &borderStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   			},
//   			tileLayout: &tileLayoutStyleProperty{
//   				gutter: &gutterStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   				margin: &marginStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		typography: &typographyProperty{
//   			fontFamilies: []interface{}{
//   				&fontProperty{
//   					fontFamily: jsii.String("fontFamily"),
//   				},
//   			},
//   		},
//   		uiColorPalette: &uIColorPaletteProperty{
//   			accent: jsii.String("accent"),
//   			accentForeground: jsii.String("accentForeground"),
//   			danger: jsii.String("danger"),
//   			dangerForeground: jsii.String("dangerForeground"),
//   			dimension: jsii.String("dimension"),
//   			dimensionForeground: jsii.String("dimensionForeground"),
//   			measure: jsii.String("measure"),
//   			measureForeground: jsii.String("measureForeground"),
//   			primaryBackground: jsii.String("primaryBackground"),
//   			primaryForeground: jsii.String("primaryForeground"),
//   			secondaryBackground: jsii.String("secondaryBackground"),
//   			secondaryForeground: jsii.String("secondaryForeground"),
//   			success: jsii.String("success"),
//   			successForeground: jsii.String("successForeground"),
//   			warning: jsii.String("warning"),
//   			warningForeground: jsii.String("warningForeground"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	versionDescription: jsii.String("versionDescription"),
//   })
//
type CfnTheme interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the theme.
	AttrArn() *string
	// The time the theme was created.
	AttrCreatedTime() *string
	// The time the theme was last updated.
	AttrLastUpdatedTime() *string
	// Theme type.
	AttrType() *string
	// The ID of the AWS account where you want to store the new theme.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// The ID of the theme that a custom theme will inherit from.
	//
	// All themes inherit from one of the starting themes defined by Amazon QuickSight. For a list of the starting themes, use `ListThemes` or choose *Themes* from within an analysis.
	BaseThemeId() *string
	SetBaseThemeId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The theme configuration, which contains the theme display properties.
	Configuration() interface{}
	SetConfiguration(val interface{})
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
	// A display name for the theme.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A valid grouping of resource permissions to apply to the new theme.
	Permissions() interface{}
	SetPermissions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A map of the key-value pairs for the resource tag or tags that you want to add to the resource.
	Tags() awscdk.TagManager
	// An ID for the theme that you want to create.
	//
	// The theme ID is unique per AWS Region in each AWS account.
	ThemeId() *string
	SetThemeId(val *string)
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
	// A description of the first version of the theme that you're creating.
	//
	// Every time `UpdateTheme` is called, a new version is created. Each version of the theme has a description of the version in the `VersionDescription` field.
	VersionDescription() *string
	SetVersionDescription(val *string)
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

// The jsii proxy struct for CfnTheme
type jsiiProxy_CfnTheme struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTheme) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AttrType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) BaseThemeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseThemeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Configuration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) ThemeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTheme) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}


// Create a new `AWS::QuickSight::Theme`.
func NewCfnTheme(scope constructs.Construct, id *string, props *CfnThemeProps) CfnTheme {
	_init_.Initialize()

	j := jsiiProxy_CfnTheme{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::QuickSight::Theme`.
func NewCfnTheme_Override(c CfnTheme, scope constructs.Construct, id *string, props *CfnThemeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTheme) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetBaseThemeId(val *string) {
	_jsii_.Set(
		j,
		"baseThemeId",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetPermissions(val interface{}) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetThemeId(val *string) {
	_jsii_.Set(
		j,
		"themeId",
		val,
	)
}

func (j *jsiiProxy_CfnTheme) SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTheme_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTheme_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
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
func CfnTheme_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTheme_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnTheme",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTheme) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTheme) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTheme) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTheme) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTheme) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTheme) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTheme) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTheme) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTheme) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTheme) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTheme) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The display options for tile borders for visuals.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   borderStyleProperty := &borderStyleProperty{
//   	show: jsii.Boolean(false),
//   }
//
type CfnTheme_BorderStyleProperty struct {
	// The option to enable display of borders for visuals.
	Show interface{} `field:"optional" json:"show" yaml:"show"`
}

// The theme colors that are used for data colors in charts.
//
// The colors description is a hexadecimal color code that consists of six alphanumerical characters, prefixed with `#` , for example #37BFF5.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataColorPaletteProperty := &dataColorPaletteProperty{
//   	colors: []*string{
//   		jsii.String("colors"),
//   	},
//   	emptyFillColor: jsii.String("emptyFillColor"),
//   	minMaxGradient: []*string{
//   		jsii.String("minMaxGradient"),
//   	},
//   }
//
type CfnTheme_DataColorPaletteProperty struct {
	// The hexadecimal codes for the colors.
	Colors *[]*string `field:"optional" json:"colors" yaml:"colors"`
	// The hexadecimal code of a color that applies to charts where a lack of data is highlighted.
	EmptyFillColor *string `field:"optional" json:"emptyFillColor" yaml:"emptyFillColor"`
	// The minimum and maximum hexadecimal codes that describe a color gradient.
	MinMaxGradient *[]*string `field:"optional" json:"minMaxGradient" yaml:"minMaxGradient"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fontProperty := &fontProperty{
//   	fontFamily: jsii.String("fontFamily"),
//   }
//
type CfnTheme_FontProperty struct {
	// `CfnTheme.FontProperty.FontFamily`.
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
}

// The display options for gutter spacing between tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gutterStyleProperty := &gutterStyleProperty{
//   	show: jsii.Boolean(false),
//   }
//
type CfnTheme_GutterStyleProperty struct {
	// This Boolean value controls whether to display a gutter space between sheet tiles.
	Show interface{} `field:"optional" json:"show" yaml:"show"`
}

// The display options for margins around the outside edge of sheets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   marginStyleProperty := &marginStyleProperty{
//   	show: jsii.Boolean(false),
//   }
//
type CfnTheme_MarginStyleProperty struct {
	// This Boolean value controls whether to display sheet margins.
	Show interface{} `field:"optional" json:"show" yaml:"show"`
}

// Permission for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePermissionProperty := &resourcePermissionProperty{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	principal: jsii.String("principal"),
//   }
//
type CfnTheme_ResourcePermissionProperty struct {
	// The IAM action to grant or revoke permissions on.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Name (ARN) of the principal. This can be one of the following:.
	//
	// - The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)
	// - The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)
	// - The ARN of an AWS account root: This is an IAM ARN rather than a Amazon QuickSight ARN. Use this option only to share resources (templates) across AWS accounts . (This is less common.)
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

// The theme display options for sheets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetStyleProperty := &sheetStyleProperty{
//   	tile: &tileStyleProperty{
//   		border: &borderStyleProperty{
//   			show: jsii.Boolean(false),
//   		},
//   	},
//   	tileLayout: &tileLayoutStyleProperty{
//   		gutter: &gutterStyleProperty{
//   			show: jsii.Boolean(false),
//   		},
//   		margin: &marginStyleProperty{
//   			show: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnTheme_SheetStyleProperty struct {
	// The display options for tiles.
	Tile interface{} `field:"optional" json:"tile" yaml:"tile"`
	// The layout options for tiles.
	TileLayout interface{} `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

// The theme configuration.
//
// This configuration contains all of the display properties for a theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   themeConfigurationProperty := &themeConfigurationProperty{
//   	dataColorPalette: &dataColorPaletteProperty{
//   		colors: []*string{
//   			jsii.String("colors"),
//   		},
//   		emptyFillColor: jsii.String("emptyFillColor"),
//   		minMaxGradient: []*string{
//   			jsii.String("minMaxGradient"),
//   		},
//   	},
//   	sheet: &sheetStyleProperty{
//   		tile: &tileStyleProperty{
//   			border: &borderStyleProperty{
//   				show: jsii.Boolean(false),
//   			},
//   		},
//   		tileLayout: &tileLayoutStyleProperty{
//   			gutter: &gutterStyleProperty{
//   				show: jsii.Boolean(false),
//   			},
//   			margin: &marginStyleProperty{
//   				show: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	typography: &typographyProperty{
//   		fontFamilies: []interface{}{
//   			&fontProperty{
//   				fontFamily: jsii.String("fontFamily"),
//   			},
//   		},
//   	},
//   	uiColorPalette: &uIColorPaletteProperty{
//   		accent: jsii.String("accent"),
//   		accentForeground: jsii.String("accentForeground"),
//   		danger: jsii.String("danger"),
//   		dangerForeground: jsii.String("dangerForeground"),
//   		dimension: jsii.String("dimension"),
//   		dimensionForeground: jsii.String("dimensionForeground"),
//   		measure: jsii.String("measure"),
//   		measureForeground: jsii.String("measureForeground"),
//   		primaryBackground: jsii.String("primaryBackground"),
//   		primaryForeground: jsii.String("primaryForeground"),
//   		secondaryBackground: jsii.String("secondaryBackground"),
//   		secondaryForeground: jsii.String("secondaryForeground"),
//   		success: jsii.String("success"),
//   		successForeground: jsii.String("successForeground"),
//   		warning: jsii.String("warning"),
//   		warningForeground: jsii.String("warningForeground"),
//   	},
//   }
//
type CfnTheme_ThemeConfigurationProperty struct {
	// Color properties that apply to chart data colors.
	DataColorPalette interface{} `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// Display options related to sheets.
	Sheet interface{} `field:"optional" json:"sheet" yaml:"sheet"`
	// `CfnTheme.ThemeConfigurationProperty.Typography`.
	Typography interface{} `field:"optional" json:"typography" yaml:"typography"`
	// Color properties that apply to the UI and to charts, excluding the colors that apply to data.
	UiColorPalette interface{} `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

// The display options for the layout of tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tileLayoutStyleProperty := &tileLayoutStyleProperty{
//   	gutter: &gutterStyleProperty{
//   		show: jsii.Boolean(false),
//   	},
//   	margin: &marginStyleProperty{
//   		show: jsii.Boolean(false),
//   	},
//   }
//
type CfnTheme_TileLayoutStyleProperty struct {
	// The gutter settings that apply between tiles.
	Gutter interface{} `field:"optional" json:"gutter" yaml:"gutter"`
	// The margin settings that apply around the outside edge of sheets.
	Margin interface{} `field:"optional" json:"margin" yaml:"margin"`
}

// Display options related to tiles on a sheet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tileStyleProperty := &tileStyleProperty{
//   	border: &borderStyleProperty{
//   		show: jsii.Boolean(false),
//   	},
//   }
//
type CfnTheme_TileStyleProperty struct {
	// The border around a tile.
	Border interface{} `field:"optional" json:"border" yaml:"border"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typographyProperty := &typographyProperty{
//   	fontFamilies: []interface{}{
//   		&fontProperty{
//   			fontFamily: jsii.String("fontFamily"),
//   		},
//   	},
//   }
//
type CfnTheme_TypographyProperty struct {
	// `CfnTheme.TypographyProperty.FontFamilies`.
	FontFamilies interface{} `field:"optional" json:"fontFamilies" yaml:"fontFamilies"`
}

// The theme colors that apply to UI and to charts, excluding data colors.
//
// The colors description is a hexadecimal color code that consists of six alphanumerical characters, prefixed with `#` , for example #37BFF5. For more information, see [Using Themes in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html) in the *Amazon QuickSight User Guide.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uIColorPaletteProperty := &uIColorPaletteProperty{
//   	accent: jsii.String("accent"),
//   	accentForeground: jsii.String("accentForeground"),
//   	danger: jsii.String("danger"),
//   	dangerForeground: jsii.String("dangerForeground"),
//   	dimension: jsii.String("dimension"),
//   	dimensionForeground: jsii.String("dimensionForeground"),
//   	measure: jsii.String("measure"),
//   	measureForeground: jsii.String("measureForeground"),
//   	primaryBackground: jsii.String("primaryBackground"),
//   	primaryForeground: jsii.String("primaryForeground"),
//   	secondaryBackground: jsii.String("secondaryBackground"),
//   	secondaryForeground: jsii.String("secondaryForeground"),
//   	success: jsii.String("success"),
//   	successForeground: jsii.String("successForeground"),
//   	warning: jsii.String("warning"),
//   	warningForeground: jsii.String("warningForeground"),
//   }
//
type CfnTheme_UIColorPaletteProperty struct {
	// This color is that applies to selected states and buttons.
	Accent *string `field:"optional" json:"accent" yaml:"accent"`
	// The foreground color that applies to any text or other elements that appear over the accent color.
	AccentForeground *string `field:"optional" json:"accentForeground" yaml:"accentForeground"`
	// The color that applies to error messages.
	Danger *string `field:"optional" json:"danger" yaml:"danger"`
	// The foreground color that applies to any text or other elements that appear over the error color.
	DangerForeground *string `field:"optional" json:"dangerForeground" yaml:"dangerForeground"`
	// The color that applies to the names of fields that are identified as dimensions.
	Dimension *string `field:"optional" json:"dimension" yaml:"dimension"`
	// The foreground color that applies to any text or other elements that appear over the dimension color.
	DimensionForeground *string `field:"optional" json:"dimensionForeground" yaml:"dimensionForeground"`
	// The color that applies to the names of fields that are identified as measures.
	Measure *string `field:"optional" json:"measure" yaml:"measure"`
	// The foreground color that applies to any text or other elements that appear over the measure color.
	MeasureForeground *string `field:"optional" json:"measureForeground" yaml:"measureForeground"`
	// The background color that applies to visuals and other high emphasis UI.
	PrimaryBackground *string `field:"optional" json:"primaryBackground" yaml:"primaryBackground"`
	// The color of text and other foreground elements that appear over the primary background regions, such as grid lines, borders, table banding, icons, and so on.
	PrimaryForeground *string `field:"optional" json:"primaryForeground" yaml:"primaryForeground"`
	// The background color that applies to the sheet background and sheet controls.
	SecondaryBackground *string `field:"optional" json:"secondaryBackground" yaml:"secondaryBackground"`
	// The foreground color that applies to any sheet title, sheet control text, or UI that appears over the secondary background.
	SecondaryForeground *string `field:"optional" json:"secondaryForeground" yaml:"secondaryForeground"`
	// The color that applies to success messages, for example the check mark for a successful download.
	Success *string `field:"optional" json:"success" yaml:"success"`
	// The foreground color that applies to any text or other elements that appear over the success color.
	SuccessForeground *string `field:"optional" json:"successForeground" yaml:"successForeground"`
	// This color that applies to warning and informational messages.
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
	// The foreground color that applies to any text or other elements that appear over the warning color.
	WarningForeground *string `field:"optional" json:"warningForeground" yaml:"warningForeground"`
}

// Properties for defining a `CfnTheme`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnThemeProps := &cfnThemeProps{
//   	awsAccountId: jsii.String("awsAccountId"),
//   	themeId: jsii.String("themeId"),
//
//   	// the properties below are optional
//   	baseThemeId: jsii.String("baseThemeId"),
//   	configuration: &themeConfigurationProperty{
//   		dataColorPalette: &dataColorPaletteProperty{
//   			colors: []*string{
//   				jsii.String("colors"),
//   			},
//   			emptyFillColor: jsii.String("emptyFillColor"),
//   			minMaxGradient: []*string{
//   				jsii.String("minMaxGradient"),
//   			},
//   		},
//   		sheet: &sheetStyleProperty{
//   			tile: &tileStyleProperty{
//   				border: &borderStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   			},
//   			tileLayout: &tileLayoutStyleProperty{
//   				gutter: &gutterStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   				margin: &marginStyleProperty{
//   					show: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		typography: &typographyProperty{
//   			fontFamilies: []interface{}{
//   				&fontProperty{
//   					fontFamily: jsii.String("fontFamily"),
//   				},
//   			},
//   		},
//   		uiColorPalette: &uIColorPaletteProperty{
//   			accent: jsii.String("accent"),
//   			accentForeground: jsii.String("accentForeground"),
//   			danger: jsii.String("danger"),
//   			dangerForeground: jsii.String("dangerForeground"),
//   			dimension: jsii.String("dimension"),
//   			dimensionForeground: jsii.String("dimensionForeground"),
//   			measure: jsii.String("measure"),
//   			measureForeground: jsii.String("measureForeground"),
//   			primaryBackground: jsii.String("primaryBackground"),
//   			primaryForeground: jsii.String("primaryForeground"),
//   			secondaryBackground: jsii.String("secondaryBackground"),
//   			secondaryForeground: jsii.String("secondaryForeground"),
//   			success: jsii.String("success"),
//   			successForeground: jsii.String("successForeground"),
//   			warning: jsii.String("warning"),
//   			warningForeground: jsii.String("warningForeground"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	permissions: []interface{}{
//   		&resourcePermissionProperty{
//   			actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			principal: jsii.String("principal"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	versionDescription: jsii.String("versionDescription"),
//   }
//
type CfnThemeProps struct {
	// The ID of the AWS account where you want to store the new theme.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// An ID for the theme that you want to create.
	//
	// The theme ID is unique per AWS Region in each AWS account.
	ThemeId *string `field:"required" json:"themeId" yaml:"themeId"`
	// The ID of the theme that a custom theme will inherit from.
	//
	// All themes inherit from one of the starting themes defined by Amazon QuickSight. For a list of the starting themes, use `ListThemes` or choose *Themes* from within an analysis.
	BaseThemeId *string `field:"optional" json:"baseThemeId" yaml:"baseThemeId"`
	// The theme configuration, which contains the theme display properties.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A display name for the theme.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A valid grouping of resource permissions to apply to the new theme.
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// A map of the key-value pairs for the resource tag or tags that you want to add to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A description of the first version of the theme that you're creating.
	//
	// Every time `UpdateTheme` is called, a new version is created. Each version of the theme has a description of the version in the `VersionDescription` field.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

