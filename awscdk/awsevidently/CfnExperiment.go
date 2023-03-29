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
//   cfnExperiment := awscdk.Aws_evidently.NewCfnExperiment(this, jsii.String("MyCfnExperiment"), &CfnExperimentProps{
//   	MetricGoals: []interface{}{
//   		&MetricGoalObjectProperty{
//   			DesiredChange: jsii.String("desiredChange"),
//   			EntityIdKey: jsii.String("entityIdKey"),
//   			MetricName: jsii.String("metricName"),
//   			ValueKey: jsii.String("valueKey"),
//
//   			// the properties below are optional
//   			EventPattern: jsii.String("eventPattern"),
//   			UnitLabel: jsii.String("unitLabel"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OnlineAbConfig: &OnlineAbConfigObjectProperty{
//   		ControlTreatmentName: jsii.String("controlTreatmentName"),
//   		TreatmentWeights: []interface{}{
//   			&TreatmentToWeightProperty{
//   				SplitWeight: jsii.Number(123),
//   				Treatment: jsii.String("treatment"),
//   			},
//   		},
//   	},
//   	Project: jsii.String("project"),
//   	Treatments: []interface{}{
//   		&TreatmentObjectProperty{
//   			Feature: jsii.String("feature"),
//   			TreatmentName: jsii.String("treatmentName"),
//   			Variation: jsii.String("variation"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	RandomizationSalt: jsii.String("randomizationSalt"),
//   	RemoveSegment: jsii.Boolean(false),
//   	RunningStatus: &RunningStatusObjectProperty{
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		AnalysisCompleteTime: jsii.String("analysisCompleteTime"),
//   		DesiredState: jsii.String("desiredState"),
//   		Reason: jsii.String("reason"),
//   	},
//   	SamplingRate: jsii.Number(123),
//   	Segment: jsii.String("segment"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
	// Set this to `true` to remove the segment that is associated with this experiment.
	//
	// You can't use this parameter if the experiment is currently running.
	RemoveSegment() interface{}
	SetRemoveSegment(val interface{})
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
	// Specifies an audience *segment* to use in the experiment.
	//
	// When a segment is used in an experiment, only user sessions that match the segment pattern are used in the experiment.
	//
	// For more information, see [Segment rule pattern syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments-syntax.html) .
	Segment() *string
	SetSegment(val *string)
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

func (j *jsiiProxy_CfnExperiment) RemoveSegment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeSegment",
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

func (j *jsiiProxy_CfnExperiment) Segment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segment",
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

func (j *jsiiProxy_CfnExperiment) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Evidently::Experiment`.
func NewCfnExperiment(scope constructs.Construct, id *string, props *CfnExperimentProps) CfnExperiment {
	_init_.Initialize()

	if err := validateNewCfnExperimentParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnExperiment)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetMetricGoals(val interface{}) {
	if err := j.validateSetMetricGoalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricGoals",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetOnlineAbConfig(val interface{}) {
	if err := j.validateSetOnlineAbConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlineAbConfig",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetRandomizationSalt(val *string) {
	_jsii_.Set(
		j,
		"randomizationSalt",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetRemoveSegment(val interface{}) {
	if err := j.validateSetRemoveSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeSegment",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetRunningStatus(val interface{}) {
	if err := j.validateSetRunningStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runningStatus",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetSamplingRate(val *float64) {
	_jsii_.Set(
		j,
		"samplingRate",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetSegment(val *string) {
	_jsii_.Set(
		j,
		"segment",
		val,
	)
}

func (j *jsiiProxy_CfnExperiment)SetTreatments(val interface{}) {
	if err := j.validateSetTreatmentsParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateCfnExperiment_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnExperiment_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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
func CfnExperiment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExperiment_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnExperiment) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnExperiment) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnExperiment) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnExperiment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnExperiment) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnExperiment) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnExperiment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnExperiment) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnExperiment) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnExperiment) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnExperiment) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperiment) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnExperiment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnExperiment) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnExperiment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnExperiment) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

