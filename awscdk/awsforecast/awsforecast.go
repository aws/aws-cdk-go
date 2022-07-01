package awsforecast

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsforecast/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Forecast::Dataset`.
//
// Creates an Amazon Forecast dataset. The information about the dataset that you provide helps Forecast understand how to consume the data for model training. This includes the following:
//
// - *`DataFrequency`* - How frequently your historical time-series data is collected.
// - *`Domain`* and *`DatasetType`* - Each dataset has an associated dataset domain and a type within the domain. Amazon Forecast provides a list of predefined domains and types within each domain. For each unique dataset domain and type within the domain, Amazon Forecast requires your data to include a minimum set of predefined fields.
// - *`Schema`* - A schema specifies the fields in the dataset, including the field name and data type.
//
// After creating a dataset, you import your training data into it and add the dataset to a dataset group. You use the dataset group to create a predictor. For more information, see [Importing datasets](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html) .
//
// To get a list of all your datasets, use the [ListDatasets](https://docs.aws.amazon.com/forecast/latest/dg/API_ListDatasets.html) operation.
//
// For example Forecast datasets, see the [Amazon Forecast Sample GitHub repository](https://docs.aws.amazon.com/https://github.com/aws-samples/amazon-forecast-samples) .
//
// > The `Status` of a dataset must be `ACTIVE` before you can import training data. Use the [DescribeDataset](https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDataset.html) operation to get the status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryptionConfig interface{}
//   var schema interface{}
//   var tags interface{}
//
//   cfnDataset := awscdk.Aws_forecast.NewCfnDataset(this, jsii.String("MyCfnDataset"), &cfnDatasetProps{
//   	datasetName: jsii.String("datasetName"),
//   	datasetType: jsii.String("datasetType"),
//   	domain: jsii.String("domain"),
//   	schema: schema,
//
//   	// the properties below are optional
//   	dataFrequency: jsii.String("dataFrequency"),
//   	encryptionConfig: encryptionConfig,
//   	tags: []interface{}{
//   		tags,
//   	},
//   })
//
type CfnDataset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the dataset.
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
	// The frequency of data collection. This parameter is required for RELATED_TIME_SERIES datasets.
	//
	// Valid intervals are Y (Year), M (Month), W (Week), D (Day), H (Hour), 30min (30 minutes), 15min (15 minutes), 10min (10 minutes), 5min (5 minutes), and 1min (1 minute). For example, "D" indicates every day and "15min" indicates every 15 minutes.
	DataFrequency() *string
	SetDataFrequency(val *string)
	// The name of the dataset.
	DatasetName() *string
	SetDatasetName(val *string)
	// The dataset type.
	DatasetType() *string
	SetDatasetType(val *string)
	// The domain associated with the dataset.
	Domain() *string
	SetDomain(val *string)
	// A Key Management Service (KMS) key and the Identity and Access Management (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig() interface{}
	SetEncryptionConfig(val interface{})
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
	// The schema for the dataset.
	//
	// The schema attributes and their order must match the fields in your data. The dataset `Domain` and `DatasetType` that you choose determine the minimum required fields in your training data. For information about the required fields for a specific dataset domain and type, see [Dataset Domains and Dataset Types](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-domains-ds-types.html) .
	Schema() interface{}
	SetSchema(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() *[]interface{}
	SetTags(val *[]interface{})
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

// The jsii proxy struct for CfnDataset
type jsiiProxy_CfnDataset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataset) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) DataFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) DatasetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) EncryptionConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Schema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Tags() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Forecast::Dataset`.
func NewCfnDataset(scope constructs.Construct, id *string, props *CfnDatasetProps) CfnDataset {
	_init_.Initialize()

	j := jsiiProxy_CfnDataset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_forecast.CfnDataset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Forecast::Dataset`.
func NewCfnDataset_Override(c CfnDataset, scope constructs.Construct, id *string, props *CfnDatasetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_forecast.CfnDataset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataset) SetDataFrequency(val *string) {
	_jsii_.Set(
		j,
		"dataFrequency",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetDatasetName(val *string) {
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetDatasetType(val *string) {
	_jsii_.Set(
		j,
		"datasetType",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetEncryptionConfig(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetSchema(val interface{}) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetTags(val *[]interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_forecast.CfnDataset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_forecast.CfnDataset",
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
func CfnDataset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_forecast.CfnDataset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_forecast.CfnDataset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataset) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataset) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataset) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataset) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataset) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataset) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::Forecast::DatasetGroup`.
//
// Creates a dataset group, which holds a collection of related datasets. You can add datasets to the dataset group when you create the dataset group, or later by using the [UpdateDatasetGroup](https://docs.aws.amazon.com/forecast/latest/dg/API_UpdateDatasetGroup.html) operation.
//
// After creating a dataset group and adding datasets, you use the dataset group when you create a predictor. For more information, see [Dataset groups](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html) .
//
// To get a list of all your datasets groups, use the [ListDatasetGroups](https://docs.aws.amazon.com/forecast/latest/dg/API_ListDatasetGroups.html) operation.
//
// > The `Status` of a dataset group must be `ACTIVE` before you can use the dataset group to create a predictor. To get the status, use the [DescribeDatasetGroup](https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetGroup.html) operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetGroup := awscdk.Aws_forecast.NewCfnDatasetGroup(this, jsii.String("MyCfnDatasetGroup"), &cfnDatasetGroupProps{
//   	datasetGroupName: jsii.String("datasetGroupName"),
//   	domain: jsii.String("domain"),
//
//   	// the properties below are optional
//   	datasetArns: []*string{
//   		jsii.String("datasetArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDatasetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the dataset group.
	AttrDatasetGroupArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.
	DatasetArns() *[]*string
	SetDatasetArns(val *[]*string)
	// The name of the dataset group.
	DatasetGroupName() *string
	SetDatasetGroupName(val *string)
	// The domain associated with the dataset group.
	//
	// When you add a dataset to a dataset group, this value and the value specified for the `Domain` parameter of the [CreateDataset](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html) operation must match.
	//
	// The `Domain` and `DatasetType` that you choose determine the fields that must be present in training data that you import to a dataset. For example, if you choose the `RETAIL` domain and `TARGET_TIME_SERIES` as the `DatasetType` , Amazon Forecast requires that `item_id` , `timestamp` , and `demand` fields are present in your data. For more information, see [Dataset groups](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html) .
	Domain() *string
	SetDomain(val *string)
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

// The jsii proxy struct for CfnDatasetGroup
type jsiiProxy_CfnDatasetGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDatasetGroup) AttrDatasetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDatasetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) DatasetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"datasetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) DatasetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatasetGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Forecast::DatasetGroup`.
func NewCfnDatasetGroup(scope constructs.Construct, id *string, props *CfnDatasetGroupProps) CfnDatasetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDatasetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_forecast.CfnDatasetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Forecast::DatasetGroup`.
func NewCfnDatasetGroup_Override(c CfnDatasetGroup, scope constructs.Construct, id *string, props *CfnDatasetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_forecast.CfnDatasetGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDatasetGroup) SetDatasetArns(val *[]*string) {
	_jsii_.Set(
		j,
		"datasetArns",
		val,
	)
}

func (j *jsiiProxy_CfnDatasetGroup) SetDatasetGroupName(val *string) {
	_jsii_.Set(
		j,
		"datasetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDatasetGroup) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDatasetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_forecast.CfnDatasetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDatasetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_forecast.CfnDatasetGroup",
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
func CfnDatasetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_forecast.CfnDatasetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatasetGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_forecast.CfnDatasetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDatasetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatasetGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatasetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDatasetGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatasetGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatasetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatasetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDatasetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetGroupProps := &cfnDatasetGroupProps{
//   	datasetGroupName: jsii.String("datasetGroupName"),
//   	domain: jsii.String("domain"),
//
//   	// the properties below are optional
//   	datasetArns: []*string{
//   		jsii.String("datasetArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDatasetGroupProps struct {
	// The name of the dataset group.
	DatasetGroupName *string `field:"required" json:"datasetGroupName" yaml:"datasetGroupName"`
	// The domain associated with the dataset group.
	//
	// When you add a dataset to a dataset group, this value and the value specified for the `Domain` parameter of the [CreateDataset](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html) operation must match.
	//
	// The `Domain` and `DatasetType` that you choose determine the fields that must be present in training data that you import to a dataset. For example, if you choose the `RETAIL` domain and `TARGET_TIME_SERIES` as the `DatasetType` , Amazon Forecast requires that `item_id` , `timestamp` , and `demand` fields are present in your data. For more information, see [Dataset groups](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html) .
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.
	DatasetArns *[]*string `field:"optional" json:"datasetArns" yaml:"datasetArns"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryptionConfig interface{}
//   var schema interface{}
//   var tags interface{}
//
//   cfnDatasetProps := &cfnDatasetProps{
//   	datasetName: jsii.String("datasetName"),
//   	datasetType: jsii.String("datasetType"),
//   	domain: jsii.String("domain"),
//   	schema: schema,
//
//   	// the properties below are optional
//   	dataFrequency: jsii.String("dataFrequency"),
//   	encryptionConfig: encryptionConfig,
//   	tags: []interface{}{
//   		tags,
//   	},
//   }
//
type CfnDatasetProps struct {
	// The name of the dataset.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// The dataset type.
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The domain associated with the dataset.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The schema for the dataset.
	//
	// The schema attributes and their order must match the fields in your data. The dataset `Domain` and `DatasetType` that you choose determine the minimum required fields in your training data. For information about the required fields for a specific dataset domain and type, see [Dataset Domains and Dataset Types](https://docs.aws.amazon.com/forecast/latest/dg/howitworks-domains-ds-types.html) .
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
	// The frequency of data collection. This parameter is required for RELATED_TIME_SERIES datasets.
	//
	// Valid intervals are Y (Year), M (Month), W (Week), D (Day), H (Hour), 30min (30 minutes), 15min (15 minutes), 10min (10 minutes), 5min (5 minutes), and 1min (1 minute). For example, "D" indicates every day and "15min" indicates every 15 minutes.
	DataFrequency *string `field:"optional" json:"dataFrequency" yaml:"dataFrequency"`
	// A Key Management Service (KMS) key and the Identity and Access Management (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]interface{} `field:"optional" json:"tags" yaml:"tags"`
}

