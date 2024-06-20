package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new Q topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTopic := awscdk.Aws_quicksight.NewCfnTopic(this, jsii.String("MyCfnTopic"), &CfnTopicProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DataSets: []interface{}{
//   		&DatasetMetadataProperty{
//   			DatasetArn: jsii.String("datasetArn"),
//
//   			// the properties below are optional
//   			CalculatedFields: []interface{}{
//   				&TopicCalculatedFieldProperty{
//   					CalculatedFieldName: jsii.String("calculatedFieldName"),
//   					Expression: jsii.String("expression"),
//
//   					// the properties below are optional
//   					Aggregation: jsii.String("aggregation"),
//   					AllowedAggregations: []*string{
//   						jsii.String("allowedAggregations"),
//   					},
//   					CalculatedFieldDescription: jsii.String("calculatedFieldDescription"),
//   					CalculatedFieldSynonyms: []*string{
//   						jsii.String("calculatedFieldSynonyms"),
//   					},
//   					CellValueSynonyms: []interface{}{
//   						&CellValueSynonymProperty{
//   							CellValue: jsii.String("cellValue"),
//   							Synonyms: []*string{
//   								jsii.String("synonyms"),
//   							},
//   						},
//   					},
//   					ColumnDataRole: jsii.String("columnDataRole"),
//   					ComparativeOrder: &ComparativeOrderProperty{
//   						SpecifedOrder: []*string{
//   							jsii.String("specifedOrder"),
//   						},
//   						TreatUndefinedSpecifiedValues: jsii.String("treatUndefinedSpecifiedValues"),
//   						UseOrdering: jsii.String("useOrdering"),
//   					},
//   					DefaultFormatting: &DefaultFormattingProperty{
//   						DisplayFormat: jsii.String("displayFormat"),
//   						DisplayFormatOptions: &DisplayFormatOptionsProperty{
//   							BlankCellFormat: jsii.String("blankCellFormat"),
//   							CurrencySymbol: jsii.String("currencySymbol"),
//   							DateFormat: jsii.String("dateFormat"),
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							FractionDigits: jsii.Number(123),
//   							GroupingSeparator: jsii.String("groupingSeparator"),
//   							NegativeFormat: &NegativeFormatProperty{
//   								Prefix: jsii.String("prefix"),
//   								Suffix: jsii.String("suffix"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							Suffix: jsii.String("suffix"),
//   							UnitScaler: jsii.String("unitScaler"),
//   							UseBlankCellFormat: jsii.Boolean(false),
//   							UseGrouping: jsii.Boolean(false),
//   						},
//   					},
//   					DisableIndexing: jsii.Boolean(false),
//   					IsIncludedInTopic: jsii.Boolean(false),
//   					NeverAggregateInFilter: jsii.Boolean(false),
//   					NonAdditive: jsii.Boolean(false),
//   					NotAllowedAggregations: []*string{
//   						jsii.String("notAllowedAggregations"),
//   					},
//   					SemanticType: &SemanticTypeProperty{
//   						FalseyCellValue: jsii.String("falseyCellValue"),
//   						FalseyCellValueSynonyms: []*string{
//   							jsii.String("falseyCellValueSynonyms"),
//   						},
//   						SubTypeName: jsii.String("subTypeName"),
//   						TruthyCellValue: jsii.String("truthyCellValue"),
//   						TruthyCellValueSynonyms: []*string{
//   							jsii.String("truthyCellValueSynonyms"),
//   						},
//   						TypeName: jsii.String("typeName"),
//   						TypeParameters: map[string]*string{
//   							"typeParametersKey": jsii.String("typeParameters"),
//   						},
//   					},
//   					TimeGranularity: jsii.String("timeGranularity"),
//   				},
//   			},
//   			Columns: []interface{}{
//   				&TopicColumnProperty{
//   					ColumnName: jsii.String("columnName"),
//
//   					// the properties below are optional
//   					Aggregation: jsii.String("aggregation"),
//   					AllowedAggregations: []*string{
//   						jsii.String("allowedAggregations"),
//   					},
//   					CellValueSynonyms: []interface{}{
//   						&CellValueSynonymProperty{
//   							CellValue: jsii.String("cellValue"),
//   							Synonyms: []*string{
//   								jsii.String("synonyms"),
//   							},
//   						},
//   					},
//   					ColumnDataRole: jsii.String("columnDataRole"),
//   					ColumnDescription: jsii.String("columnDescription"),
//   					ColumnFriendlyName: jsii.String("columnFriendlyName"),
//   					ColumnSynonyms: []*string{
//   						jsii.String("columnSynonyms"),
//   					},
//   					ComparativeOrder: &ComparativeOrderProperty{
//   						SpecifedOrder: []*string{
//   							jsii.String("specifedOrder"),
//   						},
//   						TreatUndefinedSpecifiedValues: jsii.String("treatUndefinedSpecifiedValues"),
//   						UseOrdering: jsii.String("useOrdering"),
//   					},
//   					DefaultFormatting: &DefaultFormattingProperty{
//   						DisplayFormat: jsii.String("displayFormat"),
//   						DisplayFormatOptions: &DisplayFormatOptionsProperty{
//   							BlankCellFormat: jsii.String("blankCellFormat"),
//   							CurrencySymbol: jsii.String("currencySymbol"),
//   							DateFormat: jsii.String("dateFormat"),
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							FractionDigits: jsii.Number(123),
//   							GroupingSeparator: jsii.String("groupingSeparator"),
//   							NegativeFormat: &NegativeFormatProperty{
//   								Prefix: jsii.String("prefix"),
//   								Suffix: jsii.String("suffix"),
//   							},
//   							Prefix: jsii.String("prefix"),
//   							Suffix: jsii.String("suffix"),
//   							UnitScaler: jsii.String("unitScaler"),
//   							UseBlankCellFormat: jsii.Boolean(false),
//   							UseGrouping: jsii.Boolean(false),
//   						},
//   					},
//   					DisableIndexing: jsii.Boolean(false),
//   					IsIncludedInTopic: jsii.Boolean(false),
//   					NeverAggregateInFilter: jsii.Boolean(false),
//   					NonAdditive: jsii.Boolean(false),
//   					NotAllowedAggregations: []*string{
//   						jsii.String("notAllowedAggregations"),
//   					},
//   					SemanticType: &SemanticTypeProperty{
//   						FalseyCellValue: jsii.String("falseyCellValue"),
//   						FalseyCellValueSynonyms: []*string{
//   							jsii.String("falseyCellValueSynonyms"),
//   						},
//   						SubTypeName: jsii.String("subTypeName"),
//   						TruthyCellValue: jsii.String("truthyCellValue"),
//   						TruthyCellValueSynonyms: []*string{
//   							jsii.String("truthyCellValueSynonyms"),
//   						},
//   						TypeName: jsii.String("typeName"),
//   						TypeParameters: map[string]*string{
//   							"typeParametersKey": jsii.String("typeParameters"),
//   						},
//   					},
//   					TimeGranularity: jsii.String("timeGranularity"),
//   				},
//   			},
//   			DataAggregation: &DataAggregationProperty{
//   				DatasetRowDateGranularity: jsii.String("datasetRowDateGranularity"),
//   				DefaultDateColumnName: jsii.String("defaultDateColumnName"),
//   			},
//   			DatasetDescription: jsii.String("datasetDescription"),
//   			DatasetName: jsii.String("datasetName"),
//   			Filters: []interface{}{
//   				&TopicFilterProperty{
//   					FilterName: jsii.String("filterName"),
//   					OperandFieldName: jsii.String("operandFieldName"),
//
//   					// the properties below are optional
//   					CategoryFilter: &TopicCategoryFilterProperty{
//   						CategoryFilterFunction: jsii.String("categoryFilterFunction"),
//   						CategoryFilterType: jsii.String("categoryFilterType"),
//   						Constant: &TopicCategoryFilterConstantProperty{
//   							CollectiveConstant: &CollectiveConstantProperty{
//   								ValueList: []*string{
//   									jsii.String("valueList"),
//   								},
//   							},
//   							ConstantType: jsii.String("constantType"),
//   							SingularConstant: jsii.String("singularConstant"),
//   						},
//   						Inverse: jsii.Boolean(false),
//   					},
//   					DateRangeFilter: &TopicDateRangeFilterProperty{
//   						Constant: &TopicRangeFilterConstantProperty{
//   							ConstantType: jsii.String("constantType"),
//   							RangeConstant: &RangeConstantProperty{
//   								Maximum: jsii.String("maximum"),
//   								Minimum: jsii.String("minimum"),
//   							},
//   						},
//   						Inclusive: jsii.Boolean(false),
//   					},
//   					FilterClass: jsii.String("filterClass"),
//   					FilterDescription: jsii.String("filterDescription"),
//   					FilterSynonyms: []*string{
//   						jsii.String("filterSynonyms"),
//   					},
//   					FilterType: jsii.String("filterType"),
//   					NumericEqualityFilter: &TopicNumericEqualityFilterProperty{
//   						Aggregation: jsii.String("aggregation"),
//   						Constant: &TopicSingularFilterConstantProperty{
//   							ConstantType: jsii.String("constantType"),
//   							SingularConstant: jsii.String("singularConstant"),
//   						},
//   					},
//   					NumericRangeFilter: &TopicNumericRangeFilterProperty{
//   						Aggregation: jsii.String("aggregation"),
//   						Constant: &TopicRangeFilterConstantProperty{
//   							ConstantType: jsii.String("constantType"),
//   							RangeConstant: &RangeConstantProperty{
//   								Maximum: jsii.String("maximum"),
//   								Minimum: jsii.String("minimum"),
//   							},
//   						},
//   						Inclusive: jsii.Boolean(false),
//   					},
//   					RelativeDateFilter: &TopicRelativeDateFilterProperty{
//   						Constant: &TopicSingularFilterConstantProperty{
//   							ConstantType: jsii.String("constantType"),
//   							SingularConstant: jsii.String("singularConstant"),
//   						},
//   						RelativeDateFilterFunction: jsii.String("relativeDateFilterFunction"),
//   						TimeGranularity: jsii.String("timeGranularity"),
//   					},
//   				},
//   			},
//   			NamedEntities: []interface{}{
//   				&TopicNamedEntityProperty{
//   					EntityName: jsii.String("entityName"),
//
//   					// the properties below are optional
//   					Definition: []interface{}{
//   						&NamedEntityDefinitionProperty{
//   							FieldName: jsii.String("fieldName"),
//   							Metric: &NamedEntityDefinitionMetricProperty{
//   								Aggregation: jsii.String("aggregation"),
//   								AggregationFunctionParameters: map[string]*string{
//   									"aggregationFunctionParametersKey": jsii.String("aggregationFunctionParameters"),
//   								},
//   							},
//   							PropertyName: jsii.String("propertyName"),
//   							PropertyRole: jsii.String("propertyRole"),
//   							PropertyUsage: jsii.String("propertyUsage"),
//   						},
//   					},
//   					EntityDescription: jsii.String("entityDescription"),
//   					EntitySynonyms: []*string{
//   						jsii.String("entitySynonyms"),
//   					},
//   					SemanticEntityType: &SemanticEntityTypeProperty{
//   						SubTypeName: jsii.String("subTypeName"),
//   						TypeName: jsii.String("typeName"),
//   						TypeParameters: map[string]*string{
//   							"typeParametersKey": jsii.String("typeParameters"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	TopicId: jsii.String("topicId"),
//   	UserExperienceVersion: jsii.String("userExperienceVersion"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-topic.html
//
type CfnTopic interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the topic.
	AttrArn() *string
	// The ID of the AWS account that you want to create a topic in.
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
	// The data sets that the topic is associated with.
	DataSets() interface{}
	SetDataSets(val interface{})
	// The description of the topic.
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
	LogicalId() *string
	// The name of the topic.
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
	// The ID for the topic.
	TopicId() *string
	SetTopicId(val *string)
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
	// The user experience version of the topic.
	UserExperienceVersion() *string
	SetUserExperienceVersion(val *string)
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
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resources-section-structure.html#resources-section-structure-logicalid
	//
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

// The jsii proxy struct for CfnTopic
type jsiiProxy_CfnTopic struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTopic) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) DataSets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) TopicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTopic) UserExperienceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userExperienceVersion",
		&returns,
	)
	return returns
}


func NewCfnTopic(scope constructs.Construct, id *string, props *CfnTopicProps) CfnTopic {
	_init_.Initialize()

	if err := validateNewCfnTopicParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTopic{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTopic",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnTopic_Override(c CfnTopic, scope constructs.Construct, id *string, props *CfnTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnTopic",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTopic)SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnTopic)SetDataSets(val interface{}) {
	if err := j.validateSetDataSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSets",
		val,
	)
}

func (j *jsiiProxy_CfnTopic)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnTopic)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTopic)SetTopicId(val *string) {
	_jsii_.Set(
		j,
		"topicId",
		val,
	)
}

func (j *jsiiProxy_CfnTopic)SetUserExperienceVersion(val *string) {
	_jsii_.Set(
		j,
		"userExperienceVersion",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTopic_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopic_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTopic",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnTopic_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopic_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTopic",
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
func CfnTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTopic_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnTopic",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTopic) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTopic) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopic) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopic) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTopic) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTopic) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTopic) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTopic) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTopic) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnTopic) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTopic) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTopic) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTopic) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTopic) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTopic) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnTopic) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTopic) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

