package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudWatch::AnomalyDetector` type specifies an anomaly detection band for a certain metric and statistic.
//
// The band represents the expected "normal" range for the metric values. Anomaly detection bands can be used for visualization of a metric's expected values, and for alarms.
//
// For more information see [Using CloudWatch anomaly detection.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Anomaly_Detection.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAnomalyDetector := awscdk.Aws_cloudwatch.NewCfnAnomalyDetector(this, jsii.String("MyCfnAnomalyDetector"), &CfnAnomalyDetectorProps{
//   	Configuration: &ConfigurationProperty{
//   		ExcludedTimeRanges: []interface{}{
//   			&RangeProperty{
//   				EndTime: jsii.String("endTime"),
//   				StartTime: jsii.String("startTime"),
//   			},
//   		},
//   		MetricTimeZone: jsii.String("metricTimeZone"),
//   	},
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MetricMathAnomalyDetector: &MetricMathAnomalyDetectorProperty{
//   		MetricDataQueries: []interface{}{
//   			&MetricDataQueryProperty{
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				AccountId: jsii.String("accountId"),
//   				Expression: jsii.String("expression"),
//   				Label: jsii.String("label"),
//   				MetricStat: &MetricStatProperty{
//   					Metric: &MetricProperty{
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//
//   						// the properties below are optional
//   						Dimensions: []interface{}{
//   							&DimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   					Period: jsii.Number(123),
//   					Stat: jsii.String("stat"),
//
//   					// the properties below are optional
//   					Unit: jsii.String("unit"),
//   				},
//   				Period: jsii.Number(123),
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Namespace: jsii.String("namespace"),
//   	SingleMetricAnomalyDetector: &SingleMetricAnomalyDetectorProperty{
//   		AccountId: jsii.String("accountId"),
//   		Dimensions: []interface{}{
//   			&DimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   		Stat: jsii.String("stat"),
//   	},
//   	Stat: jsii.String("stat"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html
//
type CfnAnomalyDetector interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Specifies details about how the anomaly detection model is to be trained, including time ranges to exclude when training and updating the model.
	Configuration() interface{}
	SetConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The dimensions of the metric associated with the anomaly detection band.
	Dimensions() interface{}
	SetDimensions(val interface{})
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
	// The CloudWatch metric math expression for this anomaly detector.
	MetricMathAnomalyDetector() interface{}
	SetMetricMathAnomalyDetector(val interface{})
	// The name of the metric associated with the anomaly detection band.
	MetricName() *string
	SetMetricName(val *string)
	// The namespace of the metric associated with the anomaly detection band.
	Namespace() *string
	SetNamespace(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The CloudWatch metric and statistic for this anomaly detector.
	SingleMetricAnomalyDetector() interface{}
	SetSingleMetricAnomalyDetector(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The statistic of the metric associated with the anomaly detection band.
	Stat() *string
	SetStat(val *string)
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

// The jsii proxy struct for CfnAnomalyDetector
type jsiiProxy_CfnAnomalyDetector struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAnomalyDetector) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Configuration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Dimensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) MetricMathAnomalyDetector() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricMathAnomalyDetector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) SingleMetricAnomalyDetector() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleMetricAnomalyDetector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnAnomalyDetector(scope constructs.Construct, id *string, props *CfnAnomalyDetectorProps) CfnAnomalyDetector {
	_init_.Initialize()

	if err := validateNewCfnAnomalyDetectorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAnomalyDetector{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnAnomalyDetector_Override(c CfnAnomalyDetector, scope constructs.Construct, id *string, props *CfnAnomalyDetectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector)SetConfiguration(val interface{}) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector)SetDimensions(val interface{}) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector)SetMetricMathAnomalyDetector(val interface{}) {
	if err := j.validateSetMetricMathAnomalyDetectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricMathAnomalyDetector",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector)SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector)SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector)SetSingleMetricAnomalyDetector(val interface{}) {
	if err := j.validateSetSingleMetricAnomalyDetectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleMetricAnomalyDetector",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector)SetStat(val *string) {
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAnomalyDetector_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnomalyDetector_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnAnomalyDetector_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnomalyDetector_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector",
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
func CfnAnomalyDetector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnomalyDetector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnomalyDetector_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudwatch.CfnAnomalyDetector",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnAnomalyDetector) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnAnomalyDetector) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnAnomalyDetector) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

