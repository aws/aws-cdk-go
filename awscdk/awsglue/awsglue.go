package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Glue::Classifier`.
//
// The `AWS::Glue::Classifier` resource creates an AWS Glue classifier that categorizes data sources and specifies schemas. For more information, see [Adding Classifiers to a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-classifier.html) and [Classifier Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-crawler-classifiers.html#aws-glue-api-crawler-classifiers-Classifier) in the *AWS Glue Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnClassifier interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	CsvClassifier() interface{}
	SetCsvClassifier(val interface{})
	GrokClassifier() interface{}
	SetGrokClassifier(val interface{})
	JsonClassifier() interface{}
	SetJsonClassifier(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	XmlClassifier() interface{}
	SetXmlClassifier(val interface{})
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnClassifier
type jsiiProxy_CfnClassifier struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClassifier) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) CsvClassifier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"csvClassifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) GrokClassifier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grokClassifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) JsonClassifier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonClassifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClassifier) XmlClassifier() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xmlClassifier",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Classifier`.
func NewCfnClassifier(scope constructs.Construct, id *string, props *CfnClassifierProps) CfnClassifier {
	_init_.Initialize()

	j := jsiiProxy_CfnClassifier{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnClassifier",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Classifier`.
func NewCfnClassifier_Override(c CfnClassifier, scope constructs.Construct, id *string, props *CfnClassifierProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnClassifier",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClassifier) SetCsvClassifier(val interface{}) {
	_jsii_.Set(
		j,
		"csvClassifier",
		val,
	)
}

func (j *jsiiProxy_CfnClassifier) SetGrokClassifier(val interface{}) {
	_jsii_.Set(
		j,
		"grokClassifier",
		val,
	)
}

func (j *jsiiProxy_CfnClassifier) SetJsonClassifier(val interface{}) {
	_jsii_.Set(
		j,
		"jsonClassifier",
		val,
	)
}

func (j *jsiiProxy_CfnClassifier) SetXmlClassifier(val interface{}) {
	_jsii_.Set(
		j,
		"xmlClassifier",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnClassifier_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnClassifier",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnClassifier_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnClassifier",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnClassifier_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnClassifier",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClassifier_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnClassifier",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnClassifier) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnClassifier) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnClassifier) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnClassifier) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnClassifier) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnClassifier) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnClassifier) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnClassifier) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnClassifier) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnClassifier) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnClassifier) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClassifier) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnClassifier) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnClassifier) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClassifier) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A classifier for custom `CSV` content.
//
// TODO: EXAMPLE
//
type CfnClassifier_CsvClassifierProperty struct {
	// Enables the processing of files that contain only one column.
	AllowSingleColumn interface{} `json:"allowSingleColumn" yaml:"allowSingleColumn"`
	// Indicates whether the CSV file contains a header.
	ContainsHeader *string `json:"containsHeader" yaml:"containsHeader"`
	// A custom symbol to denote what separates each column entry in the row.
	Delimiter *string `json:"delimiter" yaml:"delimiter"`
	// Specifies not to trim values before identifying the type of column values.
	//
	// The default value is `true` .
	DisableValueTrimming interface{} `json:"disableValueTrimming" yaml:"disableValueTrimming"`
	// A list of strings representing column names.
	Header *[]*string `json:"header" yaml:"header"`
	// The name of the classifier.
	Name *string `json:"name" yaml:"name"`
	// A custom symbol to denote what combines content into a single column value.
	//
	// It must be different from the column delimiter.
	QuoteSymbol *string `json:"quoteSymbol" yaml:"quoteSymbol"`
}

// A classifier that uses `grok` patterns.
//
// TODO: EXAMPLE
//
type CfnClassifier_GrokClassifierProperty struct {
	// An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, and so on.
	Classification *string `json:"classification" yaml:"classification"`
	// The grok pattern applied to a data store by this classifier.
	//
	// For more information, see built-in patterns in [Writing Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html) .
	GrokPattern *string `json:"grokPattern" yaml:"grokPattern"`
	// Optional custom grok patterns defined by this classifier.
	//
	// For more information, see custom patterns in [Writing Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html) .
	CustomPatterns *string `json:"customPatterns" yaml:"customPatterns"`
	// The name of the classifier.
	Name *string `json:"name" yaml:"name"`
}

// A classifier for `JSON` content.
//
// TODO: EXAMPLE
//
type CfnClassifier_JsonClassifierProperty struct {
	// A `JsonPath` string defining the JSON data for the classifier to classify.
	//
	// AWS Glue supports a subset of `JsonPath` , as described in [Writing JsonPath Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html#custom-classifier-json) .
	JsonPath *string `json:"jsonPath" yaml:"jsonPath"`
	// The name of the classifier.
	Name *string `json:"name" yaml:"name"`
}

// A classifier for `XML` content.
//
// TODO: EXAMPLE
//
type CfnClassifier_XMLClassifierProperty struct {
	// An identifier of the data format that the classifier matches.
	Classification *string `json:"classification" yaml:"classification"`
	// The XML tag designating the element that contains each record in an XML document being parsed.
	//
	// This can't identify a self-closing element (closed by `/>` ). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, `<row item_a="A" item_b="B"></row>` is okay, but `<row item_a="A" item_b="B" />` is not).
	RowTag *string `json:"rowTag" yaml:"rowTag"`
	// The name of the classifier.
	Name *string `json:"name" yaml:"name"`
}

// Properties for defining a `CfnClassifier`.
//
// TODO: EXAMPLE
//
type CfnClassifierProps struct {
	// A classifier for comma-separated values (CSV).
	CsvClassifier interface{} `json:"csvClassifier" yaml:"csvClassifier"`
	// A classifier that uses `grok` .
	GrokClassifier interface{} `json:"grokClassifier" yaml:"grokClassifier"`
	// A classifier for JSON content.
	JsonClassifier interface{} `json:"jsonClassifier" yaml:"jsonClassifier"`
	// A classifier for XML content.
	XmlClassifier interface{} `json:"xmlClassifier" yaml:"xmlClassifier"`
}

// A CloudFormation `AWS::Glue::Connection`.
//
// The `AWS::Glue::Connection` resource specifies an AWS Glue connection to a data source. For more information, see [Adding a Connection to Your Data Store](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html) and [Connection Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-connections.html#aws-glue-api-catalog-connections-Connection) in the *AWS Glue Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnConnection interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CatalogId() *string
	SetCatalogId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConnectionInput() interface{}
	SetConnectionInput(val interface{})
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConnection
type jsiiProxy_CfnConnection struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnection) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) ConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Connection`.
func NewCfnConnection(scope constructs.Construct, id *string, props *CfnConnectionProps) CfnConnection {
	_init_.Initialize()

	j := jsiiProxy_CfnConnection{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnConnection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Connection`.
func NewCfnConnection_Override(c CfnConnection, scope constructs.Construct, id *string, props *CfnConnectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnConnection",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnection) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_CfnConnection) SetConnectionInput(val interface{}) {
	_jsii_.Set(
		j,
		"connectionInput",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnection_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnConnection",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConnection_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnConnection",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnection_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnConnection",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnConnection) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnConnection) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnConnection) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnConnection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnConnection) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnConnection) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnConnection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnConnection) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnConnection) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnConnection) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnConnection) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnection) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnConnection) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnection) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A structure that is used to specify a connection to create or update.
//
// TODO: EXAMPLE
//
type CfnConnection_ConnectionInputProperty struct {
	// The type of the connection. Currently, these types are supported:.
	//
	// - `JDBC` - Designates a connection to a database through Java Database Connectivity (JDBC).
	// - `KAFKA` - Designates a connection to an Apache Kafka streaming platform.
	// - `MONGODB` - Designates a connection to a MongoDB document database.
	// - `NETWORK` - Designates a network connection to a data source within an Amazon Virtual Private Cloud environment (Amazon VPC).
	//
	// SFTP is not supported.
	ConnectionType *string `json:"connectionType" yaml:"connectionType"`
	// These key-value pairs define parameters for the connection.
	ConnectionProperties interface{} `json:"connectionProperties" yaml:"connectionProperties"`
	// The description of the connection.
	Description *string `json:"description" yaml:"description"`
	// A list of criteria that can be used in selecting this connection.
	MatchCriteria *[]*string `json:"matchCriteria" yaml:"matchCriteria"`
	// The name of the connection.
	Name *string `json:"name" yaml:"name"`
	// A map of physical connection requirements, such as virtual private cloud (VPC) and `SecurityGroup` , that are needed to successfully make this connection.
	PhysicalConnectionRequirements interface{} `json:"physicalConnectionRequirements" yaml:"physicalConnectionRequirements"`
}

// Specifies the physical requirements for a connection.
//
// TODO: EXAMPLE
//
type CfnConnection_PhysicalConnectionRequirementsProperty struct {
	// The connection's Availability Zone.
	//
	// This field is redundant because the specified subnet implies the Availability Zone to be used. Currently the field must be populated, but it will be deprecated in the future.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The security group ID list used by the connection.
	SecurityGroupIdList *[]*string `json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// The subnet ID used by the connection.
	SubnetId *string `json:"subnetId" yaml:"subnetId"`
}

// Properties for defining a `CfnConnection`.
//
// TODO: EXAMPLE
//
type CfnConnectionProps struct {
	// The ID of the data catalog to create the catalog object in.
	//
	// Currently, this should be the AWS account ID.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId` .
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The connection that you want to create.
	ConnectionInput interface{} `json:"connectionInput" yaml:"connectionInput"`
}

// A CloudFormation `AWS::Glue::Crawler`.
//
// The `AWS::Glue::Crawler` resource specifies an AWS Glue crawler. For more information, see [Cataloging Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) and [Crawler Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-crawler-crawling.html#aws-glue-api-crawler-crawling-Crawler) in the *AWS Glue Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnCrawler interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Classifiers() *[]*string
	SetClassifiers(val *[]*string)
	Configuration() *string
	SetConfiguration(val *string)
	CrawlerSecurityConfiguration() *string
	SetCrawlerSecurityConfiguration(val *string)
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	RecrawlPolicy() interface{}
	SetRecrawlPolicy(val interface{})
	Ref() *string
	Role() *string
	SetRole(val *string)
	Schedule() interface{}
	SetSchedule(val interface{})
	SchemaChangePolicy() interface{}
	SetSchemaChangePolicy(val interface{})
	Stack() awscdk.Stack
	TablePrefix() *string
	SetTablePrefix(val *string)
	Tags() awscdk.TagManager
	Targets() interface{}
	SetTargets(val interface{})
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnCrawler
type jsiiProxy_CfnCrawler struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCrawler) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Classifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) CrawlerSecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crawlerSecurityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) RecrawlPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recrawlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) SchemaChangePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaChangePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCrawler) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Crawler`.
func NewCfnCrawler(scope constructs.Construct, id *string, props *CfnCrawlerProps) CfnCrawler {
	_init_.Initialize()

	j := jsiiProxy_CfnCrawler{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Crawler`.
func NewCfnCrawler_Override(c CfnCrawler, scope constructs.Construct, id *string, props *CfnCrawlerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCrawler) SetClassifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"classifiers",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetConfiguration(val *string) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetCrawlerSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"crawlerSecurityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetRecrawlPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"recrawlPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetSchemaChangePolicy(val interface{}) {
	_jsii_.Set(
		j,
		"schemaChangePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetTablePrefix(val *string) {
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_CfnCrawler) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCrawler_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCrawler_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnCrawler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCrawler_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnCrawler",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnCrawler) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnCrawler) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnCrawler) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnCrawler) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnCrawler) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnCrawler) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnCrawler) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnCrawler) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnCrawler) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnCrawler) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCrawler) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCrawler) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnCrawler) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnCrawler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCrawler) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an AWS Glue Data Catalog target.
//
// TODO: EXAMPLE
//
type CfnCrawler_CatalogTargetProperty struct {
	// The name of the database to be synchronized.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// A list of the tables to be synchronized.
	Tables *[]*string `json:"tables" yaml:"tables"`
}

// Specifies an Amazon DynamoDB table to crawl.
//
// TODO: EXAMPLE
//
type CfnCrawler_DynamoDBTargetProperty struct {
	// The name of the DynamoDB table to crawl.
	Path *string `json:"path" yaml:"path"`
}

// Specifies a JDBC data store to crawl.
//
// TODO: EXAMPLE
//
type CfnCrawler_JdbcTargetProperty struct {
	// The name of the connection to use to connect to the JDBC target.
	ConnectionName *string `json:"connectionName" yaml:"connectionName"`
	// A list of glob patterns used to exclude from the crawl.
	//
	// For more information, see [Catalog Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) .
	Exclusions *[]*string `json:"exclusions" yaml:"exclusions"`
	// The path of the JDBC target.
	Path *string `json:"path" yaml:"path"`
}

// When crawling an Amazon S3 data source after the first crawl is complete, specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run.
//
// For more information, see [Incremental Crawls in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/incremental-crawls.html) in the developer guide.
//
// TODO: EXAMPLE
//
type CfnCrawler_RecrawlPolicyProperty struct {
	// Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run.
	//
	// A value of `CRAWL_EVERYTHING` specifies crawling the entire dataset again.
	//
	// A value of `CRAWL_NEW_FOLDERS_ONLY` specifies crawling only folders that were added since the last crawler run.
	//
	// A value of `CRAWL_EVENT_MODE` specifies crawling only the changes identified by Amazon S3 events.
	RecrawlBehavior *string `json:"recrawlBehavior" yaml:"recrawlBehavior"`
}

// Specifies a data store in Amazon Simple Storage Service (Amazon S3).
//
// TODO: EXAMPLE
//
type CfnCrawler_S3TargetProperty struct {
	// The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC).
	ConnectionName *string `json:"connectionName" yaml:"connectionName"`
	// A list of glob patterns used to exclude from the crawl.
	//
	// For more information, see [Catalog Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) .
	Exclusions *[]*string `json:"exclusions" yaml:"exclusions"`
	// The path to the Amazon S3 target.
	Path *string `json:"path" yaml:"path"`
}

// A scheduling object using a `cron` statement to schedule an event.
//
// TODO: EXAMPLE
//
type CfnCrawler_ScheduleProperty struct {
	// A `cron` expression used to specify the schedule.
	//
	// For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html) . For example, to run something every day at 12:15 UTC, specify `cron(15 12 * * ? *)` .
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
}

// A policy that specifies update and deletion behaviors for the crawler.
//
// TODO: EXAMPLE
//
type CfnCrawler_SchemaChangePolicyProperty struct {
	// The deletion behavior when the crawler finds a deleted object.
	DeleteBehavior *string `json:"deleteBehavior" yaml:"deleteBehavior"`
	// The update behavior when the crawler finds a changed schema.
	UpdateBehavior *string `json:"updateBehavior" yaml:"updateBehavior"`
}

// Specifies data stores to crawl.
//
// TODO: EXAMPLE
//
type CfnCrawler_TargetsProperty struct {
	// Specifies AWS Glue Data Catalog targets.
	CatalogTargets interface{} `json:"catalogTargets" yaml:"catalogTargets"`
	// Specifies Amazon DynamoDB targets.
	DynamoDbTargets interface{} `json:"dynamoDbTargets" yaml:"dynamoDbTargets"`
	// Specifies JDBC targets.
	JdbcTargets interface{} `json:"jdbcTargets" yaml:"jdbcTargets"`
	// Specifies Amazon Simple Storage Service (Amazon S3) targets.
	S3Targets interface{} `json:"s3Targets" yaml:"s3Targets"`
}

// Properties for defining a `CfnCrawler`.
//
// TODO: EXAMPLE
//
type CfnCrawlerProps struct {
	// The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.
	Role *string `json:"role" yaml:"role"`
	// A collection of targets to crawl.
	Targets interface{} `json:"targets" yaml:"targets"`
	// A list of UTF-8 strings that specify the custom classifiers that are associated with the crawler.
	Classifiers *[]*string `json:"classifiers" yaml:"classifiers"`
	// Crawler configuration information.
	//
	// This versioned JSON string allows users to specify aspects of a crawler's behavior. For more information, see [Configuring a Crawler](https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html) .
	Configuration *string `json:"configuration" yaml:"configuration"`
	// The name of the `SecurityConfiguration` structure to be used by this crawler.
	CrawlerSecurityConfiguration *string `json:"crawlerSecurityConfiguration" yaml:"crawlerSecurityConfiguration"`
	// The name of the database in which the crawler's output is stored.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// A description of the crawler.
	Description *string `json:"description" yaml:"description"`
	// The name of the crawler.
	Name *string `json:"name" yaml:"name"`
	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.
	RecrawlPolicy interface{} `json:"recrawlPolicy" yaml:"recrawlPolicy"`
	// For scheduled crawlers, the schedule when the crawler runs.
	Schedule interface{} `json:"schedule" yaml:"schedule"`
	// The policy that specifies update and delete behaviors for the crawler.
	SchemaChangePolicy interface{} `json:"schemaChangePolicy" yaml:"schemaChangePolicy"`
	// The prefix added to the names of tables that are created.
	TablePrefix *string `json:"tablePrefix" yaml:"tablePrefix"`
	// The tags to use with this crawler.
	Tags interface{} `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Glue::DataCatalogEncryptionSettings`.
//
// Sets the security configuration for a specified catalog. After the configuration has been set, the specified encryption is applied to every catalog write thereafter.
//
// TODO: EXAMPLE
//
type CfnDataCatalogEncryptionSettings interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CatalogId() *string
	SetCatalogId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DataCatalogEncryptionSettings() interface{}
	SetDataCatalogEncryptionSettings(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDataCatalogEncryptionSettings
type jsiiProxy_CfnDataCatalogEncryptionSettings struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) DataCatalogEncryptionSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCatalogEncryptionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::DataCatalogEncryptionSettings`.
func NewCfnDataCatalogEncryptionSettings(scope constructs.Construct, id *string, props *CfnDataCatalogEncryptionSettingsProps) CfnDataCatalogEncryptionSettings {
	_init_.Initialize()

	j := jsiiProxy_CfnDataCatalogEncryptionSettings{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnDataCatalogEncryptionSettings",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::DataCatalogEncryptionSettings`.
func NewCfnDataCatalogEncryptionSettings_Override(c CfnDataCatalogEncryptionSettings, scope constructs.Construct, id *string, props *CfnDataCatalogEncryptionSettingsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnDataCatalogEncryptionSettings",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_CfnDataCatalogEncryptionSettings) SetDataCatalogEncryptionSettings(val interface{}) {
	_jsii_.Set(
		j,
		"dataCatalogEncryptionSettings",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataCatalogEncryptionSettings_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDataCatalogEncryptionSettings",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataCatalogEncryptionSettings_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDataCatalogEncryptionSettings",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDataCatalogEncryptionSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDataCatalogEncryptionSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataCatalogEncryptionSettings_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnDataCatalogEncryptionSettings",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataCatalogEncryptionSettings) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The data structure used by the Data Catalog to encrypt the password as part of `CreateConnection` or `UpdateConnection` and store it in the `ENCRYPTED_PASSWORD` field in the connection properties.
//
// You can enable catalog encryption or only password encryption.
//
// When a `CreationConnection` request arrives containing a password, the Data Catalog first encrypts the password using your AWS KMS key. It then encrypts the whole connection object again if catalog encryption is also enabled.
//
// This encryption requires that you set AWS KMS key permissions to enable or restrict access on the password key according to your security requirements. For example, you might want only administrators to have decrypt permission on the password key.
//
// TODO: EXAMPLE
//
type CfnDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty struct {
	// An AWS KMS key that is used to encrypt the connection password.
	//
	// If connection password protection is enabled, the caller of `CreateConnection` and `UpdateConnection` needs at least `kms:Encrypt` permission on the specified AWS KMS key, to encrypt passwords before storing them in the Data Catalog. You can set the decrypt permission to enable or restrict access on the password key according to your security requirements.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// When the `ReturnConnectionPasswordEncrypted` flag is set to "true", passwords remain encrypted in the responses of `GetConnection` and `GetConnections` .
	//
	// This encryption takes effect independently from catalog encryption.
	ReturnConnectionPasswordEncrypted interface{} `json:"returnConnectionPasswordEncrypted" yaml:"returnConnectionPasswordEncrypted"`
}

// Contains configuration information for maintaining Data Catalog security.
//
// TODO: EXAMPLE
//
type CfnDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty struct {
	// When connection password protection is enabled, the Data Catalog uses a customer-provided key to encrypt the password as part of `CreateConnection` or `UpdateConnection` and store it in the `ENCRYPTED_PASSWORD` field in the connection properties.
	//
	// You can enable catalog encryption or only password encryption.
	ConnectionPasswordEncryption interface{} `json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// Specifies the encryption-at-rest configuration for the Data Catalog.
	EncryptionAtRest interface{} `json:"encryptionAtRest" yaml:"encryptionAtRest"`
}

// Specifies the encryption-at-rest configuration for the Data Catalog.
//
// TODO: EXAMPLE
//
type CfnDataCatalogEncryptionSettings_EncryptionAtRestProperty struct {
	// The encryption-at-rest mode for encrypting Data Catalog data.
	CatalogEncryptionMode *string `json:"catalogEncryptionMode" yaml:"catalogEncryptionMode"`
	// The ID of the AWS KMS key to use for encryption at rest.
	SseAwsKmsKeyId *string `json:"sseAwsKmsKeyId" yaml:"sseAwsKmsKeyId"`
}

// Properties for defining a `CfnDataCatalogEncryptionSettings`.
//
// TODO: EXAMPLE
//
type CfnDataCatalogEncryptionSettingsProps struct {
	// The ID of the Data Catalog in which the settings are created.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Contains configuration information for maintaining Data Catalog security.
	DataCatalogEncryptionSettings interface{} `json:"dataCatalogEncryptionSettings" yaml:"dataCatalogEncryptionSettings"`
}

// A CloudFormation `AWS::Glue::Database`.
//
// The `AWS::Glue::Database` resource specifies a logical grouping of tables in AWS Glue . For more information, see [Defining a Database in Your Data Catalog](https://docs.aws.amazon.com/glue/latest/dg/define-database.html) and [Database Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-databases.html#aws-glue-api-catalog-databases-Database) in the *AWS Glue Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnDatabase interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CatalogId() *string
	SetCatalogId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatabaseInput() interface{}
	SetDatabaseInput(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDatabase
type jsiiProxy_CfnDatabase struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDatabase) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Database`.
func NewCfnDatabase(scope constructs.Construct, id *string, props *CfnDatabaseProps) CfnDatabase {
	_init_.Initialize()

	j := jsiiProxy_CfnDatabase{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnDatabase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Database`.
func NewCfnDatabase_Override(c CfnDatabase, scope constructs.Construct, id *string, props *CfnDatabaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnDatabase",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDatabase) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetDatabaseInput(val interface{}) {
	_jsii_.Set(
		j,
		"databaseInput",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDatabase_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDatabase",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDatabase_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDatabase",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatabase_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnDatabase",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDatabase) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDatabase) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDatabase) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnDatabase) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDatabase) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDatabase) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDatabase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDatabase) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDatabase) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDatabase) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDatabase) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDatabase) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDatabase) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatabase) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The AWS Lake Formation principal.
//
// TODO: EXAMPLE
//
type CfnDatabase_DataLakePrincipalProperty struct {
	// An identifier for the AWS Lake Formation principal.
	DataLakePrincipalIdentifier *string `json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

// A structure that describes a target database for resource linking.
//
// TODO: EXAMPLE
//
type CfnDatabase_DatabaseIdentifierProperty struct {
	// The ID of the Data Catalog in which the database resides.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
}

// The structure used to create or update a database.
//
// TODO: EXAMPLE
//
type CfnDatabase_DatabaseInputProperty struct {
	// Creates a set of default permissions on the table for principals.
	CreateTableDefaultPermissions interface{} `json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// A description of the database.
	Description *string `json:"description" yaml:"description"`
	// The location of the database (for example, an HDFS path).
	LocationUri *string `json:"locationUri" yaml:"locationUri"`
	// The name of the database.
	//
	// For Hive compatibility, this is folded to lowercase when it is stored.
	Name *string `json:"name" yaml:"name"`
	// These key-value pairs define parameters and properties of the database.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// A `DatabaseIdentifier` structure that describes a target database for resource linking.
	TargetDatabase interface{} `json:"targetDatabase" yaml:"targetDatabase"`
}

// TODO: EXAMPLE
//
type CfnDatabase_PrincipalPrivilegesProperty struct {
	// `CfnDatabase.PrincipalPrivilegesProperty.Permissions`.
	Permissions *[]*string `json:"permissions" yaml:"permissions"`
	// `CfnDatabase.PrincipalPrivilegesProperty.Principal`.
	Principal interface{} `json:"principal" yaml:"principal"`
}

// Properties for defining a `CfnDatabase`.
//
// TODO: EXAMPLE
//
type CfnDatabaseProps struct {
	// The AWS account ID for the account in which to create the catalog object.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId`
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The metadata for the database.
	DatabaseInput interface{} `json:"databaseInput" yaml:"databaseInput"`
}

// A CloudFormation `AWS::Glue::DevEndpoint`.
//
// The `AWS::Glue::DevEndpoint` resource specifies a development endpoint where a developer can remotely debug ETL scripts for AWS Glue . For more information, see [DevEndpoint Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-dev-endpoint.html#aws-glue-api-jobs-dev-endpoint-DevEndpoint) in the AWS Glue Developer Guide.
//
// TODO: EXAMPLE
//
type CfnDevEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Arguments() interface{}
	SetArguments(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EndpointName() *string
	SetEndpointName(val *string)
	ExtraJarsS3Path() *string
	SetExtraJarsS3Path(val *string)
	ExtraPythonLibsS3Path() *string
	SetExtraPythonLibsS3Path(val *string)
	GlueVersion() *string
	SetGlueVersion(val *string)
	LogicalId() *string
	Node() constructs.Node
	NumberOfNodes() *float64
	SetNumberOfNodes(val *float64)
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	PublicKey() *string
	SetPublicKey(val *string)
	PublicKeys() *[]*string
	SetPublicKeys(val *[]*string)
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	Stack() awscdk.Stack
	SubnetId() *string
	SetSubnetId(val *string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	WorkerType() *string
	SetWorkerType(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDevEndpoint
type jsiiProxy_CfnDevEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDevEndpoint) Arguments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) ExtraJarsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraJarsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) ExtraPythonLibsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraPythonLibsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) NumberOfNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) PublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDevEndpoint) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::DevEndpoint`.
func NewCfnDevEndpoint(scope constructs.Construct, id *string, props *CfnDevEndpointProps) CfnDevEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CfnDevEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnDevEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::DevEndpoint`.
func NewCfnDevEndpoint_Override(c CfnDevEndpoint, scope constructs.Construct, id *string, props *CfnDevEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnDevEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetArguments(val interface{}) {
	_jsii_.Set(
		j,
		"arguments",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetEndpointName(val *string) {
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetExtraJarsS3Path(val *string) {
	_jsii_.Set(
		j,
		"extraJarsS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetExtraPythonLibsS3Path(val *string) {
	_jsii_.Set(
		j,
		"extraPythonLibsS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetGlueVersion(val *string) {
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetNumberOfNodes(val *float64) {
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetPublicKey(val *string) {
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetPublicKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"publicKeys",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_CfnDevEndpoint) SetWorkerType(val *string) {
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDevEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDevEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDevEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDevEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnDevEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnDevEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDevEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnDevEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDevEndpoint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnDevEndpoint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDevEndpoint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnDevEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDevEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDevEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDevEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnDevEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnDevEndpoint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnDevEndpoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDevEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDevEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnDevEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnDevEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDevEndpoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDevEndpoint`.
//
// TODO: EXAMPLE
//
type CfnDevEndpointProps struct {
	// The Amazon Resource Name (ARN) of the IAM role used in this `DevEndpoint` .
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// A map of arguments used to configure the `DevEndpoint` .
	//
	// Valid arguments are:
	//
	// - `"--enable-glue-datacatalog": ""`
	// - `"GLUE_PYTHON_VERSION": "3"`
	// - `"GLUE_PYTHON_VERSION": "2"`
	//
	// You can specify a version of Python support for development endpoints by using the `Arguments` parameter in the `CreateDevEndpoint` or `UpdateDevEndpoint` APIs. If no arguments are provided, the version defaults to Python 2.
	Arguments interface{} `json:"arguments" yaml:"arguments"`
	// The name of the `DevEndpoint` .
	EndpointName *string `json:"endpointName" yaml:"endpointName"`
	// The path to one or more Java `.jar` files in an S3 bucket that should be loaded in your `DevEndpoint` .
	//
	// > You can only use pure Java/Scala libraries with a `DevEndpoint` .
	ExtraJarsS3Path *string `json:"extraJarsS3Path" yaml:"extraJarsS3Path"`
	// The paths to one or more Python libraries in an Amazon S3 bucket that should be loaded in your `DevEndpoint` .
	//
	// Multiple values must be complete paths separated by a comma.
	//
	// > You can only use pure Python libraries with a `DevEndpoint` . Libraries that rely on C extensions, such as the [pandas](https://docs.aws.amazon.com/http://pandas.pydata.org/) Python data analysis library, are not currently supported.
	ExtraPythonLibsS3Path *string `json:"extraPythonLibsS3Path" yaml:"extraPythonLibsS3Path"`
	// The AWS Glue version determines the versions of Apache Spark and Python that AWS Glue supports.
	//
	// The Python version indicates the version supported for running your ETL scripts on development endpoints.
	//
	// For more information about the available AWS Glue versions and corresponding Spark and Python versions, see [Glue version](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) in the developer guide.
	//
	// Development endpoints that are created without specifying a Glue version default to Glue 0.9.
	//
	// You can specify a version of Python support for development endpoints by using the `Arguments` parameter in the `CreateDevEndpoint` or `UpdateDevEndpoint` APIs. If no arguments are provided, the version defaults to Python 2.
	GlueVersion *string `json:"glueVersion" yaml:"glueVersion"`
	// The number of AWS Glue Data Processing Units (DPUs) allocated to this `DevEndpoint` .
	NumberOfNodes *float64 `json:"numberOfNodes" yaml:"numberOfNodes"`
	// The number of workers of a defined `workerType` that are allocated to the development endpoint.
	//
	// The maximum number of workers you can define are 299 for `G.1X` , and 149 for `G.2X` .
	NumberOfWorkers *float64 `json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The public key to be used by this `DevEndpoint` for authentication.
	//
	// This attribute is provided for backward compatibility because the recommended attribute to use is public keys.
	PublicKey *string `json:"publicKey" yaml:"publicKey"`
	// A list of public keys to be used by the `DevEndpoints` for authentication.
	//
	// Using this attribute is preferred over a single public key because the public keys allow you to have a different private key per client.
	//
	// > If you previously created an endpoint with a public key, you must remove that key to be able to set a list of public keys. Call the `UpdateDevEndpoint` API operation with the public key content in the `deletePublicKeys` attribute, and the list of new keys in the `addPublicKeys` attribute.
	PublicKeys *[]*string `json:"publicKeys" yaml:"publicKeys"`
	// The name of the `SecurityConfiguration` structure to be used with this `DevEndpoint` .
	SecurityConfiguration *string `json:"securityConfiguration" yaml:"securityConfiguration"`
	// A list of security group identifiers used in this `DevEndpoint` .
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet ID for this `DevEndpoint` .
	SubnetId *string `json:"subnetId" yaml:"subnetId"`
	// The tags to use with this DevEndpoint.
	Tags interface{} `json:"tags" yaml:"tags"`
	// The type of predefined worker that is allocated to the development endpoint.
	//
	// Accepts a value of Standard, G.1X, or G.2X.
	//
	// - For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	// - For the `G.1X` worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	// - For the `G.2X` worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	//
	// Known issue: when a development endpoint is created with the `G.2X` `WorkerType` configuration, the Spark drivers for the development endpoint will run on 4 vCPU, 16 GB of memory, and a 64 GB disk.
	WorkerType *string `json:"workerType" yaml:"workerType"`
}

// A CloudFormation `AWS::Glue::Job`.
//
// The `AWS::Glue::Job` resource specifies an AWS Glue job in the data catalog. For more information, see [Adding Jobs in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) and [Job Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html#aws-glue-api-jobs-job-Job) in the *AWS Glue Developer Guide.*
//
// TODO: EXAMPLE
//
type CfnJob interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AllocatedCapacity() *float64
	SetAllocatedCapacity(val *float64)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Command() interface{}
	SetCommand(val interface{})
	Connections() interface{}
	SetConnections(val interface{})
	CreationStack() *[]*string
	DefaultArguments() interface{}
	SetDefaultArguments(val interface{})
	Description() *string
	SetDescription(val *string)
	ExecutionProperty() interface{}
	SetExecutionProperty(val interface{})
	GlueVersion() *string
	SetGlueVersion(val *string)
	LogicalId() *string
	LogUri() *string
	SetLogUri(val *string)
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	NotificationProperty() interface{}
	SetNotificationProperty(val interface{})
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	Ref() *string
	Role() *string
	SetRole(val *string)
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Timeout() *float64
	SetTimeout(val *float64)
	UpdatedProperites() *map[string]interface{}
	WorkerType() *string
	SetWorkerType(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnJob
type jsiiProxy_CfnJob struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJob) AllocatedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedCapacity",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnJob) Command() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Connections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connections",
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

func (j *jsiiProxy_CfnJob) DefaultArguments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ExecutionProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
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

func (j *jsiiProxy_CfnJob) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
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

func (j *jsiiProxy_CfnJob) NotificationProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
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

func (j *jsiiProxy_CfnJob) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
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

func (j *jsiiProxy_CfnJob) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Job`.
func NewCfnJob(scope constructs.Construct, id *string, props *CfnJobProps) CfnJob {
	_init_.Initialize()

	j := jsiiProxy_CfnJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Job`.
func NewCfnJob_Override(c CfnJob, scope constructs.Construct, id *string, props *CfnJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnJob",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJob) SetAllocatedCapacity(val *float64) {
	_jsii_.Set(
		j,
		"allocatedCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetCommand(val interface{}) {
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetConnections(val interface{}) {
	_jsii_.Set(
		j,
		"connections",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetDefaultArguments(val interface{}) {
	_jsii_.Set(
		j,
		"defaultArguments",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetExecutionProperty(val interface{}) {
	_jsii_.Set(
		j,
		"executionProperty",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetGlueVersion(val *string) {
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetLogUri(val *string) {
	_jsii_.Set(
		j,
		"logUri",
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

func (j *jsiiProxy_CfnJob) SetNotificationProperty(val interface{}) {
	_jsii_.Set(
		j,
		"notificationProperty",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnJob) SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
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

func (j *jsiiProxy_CfnJob) SetWorkerType(val *string) {
	_jsii_.Set(
		j,
		"workerType",
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
		"aws-cdk-lib.aws_glue.CfnJob",
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
		"aws-cdk-lib.aws_glue.CfnJob",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnJob",
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
		"aws-cdk-lib.aws_glue.CfnJob",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnJob) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnJob) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnJob) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnJob) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnJob) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnJob) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnJob) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

// Specifies the connections used by a job.
//
// TODO: EXAMPLE
//
type CfnJob_ConnectionsListProperty struct {
	// A list of connections used by the job.
	Connections *[]*string `json:"connections" yaml:"connections"`
}

// An execution property of a job.
//
// TODO: EXAMPLE
//
type CfnJob_ExecutionPropertyProperty struct {
	// The maximum number of concurrent runs allowed for the job.
	//
	// The default is 1. An error is returned when this threshold is reached. The maximum value you can specify is controlled by a service limit.
	MaxConcurrentRuns *float64 `json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
}

// Specifies code executed when a job is run.
//
// TODO: EXAMPLE
//
type CfnJob_JobCommandProperty struct {
	// The name of the job command.
	//
	// For an Apache Spark ETL job, this must be `glueetl` . For a Python shell job, it must be `pythonshell` .
	Name *string `json:"name" yaml:"name"`
	// The Python version being used to execute a Python shell job.
	//
	// Allowed values are 2 or 3.
	PythonVersion *string `json:"pythonVersion" yaml:"pythonVersion"`
	// Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job (required).
	ScriptLocation *string `json:"scriptLocation" yaml:"scriptLocation"`
}

// Specifies configuration properties of a notification.
//
// TODO: EXAMPLE
//
type CfnJob_NotificationPropertyProperty struct {
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	NotifyDelayAfter *float64 `json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

// Properties for defining a `CfnJob`.
//
// TODO: EXAMPLE
//
type CfnJobProps struct {
	// The code that executes a job.
	Command interface{} `json:"command" yaml:"command"`
	// The name or Amazon Resource Name (ARN) of the IAM role associated with this job.
	Role *string `json:"role" yaml:"role"`
	// The number of capacity units that are allocated to this job.
	AllocatedCapacity *float64 `json:"allocatedCapacity" yaml:"allocatedCapacity"`
	// The connections used for this job.
	Connections interface{} `json:"connections" yaml:"connections"`
	// The default arguments for this job, specified as name-value pairs.
	//
	// You can specify arguments here that your own job-execution script consumes, in addition to arguments that AWS Glue itself consumes.
	//
	// For information about how to specify and consume your own job arguments, see [Calling AWS Glue APIs in Python](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) in the *AWS Glue Developer Guide* .
	//
	// For information about the key-value pairs that AWS Glue consumes to set up your job, see [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) in the *AWS Glue Developer Guide* .
	DefaultArguments interface{} `json:"defaultArguments" yaml:"defaultArguments"`
	// A description of the job.
	Description *string `json:"description" yaml:"description"`
	// The maximum number of concurrent runs that are allowed for this job.
	ExecutionProperty interface{} `json:"executionProperty" yaml:"executionProperty"`
	// Glue version determines the versions of Apache Spark and Python that AWS Glue supports.
	//
	// The Python version indicates the version supported for jobs of type Spark.
	//
	// For more information about the available AWS Glue versions and corresponding Spark and Python versions, see [Glue version](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) in the developer guide.
	//
	// Jobs that are created without specifying a Glue version default to Glue 0.9.
	GlueVersion *string `json:"glueVersion" yaml:"glueVersion"`
	// This field is reserved for future use.
	LogUri *string `json:"logUri" yaml:"logUri"`
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	//
	// Do not set `Max Capacity` if using `WorkerType` and `NumberOfWorkers` .
	//
	// The value that can be allocated for `MaxCapacity` depends on whether you are running a Python shell job or an Apache Spark ETL job:
	//
	// - When you specify a Python shell job ( `JobCommand.Name` ="pythonshell"), you can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	// - When you specify an Apache Spark ETL job ( `JobCommand.Name` ="glueetl"), you can allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type cannot have a fractional DPU allocation.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry this job after a JobRun fails.
	MaxRetries *float64 `json:"maxRetries" yaml:"maxRetries"`
	// The name you assign to this job definition.
	Name *string `json:"name" yaml:"name"`
	// Specifies configuration properties of a notification.
	NotificationProperty interface{} `json:"notificationProperty" yaml:"notificationProperty"`
	// The number of workers of a defined `workerType` that are allocated when a job runs.
	//
	// The maximum number of workers you can define are 299 for `G.1X` , and 149 for `G.2X` .
	NumberOfWorkers *float64 `json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The name of the `SecurityConfiguration` structure to be used with this job.
	SecurityConfiguration *string `json:"securityConfiguration" yaml:"securityConfiguration"`
	// The tags to use with this job.
	Tags interface{} `json:"tags" yaml:"tags"`
	// The job timeout in minutes.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	Timeout *float64 `json:"timeout" yaml:"timeout"`
	// The type of predefined worker that is allocated when a job runs.
	//
	// Accepts a value of Standard, G.1X, or G.2X.
	//
	// - For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	// - For the `G.1X` worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	// - For the `G.2X` worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	WorkerType *string `json:"workerType" yaml:"workerType"`
}

// A CloudFormation `AWS::Glue::MLTransform`.
//
// The AWS::Glue::MLTransform is an AWS Glue resource type that manages machine learning transforms.
//
// TODO: EXAMPLE
//
type CfnMLTransform interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	GlueVersion() *string
	SetGlueVersion(val *string)
	InputRecordTables() interface{}
	SetInputRecordTables(val interface{})
	LogicalId() *string
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	Ref() *string
	Role() *string
	SetRole(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Timeout() *float64
	SetTimeout(val *float64)
	TransformEncryption() interface{}
	SetTransformEncryption(val interface{})
	TransformParameters() interface{}
	SetTransformParameters(val interface{})
	UpdatedProperites() *map[string]interface{}
	WorkerType() *string
	SetWorkerType(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMLTransform
type jsiiProxy_CfnMLTransform struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMLTransform) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) InputRecordTables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputRecordTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) TransformEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) TransformParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMLTransform) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::MLTransform`.
func NewCfnMLTransform(scope constructs.Construct, id *string, props *CfnMLTransformProps) CfnMLTransform {
	_init_.Initialize()

	j := jsiiProxy_CfnMLTransform{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnMLTransform",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::MLTransform`.
func NewCfnMLTransform_Override(c CfnMLTransform, scope constructs.Construct, id *string, props *CfnMLTransformProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnMLTransform",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetGlueVersion(val *string) {
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetInputRecordTables(val interface{}) {
	_jsii_.Set(
		j,
		"inputRecordTables",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetTransformEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"transformEncryption",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetTransformParameters(val interface{}) {
	_jsii_.Set(
		j,
		"transformParameters",
		val,
	)
}

func (j *jsiiProxy_CfnMLTransform) SetWorkerType(val *string) {
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMLTransform_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnMLTransform",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMLTransform_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnMLTransform",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnMLTransform_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnMLTransform",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMLTransform_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnMLTransform",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnMLTransform) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnMLTransform) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnMLTransform) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnMLTransform) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnMLTransform) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnMLTransform) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnMLTransform) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnMLTransform) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnMLTransform) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnMLTransform) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnMLTransform) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMLTransform) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnMLTransform) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnMLTransform) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMLTransform) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The parameters to configure the find matches transform.
//
// TODO: EXAMPLE
//
type CfnMLTransform_FindMatchesParametersProperty struct {
	// The name of a column that uniquely identifies rows in the source table.
	//
	// Used to help identify matching records.
	PrimaryKeyColumnName *string `json:"primaryKeyColumnName" yaml:"primaryKeyColumnName"`
	// The value that is selected when tuning your transform for a balance between accuracy and cost.
	//
	// A value of 0.5 means that the system balances accuracy and cost concerns. A value of 1.0 means a bias purely for accuracy, which typically results in a higher cost, sometimes substantially higher. A value of 0.0 means a bias purely for cost, which results in a less accurate `FindMatches` transform, sometimes with unacceptable accuracy.
	//
	// Accuracy measures how well the transform finds true positives and true negatives. Increasing accuracy requires more machine resources and cost. But it also results in increased recall.
	//
	// Cost measures how many compute resources, and thus money, are consumed to run the transform.
	AccuracyCostTradeoff *float64 `json:"accuracyCostTradeoff" yaml:"accuracyCostTradeoff"`
	// The value to switch on or off to force the output to match the provided labels from users.
	//
	// If the value is `True` , the `find matches` transform forces the output to match the provided labels. The results override the normal conflation results. If the value is `False` , the `find matches` transform does not ensure all the labels provided are respected, and the results rely on the trained model.
	//
	// Note that setting this value to true may increase the conflation execution time.
	EnforceProvidedLabels interface{} `json:"enforceProvidedLabels" yaml:"enforceProvidedLabels"`
	// The value selected when tuning your transform for a balance between precision and recall.
	//
	// A value of 0.5 means no preference; a value of 1.0 means a bias purely for precision, and a value of 0.0 means a bias for recall. Because this is a tradeoff, choosing values close to 1.0 means very low recall, and choosing values close to 0.0 results in very low precision.
	//
	// The precision metric indicates how often your model is correct when it predicts a match.
	//
	// The recall metric indicates that for an actual match, how often your model predicts the match.
	PrecisionRecallTradeoff *float64 `json:"precisionRecallTradeoff" yaml:"precisionRecallTradeoff"`
}

// The database and table in the AWS Glue Data Catalog that is used for input or output data.
//
// TODO: EXAMPLE
//
type CfnMLTransform_GlueTablesProperty struct {
	// A database name in the AWS Glue Data Catalog .
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// A table name in the AWS Glue Data Catalog .
	TableName *string `json:"tableName" yaml:"tableName"`
	// A unique identifier for the AWS Glue Data Catalog .
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The name of the connection to the AWS Glue Data Catalog .
	ConnectionName *string `json:"connectionName" yaml:"connectionName"`
}

// A list of AWS Glue table definitions used by the transform.
//
// TODO: EXAMPLE
//
type CfnMLTransform_InputRecordTablesProperty struct {
	// The database and table in the AWS Glue Data Catalog that is used for input or output data.
	GlueTables interface{} `json:"glueTables" yaml:"glueTables"`
}

// The encryption-at-rest settings of the transform that apply to accessing user data.
//
// TODO: EXAMPLE
//
type CfnMLTransform_MLUserDataEncryptionProperty struct {
	// The encryption mode applied to user data. Valid values are:.
	//
	// - DISABLED: encryption is disabled.
	// - SSEKMS: use of server-side encryption with AWS Key Management Service (SSE-KMS) for user data
	// stored in Amazon S3.
	MlUserDataEncryptionMode *string `json:"mlUserDataEncryptionMode" yaml:"mlUserDataEncryptionMode"`
	// The ID for the customer-provided KMS key.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
}

// The encryption-at-rest settings of the transform that apply to accessing user data.
//
// Machine learning
// transforms can access user data encrypted in Amazon S3 using KMS.
//
// Additionally, imported labels and trained transforms can now be encrypted using a customer provided
// KMS key.
//
// TODO: EXAMPLE
//
type CfnMLTransform_TransformEncryptionProperty struct {
	// The encryption-at-rest settings of the transform that apply to accessing user data.
	MlUserDataEncryption interface{} `json:"mlUserDataEncryption" yaml:"mlUserDataEncryption"`
	// The name of the security configuration.
	TaskRunSecurityConfigurationName *string `json:"taskRunSecurityConfigurationName" yaml:"taskRunSecurityConfigurationName"`
}

// The algorithm-specific parameters that are associated with the machine learning transform.
//
// TODO: EXAMPLE
//
type CfnMLTransform_TransformParametersProperty struct {
	// The type of machine learning transform. `FIND_MATCHES` is the only option.
	//
	// For information about the types of machine learning transforms, see [Creating Machine Learning Transforms](https://docs.aws.amazon.com/glue/latest/dg/add-job-machine-learning-transform.html) .
	TransformType *string `json:"transformType" yaml:"transformType"`
	// The parameters for the find matches algorithm.
	FindMatchesParameters interface{} `json:"findMatchesParameters" yaml:"findMatchesParameters"`
}

// Properties for defining a `CfnMLTransform`.
//
// TODO: EXAMPLE
//
type CfnMLTransformProps struct {
	// A list of AWS Glue table definitions used by the transform.
	InputRecordTables interface{} `json:"inputRecordTables" yaml:"inputRecordTables"`
	// The name or Amazon Resource Name (ARN) of the IAM role with the required permissions.
	//
	// The required permissions include both AWS Glue service role permissions to AWS Glue resources, and Amazon S3 permissions required by the transform.
	//
	// - This role needs AWS Glue service role permissions to allow access to resources in AWS Glue . See [Attach a Policy to IAM Users That Access AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/attach-policy-iam-user.html) .
	// - This role needs permission to your Amazon Simple Storage Service (Amazon S3) sources, targets, temporary directory, scripts, and any libraries used by the task run for this transform.
	Role *string `json:"role" yaml:"role"`
	// The algorithm-specific parameters that are associated with the machine learning transform.
	TransformParameters interface{} `json:"transformParameters" yaml:"transformParameters"`
	// A user-defined, long-form description text for the machine learning transform.
	Description *string `json:"description" yaml:"description"`
	// This value determines which version of AWS Glue this machine learning transform is compatible with.
	//
	// Glue 1.0 is recommended for most customers. If the value is not set, the Glue compatibility defaults to Glue 0.9. For more information, see [AWS Glue Versions](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html#release-notes-versions) in the developer guide.
	GlueVersion *string `json:"glueVersion" yaml:"glueVersion"`
	// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform.
	//
	// You can allocate from 2 to 100 DPUs; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory. For more information, see the [AWS Glue pricing page](https://docs.aws.amazon.com/glue/pricing/) .
	//
	// `MaxCapacity` is a mutually exclusive option with `NumberOfWorkers` and `WorkerType` .
	//
	// - If either `NumberOfWorkers` or `WorkerType` is set, then `MaxCapacity` cannot be set.
	// - If `MaxCapacity` is set then neither `NumberOfWorkers` or `WorkerType` can be set.
	// - If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).
	// - `MaxCapacity` and `NumberOfWorkers` must both be at least 1.
	//
	// When the `WorkerType` field is set to a value other than `Standard` , the `MaxCapacity` field is set automatically and becomes read-only.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry after an `MLTaskRun` of the machine learning transform fails.
	MaxRetries *float64 `json:"maxRetries" yaml:"maxRetries"`
	// A user-defined name for the machine learning transform. Names are required to be unique. `Name` is optional:.
	//
	// - If you supply `Name` , the stack cannot be repeatedly created.
	// - If `Name` is not provided, a randomly generated name will be used instead.
	Name *string `json:"name" yaml:"name"`
	// The number of workers of a defined `workerType` that are allocated when a task of the transform runs.
	//
	// If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).
	NumberOfWorkers *float64 `json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The tags to use with this machine learning transform.
	//
	// You may use tags to limit access to the machine learning transform. For more information about tags in AWS Glue , see [AWS Tags in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in the developer guide.
	Tags interface{} `json:"tags" yaml:"tags"`
	// The timeout in minutes of the machine learning transform.
	Timeout *float64 `json:"timeout" yaml:"timeout"`
	// The encryption-at-rest settings of the transform that apply to accessing user data.
	//
	// Machine learning
	// transforms can access user data encrypted in Amazon S3 using KMS.
	//
	// Additionally, imported labels and trained transforms can now be encrypted using a customer provided
	// KMS key.
	TransformEncryption interface{} `json:"transformEncryption" yaml:"transformEncryption"`
	// The type of predefined worker that is allocated when a task of this transform runs.
	//
	// Accepts a value of Standard, G.1X, or G.2X.
	//
	// - For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	// - For the `G.1X` worker type, each worker provides 4 vCPU, 16 GB of memory and a 64GB disk, and 1 executor per worker.
	// - For the `G.2X` worker type, each worker provides 8 vCPU, 32 GB of memory and a 128GB disk, and 1 executor per worker.
	//
	// `MaxCapacity` is a mutually exclusive option with `NumberOfWorkers` and `WorkerType` .
	//
	// - If either `NumberOfWorkers` or `WorkerType` is set, then `MaxCapacity` cannot be set.
	// - If `MaxCapacity` is set then neither `NumberOfWorkers` or `WorkerType` can be set.
	// - If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).
	// - `MaxCapacity` and `NumberOfWorkers` must both be at least 1.
	WorkerType *string `json:"workerType" yaml:"workerType"`
}

// A CloudFormation `AWS::Glue::Partition`.
//
// The `AWS::Glue::Partition` resource creates an AWS Glue partition, which represents a slice of table data. For more information, see [CreatePartition Action](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-partitions.html#aws-glue-api-catalog-partitions-CreatePartition) and [Partition Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-partitions.html#aws-glue-api-catalog-partitions-Partition) in the *AWS Glue Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnPartition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CatalogId() *string
	SetCatalogId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	LogicalId() *string
	Node() constructs.Node
	PartitionInput() interface{}
	SetPartitionInput(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	TableName() *string
	SetTableName(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPartition
type jsiiProxy_CfnPartition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPartition) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) PartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPartition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Partition`.
func NewCfnPartition(scope constructs.Construct, id *string, props *CfnPartitionProps) CfnPartition {
	_init_.Initialize()

	j := jsiiProxy_CfnPartition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnPartition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Partition`.
func NewCfnPartition_Override(c CfnPartition, scope constructs.Construct, id *string, props *CfnPartitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnPartition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPartition) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_CfnPartition) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnPartition) SetPartitionInput(val interface{}) {
	_jsii_.Set(
		j,
		"partitionInput",
		val,
	)
}

func (j *jsiiProxy_CfnPartition) SetTableName(val *string) {
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
func CfnPartition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnPartition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPartition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnPartition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnPartition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnPartition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPartition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnPartition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPartition) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnPartition) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnPartition) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnPartition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPartition) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPartition) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnPartition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnPartition) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnPartition) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnPartition) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPartition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPartition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnPartition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnPartition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPartition) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A column in a `Table` .
//
// TODO: EXAMPLE
//
type CfnPartition_ColumnProperty struct {
	// The name of the `Column` .
	Name *string `json:"name" yaml:"name"`
	// A free-form text comment.
	Comment *string `json:"comment" yaml:"comment"`
	// The data type of the `Column` .
	Type *string `json:"type" yaml:"type"`
}

// Specifies the sort order of a sorted column.
//
// TODO: EXAMPLE
//
type CfnPartition_OrderProperty struct {
	// The name of the column.
	Column *string `json:"column" yaml:"column"`
	// Indicates that the column is sorted in ascending order ( `== 1` ), or in descending order ( `==0` ).
	SortOrder *float64 `json:"sortOrder" yaml:"sortOrder"`
}

// The structure used to create and update a partition.
//
// TODO: EXAMPLE
//
type CfnPartition_PartitionInputProperty struct {
	// The values of the partition.
	//
	// Although this parameter is not required by the SDK, you must specify this parameter for a valid input.
	//
	// The values for the keys for the new partition must be passed as an array of String objects that must be ordered in the same order as the partition keys appearing in the Amazon S3 prefix. Otherwise AWS Glue will add the values to the wrong keys.
	Values *[]*string `json:"values" yaml:"values"`
	// These key-value pairs define partition parameters.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// Provides information about the physical location where the partition is stored.
	StorageDescriptor interface{} `json:"storageDescriptor" yaml:"storageDescriptor"`
}

// A structure that contains schema identity fields.
//
// Either this or the `SchemaVersionId` has to be
// provided.
//
// TODO: EXAMPLE
//
type CfnPartition_SchemaIdProperty struct {
	// The name of the schema registry that contains the schema.
	RegistryName *string `json:"registryName" yaml:"registryName"`
	// The Amazon Resource Name (ARN) of the schema.
	//
	// One of `SchemaArn` or `SchemaName` has to be
	// provided.
	SchemaArn *string `json:"schemaArn" yaml:"schemaArn"`
	// The name of the schema.
	//
	// One of `SchemaArn` or `SchemaName` has to be provided.
	SchemaName *string `json:"schemaName" yaml:"schemaName"`
}

// An object that references a schema stored in the AWS Glue Schema Registry.
//
// TODO: EXAMPLE
//
type CfnPartition_SchemaReferenceProperty struct {
	// A structure that contains schema identity fields.
	//
	// Either this or the `SchemaVersionId` has to be
	// provided.
	SchemaId interface{} `json:"schemaId" yaml:"schemaId"`
	// The unique ID assigned to a version of the schema.
	//
	// Either this or the `SchemaId` has to be provided.
	SchemaVersionId *string `json:"schemaVersionId" yaml:"schemaVersionId"`
	// The version number of the schema.
	SchemaVersionNumber *float64 `json:"schemaVersionNumber" yaml:"schemaVersionNumber"`
}

// Information about a serialization/deserialization program (SerDe) that serves as an extractor and loader.
//
// TODO: EXAMPLE
//
type CfnPartition_SerdeInfoProperty struct {
	// Name of the SerDe.
	Name *string `json:"name" yaml:"name"`
	// These key-value pairs define initialization parameters for the SerDe.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// Usually the class that implements the SerDe.
	//
	// An example is `org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe` .
	SerializationLibrary *string `json:"serializationLibrary" yaml:"serializationLibrary"`
}

// Specifies skewed values in a table.
//
// Skewed values are those that occur with very high frequency.
//
// TODO: EXAMPLE
//
type CfnPartition_SkewedInfoProperty struct {
	// A list of names of columns that contain skewed values.
	SkewedColumnNames *[]*string `json:"skewedColumnNames" yaml:"skewedColumnNames"`
	// A mapping of skewed values to the columns that contain them.
	SkewedColumnValueLocationMaps interface{} `json:"skewedColumnValueLocationMaps" yaml:"skewedColumnValueLocationMaps"`
	// A list of values that appear so frequently as to be considered skewed.
	SkewedColumnValues *[]*string `json:"skewedColumnValues" yaml:"skewedColumnValues"`
}

// Describes the physical storage of table data.
//
// TODO: EXAMPLE
//
type CfnPartition_StorageDescriptorProperty struct {
	// A list of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns *[]*string `json:"bucketColumns" yaml:"bucketColumns"`
	// A list of the `Columns` in the table.
	Columns interface{} `json:"columns" yaml:"columns"`
	// `True` if the data in the table is compressed, or `False` if not.
	Compressed interface{} `json:"compressed" yaml:"compressed"`
	// The input format: `SequenceFileInputFormat` (binary), or `TextInputFormat` , or a custom format.
	InputFormat *string `json:"inputFormat" yaml:"inputFormat"`
	// The physical location of the table.
	//
	// By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location *string `json:"location" yaml:"location"`
	// The number of buckets.
	//
	// You must specify this property if the partition contains any dimension columns.
	NumberOfBuckets *float64 `json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// The output format: `SequenceFileOutputFormat` (binary), or `IgnoreKeyTextOutputFormat` , or a custom format.
	OutputFormat *string `json:"outputFormat" yaml:"outputFormat"`
	// The user-supplied properties in key-value form.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// An object that references a schema stored in the AWS Glue Schema Registry.
	SchemaReference interface{} `json:"schemaReference" yaml:"schemaReference"`
	// The serialization/deserialization (SerDe) information.
	SerdeInfo interface{} `json:"serdeInfo" yaml:"serdeInfo"`
	// The information about values that appear frequently in a column (skewed values).
	SkewedInfo interface{} `json:"skewedInfo" yaml:"skewedInfo"`
	// A list specifying the sort order of each bucket in the table.
	SortColumns interface{} `json:"sortColumns" yaml:"sortColumns"`
	// `True` if the table data is stored in subdirectories, or `False` if not.
	StoredAsSubDirectories interface{} `json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

// Properties for defining a `CfnPartition`.
//
// TODO: EXAMPLE
//
type CfnPartitionProps struct {
	// The AWS account ID of the catalog in which the partion is to be created.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId`
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database in which to create the partition.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The structure used to create and update a partition.
	PartitionInput interface{} `json:"partitionInput" yaml:"partitionInput"`
	// The name of the metadata table in which the partition is to be created.
	TableName *string `json:"tableName" yaml:"tableName"`
}

// A CloudFormation `AWS::Glue::Registry`.
//
// The AWS::Glue::Registry is an AWS Glue resource type that manages registries of schemas in the AWS Glue Schema Registry.
//
// TODO: EXAMPLE
//
type CfnRegistry interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRegistry
type jsiiProxy_CfnRegistry struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRegistry) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRegistry) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Registry`.
func NewCfnRegistry(scope constructs.Construct, id *string, props *CfnRegistryProps) CfnRegistry {
	_init_.Initialize()

	j := jsiiProxy_CfnRegistry{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnRegistry",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Registry`.
func NewCfnRegistry_Override(c CfnRegistry, scope constructs.Construct, id *string, props *CfnRegistryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnRegistry",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRegistry) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRegistry) SetName(val *string) {
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
func CfnRegistry_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnRegistry",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRegistry_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnRegistry",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnRegistry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnRegistry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRegistry_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnRegistry",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRegistry) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnRegistry) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRegistry) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnRegistry) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRegistry) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRegistry) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnRegistry) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnRegistry) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnRegistry) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnRegistry) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRegistry) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRegistry) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnRegistry) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnRegistry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRegistry) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnRegistry`.
//
// TODO: EXAMPLE
//
type CfnRegistryProps struct {
	// The name of the registry.
	Name *string `json:"name" yaml:"name"`
	// A description of the registry.
	Description *string `json:"description" yaml:"description"`
	// AWS tags that contain a key value pair and may be searched by console, command line, or API.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Glue::Schema`.
//
// The `AWS::Glue::Schema` is an AWS Glue resource type that manages schemas in the AWS Glue Schema Registry.
//
// TODO: EXAMPLE
//
type CfnSchema interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrInitialSchemaVersionId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CheckpointVersion() interface{}
	SetCheckpointVersion(val interface{})
	Compatibility() *string
	SetCompatibility(val *string)
	CreationStack() *[]*string
	DataFormat() *string
	SetDataFormat(val *string)
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Registry() interface{}
	SetRegistry(val interface{})
	SchemaDefinition() *string
	SetSchemaDefinition(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSchema
type jsiiProxy_CfnSchema struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSchema) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) AttrInitialSchemaVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrInitialSchemaVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) CheckpointVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) Compatibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) Registry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) SchemaDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchema) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Schema`.
func NewCfnSchema(scope constructs.Construct, id *string, props *CfnSchemaProps) CfnSchema {
	_init_.Initialize()

	j := jsiiProxy_CfnSchema{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnSchema",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Schema`.
func NewCfnSchema_Override(c CfnSchema, scope constructs.Construct, id *string, props *CfnSchemaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnSchema",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSchema) SetCheckpointVersion(val interface{}) {
	_jsii_.Set(
		j,
		"checkpointVersion",
		val,
	)
}

func (j *jsiiProxy_CfnSchema) SetCompatibility(val *string) {
	_jsii_.Set(
		j,
		"compatibility",
		val,
	)
}

func (j *jsiiProxy_CfnSchema) SetDataFormat(val *string) {
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_CfnSchema) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnSchema) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSchema) SetRegistry(val interface{}) {
	_jsii_.Set(
		j,
		"registry",
		val,
	)
}

func (j *jsiiProxy_CfnSchema) SetSchemaDefinition(val *string) {
	_jsii_.Set(
		j,
		"schemaDefinition",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSchema_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchema",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSchema_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchema",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnSchema_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchema",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchema_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnSchema",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSchema) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnSchema) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSchema) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnSchema) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSchema) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSchema) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnSchema) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnSchema) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSchema) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnSchema) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSchema) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSchema) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnSchema) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnSchema) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchema) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies a registry in the AWS Glue Schema Registry.
//
// TODO: EXAMPLE
//
type CfnSchema_RegistryProperty struct {
	// The Amazon Resource Name (ARN) of the registry.
	Arn *string `json:"arn" yaml:"arn"`
	// The name of the registry.
	Name *string `json:"name" yaml:"name"`
}

// Specifies the version of a schema.
//
// TODO: EXAMPLE
//
type CfnSchema_SchemaVersionProperty struct {
	// Indicates if this version is the latest version of the schema.
	IsLatest interface{} `json:"isLatest" yaml:"isLatest"`
	// The version number of the schema.
	VersionNumber *float64 `json:"versionNumber" yaml:"versionNumber"`
}

// Properties for defining a `CfnSchema`.
//
// TODO: EXAMPLE
//
type CfnSchemaProps struct {
	// The compatibility mode of the schema.
	Compatibility *string `json:"compatibility" yaml:"compatibility"`
	// The data format of the schema definition.
	//
	// Currently only `AVRO` is supported.
	DataFormat *string `json:"dataFormat" yaml:"dataFormat"`
	// Name of the schema to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.
	//
	// No whitespace.
	Name *string `json:"name" yaml:"name"`
	// The schema definition using the `DataFormat` setting for `SchemaName` .
	SchemaDefinition *string `json:"schemaDefinition" yaml:"schemaDefinition"`
	// Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema.
	//
	// This is only required for updating a checkpoint.
	CheckpointVersion interface{} `json:"checkpointVersion" yaml:"checkpointVersion"`
	// A description of the schema if specified when created.
	Description *string `json:"description" yaml:"description"`
	// The registry where a schema is stored.
	Registry interface{} `json:"registry" yaml:"registry"`
	// AWS tags that contain a key value pair and may be searched by console, command line, or API.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Glue::SchemaVersion`.
//
// The `AWS::Glue::SchemaVersion` is an AWS Glue resource type that manages schema versions of schemas in the AWS Glue Schema Registry.
//
// TODO: EXAMPLE
//
type CfnSchemaVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrVersionId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Schema() interface{}
	SetSchema(val interface{})
	SchemaDefinition() *string
	SetSchemaDefinition(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSchemaVersion
type jsiiProxy_CfnSchemaVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSchemaVersion) AttrVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) Schema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) SchemaDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::SchemaVersion`.
func NewCfnSchemaVersion(scope constructs.Construct, id *string, props *CfnSchemaVersionProps) CfnSchemaVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnSchemaVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnSchemaVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::SchemaVersion`.
func NewCfnSchemaVersion_Override(c CfnSchemaVersion, scope constructs.Construct, id *string, props *CfnSchemaVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnSchemaVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSchemaVersion) SetSchema(val interface{}) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_CfnSchemaVersion) SetSchemaDefinition(val *string) {
	_jsii_.Set(
		j,
		"schemaDefinition",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSchemaVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchemaVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSchemaVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchemaVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnSchemaVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchemaVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchemaVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnSchemaVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSchemaVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnSchemaVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSchemaVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnSchemaVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSchemaVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSchemaVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnSchemaVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnSchemaVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSchemaVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnSchemaVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSchemaVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSchemaVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnSchemaVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnSchemaVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchemaVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A wrapper structure to contain schema identity fields.
//
// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
//
// TODO: EXAMPLE
//
type CfnSchemaVersion_SchemaProperty struct {
	// The name of the registry where the schema is stored.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	RegistryName *string `json:"registryName" yaml:"registryName"`
	// The Amazon Resource Name (ARN) of the schema.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	SchemaArn *string `json:"schemaArn" yaml:"schemaArn"`
	// The name of the schema.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	SchemaName *string `json:"schemaName" yaml:"schemaName"`
}

// A CloudFormation `AWS::Glue::SchemaVersionMetadata`.
//
// The `AWS::Glue::SchemaVersionMetadata` is an AWS Glue resource type that defines the metadata key-value pairs for a schema version in AWS Glue Schema Registry.
//
// TODO: EXAMPLE
//
type CfnSchemaVersionMetadata interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Key() *string
	SetKey(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SchemaVersionId() *string
	SetSchemaVersionId(val *string)
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	Value() *string
	SetValue(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSchemaVersionMetadata
type jsiiProxy_CfnSchemaVersionMetadata struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) SchemaVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::SchemaVersionMetadata`.
func NewCfnSchemaVersionMetadata(scope constructs.Construct, id *string, props *CfnSchemaVersionMetadataProps) CfnSchemaVersionMetadata {
	_init_.Initialize()

	j := jsiiProxy_CfnSchemaVersionMetadata{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnSchemaVersionMetadata",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::SchemaVersionMetadata`.
func NewCfnSchemaVersionMetadata_Override(c CfnSchemaVersionMetadata, scope constructs.Construct, id *string, props *CfnSchemaVersionMetadataProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnSchemaVersionMetadata",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) SetSchemaVersionId(val *string) {
	_jsii_.Set(
		j,
		"schemaVersionId",
		val,
	)
}

func (j *jsiiProxy_CfnSchemaVersionMetadata) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSchemaVersionMetadata_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchemaVersionMetadata",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSchemaVersionMetadata_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchemaVersionMetadata",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnSchemaVersionMetadata_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSchemaVersionMetadata",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchemaVersionMetadata_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnSchemaVersionMetadata",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSchemaVersionMetadata) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnSchemaVersionMetadata) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSchemaVersionMetadata) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnSchemaVersionMetadata) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSchemaVersionMetadata) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSchemaVersionMetadata) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnSchemaVersionMetadata) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnSchemaVersionMetadata) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSchemaVersionMetadata) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnSchemaVersionMetadata) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSchemaVersionMetadata) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSchemaVersionMetadata) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnSchemaVersionMetadata) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnSchemaVersionMetadata) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSchemaVersionMetadata) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSchemaVersionMetadata`.
//
// TODO: EXAMPLE
//
type CfnSchemaVersionMetadataProps struct {
	// A metadata key in a key-value pair for metadata.
	Key *string `json:"key" yaml:"key"`
	// The version number of the schema.
	SchemaVersionId *string `json:"schemaVersionId" yaml:"schemaVersionId"`
	// A metadata key's corresponding value.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnSchemaVersion`.
//
// TODO: EXAMPLE
//
type CfnSchemaVersionProps struct {
	// The schema that includes the schema version.
	Schema interface{} `json:"schema" yaml:"schema"`
	// The schema definition for the schema version.
	SchemaDefinition *string `json:"schemaDefinition" yaml:"schemaDefinition"`
}

// A CloudFormation `AWS::Glue::SecurityConfiguration`.
//
// Creates a new security configuration. A security configuration is a set of security properties that can be used by AWS Glue . You can use a security configuration to encrypt data at rest. For information about using security configurations in AWS Glue , see [Encrypting Data Written by Crawlers, Jobs, and Development Endpoints](https://docs.aws.amazon.com/glue/latest/dg/encryption-security-configuration.html) .
//
// TODO: EXAMPLE
//
type CfnSecurityConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EncryptionConfiguration() interface{}
	SetEncryptionConfiguration(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnSecurityConfiguration
type jsiiProxy_CfnSecurityConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) EncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::SecurityConfiguration`.
func NewCfnSecurityConfiguration(scope constructs.Construct, id *string, props *CfnSecurityConfigurationProps) CfnSecurityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnSecurityConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnSecurityConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::SecurityConfiguration`.
func NewCfnSecurityConfiguration_Override(c CfnSecurityConfiguration, scope constructs.Construct, id *string, props *CfnSecurityConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnSecurityConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSecurityConfiguration) SetEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityConfiguration) SetName(val *string) {
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
func CfnSecurityConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSecurityConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSecurityConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSecurityConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnSecurityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnSecurityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnSecurityConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSecurityConfiguration) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnSecurityConfiguration) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSecurityConfiguration) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnSecurityConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSecurityConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSecurityConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnSecurityConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnSecurityConfiguration) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnSecurityConfiguration) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnSecurityConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSecurityConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnSecurityConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnSecurityConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies how Amazon CloudWatch data should be encrypted.
//
// TODO: EXAMPLE
//
type CfnSecurityConfiguration_CloudWatchEncryptionProperty struct {
	// The encryption mode to use for CloudWatch data.
	CloudWatchEncryptionMode *string `json:"cloudWatchEncryptionMode" yaml:"cloudWatchEncryptionMode"`
	// The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

// Specifies an encryption configuration.
//
// TODO: EXAMPLE
//
type CfnSecurityConfiguration_EncryptionConfigurationProperty struct {
	// The encryption configuration for Amazon CloudWatch.
	CloudWatchEncryption interface{} `json:"cloudWatchEncryption" yaml:"cloudWatchEncryption"`
	// The encryption configuration for job bookmarks.
	JobBookmarksEncryption interface{} `json:"jobBookmarksEncryption" yaml:"jobBookmarksEncryption"`
	// The encyption configuration for Amazon Simple Storage Service (Amazon S3) data.
	S3Encryptions interface{} `json:"s3Encryptions" yaml:"s3Encryptions"`
}

// Specifies how job bookmark data should be encrypted.
//
// TODO: EXAMPLE
//
type CfnSecurityConfiguration_JobBookmarksEncryptionProperty struct {
	// The encryption mode to use for job bookmarks data.
	JobBookmarksEncryptionMode *string `json:"jobBookmarksEncryptionMode" yaml:"jobBookmarksEncryptionMode"`
	// The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

// Specifies how Amazon Simple Storage Service (Amazon S3) data should be encrypted.
//
// TODO: EXAMPLE
//
type CfnSecurityConfiguration_S3EncryptionProperty struct {
	// The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The encryption mode to use for Amazon S3 data.
	S3EncryptionMode *string `json:"s3EncryptionMode" yaml:"s3EncryptionMode"`
}

// Properties for defining a `CfnSecurityConfiguration`.
//
// TODO: EXAMPLE
//
type CfnSecurityConfigurationProps struct {
	// The encryption configuration associated with this security configuration.
	EncryptionConfiguration interface{} `json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The name of the security configuration.
	Name *string `json:"name" yaml:"name"`
}

// A CloudFormation `AWS::Glue::Table`.
//
// The `AWS::Glue::Table` resource specifies tabular data in the AWS Glue data catalog. For more information, see [Defining Tables in the AWS Glue Data Catalog](https://docs.aws.amazon.com/glue/latest/dg/tables-described.html) and [Table Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-Table) in the *AWS Glue Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnTable interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CatalogId() *string
	SetCatalogId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	TableInput() interface{}
	SetTableInput(val interface{})
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnTable
type jsiiProxy_CfnTable struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTable) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Table`.
func NewCfnTable(scope constructs.Construct, id *string, props *CfnTableProps) CfnTable {
	_init_.Initialize()

	j := jsiiProxy_CfnTable{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Table`.
func NewCfnTable_Override(c CfnTable, scope constructs.Construct, id *string, props *CfnTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnTable",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTable) SetCatalogId(val *string) {
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetTableInput(val interface{}) {
	_jsii_.Set(
		j,
		"tableInput",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTable_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnTable",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTable_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnTable",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTable_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnTable",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTable) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnTable) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnTable) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnTable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTable) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTable) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnTable) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnTable) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnTable) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTable) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnTable) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTable) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A column in a `Table` .
//
// TODO: EXAMPLE
//
type CfnTable_ColumnProperty struct {
	// The name of the `Column` .
	Name *string `json:"name" yaml:"name"`
	// A free-form text comment.
	Comment *string `json:"comment" yaml:"comment"`
	// The data type of the `Column` .
	Type *string `json:"type" yaml:"type"`
}

// Specifies the sort order of a sorted column.
//
// TODO: EXAMPLE
//
type CfnTable_OrderProperty struct {
	// The name of the column.
	Column *string `json:"column" yaml:"column"`
	// Indicates that the column is sorted in ascending order ( `== 1` ), or in descending order ( `==0` ).
	SortOrder *float64 `json:"sortOrder" yaml:"sortOrder"`
}

// A structure that contains schema identity fields.
//
// Either this or the `SchemaVersionId` has to be
// provided.
//
// TODO: EXAMPLE
//
type CfnTable_SchemaIdProperty struct {
	// The name of the schema registry that contains the schema.
	RegistryName *string `json:"registryName" yaml:"registryName"`
	// The Amazon Resource Name (ARN) of the schema.
	//
	// One of `SchemaArn` or `SchemaName` has to be
	// provided.
	SchemaArn *string `json:"schemaArn" yaml:"schemaArn"`
	// The name of the schema.
	//
	// One of `SchemaArn` or `SchemaName` has to be provided.
	SchemaName *string `json:"schemaName" yaml:"schemaName"`
}

// An object that references a schema stored in the AWS Glue Schema Registry.
//
// TODO: EXAMPLE
//
type CfnTable_SchemaReferenceProperty struct {
	// A structure that contains schema identity fields.
	//
	// Either this or the `SchemaVersionId` has to be
	// provided.
	SchemaId interface{} `json:"schemaId" yaml:"schemaId"`
	// The unique ID assigned to a version of the schema.
	//
	// Either this or the `SchemaId` has to be provided.
	SchemaVersionId *string `json:"schemaVersionId" yaml:"schemaVersionId"`
	// The version number of the schema.
	SchemaVersionNumber *float64 `json:"schemaVersionNumber" yaml:"schemaVersionNumber"`
}

// Information about a serialization/deserialization program (SerDe) that serves as an extractor and loader.
//
// TODO: EXAMPLE
//
type CfnTable_SerdeInfoProperty struct {
	// Name of the SerDe.
	Name *string `json:"name" yaml:"name"`
	// These key-value pairs define initialization parameters for the SerDe.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// Usually the class that implements the SerDe.
	//
	// An example is `org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe` .
	SerializationLibrary *string `json:"serializationLibrary" yaml:"serializationLibrary"`
}

// Specifies skewed values in a table.
//
// Skewed values are those that occur with very high frequency.
//
// TODO: EXAMPLE
//
type CfnTable_SkewedInfoProperty struct {
	// A list of names of columns that contain skewed values.
	SkewedColumnNames *[]*string `json:"skewedColumnNames" yaml:"skewedColumnNames"`
	// A mapping of skewed values to the columns that contain them.
	SkewedColumnValueLocationMaps interface{} `json:"skewedColumnValueLocationMaps" yaml:"skewedColumnValueLocationMaps"`
	// A list of values that appear so frequently as to be considered skewed.
	SkewedColumnValues *[]*string `json:"skewedColumnValues" yaml:"skewedColumnValues"`
}

// Describes the physical storage of table data.
//
// TODO: EXAMPLE
//
type CfnTable_StorageDescriptorProperty struct {
	// A list of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns *[]*string `json:"bucketColumns" yaml:"bucketColumns"`
	// A list of the `Columns` in the table.
	Columns interface{} `json:"columns" yaml:"columns"`
	// `True` if the data in the table is compressed, or `False` if not.
	Compressed interface{} `json:"compressed" yaml:"compressed"`
	// The input format: `SequenceFileInputFormat` (binary), or `TextInputFormat` , or a custom format.
	InputFormat *string `json:"inputFormat" yaml:"inputFormat"`
	// The physical location of the table.
	//
	// By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location *string `json:"location" yaml:"location"`
	// Must be specified if the table contains any dimension columns.
	NumberOfBuckets *float64 `json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// The output format: `SequenceFileOutputFormat` (binary), or `IgnoreKeyTextOutputFormat` , or a custom format.
	OutputFormat *string `json:"outputFormat" yaml:"outputFormat"`
	// The user-supplied properties in key-value form.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// An object that references a schema stored in the AWS Glue Schema Registry.
	SchemaReference interface{} `json:"schemaReference" yaml:"schemaReference"`
	// The serialization/deserialization (SerDe) information.
	SerdeInfo interface{} `json:"serdeInfo" yaml:"serdeInfo"`
	// The information about values that appear frequently in a column (skewed values).
	SkewedInfo interface{} `json:"skewedInfo" yaml:"skewedInfo"`
	// A list specifying the sort order of each bucket in the table.
	SortColumns interface{} `json:"sortColumns" yaml:"sortColumns"`
	// `True` if the table data is stored in subdirectories, or `False` if not.
	StoredAsSubDirectories interface{} `json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}

// A structure that describes a target table for resource linking.
//
// TODO: EXAMPLE
//
type CfnTable_TableIdentifierProperty struct {
	// The ID of the Data Catalog in which the table resides.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database that contains the target table.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The name of the target table.
	Name *string `json:"name" yaml:"name"`
}

// A structure used to define a table.
//
// TODO: EXAMPLE
//
type CfnTable_TableInputProperty struct {
	// A description of the table.
	Description *string `json:"description" yaml:"description"`
	// The table name.
	//
	// For Hive compatibility, this is folded to lowercase when it is stored.
	Name *string `json:"name" yaml:"name"`
	// The table owner.
	Owner *string `json:"owner" yaml:"owner"`
	// These key-value pairs define properties associated with the table.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// A list of columns by which the table is partitioned. Only primitive types are supported as partition keys.
	//
	// When you create a table used by Amazon Athena, and you do not specify any `partitionKeys` , you must at least set the value of `partitionKeys` to an empty list. For example:
	//
	// `"PartitionKeys": []`
	PartitionKeys interface{} `json:"partitionKeys" yaml:"partitionKeys"`
	// The retention time for this table.
	Retention *float64 `json:"retention" yaml:"retention"`
	// A storage descriptor containing information about the physical storage of this table.
	StorageDescriptor interface{} `json:"storageDescriptor" yaml:"storageDescriptor"`
	// The type of this table ( `EXTERNAL_TABLE` , `VIRTUAL_VIEW` , etc.).
	TableType *string `json:"tableType" yaml:"tableType"`
	// A `TableIdentifier` structure that describes a target table for resource linking.
	TargetTable interface{} `json:"targetTable" yaml:"targetTable"`
	// If the table is a view, the expanded text of the view;
	//
	// otherwise `null` .
	ViewExpandedText *string `json:"viewExpandedText" yaml:"viewExpandedText"`
	// If the table is a view, the original text of the view;
	//
	// otherwise `null` .
	ViewOriginalText *string `json:"viewOriginalText" yaml:"viewOriginalText"`
}

// Properties for defining a `CfnTable`.
//
// TODO: EXAMPLE
//
type CfnTableProps struct {
	// The ID of the Data Catalog in which to create the `Table` .
	//
	// If none is supplied, the AWS account ID is used by default.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The name of the database where the table metadata resides.
	//
	// For Hive compatibility, this must be all lowercase.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// A structure used to define a table.
	TableInput interface{} `json:"tableInput" yaml:"tableInput"`
}

// A CloudFormation `AWS::Glue::Trigger`.
//
// The `AWS::Glue::Trigger` resource specifies triggers that run AWS Glue jobs. For more information, see [Triggering Jobs in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/trigger-job.html) and [Trigger Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-trigger.html#aws-glue-api-jobs-trigger-Trigger) in the *AWS Glue Developer Guide* .
//
// TODO: EXAMPLE
//
type CfnTrigger interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Actions() interface{}
	SetActions(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Predicate() interface{}
	SetPredicate(val interface{})
	Ref() *string
	Schedule() *string
	SetSchedule(val *string)
	Stack() awscdk.Stack
	StartOnCreation() interface{}
	SetStartOnCreation(val interface{})
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
	UpdatedProperites() *map[string]interface{}
	WorkflowName() *string
	SetWorkflowName(val *string)
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnTrigger
type jsiiProxy_CfnTrigger struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTrigger) Actions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Predicate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"predicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) StartOnCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrigger) WorkflowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowName",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Trigger`.
func NewCfnTrigger(scope constructs.Construct, id *string, props *CfnTriggerProps) CfnTrigger {
	_init_.Initialize()

	j := jsiiProxy_CfnTrigger{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnTrigger",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Trigger`.
func NewCfnTrigger_Override(c CfnTrigger, scope constructs.Construct, id *string, props *CfnTriggerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnTrigger",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTrigger) SetActions(val interface{}) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CfnTrigger) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnTrigger) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTrigger) SetPredicate(val interface{}) {
	_jsii_.Set(
		j,
		"predicate",
		val,
	)
}

func (j *jsiiProxy_CfnTrigger) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnTrigger) SetStartOnCreation(val interface{}) {
	_jsii_.Set(
		j,
		"startOnCreation",
		val,
	)
}

func (j *jsiiProxy_CfnTrigger) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnTrigger) SetWorkflowName(val *string) {
	_jsii_.Set(
		j,
		"workflowName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTrigger_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnTrigger",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTrigger_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnTrigger",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnTrigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnTrigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrigger_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnTrigger",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTrigger) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnTrigger) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnTrigger) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnTrigger) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTrigger) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTrigger) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnTrigger) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnTrigger) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnTrigger) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnTrigger) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTrigger) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTrigger) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnTrigger) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnTrigger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrigger) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Defines an action to be initiated by a trigger.
//
// TODO: EXAMPLE
//
type CfnTrigger_ActionProperty struct {
	// The job arguments used when this trigger fires.
	//
	// For this job run, they replace the default arguments set in the job definition itself.
	//
	// You can specify arguments here that your own job-execution script consumes, in addition to arguments that AWS Glue itself consumes.
	//
	// For information about how to specify and consume your own job arguments, see [Calling AWS Glue APIs in Python](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) in the *AWS Glue Developer Guide* .
	//
	// For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the developer guide.
	Arguments interface{} `json:"arguments" yaml:"arguments"`
	// The name of the crawler to be used with this action.
	CrawlerName *string `json:"crawlerName" yaml:"crawlerName"`
	// The name of a job to be executed.
	JobName *string `json:"jobName" yaml:"jobName"`
	// Specifies configuration properties of a job run notification.
	NotificationProperty interface{} `json:"notificationProperty" yaml:"notificationProperty"`
	// The name of the `SecurityConfiguration` structure to be used with this action.
	SecurityConfiguration *string `json:"securityConfiguration" yaml:"securityConfiguration"`
	// The `JobRun` timeout in minutes.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours). This overrides the timeout value set in the parent job.
	Timeout *float64 `json:"timeout" yaml:"timeout"`
}

// Defines a condition under which a trigger fires.
//
// TODO: EXAMPLE
//
type CfnTrigger_ConditionProperty struct {
	// The name of the crawler to which this condition applies.
	CrawlerName *string `json:"crawlerName" yaml:"crawlerName"`
	// The state of the crawler to which this condition applies.
	CrawlState *string `json:"crawlState" yaml:"crawlState"`
	// The name of the job whose `JobRuns` this condition applies to, and on which this trigger waits.
	JobName *string `json:"jobName" yaml:"jobName"`
	// A logical operator.
	LogicalOperator *string `json:"logicalOperator" yaml:"logicalOperator"`
	// The condition state.
	//
	// Currently, the values supported are `SUCCEEDED` , `STOPPED` , `TIMEOUT` , and `FAILED` .
	State *string `json:"state" yaml:"state"`
}

// Specifies configuration properties of a job run notification.
//
// TODO: EXAMPLE
//
type CfnTrigger_NotificationPropertyProperty struct {
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	NotifyDelayAfter *float64 `json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

// Defines the predicate of the trigger, which determines when it fires.
//
// TODO: EXAMPLE
//
type CfnTrigger_PredicateProperty struct {
	// A list of the conditions that determine when the trigger will fire.
	Conditions interface{} `json:"conditions" yaml:"conditions"`
	// An optional field if only one condition is listed.
	//
	// If multiple conditions are listed, then this field is required.
	Logical *string `json:"logical" yaml:"logical"`
}

// Properties for defining a `CfnTrigger`.
//
// TODO: EXAMPLE
//
type CfnTriggerProps struct {
	// The actions initiated by this trigger.
	Actions interface{} `json:"actions" yaml:"actions"`
	// The type of trigger that this is.
	Type *string `json:"type" yaml:"type"`
	// A description of this trigger.
	Description *string `json:"description" yaml:"description"`
	// The name of the trigger.
	Name *string `json:"name" yaml:"name"`
	// The predicate of this trigger, which defines when it will fire.
	Predicate interface{} `json:"predicate" yaml:"predicate"`
	// A `cron` expression used to specify the schedule.
	//
	// For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html) in the *AWS Glue Developer Guide* . For example, to run something every day at 12:15 UTC, specify `cron(15 12 * * ? *)` .
	Schedule *string `json:"schedule" yaml:"schedule"`
	// Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created.
	//
	// True is not supported for `ON_DEMAND` triggers.
	StartOnCreation interface{} `json:"startOnCreation" yaml:"startOnCreation"`
	// The tags to use with this trigger.
	Tags interface{} `json:"tags" yaml:"tags"`
	// The name of the workflow associated with the trigger.
	WorkflowName *string `json:"workflowName" yaml:"workflowName"`
}

// A CloudFormation `AWS::Glue::Workflow`.
//
// The `AWS::Glue::Workflow` is an AWS Glue resource type that manages AWS Glue workflows. A workflow is a container for a set of related jobs, crawlers, and triggers in AWS Glue . Using a workflow, you can design a complex multi-job extract, transform, and load (ETL) activity that AWS Glue can execute and track as single entity.
//
// TODO: EXAMPLE
//
type CfnWorkflow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DefaultRunProperties() interface{}
	SetDefaultRunProperties(val interface{})
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnWorkflow
type jsiiProxy_CfnWorkflow struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWorkflow) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) DefaultRunProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRunProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflow) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Workflow`.
func NewCfnWorkflow(scope constructs.Construct, id *string, props *CfnWorkflowProps) CfnWorkflow {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkflow{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnWorkflow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Workflow`.
func NewCfnWorkflow_Override(c CfnWorkflow, scope constructs.Construct, id *string, props *CfnWorkflowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnWorkflow",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWorkflow) SetDefaultRunProperties(val interface{}) {
	_jsii_.Set(
		j,
		"defaultRunProperties",
		val,
	)
}

func (j *jsiiProxy_CfnWorkflow) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnWorkflow) SetName(val *string) {
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
func CfnWorkflow_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnWorkflow",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWorkflow_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnWorkflow",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkflow_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnWorkflow",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnWorkflow) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnWorkflow) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnWorkflow) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnWorkflow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnWorkflow) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnWorkflow) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnWorkflow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnWorkflow) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnWorkflow) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnWorkflow) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnWorkflow) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWorkflow) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnWorkflow) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnWorkflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkflow) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnWorkflow`.
//
// TODO: EXAMPLE
//
type CfnWorkflowProps struct {
	// A collection of properties to be used as part of each execution of the workflow.
	DefaultRunProperties interface{} `json:"defaultRunProperties" yaml:"defaultRunProperties"`
	// A description of the workflow.
	Description *string `json:"description" yaml:"description"`
	// The name of the workflow representing the flow.
	Name *string `json:"name" yaml:"name"`
	// The tags to use with this workflow.
	Tags interface{} `json:"tags" yaml:"tags"`
}

