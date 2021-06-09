package awskinesisanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::KinesisAnalytics::Application`.
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationCode() *string
	SetApplicationCode(val *string)
	ApplicationDescription() *string
	SetApplicationDescription(val *string)
	ApplicationName() *string
	SetApplicationName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Inputs() interface{}
	SetInputs(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplication
type jsiiProxy_CfnApplication struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplication) ApplicationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) ApplicationDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Inputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisAnalytics::Application`.
func NewCfnApplication(scope awscdk.Construct, id *string, props *CfnApplicationProps) CfnApplication {
	_init_.Initialize()

	j := jsiiProxy_CfnApplication{}

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisAnalytics::Application`.
func NewCfnApplication_Override(c CfnApplication, scope awscdk.Construct, id *string, props *CfnApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplication",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplication) SetApplicationCode(val *string) {
	_jsii_.Set(
		j,
		"applicationCode",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetApplicationDescription(val *string) {
	_jsii_.Set(
		j,
		"applicationDescription",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetInputs(val interface{}) {
	_jsii_.Set(
		j,
		"inputs",
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
func CfnApplication_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplication",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApplication_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplication",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplication_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisanalytics.CfnApplication",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApplication) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnApplication) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnApplication) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApplication) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnApplication) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnApplication) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnApplication) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApplication) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplication) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplication) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplication) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnApplication) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplication) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApplication_CSVMappingParametersProperty struct {
	// `CfnApplication.CSVMappingParametersProperty.RecordColumnDelimiter`.
	RecordColumnDelimiter *string `json:"recordColumnDelimiter"`
	// `CfnApplication.CSVMappingParametersProperty.RecordRowDelimiter`.
	RecordRowDelimiter *string `json:"recordRowDelimiter"`
}

type CfnApplication_InputLambdaProcessorProperty struct {
	// `CfnApplication.InputLambdaProcessorProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
	// `CfnApplication.InputLambdaProcessorProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
}

type CfnApplication_InputParallelismProperty struct {
	// `CfnApplication.InputParallelismProperty.Count`.
	Count *float64 `json:"count"`
}

type CfnApplication_InputProcessingConfigurationProperty struct {
	// `CfnApplication.InputProcessingConfigurationProperty.InputLambdaProcessor`.
	InputLambdaProcessor interface{} `json:"inputLambdaProcessor"`
}

type CfnApplication_InputProperty struct {
	// `CfnApplication.InputProperty.InputSchema`.
	InputSchema interface{} `json:"inputSchema"`
	// `CfnApplication.InputProperty.NamePrefix`.
	NamePrefix *string `json:"namePrefix"`
	// `CfnApplication.InputProperty.InputParallelism`.
	InputParallelism interface{} `json:"inputParallelism"`
	// `CfnApplication.InputProperty.InputProcessingConfiguration`.
	InputProcessingConfiguration interface{} `json:"inputProcessingConfiguration"`
	// `CfnApplication.InputProperty.KinesisFirehoseInput`.
	KinesisFirehoseInput interface{} `json:"kinesisFirehoseInput"`
	// `CfnApplication.InputProperty.KinesisStreamsInput`.
	KinesisStreamsInput interface{} `json:"kinesisStreamsInput"`
}

type CfnApplication_InputSchemaProperty struct {
	// `CfnApplication.InputSchemaProperty.RecordColumns`.
	RecordColumns interface{} `json:"recordColumns"`
	// `CfnApplication.InputSchemaProperty.RecordFormat`.
	RecordFormat interface{} `json:"recordFormat"`
	// `CfnApplication.InputSchemaProperty.RecordEncoding`.
	RecordEncoding *string `json:"recordEncoding"`
}

type CfnApplication_JSONMappingParametersProperty struct {
	// `CfnApplication.JSONMappingParametersProperty.RecordRowPath`.
	RecordRowPath *string `json:"recordRowPath"`
}

type CfnApplication_KinesisFirehoseInputProperty struct {
	// `CfnApplication.KinesisFirehoseInputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
	// `CfnApplication.KinesisFirehoseInputProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
}

type CfnApplication_KinesisStreamsInputProperty struct {
	// `CfnApplication.KinesisStreamsInputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
	// `CfnApplication.KinesisStreamsInputProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
}

type CfnApplication_MappingParametersProperty struct {
	// `CfnApplication.MappingParametersProperty.CSVMappingParameters`.
	CsvMappingParameters interface{} `json:"csvMappingParameters"`
	// `CfnApplication.MappingParametersProperty.JSONMappingParameters`.
	JsonMappingParameters interface{} `json:"jsonMappingParameters"`
}

type CfnApplication_RecordColumnProperty struct {
	// `CfnApplication.RecordColumnProperty.Name`.
	Name *string `json:"name"`
	// `CfnApplication.RecordColumnProperty.SqlType`.
	SqlType *string `json:"sqlType"`
	// `CfnApplication.RecordColumnProperty.Mapping`.
	Mapping *string `json:"mapping"`
}

type CfnApplication_RecordFormatProperty struct {
	// `CfnApplication.RecordFormatProperty.RecordFormatType`.
	RecordFormatType *string `json:"recordFormatType"`
	// `CfnApplication.RecordFormatProperty.MappingParameters`.
	MappingParameters interface{} `json:"mappingParameters"`
}

// A CloudFormation `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
type CfnApplicationCloudWatchLoggingOptionV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationName() *string
	SetApplicationName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CloudWatchLoggingOption() interface{}
	SetCloudWatchLoggingOption(val interface{})
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplicationCloudWatchLoggingOptionV2
type jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) CloudWatchLoggingOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudWatchLoggingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
func NewCfnApplicationCloudWatchLoggingOptionV2(scope awscdk.Construct, id *string, props *CfnApplicationCloudWatchLoggingOptionV2Props) CfnApplicationCloudWatchLoggingOptionV2 {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2{}

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationCloudWatchLoggingOptionV2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
func NewCfnApplicationCloudWatchLoggingOptionV2_Override(c CfnApplicationCloudWatchLoggingOptionV2, scope awscdk.Construct, id *string, props *CfnApplicationCloudWatchLoggingOptionV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationCloudWatchLoggingOptionV2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) SetCloudWatchLoggingOption(val interface{}) {
	_jsii_.Set(
		j,
		"cloudWatchLoggingOption",
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
func CfnApplicationCloudWatchLoggingOptionV2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationCloudWatchLoggingOptionV2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApplicationCloudWatchLoggingOptionV2_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationCloudWatchLoggingOptionV2",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApplicationCloudWatchLoggingOptionV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationCloudWatchLoggingOptionV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationCloudWatchLoggingOptionV2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisanalytics.CfnApplicationCloudWatchLoggingOptionV2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApplicationCloudWatchLoggingOptionV2_CloudWatchLoggingOptionProperty struct {
	// `CfnApplicationCloudWatchLoggingOptionV2.CloudWatchLoggingOptionProperty.LogStreamARN`.
	LogStreamArn *string `json:"logStreamArn"`
}

// Properties for defining a `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
type CfnApplicationCloudWatchLoggingOptionV2Props struct {
	// `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
	// `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption.CloudWatchLoggingOption`.
	CloudWatchLoggingOption interface{} `json:"cloudWatchLoggingOption"`
}

// A CloudFormation `AWS::KinesisAnalytics::ApplicationOutput`.
type CfnApplicationOutput interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationName() *string
	SetApplicationName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Output() interface{}
	SetOutput(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplicationOutput
type jsiiProxy_CfnApplicationOutput struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplicationOutput) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) Output() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutput) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisAnalytics::ApplicationOutput`.
func NewCfnApplicationOutput(scope awscdk.Construct, id *string, props *CfnApplicationOutputProps) CfnApplicationOutput {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationOutput{}

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutput",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisAnalytics::ApplicationOutput`.
func NewCfnApplicationOutput_Override(c CfnApplicationOutput, scope awscdk.Construct, id *string, props *CfnApplicationOutputProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutput",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplicationOutput) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationOutput) SetOutput(val interface{}) {
	_jsii_.Set(
		j,
		"output",
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
func CfnApplicationOutput_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutput",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApplicationOutput_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutput",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApplicationOutput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationOutput_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutput",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApplicationOutput) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationOutput) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApplicationOutput) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApplicationOutput_DestinationSchemaProperty struct {
	// `CfnApplicationOutput.DestinationSchemaProperty.RecordFormatType`.
	RecordFormatType *string `json:"recordFormatType"`
}

type CfnApplicationOutput_KinesisFirehoseOutputProperty struct {
	// `CfnApplicationOutput.KinesisFirehoseOutputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
	// `CfnApplicationOutput.KinesisFirehoseOutputProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
}

type CfnApplicationOutput_KinesisStreamsOutputProperty struct {
	// `CfnApplicationOutput.KinesisStreamsOutputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
	// `CfnApplicationOutput.KinesisStreamsOutputProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
}

type CfnApplicationOutput_LambdaOutputProperty struct {
	// `CfnApplicationOutput.LambdaOutputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
	// `CfnApplicationOutput.LambdaOutputProperty.RoleARN`.
	RoleArn *string `json:"roleArn"`
}

type CfnApplicationOutput_OutputProperty struct {
	// `CfnApplicationOutput.OutputProperty.DestinationSchema`.
	DestinationSchema interface{} `json:"destinationSchema"`
	// `CfnApplicationOutput.OutputProperty.KinesisFirehoseOutput`.
	KinesisFirehoseOutput interface{} `json:"kinesisFirehoseOutput"`
	// `CfnApplicationOutput.OutputProperty.KinesisStreamsOutput`.
	KinesisStreamsOutput interface{} `json:"kinesisStreamsOutput"`
	// `CfnApplicationOutput.OutputProperty.LambdaOutput`.
	LambdaOutput interface{} `json:"lambdaOutput"`
	// `CfnApplicationOutput.OutputProperty.Name`.
	Name *string `json:"name"`
}

// Properties for defining a `AWS::KinesisAnalytics::ApplicationOutput`.
type CfnApplicationOutputProps struct {
	// `AWS::KinesisAnalytics::ApplicationOutput.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
	// `AWS::KinesisAnalytics::ApplicationOutput.Output`.
	Output interface{} `json:"output"`
}

// A CloudFormation `AWS::KinesisAnalyticsV2::ApplicationOutput`.
type CfnApplicationOutputV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationName() *string
	SetApplicationName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Output() interface{}
	SetOutput(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplicationOutputV2
type jsiiProxy_CfnApplicationOutputV2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplicationOutputV2) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) Output() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputV2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisAnalyticsV2::ApplicationOutput`.
func NewCfnApplicationOutputV2(scope awscdk.Construct, id *string, props *CfnApplicationOutputV2Props) CfnApplicationOutputV2 {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationOutputV2{}

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutputV2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisAnalyticsV2::ApplicationOutput`.
func NewCfnApplicationOutputV2_Override(c CfnApplicationOutputV2, scope awscdk.Construct, id *string, props *CfnApplicationOutputV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutputV2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplicationOutputV2) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationOutputV2) SetOutput(val interface{}) {
	_jsii_.Set(
		j,
		"output",
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
func CfnApplicationOutputV2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutputV2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApplicationOutputV2_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutputV2",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApplicationOutputV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutputV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationOutputV2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisanalytics.CfnApplicationOutputV2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApplicationOutputV2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApplicationOutputV2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApplicationOutputV2_DestinationSchemaProperty struct {
	// `CfnApplicationOutputV2.DestinationSchemaProperty.RecordFormatType`.
	RecordFormatType *string `json:"recordFormatType"`
}

type CfnApplicationOutputV2_KinesisFirehoseOutputProperty struct {
	// `CfnApplicationOutputV2.KinesisFirehoseOutputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
}

type CfnApplicationOutputV2_KinesisStreamsOutputProperty struct {
	// `CfnApplicationOutputV2.KinesisStreamsOutputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
}

type CfnApplicationOutputV2_LambdaOutputProperty struct {
	// `CfnApplicationOutputV2.LambdaOutputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
}

type CfnApplicationOutputV2_OutputProperty struct {
	// `CfnApplicationOutputV2.OutputProperty.DestinationSchema`.
	DestinationSchema interface{} `json:"destinationSchema"`
	// `CfnApplicationOutputV2.OutputProperty.KinesisFirehoseOutput`.
	KinesisFirehoseOutput interface{} `json:"kinesisFirehoseOutput"`
	// `CfnApplicationOutputV2.OutputProperty.KinesisStreamsOutput`.
	KinesisStreamsOutput interface{} `json:"kinesisStreamsOutput"`
	// `CfnApplicationOutputV2.OutputProperty.LambdaOutput`.
	LambdaOutput interface{} `json:"lambdaOutput"`
	// `CfnApplicationOutputV2.OutputProperty.Name`.
	Name *string `json:"name"`
}

// Properties for defining a `AWS::KinesisAnalyticsV2::ApplicationOutput`.
type CfnApplicationOutputV2Props struct {
	// `AWS::KinesisAnalyticsV2::ApplicationOutput.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
	// `AWS::KinesisAnalyticsV2::ApplicationOutput.Output`.
	Output interface{} `json:"output"`
}

// Properties for defining a `AWS::KinesisAnalytics::Application`.
type CfnApplicationProps struct {
	// `AWS::KinesisAnalytics::Application.Inputs`.
	Inputs interface{} `json:"inputs"`
	// `AWS::KinesisAnalytics::Application.ApplicationCode`.
	ApplicationCode *string `json:"applicationCode"`
	// `AWS::KinesisAnalytics::Application.ApplicationDescription`.
	ApplicationDescription *string `json:"applicationDescription"`
	// `AWS::KinesisAnalytics::Application.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
}

// A CloudFormation `AWS::KinesisAnalytics::ApplicationReferenceDataSource`.
type CfnApplicationReferenceDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationName() *string
	SetApplicationName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	ReferenceDataSource() interface{}
	SetReferenceDataSource(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplicationReferenceDataSource
type jsiiProxy_CfnApplicationReferenceDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) ReferenceDataSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referenceDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisAnalytics::ApplicationReferenceDataSource`.
func NewCfnApplicationReferenceDataSource(scope awscdk.Construct, id *string, props *CfnApplicationReferenceDataSourceProps) CfnApplicationReferenceDataSource {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationReferenceDataSource{}

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisAnalytics::ApplicationReferenceDataSource`.
func NewCfnApplicationReferenceDataSource_Override(c CfnApplicationReferenceDataSource, scope awscdk.Construct, id *string, props *CfnApplicationReferenceDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationReferenceDataSource) SetReferenceDataSource(val interface{}) {
	_jsii_.Set(
		j,
		"referenceDataSource",
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
func CfnApplicationReferenceDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApplicationReferenceDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSource",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApplicationReferenceDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationReferenceDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApplicationReferenceDataSource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApplicationReferenceDataSource_CSVMappingParametersProperty struct {
	// `CfnApplicationReferenceDataSource.CSVMappingParametersProperty.RecordColumnDelimiter`.
	RecordColumnDelimiter *string `json:"recordColumnDelimiter"`
	// `CfnApplicationReferenceDataSource.CSVMappingParametersProperty.RecordRowDelimiter`.
	RecordRowDelimiter *string `json:"recordRowDelimiter"`
}

type CfnApplicationReferenceDataSource_JSONMappingParametersProperty struct {
	// `CfnApplicationReferenceDataSource.JSONMappingParametersProperty.RecordRowPath`.
	RecordRowPath *string `json:"recordRowPath"`
}

type CfnApplicationReferenceDataSource_MappingParametersProperty struct {
	// `CfnApplicationReferenceDataSource.MappingParametersProperty.CSVMappingParameters`.
	CsvMappingParameters interface{} `json:"csvMappingParameters"`
	// `CfnApplicationReferenceDataSource.MappingParametersProperty.JSONMappingParameters`.
	JsonMappingParameters interface{} `json:"jsonMappingParameters"`
}

type CfnApplicationReferenceDataSource_RecordColumnProperty struct {
	// `CfnApplicationReferenceDataSource.RecordColumnProperty.Name`.
	Name *string `json:"name"`
	// `CfnApplicationReferenceDataSource.RecordColumnProperty.SqlType`.
	SqlType *string `json:"sqlType"`
	// `CfnApplicationReferenceDataSource.RecordColumnProperty.Mapping`.
	Mapping *string `json:"mapping"`
}

type CfnApplicationReferenceDataSource_RecordFormatProperty struct {
	// `CfnApplicationReferenceDataSource.RecordFormatProperty.RecordFormatType`.
	RecordFormatType *string `json:"recordFormatType"`
	// `CfnApplicationReferenceDataSource.RecordFormatProperty.MappingParameters`.
	MappingParameters interface{} `json:"mappingParameters"`
}

type CfnApplicationReferenceDataSource_ReferenceDataSourceProperty struct {
	// `CfnApplicationReferenceDataSource.ReferenceDataSourceProperty.ReferenceSchema`.
	ReferenceSchema interface{} `json:"referenceSchema"`
	// `CfnApplicationReferenceDataSource.ReferenceDataSourceProperty.S3ReferenceDataSource`.
	S3ReferenceDataSource interface{} `json:"s3ReferenceDataSource"`
	// `CfnApplicationReferenceDataSource.ReferenceDataSourceProperty.TableName`.
	TableName *string `json:"tableName"`
}

type CfnApplicationReferenceDataSource_ReferenceSchemaProperty struct {
	// `CfnApplicationReferenceDataSource.ReferenceSchemaProperty.RecordColumns`.
	RecordColumns interface{} `json:"recordColumns"`
	// `CfnApplicationReferenceDataSource.ReferenceSchemaProperty.RecordFormat`.
	RecordFormat interface{} `json:"recordFormat"`
	// `CfnApplicationReferenceDataSource.ReferenceSchemaProperty.RecordEncoding`.
	RecordEncoding *string `json:"recordEncoding"`
}

type CfnApplicationReferenceDataSource_S3ReferenceDataSourceProperty struct {
	// `CfnApplicationReferenceDataSource.S3ReferenceDataSourceProperty.BucketARN`.
	BucketArn *string `json:"bucketArn"`
	// `CfnApplicationReferenceDataSource.S3ReferenceDataSourceProperty.FileKey`.
	FileKey *string `json:"fileKey"`
	// `CfnApplicationReferenceDataSource.S3ReferenceDataSourceProperty.ReferenceRoleARN`.
	ReferenceRoleArn *string `json:"referenceRoleArn"`
}

// Properties for defining a `AWS::KinesisAnalytics::ApplicationReferenceDataSource`.
type CfnApplicationReferenceDataSourceProps struct {
	// `AWS::KinesisAnalytics::ApplicationReferenceDataSource.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
	// `AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource`.
	ReferenceDataSource interface{} `json:"referenceDataSource"`
}

// A CloudFormation `AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource`.
type CfnApplicationReferenceDataSourceV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationName() *string
	SetApplicationName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	ReferenceDataSource() interface{}
	SetReferenceDataSource(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplicationReferenceDataSourceV2
type jsiiProxy_CfnApplicationReferenceDataSourceV2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) ReferenceDataSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referenceDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource`.
func NewCfnApplicationReferenceDataSourceV2(scope awscdk.Construct, id *string, props *CfnApplicationReferenceDataSourceV2Props) CfnApplicationReferenceDataSourceV2 {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationReferenceDataSourceV2{}

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSourceV2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource`.
func NewCfnApplicationReferenceDataSourceV2_Override(c CfnApplicationReferenceDataSourceV2, scope awscdk.Construct, id *string, props *CfnApplicationReferenceDataSourceV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSourceV2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourceV2) SetReferenceDataSource(val interface{}) {
	_jsii_.Set(
		j,
		"referenceDataSource",
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
func CfnApplicationReferenceDataSourceV2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSourceV2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApplicationReferenceDataSourceV2_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSourceV2",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApplicationReferenceDataSourceV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSourceV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationReferenceDataSourceV2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisanalytics.CfnApplicationReferenceDataSourceV2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApplicationReferenceDataSourceV2_CSVMappingParametersProperty struct {
	// `CfnApplicationReferenceDataSourceV2.CSVMappingParametersProperty.RecordColumnDelimiter`.
	RecordColumnDelimiter *string `json:"recordColumnDelimiter"`
	// `CfnApplicationReferenceDataSourceV2.CSVMappingParametersProperty.RecordRowDelimiter`.
	RecordRowDelimiter *string `json:"recordRowDelimiter"`
}

type CfnApplicationReferenceDataSourceV2_JSONMappingParametersProperty struct {
	// `CfnApplicationReferenceDataSourceV2.JSONMappingParametersProperty.RecordRowPath`.
	RecordRowPath *string `json:"recordRowPath"`
}

type CfnApplicationReferenceDataSourceV2_MappingParametersProperty struct {
	// `CfnApplicationReferenceDataSourceV2.MappingParametersProperty.CSVMappingParameters`.
	CsvMappingParameters interface{} `json:"csvMappingParameters"`
	// `CfnApplicationReferenceDataSourceV2.MappingParametersProperty.JSONMappingParameters`.
	JsonMappingParameters interface{} `json:"jsonMappingParameters"`
}

type CfnApplicationReferenceDataSourceV2_RecordColumnProperty struct {
	// `CfnApplicationReferenceDataSourceV2.RecordColumnProperty.Name`.
	Name *string `json:"name"`
	// `CfnApplicationReferenceDataSourceV2.RecordColumnProperty.SqlType`.
	SqlType *string `json:"sqlType"`
	// `CfnApplicationReferenceDataSourceV2.RecordColumnProperty.Mapping`.
	Mapping *string `json:"mapping"`
}

type CfnApplicationReferenceDataSourceV2_RecordFormatProperty struct {
	// `CfnApplicationReferenceDataSourceV2.RecordFormatProperty.RecordFormatType`.
	RecordFormatType *string `json:"recordFormatType"`
	// `CfnApplicationReferenceDataSourceV2.RecordFormatProperty.MappingParameters`.
	MappingParameters interface{} `json:"mappingParameters"`
}

type CfnApplicationReferenceDataSourceV2_ReferenceDataSourceProperty struct {
	// `CfnApplicationReferenceDataSourceV2.ReferenceDataSourceProperty.ReferenceSchema`.
	ReferenceSchema interface{} `json:"referenceSchema"`
	// `CfnApplicationReferenceDataSourceV2.ReferenceDataSourceProperty.S3ReferenceDataSource`.
	S3ReferenceDataSource interface{} `json:"s3ReferenceDataSource"`
	// `CfnApplicationReferenceDataSourceV2.ReferenceDataSourceProperty.TableName`.
	TableName *string `json:"tableName"`
}

type CfnApplicationReferenceDataSourceV2_ReferenceSchemaProperty struct {
	// `CfnApplicationReferenceDataSourceV2.ReferenceSchemaProperty.RecordColumns`.
	RecordColumns interface{} `json:"recordColumns"`
	// `CfnApplicationReferenceDataSourceV2.ReferenceSchemaProperty.RecordFormat`.
	RecordFormat interface{} `json:"recordFormat"`
	// `CfnApplicationReferenceDataSourceV2.ReferenceSchemaProperty.RecordEncoding`.
	RecordEncoding *string `json:"recordEncoding"`
}

type CfnApplicationReferenceDataSourceV2_S3ReferenceDataSourceProperty struct {
	// `CfnApplicationReferenceDataSourceV2.S3ReferenceDataSourceProperty.BucketARN`.
	BucketArn *string `json:"bucketArn"`
	// `CfnApplicationReferenceDataSourceV2.S3ReferenceDataSourceProperty.FileKey`.
	FileKey *string `json:"fileKey"`
}

// Properties for defining a `AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource`.
type CfnApplicationReferenceDataSourceV2Props struct {
	// `AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
	// `AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource.ReferenceDataSource`.
	ReferenceDataSource interface{} `json:"referenceDataSource"`
}

// A CloudFormation `AWS::KinesisAnalyticsV2::Application`.
type CfnApplicationV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationConfiguration() interface{}
	SetApplicationConfiguration(val interface{})
	ApplicationDescription() *string
	SetApplicationDescription(val *string)
	ApplicationName() *string
	SetApplicationName(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	RuntimeEnvironment() *string
	SetRuntimeEnvironment(val *string)
	ServiceExecutionRole() *string
	SetServiceExecutionRole(val *string)
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplicationV2
type jsiiProxy_CfnApplicationV2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) RuntimeEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ServiceExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisAnalyticsV2::Application`.
func NewCfnApplicationV2(scope awscdk.Construct, id *string, props *CfnApplicationV2Props) CfnApplicationV2 {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationV2{}

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationV2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisAnalyticsV2::Application`.
func NewCfnApplicationV2_Override(c CfnApplicationV2, scope awscdk.Construct, id *string, props *CfnApplicationV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisanalytics.CfnApplicationV2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplicationV2) SetApplicationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"applicationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2) SetApplicationDescription(val *string) {
	_jsii_.Set(
		j,
		"applicationDescription",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2) SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2) SetRuntimeEnvironment(val *string) {
	_jsii_.Set(
		j,
		"runtimeEnvironment",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2) SetServiceExecutionRole(val *string) {
	_jsii_.Set(
		j,
		"serviceExecutionRole",
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
func CfnApplicationV2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationV2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApplicationV2_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationV2",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApplicationV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisanalytics.CfnApplicationV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationV2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisanalytics.CfnApplicationV2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApplicationV2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationV2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnApplicationV2) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApplicationV2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApplicationV2_ApplicationCodeConfigurationProperty struct {
	// `CfnApplicationV2.ApplicationCodeConfigurationProperty.CodeContent`.
	CodeContent interface{} `json:"codeContent"`
	// `CfnApplicationV2.ApplicationCodeConfigurationProperty.CodeContentType`.
	CodeContentType *string `json:"codeContentType"`
}

type CfnApplicationV2_ApplicationConfigurationProperty struct {
	// `CfnApplicationV2.ApplicationConfigurationProperty.ApplicationCodeConfiguration`.
	ApplicationCodeConfiguration interface{} `json:"applicationCodeConfiguration"`
	// `CfnApplicationV2.ApplicationConfigurationProperty.ApplicationSnapshotConfiguration`.
	ApplicationSnapshotConfiguration interface{} `json:"applicationSnapshotConfiguration"`
	// `CfnApplicationV2.ApplicationConfigurationProperty.EnvironmentProperties`.
	EnvironmentProperties interface{} `json:"environmentProperties"`
	// `CfnApplicationV2.ApplicationConfigurationProperty.FlinkApplicationConfiguration`.
	FlinkApplicationConfiguration interface{} `json:"flinkApplicationConfiguration"`
	// `CfnApplicationV2.ApplicationConfigurationProperty.SqlApplicationConfiguration`.
	SqlApplicationConfiguration interface{} `json:"sqlApplicationConfiguration"`
}

type CfnApplicationV2_ApplicationSnapshotConfigurationProperty struct {
	// `CfnApplicationV2.ApplicationSnapshotConfigurationProperty.SnapshotsEnabled`.
	SnapshotsEnabled interface{} `json:"snapshotsEnabled"`
}

type CfnApplicationV2_CSVMappingParametersProperty struct {
	// `CfnApplicationV2.CSVMappingParametersProperty.RecordColumnDelimiter`.
	RecordColumnDelimiter *string `json:"recordColumnDelimiter"`
	// `CfnApplicationV2.CSVMappingParametersProperty.RecordRowDelimiter`.
	RecordRowDelimiter *string `json:"recordRowDelimiter"`
}

type CfnApplicationV2_CheckpointConfigurationProperty struct {
	// `CfnApplicationV2.CheckpointConfigurationProperty.ConfigurationType`.
	ConfigurationType *string `json:"configurationType"`
	// `CfnApplicationV2.CheckpointConfigurationProperty.CheckpointingEnabled`.
	CheckpointingEnabled interface{} `json:"checkpointingEnabled"`
	// `CfnApplicationV2.CheckpointConfigurationProperty.CheckpointInterval`.
	CheckpointInterval *float64 `json:"checkpointInterval"`
	// `CfnApplicationV2.CheckpointConfigurationProperty.MinPauseBetweenCheckpoints`.
	MinPauseBetweenCheckpoints *float64 `json:"minPauseBetweenCheckpoints"`
}

type CfnApplicationV2_CodeContentProperty struct {
	// `CfnApplicationV2.CodeContentProperty.S3ContentLocation`.
	S3ContentLocation interface{} `json:"s3ContentLocation"`
	// `CfnApplicationV2.CodeContentProperty.TextContent`.
	TextContent *string `json:"textContent"`
	// `CfnApplicationV2.CodeContentProperty.ZipFileContent`.
	ZipFileContent *string `json:"zipFileContent"`
}

type CfnApplicationV2_EnvironmentPropertiesProperty struct {
	// `CfnApplicationV2.EnvironmentPropertiesProperty.PropertyGroups`.
	PropertyGroups interface{} `json:"propertyGroups"`
}

type CfnApplicationV2_FlinkApplicationConfigurationProperty struct {
	// `CfnApplicationV2.FlinkApplicationConfigurationProperty.CheckpointConfiguration`.
	CheckpointConfiguration interface{} `json:"checkpointConfiguration"`
	// `CfnApplicationV2.FlinkApplicationConfigurationProperty.MonitoringConfiguration`.
	MonitoringConfiguration interface{} `json:"monitoringConfiguration"`
	// `CfnApplicationV2.FlinkApplicationConfigurationProperty.ParallelismConfiguration`.
	ParallelismConfiguration interface{} `json:"parallelismConfiguration"`
}

type CfnApplicationV2_InputLambdaProcessorProperty struct {
	// `CfnApplicationV2.InputLambdaProcessorProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
}

type CfnApplicationV2_InputParallelismProperty struct {
	// `CfnApplicationV2.InputParallelismProperty.Count`.
	Count *float64 `json:"count"`
}

type CfnApplicationV2_InputProcessingConfigurationProperty struct {
	// `CfnApplicationV2.InputProcessingConfigurationProperty.InputLambdaProcessor`.
	InputLambdaProcessor interface{} `json:"inputLambdaProcessor"`
}

type CfnApplicationV2_InputProperty struct {
	// `CfnApplicationV2.InputProperty.InputSchema`.
	InputSchema interface{} `json:"inputSchema"`
	// `CfnApplicationV2.InputProperty.NamePrefix`.
	NamePrefix *string `json:"namePrefix"`
	// `CfnApplicationV2.InputProperty.InputParallelism`.
	InputParallelism interface{} `json:"inputParallelism"`
	// `CfnApplicationV2.InputProperty.InputProcessingConfiguration`.
	InputProcessingConfiguration interface{} `json:"inputProcessingConfiguration"`
	// `CfnApplicationV2.InputProperty.KinesisFirehoseInput`.
	KinesisFirehoseInput interface{} `json:"kinesisFirehoseInput"`
	// `CfnApplicationV2.InputProperty.KinesisStreamsInput`.
	KinesisStreamsInput interface{} `json:"kinesisStreamsInput"`
}

type CfnApplicationV2_InputSchemaProperty struct {
	// `CfnApplicationV2.InputSchemaProperty.RecordColumns`.
	RecordColumns interface{} `json:"recordColumns"`
	// `CfnApplicationV2.InputSchemaProperty.RecordFormat`.
	RecordFormat interface{} `json:"recordFormat"`
	// `CfnApplicationV2.InputSchemaProperty.RecordEncoding`.
	RecordEncoding *string `json:"recordEncoding"`
}

type CfnApplicationV2_JSONMappingParametersProperty struct {
	// `CfnApplicationV2.JSONMappingParametersProperty.RecordRowPath`.
	RecordRowPath *string `json:"recordRowPath"`
}

type CfnApplicationV2_KinesisFirehoseInputProperty struct {
	// `CfnApplicationV2.KinesisFirehoseInputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
}

type CfnApplicationV2_KinesisStreamsInputProperty struct {
	// `CfnApplicationV2.KinesisStreamsInputProperty.ResourceARN`.
	ResourceArn *string `json:"resourceArn"`
}

type CfnApplicationV2_MappingParametersProperty struct {
	// `CfnApplicationV2.MappingParametersProperty.CSVMappingParameters`.
	CsvMappingParameters interface{} `json:"csvMappingParameters"`
	// `CfnApplicationV2.MappingParametersProperty.JSONMappingParameters`.
	JsonMappingParameters interface{} `json:"jsonMappingParameters"`
}

type CfnApplicationV2_MonitoringConfigurationProperty struct {
	// `CfnApplicationV2.MonitoringConfigurationProperty.ConfigurationType`.
	ConfigurationType *string `json:"configurationType"`
	// `CfnApplicationV2.MonitoringConfigurationProperty.LogLevel`.
	LogLevel *string `json:"logLevel"`
	// `CfnApplicationV2.MonitoringConfigurationProperty.MetricsLevel`.
	MetricsLevel *string `json:"metricsLevel"`
}

type CfnApplicationV2_ParallelismConfigurationProperty struct {
	// `CfnApplicationV2.ParallelismConfigurationProperty.ConfigurationType`.
	ConfigurationType *string `json:"configurationType"`
	// `CfnApplicationV2.ParallelismConfigurationProperty.AutoScalingEnabled`.
	AutoScalingEnabled interface{} `json:"autoScalingEnabled"`
	// `CfnApplicationV2.ParallelismConfigurationProperty.Parallelism`.
	Parallelism *float64 `json:"parallelism"`
	// `CfnApplicationV2.ParallelismConfigurationProperty.ParallelismPerKPU`.
	ParallelismPerKpu *float64 `json:"parallelismPerKpu"`
}

type CfnApplicationV2_PropertyGroupProperty struct {
	// `CfnApplicationV2.PropertyGroupProperty.PropertyGroupId`.
	PropertyGroupId *string `json:"propertyGroupId"`
	// `CfnApplicationV2.PropertyGroupProperty.PropertyMap`.
	PropertyMap interface{} `json:"propertyMap"`
}

type CfnApplicationV2_RecordColumnProperty struct {
	// `CfnApplicationV2.RecordColumnProperty.Name`.
	Name *string `json:"name"`
	// `CfnApplicationV2.RecordColumnProperty.SqlType`.
	SqlType *string `json:"sqlType"`
	// `CfnApplicationV2.RecordColumnProperty.Mapping`.
	Mapping *string `json:"mapping"`
}

type CfnApplicationV2_RecordFormatProperty struct {
	// `CfnApplicationV2.RecordFormatProperty.RecordFormatType`.
	RecordFormatType *string `json:"recordFormatType"`
	// `CfnApplicationV2.RecordFormatProperty.MappingParameters`.
	MappingParameters interface{} `json:"mappingParameters"`
}

type CfnApplicationV2_S3ContentLocationProperty struct {
	// `CfnApplicationV2.S3ContentLocationProperty.BucketARN`.
	BucketArn *string `json:"bucketArn"`
	// `CfnApplicationV2.S3ContentLocationProperty.FileKey`.
	FileKey *string `json:"fileKey"`
	// `CfnApplicationV2.S3ContentLocationProperty.ObjectVersion`.
	ObjectVersion *string `json:"objectVersion"`
}

type CfnApplicationV2_SqlApplicationConfigurationProperty struct {
	// `CfnApplicationV2.SqlApplicationConfigurationProperty.Inputs`.
	Inputs interface{} `json:"inputs"`
}

// Properties for defining a `AWS::KinesisAnalyticsV2::Application`.
type CfnApplicationV2Props struct {
	// `AWS::KinesisAnalyticsV2::Application.RuntimeEnvironment`.
	RuntimeEnvironment *string `json:"runtimeEnvironment"`
	// `AWS::KinesisAnalyticsV2::Application.ServiceExecutionRole`.
	ServiceExecutionRole *string `json:"serviceExecutionRole"`
	// `AWS::KinesisAnalyticsV2::Application.ApplicationConfiguration`.
	ApplicationConfiguration interface{} `json:"applicationConfiguration"`
	// `AWS::KinesisAnalyticsV2::Application.ApplicationDescription`.
	ApplicationDescription *string `json:"applicationDescription"`
	// `AWS::KinesisAnalyticsV2::Application.ApplicationName`.
	ApplicationName *string `json:"applicationName"`
	// `AWS::KinesisAnalyticsV2::Application.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

