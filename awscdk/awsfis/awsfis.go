package awsfis

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::FIS::ExperimentTemplate`.
//
// Specifies an experiment template.
//
// An experiment template includes the following components:
//
// - *Targets* : A target can be a specific resource in your AWS environment, or one or more resources that match criteria that you specify, for example, resources that have specific tags.
// - *Actions* : The actions to carry out on the target. You can specify multiple actions, the duration of each action, and when to start each action during an experiment.
// - *Stop conditions* : If a stop condition is triggered while an experiment is running, the experiment is automatically stopped. You can define a stop condition as a CloudWatch alarm.
//
// For more information, see [Experiment templates](https://docs.aws.amazon.com/fis/latest/userguide/experiment-templates.html) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fis "github.com/aws/aws-cdk-go/awscdk/aws_fis"
//
//   var cloudWatchLogsConfiguration interface{}
//   var s3Configuration interface{}
//   cfnExperimentTemplate := fis.NewCfnExperimentTemplate(this, jsii.String("MyCfnExperimentTemplate"), &cfnExperimentTemplateProps{
//   	description: jsii.String("description"),
//   	roleArn: jsii.String("roleArn"),
//   	stopConditions: []interface{}{
//   		&experimentTemplateStopConditionProperty{
//   			source: jsii.String("source"),
//
//   			// the properties below are optional
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	targets: map[string]interface{}{
//   		"targetsKey": &ExperimentTemplateTargetProperty{
//   			"resourceType": jsii.String("resourceType"),
//   			"selectionMode": jsii.String("selectionMode"),
//
//   			// the properties below are optional
//   			"filters": []interface{}{
//   				&ExperimentTemplateTargetFilterProperty{
//   					"path": jsii.String("path"),
//   					"values": []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			"parameters": map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			"resourceArns": []*string{
//   				jsii.String("resourceArns"),
//   			},
//   			"resourceTags": map[string]*string{
//   				"resourceTagsKey": jsii.String("resourceTags"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	actions: map[string]interface{}{
//   		"actionsKey": &ExperimentTemplateActionProperty{
//   			"actionId": jsii.String("actionId"),
//
//   			// the properties below are optional
//   			"description": jsii.String("description"),
//   			"parameters": map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			"startAfter": []*string{
//   				jsii.String("startAfter"),
//   			},
//   			"targets": map[string]*string{
//   				"targetsKey": jsii.String("targets"),
//   			},
//   		},
//   	},
//   	logConfiguration: &experimentTemplateLogConfigurationProperty{
//   		logSchemaVersion: jsii.Number(123),
//
//   		// the properties below are optional
//   		cloudWatchLogsConfiguration: cloudWatchLogsConfiguration,
//   		s3Configuration: s3Configuration,
//   	},
//   })
//
type CfnExperimentTemplate interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The actions for the experiment.
	Actions() interface{}
	SetActions(val interface{})
	// The ID of the experiment template.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description for the experiment template.
	Description() *string
	SetDescription(val *string)
	// The configuration for experiment logging.
	LogConfiguration() interface{}
	SetLogConfiguration(val interface{})
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
	// The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The stop conditions.
	StopConditions() interface{}
	SetStopConditions(val interface{})
	// The tags to apply to the experiment template.
	Tags() awscdk.TagManager
	// The targets for the experiment.
	Targets() interface{}
	SetTargets(val interface{})
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

// The jsii proxy struct for CfnExperimentTemplate
type jsiiProxy_CfnExperimentTemplate struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnExperimentTemplate) Actions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) LogConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) StopConditions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stopConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentTemplate) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::FIS::ExperimentTemplate`.
func NewCfnExperimentTemplate(scope constructs.Construct, id *string, props *CfnExperimentTemplateProps) CfnExperimentTemplate {
	_init_.Initialize()

	j := jsiiProxy_CfnExperimentTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::FIS::ExperimentTemplate`.
func NewCfnExperimentTemplate_Override(c CfnExperimentTemplate, scope constructs.Construct, id *string, props *CfnExperimentTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetActions(val interface{}) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetLogConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"logConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetStopConditions(val interface{}) {
	_jsii_.Set(
		j,
		"stopConditions",
		val,
	)
}

func (j *jsiiProxy_CfnExperimentTemplate) SetTargets(val interface{}) {
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
func CfnExperimentTemplate_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnExperimentTemplate_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
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
func CfnExperimentTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExperimentTemplate_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_fis.CfnExperimentTemplate",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExperimentTemplate) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperimentTemplate) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperimentTemplate) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnExperimentTemplate) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperimentTemplate) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperimentTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperimentTemplate) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies an action for an experiment template.
//
// For more information, see [Actions](https://docs.aws.amazon.com/fis/latest/userguide/actions.html) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fis "github.com/aws/aws-cdk-go/awscdk/aws_fis"
//   experimentTemplateActionProperty := &experimentTemplateActionProperty{
//   	actionId: jsii.String("actionId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	startAfter: []*string{
//   		jsii.String("startAfter"),
//   	},
//   	targets: map[string]*string{
//   		"targetsKey": jsii.String("targets"),
//   	},
//   }
//
type CfnExperimentTemplate_ExperimentTemplateActionProperty struct {
	// The ID of the action.
	//
	// The format of the action ID is: aws: *service-name* : *action-type* .
	ActionId *string `json:"actionId" yaml:"actionId"`
	// A description for the action.
	Description *string `json:"description" yaml:"description"`
	// The parameters for the action, if applicable.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// The name of the action that must be completed before the current action starts.
	//
	// Omit this parameter to run the action at the start of the experiment.
	StartAfter *[]*string `json:"startAfter" yaml:"startAfter"`
	// The targets for the action.
	Targets interface{} `json:"targets" yaml:"targets"`
}

// Specifies the configuration for experiment logging.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fis "github.com/aws/aws-cdk-go/awscdk/aws_fis"
//
//   var cloudWatchLogsConfiguration interface{}
//   var s3Configuration interface{}
//   experimentTemplateLogConfigurationProperty := &experimentTemplateLogConfigurationProperty{
//   	logSchemaVersion: jsii.Number(123),
//
//   	// the properties below are optional
//   	cloudWatchLogsConfiguration: cloudWatchLogsConfiguration,
//   	s3Configuration: s3Configuration,
//   }
//
type CfnExperimentTemplate_ExperimentTemplateLogConfigurationProperty struct {
	// The schema version.
	//
	// The supported value is 1.
	LogSchemaVersion *float64 `json:"logSchemaVersion" yaml:"logSchemaVersion"`
	// The configuration for experiment logging to Amazon CloudWatch Logs. The supported field is `logGroupArn` . For example:.
	//
	// `{"logGroupArn": "aws:arn:logs: *region_name* : *account_id* :log-group: *log_group_name* "}`.
	CloudWatchLogsConfiguration interface{} `json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// The configuration for experiment logging to Amazon S3. The following fields are supported:.
	//
	// - `bucketName` - The name of the destination bucket.
	// - `prefix` - An optional bucket prefix.
	//
	// For example:
	//
	// `{"bucketName": " *my-s3-bucket* ", "prefix": " *log-folder* "}`.
	S3Configuration interface{} `json:"s3Configuration" yaml:"s3Configuration"`
}

// Specifies a stop condition for an experiment template.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fis "github.com/aws/aws-cdk-go/awscdk/aws_fis"
//   experimentTemplateStopConditionProperty := &experimentTemplateStopConditionProperty{
//   	source: jsii.String("source"),
//
//   	// the properties below are optional
//   	value: jsii.String("value"),
//   }
//
type CfnExperimentTemplate_ExperimentTemplateStopConditionProperty struct {
	// The source for the stop condition.
	//
	// Specify `aws:cloudwatch:alarm` if the stop condition is defined by a CloudWatch alarm. Specify `none` if there is no stop condition.
	Source *string `json:"source" yaml:"source"`
	// The Amazon Resource Name (ARN) of the CloudWatch alarm.
	//
	// This is required if the source is a CloudWatch alarm.
	Value *string `json:"value" yaml:"value"`
}

// Specifies a filter used for the target resource input in an experiment template.
//
// For more information, see [Resource filters](https://docs.aws.amazon.com/fis/latest/userguide/targets.html#target-filters) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fis "github.com/aws/aws-cdk-go/awscdk/aws_fis"
//   experimentTemplateTargetFilterProperty := &experimentTemplateTargetFilterProperty{
//   	path: jsii.String("path"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnExperimentTemplate_ExperimentTemplateTargetFilterProperty struct {
	// The attribute path for the filter.
	Path *string `json:"path" yaml:"path"`
	// The attribute values for the filter.
	Values *[]*string `json:"values" yaml:"values"`
}

// Specifies a target for an experiment.
//
// You must specify at least one Amazon Resource Name (ARN) or at least one resource tag. You cannot specify both ARNs and tags.
//
// For more information, see [Targets](https://docs.aws.amazon.com/fis/latest/userguide/targets.html) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fis "github.com/aws/aws-cdk-go/awscdk/aws_fis"
//   experimentTemplateTargetProperty := &experimentTemplateTargetProperty{
//   	resourceType: jsii.String("resourceType"),
//   	selectionMode: jsii.String("selectionMode"),
//
//   	// the properties below are optional
//   	filters: []interface{}{
//   		&experimentTemplateTargetFilterProperty{
//   			path: jsii.String("path"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	resourceTags: map[string]*string{
//   		"resourceTagsKey": jsii.String("resourceTags"),
//   	},
//   }
//
type CfnExperimentTemplate_ExperimentTemplateTargetProperty struct {
	// The resource type.
	//
	// The resource type must be supported for the specified action.
	ResourceType *string `json:"resourceType" yaml:"resourceType"`
	// Scopes the identified resources to a specific count of the resources at random, or a percentage of the resources.
	//
	// All identified resources are included in the target.
	//
	// - ALL - Run the action on all identified targets. This is the default.
	// - COUNT(n) - Run the action on the specified number of targets, chosen from the identified targets at random. For example, COUNT(1) selects one of the targets.
	// - PERCENT(n) - Run the action on the specified percentage of targets, chosen from the identified targets at random. For example, PERCENT(25) selects 25% of the targets.
	SelectionMode *string `json:"selectionMode" yaml:"selectionMode"`
	// The filters to apply to identify target resources using specific attributes.
	Filters interface{} `json:"filters" yaml:"filters"`
	// The parameters for the resource type.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// The Amazon Resource Names (ARNs) of the resources.
	ResourceArns *[]*string `json:"resourceArns" yaml:"resourceArns"`
	// The tags for the target resources.
	ResourceTags interface{} `json:"resourceTags" yaml:"resourceTags"`
}

// Properties for defining a `CfnExperimentTemplate`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import fis "github.com/aws/aws-cdk-go/awscdk/aws_fis"
//
//   var cloudWatchLogsConfiguration interface{}
//   var s3Configuration interface{}
//   cfnExperimentTemplateProps := &cfnExperimentTemplateProps{
//   	description: jsii.String("description"),
//   	roleArn: jsii.String("roleArn"),
//   	stopConditions: []interface{}{
//   		&experimentTemplateStopConditionProperty{
//   			source: jsii.String("source"),
//
//   			// the properties below are optional
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	targets: map[string]interface{}{
//   		"targetsKey": &ExperimentTemplateTargetProperty{
//   			"resourceType": jsii.String("resourceType"),
//   			"selectionMode": jsii.String("selectionMode"),
//
//   			// the properties below are optional
//   			"filters": []interface{}{
//   				&ExperimentTemplateTargetFilterProperty{
//   					"path": jsii.String("path"),
//   					"values": []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			"parameters": map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			"resourceArns": []*string{
//   				jsii.String("resourceArns"),
//   			},
//   			"resourceTags": map[string]*string{
//   				"resourceTagsKey": jsii.String("resourceTags"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	actions: map[string]interface{}{
//   		"actionsKey": &ExperimentTemplateActionProperty{
//   			"actionId": jsii.String("actionId"),
//
//   			// the properties below are optional
//   			"description": jsii.String("description"),
//   			"parameters": map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			"startAfter": []*string{
//   				jsii.String("startAfter"),
//   			},
//   			"targets": map[string]*string{
//   				"targetsKey": jsii.String("targets"),
//   			},
//   		},
//   	},
//   	logConfiguration: &experimentTemplateLogConfigurationProperty{
//   		logSchemaVersion: jsii.Number(123),
//
//   		// the properties below are optional
//   		cloudWatchLogsConfiguration: cloudWatchLogsConfiguration,
//   		s3Configuration: s3Configuration,
//   	},
//   }
//
type CfnExperimentTemplateProps struct {
	// A description for the experiment template.
	Description *string `json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The stop conditions.
	StopConditions interface{} `json:"stopConditions" yaml:"stopConditions"`
	// The tags to apply to the experiment template.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// The targets for the experiment.
	Targets interface{} `json:"targets" yaml:"targets"`
	// The actions for the experiment.
	Actions interface{} `json:"actions" yaml:"actions"`
	// The configuration for experiment logging.
	LogConfiguration interface{} `json:"logConfiguration" yaml:"logConfiguration"`
}

