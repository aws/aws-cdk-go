package awsdatapipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatapipeline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DataPipeline::Pipeline`.
//
// The AWS::DataPipeline::Pipeline resource specifies a data pipeline that you can use to automate the movement and transformation of data. In each pipeline, you define pipeline objects, such as activities, schedules, data nodes, and resources. For information about pipeline objects and components that you can use, see [Pipeline Object Reference](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-pipeline-objects.html) in the *AWS Data Pipeline Developer Guide* .
//
// The `AWS::DataPipeline::Pipeline` resource adds tasks, schedules, and preconditions to the specified pipeline. You can use `PutPipelineDefinition` to populate a new pipeline.
//
// `PutPipelineDefinition` also validates the configuration as it adds it to the pipeline. Changes to the pipeline are saved unless one of the following validation errors exist in the pipeline.
//
// - An object is missing a name or identifier field.
// - A string or reference field is empty.
// - The number of objects in the pipeline exceeds the allowed maximum number of objects.
// - The pipeline is in a FINISHED state.
//
// Pipeline object definitions are passed to the [PutPipelineDefinition](https://docs.aws.amazon.com/datapipeline/latest/APIReference/API_PutPipelineDefinition.html) action and returned by the [GetPipelineDefinition](https://docs.aws.amazon.com/datapipeline/latest/APIReference/API_GetPipelineDefinition.html) action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipeline := awscdk.Aws_datapipeline.NewCfnPipeline(this, jsii.String("MyCfnPipeline"), &cfnPipelineProps{
//   	name: jsii.String("name"),
//   	parameterObjects: []interface{}{
//   		&parameterObjectProperty{
//   			attributes: []interface{}{
//   				&parameterAttributeProperty{
//   					key: jsii.String("key"),
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			id: jsii.String("id"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	activate: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	parameterValues: []interface{}{
//   		&parameterValueProperty{
//   			id: jsii.String("id"),
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	pipelineObjects: []interface{}{
//   		&pipelineObjectProperty{
//   			fields: []interface{}{
//   				&fieldProperty{
//   					key: jsii.String("key"),
//
//   					// the properties below are optional
//   					refValue: jsii.String("refValue"),
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	pipelineTags: []interface{}{
//   		&pipelineTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnPipeline interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Indicates whether to validate and start the pipeline or stop an active pipeline.
	//
	// By default, the value is set to `true` .
	Activate() interface{}
	SetActivate(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the pipeline.
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
	// The name of the pipeline.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The parameter objects used with the pipeline.
	ParameterObjects() interface{}
	SetParameterObjects(val interface{})
	// The parameter values used with the pipeline.
	ParameterValues() interface{}
	SetParameterValues(val interface{})
	// The objects that define the pipeline.
	//
	// These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see [Editing Your Pipeline](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-manage-pipeline-modify-console.html) in the *AWS Data Pipeline Developer Guide* .
	PipelineObjects() interface{}
	SetPipelineObjects(val interface{})
	// A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions.
	//
	// For more information, see [Controlling Access to Pipelines and Resources](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html) in the *AWS Data Pipeline Developer Guide* .
	PipelineTags() interface{}
	SetPipelineTags(val interface{})
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

// The jsii proxy struct for CfnPipeline
type jsiiProxy_CfnPipeline struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPipeline) Activate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) ParameterObjects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) ParameterValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineObjects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) PipelineTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeline) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataPipeline::Pipeline`.
func NewCfnPipeline(scope constructs.Construct, id *string, props *CfnPipelineProps) CfnPipeline {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeline{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datapipeline.CfnPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataPipeline::Pipeline`.
func NewCfnPipeline_Override(c CfnPipeline, scope constructs.Construct, id *string, props *CfnPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datapipeline.CfnPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPipeline) SetActivate(val interface{}) {
	_jsii_.Set(
		j,
		"activate",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetParameterObjects(val interface{}) {
	_jsii_.Set(
		j,
		"parameterObjects",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetParameterValues(val interface{}) {
	_jsii_.Set(
		j,
		"parameterValues",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineObjects(val interface{}) {
	_jsii_.Set(
		j,
		"pipelineObjects",
		val,
	)
}

func (j *jsiiProxy_CfnPipeline) SetPipelineTags(val interface{}) {
	_jsii_.Set(
		j,
		"pipelineTags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPipeline_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datapipeline.CfnPipeline",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPipeline_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datapipeline.CfnPipeline",
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
func CfnPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datapipeline.CfnPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipeline_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datapipeline.CfnPipeline",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipeline) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPipeline) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPipeline) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPipeline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPipeline) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPipeline) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPipeline) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPipeline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPipeline) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipeline) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A key-value pair that describes a property of a `PipelineObject` .
//
// The value is specified as either a string value ( `StringValue` ) or a reference to another object ( `RefValue` ) but not as both. To view fields for a data pipeline object, see [Pipeline Object Reference](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-pipeline-objects.html) in the *AWS Data Pipeline Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldProperty := &fieldProperty{
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	refValue: jsii.String("refValue"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnPipeline_FieldProperty struct {
	// Specifies the name of a field for a particular object.
	//
	// To view valid values for a particular field, see [Pipeline Object Reference](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-pipeline-objects.html) in the *AWS Data Pipeline Developer Guide* .
	Key *string `field:"required" json:"key" yaml:"key"`
	// A field value that you specify as an identifier of another object in the same pipeline definition.
	//
	// > You can specify the field value as either a string value ( `StringValue` ) or a reference to another object ( `RefValue` ), but not both.
	//
	// Required if the key that you are using requires it.
	RefValue *string `field:"optional" json:"refValue" yaml:"refValue"`
	// A field value that you specify as a string.
	//
	// To view valid values for a particular field, see [Pipeline Object Reference](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-pipeline-objects.html) in the *AWS Data Pipeline Developer Guide* .
	//
	// > You can specify the field value as either a string value ( `StringValue` ) or a reference to another object ( `RefValue` ), but not both.
	//
	// Required if the key that you are using requires it.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

// `Attribute` is a property of `ParameterObject` that defines the attributes of a parameter object as key-value pairs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterAttributeProperty := &parameterAttributeProperty{
//   	key: jsii.String("key"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnPipeline_ParameterAttributeProperty struct {
	// The field identifier.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The field value, expressed as a String.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}

// Contains information about a parameter object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterObjectProperty := &parameterObjectProperty{
//   	attributes: []interface{}{
//   		&parameterAttributeProperty{
//   			key: jsii.String("key"),
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	id: jsii.String("id"),
//   }
//
type CfnPipeline_ParameterObjectProperty struct {
	// The attributes of the parameter object.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The ID of the parameter object.
	Id *string `field:"required" json:"id" yaml:"id"`
}

// A value or list of parameter values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterValueProperty := &parameterValueProperty{
//   	id: jsii.String("id"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnPipeline_ParameterValueProperty struct {
	// The ID of the parameter value.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The field value, expressed as a String.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}

// PipelineObject is property of the AWS::DataPipeline::Pipeline resource that contains information about a pipeline object.
//
// This can be a logical, physical, or physical attempt pipeline object. The complete set of components of a pipeline defines the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineObjectProperty := &pipelineObjectProperty{
//   	fields: []interface{}{
//   		&fieldProperty{
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			refValue: jsii.String("refValue"),
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   }
//
type CfnPipeline_PipelineObjectProperty struct {
	// Key-value pairs that define the properties of the object.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The ID of the object.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The name of the object.
	Name *string `field:"required" json:"name" yaml:"name"`
}

// A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions.
//
// For more information, see [Controlling Access to Pipelines and Resources](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html) in the *AWS Data Pipeline Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineTagProperty := &pipelineTagProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnPipeline_PipelineTagProperty struct {
	// The key name of a tag.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value to associate with the key name.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineProps := &cfnPipelineProps{
//   	name: jsii.String("name"),
//   	parameterObjects: []interface{}{
//   		&parameterObjectProperty{
//   			attributes: []interface{}{
//   				&parameterAttributeProperty{
//   					key: jsii.String("key"),
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			id: jsii.String("id"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	activate: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	parameterValues: []interface{}{
//   		&parameterValueProperty{
//   			id: jsii.String("id"),
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	pipelineObjects: []interface{}{
//   		&pipelineObjectProperty{
//   			fields: []interface{}{
//   				&fieldProperty{
//   					key: jsii.String("key"),
//
//   					// the properties below are optional
//   					refValue: jsii.String("refValue"),
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	pipelineTags: []interface{}{
//   		&pipelineTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// The name of the pipeline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parameter objects used with the pipeline.
	ParameterObjects interface{} `field:"required" json:"parameterObjects" yaml:"parameterObjects"`
	// Indicates whether to validate and start the pipeline or stop an active pipeline.
	//
	// By default, the value is set to `true` .
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// A description of the pipeline.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The parameter values used with the pipeline.
	ParameterValues interface{} `field:"optional" json:"parameterValues" yaml:"parameterValues"`
	// The objects that define the pipeline.
	//
	// These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see [Editing Your Pipeline](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-manage-pipeline-modify-console.html) in the *AWS Data Pipeline Developer Guide* .
	PipelineObjects interface{} `field:"optional" json:"pipelineObjects" yaml:"pipelineObjects"`
	// A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions.
	//
	// For more information, see [Controlling Access to Pipelines and Resources](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html) in the *AWS Data Pipeline Developer Guide* .
	PipelineTags interface{} `field:"optional" json:"pipelineTags" yaml:"pipelineTags"`
}

