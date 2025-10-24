package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoT::SecurityProfile` resource to create a Device Defender security profile.
//
// For API reference, see [CreateSecurityProfile](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateSecurityProfile.html) and for general information, see [Detect](https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-detect.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityProfile := awscdk.Aws_iot.NewCfnSecurityProfile(this, jsii.String("MyCfnSecurityProfile"), &CfnSecurityProfileProps{
//   	AdditionalMetricsToRetainV2: []interface{}{
//   		&MetricToRetainProperty{
//   			Metric: jsii.String("metric"),
//
//   			// the properties below are optional
//   			ExportMetric: jsii.Boolean(false),
//   			MetricDimension: &MetricDimensionProperty{
//   				DimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				Operator: jsii.String("operator"),
//   			},
//   		},
//   	},
//   	AlertTargets: map[string]interface{}{
//   		"alertTargetsKey": &AlertTargetProperty{
//   			"alertTargetArn": jsii.String("alertTargetArn"),
//   			"roleArn": jsii.String("roleArn"),
//   		},
//   	},
//   	Behaviors: []interface{}{
//   		&BehaviorProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Criteria: &BehaviorCriteriaProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				ConsecutiveDatapointsToAlarm: jsii.Number(123),
//   				ConsecutiveDatapointsToClear: jsii.Number(123),
//   				DurationSeconds: jsii.Number(123),
//   				MlDetectionConfig: &MachineLearningDetectionConfigProperty{
//   					ConfidenceLevel: jsii.String("confidenceLevel"),
//   				},
//   				StatisticalThreshold: &StatisticalThresholdProperty{
//   					Statistic: jsii.String("statistic"),
//   				},
//   				Value: &MetricValueProperty{
//   					Cidrs: []*string{
//   						jsii.String("cidrs"),
//   					},
//   					Count: jsii.String("count"),
//   					Number: jsii.Number(123),
//   					Numbers: []interface{}{
//   						jsii.Number(123),
//   					},
//   					Ports: []interface{}{
//   						jsii.Number(123),
//   					},
//   					Strings: []*string{
//   						jsii.String("strings"),
//   					},
//   				},
//   			},
//   			ExportMetric: jsii.Boolean(false),
//   			Metric: jsii.String("metric"),
//   			MetricDimension: &MetricDimensionProperty{
//   				DimensionName: jsii.String("dimensionName"),
//
//   				// the properties below are optional
//   				Operator: jsii.String("operator"),
//   			},
//   			SuppressAlerts: jsii.Boolean(false),
//   		},
//   	},
//   	MetricsExportConfig: &MetricsExportConfigProperty{
//   		MqttTopic: jsii.String("mqttTopic"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SecurityProfileDescription: jsii.String("securityProfileDescription"),
//   	SecurityProfileName: jsii.String("securityProfileName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-securityprofile.html
//
type CfnSecurityProfile interface {
	awscdk.CfnResource
	ISecurityProfileRef
	awscdk.IInspectable
	awscdk.ITaggable
	// A list of metrics whose data is retained (stored).
	AdditionalMetricsToRetainV2() interface{}
	SetAdditionalMetricsToRetainV2(val interface{})
	// Specifies the destinations to which alerts are sent.
	AlertTargets() interface{}
	SetAlertTargets(val interface{})
	// The Amazon Resource Name (ARN) of the security profile.
	AttrSecurityProfileArn() *string
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors() interface{}
	SetBehaviors(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
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
	LogicalId() *string
	// Specifies the MQTT topic and role ARN required for metric export.
	MetricsExportConfig() interface{}
	SetMetricsExportConfig(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A description of the security profile.
	SecurityProfileDescription() *string
	SetSecurityProfileDescription(val *string)
	// The name you gave to the security profile.
	SecurityProfileName() *string
	SetSecurityProfileName(val *string)
	// A reference to a SecurityProfile resource.
	SecurityProfileRef() *SecurityProfileReference
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Metadata that can be used to manage the security profile.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The ARN of the target (thing group) to which the security profile is attached.
	TargetArns() *[]*string
	SetTargetArns(val *[]*string)
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

// The jsii proxy struct for CfnSecurityProfile
type jsiiProxy_CfnSecurityProfile struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_ISecurityProfileRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnSecurityProfile) AdditionalMetricsToRetainV2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalMetricsToRetainV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) AlertTargets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) AttrSecurityProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSecurityProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Behaviors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"behaviors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) MetricsExportConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsExportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) SecurityProfileDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) SecurityProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) SecurityProfileRef() *SecurityProfileReference {
	var returns *SecurityProfileReference
	_jsii_.Get(
		j,
		"securityProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) TargetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityProfile) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnSecurityProfile(scope constructs.Construct, id *string, props *CfnSecurityProfileProps) CfnSecurityProfile {
	_init_.Initialize()

	if err := validateNewCfnSecurityProfileParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityProfile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnSecurityProfile_Override(c CfnSecurityProfile, scope constructs.Construct, id *string, props *CfnSecurityProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSecurityProfile)SetAdditionalMetricsToRetainV2(val interface{}) {
	if err := j.validateSetAdditionalMetricsToRetainV2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalMetricsToRetainV2",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile)SetAlertTargets(val interface{}) {
	if err := j.validateSetAlertTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertTargets",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile)SetBehaviors(val interface{}) {
	if err := j.validateSetBehaviorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behaviors",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile)SetMetricsExportConfig(val interface{}) {
	if err := j.validateSetMetricsExportConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsExportConfig",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile)SetSecurityProfileDescription(val *string) {
	_jsii_.Set(
		j,
		"securityProfileDescription",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile)SetSecurityProfileName(val *string) {
	_jsii_.Set(
		j,
		"securityProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityProfile)SetTargetArns(val *[]*string) {
	_jsii_.Set(
		j,
		"targetArns",
		val,
	)
}

// Creates a new ISecurityProfileRef from an ARN.
func CfnSecurityProfile_FromSecurityProfileArn(scope constructs.Construct, id *string, arn *string) ISecurityProfileRef {
	_init_.Initialize()

	if err := validateCfnSecurityProfile_FromSecurityProfileArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns ISecurityProfileRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"fromSecurityProfileArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Creates a new ISecurityProfileRef from a securityProfileName.
func CfnSecurityProfile_FromSecurityProfileName(scope constructs.Construct, id *string, securityProfileName *string) ISecurityProfileRef {
	_init_.Initialize()

	if err := validateCfnSecurityProfile_FromSecurityProfileNameParameters(scope, id, securityProfileName); err != nil {
		panic(err)
	}
	var returns ISecurityProfileRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"fromSecurityProfileName",
		[]interface{}{scope, id, securityProfileName},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSecurityProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityProfile_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnSecurityProfile_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityProfile_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
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
func CfnSecurityProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iot.CfnSecurityProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnSecurityProfile) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnSecurityProfile) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnSecurityProfile) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnSecurityProfile) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSecurityProfile) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

