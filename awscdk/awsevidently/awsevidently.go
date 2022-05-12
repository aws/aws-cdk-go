package awsevidently

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Evidently::Experiment`.
//
// Creates or updates an Evidently *experiment* . Before you create an experiment, you must create the feature to use for the experiment.
//
// An experiment helps you make feature design decisions based on evidence and data. An experiment can test as many as five variations at once. Evidently collects experiment data and analyzes it by statistical methods, and provides clear recommendations about which variations perform better.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExperiment := awscdk.Aws_evidently.NewCfnExperiment(this, jsii.String("MyCfnExperiment"), &cfnExperimentProps{
//   	metricGoals: []interface{}{
//   		&metricGoalObjectProperty{
//   			desiredChange: jsii.String("desiredChange"),
//   			entityIdKey: jsii.String("entityIdKey"),
//   			eventPattern: jsii.String("eventPattern"),
//   			metricName: jsii.String("metricName"),
//   			valueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			unitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	onlineAbConfig: &onlineAbConfigObjectProperty{
//   		controlTreatmentName: jsii.String("controlTreatmentName"),
//   		treatmentWeights: []interface{}{
//   			&treatmentToWeightProperty{
//   				splitWeight: jsii.Number(123),
//   				treatment: jsii.String("treatment"),
//   			},
//   		},
//   	},
//   	project: jsii.String("project"),
//   	treatments: []interface{}{
//   		&treatmentObjectProperty{
//   			feature: jsii.String("feature"),
//   			treatmentName: jsii.String("treatmentName"),
//   			variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	randomizationSalt: jsii.String("randomizationSalt"),
//   	runningStatus: &runningStatusObjectProperty{
//   		analysisCompleteTime: jsii.String("analysisCompleteTime"),
//   		desiredState: jsii.String("desiredState"),
//   		reason: jsii.String("reason"),
//   		status: jsii.String("status"),
//   	},
//   	samplingRate: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnExperiment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the experiment.
	//
	// For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/experiment/myExperiment`.
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
	// An optional description of the experiment.
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
	// An array of structures that defines the metrics used for the experiment, and whether a higher or lower value for each metric is the goal.
	//
	// You can use up to three metrics in an experiment.
	MetricGoals() interface{}
	SetMetricGoals(val interface{})
	// A name for the new experiment.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A structure that contains the configuration of which variation to use as the "control" version.
	//
	// The "control" version is used for comparison with other variations. This structure also specifies how much experiment traffic is allocated to each variation.
	OnlineAbConfig() interface{}
	SetOnlineAbConfig(val interface{})
	// The name or the ARN of the project where this experiment is to be created.
	Project() *string
	SetProject(val *string)
	// When Evidently assigns a particular user session to an experiment, it must use a randomization ID to determine which variation the user session is served.
	//
	// This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the experiment name as the `randomizationSalt` .
	RandomizationSalt() *string
	SetRandomizationSalt(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A structure that you can use to start and stop the experiment.
	RunningStatus() interface{}
	SetRunningStatus(val interface{})
	// The portion of the available audience that you want to allocate to this experiment, in thousandths of a percent.
	//
	// The available audience is the total audience minus the audience that you have allocated to overrides or current launches of this feature.
	//
	// This is represented in thousandths of a percent. For example, specify 10,000 to allocate 10% of the available audience.
	SamplingRate() *float64
	SetSamplingRate(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Assigns one or more tags (key-value pairs) to the experiment.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with an experiment.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags() awscdk.TagManager
	// An array of structures that describe the configuration of each feature variation used in the experiment.
	Treatments() interface{}
	SetTreatments(val interface{})
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

// The jsii proxy struct for CfnExperiment
type jsiiProxy_CfnExperiment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnExperiment) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) MetricGoals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricGoals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) OnlineAbConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlineAbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) RandomizationSalt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"randomizationSalt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) RunningStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runningStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) SamplingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) Treatments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperiment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Evidently::Experiment`.
func NewCfnExperiment(scope constructs.Construct, id *string, props *CfnExperimentProps) CfnExperiment {
	_init_.Initialize()

	j := jsiiProxy_CfnExperiment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_evidently.CfnExperiment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Evidently::Experiment`.
func NewCfnExperiment_Override(c CfnExperiment, scope constructs.Construct, id *string, props *CfnExperimentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_evidently.CfnExperiment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnExperiment) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment) SetMetricGoals(val interface{}) {
	_jsii_.Set(
		j,
		"metricGoals",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment) SetOnlineAbConfig(val interface{}) {
	_jsii_.Set(
		j,
		"onlineAbConfig",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment) SetProject(val *string) {
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment) SetRandomizationSalt(val *string) {
	_jsii_.Set(
		j,
		"randomizationSalt",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment) SetRunningStatus(val interface{}) {
	_jsii_.Set(
		j,
		"runningStatus",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment) SetSamplingRate(val *float64) {
	_jsii_.Set(
		j,
		"samplingRate",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment) SetTreatments(val interface{}) {
	_jsii_.Set(
		j,
		"treatments",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnExperiment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnExperiment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnExperiment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnExperiment",
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
func CfnExperiment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnExperiment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExperiment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_evidently.CfnExperiment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExperiment) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnExperiment) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnExperiment) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnExperiment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnExperiment) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnExperiment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnExperiment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnExperiment) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperiment) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperiment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnExperiment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnExperiment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperiment) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperiment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperiment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Use this structure to tell Evidently whether higher or lower values are desired for a metric that is used in an experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricGoalObjectProperty := &metricGoalObjectProperty{
//   	desiredChange: jsii.String("desiredChange"),
//   	entityIdKey: jsii.String("entityIdKey"),
//   	eventPattern: jsii.String("eventPattern"),
//   	metricName: jsii.String("metricName"),
//   	valueKey: jsii.String("valueKey"),
//
//   	// the properties below are optional
//   	unitLabel: jsii.String("unitLabel"),
//   }
//
type CfnExperiment_MetricGoalObjectProperty struct {
	// `INCREASE` means that a variation with a higher number for this metric is performing better.
	//
	// `DECREASE` means that a variation with a lower number for this metric is performing better.
	DesiredChange *string `field:"required" json:"desiredChange" yaml:"desiredChange"`
	// The entity, such as a user or session, that does an action that causes a metric value to be recorded.
	//
	// An example is `userDetails.userID` .
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// The EventBridge event pattern that defines how the metric is recorded.
	//
	// For more information about EventBridge event patterns, see [Amazon EventBridge event patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html) .
	EventPattern *string `field:"required" json:"eventPattern" yaml:"eventPattern"`
	// A name for the metric.
	//
	// It can include up to 255 characters.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The JSON path to reference the numerical metric value in the event.
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// A label for the units that the metric is measuring.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

// A structure that contains the configuration of which variation to use as the "control" version.
//
// The "control" version is used for comparison with other variations. This structure also specifies how much experiment traffic is allocated to each variation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onlineAbConfigObjectProperty := &onlineAbConfigObjectProperty{
//   	controlTreatmentName: jsii.String("controlTreatmentName"),
//   	treatmentWeights: []interface{}{
//   		&treatmentToWeightProperty{
//   			splitWeight: jsii.Number(123),
//   			treatment: jsii.String("treatment"),
//   		},
//   	},
//   }
//
type CfnExperiment_OnlineAbConfigObjectProperty struct {
	// The name of the variation that is to be the default variation that the other variations are compared to.
	ControlTreatmentName *string `field:"optional" json:"controlTreatmentName" yaml:"controlTreatmentName"`
	// A set of key-value pairs.
	//
	// The keys are treatment names, and the values are the portion of experiment traffic to be assigned to that treatment. Specify the traffic portion in thousandths of a percent, so 20,000 for a variation would allocate 20% of the experiment traffic to that variation.
	TreatmentWeights interface{} `field:"optional" json:"treatmentWeights" yaml:"treatmentWeights"`
}

// Use this structure to start and stop the experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runningStatusObjectProperty := &runningStatusObjectProperty{
//   	analysisCompleteTime: jsii.String("analysisCompleteTime"),
//   	desiredState: jsii.String("desiredState"),
//   	reason: jsii.String("reason"),
//   	status: jsii.String("status"),
//   }
//
type CfnExperiment_RunningStatusObjectProperty struct {
	// If you are using AWS CloudFormation to start the experiment, use this field to specify when the experiment is to end.
	//
	// The format is as a UNIX timestamp. For more information about this format, see [The Current Epoch Unix Timestamp](https://docs.aws.amazon.com/https://www.unixtimestamp.com/index.php) .
	AnalysisCompleteTime *string `field:"optional" json:"analysisCompleteTime" yaml:"analysisCompleteTime"`
	// If you are using AWS CloudFormation to stop this experiment, specify either `COMPLETED` or `CANCELLED` here to indicate how to classify this experiment.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// If you are using AWS CloudFormation to stop this experiment, this is an optional field that you can use to record why the experiment is being stopped or cancelled.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// To start the experiment now, specify `START` for this parameter.
	//
	// If this experiment is currently running and you want to stop it now, specify `STOP` .
	Status *string `field:"optional" json:"status" yaml:"status"`
}

// A structure that defines one treatment in an experiment.
//
// A treatment is a variation of the feature that you are including in the experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   treatmentObjectProperty := &treatmentObjectProperty{
//   	feature: jsii.String("feature"),
//   	treatmentName: jsii.String("treatmentName"),
//   	variation: jsii.String("variation"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnExperiment_TreatmentObjectProperty struct {
	// The name of the feature for this experiment.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// A name for this treatment.
	//
	// It can include up to 127 characters.
	TreatmentName *string `field:"required" json:"treatmentName" yaml:"treatmentName"`
	// The name of the variation to use for this treatment.
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// The description of the treatment.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// This structure defines how much experiment traffic to allocate to one treatment used in the experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   treatmentToWeightProperty := &treatmentToWeightProperty{
//   	splitWeight: jsii.Number(123),
//   	treatment: jsii.String("treatment"),
//   }
//
type CfnExperiment_TreatmentToWeightProperty struct {
	// The portion of experiment traffic to allocate to this treatment.
	//
	// Specify the traffic portion in thousandths of a percent, so 20,000 allocated to a treatment would allocate 20% of the experiment traffic to that treatment.
	SplitWeight *float64 `field:"required" json:"splitWeight" yaml:"splitWeight"`
	// The name of the treatment.
	Treatment *string `field:"required" json:"treatment" yaml:"treatment"`
}

// Properties for defining a `CfnExperiment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExperimentProps := &cfnExperimentProps{
//   	metricGoals: []interface{}{
//   		&metricGoalObjectProperty{
//   			desiredChange: jsii.String("desiredChange"),
//   			entityIdKey: jsii.String("entityIdKey"),
//   			eventPattern: jsii.String("eventPattern"),
//   			metricName: jsii.String("metricName"),
//   			valueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			unitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	onlineAbConfig: &onlineAbConfigObjectProperty{
//   		controlTreatmentName: jsii.String("controlTreatmentName"),
//   		treatmentWeights: []interface{}{
//   			&treatmentToWeightProperty{
//   				splitWeight: jsii.Number(123),
//   				treatment: jsii.String("treatment"),
//   			},
//   		},
//   	},
//   	project: jsii.String("project"),
//   	treatments: []interface{}{
//   		&treatmentObjectProperty{
//   			feature: jsii.String("feature"),
//   			treatmentName: jsii.String("treatmentName"),
//   			variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	randomizationSalt: jsii.String("randomizationSalt"),
//   	runningStatus: &runningStatusObjectProperty{
//   		analysisCompleteTime: jsii.String("analysisCompleteTime"),
//   		desiredState: jsii.String("desiredState"),
//   		reason: jsii.String("reason"),
//   		status: jsii.String("status"),
//   	},
//   	samplingRate: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnExperimentProps struct {
	// An array of structures that defines the metrics used for the experiment, and whether a higher or lower value for each metric is the goal.
	//
	// You can use up to three metrics in an experiment.
	MetricGoals interface{} `field:"required" json:"metricGoals" yaml:"metricGoals"`
	// A name for the new experiment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A structure that contains the configuration of which variation to use as the "control" version.
	//
	// The "control" version is used for comparison with other variations. This structure also specifies how much experiment traffic is allocated to each variation.
	OnlineAbConfig interface{} `field:"required" json:"onlineAbConfig" yaml:"onlineAbConfig"`
	// The name or the ARN of the project where this experiment is to be created.
	Project *string `field:"required" json:"project" yaml:"project"`
	// An array of structures that describe the configuration of each feature variation used in the experiment.
	Treatments interface{} `field:"required" json:"treatments" yaml:"treatments"`
	// An optional description of the experiment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When Evidently assigns a particular user session to an experiment, it must use a randomization ID to determine which variation the user session is served.
	//
	// This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the experiment name as the `randomizationSalt` .
	RandomizationSalt *string `field:"optional" json:"randomizationSalt" yaml:"randomizationSalt"`
	// A structure that you can use to start and stop the experiment.
	RunningStatus interface{} `field:"optional" json:"runningStatus" yaml:"runningStatus"`
	// The portion of the available audience that you want to allocate to this experiment, in thousandths of a percent.
	//
	// The available audience is the total audience minus the audience that you have allocated to overrides or current launches of this feature.
	//
	// This is represented in thousandths of a percent. For example, specify 10,000 to allocate 10% of the available audience.
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
	// Assigns one or more tags (key-value pairs) to the experiment.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with an experiment.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Evidently::Feature`.
//
// Creates or updates an Evidently *feature* that you want to launch or test. You can define up to five variations of a feature, and use these variations in your launches and experiments. A feature must be created in a project. For information about creating a project, see [CreateProject](https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_CreateProject.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFeature := awscdk.Aws_evidently.NewCfnFeature(this, jsii.String("MyCfnFeature"), &cfnFeatureProps{
//   	name: jsii.String("name"),
//   	project: jsii.String("project"),
//   	variations: []interface{}{
//   		&variationObjectProperty{
//   			booleanValue: jsii.Boolean(false),
//   			doubleValue: jsii.Number(123),
//   			longValue: jsii.Number(123),
//   			stringValue: jsii.String("stringValue"),
//   			variationName: jsii.String("variationName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	defaultVariation: jsii.String("defaultVariation"),
//   	description: jsii.String("description"),
//   	entityOverrides: []interface{}{
//   		&entityOverrideProperty{
//   			entityId: jsii.String("entityId"),
//   			variation: jsii.String("variation"),
//   		},
//   	},
//   	evaluationStrategy: jsii.String("evaluationStrategy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFeature interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the feature.
	//
	// For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/feature/myFeature` .
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
	// The name of the variation to use as the default variation.
	//
	// The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature.
	//
	// This variation must also be listed in the `Variations` structure.
	//
	// If you omit `DefaultVariation` , the first variation listed in the `Variations` structure is used as the default variation.
	DefaultVariation() *string
	SetDefaultVariation(val *string)
	// An optional description of the feature.
	Description() *string
	SetDescription(val *string)
	// Specify users that should always be served a specific variation of a feature.
	//
	// Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
	EntityOverrides() interface{}
	SetEntityOverrides(val interface{})
	// Specify `ALL_RULES` to activate the traffic allocation specified by any ongoing launches or experiments.
	//
	// Specify `DEFAULT_VARIATION` to serve the default variation to all users instead.
	EvaluationStrategy() *string
	SetEvaluationStrategy(val *string)
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
	// The name for the feature.
	//
	// It can include up to 127 characters.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The name or ARN of the project that is to contain the new feature.
	Project() *string
	SetProject(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Assigns one or more tags (key-value pairs) to the feature.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a feature.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// An array of structures that contain the configuration of the feature's different variations.
	//
	// Each `VariationObject` in the `Variations` array for a feature must have the same type of value ( `BooleanValue` , `DoubleValue` , `LongValue` or `StringValue` ).
	Variations() interface{}
	SetVariations(val interface{})
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

// The jsii proxy struct for CfnFeature
type jsiiProxy_CfnFeature struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFeature) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) DefaultVariation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVariation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) EntityOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) EvaluationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeature) Variations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variations",
		&returns,
	)
	return returns
}


// Create a new `AWS::Evidently::Feature`.
func NewCfnFeature(scope constructs.Construct, id *string, props *CfnFeatureProps) CfnFeature {
	_init_.Initialize()

	j := jsiiProxy_CfnFeature{}

	_jsii_.Create(
		"aws-cdk-lib.aws_evidently.CfnFeature",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Evidently::Feature`.
func NewCfnFeature_Override(c CfnFeature, scope constructs.Construct, id *string, props *CfnFeatureProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_evidently.CfnFeature",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFeature) SetDefaultVariation(val *string) {
	_jsii_.Set(
		j,
		"defaultVariation",
		val,
	)
}

func (j *jsiiProxy_CfnFeature) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFeature) SetEntityOverrides(val interface{}) {
	_jsii_.Set(
		j,
		"entityOverrides",
		val,
	)
}

func (j *jsiiProxy_CfnFeature) SetEvaluationStrategy(val *string) {
	_jsii_.Set(
		j,
		"evaluationStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnFeature) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFeature) SetProject(val *string) {
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CfnFeature) SetVariations(val interface{}) {
	_jsii_.Set(
		j,
		"variations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFeature_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnFeature",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFeature_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnFeature",
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
func CfnFeature_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnFeature",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFeature_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_evidently.CfnFeature",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFeature) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFeature) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFeature) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFeature) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFeature) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFeature) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFeature) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFeature) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeature) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeature) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFeature) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFeature) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeature) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeature) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFeature) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A set of key-value pairs that specify users who should always be served a specific variation of a feature.
//
// Each key specifies a user using their user ID, account ID, or some other identifier. The value specifies the name of the variation that the user is to be served.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityOverrideProperty := &entityOverrideProperty{
//   	entityId: jsii.String("entityId"),
//   	variation: jsii.String("variation"),
//   }
//
type CfnFeature_EntityOverrideProperty struct {
	// The entity ID to be served the variation specified in `Variation` .
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The name of the variation to serve to the user session that matches the `EntityId` .
	Variation *string `field:"optional" json:"variation" yaml:"variation"`
}

// This structure contains the name and variation value of one variation of a feature.
//
// It can contain only one of the following parameters: `BooleanValue` , `DoubleValue` , `LongValue` or `StringValue` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variationObjectProperty := &variationObjectProperty{
//   	booleanValue: jsii.Boolean(false),
//   	doubleValue: jsii.Number(123),
//   	longValue: jsii.Number(123),
//   	stringValue: jsii.String("stringValue"),
//   	variationName: jsii.String("variationName"),
//   }
//
type CfnFeature_VariationObjectProperty struct {
	// The value assigned to this variation, if the variation type is boolean.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// The value assigned to this variation, if the variation type is a double.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The value assigned to this variation, if the variation type is a long.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// The value assigned to this variation, if the variation type is a string.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// A name for the variation.
	//
	// It can include up to 127 characters.
	VariationName *string `field:"optional" json:"variationName" yaml:"variationName"`
}

// Properties for defining a `CfnFeature`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFeatureProps := &cfnFeatureProps{
//   	name: jsii.String("name"),
//   	project: jsii.String("project"),
//   	variations: []interface{}{
//   		&variationObjectProperty{
//   			booleanValue: jsii.Boolean(false),
//   			doubleValue: jsii.Number(123),
//   			longValue: jsii.Number(123),
//   			stringValue: jsii.String("stringValue"),
//   			variationName: jsii.String("variationName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	defaultVariation: jsii.String("defaultVariation"),
//   	description: jsii.String("description"),
//   	entityOverrides: []interface{}{
//   		&entityOverrideProperty{
//   			entityId: jsii.String("entityId"),
//   			variation: jsii.String("variation"),
//   		},
//   	},
//   	evaluationStrategy: jsii.String("evaluationStrategy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFeatureProps struct {
	// The name for the feature.
	//
	// It can include up to 127 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name or ARN of the project that is to contain the new feature.
	Project *string `field:"required" json:"project" yaml:"project"`
	// An array of structures that contain the configuration of the feature's different variations.
	//
	// Each `VariationObject` in the `Variations` array for a feature must have the same type of value ( `BooleanValue` , `DoubleValue` , `LongValue` or `StringValue` ).
	Variations interface{} `field:"required" json:"variations" yaml:"variations"`
	// The name of the variation to use as the default variation.
	//
	// The default variation is served to users who are not allocated to any ongoing launches or experiments of this feature.
	//
	// This variation must also be listed in the `Variations` structure.
	//
	// If you omit `DefaultVariation` , the first variation listed in the `Variations` structure is used as the default variation.
	DefaultVariation *string `field:"optional" json:"defaultVariation" yaml:"defaultVariation"`
	// An optional description of the feature.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specify users that should always be served a specific variation of a feature.
	//
	// Each user is specified by a key-value pair . For each key, specify a user by entering their user ID, account ID, or some other identifier. For the value, specify the name of the variation that they are to be served.
	EntityOverrides interface{} `field:"optional" json:"entityOverrides" yaml:"entityOverrides"`
	// Specify `ALL_RULES` to activate the traffic allocation specified by any ongoing launches or experiments.
	//
	// Specify `DEFAULT_VARIATION` to serve the default variation to all users instead.
	EvaluationStrategy *string `field:"optional" json:"evaluationStrategy" yaml:"evaluationStrategy"`
	// Assigns one or more tags (key-value pairs) to the feature.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a feature.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Evidently::Launch`.
//
// Creates or updates a *launch* of a given feature. Before you create a launch, you must create the feature to use for the launch.
//
// You can use a launch to safely validate new features by serving them to a specified percentage of your users while you roll out the feature. You can monitor the performance of the new feature to help you decide when to ramp up traffic to more users. This helps you reduce risk and identify unintended consequences before you fully launch the feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunch := awscdk.Aws_evidently.NewCfnLaunch(this, jsii.String("MyCfnLaunch"), &cfnLaunchProps{
//   	groups: []interface{}{
//   		&launchGroupObjectProperty{
//   			feature: jsii.String("feature"),
//   			groupName: jsii.String("groupName"),
//   			variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	project: jsii.String("project"),
//   	scheduledSplitsConfig: []interface{}{
//   		&stepConfigProperty{
//   			groupWeights: []interface{}{
//   				&groupToWeightProperty{
//   					groupName: jsii.String("groupName"),
//   					splitWeight: jsii.Number(123),
//   				},
//   			},
//   			startTime: jsii.String("startTime"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	executionStatus: &executionStatusObjectProperty{
//   		status: jsii.String("status"),
//
//   		// the properties below are optional
//   		desiredState: jsii.String("desiredState"),
//   		reason: jsii.String("reason"),
//   	},
//   	metricMonitors: []interface{}{
//   		&metricDefinitionObjectProperty{
//   			entityIdKey: jsii.String("entityIdKey"),
//   			eventPattern: jsii.String("eventPattern"),
//   			metricName: jsii.String("metricName"),
//   			valueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			unitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	randomizationSalt: jsii.String("randomizationSalt"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnLaunch interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the launch.
	//
	// For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/launch/myLaunch`.
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
	// An optional description for the launch.
	Description() *string
	SetDescription(val *string)
	// A structure that you can use to start and stop the launch.
	ExecutionStatus() interface{}
	SetExecutionStatus(val interface{})
	// An array of structures that contains the feature and variations that are to be used for the launch.
	//
	// You can up to five launch groups in a launch.
	Groups() interface{}
	SetGroups(val interface{})
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
	// An array of structures that define the metrics that will be used to monitor the launch performance.
	//
	// You can have up to three metric monitors in the array.
	MetricMonitors() interface{}
	SetMetricMonitors(val interface{})
	// The name for the launch.
	//
	// It can include up to 127 characters.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The name or ARN of the project that you want to create the launch in.
	Project() *string
	SetProject(val *string)
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served.
	//
	// This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the launch name as the `randomizationsSalt` .
	RandomizationSalt() *string
	SetRandomizationSalt(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// An array of structures that define the traffic allocation percentages among the feature variations during each step of the launch.
	ScheduledSplitsConfig() interface{}
	SetScheduledSplitsConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Assigns one or more tags (key-value pairs) to the launch.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a launch.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
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

// The jsii proxy struct for CfnLaunch
type jsiiProxy_CfnLaunch struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLaunch) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) ExecutionStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) Groups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) MetricMonitors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricMonitors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) RandomizationSalt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"randomizationSalt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) ScheduledSplitsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledSplitsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLaunch) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Evidently::Launch`.
func NewCfnLaunch(scope constructs.Construct, id *string, props *CfnLaunchProps) CfnLaunch {
	_init_.Initialize()

	j := jsiiProxy_CfnLaunch{}

	_jsii_.Create(
		"aws-cdk-lib.aws_evidently.CfnLaunch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Evidently::Launch`.
func NewCfnLaunch_Override(c CfnLaunch, scope constructs.Construct, id *string, props *CfnLaunchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_evidently.CfnLaunch",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLaunch) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLaunch) SetExecutionStatus(val interface{}) {
	_jsii_.Set(
		j,
		"executionStatus",
		val,
	)
}

func (j *jsiiProxy_CfnLaunch) SetGroups(val interface{}) {
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_CfnLaunch) SetMetricMonitors(val interface{}) {
	_jsii_.Set(
		j,
		"metricMonitors",
		val,
	)
}

func (j *jsiiProxy_CfnLaunch) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnLaunch) SetProject(val *string) {
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CfnLaunch) SetRandomizationSalt(val *string) {
	_jsii_.Set(
		j,
		"randomizationSalt",
		val,
	)
}

func (j *jsiiProxy_CfnLaunch) SetScheduledSplitsConfig(val interface{}) {
	_jsii_.Set(
		j,
		"scheduledSplitsConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLaunch_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnLaunch",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLaunch_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnLaunch",
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
func CfnLaunch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnLaunch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLaunch_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_evidently.CfnLaunch",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLaunch) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLaunch) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLaunch) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLaunch) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLaunch) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLaunch) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLaunch) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLaunch) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunch) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunch) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLaunch) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLaunch) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunch) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLaunch) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Use this structure to start and stop the launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionStatusObjectProperty := &executionStatusObjectProperty{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	desiredState: jsii.String("desiredState"),
//   	reason: jsii.String("reason"),
//   }
//
type CfnLaunch_ExecutionStatusObjectProperty struct {
	// To start the launch now, specify `START` for this parameter.
	//
	// If this launch is currently running and you want to stop it now, specify `STOP` .
	Status *string `field:"required" json:"status" yaml:"status"`
	// If you are using AWS CloudFormation to stop this launch, specify either `COMPLETED` or `CANCELLED` here to indicate how to classify this experiment.
	//
	// If you omit this parameter, the default of `COMPLETED` is used.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// If you are using AWS CloudFormation to stop this launch, this is an optional field that you can use to record why the launch is being stopped or cancelled.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

// A structure containing the percentage of launch traffic to allocate to one launch group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupToWeightProperty := &groupToWeightProperty{
//   	groupName: jsii.String("groupName"),
//   	splitWeight: jsii.Number(123),
//   }
//
type CfnLaunch_GroupToWeightProperty struct {
	// The name of the launch group.
	//
	// It can include up to 127 characters.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The portion of launch traffic to allocate to this launch group.
	//
	// This is represented in thousandths of a percent. For example, specify 20,000 to allocate 20% of the launch audience to this launch group.
	SplitWeight *float64 `field:"required" json:"splitWeight" yaml:"splitWeight"`
}

// A structure that defines one launch group in a launch.
//
// A launch group is a variation of the feature that you are including in the launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchGroupObjectProperty := &launchGroupObjectProperty{
//   	feature: jsii.String("feature"),
//   	groupName: jsii.String("groupName"),
//   	variation: jsii.String("variation"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnLaunch_LaunchGroupObjectProperty struct {
	// The feature that this launch is using.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// A name for this launch group.
	//
	// It can include up to 127 characters.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The feature variation to use for this launch group.
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// A description of the launch group.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

// This structure defines a metric that you want to use to evaluate the variations during a launch or experiment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDefinitionObjectProperty := &metricDefinitionObjectProperty{
//   	entityIdKey: jsii.String("entityIdKey"),
//   	eventPattern: jsii.String("eventPattern"),
//   	metricName: jsii.String("metricName"),
//   	valueKey: jsii.String("valueKey"),
//
//   	// the properties below are optional
//   	unitLabel: jsii.String("unitLabel"),
//   }
//
type CfnLaunch_MetricDefinitionObjectProperty struct {
	// The entity, such as a user or session, that does an action that causes a metric value to be recorded.
	//
	// An example is `userDetails.userID` .
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// The EventBridge event pattern that defines how the metric is recorded.
	//
	// For more information about EventBridge event patterns, see [Amazon EventBridge event patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html) .
	EventPattern *string `field:"required" json:"eventPattern" yaml:"eventPattern"`
	// A name for the metric.
	//
	// It can include up to 255 characters.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The value that is tracked to produce the metric.
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// A label for the units that the metric is measuring.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

// A structure that defines when each step of the launch is to start, and how much launch traffic is to be allocated to each variation during each step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepConfigProperty := &stepConfigProperty{
//   	groupWeights: []interface{}{
//   		&groupToWeightProperty{
//   			groupName: jsii.String("groupName"),
//   			splitWeight: jsii.Number(123),
//   		},
//   	},
//   	startTime: jsii.String("startTime"),
//   }
//
type CfnLaunch_StepConfigProperty struct {
	// An array of structures that define how much launch traffic to allocate to each launch group during this step of the launch.
	GroupWeights interface{} `field:"required" json:"groupWeights" yaml:"groupWeights"`
	// The date and time to start this step of the launch.
	//
	// Use UTC format, `yyyy-MM-ddTHH:mm:ssZ` . For example, `2025-11-25T23:59:59Z`
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

// Properties for defining a `CfnLaunch`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchProps := &cfnLaunchProps{
//   	groups: []interface{}{
//   		&launchGroupObjectProperty{
//   			feature: jsii.String("feature"),
//   			groupName: jsii.String("groupName"),
//   			variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	project: jsii.String("project"),
//   	scheduledSplitsConfig: []interface{}{
//   		&stepConfigProperty{
//   			groupWeights: []interface{}{
//   				&groupToWeightProperty{
//   					groupName: jsii.String("groupName"),
//   					splitWeight: jsii.Number(123),
//   				},
//   			},
//   			startTime: jsii.String("startTime"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	executionStatus: &executionStatusObjectProperty{
//   		status: jsii.String("status"),
//
//   		// the properties below are optional
//   		desiredState: jsii.String("desiredState"),
//   		reason: jsii.String("reason"),
//   	},
//   	metricMonitors: []interface{}{
//   		&metricDefinitionObjectProperty{
//   			entityIdKey: jsii.String("entityIdKey"),
//   			eventPattern: jsii.String("eventPattern"),
//   			metricName: jsii.String("metricName"),
//   			valueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			unitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	randomizationSalt: jsii.String("randomizationSalt"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLaunchProps struct {
	// An array of structures that contains the feature and variations that are to be used for the launch.
	//
	// You can up to five launch groups in a launch.
	Groups interface{} `field:"required" json:"groups" yaml:"groups"`
	// The name for the launch.
	//
	// It can include up to 127 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name or ARN of the project that you want to create the launch in.
	Project *string `field:"required" json:"project" yaml:"project"`
	// An array of structures that define the traffic allocation percentages among the feature variations during each step of the launch.
	ScheduledSplitsConfig interface{} `field:"required" json:"scheduledSplitsConfig" yaml:"scheduledSplitsConfig"`
	// An optional description for the launch.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A structure that you can use to start and stop the launch.
	ExecutionStatus interface{} `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// An array of structures that define the metrics that will be used to monitor the launch performance.
	//
	// You can have up to three metric monitors in the array.
	MetricMonitors interface{} `field:"optional" json:"metricMonitors" yaml:"metricMonitors"`
	// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served.
	//
	// This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the launch name as the `randomizationsSalt` .
	RandomizationSalt *string `field:"optional" json:"randomizationSalt" yaml:"randomizationSalt"`
	// Assigns one or more tags (key-value pairs) to the launch.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a launch.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Evidently::Project`.
//
// Creates a project, which is the logical object in Evidently that can contain features, launches, and experiments. Use projects to group similar features together.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProject := awscdk.Aws_evidently.NewCfnProject(this, jsii.String("MyCfnProject"), &cfnProjectProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dataDelivery: &dataDeliveryObjectProperty{
//   		logGroup: jsii.String("logGroup"),
//   		s3: &s3DestinationProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   	description: jsii.String("description"),
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
	// The ARN of the project.
	//
	// For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject`.
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
	// A structure that contains information about where Evidently is to store evaluation events for longer term storage, if you choose to do so.
	//
	// If you choose not to store these events, Evidently deletes them after using them to produce metrics and other experiment results that you can view.
	//
	// You can't specify both `CloudWatchLogs` and `S3Destination` in the same operation.
	DataDelivery() interface{}
	SetDataDelivery(val interface{})
	// An optional description of the project.
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
	// The name for the project.
	//
	// It can include up to 127 characters.
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
	// Assigns one or more tags (key-value pairs) to the project.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a project.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
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

// The jsii proxy struct for CfnProject
type jsiiProxy_CfnProject struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnProject) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnProject) DataDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
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

func (j *jsiiProxy_CfnProject) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
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


// Create a new `AWS::Evidently::Project`.
func NewCfnProject(scope constructs.Construct, id *string, props *CfnProjectProps) CfnProject {
	_init_.Initialize()

	j := jsiiProxy_CfnProject{}

	_jsii_.Create(
		"aws-cdk-lib.aws_evidently.CfnProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Evidently::Project`.
func NewCfnProject_Override(c CfnProject, scope constructs.Construct, id *string, props *CfnProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_evidently.CfnProject",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProject) SetDataDelivery(val interface{}) {
	_jsii_.Set(
		j,
		"dataDelivery",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
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
		"aws-cdk-lib.aws_evidently.CfnProject",
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
		"aws-cdk-lib.aws_evidently.CfnProject",
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
func CfnProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_evidently.CfnProject",
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
		"aws-cdk-lib.aws_evidently.CfnProject",
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

// A structure that contains information about where Evidently is to store evaluation events for longer term storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataDeliveryObjectProperty := &dataDeliveryObjectProperty{
//   	logGroup: jsii.String("logGroup"),
//   	s3: &s3DestinationProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnProject_DataDeliveryObjectProperty struct {
	// If the project stores evaluation events in CloudWatch Logs , this structure stores the log group name.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// If the project stores evaluation events in an Amazon S3 bucket, this structure stores the bucket name and bucket prefix.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

// If the project stores evaluation events in an Amazon S3 bucket, this structure stores the bucket name and bucket prefix.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationProperty := &s3DestinationProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnProject_S3DestinationProperty struct {
	// The name of the bucket in which Evidently stores evaluation events.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The bucket prefix in which Evidently stores evaluation events.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &cfnProjectProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dataDelivery: &dataDeliveryObjectProperty{
//   		logGroup: jsii.String("logGroup"),
//   		s3: &s3DestinationProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProjectProps struct {
	// The name for the project.
	//
	// It can include up to 127 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A structure that contains information about where Evidently is to store evaluation events for longer term storage, if you choose to do so.
	//
	// If you choose not to store these events, Evidently deletes them after using them to produce metrics and other experiment results that you can view.
	//
	// You can't specify both `CloudWatchLogs` and `S3Destination` in the same operation.
	DataDelivery interface{} `field:"optional" json:"dataDelivery" yaml:"dataDelivery"`
	// An optional description of the project.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Assigns one or more tags (key-value pairs) to the project.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a project.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

