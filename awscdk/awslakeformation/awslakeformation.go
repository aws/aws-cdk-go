package awslakeformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::LakeFormation::DataCellsFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allRowsWildcard interface{}
//
//   cfnDataCellsFilter := awscdk.Aws_lakeformation.NewCfnDataCellsFilter(this, jsii.String("MyCfnDataCellsFilter"), &cfnDataCellsFilterProps{
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   	tableCatalogId: jsii.String("tableCatalogId"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	columnWildcard: &columnWildcardProperty{
//   		excludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   	rowFilter: &rowFilterProperty{
//   		allRowsWildcard: allRowsWildcard,
//   		filterExpression: jsii.String("filterExpression"),
//   	},
//   })
//
type CfnDataCellsFilter interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::LakeFormation::DataCellsFilter.ColumnNames`.
	ColumnNames() *[]*string
	SetColumnNames(val *[]*string)
	// `AWS::LakeFormation::DataCellsFilter.ColumnWildcard`.
	ColumnWildcard() interface{}
	SetColumnWildcard(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::LakeFormation::DataCellsFilter.DatabaseName`.
	DatabaseName() *string
	SetDatabaseName(val *string)
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
	// `AWS::LakeFormation::DataCellsFilter.Name`.
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
	// `AWS::LakeFormation::DataCellsFilter.RowFilter`.
	RowFilter() interface{}
	SetRowFilter(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::LakeFormation::DataCellsFilter.TableCatalogId`.
	TableCatalogId() *string
	SetTableCatalogId(val *string)
	// `AWS::LakeFormation::DataCellsFilter.TableName`.
	TableName() *string
	SetTableName(val *string)
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

// The jsii proxy struct for CfnDataCellsFilter
type jsiiProxy_CfnDataCellsFilter struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataCellsFilter) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) ColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) ColumnWildcard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnWildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) RowFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) TableCatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableCatalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCellsFilter) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::DataCellsFilter`.
func NewCfnDataCellsFilter(scope awscdk.Construct, id *string, props *CfnDataCellsFilterProps) CfnDataCellsFilter {
	_init_.Initialize()

	j := jsiiProxy_CfnDataCellsFilter{}

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnDataCellsFilter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::DataCellsFilter`.
func NewCfnDataCellsFilter_Override(c CfnDataCellsFilter, scope awscdk.Construct, id *string, props *CfnDataCellsFilterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnDataCellsFilter",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataCellsFilter) SetColumnNames(val *[]*string) {
	_jsii_.Set(
		j,
		"columnNames",
		val,
	)
}

func (j *jsiiProxy_CfnDataCellsFilter) SetColumnWildcard(val interface{}) {
	_jsii_.Set(
		j,
		"columnWildcard",
		val,
	)
}

func (j *jsiiProxy_CfnDataCellsFilter) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnDataCellsFilter) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataCellsFilter) SetRowFilter(val interface{}) {
	_jsii_.Set(
		j,
		"rowFilter",
		val,
	)
}

func (j *jsiiProxy_CfnDataCellsFilter) SetTableCatalogId(val *string) {
	_jsii_.Set(
		j,
		"tableCatalogId",
		val,
	)
}

func (j *jsiiProxy_CfnDataCellsFilter) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
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
func CfnDataCellsFilter_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnDataCellsFilter",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDataCellsFilter_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnDataCellsFilter",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDataCellsFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnDataCellsFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataCellsFilter_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lakeformation.CfnDataCellsFilter",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataCellsFilter) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataCellsFilter) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataCellsFilter) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataCellsFilter) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataCellsFilter) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataCellsFilter) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDataCellsFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataCellsFilter) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataCellsFilter) ValidateProperties(_properties interface{}) {
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
//   columnWildcardProperty := &columnWildcardProperty{
//   	excludedColumnNames: []*string{
//   		jsii.String("excludedColumnNames"),
//   	},
//   }
//
type CfnDataCellsFilter_ColumnWildcardProperty struct {
	// `CfnDataCellsFilter.ColumnWildcardProperty.ExcludedColumnNames`.
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allRowsWildcard interface{}
//
//   rowFilterProperty := &rowFilterProperty{
//   	allRowsWildcard: allRowsWildcard,
//   	filterExpression: jsii.String("filterExpression"),
//   }
//
type CfnDataCellsFilter_RowFilterProperty struct {
	// `CfnDataCellsFilter.RowFilterProperty.AllRowsWildcard`.
	AllRowsWildcard interface{} `field:"optional" json:"allRowsWildcard" yaml:"allRowsWildcard"`
	// `CfnDataCellsFilter.RowFilterProperty.FilterExpression`.
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
}

// Properties for defining a `CfnDataCellsFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var allRowsWildcard interface{}
//
//   cfnDataCellsFilterProps := &cfnDataCellsFilterProps{
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   	tableCatalogId: jsii.String("tableCatalogId"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	columnWildcard: &columnWildcardProperty{
//   		excludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   	rowFilter: &rowFilterProperty{
//   		allRowsWildcard: allRowsWildcard,
//   		filterExpression: jsii.String("filterExpression"),
//   	},
//   }
//
type CfnDataCellsFilterProps struct {
	// `AWS::LakeFormation::DataCellsFilter.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `AWS::LakeFormation::DataCellsFilter.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::LakeFormation::DataCellsFilter.TableCatalogId`.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// `AWS::LakeFormation::DataCellsFilter.TableName`.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// `AWS::LakeFormation::DataCellsFilter.ColumnNames`.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// `AWS::LakeFormation::DataCellsFilter.ColumnWildcard`.
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
	// `AWS::LakeFormation::DataCellsFilter.RowFilter`.
	RowFilter interface{} `field:"optional" json:"rowFilter" yaml:"rowFilter"`
}

// A CloudFormation `AWS::LakeFormation::DataLakeSettings`.
//
// The `AWS::LakeFormation::DataLakeSettings` resource is an AWS Lake Formation resource type that manages the data lake settings for your account. Note that the CloudFormation template only supports updating the `Admins` list. It does not support updating the [CreateDatabaseDefaultPermissions](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-settings.html#aws-lake-formation-api-aws-lake-formation-api-settings-DataLakeSettings) or [CreateTableDefaultPermissions](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-settings.html#aws-lake-formation-api-aws-lake-formation-api-settings-DataLakeSettings) . Those permissions can only be edited in the DataLakeSettings resource via the API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataLakeSettings := awscdk.Aws_lakeformation.NewCfnDataLakeSettings(this, jsii.String("MyCfnDataLakeSettings"), &cfnDataLakeSettingsProps{
//   	admins: []interface{}{
//   		&dataLakePrincipalProperty{
//   			dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	trustedResourceOwners: []*string{
//   		jsii.String("trustedResourceOwners"),
//   	},
//   })
//
type CfnDataLakeSettings interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A list of AWS Lake Formation principals.
	Admins() interface{}
	SetAdmins(val interface{})
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::LakeFormation::DataLakeSettings.TrustedResourceOwners`.
	TrustedResourceOwners() *[]*string
	SetTrustedResourceOwners(val *[]*string)
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

// The jsii proxy struct for CfnDataLakeSettings
type jsiiProxy_CfnDataLakeSettings struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataLakeSettings) Admins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"admins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) TrustedResourceOwners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::DataLakeSettings`.
func NewCfnDataLakeSettings(scope awscdk.Construct, id *string, props *CfnDataLakeSettingsProps) CfnDataLakeSettings {
	_init_.Initialize()

	j := jsiiProxy_CfnDataLakeSettings{}

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnDataLakeSettings",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::DataLakeSettings`.
func NewCfnDataLakeSettings_Override(c CfnDataLakeSettings, scope awscdk.Construct, id *string, props *CfnDataLakeSettingsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnDataLakeSettings",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings) SetAdmins(val interface{}) {
	_jsii_.Set(
		j,
		"admins",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings) SetTrustedResourceOwners(val *[]*string) {
	_jsii_.Set(
		j,
		"trustedResourceOwners",
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
func CfnDataLakeSettings_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnDataLakeSettings",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDataLakeSettings_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnDataLakeSettings",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDataLakeSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnDataLakeSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataLakeSettings_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lakeformation.CfnDataLakeSettings",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The Lake Formation principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakePrincipalProperty := &dataLakePrincipalProperty{
//   	dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
type CfnDataLakeSettings_DataLakePrincipalProperty struct {
	// An identifier for the Lake Formation principal.
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

// Properties for defining a `CfnDataLakeSettings`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataLakeSettingsProps := &cfnDataLakeSettingsProps{
//   	admins: []interface{}{
//   		&dataLakePrincipalProperty{
//   			dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	trustedResourceOwners: []*string{
//   		jsii.String("trustedResourceOwners"),
//   	},
//   }
//
type CfnDataLakeSettingsProps struct {
	// A list of AWS Lake Formation principals.
	Admins interface{} `field:"optional" json:"admins" yaml:"admins"`
	// `AWS::LakeFormation::DataLakeSettings.TrustedResourceOwners`.
	TrustedResourceOwners *[]*string `field:"optional" json:"trustedResourceOwners" yaml:"trustedResourceOwners"`
}

// A CloudFormation `AWS::LakeFormation::Permissions`.
//
// The `AWS::LakeFormation::Permissions` resource represents the permissions that a principal has on an AWS Glue Data Catalog resource (such as AWS Glue database or AWS Glue tables). When you upload a permissions stack, the permissions are granted to the principal and when you remove the stack, the permissions are revoked from the principal. If you remove a stack, and the principal does not have the permissions referenced in the stack then AWS Lake Formation will throw an error because you can’t call revoke on non-existing permissions. To successfully remove the stack, you’ll need to regrant those permissions and then remove the stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissions := awscdk.Aws_lakeformation.NewCfnPermissions(this, jsii.String("MyCfnPermissions"), &cfnPermissionsProps{
//   	dataLakePrincipal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	resource: &resourceProperty{
//   		databaseResource: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		dataLocationResource: &dataLocationResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			s3Resource: jsii.String("s3Resource"),
//   		},
//   		tableResource: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   			tableWildcard: &tableWildcardProperty{
//   			},
//   		},
//   		tableWithColumnsResource: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			columnWildcard: &columnWildcardProperty{
//   				excludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	permissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   })
//
type CfnPermissions interface {
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
	// The AWS Lake Formation principal.
	DataLakePrincipal() interface{}
	SetDataLakePrincipal(val interface{})
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
	// The permissions granted or revoked.
	Permissions() *[]*string
	SetPermissions(val *[]*string)
	// Indicates whether to grant the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption() *[]*string
	SetPermissionsWithGrantOption(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A structure for the resource.
	Resource() interface{}
	SetResource(val interface{})
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

// The jsii proxy struct for CfnPermissions
type jsiiProxy_CfnPermissions struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPermissions) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) DataLakePrincipal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLakePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Resource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::Permissions`.
func NewCfnPermissions(scope awscdk.Construct, id *string, props *CfnPermissionsProps) CfnPermissions {
	_init_.Initialize()

	j := jsiiProxy_CfnPermissions{}

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnPermissions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::Permissions`.
func NewCfnPermissions_Override(c CfnPermissions, scope awscdk.Construct, id *string, props *CfnPermissionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnPermissions",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPermissions) SetDataLakePrincipal(val interface{}) {
	_jsii_.Set(
		j,
		"dataLakePrincipal",
		val,
	)
}

func (j *jsiiProxy_CfnPermissions) SetPermissions(val *[]*string) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnPermissions) SetPermissionsWithGrantOption(val *[]*string) {
	_jsii_.Set(
		j,
		"permissionsWithGrantOption",
		val,
	)
}

func (j *jsiiProxy_CfnPermissions) SetResource(val interface{}) {
	_jsii_.Set(
		j,
		"resource",
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
func CfnPermissions_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnPermissions",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPermissions_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnPermissions",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPermissions_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lakeformation.CfnPermissions",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPermissions) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPermissions) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPermissions) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPermissions) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPermissions) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPermissions) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPermissions) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPermissions) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPermissions) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPermissions) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPermissions) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPermissions) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPermissions) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A wildcard object, consisting of an optional list of excluded column names or indexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnWildcardProperty := &columnWildcardProperty{
//   	excludedColumnNames: []*string{
//   		jsii.String("excludedColumnNames"),
//   	},
//   }
//
type CfnPermissions_ColumnWildcardProperty struct {
	// Excludes column names.
	//
	// Any column with this name will be excluded.
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

// The Lake Formation principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakePrincipalProperty := &dataLakePrincipalProperty{
//   	dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
type CfnPermissions_DataLakePrincipalProperty struct {
	// An identifier for the Lake Formation principal.
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

// A structure for a data location object where permissions are granted or revoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLocationResourceProperty := &dataLocationResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	s3Resource: jsii.String("s3Resource"),
//   }
//
type CfnPermissions_DataLocationResourceProperty struct {
	// `CfnPermissions.DataLocationResourceProperty.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Currently not supported by AWS CloudFormation .
	S3Resource *string `field:"optional" json:"s3Resource" yaml:"s3Resource"`
}

// A structure for the database object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseResourceProperty := &databaseResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	name: jsii.String("name"),
//   }
//
type CfnPermissions_DatabaseResourceProperty struct {
	// `CfnPermissions.DatabaseResourceProperty.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database resource.
	//
	// Unique to the Data Catalog.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// A structure for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceProperty := &resourceProperty{
//   	databaseResource: &databaseResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		name: jsii.String("name"),
//   	},
//   	dataLocationResource: &dataLocationResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		s3Resource: jsii.String("s3Resource"),
//   	},
//   	tableResource: &tableResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   		tableWildcard: &tableWildcardProperty{
//   		},
//   	},
//   	tableWithColumnsResource: &tableWithColumnsResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		columnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		columnWildcard: &columnWildcardProperty{
//   			excludedColumnNames: []*string{
//   				jsii.String("excludedColumnNames"),
//   			},
//   		},
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnPermissions_ResourceProperty struct {
	// A structure for the database object.
	DatabaseResource interface{} `field:"optional" json:"databaseResource" yaml:"databaseResource"`
	// A structure for a data location object where permissions are granted or revoked.
	DataLocationResource interface{} `field:"optional" json:"dataLocationResource" yaml:"dataLocationResource"`
	// A structure for the table object.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	TableResource interface{} `field:"optional" json:"tableResource" yaml:"tableResource"`
	// Currently not supported by AWS CloudFormation .
	TableWithColumnsResource interface{} `field:"optional" json:"tableWithColumnsResource" yaml:"tableWithColumnsResource"`
}

// A structure for the table object.
//
// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableResourceProperty := &tableResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   	tableWildcard: &tableWildcardProperty{
//   	},
//   }
//
type CfnPermissions_TableResourceProperty struct {
	// `CfnPermissions.TableResourceProperty.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database for the table.
	//
	// Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the table.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An empty object representing all tables under a database.
	//
	// If this field is specified instead of the `Name` field, all tables under `DatabaseName` will have permission changes applied.
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}

// A wildcard object representing every table under a database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableWildcardProperty := &tableWildcardProperty{
//   }
//
type CfnPermissions_TableWildcardProperty struct {
}

// A structure for a table with columns object. This object is only used when granting a SELECT permission.
//
// This object must take a value for at least one of `ColumnsNames` , `ColumnsIndexes` , or `ColumnsWildcard` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableWithColumnsResourceProperty := &tableWithColumnsResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	columnWildcard: &columnWildcardProperty{
//   		excludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   }
//
type CfnPermissions_TableWithColumnsResourceProperty struct {
	// `CfnPermissions.TableWithColumnsResourceProperty.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The list of column names for the table.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// A wildcard specified by a `ColumnWildcard` object.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
	// The name of the database for the table with columns resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the table resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Properties for defining a `CfnPermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionsProps := &cfnPermissionsProps{
//   	dataLakePrincipal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	resource: &resourceProperty{
//   		databaseResource: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		dataLocationResource: &dataLocationResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			s3Resource: jsii.String("s3Resource"),
//   		},
//   		tableResource: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   			tableWildcard: &tableWildcardProperty{
//   			},
//   		},
//   		tableWithColumnsResource: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			columnWildcard: &columnWildcardProperty{
//   				excludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	permissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   }
//
type CfnPermissionsProps struct {
	// The AWS Lake Formation principal.
	DataLakePrincipal interface{} `field:"required" json:"dataLakePrincipal" yaml:"dataLakePrincipal"`
	// A structure for the resource.
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
	// The permissions granted or revoked.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Indicates whether to grant the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption *[]*string `field:"optional" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
}

// A CloudFormation `AWS::LakeFormation::PrincipalPermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnPrincipalPermissions := awscdk.Aws_lakeformation.NewCfnPrincipalPermissions(this, jsii.String("MyCfnPrincipalPermissions"), &cfnPrincipalPermissionsProps{
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	permissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   	principal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	resource: &resourceProperty{
//   		catalog: catalog,
//   		database: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		dataCellsFilter: &dataCellsFilterResourceProperty{
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   			tableCatalogId: jsii.String("tableCatalogId"),
//   			tableName: jsii.String("tableName"),
//   		},
//   		dataLocation: &dataLocationResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		lfTag: &lFTagKeyResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			tagKey: jsii.String("tagKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   		lfTagPolicy: &lFTagPolicyResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			expression: []interface{}{
//   				&lFTagProperty{
//   					tagKey: jsii.String("tagKey"),
//   					tagValues: []*string{
//   						jsii.String("tagValues"),
//   					},
//   				},
//   			},
//   			resourceType: jsii.String("resourceType"),
//   		},
//   		table: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   			tableWildcard: tableWildcard,
//   		},
//   		tableWithColumns: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			columnWildcard: &columnWildcardProperty{
//   				excludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	catalog: jsii.String("catalog"),
//   })
//
type CfnPrincipalPermissions interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrPrincipalIdentifier() *string
	AttrResourceIdentifier() *string
	// `AWS::LakeFormation::PrincipalPermissions.Catalog`.
	Catalog() *string
	SetCatalog(val *string)
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
	// `AWS::LakeFormation::PrincipalPermissions.Permissions`.
	Permissions() *[]*string
	SetPermissions(val *[]*string)
	// `AWS::LakeFormation::PrincipalPermissions.PermissionsWithGrantOption`.
	PermissionsWithGrantOption() *[]*string
	SetPermissionsWithGrantOption(val *[]*string)
	// `AWS::LakeFormation::PrincipalPermissions.Principal`.
	Principal() interface{}
	SetPrincipal(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::LakeFormation::PrincipalPermissions.Resource`.
	Resource() interface{}
	SetResource(val interface{})
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

// The jsii proxy struct for CfnPrincipalPermissions
type jsiiProxy_CfnPrincipalPermissions struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPrincipalPermissions) AttrPrincipalIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrincipalIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) AttrResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Principal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Resource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPrincipalPermissions) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::PrincipalPermissions`.
func NewCfnPrincipalPermissions(scope awscdk.Construct, id *string, props *CfnPrincipalPermissionsProps) CfnPrincipalPermissions {
	_init_.Initialize()

	j := jsiiProxy_CfnPrincipalPermissions{}

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnPrincipalPermissions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::PrincipalPermissions`.
func NewCfnPrincipalPermissions_Override(c CfnPrincipalPermissions, scope awscdk.Construct, id *string, props *CfnPrincipalPermissionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnPrincipalPermissions",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions) SetCatalog(val *string) {
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions) SetPermissions(val *[]*string) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions) SetPermissionsWithGrantOption(val *[]*string) {
	_jsii_.Set(
		j,
		"permissionsWithGrantOption",
		val,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions) SetPrincipal(val interface{}) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_CfnPrincipalPermissions) SetResource(val interface{}) {
	_jsii_.Set(
		j,
		"resource",
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
func CfnPrincipalPermissions_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnPrincipalPermissions",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnPrincipalPermissions_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnPrincipalPermissions",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnPrincipalPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnPrincipalPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPrincipalPermissions_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lakeformation.CfnPrincipalPermissions",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnPrincipalPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPrincipalPermissions) ValidateProperties(_properties interface{}) {
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
//   columnWildcardProperty := &columnWildcardProperty{
//   	excludedColumnNames: []*string{
//   		jsii.String("excludedColumnNames"),
//   	},
//   }
//
type CfnPrincipalPermissions_ColumnWildcardProperty struct {
	// `CfnPrincipalPermissions.ColumnWildcardProperty.ExcludedColumnNames`.
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCellsFilterResourceProperty := &dataCellsFilterResourceProperty{
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   	tableCatalogId: jsii.String("tableCatalogId"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnPrincipalPermissions_DataCellsFilterResourceProperty struct {
	// `CfnPrincipalPermissions.DataCellsFilterResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnPrincipalPermissions.DataCellsFilterResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnPrincipalPermissions.DataCellsFilterResourceProperty.TableCatalogId`.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// `CfnPrincipalPermissions.DataCellsFilterResourceProperty.TableName`.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakePrincipalProperty := &dataLakePrincipalProperty{
//   	dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
type CfnPrincipalPermissions_DataLakePrincipalProperty struct {
	// `CfnPrincipalPermissions.DataLakePrincipalProperty.DataLakePrincipalIdentifier`.
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLocationResourceProperty := &dataLocationResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnPrincipalPermissions_DataLocationResourceProperty struct {
	// `CfnPrincipalPermissions.DataLocationResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.DataLocationResourceProperty.ResourceArn`.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseResourceProperty := &databaseResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	name: jsii.String("name"),
//   }
//
type CfnPrincipalPermissions_DatabaseResourceProperty struct {
	// `CfnPrincipalPermissions.DatabaseResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.DatabaseResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagKeyResourceProperty := &lFTagKeyResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnPrincipalPermissions_LFTagKeyResourceProperty struct {
	// `CfnPrincipalPermissions.LFTagKeyResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.LFTagKeyResourceProperty.TagKey`.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// `CfnPrincipalPermissions.LFTagKeyResourceProperty.TagValues`.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagPolicyResourceProperty := &lFTagPolicyResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	expression: []interface{}{
//   		&lFTagProperty{
//   			tagKey: jsii.String("tagKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   	resourceType: jsii.String("resourceType"),
//   }
//
type CfnPrincipalPermissions_LFTagPolicyResourceProperty struct {
	// `CfnPrincipalPermissions.LFTagPolicyResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.LFTagPolicyResourceProperty.Expression`.
	Expression interface{} `field:"required" json:"expression" yaml:"expression"`
	// `CfnPrincipalPermissions.LFTagPolicyResourceProperty.ResourceType`.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagProperty := &lFTagProperty{
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnPrincipalPermissions_LFTagProperty struct {
	// `CfnPrincipalPermissions.LFTagProperty.TagKey`.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// `CfnPrincipalPermissions.LFTagProperty.TagValues`.
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   resourceProperty := &resourceProperty{
//   	catalog: catalog,
//   	database: &databaseResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		name: jsii.String("name"),
//   	},
//   	dataCellsFilter: &dataCellsFilterResourceProperty{
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   		tableCatalogId: jsii.String("tableCatalogId"),
//   		tableName: jsii.String("tableName"),
//   	},
//   	dataLocation: &dataLocationResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   	lfTag: &lFTagKeyResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		tagKey: jsii.String("tagKey"),
//   		tagValues: []*string{
//   			jsii.String("tagValues"),
//   		},
//   	},
//   	lfTagPolicy: &lFTagPolicyResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		expression: []interface{}{
//   			&lFTagProperty{
//   				tagKey: jsii.String("tagKey"),
//   				tagValues: []*string{
//   					jsii.String("tagValues"),
//   				},
//   			},
//   		},
//   		resourceType: jsii.String("resourceType"),
//   	},
//   	table: &tableResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   		tableWildcard: tableWildcard,
//   	},
//   	tableWithColumns: &tableWithColumnsResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		columnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		columnWildcard: &columnWildcardProperty{
//   			excludedColumnNames: []*string{
//   				jsii.String("excludedColumnNames"),
//   			},
//   		},
//   	},
//   }
//
type CfnPrincipalPermissions_ResourceProperty struct {
	// `CfnPrincipalPermissions.ResourceProperty.Catalog`.
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// `CfnPrincipalPermissions.ResourceProperty.Database`.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// `CfnPrincipalPermissions.ResourceProperty.DataCellsFilter`.
	DataCellsFilter interface{} `field:"optional" json:"dataCellsFilter" yaml:"dataCellsFilter"`
	// `CfnPrincipalPermissions.ResourceProperty.DataLocation`.
	DataLocation interface{} `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// `CfnPrincipalPermissions.ResourceProperty.LFTag`.
	LfTag interface{} `field:"optional" json:"lfTag" yaml:"lfTag"`
	// `CfnPrincipalPermissions.ResourceProperty.LFTagPolicy`.
	LfTagPolicy interface{} `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// `CfnPrincipalPermissions.ResourceProperty.Table`.
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// `CfnPrincipalPermissions.ResourceProperty.TableWithColumns`.
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tableWildcard interface{}
//
//   tableResourceProperty := &tableResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tableWildcard: tableWildcard,
//   }
//
type CfnPrincipalPermissions_TableResourceProperty struct {
	// `CfnPrincipalPermissions.TableResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.TableResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnPrincipalPermissions.TableResourceProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnPrincipalPermissions.TableResourceProperty.TableWildcard`.
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableWithColumnsResourceProperty := &tableWithColumnsResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	columnWildcard: &columnWildcardProperty{
//   		excludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   }
//
type CfnPrincipalPermissions_TableWithColumnsResourceProperty struct {
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.ColumnNames`.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// `CfnPrincipalPermissions.TableWithColumnsResourceProperty.ColumnWildcard`.
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
}

// Properties for defining a `CfnPrincipalPermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnPrincipalPermissionsProps := &cfnPrincipalPermissionsProps{
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	permissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   	principal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	resource: &resourceProperty{
//   		catalog: catalog,
//   		database: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		dataCellsFilter: &dataCellsFilterResourceProperty{
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   			tableCatalogId: jsii.String("tableCatalogId"),
//   			tableName: jsii.String("tableName"),
//   		},
//   		dataLocation: &dataLocationResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		lfTag: &lFTagKeyResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			tagKey: jsii.String("tagKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   		lfTagPolicy: &lFTagPolicyResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			expression: []interface{}{
//   				&lFTagProperty{
//   					tagKey: jsii.String("tagKey"),
//   					tagValues: []*string{
//   						jsii.String("tagValues"),
//   					},
//   				},
//   			},
//   			resourceType: jsii.String("resourceType"),
//   		},
//   		table: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   			tableWildcard: tableWildcard,
//   		},
//   		tableWithColumns: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			columnWildcard: &columnWildcardProperty{
//   				excludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	catalog: jsii.String("catalog"),
//   }
//
type CfnPrincipalPermissionsProps struct {
	// `AWS::LakeFormation::PrincipalPermissions.Permissions`.
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// `AWS::LakeFormation::PrincipalPermissions.PermissionsWithGrantOption`.
	PermissionsWithGrantOption *[]*string `field:"required" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// `AWS::LakeFormation::PrincipalPermissions.Principal`.
	Principal interface{} `field:"required" json:"principal" yaml:"principal"`
	// `AWS::LakeFormation::PrincipalPermissions.Resource`.
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
	// `AWS::LakeFormation::PrincipalPermissions.Catalog`.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
}

// A CloudFormation `AWS::LakeFormation::Resource`.
//
// The `AWS::LakeFormation::Resource` represents the data (Amazon S3 buckets and folders) that is being registered with AWS Lake Formation . When a `Resource` type CloudFormation template is uploaded, an AWS Lake Formation [`RegisterResource`](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-credential-vending.html#aws-lake-formation-api-credential-vending-RegisterResource) API call is made to register the resource. When a `Resource` type CloudFormation template is removed, the AWS Lake Formation [`DeregisterResource`](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-credential-vending.html#aws-lake-formation-api-credential-vending-DeregisterResource) API is called.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResource := awscdk.Aws_lakeformation.NewCfnResource(this, jsii.String("MyCfnResource"), &cfnResourceProps{
//   	resourceArn: jsii.String("resourceArn"),
//   	useServiceLinkedRole: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	roleArn: jsii.String("roleArn"),
//   })
//
type CfnResource interface {
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
	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn() *string
	SetResourceArn(val *string)
	// The IAM role that registered a resource.
	RoleArn() *string
	SetRoleArn(val *string)
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
	// Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog.
	UseServiceLinkedRole() interface{}
	SetUseServiceLinkedRole(val interface{})
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

// The jsii proxy struct for CfnResource
type jsiiProxy_CfnResource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) UseServiceLinkedRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useServiceLinkedRole",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::Resource`.
func NewCfnResource(scope awscdk.Construct, id *string, props *CfnResourceProps) CfnResource {
	_init_.Initialize()

	j := jsiiProxy_CfnResource{}

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::Resource`.
func NewCfnResource_Override(c CfnResource, scope awscdk.Construct, id *string, props *CfnResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnResource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResource) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnResource) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnResource) SetUseServiceLinkedRole(val interface{}) {
	_jsii_.Set(
		j,
		"useServiceLinkedRole",
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
func CfnResource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnResource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnResource",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lakeformation.CfnResource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResource) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResource) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResource) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResource) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceProps := &cfnResourceProps{
//   	resourceArn: jsii.String("resourceArn"),
//   	useServiceLinkedRole: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnResourceProps struct {
	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog.
	UseServiceLinkedRole interface{} `field:"required" json:"useServiceLinkedRole" yaml:"useServiceLinkedRole"`
	// The IAM role that registered a resource.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

// A CloudFormation `AWS::LakeFormation::Tag`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTag := awscdk.Aws_lakeformation.NewCfnTag(this, jsii.String("MyCfnTag"), &cfnTagProps{
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//
//   	// the properties below are optional
//   	catalogId: jsii.String("catalogId"),
//   })
//
type CfnTag interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::LakeFormation::Tag.CatalogId`.
	CatalogId() *string
	SetCatalogId(val *string)
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
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// `AWS::LakeFormation::Tag.TagKey`.
	TagKey() *string
	SetTagKey(val *string)
	// `AWS::LakeFormation::Tag.TagValues`.
	TagValues() *[]*string
	SetTagValues(val *[]*string)
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

// The jsii proxy struct for CfnTag
type jsiiProxy_CfnTag struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTag) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) TagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) TagValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTag) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::Tag`.
func NewCfnTag(scope awscdk.Construct, id *string, props *CfnTagProps) CfnTag {
	_init_.Initialize()

	j := jsiiProxy_CfnTag{}

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnTag",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::Tag`.
func NewCfnTag_Override(c CfnTag, scope awscdk.Construct, id *string, props *CfnTagProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnTag",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTag) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_CfnTag) SetTagKey(val *string) {
	_jsii_.Set(
		j,
		"tagKey",
		val,
	)
}

func (j *jsiiProxy_CfnTag) SetTagValues(val *[]*string) {
	_jsii_.Set(
		j,
		"tagValues",
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
func CfnTag_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnTag",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTag_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnTag",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTag_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnTag",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTag_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lakeformation.CfnTag",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTag) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTag) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTag) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTag) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTag) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTag) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTag) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTag) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTag) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTag) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTag) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTag) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTag) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTag) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTag) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTag) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTag) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTag) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTag) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTag) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTag) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::LakeFormation::TagAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnTagAssociation := awscdk.Aws_lakeformation.NewCfnTagAssociation(this, jsii.String("MyCfnTagAssociation"), &cfnTagAssociationProps{
//   	lfTags: []interface{}{
//   		&lFTagPairProperty{
//   			catalogId: jsii.String("catalogId"),
//   			tagKey: jsii.String("tagKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   	resource: &resourceProperty{
//   		catalog: catalog,
//   		database: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		table: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   			tableWildcard: tableWildcard,
//   		},
//   		tableWithColumns: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   })
//
type CfnTagAssociation interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrResourceIdentifier() *string
	AttrTagsIdentifier() *string
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
	// `AWS::LakeFormation::TagAssociation.LFTags`.
	LfTags() interface{}
	SetLfTags(val interface{})
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
	// `AWS::LakeFormation::TagAssociation.Resource`.
	Resource() interface{}
	SetResource(val interface{})
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

// The jsii proxy struct for CfnTagAssociation
type jsiiProxy_CfnTagAssociation struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTagAssociation) AttrResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) AttrTagsIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTagsIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) LfTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) Resource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTagAssociation) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::TagAssociation`.
func NewCfnTagAssociation(scope awscdk.Construct, id *string, props *CfnTagAssociationProps) CfnTagAssociation {
	_init_.Initialize()

	j := jsiiProxy_CfnTagAssociation{}

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnTagAssociation",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::TagAssociation`.
func NewCfnTagAssociation_Override(c CfnTagAssociation, scope awscdk.Construct, id *string, props *CfnTagAssociationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lakeformation.CfnTagAssociation",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTagAssociation) SetLfTags(val interface{}) {
	_jsii_.Set(
		j,
		"lfTags",
		val,
	)
}

func (j *jsiiProxy_CfnTagAssociation) SetResource(val interface{}) {
	_jsii_.Set(
		j,
		"resource",
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
func CfnTagAssociation_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnTagAssociation",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTagAssociation_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnTagAssociation",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTagAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lakeformation.CfnTagAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTagAssociation_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_lakeformation.CfnTagAssociation",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTagAssociation) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTagAssociation) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTagAssociation) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTagAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTagAssociation) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTagAssociation) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTagAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTagAssociation) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagAssociation) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagAssociation) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTagAssociation) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTagAssociation) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTagAssociation) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTagAssociation) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTagAssociation) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagAssociation) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagAssociation) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTagAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagAssociation) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTagAssociation) ValidateProperties(_properties interface{}) {
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
//   databaseResourceProperty := &databaseResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	name: jsii.String("name"),
//   }
//
type CfnTagAssociation_DatabaseResourceProperty struct {
	// `CfnTagAssociation.DatabaseResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnTagAssociation.DatabaseResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagPairProperty := &lFTagPairProperty{
//   	catalogId: jsii.String("catalogId"),
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnTagAssociation_LFTagPairProperty struct {
	// `CfnTagAssociation.LFTagPairProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnTagAssociation.LFTagPairProperty.TagKey`.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// `CfnTagAssociation.LFTagPairProperty.TagValues`.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   resourceProperty := &resourceProperty{
//   	catalog: catalog,
//   	database: &databaseResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		name: jsii.String("name"),
//   	},
//   	table: &tableResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   		tableWildcard: tableWildcard,
//   	},
//   	tableWithColumns: &tableWithColumnsResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		columnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnTagAssociation_ResourceProperty struct {
	// `CfnTagAssociation.ResourceProperty.Catalog`.
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// `CfnTagAssociation.ResourceProperty.Database`.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// `CfnTagAssociation.ResourceProperty.Table`.
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// `CfnTagAssociation.ResourceProperty.TableWithColumns`.
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tableWildcard interface{}
//
//   tableResourceProperty := &tableResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tableWildcard: tableWildcard,
//   }
//
type CfnTagAssociation_TableResourceProperty struct {
	// `CfnTagAssociation.TableResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnTagAssociation.TableResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnTagAssociation.TableResourceProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnTagAssociation.TableResourceProperty.TableWildcard`.
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableWithColumnsResourceProperty := &tableWithColumnsResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   }
//
type CfnTagAssociation_TableWithColumnsResourceProperty struct {
	// `CfnTagAssociation.TableWithColumnsResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnTagAssociation.TableWithColumnsResourceProperty.ColumnNames`.
	ColumnNames *[]*string `field:"required" json:"columnNames" yaml:"columnNames"`
	// `CfnTagAssociation.TableWithColumnsResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnTagAssociation.TableWithColumnsResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// Properties for defining a `CfnTagAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnTagAssociationProps := &cfnTagAssociationProps{
//   	lfTags: []interface{}{
//   		&lFTagPairProperty{
//   			catalogId: jsii.String("catalogId"),
//   			tagKey: jsii.String("tagKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   	resource: &resourceProperty{
//   		catalog: catalog,
//   		database: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		table: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   			tableWildcard: tableWildcard,
//   		},
//   		tableWithColumns: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnTagAssociationProps struct {
	// `AWS::LakeFormation::TagAssociation.LFTags`.
	LfTags interface{} `field:"required" json:"lfTags" yaml:"lfTags"`
	// `AWS::LakeFormation::TagAssociation.Resource`.
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
}

// Properties for defining a `CfnTag`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTagProps := &cfnTagProps{
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//
//   	// the properties below are optional
//   	catalogId: jsii.String("catalogId"),
//   }
//
type CfnTagProps struct {
	// `AWS::LakeFormation::Tag.TagKey`.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// `AWS::LakeFormation::Tag.TagValues`.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
	// `AWS::LakeFormation::Tag.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

