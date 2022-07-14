package awsdatabrew

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DataBrew::Dataset`.
//
// Specifies a new DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataset := awscdk.Aws_databrew.NewCfnDataset(this, jsii.String("MyCfnDataset"), &cfnDatasetProps{
//   	input: &inputProperty{
//   		databaseInputDefinition: &databaseInputDefinitionProperty{
//   			glueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			databaseTableName: jsii.String("databaseTableName"),
//   			queryString: jsii.String("queryString"),
//   			tempDirectory: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   			},
//   		},
//   		dataCatalogInputDefinition: &dataCatalogInputDefinitionProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
//   			tempDirectory: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   			},
//   		},
//   		metadata: &metadataProperty{
//   			sourceArn: jsii.String("sourceArn"),
//   		},
//   		s3InputDefinition: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	format: jsii.String("format"),
//   	formatOptions: &formatOptionsProperty{
//   		csv: &csvOptionsProperty{
//   			delimiter: jsii.String("delimiter"),
//   			headerRow: jsii.Boolean(false),
//   		},
//   		excel: &excelOptionsProperty{
//   			headerRow: jsii.Boolean(false),
//   			sheetIndexes: []interface{}{
//   				jsii.Number(123),
//   			},
//   			sheetNames: []*string{
//   				jsii.String("sheetNames"),
//   			},
//   		},
//   		json: &jsonOptionsProperty{
//   			multiLine: jsii.Boolean(false),
//   		},
//   	},
//   	pathOptions: &pathOptionsProperty{
//   		filesLimit: &filesLimitProperty{
//   			maxFiles: jsii.Number(123),
//
//   			// the properties below are optional
//   			order: jsii.String("order"),
//   			orderedBy: jsii.String("orderedBy"),
//   		},
//   		lastModifiedDateCondition: &filterExpressionProperty{
//   			expression: jsii.String("expression"),
//   			valuesMap: []interface{}{
//   				&filterValueProperty{
//   					value: jsii.String("value"),
//   					valueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   		parameters: []interface{}{
//   			&pathParameterProperty{
//   				datasetParameter: &datasetParameterProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					createColumn: jsii.Boolean(false),
//   					datetimeOptions: &datetimeOptionsProperty{
//   						format: jsii.String("format"),
//
//   						// the properties below are optional
//   						localeCode: jsii.String("localeCode"),
//   						timezoneOffset: jsii.String("timezoneOffset"),
//   					},
//   					filter: &filterExpressionProperty{
//   						expression: jsii.String("expression"),
//   						valuesMap: []interface{}{
//   							&filterValueProperty{
//   								value: jsii.String("value"),
//   								valueReference: jsii.String("valueReference"),
//   							},
//   						},
//   					},
//   				},
//   				pathParameterName: jsii.String("pathParameterName"),
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
type CfnDataset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The file format of a dataset that is created from an Amazon S3 file or folder.
	Format() *string
	SetFormat(val *string)
	// A set of options that define how DataBrew interprets the data in the dataset.
	FormatOptions() interface{}
	SetFormatOptions(val interface{})
	// Information on how DataBrew can find the dataset, in either the AWS Glue Data Catalog or Amazon S3 .
	Input() interface{}
	SetInput(val interface{})
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
	// The unique name of the dataset.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A set of options that defines how DataBrew interprets an Amazon S3 path of the dataset.
	PathOptions() interface{}
	SetPathOptions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata tags that have been applied to the dataset.
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

// The jsii proxy struct for CfnDataset
type jsiiProxy_CfnDataset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnDataset) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) FormatOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formatOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"input",
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

func (j *jsiiProxy_CfnDataset) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
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

func (j *jsiiProxy_CfnDataset) PathOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathOptions",
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

func (j *jsiiProxy_CfnDataset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
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

func (j *jsiiProxy_CfnDataset) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Dataset`.
func NewCfnDataset(scope constructs.Construct, id *string, props *CfnDatasetProps) CfnDataset {
	_init_.Initialize()

	j := jsiiProxy_CfnDataset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnDataset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Dataset`.
func NewCfnDataset_Override(c CfnDataset, scope constructs.Construct, id *string, props *CfnDatasetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnDataset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataset) SetFormat(val *string) {
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetFormatOptions(val interface{}) {
	_jsii_.Set(
		j,
		"formatOptions",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetInput(val interface{}) {
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataset) SetPathOptions(val interface{}) {
	_jsii_.Set(
		j,
		"pathOptions",
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
		"aws-cdk-lib.aws_databrew.CfnDataset",
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
		"aws-cdk-lib.aws_databrew.CfnDataset",
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
		"aws-cdk-lib.aws_databrew.CfnDataset",
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
		"aws-cdk-lib.aws_databrew.CfnDataset",
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

// Represents a set of options that define how DataBrew will read a comma-separated value (CSV) file when creating a dataset from that file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvOptionsProperty := &csvOptionsProperty{
//   	delimiter: jsii.String("delimiter"),
//   	headerRow: jsii.Boolean(false),
//   }
//
type CfnDataset_CsvOptionsProperty struct {
	// A single character that specifies the delimiter being used in the CSV file.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// A variable that specifies whether the first row in the file is parsed as the header.
	//
	// If this value is false, column names are auto-generated.
	HeaderRow interface{} `field:"optional" json:"headerRow" yaml:"headerRow"`
}

// Represents how metadata stored in the AWS Glue Data Catalog is defined in a DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogInputDefinitionProperty := &dataCatalogInputDefinitionProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	tableName: jsii.String("tableName"),
//   	tempDirectory: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataset_DataCatalogInputDefinitionProperty struct {
	// The unique identifier of the AWS account that holds the Data Catalog that stores the data.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of a database in the Data Catalog.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of a database table in the Data Catalog.
	//
	// This table corresponds to a DataBrew dataset.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// An Amazon location that AWS Glue Data Catalog can use as a temporary directory.
	TempDirectory interface{} `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

// Connection information for dataset input files stored in a database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseInputDefinitionProperty := &databaseInputDefinitionProperty{
//   	glueConnectionName: jsii.String("glueConnectionName"),
//
//   	// the properties below are optional
//   	databaseTableName: jsii.String("databaseTableName"),
//   	queryString: jsii.String("queryString"),
//   	tempDirectory: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataset_DatabaseInputDefinitionProperty struct {
	// The AWS Glue Connection that stores the connection information for the target database.
	GlueConnectionName *string `field:"required" json:"glueConnectionName" yaml:"glueConnectionName"`
	// The table within the target database.
	DatabaseTableName *string `field:"optional" json:"databaseTableName" yaml:"databaseTableName"`
	// Custom SQL to run against the provided AWS Glue connection.
	//
	// This SQL will be used as the input for DataBrew projects and jobs.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// An Amazon location that AWS Glue Data Catalog can use as a temporary directory.
	TempDirectory interface{} `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

// Represents a dataset paramater that defines type and conditions for a parameter in the Amazon S3 path of the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetParameterProperty := &datasetParameterProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	createColumn: jsii.Boolean(false),
//   	datetimeOptions: &datetimeOptionsProperty{
//   		format: jsii.String("format"),
//
//   		// the properties below are optional
//   		localeCode: jsii.String("localeCode"),
//   		timezoneOffset: jsii.String("timezoneOffset"),
//   	},
//   	filter: &filterExpressionProperty{
//   		expression: jsii.String("expression"),
//   		valuesMap: []interface{}{
//   			&filterValueProperty{
//   				value: jsii.String("value"),
//   				valueReference: jsii.String("valueReference"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataset_DatasetParameterProperty struct {
	// The name of the parameter that is used in the dataset's Amazon S3 path.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the dataset parameter, can be one of a 'String', 'Number' or 'Datetime'.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Optional boolean value that defines whether the captured value of this parameter should be loaded as an additional column in the dataset.
	CreateColumn interface{} `field:"optional" json:"createColumn" yaml:"createColumn"`
	// Additional parameter options such as a format and a timezone.
	//
	// Required for datetime parameters.
	DatetimeOptions interface{} `field:"optional" json:"datetimeOptions" yaml:"datetimeOptions"`
	// The optional filter expression structure to apply additional matching criteria to the parameter.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

// Represents additional options for correct interpretation of datetime parameters used in the Amazon S3 path of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datetimeOptionsProperty := &datetimeOptionsProperty{
//   	format: jsii.String("format"),
//
//   	// the properties below are optional
//   	localeCode: jsii.String("localeCode"),
//   	timezoneOffset: jsii.String("timezoneOffset"),
//   }
//
type CfnDataset_DatetimeOptionsProperty struct {
	// Required option, that defines the datetime format used for a date parameter in the Amazon S3 path.
	//
	// Should use only supported datetime specifiers and separation characters, all litera a-z or A-Z character should be escaped with single quotes. E.g. "MM.dd.yyyy-'at'-HH:mm".
	Format *string `field:"required" json:"format" yaml:"format"`
	// Optional value for a non-US locale code, needed for correct interpretation of some date formats.
	LocaleCode *string `field:"optional" json:"localeCode" yaml:"localeCode"`
	// Optional value for a timezone offset of the datetime parameter value in the Amazon S3 path.
	//
	// Shouldn't be used if Format for this parameter includes timezone fields. If no offset specified, UTC is assumed.
	TimezoneOffset *string `field:"optional" json:"timezoneOffset" yaml:"timezoneOffset"`
}

// Represents a set of options that define how DataBrew will interpret a Microsoft Excel file when creating a dataset from that file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   excelOptionsProperty := &excelOptionsProperty{
//   	headerRow: jsii.Boolean(false),
//   	sheetIndexes: []interface{}{
//   		jsii.Number(123),
//   	},
//   	sheetNames: []*string{
//   		jsii.String("sheetNames"),
//   	},
//   }
//
type CfnDataset_ExcelOptionsProperty struct {
	// A variable that specifies whether the first row in the file is parsed as the header.
	//
	// If this value is false, column names are auto-generated.
	HeaderRow interface{} `field:"optional" json:"headerRow" yaml:"headerRow"`
	// One or more sheet numbers in the Excel file that will be included in the dataset.
	SheetIndexes interface{} `field:"optional" json:"sheetIndexes" yaml:"sheetIndexes"`
	// One or more named sheets in the Excel file that will be included in the dataset.
	SheetNames *[]*string `field:"optional" json:"sheetNames" yaml:"sheetNames"`
}

// Represents a limit imposed on number of Amazon S3 files that should be selected for a dataset from a connected Amazon S3 path.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filesLimitProperty := &filesLimitProperty{
//   	maxFiles: jsii.Number(123),
//
//   	// the properties below are optional
//   	order: jsii.String("order"),
//   	orderedBy: jsii.String("orderedBy"),
//   }
//
type CfnDataset_FilesLimitProperty struct {
	// The number of Amazon S3 files to select.
	MaxFiles *float64 `field:"required" json:"maxFiles" yaml:"maxFiles"`
	// A criteria to use for Amazon S3 files sorting before their selection.
	//
	// By default uses DESCENDING order, i.e. most recent files are selected first. Anotherpossible value is ASCENDING.
	Order *string `field:"optional" json:"order" yaml:"order"`
	// A criteria to use for Amazon S3 files sorting before their selection.
	//
	// By default uses LAST_MODIFIED_DATE as a sorting criteria. Currently it's the only allowed value.
	OrderedBy *string `field:"optional" json:"orderedBy" yaml:"orderedBy"`
}

// Represents a structure for defining parameter conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterExpressionProperty := &filterExpressionProperty{
//   	expression: jsii.String("expression"),
//   	valuesMap: []interface{}{
//   		&filterValueProperty{
//   			value: jsii.String("value"),
//   			valueReference: jsii.String("valueReference"),
//   		},
//   	},
//   }
//
type CfnDataset_FilterExpressionProperty struct {
	// The expression which includes condition names followed by substitution variables, possibly grouped and combined with other conditions.
	//
	// For example, "(starts_with :prefix1 or starts_with :prefix2) and (ends_with :suffix1 or ends_with :suffix2)". Substitution variables should start with ':' symbol.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The map of substitution variable names to their values used in this filter expression.
	ValuesMap interface{} `field:"required" json:"valuesMap" yaml:"valuesMap"`
}

// Represents a single entry in the `ValuesMap` of a `FilterExpression` .
//
// A `FilterValue` associates the name of a substitution variable in an expression to its value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterValueProperty := &filterValueProperty{
//   	value: jsii.String("value"),
//   	valueReference: jsii.String("valueReference"),
//   }
//
type CfnDataset_FilterValueProperty struct {
	// The value to be associated with the substitution variable.
	Value *string `field:"required" json:"value" yaml:"value"`
	// The substitution variable reference.
	ValueReference *string `field:"required" json:"valueReference" yaml:"valueReference"`
}

// Represents a set of options that define the structure of either comma-separated value (CSV), Excel, or JSON input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formatOptionsProperty := &formatOptionsProperty{
//   	csv: &csvOptionsProperty{
//   		delimiter: jsii.String("delimiter"),
//   		headerRow: jsii.Boolean(false),
//   	},
//   	excel: &excelOptionsProperty{
//   		headerRow: jsii.Boolean(false),
//   		sheetIndexes: []interface{}{
//   			jsii.Number(123),
//   		},
//   		sheetNames: []*string{
//   			jsii.String("sheetNames"),
//   		},
//   	},
//   	json: &jsonOptionsProperty{
//   		multiLine: jsii.Boolean(false),
//   	},
//   }
//
type CfnDataset_FormatOptionsProperty struct {
	// Options that define how CSV input is to be interpreted by DataBrew.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// Options that define how Excel input is to be interpreted by DataBrew.
	Excel interface{} `field:"optional" json:"excel" yaml:"excel"`
	// Options that define how JSON input is to be interpreted by DataBrew.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
}

// Represents information on how DataBrew can find data, in either the AWS Glue Data Catalog or Amazon S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &inputProperty{
//   	databaseInputDefinition: &databaseInputDefinitionProperty{
//   		glueConnectionName: jsii.String("glueConnectionName"),
//
//   		// the properties below are optional
//   		databaseTableName: jsii.String("databaseTableName"),
//   		queryString: jsii.String("queryString"),
//   		tempDirectory: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   		},
//   	},
//   	dataCatalogInputDefinition: &dataCatalogInputDefinitionProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		tableName: jsii.String("tableName"),
//   		tempDirectory: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   		},
//   	},
//   	metadata: &metadataProperty{
//   		sourceArn: jsii.String("sourceArn"),
//   	},
//   	s3InputDefinition: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataset_InputProperty struct {
	// Connection information for dataset input files stored in a database.
	DatabaseInputDefinition interface{} `field:"optional" json:"databaseInputDefinition" yaml:"databaseInputDefinition"`
	// The AWS Glue Data Catalog parameters for the data.
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// Contains additional resource information needed for specific datasets.
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The Amazon S3 location where the data is stored.
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

// Represents the JSON-specific options that define how input is to be interpreted by AWS Glue DataBrew .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonOptionsProperty := &jsonOptionsProperty{
//   	multiLine: jsii.Boolean(false),
//   }
//
type CfnDataset_JsonOptionsProperty struct {
	// A value that specifies whether JSON input contains embedded new line characters.
	MultiLine interface{} `field:"optional" json:"multiLine" yaml:"multiLine"`
}

// Contains additional resource information needed for specific datasets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataProperty := &metadataProperty{
//   	sourceArn: jsii.String("sourceArn"),
//   }
//
type CfnDataset_MetadataProperty struct {
	// The Amazon Resource Name (ARN) associated with the dataset.
	//
	// Currently, DataBrew only supports ARNs from Amazon AppFlow.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

// Represents a set of options that define how DataBrew selects files for a given Amazon S3 path in a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathOptionsProperty := &pathOptionsProperty{
//   	filesLimit: &filesLimitProperty{
//   		maxFiles: jsii.Number(123),
//
//   		// the properties below are optional
//   		order: jsii.String("order"),
//   		orderedBy: jsii.String("orderedBy"),
//   	},
//   	lastModifiedDateCondition: &filterExpressionProperty{
//   		expression: jsii.String("expression"),
//   		valuesMap: []interface{}{
//   			&filterValueProperty{
//   				value: jsii.String("value"),
//   				valueReference: jsii.String("valueReference"),
//   			},
//   		},
//   	},
//   	parameters: []interface{}{
//   		&pathParameterProperty{
//   			datasetParameter: &datasetParameterProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				createColumn: jsii.Boolean(false),
//   				datetimeOptions: &datetimeOptionsProperty{
//   					format: jsii.String("format"),
//
//   					// the properties below are optional
//   					localeCode: jsii.String("localeCode"),
//   					timezoneOffset: jsii.String("timezoneOffset"),
//   				},
//   				filter: &filterExpressionProperty{
//   					expression: jsii.String("expression"),
//   					valuesMap: []interface{}{
//   						&filterValueProperty{
//   							value: jsii.String("value"),
//   							valueReference: jsii.String("valueReference"),
//   						},
//   					},
//   				},
//   			},
//   			pathParameterName: jsii.String("pathParameterName"),
//   		},
//   	},
//   }
//
type CfnDataset_PathOptionsProperty struct {
	// If provided, this structure imposes a limit on a number of files that should be selected.
	FilesLimit interface{} `field:"optional" json:"filesLimit" yaml:"filesLimit"`
	// If provided, this structure defines a date range for matching Amazon S3 objects based on their LastModifiedDate attribute in Amazon S3 .
	LastModifiedDateCondition interface{} `field:"optional" json:"lastModifiedDateCondition" yaml:"lastModifiedDateCondition"`
	// A structure that maps names of parameters used in the Amazon S3 path of a dataset to their definitions.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// Represents a single entry in the path parameters of a dataset.
//
// Each `PathParameter` consists of a name and a parameter definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathParameterProperty := &pathParameterProperty{
//   	datasetParameter: &datasetParameterProperty{
//   		name: jsii.String("name"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		createColumn: jsii.Boolean(false),
//   		datetimeOptions: &datetimeOptionsProperty{
//   			format: jsii.String("format"),
//
//   			// the properties below are optional
//   			localeCode: jsii.String("localeCode"),
//   			timezoneOffset: jsii.String("timezoneOffset"),
//   		},
//   		filter: &filterExpressionProperty{
//   			expression: jsii.String("expression"),
//   			valuesMap: []interface{}{
//   				&filterValueProperty{
//   					value: jsii.String("value"),
//   					valueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   	},
//   	pathParameterName: jsii.String("pathParameterName"),
//   }
//
type CfnDataset_PathParameterProperty struct {
	// The path parameter definition.
	DatasetParameter interface{} `field:"required" json:"datasetParameter" yaml:"datasetParameter"`
	// The name of the path parameter.
	PathParameterName *string `field:"required" json:"pathParameterName" yaml:"pathParameterName"`
}

// Represents an Amazon S3 location (bucket name, bucket owner, and object key) where DataBrew can read input data, or write output from a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	key: jsii.String("key"),
//   }
//
type CfnDataset_S3LocationProperty struct {
	// The Amazon S3 bucket name.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The unique name of the object in the bucket.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetProps := &cfnDatasetProps{
//   	input: &inputProperty{
//   		databaseInputDefinition: &databaseInputDefinitionProperty{
//   			glueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			databaseTableName: jsii.String("databaseTableName"),
//   			queryString: jsii.String("queryString"),
//   			tempDirectory: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   			},
//   		},
//   		dataCatalogInputDefinition: &dataCatalogInputDefinitionProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
//   			tempDirectory: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   			},
//   		},
//   		metadata: &metadataProperty{
//   			sourceArn: jsii.String("sourceArn"),
//   		},
//   		s3InputDefinition: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	format: jsii.String("format"),
//   	formatOptions: &formatOptionsProperty{
//   		csv: &csvOptionsProperty{
//   			delimiter: jsii.String("delimiter"),
//   			headerRow: jsii.Boolean(false),
//   		},
//   		excel: &excelOptionsProperty{
//   			headerRow: jsii.Boolean(false),
//   			sheetIndexes: []interface{}{
//   				jsii.Number(123),
//   			},
//   			sheetNames: []*string{
//   				jsii.String("sheetNames"),
//   			},
//   		},
//   		json: &jsonOptionsProperty{
//   			multiLine: jsii.Boolean(false),
//   		},
//   	},
//   	pathOptions: &pathOptionsProperty{
//   		filesLimit: &filesLimitProperty{
//   			maxFiles: jsii.Number(123),
//
//   			// the properties below are optional
//   			order: jsii.String("order"),
//   			orderedBy: jsii.String("orderedBy"),
//   		},
//   		lastModifiedDateCondition: &filterExpressionProperty{
//   			expression: jsii.String("expression"),
//   			valuesMap: []interface{}{
//   				&filterValueProperty{
//   					value: jsii.String("value"),
//   					valueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   		parameters: []interface{}{
//   			&pathParameterProperty{
//   				datasetParameter: &datasetParameterProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					createColumn: jsii.Boolean(false),
//   					datetimeOptions: &datetimeOptionsProperty{
//   						format: jsii.String("format"),
//
//   						// the properties below are optional
//   						localeCode: jsii.String("localeCode"),
//   						timezoneOffset: jsii.String("timezoneOffset"),
//   					},
//   					filter: &filterExpressionProperty{
//   						expression: jsii.String("expression"),
//   						valuesMap: []interface{}{
//   							&filterValueProperty{
//   								value: jsii.String("value"),
//   								valueReference: jsii.String("valueReference"),
//   							},
//   						},
//   					},
//   				},
//   				pathParameterName: jsii.String("pathParameterName"),
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
type CfnDatasetProps struct {
	// Information on how DataBrew can find the dataset, in either the AWS Glue Data Catalog or Amazon S3 .
	Input interface{} `field:"required" json:"input" yaml:"input"`
	// The unique name of the dataset.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The file format of a dataset that is created from an Amazon S3 file or folder.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// A set of options that define how DataBrew interprets the data in the dataset.
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// A set of options that defines how DataBrew interprets an Amazon S3 path of the dataset.
	PathOptions interface{} `field:"optional" json:"pathOptions" yaml:"pathOptions"`
	// Metadata tags that have been applied to the dataset.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataBrew::Job`.
//
// Specifies a new DataBrew job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnJob := awscdk.Aws_databrew.NewCfnJob(this, jsii.String("MyCfnJob"), &cfnJobProps{
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	databaseOutputs: []interface{}{
//   		&databaseOutputProperty{
//   			databaseOptions: &databaseTableOutputOptionsProperty{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				tempDirectory: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			glueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			databaseOutputMode: jsii.String("databaseOutputMode"),
//   		},
//   	},
//   	dataCatalogOutputs: []interface{}{
//   		&dataCatalogOutputProperty{
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			catalogId: jsii.String("catalogId"),
//   			databaseOptions: &databaseTableOutputOptionsProperty{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				tempDirectory: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			overwrite: jsii.Boolean(false),
//   			s3Options: &s3TableOutputOptionsProperty{
//   				location: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   		},
//   	},
//   	datasetName: jsii.String("datasetName"),
//   	encryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	encryptionMode: jsii.String("encryptionMode"),
//   	jobSample: &jobSampleProperty{
//   		mode: jsii.String("mode"),
//   		size: jsii.Number(123),
//   	},
//   	logSubscription: jsii.String("logSubscription"),
//   	maxCapacity: jsii.Number(123),
//   	maxRetries: jsii.Number(123),
//   	outputLocation: &outputLocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		key: jsii.String("key"),
//   	},
//   	outputs: []interface{}{
//   		&outputProperty{
//   			location: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				bucketOwner: jsii.String("bucketOwner"),
//   				key: jsii.String("key"),
//   			},
//
//   			// the properties below are optional
//   			compressionFormat: jsii.String("compressionFormat"),
//   			format: jsii.String("format"),
//   			formatOptions: &outputFormatOptionsProperty{
//   				csv: &csvOutputOptionsProperty{
//   					delimiter: jsii.String("delimiter"),
//   				},
//   			},
//   			maxOutputFiles: jsii.Number(123),
//   			overwrite: jsii.Boolean(false),
//   			partitionColumns: []*string{
//   				jsii.String("partitionColumns"),
//   			},
//   		},
//   	},
//   	profileConfiguration: &profileConfigurationProperty{
//   		columnStatisticsConfigurations: []interface{}{
//   			&columnStatisticsConfigurationProperty{
//   				statistics: &statisticsConfigurationProperty{
//   					includedStatistics: []*string{
//   						jsii.String("includedStatistics"),
//   					},
//   					overrides: []interface{}{
//   						&statisticOverrideProperty{
//   							parameters: parameters,
//   							statistic: jsii.String("statistic"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				selectors: []interface{}{
//   					&columnSelectorProperty{
//   						name: jsii.String("name"),
//   						regex: jsii.String("regex"),
//   					},
//   				},
//   			},
//   		},
//   		datasetStatisticsConfiguration: &statisticsConfigurationProperty{
//   			includedStatistics: []*string{
//   				jsii.String("includedStatistics"),
//   			},
//   			overrides: []interface{}{
//   				&statisticOverrideProperty{
//   					parameters: parameters,
//   					statistic: jsii.String("statistic"),
//   				},
//   			},
//   		},
//   		entityDetectorConfiguration: &entityDetectorConfigurationProperty{
//   			entityTypes: []*string{
//   				jsii.String("entityTypes"),
//   			},
//
//   			// the properties below are optional
//   			allowedStatistics: &allowedStatisticsProperty{
//   				statistics: []*string{
//   					jsii.String("statistics"),
//   				},
//   			},
//   		},
//   		profileColumns: []interface{}{
//   			&columnSelectorProperty{
//   				name: jsii.String("name"),
//   				regex: jsii.String("regex"),
//   			},
//   		},
//   	},
//   	projectName: jsii.String("projectName"),
//   	recipe: &recipeProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeout: jsii.Number(123),
//   	validationConfigurations: []interface{}{
//   		&validationConfigurationProperty{
//   			rulesetArn: jsii.String("rulesetArn"),
//
//   			// the properties below are optional
//   			validationMode: jsii.String("validationMode"),
//   		},
//   	},
//   })
//
type CfnJob interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.
	DatabaseOutputs() interface{}
	SetDatabaseOutputs(val interface{})
	// One or more artifacts that represent the AWS Glue Data Catalog output from running the job.
	DataCatalogOutputs() interface{}
	SetDataCatalogOutputs(val interface{})
	// A dataset that the job is to process.
	DatasetName() *string
	SetDatasetName(val *string)
	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the job output.
	//
	// For more information, see [Encrypting data written by DataBrew jobs](https://docs.aws.amazon.com/databrew/latest/dg/encryption-security-configuration.html)
	EncryptionKeyArn() *string
	SetEncryptionKeyArn(val *string)
	// The encryption mode for the job, which can be one of the following:.
	//
	// - `SSE-KMS` - Server-side encryption with keys managed by AWS KMS .
	// - `SSE-S3` - Server-side encryption with keys managed by Amazon S3.
	EncryptionMode() *string
	SetEncryptionMode(val *string)
	// A sample configuration for profile jobs only, which determines the number of rows on which the profile job is run.
	//
	// If a `JobSample` value isn't provided, the default value is used. The default value is CUSTOM_ROWS for the mode parameter and 20,000 for the size parameter.
	JobSample() interface{}
	SetJobSample(val interface{})
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
	// The current status of Amazon CloudWatch logging for the job.
	LogSubscription() *string
	SetLogSubscription(val *string)
	// The maximum number of nodes that can be consumed when the job processes data.
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	// The maximum number of times to retry the job after a job run fails.
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	// The unique name of the job.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// `AWS::DataBrew::Job.OutputLocation`.
	OutputLocation() interface{}
	SetOutputLocation(val interface{})
	// One or more artifacts that represent output from running the job.
	Outputs() interface{}
	SetOutputs(val interface{})
	// Configuration for profile jobs.
	//
	// Configuration can be used to select columns, do evaluations, and override default parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all supported columns.
	ProfileConfiguration() interface{}
	SetProfileConfiguration(val interface{})
	// The name of the project that the job is associated with.
	ProjectName() *string
	SetProjectName(val *string)
	// A series of data transformation steps that the job runs.
	Recipe() interface{}
	SetRecipe(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the role to be assumed for this job.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata tags that have been applied to the job.
	Tags() awscdk.TagManager
	// The job's timeout in minutes.
	//
	// A job that attempts to run longer than this timeout period ends with a status of `TIMEOUT` .
	Timeout() *float64
	SetTimeout(val *float64)
	// The job type of the job, which must be one of the following:.
	//
	// - `PROFILE` - A job to analyze a dataset, to determine its size, data types, data distribution, and more.
	// - `RECIPE` - A job to apply one or more transformations to a dataset.
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
	// List of validation configurations that are applied to the profile job.
	ValidationConfigurations() interface{}
	SetValidationConfigurations(val interface{})
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

// The jsii proxy struct for CfnJob
type jsiiProxy_CfnJob struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJob) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DatabaseOutputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DataCatalogOutputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCatalogOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) JobSample() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobSample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) LogSubscription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) OutputLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Outputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ProfileConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Recipe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recipe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ValidationConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationConfigurations",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Job`.
func NewCfnJob(scope constructs.Construct, id *string, props *CfnJobProps) CfnJob {
	_init_.Initialize()

	j := jsiiProxy_CfnJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Job`.
func NewCfnJob_Override(c CfnJob, scope constructs.Construct, id *string, props *CfnJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnJob",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJob) SetDatabaseOutputs(val interface{}) {
	_jsii_.Set(
		j,
		"databaseOutputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetDataCatalogOutputs(val interface{}) {
	_jsii_.Set(
		j,
		"dataCatalogOutputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetDatasetName(val *string) {
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetEncryptionKeyArn(val *string) {
	_jsii_.Set(
		j,
		"encryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetJobSample(val interface{}) {
	_jsii_.Set(
		j,
		"jobSample",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetLogSubscription(val *string) {
	_jsii_.Set(
		j,
		"logSubscription",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetOutputLocation(val interface{}) {
	_jsii_.Set(
		j,
		"outputLocation",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetOutputs(val interface{}) {
	_jsii_.Set(
		j,
		"outputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetProfileConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"profileConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetProjectName(val *string) {
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetRecipe(val interface{}) {
	_jsii_.Set(
		j,
		"recipe",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetValidationConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"validationConfigurations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnJob_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnJob_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
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
func CfnJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJob_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJob) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnJob) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJob) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnJob) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnJob) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnJob) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnJob) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnJob) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJob) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Configuration of statistics that are allowed to be run on columns that contain detected entities.
//
// When undefined, no statistics will be computed on columns that contain detected entities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allowedStatisticsProperty := &allowedStatisticsProperty{
//   	statistics: []*string{
//   		jsii.String("statistics"),
//   	},
//   }
//
type CfnJob_AllowedStatisticsProperty struct {
	// One or more column statistics to allow for columns that contain detected entities.
	Statistics *[]*string `field:"required" json:"statistics" yaml:"statistics"`
}

// Selector of a column from a dataset for profile job configuration.
//
// One selector includes either a column name or a regular expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnSelectorProperty := &columnSelectorProperty{
//   	name: jsii.String("name"),
//   	regex: jsii.String("regex"),
//   }
//
type CfnJob_ColumnSelectorProperty struct {
	// The name of a column from a dataset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A regular expression for selecting a column from a dataset.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

// Configuration for column evaluations for a profile job.
//
// ColumnStatisticsConfiguration can be used to select evaluations and override parameters of evaluations for particular columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   columnStatisticsConfigurationProperty := &columnStatisticsConfigurationProperty{
//   	statistics: &statisticsConfigurationProperty{
//   		includedStatistics: []*string{
//   			jsii.String("includedStatistics"),
//   		},
//   		overrides: []interface{}{
//   			&statisticOverrideProperty{
//   				parameters: parameters,
//   				statistic: jsii.String("statistic"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	selectors: []interface{}{
//   		&columnSelectorProperty{
//   			name: jsii.String("name"),
//   			regex: jsii.String("regex"),
//   		},
//   	},
//   }
//
type CfnJob_ColumnStatisticsConfigurationProperty struct {
	// Configuration for evaluations.
	//
	// Statistics can be used to select evaluations and override parameters of evaluations.
	Statistics interface{} `field:"required" json:"statistics" yaml:"statistics"`
	// List of column selectors.
	//
	// Selectors can be used to select columns from the dataset. When selectors are undefined, configuration will be applied to all supported columns.
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}

// Represents a set of options that define how DataBrew will write a comma-separated value (CSV) file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvOutputOptionsProperty := &csvOutputOptionsProperty{
//   	delimiter: jsii.String("delimiter"),
//   }
//
type CfnJob_CsvOutputOptionsProperty struct {
	// A single character that specifies the delimiter used to create CSV job output.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
}

// Represents options that specify how and where in the AWS Glue Data Catalog DataBrew writes the output generated by recipe jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogOutputProperty := &dataCatalogOutputProperty{
//   	databaseName: jsii.String("databaseName"),
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	catalogId: jsii.String("catalogId"),
//   	databaseOptions: &databaseTableOutputOptionsProperty{
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		tempDirectory: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			bucketOwner: jsii.String("bucketOwner"),
//   			key: jsii.String("key"),
//   		},
//   	},
//   	overwrite: jsii.Boolean(false),
//   	s3Options: &s3TableOutputOptionsProperty{
//   		location: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			bucketOwner: jsii.String("bucketOwner"),
//   			key: jsii.String("key"),
//   		},
//   	},
//   }
//
type CfnJob_DataCatalogOutputProperty struct {
	// The name of a database in the Data Catalog.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of a table in the Data Catalog.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// The unique identifier of the AWS account that holds the Data Catalog that stores the data.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Represents options that specify how and where DataBrew writes the database output generated by recipe jobs.
	DatabaseOptions interface{} `field:"optional" json:"databaseOptions" yaml:"databaseOptions"`
	// A value that, if true, means that any data in the location specified for output is overwritten with new output.
	//
	// Not supported with DatabaseOptions.
	Overwrite interface{} `field:"optional" json:"overwrite" yaml:"overwrite"`
	// Represents options that specify how and where DataBrew writes the Amazon S3 output generated by recipe jobs.
	S3Options interface{} `field:"optional" json:"s3Options" yaml:"s3Options"`
}

// Represents a JDBC database output object which defines the output destination for a DataBrew recipe job to write into.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseOutputProperty := &databaseOutputProperty{
//   	databaseOptions: &databaseTableOutputOptionsProperty{
//   		tableName: jsii.String("tableName"),
//
//   		// the properties below are optional
//   		tempDirectory: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			bucketOwner: jsii.String("bucketOwner"),
//   			key: jsii.String("key"),
//   		},
//   	},
//   	glueConnectionName: jsii.String("glueConnectionName"),
//
//   	// the properties below are optional
//   	databaseOutputMode: jsii.String("databaseOutputMode"),
//   }
//
type CfnJob_DatabaseOutputProperty struct {
	// Represents options that specify how and where DataBrew writes the database output generated by recipe jobs.
	DatabaseOptions interface{} `field:"required" json:"databaseOptions" yaml:"databaseOptions"`
	// The AWS Glue connection that stores the connection information for the target database.
	GlueConnectionName *string `field:"required" json:"glueConnectionName" yaml:"glueConnectionName"`
	// The output mode to write into the database.
	//
	// Currently supported option: NEW_TABLE.
	DatabaseOutputMode *string `field:"optional" json:"databaseOutputMode" yaml:"databaseOutputMode"`
}

// Represents options that specify how and where DataBrew writes the database output generated by recipe jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseTableOutputOptionsProperty := &databaseTableOutputOptionsProperty{
//   	tableName: jsii.String("tableName"),
//
//   	// the properties below are optional
//   	tempDirectory: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnJob_DatabaseTableOutputOptionsProperty struct {
	// A prefix for the name of a table DataBrew will create in the database.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Represents an Amazon S3 location (bucket name and object key) where DataBrew can store intermediate results.
	TempDirectory interface{} `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

// Configuration of entity detection for a profile job.
//
// When undefined, entity detection is disabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityDetectorConfigurationProperty := &entityDetectorConfigurationProperty{
//   	entityTypes: []*string{
//   		jsii.String("entityTypes"),
//   	},
//
//   	// the properties below are optional
//   	allowedStatistics: &allowedStatisticsProperty{
//   		statistics: []*string{
//   			jsii.String("statistics"),
//   		},
//   	},
//   }
//
type CfnJob_EntityDetectorConfigurationProperty struct {
	// Entity types to detect. Can be any of the following:.
	//
	// - USA_SSN
	// - EMAIL
	// - USA_ITIN
	// - USA_PASSPORT_NUMBER
	// - PHONE_NUMBER
	// - USA_DRIVING_LICENSE
	// - BANK_ACCOUNT
	// - CREDIT_CARD
	// - IP_ADDRESS
	// - MAC_ADDRESS
	// - USA_DEA_NUMBER
	// - USA_HCPCS_CODE
	// - USA_NATIONAL_PROVIDER_IDENTIFIER
	// - USA_NATIONAL_DRUG_CODE
	// - USA_HEALTH_INSURANCE_CLAIM_NUMBER
	// - USA_MEDICARE_BENEFICIARY_IDENTIFIER
	// - USA_CPT_CODE
	// - PERSON_NAME
	// - DATE
	//
	// The Entity type group USA_ALL is also supported, and includes all of the above entity types except PERSON_NAME and DATE.
	EntityTypes *[]*string `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// Configuration of statistics that are allowed to be run on columns that contain detected entities.
	//
	// When undefined, no statistics will be computed on columns that contain detected entities.
	AllowedStatistics interface{} `field:"optional" json:"allowedStatistics" yaml:"allowedStatistics"`
}

// A sample configuration for profile jobs only, which determines the number of rows on which the profile job is run.
//
// If a `JobSample` value isn't provided, the default is used. The default value is CUSTOM_ROWS for the mode parameter and 20,000 for the size parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobSampleProperty := &jobSampleProperty{
//   	mode: jsii.String("mode"),
//   	size: jsii.Number(123),
//   }
//
type CfnJob_JobSampleProperty struct {
	// A value that determines whether the profile job is run on the entire dataset or a specified number of rows.
	//
	// This value must be one of the following:
	//
	// - FULL_DATASET - The profile job is run on the entire dataset.
	// - CUSTOM_ROWS - The profile job is run on the number of rows specified in the `Size` parameter.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The `Size` parameter is only required when the mode is CUSTOM_ROWS.
	//
	// The profile job is run on the specified number of rows. The maximum value for size is Long.MAX_VALUE.
	//
	// Long.MAX_VALUE = 9223372036854775807
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

// Represents a set of options that define the structure of comma-separated (CSV) job output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputFormatOptionsProperty := &outputFormatOptionsProperty{
//   	csv: &csvOutputOptionsProperty{
//   		delimiter: jsii.String("delimiter"),
//   	},
//   }
//
type CfnJob_OutputFormatOptionsProperty struct {
	// Represents a set of options that define the structure of comma-separated value (CSV) job output.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
}

// The location in Amazon S3 or AWS Glue Data Catalog where the job writes its output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputLocationProperty := &outputLocationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	bucketOwner: jsii.String("bucketOwner"),
//   	key: jsii.String("key"),
//   }
//
type CfnJob_OutputLocationProperty struct {
	// The Amazon S3 bucket name.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnJob.OutputLocationProperty.BucketOwner`.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The unique name of the object in the bucket.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

// Represents options that specify how and where in Amazon S3 DataBrew writes the output generated by recipe jobs or profile jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputProperty := &outputProperty{
//   	location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		key: jsii.String("key"),
//   	},
//
//   	// the properties below are optional
//   	compressionFormat: jsii.String("compressionFormat"),
//   	format: jsii.String("format"),
//   	formatOptions: &outputFormatOptionsProperty{
//   		csv: &csvOutputOptionsProperty{
//   			delimiter: jsii.String("delimiter"),
//   		},
//   	},
//   	maxOutputFiles: jsii.Number(123),
//   	overwrite: jsii.Boolean(false),
//   	partitionColumns: []*string{
//   		jsii.String("partitionColumns"),
//   	},
//   }
//
type CfnJob_OutputProperty struct {
	// The location in Amazon S3 where the job writes its output.
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// The compression algorithm used to compress the output text of the job.
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// The data format of the output of the job.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Represents options that define how DataBrew formats job output files.
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
	// The maximum number of files to be generated by the job and written to the output folder.
	MaxOutputFiles *float64 `field:"optional" json:"maxOutputFiles" yaml:"maxOutputFiles"`
	// A value that, if true, means that any data in the location specified for output is overwritten with new output.
	Overwrite interface{} `field:"optional" json:"overwrite" yaml:"overwrite"`
	// The names of one or more partition columns for the output of the job.
	PartitionColumns *[]*string `field:"optional" json:"partitionColumns" yaml:"partitionColumns"`
}

// Configuration for profile jobs.
//
// Configuration can be used to select columns, do evaluations, and override default parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all supported columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   profileConfigurationProperty := &profileConfigurationProperty{
//   	columnStatisticsConfigurations: []interface{}{
//   		&columnStatisticsConfigurationProperty{
//   			statistics: &statisticsConfigurationProperty{
//   				includedStatistics: []*string{
//   					jsii.String("includedStatistics"),
//   				},
//   				overrides: []interface{}{
//   					&statisticOverrideProperty{
//   						parameters: parameters,
//   						statistic: jsii.String("statistic"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			selectors: []interface{}{
//   				&columnSelectorProperty{
//   					name: jsii.String("name"),
//   					regex: jsii.String("regex"),
//   				},
//   			},
//   		},
//   	},
//   	datasetStatisticsConfiguration: &statisticsConfigurationProperty{
//   		includedStatistics: []*string{
//   			jsii.String("includedStatistics"),
//   		},
//   		overrides: []interface{}{
//   			&statisticOverrideProperty{
//   				parameters: parameters,
//   				statistic: jsii.String("statistic"),
//   			},
//   		},
//   	},
//   	entityDetectorConfiguration: &entityDetectorConfigurationProperty{
//   		entityTypes: []*string{
//   			jsii.String("entityTypes"),
//   		},
//
//   		// the properties below are optional
//   		allowedStatistics: &allowedStatisticsProperty{
//   			statistics: []*string{
//   				jsii.String("statistics"),
//   			},
//   		},
//   	},
//   	profileColumns: []interface{}{
//   		&columnSelectorProperty{
//   			name: jsii.String("name"),
//   			regex: jsii.String("regex"),
//   		},
//   	},
//   }
//
type CfnJob_ProfileConfigurationProperty struct {
	// List of configurations for column evaluations.
	//
	// ColumnStatisticsConfigurations are used to select evaluations and override parameters of evaluations for particular columns. When ColumnStatisticsConfigurations is undefined, the profile job will profile all supported columns and run all supported evaluations.
	ColumnStatisticsConfigurations interface{} `field:"optional" json:"columnStatisticsConfigurations" yaml:"columnStatisticsConfigurations"`
	// Configuration for inter-column evaluations.
	//
	// Configuration can be used to select evaluations and override parameters of evaluations. When configuration is undefined, the profile job will run all supported inter-column evaluations.
	DatasetStatisticsConfiguration interface{} `field:"optional" json:"datasetStatisticsConfiguration" yaml:"datasetStatisticsConfiguration"`
	// Configuration of entity detection for a profile job.
	//
	// When undefined, entity detection is disabled.
	EntityDetectorConfiguration interface{} `field:"optional" json:"entityDetectorConfiguration" yaml:"entityDetectorConfiguration"`
	// List of column selectors.
	//
	// ProfileColumns can be used to select columns from the dataset. When ProfileColumns is undefined, the profile job will profile all supported columns.
	ProfileColumns interface{} `field:"optional" json:"profileColumns" yaml:"profileColumns"`
}

// Represents one or more actions to be performed on a DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipeProperty := &recipeProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	version: jsii.String("version"),
//   }
//
type CfnJob_RecipeProperty struct {
	// The unique name for the recipe.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier for the version for the recipe.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

// Represents an Amazon S3 location (bucket name, bucket owner, and object key) where DataBrew can read input data, or write output from a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	bucketOwner: jsii.String("bucketOwner"),
//   	key: jsii.String("key"),
//   }
//
type CfnJob_S3LocationProperty struct {
	// The Amazon S3 bucket name.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account ID of the bucket owner.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The unique name of the object in the bucket.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

// Represents options that specify how and where DataBrew writes the Amazon S3 output generated by recipe jobs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3TableOutputOptionsProperty := &s3TableOutputOptionsProperty{
//   	location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnJob_S3TableOutputOptionsProperty struct {
	// Represents an Amazon S3 location (bucket name and object key) where DataBrew can write output from a job.
	Location interface{} `field:"required" json:"location" yaml:"location"`
}

// Override of a particular evaluation for a profile job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   statisticOverrideProperty := &statisticOverrideProperty{
//   	parameters: parameters,
//   	statistic: jsii.String("statistic"),
//   }
//
type CfnJob_StatisticOverrideProperty struct {
	// A map that includes overrides of an evaluation’s parameters.
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The name of an evaluation.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}

// Configuration of evaluations for a profile job.
//
// This configuration can be used to select evaluations and override the parameters of selected evaluations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   statisticsConfigurationProperty := &statisticsConfigurationProperty{
//   	includedStatistics: []*string{
//   		jsii.String("includedStatistics"),
//   	},
//   	overrides: []interface{}{
//   		&statisticOverrideProperty{
//   			parameters: parameters,
//   			statistic: jsii.String("statistic"),
//   		},
//   	},
//   }
//
type CfnJob_StatisticsConfigurationProperty struct {
	// List of included evaluations.
	//
	// When the list is undefined, all supported evaluations will be included.
	IncludedStatistics *[]*string `field:"optional" json:"includedStatistics" yaml:"includedStatistics"`
	// List of overrides for evaluations.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

// Configuration for data quality validation.
//
// Used to select the Rulesets and Validation Mode to be used in the profile job. When ValidationConfiguration is null, the profile job will run without data quality validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validationConfigurationProperty := &validationConfigurationProperty{
//   	rulesetArn: jsii.String("rulesetArn"),
//
//   	// the properties below are optional
//   	validationMode: jsii.String("validationMode"),
//   }
//
type CfnJob_ValidationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) for the ruleset to be validated in the profile job.
	//
	// The TargetArn of the selected ruleset should be the same as the Amazon Resource Name (ARN) of the dataset that is associated with the profile job.
	RulesetArn *string `field:"required" json:"rulesetArn" yaml:"rulesetArn"`
	// Mode of data quality validation.
	//
	// Default mode is “CHECK_ALL” which verifies all rules defined in the selected ruleset.
	ValidationMode *string `field:"optional" json:"validationMode" yaml:"validationMode"`
}

// Properties for defining a `CfnJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnJobProps := &cfnJobProps{
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	databaseOutputs: []interface{}{
//   		&databaseOutputProperty{
//   			databaseOptions: &databaseTableOutputOptionsProperty{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				tempDirectory: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			glueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			databaseOutputMode: jsii.String("databaseOutputMode"),
//   		},
//   	},
//   	dataCatalogOutputs: []interface{}{
//   		&dataCatalogOutputProperty{
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			catalogId: jsii.String("catalogId"),
//   			databaseOptions: &databaseTableOutputOptionsProperty{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				tempDirectory: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			overwrite: jsii.Boolean(false),
//   			s3Options: &s3TableOutputOptionsProperty{
//   				location: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   		},
//   	},
//   	datasetName: jsii.String("datasetName"),
//   	encryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	encryptionMode: jsii.String("encryptionMode"),
//   	jobSample: &jobSampleProperty{
//   		mode: jsii.String("mode"),
//   		size: jsii.Number(123),
//   	},
//   	logSubscription: jsii.String("logSubscription"),
//   	maxCapacity: jsii.Number(123),
//   	maxRetries: jsii.Number(123),
//   	outputLocation: &outputLocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		key: jsii.String("key"),
//   	},
//   	outputs: []interface{}{
//   		&outputProperty{
//   			location: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				bucketOwner: jsii.String("bucketOwner"),
//   				key: jsii.String("key"),
//   			},
//
//   			// the properties below are optional
//   			compressionFormat: jsii.String("compressionFormat"),
//   			format: jsii.String("format"),
//   			formatOptions: &outputFormatOptionsProperty{
//   				csv: &csvOutputOptionsProperty{
//   					delimiter: jsii.String("delimiter"),
//   				},
//   			},
//   			maxOutputFiles: jsii.Number(123),
//   			overwrite: jsii.Boolean(false),
//   			partitionColumns: []*string{
//   				jsii.String("partitionColumns"),
//   			},
//   		},
//   	},
//   	profileConfiguration: &profileConfigurationProperty{
//   		columnStatisticsConfigurations: []interface{}{
//   			&columnStatisticsConfigurationProperty{
//   				statistics: &statisticsConfigurationProperty{
//   					includedStatistics: []*string{
//   						jsii.String("includedStatistics"),
//   					},
//   					overrides: []interface{}{
//   						&statisticOverrideProperty{
//   							parameters: parameters,
//   							statistic: jsii.String("statistic"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				selectors: []interface{}{
//   					&columnSelectorProperty{
//   						name: jsii.String("name"),
//   						regex: jsii.String("regex"),
//   					},
//   				},
//   			},
//   		},
//   		datasetStatisticsConfiguration: &statisticsConfigurationProperty{
//   			includedStatistics: []*string{
//   				jsii.String("includedStatistics"),
//   			},
//   			overrides: []interface{}{
//   				&statisticOverrideProperty{
//   					parameters: parameters,
//   					statistic: jsii.String("statistic"),
//   				},
//   			},
//   		},
//   		entityDetectorConfiguration: &entityDetectorConfigurationProperty{
//   			entityTypes: []*string{
//   				jsii.String("entityTypes"),
//   			},
//
//   			// the properties below are optional
//   			allowedStatistics: &allowedStatisticsProperty{
//   				statistics: []*string{
//   					jsii.String("statistics"),
//   				},
//   			},
//   		},
//   		profileColumns: []interface{}{
//   			&columnSelectorProperty{
//   				name: jsii.String("name"),
//   				regex: jsii.String("regex"),
//   			},
//   		},
//   	},
//   	projectName: jsii.String("projectName"),
//   	recipe: &recipeProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeout: jsii.Number(123),
//   	validationConfigurations: []interface{}{
//   		&validationConfigurationProperty{
//   			rulesetArn: jsii.String("rulesetArn"),
//
//   			// the properties below are optional
//   			validationMode: jsii.String("validationMode"),
//   		},
//   	},
//   }
//
type CfnJobProps struct {
	// The unique name of the job.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role to be assumed for this job.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The job type of the job, which must be one of the following:.
	//
	// - `PROFILE` - A job to analyze a dataset, to determine its size, data types, data distribution, and more.
	// - `RECIPE` - A job to apply one or more transformations to a dataset.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.
	DatabaseOutputs interface{} `field:"optional" json:"databaseOutputs" yaml:"databaseOutputs"`
	// One or more artifacts that represent the AWS Glue Data Catalog output from running the job.
	DataCatalogOutputs interface{} `field:"optional" json:"dataCatalogOutputs" yaml:"dataCatalogOutputs"`
	// A dataset that the job is to process.
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the job output.
	//
	// For more information, see [Encrypting data written by DataBrew jobs](https://docs.aws.amazon.com/databrew/latest/dg/encryption-security-configuration.html)
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The encryption mode for the job, which can be one of the following:.
	//
	// - `SSE-KMS` - Server-side encryption with keys managed by AWS KMS .
	// - `SSE-S3` - Server-side encryption with keys managed by Amazon S3.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// A sample configuration for profile jobs only, which determines the number of rows on which the profile job is run.
	//
	// If a `JobSample` value isn't provided, the default value is used. The default value is CUSTOM_ROWS for the mode parameter and 20,000 for the size parameter.
	JobSample interface{} `field:"optional" json:"jobSample" yaml:"jobSample"`
	// The current status of Amazon CloudWatch logging for the job.
	LogSubscription *string `field:"optional" json:"logSubscription" yaml:"logSubscription"`
	// The maximum number of nodes that can be consumed when the job processes data.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry the job after a job run fails.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// `AWS::DataBrew::Job.OutputLocation`.
	OutputLocation interface{} `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// One or more artifacts that represent output from running the job.
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// Configuration for profile jobs.
	//
	// Configuration can be used to select columns, do evaluations, and override default parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all supported columns.
	ProfileConfiguration interface{} `field:"optional" json:"profileConfiguration" yaml:"profileConfiguration"`
	// The name of the project that the job is associated with.
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// A series of data transformation steps that the job runs.
	Recipe interface{} `field:"optional" json:"recipe" yaml:"recipe"`
	// Metadata tags that have been applied to the job.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The job's timeout in minutes.
	//
	// A job that attempts to run longer than this timeout period ends with a status of `TIMEOUT` .
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// List of validation configurations that are applied to the profile job.
	ValidationConfigurations interface{} `field:"optional" json:"validationConfigurations" yaml:"validationConfigurations"`
}

// A CloudFormation `AWS::DataBrew::Project`.
//
// Specifies a new AWS Glue DataBrew project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProject := awscdk.Aws_databrew.NewCfnProject(this, jsii.String("MyCfnProject"), &cfnProjectProps{
//   	datasetName: jsii.String("datasetName"),
//   	name: jsii.String("name"),
//   	recipeName: jsii.String("recipeName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	sample: &sampleProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		size: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnProject interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The dataset that the project is to act upon.
	DatasetName() *string
	SetDatasetName(val *string)
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
	// The unique name of a project.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The name of a recipe that will be developed during a project session.
	RecipeName() *string
	SetRecipeName(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the role that will be assumed for this project.
	RoleArn() *string
	SetRoleArn(val *string)
	// The sample size and sampling type to apply to the data.
	//
	// If this parameter isn't specified, then the sample consists of the first 500 rows from the dataset.
	Sample() interface{}
	SetSample(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Metadata tags that have been applied to the project.
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

// The jsii proxy struct for CfnProject
type jsiiProxy_CfnProject struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnProject) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) RecipeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Sample() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Project`.
func NewCfnProject(scope constructs.Construct, id *string, props *CfnProjectProps) CfnProject {
	_init_.Initialize()

	j := jsiiProxy_CfnProject{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Project`.
func NewCfnProject_Override(c CfnProject, scope constructs.Construct, id *string, props *CfnProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnProject",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProject) SetDatasetName(val *string) {
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetRecipeName(val *string) {
	_jsii_.Set(
		j,
		"recipeName",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetSample(val interface{}) {
	_jsii_.Set(
		j,
		"sample",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnProject_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnProject",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnProject_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnProject",
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
func CfnProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProject_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnProject",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProject) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnProject) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProject) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnProject) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnProject) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnProject) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnProject) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnProject) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnProject) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Represents the sample size and sampling type for DataBrew to use for interactive data analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sampleProperty := &sampleProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	size: jsii.Number(123),
//   }
//
type CfnProject_SampleProperty struct {
	// The way in which DataBrew obtains rows from a dataset.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The number of rows in the sample.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &cfnProjectProps{
//   	datasetName: jsii.String("datasetName"),
//   	name: jsii.String("name"),
//   	recipeName: jsii.String("recipeName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	sample: &sampleProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		size: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProjectProps struct {
	// The dataset that the project is to act upon.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// The unique name of a project.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of a recipe that will be developed during a project session.
	RecipeName *string `field:"required" json:"recipeName" yaml:"recipeName"`
	// The Amazon Resource Name (ARN) of the role that will be assumed for this project.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The sample size and sampling type to apply to the data.
	//
	// If this parameter isn't specified, then the sample consists of the first 500 rows from the dataset.
	Sample interface{} `field:"optional" json:"sample" yaml:"sample"`
	// Metadata tags that have been applied to the project.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataBrew::Recipe`.
//
// Specifies a new AWS Glue DataBrew transformation recipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecipe := awscdk.Aws_databrew.NewCfnRecipe(this, jsii.String("MyCfnRecipe"), &cfnRecipeProps{
//   	name: jsii.String("name"),
//   	steps: []interface{}{
//   		&recipeStepProperty{
//   			action: &actionProperty{
//   				operation: jsii.String("operation"),
//
//   				// the properties below are optional
//   				parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			conditionExpressions: []interface{}{
//   				&conditionExpressionProperty{
//   					condition: jsii.String("condition"),
//   					targetColumn: jsii.String("targetColumn"),
//
//   					// the properties below are optional
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnRecipe interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the recipe.
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
	// The unique name for the recipe.
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
	// A list of steps that are defined by the recipe.
	Steps() interface{}
	SetSteps(val interface{})
	// Metadata tags that have been applied to the recipe.
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

// The jsii proxy struct for CfnRecipe
type jsiiProxy_CfnRecipe struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRecipe) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Steps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecipe) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Recipe`.
func NewCfnRecipe(scope constructs.Construct, id *string, props *CfnRecipeProps) CfnRecipe {
	_init_.Initialize()

	j := jsiiProxy_CfnRecipe{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Recipe`.
func NewCfnRecipe_Override(c CfnRecipe, scope constructs.Construct, id *string, props *CfnRecipeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRecipe) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRecipe) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRecipe) SetSteps(val interface{}) {
	_jsii_.Set(
		j,
		"steps",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRecipe_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRecipe_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
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
func CfnRecipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecipe_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnRecipe",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRecipe) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRecipe) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRecipe) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRecipe) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRecipe) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRecipe) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRecipe) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRecipe) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecipe) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecipe) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRecipe) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRecipe) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecipe) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecipe) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRecipe) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Represents a transformation and associated parameters that are used to apply a change to an AWS Glue DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &actionProperty{
//   	operation: jsii.String("operation"),
//
//   	// the properties below are optional
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   }
//
type CfnRecipe_ActionProperty struct {
	// The name of a valid DataBrew transformation to be performed on the data.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Contextual parameters for the transformation.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// Represents an individual condition that evaluates to true or false.
//
// Conditions are used with recipe actions. The action is only performed for column values where the condition evaluates to true.
//
// If a recipe requires more than one condition, then the recipe must specify multiple `ConditionExpression` elements. Each condition is applied to the rows in a dataset first, before the recipe action is performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionExpressionProperty := &conditionExpressionProperty{
//   	condition: jsii.String("condition"),
//   	targetColumn: jsii.String("targetColumn"),
//
//   	// the properties below are optional
//   	value: jsii.String("value"),
//   }
//
type CfnRecipe_ConditionExpressionProperty struct {
	// A specific condition to apply to a recipe action.
	//
	// For more information, see [Recipe structure](https://docs.aws.amazon.com/databrew/latest/dg/recipe-structure.html) in the *AWS Glue DataBrew Developer Guide* .
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// A column to apply this condition to.
	TargetColumn *string `field:"required" json:"targetColumn" yaml:"targetColumn"`
	// A value that the condition must evaluate to for the condition to succeed.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Represents how metadata stored in the AWS Glue Data Catalog is defined in a DataBrew dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogInputDefinitionProperty := &dataCatalogInputDefinitionProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	tableName: jsii.String("tableName"),
//   	tempDirectory: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnRecipe_DataCatalogInputDefinitionProperty struct {
	// The unique identifier of the AWS account that holds the Data Catalog that stores the data.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of a database in the Data Catalog.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of a database table in the Data Catalog.
	//
	// This table corresponds to a DataBrew dataset.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Represents an Amazon location where DataBrew can store intermediate results.
	TempDirectory interface{} `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

// Parameters that are used as inputs for various recipe actions.
//
// The parameters are specific to the context in which they're used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var input interface{}
//
//   recipeParametersProperty := &recipeParametersProperty{
//   	aggregateFunction: jsii.String("aggregateFunction"),
//   	base: jsii.String("base"),
//   	caseStatement: jsii.String("caseStatement"),
//   	categoryMap: jsii.String("categoryMap"),
//   	charsToRemove: jsii.String("charsToRemove"),
//   	collapseConsecutiveWhitespace: jsii.String("collapseConsecutiveWhitespace"),
//   	columnDataType: jsii.String("columnDataType"),
//   	columnRange: jsii.String("columnRange"),
//   	count: jsii.String("count"),
//   	customCharacters: jsii.String("customCharacters"),
//   	customStopWords: jsii.String("customStopWords"),
//   	customValue: jsii.String("customValue"),
//   	datasetsColumns: jsii.String("datasetsColumns"),
//   	dateAddValue: jsii.String("dateAddValue"),
//   	dateTimeFormat: jsii.String("dateTimeFormat"),
//   	dateTimeParameters: jsii.String("dateTimeParameters"),
//   	deleteOtherRows: jsii.String("deleteOtherRows"),
//   	delimiter: jsii.String("delimiter"),
//   	endPattern: jsii.String("endPattern"),
//   	endPosition: jsii.String("endPosition"),
//   	endValue: jsii.String("endValue"),
//   	expandContractions: jsii.String("expandContractions"),
//   	exponent: jsii.String("exponent"),
//   	falseString: jsii.String("falseString"),
//   	groupByAggFunctionOptions: jsii.String("groupByAggFunctionOptions"),
//   	groupByColumns: jsii.String("groupByColumns"),
//   	hiddenColumns: jsii.String("hiddenColumns"),
//   	ignoreCase: jsii.String("ignoreCase"),
//   	includeInSplit: jsii.String("includeInSplit"),
//   	input: input,
//   	interval: jsii.String("interval"),
//   	isText: jsii.String("isText"),
//   	joinKeys: jsii.String("joinKeys"),
//   	joinType: jsii.String("joinType"),
//   	leftColumns: jsii.String("leftColumns"),
//   	limit: jsii.String("limit"),
//   	lowerBound: jsii.String("lowerBound"),
//   	mapType: jsii.String("mapType"),
//   	modeType: jsii.String("modeType"),
//   	multiLine: jsii.Boolean(false),
//   	numRows: jsii.String("numRows"),
//   	numRowsAfter: jsii.String("numRowsAfter"),
//   	numRowsBefore: jsii.String("numRowsBefore"),
//   	orderByColumn: jsii.String("orderByColumn"),
//   	orderByColumns: jsii.String("orderByColumns"),
//   	other: jsii.String("other"),
//   	pattern: jsii.String("pattern"),
//   	patternOption1: jsii.String("patternOption1"),
//   	patternOption2: jsii.String("patternOption2"),
//   	patternOptions: jsii.String("patternOptions"),
//   	period: jsii.String("period"),
//   	position: jsii.String("position"),
//   	removeAllPunctuation: jsii.String("removeAllPunctuation"),
//   	removeAllQuotes: jsii.String("removeAllQuotes"),
//   	removeAllWhitespace: jsii.String("removeAllWhitespace"),
//   	removeCustomCharacters: jsii.String("removeCustomCharacters"),
//   	removeCustomValue: jsii.String("removeCustomValue"),
//   	removeLeadingAndTrailingPunctuation: jsii.String("removeLeadingAndTrailingPunctuation"),
//   	removeLeadingAndTrailingQuotes: jsii.String("removeLeadingAndTrailingQuotes"),
//   	removeLeadingAndTrailingWhitespace: jsii.String("removeLeadingAndTrailingWhitespace"),
//   	removeLetters: jsii.String("removeLetters"),
//   	removeNumbers: jsii.String("removeNumbers"),
//   	removeSourceColumn: jsii.String("removeSourceColumn"),
//   	removeSpecialCharacters: jsii.String("removeSpecialCharacters"),
//   	rightColumns: jsii.String("rightColumns"),
//   	sampleSize: jsii.String("sampleSize"),
//   	sampleType: jsii.String("sampleType"),
//   	secondaryInputs: []interface{}{
//   		&secondaryInputProperty{
//   			dataCatalogInputDefinition: &dataCatalogInputDefinitionProperty{
//   				catalogId: jsii.String("catalogId"),
//   				databaseName: jsii.String("databaseName"),
//   				tableName: jsii.String("tableName"),
//   				tempDirectory: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					key: jsii.String("key"),
//   				},
//   			},
//   			s3InputDefinition: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				key: jsii.String("key"),
//   			},
//   		},
//   	},
//   	secondInput: jsii.String("secondInput"),
//   	sheetIndexes: []interface{}{
//   		jsii.Number(123),
//   	},
//   	sheetNames: []*string{
//   		jsii.String("sheetNames"),
//   	},
//   	sourceColumn: jsii.String("sourceColumn"),
//   	sourceColumn1: jsii.String("sourceColumn1"),
//   	sourceColumn2: jsii.String("sourceColumn2"),
//   	sourceColumns: jsii.String("sourceColumns"),
//   	startColumnIndex: jsii.String("startColumnIndex"),
//   	startPattern: jsii.String("startPattern"),
//   	startPosition: jsii.String("startPosition"),
//   	startValue: jsii.String("startValue"),
//   	stemmingMode: jsii.String("stemmingMode"),
//   	stepCount: jsii.String("stepCount"),
//   	stepIndex: jsii.String("stepIndex"),
//   	stopWordsMode: jsii.String("stopWordsMode"),
//   	strategy: jsii.String("strategy"),
//   	targetColumn: jsii.String("targetColumn"),
//   	targetColumnNames: jsii.String("targetColumnNames"),
//   	targetDateFormat: jsii.String("targetDateFormat"),
//   	targetIndex: jsii.String("targetIndex"),
//   	timeZone: jsii.String("timeZone"),
//   	tokenizerPattern: jsii.String("tokenizerPattern"),
//   	trueString: jsii.String("trueString"),
//   	udfLang: jsii.String("udfLang"),
//   	units: jsii.String("units"),
//   	unpivotColumn: jsii.String("unpivotColumn"),
//   	upperBound: jsii.String("upperBound"),
//   	useNewDataFrame: jsii.String("useNewDataFrame"),
//   	value: jsii.String("value"),
//   	value1: jsii.String("value1"),
//   	value2: jsii.String("value2"),
//   	valueColumn: jsii.String("valueColumn"),
//   	viewFrame: jsii.String("viewFrame"),
//   }
//
type CfnRecipe_RecipeParametersProperty struct {
	// The name of an aggregation function to apply.
	AggregateFunction *string `field:"optional" json:"aggregateFunction" yaml:"aggregateFunction"`
	// The number of digits used in a counting system.
	Base *string `field:"optional" json:"base" yaml:"base"`
	// A case statement associated with a recipe.
	CaseStatement *string `field:"optional" json:"caseStatement" yaml:"caseStatement"`
	// A category map used for one-hot encoding.
	CategoryMap *string `field:"optional" json:"categoryMap" yaml:"categoryMap"`
	// Characters to remove from a step that applies one-hot encoding or tokenization.
	CharsToRemove *string `field:"optional" json:"charsToRemove" yaml:"charsToRemove"`
	// Remove any non-word non-punctuation character.
	CollapseConsecutiveWhitespace *string `field:"optional" json:"collapseConsecutiveWhitespace" yaml:"collapseConsecutiveWhitespace"`
	// The data type of the column.
	ColumnDataType *string `field:"optional" json:"columnDataType" yaml:"columnDataType"`
	// A range of columns to which a step is applied.
	ColumnRange *string `field:"optional" json:"columnRange" yaml:"columnRange"`
	// The number of times a string needs to be repeated.
	Count *string `field:"optional" json:"count" yaml:"count"`
	// One or more characters that can be substituted or removed, depending on the context.
	CustomCharacters *string `field:"optional" json:"customCharacters" yaml:"customCharacters"`
	// A list of words to ignore in a step that applies word tokenization.
	CustomStopWords *string `field:"optional" json:"customStopWords" yaml:"customStopWords"`
	// A list of custom values to use in a step that requires that you provide a value to finish the operation.
	CustomValue *string `field:"optional" json:"customValue" yaml:"customValue"`
	// A list of the dataset columns included in a project.
	DatasetsColumns *string `field:"optional" json:"datasetsColumns" yaml:"datasetsColumns"`
	// A value that specifies how many units of time to add or subtract for a date math operation.
	DateAddValue *string `field:"optional" json:"dateAddValue" yaml:"dateAddValue"`
	// A date format to apply to a date.
	DateTimeFormat *string `field:"optional" json:"dateTimeFormat" yaml:"dateTimeFormat"`
	// A set of parameters associated with a datetime.
	DateTimeParameters *string `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// Determines whether unmapped rows in a categorical mapping should be deleted.
	DeleteOtherRows *string `field:"optional" json:"deleteOtherRows" yaml:"deleteOtherRows"`
	// The delimiter to use when parsing separated values in a text file.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The end pattern to locate.
	EndPattern *string `field:"optional" json:"endPattern" yaml:"endPattern"`
	// The end position to locate.
	EndPosition *string `field:"optional" json:"endPosition" yaml:"endPosition"`
	// The end value to locate.
	EndValue *string `field:"optional" json:"endValue" yaml:"endValue"`
	// A list of word contractions and what they expand to.
	//
	// For eample: *can't* ; *cannot* ; *can not* .
	ExpandContractions *string `field:"optional" json:"expandContractions" yaml:"expandContractions"`
	// The exponent to apply in an exponential operation.
	Exponent *string `field:"optional" json:"exponent" yaml:"exponent"`
	// A value that represents `FALSE` .
	FalseString *string `field:"optional" json:"falseString" yaml:"falseString"`
	// Specifies options to apply to the `GROUP BY` used in an aggregation.
	GroupByAggFunctionOptions *string `field:"optional" json:"groupByAggFunctionOptions" yaml:"groupByAggFunctionOptions"`
	// The columns to use in the `GROUP BY` clause.
	GroupByColumns *string `field:"optional" json:"groupByColumns" yaml:"groupByColumns"`
	// A list of columns to hide.
	HiddenColumns *string `field:"optional" json:"hiddenColumns" yaml:"hiddenColumns"`
	// Indicates that lower and upper case letters are treated equally.
	IgnoreCase *string `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// Indicates if this column is participating in a split transform.
	IncludeInSplit *string `field:"optional" json:"includeInSplit" yaml:"includeInSplit"`
	// The input location to load the dataset from - Amazon S3 or AWS Glue Data Catalog .
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// The number of characters to split by.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Indicates if the content is text.
	IsText *string `field:"optional" json:"isText" yaml:"isText"`
	// The keys or columns involved in a join.
	JoinKeys *string `field:"optional" json:"joinKeys" yaml:"joinKeys"`
	// The type of join to use, for example, `INNER JOIN` , `OUTER JOIN` , and so on.
	JoinType *string `field:"optional" json:"joinType" yaml:"joinType"`
	// The columns on the left side of the join.
	LeftColumns *string `field:"optional" json:"leftColumns" yaml:"leftColumns"`
	// The number of times to perform `split` or `replaceBy` in a string.
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
	// The lower boundary for a value.
	LowerBound *string `field:"optional" json:"lowerBound" yaml:"lowerBound"`
	// The type of mappings to apply to construct a new dynamic frame.
	MapType *string `field:"optional" json:"mapType" yaml:"mapType"`
	// Determines the manner in which mode value is calculated, in case there is more than one mode value.
	//
	// Valid values: `NONE` | `AVERAGE` | `MINIMUM` | `MAXIMUM`.
	ModeType *string `field:"optional" json:"modeType" yaml:"modeType"`
	// Specifies whether JSON input contains embedded new line characters.
	MultiLine interface{} `field:"optional" json:"multiLine" yaml:"multiLine"`
	// The number of rows to consider in a window.
	NumRows *string `field:"optional" json:"numRows" yaml:"numRows"`
	// The number of rows to consider after the current row in a window.
	NumRowsAfter *string `field:"optional" json:"numRowsAfter" yaml:"numRowsAfter"`
	// The number of rows to consider before the current row in a window.
	NumRowsBefore *string `field:"optional" json:"numRowsBefore" yaml:"numRowsBefore"`
	// A column to sort the results by.
	OrderByColumn *string `field:"optional" json:"orderByColumn" yaml:"orderByColumn"`
	// The columns to sort the results by.
	OrderByColumns *string `field:"optional" json:"orderByColumns" yaml:"orderByColumns"`
	// The value to assign to unmapped cells, in categorical mapping.
	Other *string `field:"optional" json:"other" yaml:"other"`
	// The pattern to locate.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The starting pattern to split between.
	PatternOption1 *string `field:"optional" json:"patternOption1" yaml:"patternOption1"`
	// The ending pattern to split between.
	PatternOption2 *string `field:"optional" json:"patternOption2" yaml:"patternOption2"`
	// For splitting by multiple delimiters: A JSON-encoded string that lists the patterns in the format.
	//
	// For example: `[{\"pattern\":\"1\",\"includeInSplit\":true}]`.
	PatternOptions *string `field:"optional" json:"patternOptions" yaml:"patternOptions"`
	// The size of the rolling window.
	Period *string `field:"optional" json:"period" yaml:"period"`
	// The character index within a string.
	Position *string `field:"optional" json:"position" yaml:"position"`
	// If `true` , removes all of the following characters: `.` `.!` `.,` `.?`.
	RemoveAllPunctuation *string `field:"optional" json:"removeAllPunctuation" yaml:"removeAllPunctuation"`
	// If `true` , removes all single quotes and double quotes.
	RemoveAllQuotes *string `field:"optional" json:"removeAllQuotes" yaml:"removeAllQuotes"`
	// If `true` , removes all whitespaces from the value.
	RemoveAllWhitespace *string `field:"optional" json:"removeAllWhitespace" yaml:"removeAllWhitespace"`
	// If `true` , removes all chraracters specified by `CustomCharacters` .
	RemoveCustomCharacters *string `field:"optional" json:"removeCustomCharacters" yaml:"removeCustomCharacters"`
	// If `true` , removes all chraracters specified by `CustomValue` .
	RemoveCustomValue *string `field:"optional" json:"removeCustomValue" yaml:"removeCustomValue"`
	// If `true` , removes the following characters if they occur at the start or end of the value: `.` `!` `,` `?`.
	RemoveLeadingAndTrailingPunctuation *string `field:"optional" json:"removeLeadingAndTrailingPunctuation" yaml:"removeLeadingAndTrailingPunctuation"`
	// If `true` , removes single quotes and double quotes from the beginning and end of the value.
	RemoveLeadingAndTrailingQuotes *string `field:"optional" json:"removeLeadingAndTrailingQuotes" yaml:"removeLeadingAndTrailingQuotes"`
	// If `true` , removes all whitespaces from the beginning and end of the value.
	RemoveLeadingAndTrailingWhitespace *string `field:"optional" json:"removeLeadingAndTrailingWhitespace" yaml:"removeLeadingAndTrailingWhitespace"`
	// If `true` , removes all uppercase and lowercase alphabetic characters (A through Z;
	//
	// a through z).
	RemoveLetters *string `field:"optional" json:"removeLetters" yaml:"removeLetters"`
	// If `true` , removes all numeric characters (0 through 9).
	RemoveNumbers *string `field:"optional" json:"removeNumbers" yaml:"removeNumbers"`
	// If `true` , the source column will be removed after un-nesting that column.
	//
	// (Used with nested column types, such as Map, Struct, or Array.)
	RemoveSourceColumn *string `field:"optional" json:"removeSourceColumn" yaml:"removeSourceColumn"`
	// If `true` , removes all of the following characters: `!
	//
	// " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~`
	RemoveSpecialCharacters *string `field:"optional" json:"removeSpecialCharacters" yaml:"removeSpecialCharacters"`
	// The columns on the right side of a join.
	RightColumns *string `field:"optional" json:"rightColumns" yaml:"rightColumns"`
	// The number of rows in the sample.
	SampleSize *string `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// The sampling type to apply to the dataset.
	//
	// Valid values: `FIRST_N` | `LAST_N` | `RANDOM`.
	SampleType *string `field:"optional" json:"sampleType" yaml:"sampleType"`
	// A list of secondary inputs in a UNION transform.
	SecondaryInputs interface{} `field:"optional" json:"secondaryInputs" yaml:"secondaryInputs"`
	// A object value to indicate the second dataset used in a join.
	SecondInput *string `field:"optional" json:"secondInput" yaml:"secondInput"`
	// One or more sheet numbers in the Excel file, which will be included in a dataset.
	SheetIndexes interface{} `field:"optional" json:"sheetIndexes" yaml:"sheetIndexes"`
	// Oone or more named sheets in the Excel file, which will be included in a dataset.
	SheetNames *[]*string `field:"optional" json:"sheetNames" yaml:"sheetNames"`
	// A source column needed for an operation, step, or transform.
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// A source column needed for an operation, step, or transform.
	SourceColumn1 *string `field:"optional" json:"sourceColumn1" yaml:"sourceColumn1"`
	// A source column needed for an operation, step, or transform.
	SourceColumn2 *string `field:"optional" json:"sourceColumn2" yaml:"sourceColumn2"`
	// A list of source columns needed for an operation, step, or transform.
	SourceColumns *string `field:"optional" json:"sourceColumns" yaml:"sourceColumns"`
	// The index number of the first column used by an operation, step, or transform.
	StartColumnIndex *string `field:"optional" json:"startColumnIndex" yaml:"startColumnIndex"`
	// The starting pattern to locate.
	StartPattern *string `field:"optional" json:"startPattern" yaml:"startPattern"`
	// The starting position to locate.
	StartPosition *string `field:"optional" json:"startPosition" yaml:"startPosition"`
	// The starting value to locate.
	StartValue *string `field:"optional" json:"startValue" yaml:"startValue"`
	// Indicates this operation uses stems and lemmas (base words) for word tokenization.
	StemmingMode *string `field:"optional" json:"stemmingMode" yaml:"stemmingMode"`
	// The total number of transforms in this recipe.
	StepCount *string `field:"optional" json:"stepCount" yaml:"stepCount"`
	// The index ID of a step.
	StepIndex *string `field:"optional" json:"stepIndex" yaml:"stepIndex"`
	// Indicates this operation uses stop words as part of word tokenization.
	StopWordsMode *string `field:"optional" json:"stopWordsMode" yaml:"stopWordsMode"`
	// The resolution strategy to apply in resolving ambiguities.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// The column targeted by this operation.
	TargetColumn *string `field:"optional" json:"targetColumn" yaml:"targetColumn"`
	// The names to give columns altered by this operation.
	TargetColumnNames *string `field:"optional" json:"targetColumnNames" yaml:"targetColumnNames"`
	// The date format to convert to.
	TargetDateFormat *string `field:"optional" json:"targetDateFormat" yaml:"targetDateFormat"`
	// The index number of an object that is targeted by this operation.
	TargetIndex *string `field:"optional" json:"targetIndex" yaml:"targetIndex"`
	// The current timezone that you want to use for dates.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// A regex expression to use when splitting text into terms, also called words or tokens.
	TokenizerPattern *string `field:"optional" json:"tokenizerPattern" yaml:"tokenizerPattern"`
	// A value to use to represent `TRUE` .
	TrueString *string `field:"optional" json:"trueString" yaml:"trueString"`
	// The language that's used in the user-defined function.
	UdfLang *string `field:"optional" json:"udfLang" yaml:"udfLang"`
	// Specifies a unit of time.
	//
	// For example: `MINUTES` ; `SECONDS` ; `HOURS` ; etc.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// Cast columns as rows, so that each value is a different row in a single column.
	UnpivotColumn *string `field:"optional" json:"unpivotColumn" yaml:"unpivotColumn"`
	// The upper boundary for a value.
	UpperBound *string `field:"optional" json:"upperBound" yaml:"upperBound"`
	// Create a new container to hold a dataset.
	UseNewDataFrame *string `field:"optional" json:"useNewDataFrame" yaml:"useNewDataFrame"`
	// A static value that can be used in a comparison, a substitution, or in another context-specific way.
	//
	// A `Value` can be a number, string, or other datatype, depending on the recipe action in which it's used.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// A value that's used by this operation.
	Value1 *string `field:"optional" json:"value1" yaml:"value1"`
	// A value that's used by this operation.
	Value2 *string `field:"optional" json:"value2" yaml:"value2"`
	// The column that is provided as a value that's used by this operation.
	ValueColumn *string `field:"optional" json:"valueColumn" yaml:"valueColumn"`
	// The subset of rows currently available for viewing.
	ViewFrame *string `field:"optional" json:"viewFrame" yaml:"viewFrame"`
}

// Represents a single step from a DataBrew recipe to be performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recipeStepProperty := &recipeStepProperty{
//   	action: &actionProperty{
//   		operation: jsii.String("operation"),
//
//   		// the properties below are optional
//   		parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	conditionExpressions: []interface{}{
//   		&conditionExpressionProperty{
//   			condition: jsii.String("condition"),
//   			targetColumn: jsii.String("targetColumn"),
//
//   			// the properties below are optional
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRecipe_RecipeStepProperty struct {
	// The particular action to be performed in the recipe step.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// One or more conditions that must be met for the recipe step to succeed.
	//
	// > All of the conditions in the array must be met. In other words, all of the conditions must be combined using a logical AND operation.
	ConditionExpressions interface{} `field:"optional" json:"conditionExpressions" yaml:"conditionExpressions"`
}

// Represents an Amazon S3 location (bucket name, bucket owner, and object key) where DataBrew can read input data, or write output from a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	key: jsii.String("key"),
//   }
//
type CfnRecipe_S3LocationProperty struct {
	// The Amazon S3 bucket name.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The unique name of the object in the bucket.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

// Represents secondary inputs in a UNION transform.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secondaryInputProperty := &secondaryInputProperty{
//   	dataCatalogInputDefinition: &dataCatalogInputDefinitionProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		tableName: jsii.String("tableName"),
//   		tempDirectory: &s3LocationProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			key: jsii.String("key"),
//   		},
//   	},
//   	s3InputDefinition: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnRecipe_SecondaryInputProperty struct {
	// The AWS Glue Data Catalog parameters for the data.
	DataCatalogInputDefinition interface{} `field:"optional" json:"dataCatalogInputDefinition" yaml:"dataCatalogInputDefinition"`
	// The Amazon S3 location where the data is stored.
	S3InputDefinition interface{} `field:"optional" json:"s3InputDefinition" yaml:"s3InputDefinition"`
}

// Properties for defining a `CfnRecipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecipeProps := &cfnRecipeProps{
//   	name: jsii.String("name"),
//   	steps: []interface{}{
//   		&recipeStepProperty{
//   			action: &actionProperty{
//   				operation: jsii.String("operation"),
//
//   				// the properties below are optional
//   				parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			conditionExpressions: []interface{}{
//   				&conditionExpressionProperty{
//   					condition: jsii.String("condition"),
//   					targetColumn: jsii.String("targetColumn"),
//
//   					// the properties below are optional
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRecipeProps struct {
	// The unique name for the recipe.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of steps that are defined by the recipe.
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// The description of the recipe.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata tags that have been applied to the recipe.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataBrew::Ruleset`.
//
// Specifies a new ruleset that can be used in a profile job to validate the data quality of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRuleset := awscdk.Aws_databrew.NewCfnRuleset(this, jsii.String("MyCfnRuleset"), &cfnRulesetProps{
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			checkExpression: jsii.String("checkExpression"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			columnSelectors: []interface{}{
//   				&columnSelectorProperty{
//   					name: jsii.String("name"),
//   					regex: jsii.String("regex"),
//   				},
//   			},
//   			disabled: jsii.Boolean(false),
//   			substitutionMap: []interface{}{
//   				&substitutionValueProperty{
//   					value: jsii.String("value"),
//   					valueReference: jsii.String("valueReference"),
//   				},
//   			},
//   			threshold: &thresholdProperty{
//   				value: jsii.Number(123),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   				unit: jsii.String("unit"),
//   			},
//   		},
//   	},
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnRuleset interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the ruleset.
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
	// The name of the ruleset.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Contains metadata about the ruleset.
	Rules() interface{}
	SetRules(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// The Amazon Resource Name (ARN) of a resource (dataset) that the ruleset is associated with.
	TargetArn() *string
	SetTargetArn(val *string)
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

// The jsii proxy struct for CfnRuleset
type jsiiProxy_CfnRuleset struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRuleset) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleset) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Ruleset`.
func NewCfnRuleset(scope constructs.Construct, id *string, props *CfnRulesetProps) CfnRuleset {
	_init_.Initialize()

	j := jsiiProxy_CfnRuleset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Ruleset`.
func NewCfnRuleset_Override(c CfnRuleset, scope constructs.Construct, id *string, props *CfnRulesetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRuleset) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRuleset) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRuleset) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_CfnRuleset) SetTargetArn(val *string) {
	_jsii_.Set(
		j,
		"targetArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRuleset_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRuleset_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
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
func CfnRuleset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRuleset_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnRuleset",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRuleset) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRuleset) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRuleset) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRuleset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRuleset) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRuleset) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRuleset) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRuleset) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleset) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleset) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRuleset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRuleset) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleset) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleset) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Selector of a column from a dataset for profile job configuration.
//
// One selector includes either a column name or a regular expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnSelectorProperty := &columnSelectorProperty{
//   	name: jsii.String("name"),
//   	regex: jsii.String("regex"),
//   }
//
type CfnRuleset_ColumnSelectorProperty struct {
	// The name of a column from a dataset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A regular expression for selecting a column from a dataset.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

// Represents a single data quality requirement that should be validated in the scope of this dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &ruleProperty{
//   	checkExpression: jsii.String("checkExpression"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	columnSelectors: []interface{}{
//   		&columnSelectorProperty{
//   			name: jsii.String("name"),
//   			regex: jsii.String("regex"),
//   		},
//   	},
//   	disabled: jsii.Boolean(false),
//   	substitutionMap: []interface{}{
//   		&substitutionValueProperty{
//   			value: jsii.String("value"),
//   			valueReference: jsii.String("valueReference"),
//   		},
//   	},
//   	threshold: &thresholdProperty{
//   		value: jsii.Number(123),
//
//   		// the properties below are optional
//   		type: jsii.String("type"),
//   		unit: jsii.String("unit"),
//   	},
//   }
//
type CfnRuleset_RuleProperty struct {
	// The expression which includes column references, condition names followed by variable references, possibly grouped and combined with other conditions.
	//
	// For example, `(:col1 starts_with :prefix1 or :col1 starts_with :prefix2) and (:col1 ends_with :suffix1 or :col1 ends_with :suffix2)` . Column and value references are substitution variables that should start with the ':' symbol. Depending on the context, substitution variables' values can be either an actual value or a column name. These values are defined in the SubstitutionMap. If a CheckExpression starts with a column reference, then ColumnSelectors in the rule should be null. If ColumnSelectors has been defined, then there should be no columnn reference in the left side of a condition, for example, `is_between :val1 and :val2` .
	CheckExpression *string `field:"required" json:"checkExpression" yaml:"checkExpression"`
	// The name of the rule.
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of column selectors.
	//
	// Selectors can be used to select columns using a name or regular expression from the dataset. Rule will be applied to selected columns.
	ColumnSelectors interface{} `field:"optional" json:"columnSelectors" yaml:"columnSelectors"`
	// A value that specifies whether the rule is disabled.
	//
	// Once a rule is disabled, a profile job will not validate it during a job run. Default value is false.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The map of substitution variable names to their values used in a check expression.
	//
	// Variable names should start with a ':' (colon). Variable values can either be actual values or column names. To differentiate between the two, column names should be enclosed in backticks, for example, `":col1": "`Column A`".`
	SubstitutionMap interface{} `field:"optional" json:"substitutionMap" yaml:"substitutionMap"`
	// The threshold used with a non-aggregate check expression.
	//
	// Non-aggregate check expressions will be applied to each row in a specific column, and the threshold will be used to determine whether the validation succeeds.
	Threshold interface{} `field:"optional" json:"threshold" yaml:"threshold"`
}

// A key-value pair to associate an expression's substitution variable names with their values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   substitutionValueProperty := &substitutionValueProperty{
//   	value: jsii.String("value"),
//   	valueReference: jsii.String("valueReference"),
//   }
//
type CfnRuleset_SubstitutionValueProperty struct {
	// Value or column name.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Variable name.
	ValueReference *string `field:"required" json:"valueReference" yaml:"valueReference"`
}

// The threshold used with a non-aggregate check expression.
//
// The non-aggregate check expression will be applied to each row in a specific column. Then the threshold will be used to determine whether the validation succeeds.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thresholdProperty := &thresholdProperty{
//   	value: jsii.Number(123),
//
//   	// the properties below are optional
//   	type: jsii.String("type"),
//   	unit: jsii.String("unit"),
//   }
//
type CfnRuleset_ThresholdProperty struct {
	// The value of a threshold.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// The type of a threshold.
	//
	// Used for comparison of an actual count of rows that satisfy the rule to the threshold value.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Unit of threshold value.
	//
	// Can be either a COUNT or PERCENTAGE of the full sample size used for validation.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

// Properties for defining a `CfnRuleset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRulesetProps := &cfnRulesetProps{
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			checkExpression: jsii.String("checkExpression"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			columnSelectors: []interface{}{
//   				&columnSelectorProperty{
//   					name: jsii.String("name"),
//   					regex: jsii.String("regex"),
//   				},
//   			},
//   			disabled: jsii.Boolean(false),
//   			substitutionMap: []interface{}{
//   				&substitutionValueProperty{
//   					value: jsii.String("value"),
//   					valueReference: jsii.String("valueReference"),
//   				},
//   			},
//   			threshold: &thresholdProperty{
//   				value: jsii.Number(123),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   				unit: jsii.String("unit"),
//   			},
//   		},
//   	},
//   	targetArn: jsii.String("targetArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRulesetProps struct {
	// The name of the ruleset.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains metadata about the ruleset.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The Amazon Resource Name (ARN) of a resource (dataset) that the ruleset is associated with.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// The description of the ruleset.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DataBrew::Schedule`.
//
// Specifies a new schedule for one or more AWS Glue DataBrew jobs. Jobs can be run at a specific date and time, or at regular intervals.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchedule := awscdk.Aws_databrew.NewCfnSchedule(this, jsii.String("MyCfnSchedule"), &cfnScheduleProps{
//   	cronExpression: jsii.String("cronExpression"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	jobNames: []*string{
//   		jsii.String("jobNames"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnSchedule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The dates and times when the job is to run.
	//
	// For more information, see [Working with cron expressions for recipe jobs](https://docs.aws.amazon.com/databrew/latest/dg/jobs.recipe.html#jobs.cron) in the *AWS Glue DataBrew Developer Guide* .
	CronExpression() *string
	SetCronExpression(val *string)
	// A list of jobs to be run, according to the schedule.
	JobNames() *[]*string
	SetJobNames(val *[]*string)
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
	// The name of the schedule.
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
	// Metadata tags that have been applied to the schedule.
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

// The jsii proxy struct for CfnSchedule
type jsiiProxy_CfnSchedule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSchedule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) CronExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) JobNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jobNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedule) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Schedule`.
func NewCfnSchedule(scope constructs.Construct, id *string, props *CfnScheduleProps) CfnSchedule {
	_init_.Initialize()

	j := jsiiProxy_CfnSchedule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Schedule`.
func NewCfnSchedule_Override(c CfnSchedule, scope constructs.Construct, id *string, props *CfnScheduleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSchedule) SetCronExpression(val *string) {
	_jsii_.Set(
		j,
		"cronExpression",
		val,
	)
}

func (j *jsiiProxy_CfnSchedule) SetJobNames(val *[]*string) {
	_jsii_.Set(
		j,
		"jobNames",
		val,
	)
}

func (j *jsiiProxy_CfnSchedule) SetName(val *string) {
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
func CfnSchedule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSchedule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
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
func CfnSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchedule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnSchedule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSchedule) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSchedule) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSchedule) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSchedule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSchedule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSchedule) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSchedule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSchedule) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedule) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSchedule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSchedule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchedule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduleProps := &cfnScheduleProps{
//   	cronExpression: jsii.String("cronExpression"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	jobNames: []*string{
//   		jsii.String("jobNames"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnScheduleProps struct {
	// The dates and times when the job is to run.
	//
	// For more information, see [Working with cron expressions for recipe jobs](https://docs.aws.amazon.com/databrew/latest/dg/jobs.recipe.html#jobs.cron) in the *AWS Glue DataBrew Developer Guide* .
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// The name of the schedule.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of jobs to be run, according to the schedule.
	JobNames *[]*string `field:"optional" json:"jobNames" yaml:"jobNames"`
	// Metadata tags that have been applied to the schedule.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

