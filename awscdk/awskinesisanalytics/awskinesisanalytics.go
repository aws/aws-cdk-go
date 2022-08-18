package awskinesisanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::KinesisAnalytics::Application`.
//
// The `AWS::KinesisAnalytics::Application` resource creates an Amazon Kinesis Data Analytics application. For more information, see the [Amazon Kinesis Data Analytics Developer Guide](https://docs.aws.amazon.com//kinesisanalytics/latest/dev/what-is.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplication := awscdk.Aws_kinesisanalytics.NewCfnApplication(this, jsii.String("MyCfnApplication"), &cfnApplicationProps{
//   	inputs: []interface{}{
//   		&inputProperty{
//   			inputSchema: &inputSchemaProperty{
//   				recordColumns: []interface{}{
//   					&recordColumnProperty{
//   						name: jsii.String("name"),
//   						sqlType: jsii.String("sqlType"),
//
//   						// the properties below are optional
//   						mapping: jsii.String("mapping"),
//   					},
//   				},
//   				recordFormat: &recordFormatProperty{
//   					recordFormatType: jsii.String("recordFormatType"),
//
//   					// the properties below are optional
//   					mappingParameters: &mappingParametersProperty{
//   						csvMappingParameters: &cSVMappingParametersProperty{
//   							recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   							recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   						},
//   						jsonMappingParameters: &jSONMappingParametersProperty{
//   							recordRowPath: jsii.String("recordRowPath"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				recordEncoding: jsii.String("recordEncoding"),
//   			},
//   			namePrefix: jsii.String("namePrefix"),
//
//   			// the properties below are optional
//   			inputParallelism: &inputParallelismProperty{
//   				count: jsii.Number(123),
//   			},
//   			inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   				inputLambdaProcessor: &inputLambdaProcessorProperty{
//   					resourceArn: jsii.String("resourceArn"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   			},
//   			kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			kinesisStreamsInput: &kinesisStreamsInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	applicationCode: jsii.String("applicationCode"),
//   	applicationDescription: jsii.String("applicationDescription"),
//   	applicationName: jsii.String("applicationName"),
//   })
//
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// One or more SQL statements that read input data, transform it, and generate output.
	//
	// For example, you can write a SQL statement that reads data from one in-application stream, generates a running average of the number of advertisement clicks by vendor, and insert resulting rows in another in-application stream using pumps. For more information about the typical pattern, see [Application Code](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-app-code.html) .
	//
	// You can provide such series of SQL statements, where output of one statement can be used as the input for the next statement. You store intermediate results by creating in-application streams and pumps.
	//
	// Note that the application code must create the streams with names specified in the `Outputs` . For example, if your `Outputs` defines output streams named `ExampleOutputStream1` and `ExampleOutputStream2` , then your application code must create these streams.
	ApplicationCode() *string
	SetApplicationCode(val *string)
	// Summary description of the application.
	ApplicationDescription() *string
	SetApplicationDescription(val *string)
	// Name of your Amazon Kinesis Analytics application (for example, `sample-app` ).
	ApplicationName() *string
	SetApplicationName(val *string)
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
	// Use this parameter to configure the application input.
	//
	// You can configure your application to receive input from a single streaming source. In this configuration, you map this streaming source to an in-application stream that is created. Your application code can then query the in-application stream like a table (you can think of it as a constantly updating table).
	//
	// For the streaming source, you provide its Amazon Resource Name (ARN) and format of data on the stream (for example, JSON, CSV, etc.). You also must provide an IAM role that Amazon Kinesis Analytics can assume to read this stream on your behalf.
	//
	// To create the in-application stream, you need to specify a schema to transform your data into a schematized version used in SQL. In the schema, you provide the necessary mapping of the data elements in the streaming source to record columns in the in-app stream.
	Inputs() interface{}
	SetInputs(val interface{})
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

func (c *jsiiProxy_CfnApplication) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplication) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnApplication) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Provides additional mapping information when the record format uses delimiters, such as CSV.
//
// For example, the following sample records use CSV format, where the records use the *'\n'* as the row delimiter and a comma (",") as the column delimiter:
//
// `"name1", "address1"`
//
// `"name2", "address2"`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cSVMappingParametersProperty := &cSVMappingParametersProperty{
//   	recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   	recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   }
//
type CfnApplication_CSVMappingParametersProperty struct {
	// Column delimiter.
	//
	// For example, in a CSV format, a comma (",") is the typical column delimiter.
	RecordColumnDelimiter *string `field:"required" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// Row delimiter.
	//
	// For example, in a CSV format, *'\n'* is the typical row delimiter.
	RecordRowDelimiter *string `field:"required" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

// An object that contains the Amazon Resource Name (ARN) of the [AWS Lambda](https://docs.aws.amazon.com/lambda/) function that is used to preprocess records in the stream, and the ARN of the IAM role that is used to access the AWS Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLambdaProcessorProperty := &inputLambdaProcessorProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplication_InputLambdaProcessorProperty struct {
	// The ARN of the [AWS Lambda](https://docs.aws.amazon.com/lambda/) function that operates on records in the stream.
	//
	// > To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: AWS Lambda](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-lambda)
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The ARN of the IAM role that is used to access the AWS Lambda function.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

// Describes the number of in-application streams to create for a given streaming source.
//
// For information about parallelism, see [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputParallelismProperty := &inputParallelismProperty{
//   	count: jsii.Number(123),
//   }
//
type CfnApplication_InputParallelismProperty struct {
	// Number of in-application streams to create.
	//
	// For more information, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html) .
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

// Provides a description of a processor that is used to preprocess the records in the stream before being processed by your application code.
//
// Currently, the only input processor available is [AWS Lambda](https://docs.aws.amazon.com/lambda/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProcessingConfigurationProperty := &inputProcessingConfigurationProperty{
//   	inputLambdaProcessor: &inputLambdaProcessorProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   }
//
type CfnApplication_InputProcessingConfigurationProperty struct {
	// The [InputLambdaProcessor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputlambdaprocessor.html) that is used to preprocess the records in the stream before being processed by your application code.
	InputLambdaProcessor interface{} `field:"optional" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}

// When you configure the application input, you specify the streaming source, the in-application stream name that is created, and the mapping between the two.
//
// For more information, see [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &inputProperty{
//   	inputSchema: &inputSchemaProperty{
//   		recordColumns: []interface{}{
//   			&recordColumnProperty{
//   				name: jsii.String("name"),
//   				sqlType: jsii.String("sqlType"),
//
//   				// the properties below are optional
//   				mapping: jsii.String("mapping"),
//   			},
//   		},
//   		recordFormat: &recordFormatProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//
//   			// the properties below are optional
//   			mappingParameters: &mappingParametersProperty{
//   				csvMappingParameters: &cSVMappingParametersProperty{
//   					recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				jsonMappingParameters: &jSONMappingParametersProperty{
//   					recordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		recordEncoding: jsii.String("recordEncoding"),
//   	},
//   	namePrefix: jsii.String("namePrefix"),
//
//   	// the properties below are optional
//   	inputParallelism: &inputParallelismProperty{
//   		count: jsii.Number(123),
//   	},
//   	inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   		inputLambdaProcessor: &inputLambdaProcessorProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	kinesisStreamsInput: &kinesisStreamsInputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   }
//
type CfnApplication_InputProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
	//
	// Also used to describe the format of the reference data source.
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// Name prefix to use when creating an in-application stream.
	//
	// Suppose that you specify a prefix "MyInApplicationStream." Amazon Kinesis Analytics then creates one or more (as per the `InputParallelism` count you specified) in-application streams with names "MyInApplicationStream_001," "MyInApplicationStream_002," and so on.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// Describes the number of in-application streams to create.
	//
	// Data from your source is routed to these in-application input streams.
	//
	// See [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html) .
	InputParallelism interface{} `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// The [InputProcessingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html) for the input. An input processor transforms records as they are received from the stream, before the application's SQL code executes. Currently, the only input processing configuration available is [InputLambdaProcessor](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputlambdaprocessor.html) .
	InputProcessingConfiguration interface{} `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// If the streaming source is an Amazon Kinesis Firehose delivery stream, identifies the delivery stream's ARN and an IAM role that enables Amazon Kinesis Analytics to access the stream on your behalf.
	//
	// Note: Either `KinesisStreamsInput` or `KinesisFirehoseInput` is required.
	KinesisFirehoseInput interface{} `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// If the streaming source is an Amazon Kinesis stream, identifies the stream's Amazon Resource Name (ARN) and an IAM role that enables Amazon Kinesis Analytics to access the stream on your behalf.
	//
	// Note: Either `KinesisStreamsInput` or `KinesisFirehoseInput` is required.
	KinesisStreamsInput interface{} `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
//
// Also used to describe the format of the reference data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSchemaProperty := &inputSchemaProperty{
//   	recordColumns: []interface{}{
//   		&recordColumnProperty{
//   			name: jsii.String("name"),
//   			sqlType: jsii.String("sqlType"),
//
//   			// the properties below are optional
//   			mapping: jsii.String("mapping"),
//   		},
//   	},
//   	recordFormat: &recordFormatProperty{
//   		recordFormatType: jsii.String("recordFormatType"),
//
//   		// the properties below are optional
//   		mappingParameters: &mappingParametersProperty{
//   			csvMappingParameters: &cSVMappingParametersProperty{
//   				recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   				recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   			},
//   			jsonMappingParameters: &jSONMappingParametersProperty{
//   				recordRowPath: jsii.String("recordRowPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	recordEncoding: jsii.String("recordEncoding"),
//   }
//
type CfnApplication_InputSchemaProperty struct {
	// A list of `RecordColumn` objects.
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the format of the records on the streaming source.
	RecordFormat interface{} `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Specifies the encoding of the records in the streaming source.
	//
	// For example, UTF-8.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

// Provides additional mapping information when JSON is the record format on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jSONMappingParametersProperty := &jSONMappingParametersProperty{
//   	recordRowPath: jsii.String("recordRowPath"),
//   }
//
type CfnApplication_JSONMappingParametersProperty struct {
	// Path to the top-level parent that contains the records.
	RecordRowPath *string `field:"required" json:"recordRowPath" yaml:"recordRowPath"`
}

// Identifies an Amazon Kinesis Firehose delivery stream as the streaming source.
//
// You provide the delivery stream's Amazon Resource Name (ARN) and an IAM role ARN that enables Amazon Kinesis Analytics to access the stream on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseInputProperty := &kinesisFirehoseInputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplication_KinesisFirehoseInputProperty struct {
	// ARN of the input delivery stream.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the stream on your behalf.
	//
	// You need to make sure that the role has the necessary permissions to access the stream.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

// Identifies an Amazon Kinesis stream as the streaming source.
//
// You provide the stream's Amazon Resource Name (ARN) and an IAM role ARN that enables Amazon Kinesis Analytics to access the stream on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamsInputProperty := &kinesisStreamsInputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplication_KinesisStreamsInputProperty struct {
	// ARN of the input Amazon Kinesis stream to read.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the stream on your behalf.
	//
	// You need to grant the necessary permissions to this role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

// When configuring application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingParametersProperty := &mappingParametersProperty{
//   	csvMappingParameters: &cSVMappingParametersProperty{
//   		recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   		recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   	},
//   	jsonMappingParameters: &jSONMappingParametersProperty{
//   		recordRowPath: jsii.String("recordRowPath"),
//   	},
//   }
//
type CfnApplication_MappingParametersProperty struct {
	// Provides additional mapping information when the record format uses delimiters (for example, CSV).
	CsvMappingParameters interface{} `field:"optional" json:"csvMappingParameters" yaml:"csvMappingParameters"`
	// Provides additional mapping information when JSON is the record format on the streaming source.
	JsonMappingParameters interface{} `field:"optional" json:"jsonMappingParameters" yaml:"jsonMappingParameters"`
}

// Describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
//
// Also used to describe the format of the reference data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordColumnProperty := &recordColumnProperty{
//   	name: jsii.String("name"),
//   	sqlType: jsii.String("sqlType"),
//
//   	// the properties below are optional
//   	mapping: jsii.String("mapping"),
//   }
//
type CfnApplication_RecordColumnProperty struct {
	// Name of the column created in the in-application input stream or reference table.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of column created in the in-application input stream or reference table.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Reference to the data element in the streaming input or the reference data source.
	//
	// This element is required if the [RecordFormatType](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_RecordFormat.html#analytics-Type-RecordFormat-RecordFormatTypel) is `JSON` .
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

// Describes the record format and relevant mapping information that should be applied to schematize the records on the stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordFormatProperty := &recordFormatProperty{
//   	recordFormatType: jsii.String("recordFormatType"),
//
//   	// the properties below are optional
//   	mappingParameters: &mappingParametersProperty{
//   		csvMappingParameters: &cSVMappingParametersProperty{
//   			recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   			recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   		},
//   		jsonMappingParameters: &jSONMappingParametersProperty{
//   			recordRowPath: jsii.String("recordRowPath"),
//   		},
//   	},
//   }
//
type CfnApplication_RecordFormatProperty struct {
	// The type of record format.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// When configuring application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	MappingParameters interface{} `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

// A CloudFormation `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
//
// Adds an Amazon CloudWatch log stream to monitor application configuration errors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationCloudWatchLoggingOptionV2 := awscdk.Aws_kinesisanalytics.NewCfnApplicationCloudWatchLoggingOptionV2(this, jsii.String("MyCfnApplicationCloudWatchLoggingOptionV2"), &cfnApplicationCloudWatchLoggingOptionV2Props{
//   	applicationName: jsii.String("applicationName"),
//   	cloudWatchLoggingOption: &cloudWatchLoggingOptionProperty{
//   		logStreamArn: jsii.String("logStreamArn"),
//   	},
//   })
//
type CfnApplicationCloudWatchLoggingOptionV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the application.
	ApplicationName() *string
	SetApplicationName(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
	CloudWatchLoggingOption() interface{}
	SetCloudWatchLoggingOption(val interface{})
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

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionV2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLoggingOptionProperty := &cloudWatchLoggingOptionProperty{
//   	logStreamArn: jsii.String("logStreamArn"),
//   }
//
type CfnApplicationCloudWatchLoggingOptionV2_CloudWatchLoggingOptionProperty struct {
	// The ARN of the CloudWatch log to receive application messages.
	LogStreamArn *string `field:"required" json:"logStreamArn" yaml:"logStreamArn"`
}

// Properties for defining a `CfnApplicationCloudWatchLoggingOptionV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationCloudWatchLoggingOptionV2Props := &cfnApplicationCloudWatchLoggingOptionV2Props{
//   	applicationName: jsii.String("applicationName"),
//   	cloudWatchLoggingOption: &cloudWatchLoggingOptionProperty{
//   		logStreamArn: jsii.String("logStreamArn"),
//   	},
//   }
//
type CfnApplicationCloudWatchLoggingOptionV2Props struct {
	// The name of the application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
	CloudWatchLoggingOption interface{} `field:"required" json:"cloudWatchLoggingOption" yaml:"cloudWatchLoggingOption"`
}

// A CloudFormation `AWS::KinesisAnalytics::ApplicationOutput`.
//
// Adds an external destination to your Amazon Kinesis Analytics application.
//
// If you want Amazon Kinesis Analytics to deliver data from an in-application stream within your application to an external destination (such as an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an Amazon Lambda function), you add the relevant configuration to your application using this operation. You can configure one or more outputs for your application. Each output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your in-application error stream to an external destination so that you can analyze the errors. For more information, see [Understanding Application Output (Destination)](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html) .
//
// Any configuration update, including adding a streaming source using this operation, results in a new version of the application. You can use the `DescribeApplication` operation to find the current application version.
//
// For the limits on the number of application inputs and outputs you can configure, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html) .
//
// This operation requires permissions to perform the `kinesisanalytics:AddApplicationOutput` action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationOutput := awscdk.Aws_kinesisanalytics.NewCfnApplicationOutput(this, jsii.String("MyCfnApplicationOutput"), &cfnApplicationOutputProps{
//   	applicationName: jsii.String("applicationName"),
//   	output: &outputProperty{
//   		destinationSchema: &destinationSchemaProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//   		},
//
//   		// the properties below are optional
//   		kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		lambdaOutput: &lambdaOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		name: jsii.String("name"),
//   	},
//   })
//
type CfnApplicationOutput interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Name of the application to which you want to add the output configuration.
	ApplicationName() *string
	SetApplicationName(val *string)
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
	// An array of objects, each describing one output configuration.
	//
	// In the output configuration, you specify the name of an in-application stream, a destination (that is, an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda function), and record the formation to use when writing to the destination.
	Output() interface{}
	SetOutput(val interface{})
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

func (c *jsiiProxy_CfnApplicationOutput) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplicationOutput) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationOutput) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplicationOutput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplicationOutput) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplicationOutput) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplicationOutput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnApplicationOutput) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplicationOutput) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationOutput) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationOutput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnApplicationOutput) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationOutput) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes the data format when records are written to the destination.
//
// For more information, see [Configuring Application Output](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationSchemaProperty := &destinationSchemaProperty{
//   	recordFormatType: jsii.String("recordFormatType"),
//   }
//
type CfnApplicationOutput_DestinationSchemaProperty struct {
	// Specifies the format of the records on the output stream.
	RecordFormatType *string `field:"optional" json:"recordFormatType" yaml:"recordFormatType"`
}

// When configuring application output, identifies an Amazon Kinesis Firehose delivery stream as the destination.
//
// You provide the stream Amazon Resource Name (ARN) and an IAM role that enables Amazon Kinesis Analytics to write to the stream on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseOutputProperty := &kinesisFirehoseOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplicationOutput_KinesisFirehoseOutputProperty struct {
	// ARN of the destination Amazon Kinesis Firehose delivery stream to write to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the destination stream on your behalf.
	//
	// You need to grant the necessary permissions to this role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

// When configuring application output, identifies an Amazon Kinesis stream as the destination.
//
// You provide the stream Amazon Resource Name (ARN) and also an IAM role ARN that Amazon Kinesis Analytics can use to write to the stream on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamsOutputProperty := &kinesisStreamsOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplicationOutput_KinesisStreamsOutputProperty struct {
	// ARN of the destination Amazon Kinesis stream to write to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the destination stream on your behalf.
	//
	// You need to grant the necessary permissions to this role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

// When configuring application output, identifies an AWS Lambda function as the destination.
//
// You provide the function Amazon Resource Name (ARN) and also an IAM role ARN that Amazon Kinesis Analytics can use to write to the function on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaOutputProperty := &lambdaOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnApplicationOutput_LambdaOutputProperty struct {
	// Amazon Resource Name (ARN) of the destination Lambda function to write to.
	//
	// > To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: AWS Lambda](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-lambda)
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the destination function on your behalf.
	//
	// You need to grant the necessary permissions to this role.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

// Describes application output configuration in which you identify an in-application stream and a destination where you want the in-application stream data to be written.
//
// The destination can be an Amazon Kinesis stream or an Amazon Kinesis Firehose delivery stream.
//
// For limits on how many destinations an application can write and other limitations, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputProperty := &outputProperty{
//   	destinationSchema: &destinationSchemaProperty{
//   		recordFormatType: jsii.String("recordFormatType"),
//   	},
//
//   	// the properties below are optional
//   	kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	lambdaOutput: &lambdaOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnApplicationOutput_OutputProperty struct {
	// Describes the data format when records are written to the destination.
	//
	// For more information, see [Configuring Application Output](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html) .
	DestinationSchema interface{} `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Identifies an Amazon Kinesis Firehose delivery stream as the destination.
	KinesisFirehoseOutput interface{} `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// Identifies an Amazon Kinesis stream as the destination.
	KinesisStreamsOutput interface{} `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// Identifies an AWS Lambda function as the destination.
	LambdaOutput interface{} `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
	// Name of the in-application stream.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Properties for defining a `CfnApplicationOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationOutputProps := &cfnApplicationOutputProps{
//   	applicationName: jsii.String("applicationName"),
//   	output: &outputProperty{
//   		destinationSchema: &destinationSchemaProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//   		},
//
//   		// the properties below are optional
//   		kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		lambdaOutput: &lambdaOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnApplicationOutputProps struct {
	// Name of the application to which you want to add the output configuration.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// An array of objects, each describing one output configuration.
	//
	// In the output configuration, you specify the name of an in-application stream, a destination (that is, an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda function), and record the formation to use when writing to the destination.
	Output interface{} `field:"required" json:"output" yaml:"output"`
}

// A CloudFormation `AWS::KinesisAnalyticsV2::ApplicationOutput`.
//
// Adds an external destination to your SQL-based Amazon Kinesis Data Analytics application.
//
// If you want Kinesis Data Analytics to deliver data from an in-application stream within your application to an external destination (such as an Kinesis data stream, a Kinesis Data Firehose delivery stream, or an Amazon Lambda function), you add the relevant configuration to your application using this operation. You can configure one or more outputs for your application. Each output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your in-application error stream to an external destination so that you can analyze the errors.
//
// Any configuration update, including adding a streaming source using this operation, results in a new version of the application. You can use the [DescribeApplication](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_DescribeApplication.html) operation to find the current application version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationOutputV2 := awscdk.Aws_kinesisanalytics.NewCfnApplicationOutputV2(this, jsii.String("MyCfnApplicationOutputV2"), &cfnApplicationOutputV2Props{
//   	applicationName: jsii.String("applicationName"),
//   	output: &outputProperty{
//   		destinationSchema: &destinationSchemaProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//   		},
//
//   		// the properties below are optional
//   		kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		lambdaOutput: &lambdaOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		name: jsii.String("name"),
//   	},
//   })
//
type CfnApplicationOutputV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the application.
	ApplicationName() *string
	SetApplicationName(val *string)
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
	// Describes a SQL-based Kinesis Data Analytics application's output configuration, in which you identify an in-application stream and a destination where you want the in-application stream data to be written.
	//
	// The destination can be a Kinesis data stream or a Kinesis Data Firehose delivery stream.
	Output() interface{}
	SetOutput(val interface{})
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

func (c *jsiiProxy_CfnApplicationOutputV2) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnApplicationOutputV2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationOutputV2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationOutputV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnApplicationOutputV2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationOutputV2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes the data format when records are written to the destination in a SQL-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationSchemaProperty := &destinationSchemaProperty{
//   	recordFormatType: jsii.String("recordFormatType"),
//   }
//
type CfnApplicationOutputV2_DestinationSchemaProperty struct {
	// Specifies the format of the records on the output stream.
	RecordFormatType *string `field:"optional" json:"recordFormatType" yaml:"recordFormatType"`
}

// For a SQL-based Kinesis Data Analytics application, when configuring application output, identifies a Kinesis Data Firehose delivery stream as the destination.
//
// You provide the stream Amazon Resource Name (ARN) of the delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseOutputProperty := &kinesisFirehoseOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationOutputV2_KinesisFirehoseOutputProperty struct {
	// The ARN of the destination delivery stream to write to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

// When you configure a SQL-based Kinesis Data Analytics application's output, identifies a Kinesis data stream as the destination.
//
// You provide the stream Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamsOutputProperty := &kinesisStreamsOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationOutputV2_KinesisStreamsOutputProperty struct {
	// The ARN of the destination Kinesis data stream to write to.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

// When you configure a SQL-based Kinesis Data Analytics application's output, identifies an Amazon Lambda function as the destination.
//
// You provide the function Amazon Resource Name (ARN) of the Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaOutputProperty := &lambdaOutputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationOutputV2_LambdaOutputProperty struct {
	// The Amazon Resource Name (ARN) of the destination Lambda function to write to.
	//
	// > To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: Amazon Lambda](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-lambda)
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

// Describes a SQL-based Kinesis Data Analytics application's output configuration, in which you identify an in-application stream and a destination where you want the in-application stream data to be written.
//
// The destination can be a Kinesis data stream or a Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputProperty := &outputProperty{
//   	destinationSchema: &destinationSchemaProperty{
//   		recordFormatType: jsii.String("recordFormatType"),
//   	},
//
//   	// the properties below are optional
//   	kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   	kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   	lambdaOutput: &lambdaOutputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnApplicationOutputV2_OutputProperty struct {
	// Describes the data format when records are written to the destination.
	DestinationSchema interface{} `field:"required" json:"destinationSchema" yaml:"destinationSchema"`
	// Identifies a Kinesis Data Firehose delivery stream as the destination.
	KinesisFirehoseOutput interface{} `field:"optional" json:"kinesisFirehoseOutput" yaml:"kinesisFirehoseOutput"`
	// Identifies a Kinesis data stream as the destination.
	KinesisStreamsOutput interface{} `field:"optional" json:"kinesisStreamsOutput" yaml:"kinesisStreamsOutput"`
	// Identifies an Amazon Lambda function as the destination.
	LambdaOutput interface{} `field:"optional" json:"lambdaOutput" yaml:"lambdaOutput"`
	// The name of the in-application stream.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

// Properties for defining a `CfnApplicationOutputV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationOutputV2Props := &cfnApplicationOutputV2Props{
//   	applicationName: jsii.String("applicationName"),
//   	output: &outputProperty{
//   		destinationSchema: &destinationSchemaProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//   		},
//
//   		// the properties below are optional
//   		kinesisFirehoseOutput: &kinesisFirehoseOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		kinesisStreamsOutput: &kinesisStreamsOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		lambdaOutput: &lambdaOutputProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnApplicationOutputV2Props struct {
	// The name of the application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Describes a SQL-based Kinesis Data Analytics application's output configuration, in which you identify an in-application stream and a destination where you want the in-application stream data to be written.
	//
	// The destination can be a Kinesis data stream or a Kinesis Data Firehose delivery stream.
	Output interface{} `field:"required" json:"output" yaml:"output"`
}

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	inputs: []interface{}{
//   		&inputProperty{
//   			inputSchema: &inputSchemaProperty{
//   				recordColumns: []interface{}{
//   					&recordColumnProperty{
//   						name: jsii.String("name"),
//   						sqlType: jsii.String("sqlType"),
//
//   						// the properties below are optional
//   						mapping: jsii.String("mapping"),
//   					},
//   				},
//   				recordFormat: &recordFormatProperty{
//   					recordFormatType: jsii.String("recordFormatType"),
//
//   					// the properties below are optional
//   					mappingParameters: &mappingParametersProperty{
//   						csvMappingParameters: &cSVMappingParametersProperty{
//   							recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   							recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   						},
//   						jsonMappingParameters: &jSONMappingParametersProperty{
//   							recordRowPath: jsii.String("recordRowPath"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				recordEncoding: jsii.String("recordEncoding"),
//   			},
//   			namePrefix: jsii.String("namePrefix"),
//
//   			// the properties below are optional
//   			inputParallelism: &inputParallelismProperty{
//   				count: jsii.Number(123),
//   			},
//   			inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   				inputLambdaProcessor: &inputLambdaProcessorProperty{
//   					resourceArn: jsii.String("resourceArn"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   			},
//   			kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   			kinesisStreamsInput: &kinesisStreamsInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   				roleArn: jsii.String("roleArn"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	applicationCode: jsii.String("applicationCode"),
//   	applicationDescription: jsii.String("applicationDescription"),
//   	applicationName: jsii.String("applicationName"),
//   }
//
type CfnApplicationProps struct {
	// Use this parameter to configure the application input.
	//
	// You can configure your application to receive input from a single streaming source. In this configuration, you map this streaming source to an in-application stream that is created. Your application code can then query the in-application stream like a table (you can think of it as a constantly updating table).
	//
	// For the streaming source, you provide its Amazon Resource Name (ARN) and format of data on the stream (for example, JSON, CSV, etc.). You also must provide an IAM role that Amazon Kinesis Analytics can assume to read this stream on your behalf.
	//
	// To create the in-application stream, you need to specify a schema to transform your data into a schematized version used in SQL. In the schema, you provide the necessary mapping of the data elements in the streaming source to record columns in the in-app stream.
	Inputs interface{} `field:"required" json:"inputs" yaml:"inputs"`
	// One or more SQL statements that read input data, transform it, and generate output.
	//
	// For example, you can write a SQL statement that reads data from one in-application stream, generates a running average of the number of advertisement clicks by vendor, and insert resulting rows in another in-application stream using pumps. For more information about the typical pattern, see [Application Code](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-app-code.html) .
	//
	// You can provide such series of SQL statements, where output of one statement can be used as the input for the next statement. You store intermediate results by creating in-application streams and pumps.
	//
	// Note that the application code must create the streams with names specified in the `Outputs` . For example, if your `Outputs` defines output streams named `ExampleOutputStream1` and `ExampleOutputStream2` , then your application code must create these streams.
	ApplicationCode *string `field:"optional" json:"applicationCode" yaml:"applicationCode"`
	// Summary description of the application.
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// Name of your Amazon Kinesis Analytics application (for example, `sample-app` ).
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

// A CloudFormation `AWS::KinesisAnalytics::ApplicationReferenceDataSource`.
//
// Adds a reference data source to an existing application.
//
// Amazon Kinesis Analytics reads reference data (that is, an Amazon S3 object) and creates an in-application table within your application. In the request, you provide the source (S3 bucket name and object key name), name of the in-application table to create, and the necessary mapping information that describes how data in Amazon S3 object maps to columns in the resulting in-application table.
//
// For conceptual information, see [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html) . For the limits on data sources you can add to your application, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html) .
//
// This operation requires permissions to perform the `kinesisanalytics:AddApplicationOutput` action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationReferenceDataSource := awscdk.Aws_kinesisanalytics.NewCfnApplicationReferenceDataSource(this, jsii.String("MyCfnApplicationReferenceDataSource"), &cfnApplicationReferenceDataSourceProps{
//   	applicationName: jsii.String("applicationName"),
//   	referenceDataSource: &referenceDataSourceProperty{
//   		referenceSchema: &referenceSchemaProperty{
//   			recordColumns: []interface{}{
//   				&recordColumnProperty{
//   					name: jsii.String("name"),
//   					sqlType: jsii.String("sqlType"),
//
//   					// the properties below are optional
//   					mapping: jsii.String("mapping"),
//   				},
//   			},
//   			recordFormat: &recordFormatProperty{
//   				recordFormatType: jsii.String("recordFormatType"),
//
//   				// the properties below are optional
//   				mappingParameters: &mappingParametersProperty{
//   					csvMappingParameters: &cSVMappingParametersProperty{
//   						recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   						recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   					},
//   					jsonMappingParameters: &jSONMappingParametersProperty{
//   						recordRowPath: jsii.String("recordRowPath"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			recordEncoding: jsii.String("recordEncoding"),
//   		},
//
//   		// the properties below are optional
//   		s3ReferenceDataSource: &s3ReferenceDataSourceProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			fileKey: jsii.String("fileKey"),
//   			referenceRoleArn: jsii.String("referenceRoleArn"),
//   		},
//   		tableName: jsii.String("tableName"),
//   	},
//   })
//
type CfnApplicationReferenceDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Name of an existing application.
	ApplicationName() *string
	SetApplicationName(val *string)
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
	// The reference data source can be an object in your Amazon S3 bucket.
	//
	// Amazon Kinesis Analytics reads the object and copies the data into the in-application table that is created. You provide an S3 bucket, object key name, and the resulting in-application table that is created. You must also provide an IAM role with the necessary permissions that Amazon Kinesis Analytics can assume to read the object from your S3 bucket on your behalf.
	ReferenceDataSource() interface{}
	SetReferenceDataSource(val interface{})
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

func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnApplicationReferenceDataSource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationReferenceDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnApplicationReferenceDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationReferenceDataSource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Provides additional mapping information when the record format uses delimiters, such as CSV.
//
// For example, the following sample records use CSV format, where the records use the *'\n'* as the row delimiter and a comma (",") as the column delimiter:
//
// `"name1", "address1"`
//
// `"name2", "address2"`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cSVMappingParametersProperty := &cSVMappingParametersProperty{
//   	recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   	recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   }
//
type CfnApplicationReferenceDataSource_CSVMappingParametersProperty struct {
	// Column delimiter.
	//
	// For example, in a CSV format, a comma (",") is the typical column delimiter.
	RecordColumnDelimiter *string `field:"required" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// Row delimiter.
	//
	// For example, in a CSV format, *'\n'* is the typical row delimiter.
	RecordRowDelimiter *string `field:"required" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

// Provides additional mapping information when JSON is the record format on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jSONMappingParametersProperty := &jSONMappingParametersProperty{
//   	recordRowPath: jsii.String("recordRowPath"),
//   }
//
type CfnApplicationReferenceDataSource_JSONMappingParametersProperty struct {
	// Path to the top-level parent that contains the records.
	RecordRowPath *string `field:"required" json:"recordRowPath" yaml:"recordRowPath"`
}

// When configuring application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingParametersProperty := &mappingParametersProperty{
//   	csvMappingParameters: &cSVMappingParametersProperty{
//   		recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   		recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   	},
//   	jsonMappingParameters: &jSONMappingParametersProperty{
//   		recordRowPath: jsii.String("recordRowPath"),
//   	},
//   }
//
type CfnApplicationReferenceDataSource_MappingParametersProperty struct {
	// Provides additional mapping information when the record format uses delimiters (for example, CSV).
	CsvMappingParameters interface{} `field:"optional" json:"csvMappingParameters" yaml:"csvMappingParameters"`
	// Provides additional mapping information when JSON is the record format on the streaming source.
	JsonMappingParameters interface{} `field:"optional" json:"jsonMappingParameters" yaml:"jsonMappingParameters"`
}

// Describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
//
// Also used to describe the format of the reference data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordColumnProperty := &recordColumnProperty{
//   	name: jsii.String("name"),
//   	sqlType: jsii.String("sqlType"),
//
//   	// the properties below are optional
//   	mapping: jsii.String("mapping"),
//   }
//
type CfnApplicationReferenceDataSource_RecordColumnProperty struct {
	// Name of the column created in the in-application input stream or reference table.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of column created in the in-application input stream or reference table.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Reference to the data element in the streaming input or the reference data source.
	//
	// This element is required if the [RecordFormatType](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_RecordFormat.html#analytics-Type-RecordFormat-RecordFormatTypel) is `JSON` .
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

// Describes the record format and relevant mapping information that should be applied to schematize the records on the stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordFormatProperty := &recordFormatProperty{
//   	recordFormatType: jsii.String("recordFormatType"),
//
//   	// the properties below are optional
//   	mappingParameters: &mappingParametersProperty{
//   		csvMappingParameters: &cSVMappingParametersProperty{
//   			recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   			recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   		},
//   		jsonMappingParameters: &jSONMappingParametersProperty{
//   			recordRowPath: jsii.String("recordRowPath"),
//   		},
//   	},
//   }
//
type CfnApplicationReferenceDataSource_RecordFormatProperty struct {
	// The type of record format.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// When configuring application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	MappingParameters interface{} `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

// Describes the reference data source by providing the source information (S3 bucket name and object key name), the resulting in-application table name that is created, and the necessary schema to map the data elements in the Amazon S3 object to the in-application table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceDataSourceProperty := &referenceDataSourceProperty{
//   	referenceSchema: &referenceSchemaProperty{
//   		recordColumns: []interface{}{
//   			&recordColumnProperty{
//   				name: jsii.String("name"),
//   				sqlType: jsii.String("sqlType"),
//
//   				// the properties below are optional
//   				mapping: jsii.String("mapping"),
//   			},
//   		},
//   		recordFormat: &recordFormatProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//
//   			// the properties below are optional
//   			mappingParameters: &mappingParametersProperty{
//   				csvMappingParameters: &cSVMappingParametersProperty{
//   					recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				jsonMappingParameters: &jSONMappingParametersProperty{
//   					recordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		recordEncoding: jsii.String("recordEncoding"),
//   	},
//
//   	// the properties below are optional
//   	s3ReferenceDataSource: &s3ReferenceDataSourceProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		fileKey: jsii.String("fileKey"),
//   		referenceRoleArn: jsii.String("referenceRoleArn"),
//   	},
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnApplicationReferenceDataSource_ReferenceDataSourceProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
	ReferenceSchema interface{} `field:"required" json:"referenceSchema" yaml:"referenceSchema"`
	// Identifies the S3 bucket and object that contains the reference data.
	//
	// Also identifies the IAM role Amazon Kinesis Analytics can assume to read this object on your behalf. An Amazon Kinesis Analytics application loads reference data only once. If the data changes, you call the `UpdateApplication` operation to trigger reloading of data into your application.
	S3ReferenceDataSource interface{} `field:"optional" json:"s3ReferenceDataSource" yaml:"s3ReferenceDataSource"`
	// Name of the in-application table to create.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

// The ReferenceSchema property type specifies the format of the data in the reference source for a SQL-based Amazon Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceSchemaProperty := &referenceSchemaProperty{
//   	recordColumns: []interface{}{
//   		&recordColumnProperty{
//   			name: jsii.String("name"),
//   			sqlType: jsii.String("sqlType"),
//
//   			// the properties below are optional
//   			mapping: jsii.String("mapping"),
//   		},
//   	},
//   	recordFormat: &recordFormatProperty{
//   		recordFormatType: jsii.String("recordFormatType"),
//
//   		// the properties below are optional
//   		mappingParameters: &mappingParametersProperty{
//   			csvMappingParameters: &cSVMappingParametersProperty{
//   				recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   				recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   			},
//   			jsonMappingParameters: &jSONMappingParametersProperty{
//   				recordRowPath: jsii.String("recordRowPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	recordEncoding: jsii.String("recordEncoding"),
//   }
//
type CfnApplicationReferenceDataSource_ReferenceSchemaProperty struct {
	// A list of RecordColumn objects.
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the format of the records on the reference source.
	RecordFormat interface{} `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Specifies the encoding of the records in the reference source.
	//
	// For example, UTF-8.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

// Identifies the S3 bucket and object that contains the reference data.
//
// Also identifies the IAM role Amazon Kinesis Analytics can assume to read this object on your behalf.
//
// An Amazon Kinesis Analytics application loads reference data only once. If the data changes, you call the [UpdateApplication](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_UpdateApplication.html) operation to trigger reloading of data into your application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ReferenceDataSourceProperty := &s3ReferenceDataSourceProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	fileKey: jsii.String("fileKey"),
//   	referenceRoleArn: jsii.String("referenceRoleArn"),
//   }
//
type CfnApplicationReferenceDataSource_S3ReferenceDataSourceProperty struct {
	// Amazon Resource Name (ARN) of the S3 bucket.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Object key name containing reference data.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// ARN of the IAM role that the service can assume to read data on your behalf.
	//
	// This role must have permission for the `s3:GetObject` action on the object and trust policy that allows Amazon Kinesis Analytics service principal to assume this role.
	ReferenceRoleArn *string `field:"required" json:"referenceRoleArn" yaml:"referenceRoleArn"`
}

// Properties for defining a `CfnApplicationReferenceDataSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationReferenceDataSourceProps := &cfnApplicationReferenceDataSourceProps{
//   	applicationName: jsii.String("applicationName"),
//   	referenceDataSource: &referenceDataSourceProperty{
//   		referenceSchema: &referenceSchemaProperty{
//   			recordColumns: []interface{}{
//   				&recordColumnProperty{
//   					name: jsii.String("name"),
//   					sqlType: jsii.String("sqlType"),
//
//   					// the properties below are optional
//   					mapping: jsii.String("mapping"),
//   				},
//   			},
//   			recordFormat: &recordFormatProperty{
//   				recordFormatType: jsii.String("recordFormatType"),
//
//   				// the properties below are optional
//   				mappingParameters: &mappingParametersProperty{
//   					csvMappingParameters: &cSVMappingParametersProperty{
//   						recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   						recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   					},
//   					jsonMappingParameters: &jSONMappingParametersProperty{
//   						recordRowPath: jsii.String("recordRowPath"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			recordEncoding: jsii.String("recordEncoding"),
//   		},
//
//   		// the properties below are optional
//   		s3ReferenceDataSource: &s3ReferenceDataSourceProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			fileKey: jsii.String("fileKey"),
//   			referenceRoleArn: jsii.String("referenceRoleArn"),
//   		},
//   		tableName: jsii.String("tableName"),
//   	},
//   }
//
type CfnApplicationReferenceDataSourceProps struct {
	// Name of an existing application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The reference data source can be an object in your Amazon S3 bucket.
	//
	// Amazon Kinesis Analytics reads the object and copies the data into the in-application table that is created. You provide an S3 bucket, object key name, and the resulting in-application table that is created. You must also provide an IAM role with the necessary permissions that Amazon Kinesis Analytics can assume to read the object from your S3 bucket on your behalf.
	ReferenceDataSource interface{} `field:"required" json:"referenceDataSource" yaml:"referenceDataSource"`
}

// A CloudFormation `AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource`.
//
// Adds a reference data source to an existing SQL-based Kinesis Data Analytics application.
//
// Kinesis Data Analytics reads reference data (that is, an Amazon S3 object) and creates an in-application table within your application. In the request, you provide the source (S3 bucket name and object key name), name of the in-application table to create, and the necessary mapping information that describes how data in an Amazon S3 object maps to columns in the resulting in-application table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationReferenceDataSourceV2 := awscdk.Aws_kinesisanalytics.NewCfnApplicationReferenceDataSourceV2(this, jsii.String("MyCfnApplicationReferenceDataSourceV2"), &cfnApplicationReferenceDataSourceV2Props{
//   	applicationName: jsii.String("applicationName"),
//   	referenceDataSource: &referenceDataSourceProperty{
//   		referenceSchema: &referenceSchemaProperty{
//   			recordColumns: []interface{}{
//   				&recordColumnProperty{
//   					name: jsii.String("name"),
//   					sqlType: jsii.String("sqlType"),
//
//   					// the properties below are optional
//   					mapping: jsii.String("mapping"),
//   				},
//   			},
//   			recordFormat: &recordFormatProperty{
//   				recordFormatType: jsii.String("recordFormatType"),
//
//   				// the properties below are optional
//   				mappingParameters: &mappingParametersProperty{
//   					csvMappingParameters: &cSVMappingParametersProperty{
//   						recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   						recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   					},
//   					jsonMappingParameters: &jSONMappingParametersProperty{
//   						recordRowPath: jsii.String("recordRowPath"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			recordEncoding: jsii.String("recordEncoding"),
//   		},
//
//   		// the properties below are optional
//   		s3ReferenceDataSource: &s3ReferenceDataSourceProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			fileKey: jsii.String("fileKey"),
//   		},
//   		tableName: jsii.String("tableName"),
//   	},
//   })
//
type CfnApplicationReferenceDataSourceV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the application.
	ApplicationName() *string
	SetApplicationName(val *string)
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
	// For a SQL-based Kinesis Data Analytics application, describes the reference data source by providing the source information (Amazon S3 bucket name and object key name), the resulting in-application table name that is created, and the necessary schema to map the data elements in the Amazon S3 object to the in-application table.
	ReferenceDataSource() interface{}
	SetReferenceDataSource(val interface{})
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

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationReferenceDataSourceV2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// For a SQL-based Kinesis Data Analytics application, provides additional mapping information when the record format uses delimiters, such as CSV.
//
// For example, the following sample records use CSV format, where the records use the *'\n'* as the row delimiter and a comma (",") as the column delimiter:
//
// `"name1", "address1"`
//
// `"name2", "address2"`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cSVMappingParametersProperty := &cSVMappingParametersProperty{
//   	recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   	recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   }
//
type CfnApplicationReferenceDataSourceV2_CSVMappingParametersProperty struct {
	// The column delimiter.
	//
	// For example, in a CSV format, a comma (",") is the typical column delimiter.
	RecordColumnDelimiter *string `field:"required" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// The row delimiter.
	//
	// For example, in a CSV format, *'\n'* is the typical row delimiter.
	RecordRowDelimiter *string `field:"required" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

// For a SQL-based Kinesis Data Analytics application, provides additional mapping information when JSON is the record format on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jSONMappingParametersProperty := &jSONMappingParametersProperty{
//   	recordRowPath: jsii.String("recordRowPath"),
//   }
//
type CfnApplicationReferenceDataSourceV2_JSONMappingParametersProperty struct {
	// The path to the top-level parent that contains the records.
	RecordRowPath *string `field:"required" json:"recordRowPath" yaml:"recordRowPath"`
}

// When you configure a SQL-based Kinesis Data Analytics application's input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingParametersProperty := &mappingParametersProperty{
//   	csvMappingParameters: &cSVMappingParametersProperty{
//   		recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   		recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   	},
//   	jsonMappingParameters: &jSONMappingParametersProperty{
//   		recordRowPath: jsii.String("recordRowPath"),
//   	},
//   }
//
type CfnApplicationReferenceDataSourceV2_MappingParametersProperty struct {
	// Provides additional mapping information when the record format uses delimiters (for example, CSV).
	CsvMappingParameters interface{} `field:"optional" json:"csvMappingParameters" yaml:"csvMappingParameters"`
	// Provides additional mapping information when JSON is the record format on the streaming source.
	JsonMappingParameters interface{} `field:"optional" json:"jsonMappingParameters" yaml:"jsonMappingParameters"`
}

// For a SQL-based Kinesis Data Analytics application, describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
//
// Also used to describe the format of the reference data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordColumnProperty := &recordColumnProperty{
//   	name: jsii.String("name"),
//   	sqlType: jsii.String("sqlType"),
//
//   	// the properties below are optional
//   	mapping: jsii.String("mapping"),
//   }
//
type CfnApplicationReferenceDataSourceV2_RecordColumnProperty struct {
	// The name of the column that is created in the in-application input stream or reference table.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of column created in the in-application input stream or reference table.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// A reference to the data element in the streaming input or the reference data source.
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

// For a SQL-based Kinesis Data Analytics application, describes the record format and relevant mapping information that should be applied to schematize the records on the stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordFormatProperty := &recordFormatProperty{
//   	recordFormatType: jsii.String("recordFormatType"),
//
//   	// the properties below are optional
//   	mappingParameters: &mappingParametersProperty{
//   		csvMappingParameters: &cSVMappingParametersProperty{
//   			recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   			recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   		},
//   		jsonMappingParameters: &jSONMappingParametersProperty{
//   			recordRowPath: jsii.String("recordRowPath"),
//   		},
//   	},
//   }
//
type CfnApplicationReferenceDataSourceV2_RecordFormatProperty struct {
	// The type of record format.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// When you configure application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	MappingParameters interface{} `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

// For a SQL-based Kinesis Data Analytics application, describes the reference data source by providing the source information (Amazon S3 bucket name and object key name), the resulting in-application table name that is created, and the necessary schema to map the data elements in the Amazon S3 object to the in-application table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceDataSourceProperty := &referenceDataSourceProperty{
//   	referenceSchema: &referenceSchemaProperty{
//   		recordColumns: []interface{}{
//   			&recordColumnProperty{
//   				name: jsii.String("name"),
//   				sqlType: jsii.String("sqlType"),
//
//   				// the properties below are optional
//   				mapping: jsii.String("mapping"),
//   			},
//   		},
//   		recordFormat: &recordFormatProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//
//   			// the properties below are optional
//   			mappingParameters: &mappingParametersProperty{
//   				csvMappingParameters: &cSVMappingParametersProperty{
//   					recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				jsonMappingParameters: &jSONMappingParametersProperty{
//   					recordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		recordEncoding: jsii.String("recordEncoding"),
//   	},
//
//   	// the properties below are optional
//   	s3ReferenceDataSource: &s3ReferenceDataSourceProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		fileKey: jsii.String("fileKey"),
//   	},
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnApplicationReferenceDataSourceV2_ReferenceDataSourceProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
	ReferenceSchema interface{} `field:"required" json:"referenceSchema" yaml:"referenceSchema"`
	// Identifies the S3 bucket and object that contains the reference data.
	//
	// A Kinesis Data Analytics application loads reference data only once. If the data changes, you call the [UpdateApplication](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_UpdateApplication.html) operation to trigger reloading of data into your application.
	S3ReferenceDataSource interface{} `field:"optional" json:"s3ReferenceDataSource" yaml:"s3ReferenceDataSource"`
	// The name of the in-application table to create.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

// For a SQL-based Kinesis Data Analytics application, describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceSchemaProperty := &referenceSchemaProperty{
//   	recordColumns: []interface{}{
//   		&recordColumnProperty{
//   			name: jsii.String("name"),
//   			sqlType: jsii.String("sqlType"),
//
//   			// the properties below are optional
//   			mapping: jsii.String("mapping"),
//   		},
//   	},
//   	recordFormat: &recordFormatProperty{
//   		recordFormatType: jsii.String("recordFormatType"),
//
//   		// the properties below are optional
//   		mappingParameters: &mappingParametersProperty{
//   			csvMappingParameters: &cSVMappingParametersProperty{
//   				recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   				recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   			},
//   			jsonMappingParameters: &jSONMappingParametersProperty{
//   				recordRowPath: jsii.String("recordRowPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	recordEncoding: jsii.String("recordEncoding"),
//   }
//
type CfnApplicationReferenceDataSourceV2_ReferenceSchemaProperty struct {
	// A list of `RecordColumn` objects.
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the format of the records on the streaming source.
	RecordFormat interface{} `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Specifies the encoding of the records in the streaming source.
	//
	// For example, UTF-8.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

// For an SQL-based Amazon Kinesis Data Analytics application, identifies the Amazon S3 bucket and object that contains the reference data.
//
// A Kinesis Data Analytics application loads reference data only once. If the data changes, you call the [UpdateApplication](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_UpdateApplication.html) operation to trigger reloading of data into your application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ReferenceDataSourceProperty := &s3ReferenceDataSourceProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	fileKey: jsii.String("fileKey"),
//   }
//
type CfnApplicationReferenceDataSourceV2_S3ReferenceDataSourceProperty struct {
	// The Amazon Resource Name (ARN) of the S3 bucket.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The object key name containing the reference data.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
}

// Properties for defining a `CfnApplicationReferenceDataSourceV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationReferenceDataSourceV2Props := &cfnApplicationReferenceDataSourceV2Props{
//   	applicationName: jsii.String("applicationName"),
//   	referenceDataSource: &referenceDataSourceProperty{
//   		referenceSchema: &referenceSchemaProperty{
//   			recordColumns: []interface{}{
//   				&recordColumnProperty{
//   					name: jsii.String("name"),
//   					sqlType: jsii.String("sqlType"),
//
//   					// the properties below are optional
//   					mapping: jsii.String("mapping"),
//   				},
//   			},
//   			recordFormat: &recordFormatProperty{
//   				recordFormatType: jsii.String("recordFormatType"),
//
//   				// the properties below are optional
//   				mappingParameters: &mappingParametersProperty{
//   					csvMappingParameters: &cSVMappingParametersProperty{
//   						recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   						recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   					},
//   					jsonMappingParameters: &jSONMappingParametersProperty{
//   						recordRowPath: jsii.String("recordRowPath"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			recordEncoding: jsii.String("recordEncoding"),
//   		},
//
//   		// the properties below are optional
//   		s3ReferenceDataSource: &s3ReferenceDataSourceProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			fileKey: jsii.String("fileKey"),
//   		},
//   		tableName: jsii.String("tableName"),
//   	},
//   }
//
type CfnApplicationReferenceDataSourceV2Props struct {
	// The name of the application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// For a SQL-based Kinesis Data Analytics application, describes the reference data source by providing the source information (Amazon S3 bucket name and object key name), the resulting in-application table name that is created, and the necessary schema to map the data elements in the Amazon S3 object to the in-application table.
	ReferenceDataSource interface{} `field:"required" json:"referenceDataSource" yaml:"referenceDataSource"`
}

// A CloudFormation `AWS::KinesisAnalyticsV2::Application`.
//
// Creates an Amazon Kinesis Data Analytics application. For information about creating a Kinesis Data Analytics application, see [Creating an Application](https://docs.aws.amazon.com/kinesisanalytics/latest/java/getting-started.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationV2 := awscdk.Aws_kinesisanalytics.NewCfnApplicationV2(this, jsii.String("MyCfnApplicationV2"), &cfnApplicationV2Props{
//   	runtimeEnvironment: jsii.String("runtimeEnvironment"),
//   	serviceExecutionRole: jsii.String("serviceExecutionRole"),
//
//   	// the properties below are optional
//   	applicationConfiguration: &applicationConfigurationProperty{
//   		applicationCodeConfiguration: &applicationCodeConfigurationProperty{
//   			codeContent: &codeContentProperty{
//   				s3ContentLocation: &s3ContentLocationProperty{
//   					bucketArn: jsii.String("bucketArn"),
//   					fileKey: jsii.String("fileKey"),
//
//   					// the properties below are optional
//   					objectVersion: jsii.String("objectVersion"),
//   				},
//   				textContent: jsii.String("textContent"),
//   				zipFileContent: jsii.String("zipFileContent"),
//   			},
//   			codeContentType: jsii.String("codeContentType"),
//   		},
//   		applicationSnapshotConfiguration: &applicationSnapshotConfigurationProperty{
//   			snapshotsEnabled: jsii.Boolean(false),
//   		},
//   		environmentProperties: &environmentPropertiesProperty{
//   			propertyGroups: []interface{}{
//   				&propertyGroupProperty{
//   					propertyGroupId: jsii.String("propertyGroupId"),
//   					propertyMap: map[string]*string{
//   						"propertyMapKey": jsii.String("propertyMap"),
//   					},
//   				},
//   			},
//   		},
//   		flinkApplicationConfiguration: &flinkApplicationConfigurationProperty{
//   			checkpointConfiguration: &checkpointConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				checkpointingEnabled: jsii.Boolean(false),
//   				checkpointInterval: jsii.Number(123),
//   				minPauseBetweenCheckpoints: jsii.Number(123),
//   			},
//   			monitoringConfiguration: &monitoringConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				logLevel: jsii.String("logLevel"),
//   				metricsLevel: jsii.String("metricsLevel"),
//   			},
//   			parallelismConfiguration: &parallelismConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				autoScalingEnabled: jsii.Boolean(false),
//   				parallelism: jsii.Number(123),
//   				parallelismPerKpu: jsii.Number(123),
//   			},
//   		},
//   		sqlApplicationConfiguration: &sqlApplicationConfigurationProperty{
//   			inputs: []interface{}{
//   				&inputProperty{
//   					inputSchema: &inputSchemaProperty{
//   						recordColumns: []interface{}{
//   							&recordColumnProperty{
//   								name: jsii.String("name"),
//   								sqlType: jsii.String("sqlType"),
//
//   								// the properties below are optional
//   								mapping: jsii.String("mapping"),
//   							},
//   						},
//   						recordFormat: &recordFormatProperty{
//   							recordFormatType: jsii.String("recordFormatType"),
//
//   							// the properties below are optional
//   							mappingParameters: &mappingParametersProperty{
//   								csvMappingParameters: &cSVMappingParametersProperty{
//   									recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   									recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   								},
//   								jsonMappingParameters: &jSONMappingParametersProperty{
//   									recordRowPath: jsii.String("recordRowPath"),
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						recordEncoding: jsii.String("recordEncoding"),
//   					},
//   					namePrefix: jsii.String("namePrefix"),
//
//   					// the properties below are optional
//   					inputParallelism: &inputParallelismProperty{
//   						count: jsii.Number(123),
//   					},
//   					inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   						inputLambdaProcessor: &inputLambdaProcessorProperty{
//   							resourceArn: jsii.String("resourceArn"),
//   						},
//   					},
//   					kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   						resourceArn: jsii.String("resourceArn"),
//   					},
//   					kinesisStreamsInput: &kinesisStreamsInputProperty{
//   						resourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   			},
//   		},
//   		vpcConfigurations: []interface{}{
//   			&vpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		zeppelinApplicationConfiguration: &zeppelinApplicationConfigurationProperty{
//   			catalogConfiguration: &catalogConfigurationProperty{
//   				glueDataCatalogConfiguration: &glueDataCatalogConfigurationProperty{
//   					databaseArn: jsii.String("databaseArn"),
//   				},
//   			},
//   			customArtifactsConfiguration: []interface{}{
//   				&customArtifactConfigurationProperty{
//   					artifactType: jsii.String("artifactType"),
//
//   					// the properties below are optional
//   					mavenReference: &mavenReferenceProperty{
//   						artifactId: jsii.String("artifactId"),
//   						groupId: jsii.String("groupId"),
//   						version: jsii.String("version"),
//   					},
//   					s3ContentLocation: &s3ContentLocationProperty{
//   						bucketArn: jsii.String("bucketArn"),
//   						fileKey: jsii.String("fileKey"),
//
//   						// the properties below are optional
//   						objectVersion: jsii.String("objectVersion"),
//   					},
//   				},
//   			},
//   			deployAsApplicationConfiguration: &deployAsApplicationConfigurationProperty{
//   				s3ContentLocation: &s3ContentBaseLocationProperty{
//   					bucketArn: jsii.String("bucketArn"),
//
//   					// the properties below are optional
//   					basePath: jsii.String("basePath"),
//   				},
//   			},
//   			monitoringConfiguration: &zeppelinMonitoringConfigurationProperty{
//   				logLevel: jsii.String("logLevel"),
//   			},
//   		},
//   	},
//   	applicationDescription: jsii.String("applicationDescription"),
//   	applicationMaintenanceConfiguration: &applicationMaintenanceConfigurationProperty{
//   		applicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   	},
//   	applicationMode: jsii.String("applicationMode"),
//   	applicationName: jsii.String("applicationName"),
//   	runConfiguration: &runConfigurationProperty{
//   		applicationRestoreConfiguration: &applicationRestoreConfigurationProperty{
//   			applicationRestoreType: jsii.String("applicationRestoreType"),
//
//   			// the properties below are optional
//   			snapshotName: jsii.String("snapshotName"),
//   		},
//   		flinkRunConfiguration: &flinkRunConfigurationProperty{
//   			allowNonRestoredState: jsii.Boolean(false),
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
type CfnApplicationV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Use this parameter to configure the application.
	ApplicationConfiguration() interface{}
	SetApplicationConfiguration(val interface{})
	// The description of the application.
	ApplicationDescription() *string
	SetApplicationDescription(val *string)
	// `AWS::KinesisAnalyticsV2::Application.ApplicationMaintenanceConfiguration`.
	ApplicationMaintenanceConfiguration() interface{}
	SetApplicationMaintenanceConfiguration(val interface{})
	// To create a Kinesis Data Analytics Studio notebook, you must set the mode to `INTERACTIVE` .
	//
	// However, for a Kinesis Data Analytics for Apache Flink application, the mode is optional.
	ApplicationMode() *string
	SetApplicationMode(val *string)
	// The name of the application.
	ApplicationName() *string
	SetApplicationName(val *string)
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
	// `AWS::KinesisAnalyticsV2::Application.RunConfiguration`.
	RunConfiguration() interface{}
	SetRunConfiguration(val interface{})
	// The runtime environment for the application.
	RuntimeEnvironment() *string
	SetRuntimeEnvironment(val *string)
	// Specifies the IAM role that the application uses to access external resources.
	ServiceExecutionRole() *string
	SetServiceExecutionRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of one or more tags to assign to the application.
	//
	// A tag is a key-value pair that identifies an application. Note that the maximum number of application tags includes system tags. The maximum number of user-defined application tags is 50.
	Tags() awscdk.TagManager
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

func (j *jsiiProxy_CfnApplicationV2) ApplicationMaintenanceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationMaintenanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationMode",
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

func (j *jsiiProxy_CfnApplicationV2) RunConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runConfiguration",
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

func (j *jsiiProxy_CfnApplicationV2) SetApplicationMaintenanceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"applicationMaintenanceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2) SetApplicationMode(val *string) {
	_jsii_.Set(
		j,
		"applicationMode",
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

func (j *jsiiProxy_CfnApplicationV2) SetRunConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"runConfiguration",
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

func (c *jsiiProxy_CfnApplicationV2) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplicationV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnApplicationV2) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplicationV2) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplicationV2) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationV2) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

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

func (c *jsiiProxy_CfnApplicationV2) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplicationV2) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes code configuration for an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationCodeConfigurationProperty := &applicationCodeConfigurationProperty{
//   	codeContent: &codeContentProperty{
//   		s3ContentLocation: &s3ContentLocationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			fileKey: jsii.String("fileKey"),
//
//   			// the properties below are optional
//   			objectVersion: jsii.String("objectVersion"),
//   		},
//   		textContent: jsii.String("textContent"),
//   		zipFileContent: jsii.String("zipFileContent"),
//   	},
//   	codeContentType: jsii.String("codeContentType"),
//   }
//
type CfnApplicationV2_ApplicationCodeConfigurationProperty struct {
	// The location and type of the application code.
	CodeContent interface{} `field:"required" json:"codeContent" yaml:"codeContent"`
	// Specifies whether the code content is in text or zip format.
	CodeContentType *string `field:"required" json:"codeContentType" yaml:"codeContentType"`
}

// Specifies the creation parameters for a Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationConfigurationProperty := &applicationConfigurationProperty{
//   	applicationCodeConfiguration: &applicationCodeConfigurationProperty{
//   		codeContent: &codeContentProperty{
//   			s3ContentLocation: &s3ContentLocationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//   				fileKey: jsii.String("fileKey"),
//
//   				// the properties below are optional
//   				objectVersion: jsii.String("objectVersion"),
//   			},
//   			textContent: jsii.String("textContent"),
//   			zipFileContent: jsii.String("zipFileContent"),
//   		},
//   		codeContentType: jsii.String("codeContentType"),
//   	},
//   	applicationSnapshotConfiguration: &applicationSnapshotConfigurationProperty{
//   		snapshotsEnabled: jsii.Boolean(false),
//   	},
//   	environmentProperties: &environmentPropertiesProperty{
//   		propertyGroups: []interface{}{
//   			&propertyGroupProperty{
//   				propertyGroupId: jsii.String("propertyGroupId"),
//   				propertyMap: map[string]*string{
//   					"propertyMapKey": jsii.String("propertyMap"),
//   				},
//   			},
//   		},
//   	},
//   	flinkApplicationConfiguration: &flinkApplicationConfigurationProperty{
//   		checkpointConfiguration: &checkpointConfigurationProperty{
//   			configurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			checkpointingEnabled: jsii.Boolean(false),
//   			checkpointInterval: jsii.Number(123),
//   			minPauseBetweenCheckpoints: jsii.Number(123),
//   		},
//   		monitoringConfiguration: &monitoringConfigurationProperty{
//   			configurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			logLevel: jsii.String("logLevel"),
//   			metricsLevel: jsii.String("metricsLevel"),
//   		},
//   		parallelismConfiguration: &parallelismConfigurationProperty{
//   			configurationType: jsii.String("configurationType"),
//
//   			// the properties below are optional
//   			autoScalingEnabled: jsii.Boolean(false),
//   			parallelism: jsii.Number(123),
//   			parallelismPerKpu: jsii.Number(123),
//   		},
//   	},
//   	sqlApplicationConfiguration: &sqlApplicationConfigurationProperty{
//   		inputs: []interface{}{
//   			&inputProperty{
//   				inputSchema: &inputSchemaProperty{
//   					recordColumns: []interface{}{
//   						&recordColumnProperty{
//   							name: jsii.String("name"),
//   							sqlType: jsii.String("sqlType"),
//
//   							// the properties below are optional
//   							mapping: jsii.String("mapping"),
//   						},
//   					},
//   					recordFormat: &recordFormatProperty{
//   						recordFormatType: jsii.String("recordFormatType"),
//
//   						// the properties below are optional
//   						mappingParameters: &mappingParametersProperty{
//   							csvMappingParameters: &cSVMappingParametersProperty{
//   								recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   								recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   							},
//   							jsonMappingParameters: &jSONMappingParametersProperty{
//   								recordRowPath: jsii.String("recordRowPath"),
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					recordEncoding: jsii.String("recordEncoding"),
//   				},
//   				namePrefix: jsii.String("namePrefix"),
//
//   				// the properties below are optional
//   				inputParallelism: &inputParallelismProperty{
//   					count: jsii.Number(123),
//   				},
//   				inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   					inputLambdaProcessor: &inputLambdaProcessorProperty{
//   						resourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   				kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   					resourceArn: jsii.String("resourceArn"),
//   				},
//   				kinesisStreamsInput: &kinesisStreamsInputProperty{
//   					resourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   		},
//   	},
//   	vpcConfigurations: []interface{}{
//   		&vpcConfigurationProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	zeppelinApplicationConfiguration: &zeppelinApplicationConfigurationProperty{
//   		catalogConfiguration: &catalogConfigurationProperty{
//   			glueDataCatalogConfiguration: &glueDataCatalogConfigurationProperty{
//   				databaseArn: jsii.String("databaseArn"),
//   			},
//   		},
//   		customArtifactsConfiguration: []interface{}{
//   			&customArtifactConfigurationProperty{
//   				artifactType: jsii.String("artifactType"),
//
//   				// the properties below are optional
//   				mavenReference: &mavenReferenceProperty{
//   					artifactId: jsii.String("artifactId"),
//   					groupId: jsii.String("groupId"),
//   					version: jsii.String("version"),
//   				},
//   				s3ContentLocation: &s3ContentLocationProperty{
//   					bucketArn: jsii.String("bucketArn"),
//   					fileKey: jsii.String("fileKey"),
//
//   					// the properties below are optional
//   					objectVersion: jsii.String("objectVersion"),
//   				},
//   			},
//   		},
//   		deployAsApplicationConfiguration: &deployAsApplicationConfigurationProperty{
//   			s3ContentLocation: &s3ContentBaseLocationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//
//   				// the properties below are optional
//   				basePath: jsii.String("basePath"),
//   			},
//   		},
//   		monitoringConfiguration: &zeppelinMonitoringConfigurationProperty{
//   			logLevel: jsii.String("logLevel"),
//   		},
//   	},
//   }
//
type CfnApplicationV2_ApplicationConfigurationProperty struct {
	// The code location and type parameters for a Flink-based Kinesis Data Analytics application.
	ApplicationCodeConfiguration interface{} `field:"optional" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	ApplicationSnapshotConfiguration interface{} `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// Describes execution properties for a Flink-based Kinesis Data Analytics application.
	EnvironmentProperties interface{} `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// The creation and update parameters for a Flink-based Kinesis Data Analytics application.
	FlinkApplicationConfiguration interface{} `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// The creation and update parameters for a SQL-based Kinesis Data Analytics application.
	SqlApplicationConfiguration interface{} `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// `CfnApplicationV2.ApplicationConfigurationProperty.VpcConfigurations`.
	VpcConfigurations interface{} `field:"optional" json:"vpcConfigurations" yaml:"vpcConfigurations"`
	// The configuration parameters for a Kinesis Data Analytics Studio notebook.
	ZeppelinApplicationConfiguration interface{} `field:"optional" json:"zeppelinApplicationConfiguration" yaml:"zeppelinApplicationConfiguration"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationMaintenanceConfigurationProperty := &applicationMaintenanceConfigurationProperty{
//   	applicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   }
//
type CfnApplicationV2_ApplicationMaintenanceConfigurationProperty struct {
	// `CfnApplicationV2.ApplicationMaintenanceConfigurationProperty.ApplicationMaintenanceWindowStartTime`.
	ApplicationMaintenanceWindowStartTime *string `field:"required" json:"applicationMaintenanceWindowStartTime" yaml:"applicationMaintenanceWindowStartTime"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationRestoreConfigurationProperty := &applicationRestoreConfigurationProperty{
//   	applicationRestoreType: jsii.String("applicationRestoreType"),
//
//   	// the properties below are optional
//   	snapshotName: jsii.String("snapshotName"),
//   }
//
type CfnApplicationV2_ApplicationRestoreConfigurationProperty struct {
	// `CfnApplicationV2.ApplicationRestoreConfigurationProperty.ApplicationRestoreType`.
	ApplicationRestoreType *string `field:"required" json:"applicationRestoreType" yaml:"applicationRestoreType"`
	// `CfnApplicationV2.ApplicationRestoreConfigurationProperty.SnapshotName`.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSnapshotConfigurationProperty := &applicationSnapshotConfigurationProperty{
//   	snapshotsEnabled: jsii.Boolean(false),
//   }
//
type CfnApplicationV2_ApplicationSnapshotConfigurationProperty struct {
	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	SnapshotsEnabled interface{} `field:"required" json:"snapshotsEnabled" yaml:"snapshotsEnabled"`
}

// For a SQL-based Kinesis Data Analytics application, provides additional mapping information when the record format uses delimiters, such as CSV.
//
// For example, the following sample records use CSV format, where the records use the *'\n'* as the row delimiter and a comma (",") as the column delimiter:
//
// `"name1", "address1"`
//
// `"name2", "address2"`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cSVMappingParametersProperty := &cSVMappingParametersProperty{
//   	recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   	recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   }
//
type CfnApplicationV2_CSVMappingParametersProperty struct {
	// The column delimiter.
	//
	// For example, in a CSV format, a comma (",") is the typical column delimiter.
	RecordColumnDelimiter *string `field:"required" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// The row delimiter.
	//
	// For example, in a CSV format, *'\n'* is the typical row delimiter.
	RecordRowDelimiter *string `field:"required" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

// The configuration parameters for the default Amazon Glue database.
//
// You use this database for SQL queries that you write in a Kinesis Data Analytics Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catalogConfigurationProperty := &catalogConfigurationProperty{
//   	glueDataCatalogConfiguration: &glueDataCatalogConfigurationProperty{
//   		databaseArn: jsii.String("databaseArn"),
//   	},
//   }
//
type CfnApplicationV2_CatalogConfigurationProperty struct {
	// The configuration parameters for the default Amazon Glue database.
	//
	// You use this database for Apache Flink SQL queries and table API transforms that you write in a Kinesis Data Analytics Studio notebook.
	GlueDataCatalogConfiguration interface{} `field:"optional" json:"glueDataCatalogConfiguration" yaml:"glueDataCatalogConfiguration"`
}

// Describes an application's checkpointing configuration.
//
// Checkpointing is the process of persisting application state for fault tolerance. For more information, see [Checkpoints for Fault Tolerance](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/concepts/programming-model.html#checkpoints-for-fault-tolerance) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   checkpointConfigurationProperty := &checkpointConfigurationProperty{
//   	configurationType: jsii.String("configurationType"),
//
//   	// the properties below are optional
//   	checkpointingEnabled: jsii.Boolean(false),
//   	checkpointInterval: jsii.Number(123),
//   	minPauseBetweenCheckpoints: jsii.Number(123),
//   }
//
type CfnApplicationV2_CheckpointConfigurationProperty struct {
	// Describes whether the application uses Kinesis Data Analytics' default checkpointing behavior.
	//
	// You must set this property to `CUSTOM` in order to set the `CheckpointingEnabled` , `CheckpointInterval` , or `MinPauseBetweenCheckpoints` parameters.
	//
	// > If this value is set to `DEFAULT` , the application will use the following values, even if they are set to other values using APIs or application code:
	// >
	// > - *CheckpointingEnabled:* true
	// > - *CheckpointInterval:* 60000
	// > - *MinPauseBetweenCheckpoints:* 5000.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Describes whether checkpointing is enabled for a Flink-based Kinesis Data Analytics application.
	//
	// > If `CheckpointConfiguration.ConfigurationType` is `DEFAULT` , the application will use a `CheckpointingEnabled` value of `true` , even if this value is set to another value using this API or in application code.
	CheckpointingEnabled interface{} `field:"optional" json:"checkpointingEnabled" yaml:"checkpointingEnabled"`
	// Describes the interval in milliseconds between checkpoint operations.
	//
	// > If `CheckpointConfiguration.ConfigurationType` is `DEFAULT` , the application will use a `CheckpointInterval` value of 60000, even if this value is set to another value using this API or in application code.
	CheckpointInterval *float64 `field:"optional" json:"checkpointInterval" yaml:"checkpointInterval"`
	// Describes the minimum time in milliseconds after a checkpoint operation completes that a new checkpoint operation can start.
	//
	// If a checkpoint operation takes longer than the `CheckpointInterval` , the application otherwise performs continual checkpoint operations. For more information, see [Tuning Checkpointing](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/ops/state/large_state_tuning.html#tuning-checkpointing) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/) .
	//
	// > If `CheckpointConfiguration.ConfigurationType` is `DEFAULT` , the application will use a `MinPauseBetweenCheckpoints` value of 5000, even if this value is set using this API or in application code.
	MinPauseBetweenCheckpoints *float64 `field:"optional" json:"minPauseBetweenCheckpoints" yaml:"minPauseBetweenCheckpoints"`
}

// Specifies either the application code, or the location of the application code, for a Flink-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeContentProperty := &codeContentProperty{
//   	s3ContentLocation: &s3ContentLocationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		fileKey: jsii.String("fileKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   	textContent: jsii.String("textContent"),
//   	zipFileContent: jsii.String("zipFileContent"),
//   }
//
type CfnApplicationV2_CodeContentProperty struct {
	// Information about the Amazon S3 bucket that contains the application code.
	S3ContentLocation interface{} `field:"optional" json:"s3ContentLocation" yaml:"s3ContentLocation"`
	// The text-format code for a Flink-based Kinesis Data Analytics application.
	TextContent *string `field:"optional" json:"textContent" yaml:"textContent"`
	// The zip-format code for a Flink-based Kinesis Data Analytics application.
	ZipFileContent *string `field:"optional" json:"zipFileContent" yaml:"zipFileContent"`
}

// The configuration of connectors and user-defined functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customArtifactConfigurationProperty := &customArtifactConfigurationProperty{
//   	artifactType: jsii.String("artifactType"),
//
//   	// the properties below are optional
//   	mavenReference: &mavenReferenceProperty{
//   		artifactId: jsii.String("artifactId"),
//   		groupId: jsii.String("groupId"),
//   		version: jsii.String("version"),
//   	},
//   	s3ContentLocation: &s3ContentLocationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		fileKey: jsii.String("fileKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
type CfnApplicationV2_CustomArtifactConfigurationProperty struct {
	// Set this to either `UDF` or `DEPENDENCY_JAR` .
	//
	// `UDF` stands for user-defined functions. This type of artifact must be in an S3 bucket. A `DEPENDENCY_JAR` can be in either Maven or an S3 bucket.
	ArtifactType *string `field:"required" json:"artifactType" yaml:"artifactType"`
	// The parameters required to fully specify a Maven reference.
	MavenReference interface{} `field:"optional" json:"mavenReference" yaml:"mavenReference"`
	// The location of the custom artifacts.
	S3ContentLocation interface{} `field:"optional" json:"s3ContentLocation" yaml:"s3ContentLocation"`
}

// The information required to deploy a Kinesis Data Analytics Studio notebook as an application with durable state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deployAsApplicationConfigurationProperty := &deployAsApplicationConfigurationProperty{
//   	s3ContentLocation: &s3ContentBaseLocationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//
//   		// the properties below are optional
//   		basePath: jsii.String("basePath"),
//   	},
//   }
//
type CfnApplicationV2_DeployAsApplicationConfigurationProperty struct {
	// The description of an Amazon S3 object that contains the Amazon Data Analytics application, including the Amazon Resource Name (ARN) of the S3 bucket, the name of the Amazon S3 object that contains the data, and the version number of the Amazon S3 object that contains the data.
	S3ContentLocation interface{} `field:"required" json:"s3ContentLocation" yaml:"s3ContentLocation"`
}

// Describes execution properties for a Flink-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentPropertiesProperty := &environmentPropertiesProperty{
//   	propertyGroups: []interface{}{
//   		&propertyGroupProperty{
//   			propertyGroupId: jsii.String("propertyGroupId"),
//   			propertyMap: map[string]*string{
//   				"propertyMapKey": jsii.String("propertyMap"),
//   			},
//   		},
//   	},
//   }
//
type CfnApplicationV2_EnvironmentPropertiesProperty struct {
	// Describes the execution property groups.
	PropertyGroups interface{} `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
}

// Describes configuration parameters for a Flink-based Kinesis Data Analytics application or a Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flinkApplicationConfigurationProperty := &flinkApplicationConfigurationProperty{
//   	checkpointConfiguration: &checkpointConfigurationProperty{
//   		configurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		checkpointingEnabled: jsii.Boolean(false),
//   		checkpointInterval: jsii.Number(123),
//   		minPauseBetweenCheckpoints: jsii.Number(123),
//   	},
//   	monitoringConfiguration: &monitoringConfigurationProperty{
//   		configurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		logLevel: jsii.String("logLevel"),
//   		metricsLevel: jsii.String("metricsLevel"),
//   	},
//   	parallelismConfiguration: &parallelismConfigurationProperty{
//   		configurationType: jsii.String("configurationType"),
//
//   		// the properties below are optional
//   		autoScalingEnabled: jsii.Boolean(false),
//   		parallelism: jsii.Number(123),
//   		parallelismPerKpu: jsii.Number(123),
//   	},
//   }
//
type CfnApplicationV2_FlinkApplicationConfigurationProperty struct {
	// Describes an application's checkpointing configuration.
	//
	// Checkpointing is the process of persisting application state for fault tolerance. For more information, see [Checkpoints for Fault Tolerance](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/concepts/programming-model.html#checkpoints-for-fault-tolerance) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/) .
	CheckpointConfiguration interface{} `field:"optional" json:"checkpointConfiguration" yaml:"checkpointConfiguration"`
	// Describes configuration parameters for Amazon CloudWatch logging for an application.
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// Describes parameters for how an application executes multiple tasks simultaneously.
	ParallelismConfiguration interface{} `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flinkRunConfigurationProperty := &flinkRunConfigurationProperty{
//   	allowNonRestoredState: jsii.Boolean(false),
//   }
//
type CfnApplicationV2_FlinkRunConfigurationProperty struct {
	// `CfnApplicationV2.FlinkRunConfigurationProperty.AllowNonRestoredState`.
	AllowNonRestoredState interface{} `field:"optional" json:"allowNonRestoredState" yaml:"allowNonRestoredState"`
}

// The configuration of the Glue Data Catalog that you use for Apache Flink SQL queries and table API transforms that you write in an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueDataCatalogConfigurationProperty := &glueDataCatalogConfigurationProperty{
//   	databaseArn: jsii.String("databaseArn"),
//   }
//
type CfnApplicationV2_GlueDataCatalogConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the database.
	DatabaseArn *string `field:"optional" json:"databaseArn" yaml:"databaseArn"`
}

// An object that contains the Amazon Resource Name (ARN) of the Amazon Lambda function that is used to preprocess records in the stream in a SQL-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputLambdaProcessorProperty := &inputLambdaProcessorProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationV2_InputLambdaProcessorProperty struct {
	// The ARN of the Amazon Lambda function that operates on records in the stream.
	//
	// > To specify an earlier version of the Lambda function than the latest, include the Lambda function version in the Lambda function ARN. For more information about Lambda ARNs, see [Example ARNs: Amazon Lambda](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-lambda)
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

// For a SQL-based Kinesis Data Analytics application, describes the number of in-application streams to create for a given streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputParallelismProperty := &inputParallelismProperty{
//   	count: jsii.Number(123),
//   }
//
type CfnApplicationV2_InputParallelismProperty struct {
	// The number of in-application streams to create.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

// For an SQL-based Amazon Kinesis Data Analytics application, describes a processor that is used to preprocess the records in the stream before being processed by your application code.
//
// Currently, the only input processor available is [Amazon Lambda](https://docs.aws.amazon.com/lambda/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProcessingConfigurationProperty := &inputProcessingConfigurationProperty{
//   	inputLambdaProcessor: &inputLambdaProcessorProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   }
//
type CfnApplicationV2_InputProcessingConfigurationProperty struct {
	// The [InputLambdaProcessor](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_InputLambdaProcessor.html) that is used to preprocess the records in the stream before being processed by your application code.
	InputLambdaProcessor interface{} `field:"optional" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}

// When you configure the application input for a SQL-based Kinesis Data Analytics application, you specify the streaming source, the in-application stream name that is created, and the mapping between the two.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputProperty := &inputProperty{
//   	inputSchema: &inputSchemaProperty{
//   		recordColumns: []interface{}{
//   			&recordColumnProperty{
//   				name: jsii.String("name"),
//   				sqlType: jsii.String("sqlType"),
//
//   				// the properties below are optional
//   				mapping: jsii.String("mapping"),
//   			},
//   		},
//   		recordFormat: &recordFormatProperty{
//   			recordFormatType: jsii.String("recordFormatType"),
//
//   			// the properties below are optional
//   			mappingParameters: &mappingParametersProperty{
//   				csvMappingParameters: &cSVMappingParametersProperty{
//   					recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   					recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   				},
//   				jsonMappingParameters: &jSONMappingParametersProperty{
//   					recordRowPath: jsii.String("recordRowPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		recordEncoding: jsii.String("recordEncoding"),
//   	},
//   	namePrefix: jsii.String("namePrefix"),
//
//   	// the properties below are optional
//   	inputParallelism: &inputParallelismProperty{
//   		count: jsii.Number(123),
//   	},
//   	inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   		inputLambdaProcessor: &inputLambdaProcessorProperty{
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   	kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   	kinesisStreamsInput: &kinesisStreamsInputProperty{
//   		resourceArn: jsii.String("resourceArn"),
//   	},
//   }
//
type CfnApplicationV2_InputProperty struct {
	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
	//
	// Also used to describe the format of the reference data source.
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// The name prefix to use when creating an in-application stream.
	//
	// Suppose that you specify a prefix " `MyInApplicationStream` ." Kinesis Data Analytics then creates one or more (as per the `InputParallelism` count you specified) in-application streams with the names " `MyInApplicationStream_001` ," " `MyInApplicationStream_002` ," and so on.
	NamePrefix *string `field:"required" json:"namePrefix" yaml:"namePrefix"`
	// Describes the number of in-application streams to create.
	InputParallelism interface{} `field:"optional" json:"inputParallelism" yaml:"inputParallelism"`
	// The [InputProcessingConfiguration](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_InputProcessingConfiguration.html) for the input. An input processor transforms records as they are received from the stream, before the application's SQL code executes. Currently, the only input processing configuration available is [InputLambdaProcessor](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_InputLambdaProcessor.html) .
	InputProcessingConfiguration interface{} `field:"optional" json:"inputProcessingConfiguration" yaml:"inputProcessingConfiguration"`
	// If the streaming source is an Amazon Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.
	KinesisFirehoseInput interface{} `field:"optional" json:"kinesisFirehoseInput" yaml:"kinesisFirehoseInput"`
	// If the streaming source is an Amazon Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).
	KinesisStreamsInput interface{} `field:"optional" json:"kinesisStreamsInput" yaml:"kinesisStreamsInput"`
}

// For a SQL-based Kinesis Data Analytics application, describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSchemaProperty := &inputSchemaProperty{
//   	recordColumns: []interface{}{
//   		&recordColumnProperty{
//   			name: jsii.String("name"),
//   			sqlType: jsii.String("sqlType"),
//
//   			// the properties below are optional
//   			mapping: jsii.String("mapping"),
//   		},
//   	},
//   	recordFormat: &recordFormatProperty{
//   		recordFormatType: jsii.String("recordFormatType"),
//
//   		// the properties below are optional
//   		mappingParameters: &mappingParametersProperty{
//   			csvMappingParameters: &cSVMappingParametersProperty{
//   				recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   				recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   			},
//   			jsonMappingParameters: &jSONMappingParametersProperty{
//   				recordRowPath: jsii.String("recordRowPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	recordEncoding: jsii.String("recordEncoding"),
//   }
//
type CfnApplicationV2_InputSchemaProperty struct {
	// A list of `RecordColumn` objects.
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// Specifies the format of the records on the streaming source.
	RecordFormat interface{} `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Specifies the encoding of the records in the streaming source.
	//
	// For example, UTF-8.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

// For a SQL-based Kinesis Data Analytics application, provides additional mapping information when JSON is the record format on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jSONMappingParametersProperty := &jSONMappingParametersProperty{
//   	recordRowPath: jsii.String("recordRowPath"),
//   }
//
type CfnApplicationV2_JSONMappingParametersProperty struct {
	// The path to the top-level parent that contains the records.
	RecordRowPath *string `field:"required" json:"recordRowPath" yaml:"recordRowPath"`
}

// For a SQL-based Kinesis Data Analytics application, identifies a Kinesis Data Firehose delivery stream as the streaming source.
//
// You provide the delivery stream's Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseInputProperty := &kinesisFirehoseInputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationV2_KinesisFirehoseInputProperty struct {
	// The Amazon Resource Name (ARN) of the delivery stream.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

// Identifies a Kinesis data stream as the streaming source.
//
// You provide the stream's Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamsInputProperty := &kinesisStreamsInputProperty{
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnApplicationV2_KinesisStreamsInputProperty struct {
	// The ARN of the input Kinesis data stream to read.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

// When you configure a SQL-based Kinesis Data Analytics application's input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingParametersProperty := &mappingParametersProperty{
//   	csvMappingParameters: &cSVMappingParametersProperty{
//   		recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   		recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   	},
//   	jsonMappingParameters: &jSONMappingParametersProperty{
//   		recordRowPath: jsii.String("recordRowPath"),
//   	},
//   }
//
type CfnApplicationV2_MappingParametersProperty struct {
	// Provides additional mapping information when the record format uses delimiters (for example, CSV).
	CsvMappingParameters interface{} `field:"optional" json:"csvMappingParameters" yaml:"csvMappingParameters"`
	// Provides additional mapping information when JSON is the record format on the streaming source.
	JsonMappingParameters interface{} `field:"optional" json:"jsonMappingParameters" yaml:"jsonMappingParameters"`
}

// The information required to specify a Maven reference.
//
// You can use Maven references to specify dependency JAR files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mavenReferenceProperty := &mavenReferenceProperty{
//   	artifactId: jsii.String("artifactId"),
//   	groupId: jsii.String("groupId"),
//   	version: jsii.String("version"),
//   }
//
type CfnApplicationV2_MavenReferenceProperty struct {
	// The artifact ID of the Maven reference.
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
	// The group ID of the Maven reference.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The version of the Maven reference.
	Version *string `field:"required" json:"version" yaml:"version"`
}

// Describes configuration parameters for Amazon CloudWatch logging for a Java-based Kinesis Data Analytics application.
//
// For more information about CloudWatch logging, see [Monitoring](https://docs.aws.amazon.com/kinesisanalytics/latest/java/monitoring-overview) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringConfigurationProperty := &monitoringConfigurationProperty{
//   	configurationType: jsii.String("configurationType"),
//
//   	// the properties below are optional
//   	logLevel: jsii.String("logLevel"),
//   	metricsLevel: jsii.String("metricsLevel"),
//   }
//
type CfnApplicationV2_MonitoringConfigurationProperty struct {
	// Describes whether to use the default CloudWatch logging configuration for an application.
	//
	// You must set this property to `CUSTOM` in order to set the `LogLevel` or `MetricsLevel` parameters.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Describes the verbosity of the CloudWatch Logs for an application.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Describes the granularity of the CloudWatch Logs for an application.
	//
	// The `Parallelism` level is not recommended for applications with a Parallelism over 64 due to excessive costs.
	MetricsLevel *string `field:"optional" json:"metricsLevel" yaml:"metricsLevel"`
}

// Describes parameters for how a Flink-based Kinesis Data Analytics application executes multiple tasks simultaneously.
//
// For more information about parallelism, see [Parallel Execution](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/dev/parallel.html) in the [Apache Flink Documentation](https://docs.aws.amazon.com/https://ci.apache.org/projects/flink/flink-docs-release-1.8/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parallelismConfigurationProperty := &parallelismConfigurationProperty{
//   	configurationType: jsii.String("configurationType"),
//
//   	// the properties below are optional
//   	autoScalingEnabled: jsii.Boolean(false),
//   	parallelism: jsii.Number(123),
//   	parallelismPerKpu: jsii.Number(123),
//   }
//
type CfnApplicationV2_ParallelismConfigurationProperty struct {
	// Describes whether the application uses the default parallelism for the Kinesis Data Analytics service.
	//
	// You must set this property to `CUSTOM` in order to change your application's `AutoScalingEnabled` , `Parallelism` , or `ParallelismPerKPU` properties.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Describes whether the Kinesis Data Analytics service can increase the parallelism of the application in response to increased throughput.
	AutoScalingEnabled interface{} `field:"optional" json:"autoScalingEnabled" yaml:"autoScalingEnabled"`
	// Describes the initial number of parallel tasks that a Java-based Kinesis Data Analytics application can perform.
	//
	// The Kinesis Data Analytics service can increase this number automatically if [ParallelismConfiguration:AutoScalingEnabled](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_ParallelismConfiguration.html#kinesisanalytics-Type-ParallelismConfiguration-AutoScalingEnabled.html) is set to `true` .
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// Describes the number of parallel tasks that a Java-based Kinesis Data Analytics application can perform per Kinesis Processing Unit (KPU) used by the application.
	//
	// For more information about KPUs, see [Amazon Kinesis Data Analytics Pricing](https://docs.aws.amazon.com/kinesis/data-analytics/pricing/) .
	ParallelismPerKpu *float64 `field:"optional" json:"parallelismPerKpu" yaml:"parallelismPerKpu"`
}

// Property key-value pairs passed into an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyGroupProperty := &propertyGroupProperty{
//   	propertyGroupId: jsii.String("propertyGroupId"),
//   	propertyMap: map[string]*string{
//   		"propertyMapKey": jsii.String("propertyMap"),
//   	},
//   }
//
type CfnApplicationV2_PropertyGroupProperty struct {
	// Describes the key of an application execution property key-value pair.
	PropertyGroupId *string `field:"optional" json:"propertyGroupId" yaml:"propertyGroupId"`
	// Describes the value of an application execution property key-value pair.
	PropertyMap interface{} `field:"optional" json:"propertyMap" yaml:"propertyMap"`
}

// For a SQL-based Kinesis Data Analytics application, describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
//
// Also used to describe the format of the reference data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordColumnProperty := &recordColumnProperty{
//   	name: jsii.String("name"),
//   	sqlType: jsii.String("sqlType"),
//
//   	// the properties below are optional
//   	mapping: jsii.String("mapping"),
//   }
//
type CfnApplicationV2_RecordColumnProperty struct {
	// The name of the column that is created in the in-application input stream or reference table.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of column created in the in-application input stream or reference table.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// A reference to the data element in the streaming input or the reference data source.
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

// For a SQL-based Kinesis Data Analytics application, describes the record format and relevant mapping information that should be applied to schematize the records on the stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordFormatProperty := &recordFormatProperty{
//   	recordFormatType: jsii.String("recordFormatType"),
//
//   	// the properties below are optional
//   	mappingParameters: &mappingParametersProperty{
//   		csvMappingParameters: &cSVMappingParametersProperty{
//   			recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   			recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   		},
//   		jsonMappingParameters: &jSONMappingParametersProperty{
//   			recordRowPath: jsii.String("recordRowPath"),
//   		},
//   	},
//   }
//
type CfnApplicationV2_RecordFormatProperty struct {
	// The type of record format.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// When you configure application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	MappingParameters interface{} `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runConfigurationProperty := &runConfigurationProperty{
//   	applicationRestoreConfiguration: &applicationRestoreConfigurationProperty{
//   		applicationRestoreType: jsii.String("applicationRestoreType"),
//
//   		// the properties below are optional
//   		snapshotName: jsii.String("snapshotName"),
//   	},
//   	flinkRunConfiguration: &flinkRunConfigurationProperty{
//   		allowNonRestoredState: jsii.Boolean(false),
//   	},
//   }
//
type CfnApplicationV2_RunConfigurationProperty struct {
	// `CfnApplicationV2.RunConfigurationProperty.ApplicationRestoreConfiguration`.
	ApplicationRestoreConfiguration interface{} `field:"optional" json:"applicationRestoreConfiguration" yaml:"applicationRestoreConfiguration"`
	// `CfnApplicationV2.RunConfigurationProperty.FlinkRunConfiguration`.
	FlinkRunConfiguration interface{} `field:"optional" json:"flinkRunConfiguration" yaml:"flinkRunConfiguration"`
}

// The base location of the Amazon Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ContentBaseLocationProperty := &s3ContentBaseLocationProperty{
//   	bucketArn: jsii.String("bucketArn"),
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   }
//
type CfnApplicationV2_S3ContentBaseLocationProperty struct {
	// The Amazon Resource Name (ARN) of the S3 bucket.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The base path for the S3 bucket.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
}

// The location of an application or a custom artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ContentLocationProperty := &s3ContentLocationProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	fileKey: jsii.String("fileKey"),
//
//   	// the properties below are optional
//   	objectVersion: jsii.String("objectVersion"),
//   }
//
type CfnApplicationV2_S3ContentLocationProperty struct {
	// The Amazon Resource Name (ARN) for the S3 bucket containing the application code.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The file key for the object containing the application code.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// The version of the object containing the application code.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

// Describes the inputs, outputs, and reference data sources for a SQL-based Kinesis Data Analytics application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqlApplicationConfigurationProperty := &sqlApplicationConfigurationProperty{
//   	inputs: []interface{}{
//   		&inputProperty{
//   			inputSchema: &inputSchemaProperty{
//   				recordColumns: []interface{}{
//   					&recordColumnProperty{
//   						name: jsii.String("name"),
//   						sqlType: jsii.String("sqlType"),
//
//   						// the properties below are optional
//   						mapping: jsii.String("mapping"),
//   					},
//   				},
//   				recordFormat: &recordFormatProperty{
//   					recordFormatType: jsii.String("recordFormatType"),
//
//   					// the properties below are optional
//   					mappingParameters: &mappingParametersProperty{
//   						csvMappingParameters: &cSVMappingParametersProperty{
//   							recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   							recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   						},
//   						jsonMappingParameters: &jSONMappingParametersProperty{
//   							recordRowPath: jsii.String("recordRowPath"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				recordEncoding: jsii.String("recordEncoding"),
//   			},
//   			namePrefix: jsii.String("namePrefix"),
//
//   			// the properties below are optional
//   			inputParallelism: &inputParallelismProperty{
//   				count: jsii.Number(123),
//   			},
//   			inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   				inputLambdaProcessor: &inputLambdaProcessorProperty{
//   					resourceArn: jsii.String("resourceArn"),
//   				},
//   			},
//   			kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   			},
//   			kinesisStreamsInput: &kinesisStreamsInputProperty{
//   				resourceArn: jsii.String("resourceArn"),
//   			},
//   		},
//   	},
//   }
//
type CfnApplicationV2_SqlApplicationConfigurationProperty struct {
	// The array of [Input](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_Input.html) objects describing the input streams used by the application.
	Inputs interface{} `field:"optional" json:"inputs" yaml:"inputs"`
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &vpcConfigurationProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnApplicationV2_VpcConfigurationProperty struct {
	// `CfnApplicationV2.VpcConfigurationProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `CfnApplicationV2.VpcConfigurationProperty.SubnetIds`.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

// The configuration of a Kinesis Data Analytics Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zeppelinApplicationConfigurationProperty := &zeppelinApplicationConfigurationProperty{
//   	catalogConfiguration: &catalogConfigurationProperty{
//   		glueDataCatalogConfiguration: &glueDataCatalogConfigurationProperty{
//   			databaseArn: jsii.String("databaseArn"),
//   		},
//   	},
//   	customArtifactsConfiguration: []interface{}{
//   		&customArtifactConfigurationProperty{
//   			artifactType: jsii.String("artifactType"),
//
//   			// the properties below are optional
//   			mavenReference: &mavenReferenceProperty{
//   				artifactId: jsii.String("artifactId"),
//   				groupId: jsii.String("groupId"),
//   				version: jsii.String("version"),
//   			},
//   			s3ContentLocation: &s3ContentLocationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//   				fileKey: jsii.String("fileKey"),
//
//   				// the properties below are optional
//   				objectVersion: jsii.String("objectVersion"),
//   			},
//   		},
//   	},
//   	deployAsApplicationConfiguration: &deployAsApplicationConfigurationProperty{
//   		s3ContentLocation: &s3ContentBaseLocationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//
//   			// the properties below are optional
//   			basePath: jsii.String("basePath"),
//   		},
//   	},
//   	monitoringConfiguration: &zeppelinMonitoringConfigurationProperty{
//   		logLevel: jsii.String("logLevel"),
//   	},
//   }
//
type CfnApplicationV2_ZeppelinApplicationConfigurationProperty struct {
	// The Amazon Glue Data Catalog that you use in queries in a Kinesis Data Analytics Studio notebook.
	CatalogConfiguration interface{} `field:"optional" json:"catalogConfiguration" yaml:"catalogConfiguration"`
	// A list of `CustomArtifactConfiguration` objects.
	CustomArtifactsConfiguration interface{} `field:"optional" json:"customArtifactsConfiguration" yaml:"customArtifactsConfiguration"`
	// The information required to deploy a Kinesis Data Analytics Studio notebook as an application with durable state.
	DeployAsApplicationConfiguration interface{} `field:"optional" json:"deployAsApplicationConfiguration" yaml:"deployAsApplicationConfiguration"`
	// The monitoring configuration of a Kinesis Data Analytics Studio notebook.
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
}

// Describes configuration parameters for Amazon CloudWatch logging for a Kinesis Data Analytics Studio notebook.
//
// For more information about CloudWatch logging, see [Monitoring](https://docs.aws.amazon.com/kinesisanalytics/latest/java/monitoring-overview.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zeppelinMonitoringConfigurationProperty := &zeppelinMonitoringConfigurationProperty{
//   	logLevel: jsii.String("logLevel"),
//   }
//
type CfnApplicationV2_ZeppelinMonitoringConfigurationProperty struct {
	// The verbosity of the CloudWatch Logs for an application.
	//
	// You can set it to `INFO` , `WARN` , `ERROR` , or `DEBUG` .
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

// Properties for defining a `CfnApplicationV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationV2Props := &cfnApplicationV2Props{
//   	runtimeEnvironment: jsii.String("runtimeEnvironment"),
//   	serviceExecutionRole: jsii.String("serviceExecutionRole"),
//
//   	// the properties below are optional
//   	applicationConfiguration: &applicationConfigurationProperty{
//   		applicationCodeConfiguration: &applicationCodeConfigurationProperty{
//   			codeContent: &codeContentProperty{
//   				s3ContentLocation: &s3ContentLocationProperty{
//   					bucketArn: jsii.String("bucketArn"),
//   					fileKey: jsii.String("fileKey"),
//
//   					// the properties below are optional
//   					objectVersion: jsii.String("objectVersion"),
//   				},
//   				textContent: jsii.String("textContent"),
//   				zipFileContent: jsii.String("zipFileContent"),
//   			},
//   			codeContentType: jsii.String("codeContentType"),
//   		},
//   		applicationSnapshotConfiguration: &applicationSnapshotConfigurationProperty{
//   			snapshotsEnabled: jsii.Boolean(false),
//   		},
//   		environmentProperties: &environmentPropertiesProperty{
//   			propertyGroups: []interface{}{
//   				&propertyGroupProperty{
//   					propertyGroupId: jsii.String("propertyGroupId"),
//   					propertyMap: map[string]*string{
//   						"propertyMapKey": jsii.String("propertyMap"),
//   					},
//   				},
//   			},
//   		},
//   		flinkApplicationConfiguration: &flinkApplicationConfigurationProperty{
//   			checkpointConfiguration: &checkpointConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				checkpointingEnabled: jsii.Boolean(false),
//   				checkpointInterval: jsii.Number(123),
//   				minPauseBetweenCheckpoints: jsii.Number(123),
//   			},
//   			monitoringConfiguration: &monitoringConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				logLevel: jsii.String("logLevel"),
//   				metricsLevel: jsii.String("metricsLevel"),
//   			},
//   			parallelismConfiguration: &parallelismConfigurationProperty{
//   				configurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				autoScalingEnabled: jsii.Boolean(false),
//   				parallelism: jsii.Number(123),
//   				parallelismPerKpu: jsii.Number(123),
//   			},
//   		},
//   		sqlApplicationConfiguration: &sqlApplicationConfigurationProperty{
//   			inputs: []interface{}{
//   				&inputProperty{
//   					inputSchema: &inputSchemaProperty{
//   						recordColumns: []interface{}{
//   							&recordColumnProperty{
//   								name: jsii.String("name"),
//   								sqlType: jsii.String("sqlType"),
//
//   								// the properties below are optional
//   								mapping: jsii.String("mapping"),
//   							},
//   						},
//   						recordFormat: &recordFormatProperty{
//   							recordFormatType: jsii.String("recordFormatType"),
//
//   							// the properties below are optional
//   							mappingParameters: &mappingParametersProperty{
//   								csvMappingParameters: &cSVMappingParametersProperty{
//   									recordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   									recordRowDelimiter: jsii.String("recordRowDelimiter"),
//   								},
//   								jsonMappingParameters: &jSONMappingParametersProperty{
//   									recordRowPath: jsii.String("recordRowPath"),
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						recordEncoding: jsii.String("recordEncoding"),
//   					},
//   					namePrefix: jsii.String("namePrefix"),
//
//   					// the properties below are optional
//   					inputParallelism: &inputParallelismProperty{
//   						count: jsii.Number(123),
//   					},
//   					inputProcessingConfiguration: &inputProcessingConfigurationProperty{
//   						inputLambdaProcessor: &inputLambdaProcessorProperty{
//   							resourceArn: jsii.String("resourceArn"),
//   						},
//   					},
//   					kinesisFirehoseInput: &kinesisFirehoseInputProperty{
//   						resourceArn: jsii.String("resourceArn"),
//   					},
//   					kinesisStreamsInput: &kinesisStreamsInputProperty{
//   						resourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   			},
//   		},
//   		vpcConfigurations: []interface{}{
//   			&vpcConfigurationProperty{
//   				securityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				subnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		zeppelinApplicationConfiguration: &zeppelinApplicationConfigurationProperty{
//   			catalogConfiguration: &catalogConfigurationProperty{
//   				glueDataCatalogConfiguration: &glueDataCatalogConfigurationProperty{
//   					databaseArn: jsii.String("databaseArn"),
//   				},
//   			},
//   			customArtifactsConfiguration: []interface{}{
//   				&customArtifactConfigurationProperty{
//   					artifactType: jsii.String("artifactType"),
//
//   					// the properties below are optional
//   					mavenReference: &mavenReferenceProperty{
//   						artifactId: jsii.String("artifactId"),
//   						groupId: jsii.String("groupId"),
//   						version: jsii.String("version"),
//   					},
//   					s3ContentLocation: &s3ContentLocationProperty{
//   						bucketArn: jsii.String("bucketArn"),
//   						fileKey: jsii.String("fileKey"),
//
//   						// the properties below are optional
//   						objectVersion: jsii.String("objectVersion"),
//   					},
//   				},
//   			},
//   			deployAsApplicationConfiguration: &deployAsApplicationConfigurationProperty{
//   				s3ContentLocation: &s3ContentBaseLocationProperty{
//   					bucketArn: jsii.String("bucketArn"),
//
//   					// the properties below are optional
//   					basePath: jsii.String("basePath"),
//   				},
//   			},
//   			monitoringConfiguration: &zeppelinMonitoringConfigurationProperty{
//   				logLevel: jsii.String("logLevel"),
//   			},
//   		},
//   	},
//   	applicationDescription: jsii.String("applicationDescription"),
//   	applicationMaintenanceConfiguration: &applicationMaintenanceConfigurationProperty{
//   		applicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   	},
//   	applicationMode: jsii.String("applicationMode"),
//   	applicationName: jsii.String("applicationName"),
//   	runConfiguration: &runConfigurationProperty{
//   		applicationRestoreConfiguration: &applicationRestoreConfigurationProperty{
//   			applicationRestoreType: jsii.String("applicationRestoreType"),
//
//   			// the properties below are optional
//   			snapshotName: jsii.String("snapshotName"),
//   		},
//   		flinkRunConfiguration: &flinkRunConfigurationProperty{
//   			allowNonRestoredState: jsii.Boolean(false),
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
type CfnApplicationV2Props struct {
	// The runtime environment for the application.
	RuntimeEnvironment *string `field:"required" json:"runtimeEnvironment" yaml:"runtimeEnvironment"`
	// Specifies the IAM role that the application uses to access external resources.
	ServiceExecutionRole *string `field:"required" json:"serviceExecutionRole" yaml:"serviceExecutionRole"`
	// Use this parameter to configure the application.
	ApplicationConfiguration interface{} `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// The description of the application.
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// `AWS::KinesisAnalyticsV2::Application.ApplicationMaintenanceConfiguration`.
	ApplicationMaintenanceConfiguration interface{} `field:"optional" json:"applicationMaintenanceConfiguration" yaml:"applicationMaintenanceConfiguration"`
	// To create a Kinesis Data Analytics Studio notebook, you must set the mode to `INTERACTIVE` .
	//
	// However, for a Kinesis Data Analytics for Apache Flink application, the mode is optional.
	ApplicationMode *string `field:"optional" json:"applicationMode" yaml:"applicationMode"`
	// The name of the application.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// `AWS::KinesisAnalyticsV2::Application.RunConfiguration`.
	RunConfiguration interface{} `field:"optional" json:"runConfiguration" yaml:"runConfiguration"`
	// A list of one or more tags to assign to the application.
	//
	// A tag is a key-value pair that identifies an application. Note that the maximum number of application tags includes system tags. The maximum number of user-defined application tags is 50.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

