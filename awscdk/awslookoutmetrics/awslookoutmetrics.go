package awslookoutmetrics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslookoutmetrics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::LookoutMetrics::Alert`.
//
// TODO: EXAMPLE
//
type CfnAlert interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Action() interface{}
	SetAction(val interface{})
	AlertDescription() *string
	SetAlertDescription(val *string)
	AlertName() *string
	SetAlertName(val *string)
	AlertSensitivityThreshold() *float64
	SetAlertSensitivityThreshold(val *float64)
	AnomalyDetectorArn() *string
	SetAnomalyDetectorArn(val *string)
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
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

// The jsii proxy struct for CfnAlert
type jsiiProxy_CfnAlert struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAlert) Action() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) AlertDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) AlertName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) AlertSensitivityThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"alertSensitivityThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) AnomalyDetectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlert) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LookoutMetrics::Alert`.
func NewCfnAlert(scope constructs.Construct, id *string, props *CfnAlertProps) CfnAlert {
	_init_.Initialize()

	j := jsiiProxy_CfnAlert{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAlert",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LookoutMetrics::Alert`.
func NewCfnAlert_Override(c CfnAlert, scope constructs.Construct, id *string, props *CfnAlertProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAlert",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAlert) SetAction(val interface{}) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CfnAlert) SetAlertDescription(val *string) {
	_jsii_.Set(
		j,
		"alertDescription",
		val,
	)
}

func (j *jsiiProxy_CfnAlert) SetAlertName(val *string) {
	_jsii_.Set(
		j,
		"alertName",
		val,
	)
}

func (j *jsiiProxy_CfnAlert) SetAlertSensitivityThreshold(val *float64) {
	_jsii_.Set(
		j,
		"alertSensitivityThreshold",
		val,
	)
}

func (j *jsiiProxy_CfnAlert) SetAnomalyDetectorArn(val *string) {
	_jsii_.Set(
		j,
		"anomalyDetectorArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAlert_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAlert",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAlert_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAlert",
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
func CfnAlert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAlert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlert_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAlert",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAlert) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAlert) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAlert) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAlert) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAlert) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAlert) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnAlert) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAlert) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAlert) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAlert) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAlert) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAlert) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAlert) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnAlert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlert) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnAlert_ActionProperty struct {
	// `CfnAlert.ActionProperty.LambdaConfiguration`.
	LambdaConfiguration interface{} `json:"lambdaConfiguration"`
	// `CfnAlert.ActionProperty.SNSConfiguration`.
	SnsConfiguration interface{} `json:"snsConfiguration"`
}

// TODO: EXAMPLE
//
type CfnAlert_LambdaConfigurationProperty struct {
	// `CfnAlert.LambdaConfigurationProperty.LambdaArn`.
	LambdaArn *string `json:"lambdaArn"`
	// `CfnAlert.LambdaConfigurationProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
}

// TODO: EXAMPLE
//
type CfnAlert_SNSConfigurationProperty struct {
	// `CfnAlert.SNSConfigurationProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnAlert.SNSConfigurationProperty.SnsTopicArn`.
	SnsTopicArn *string `json:"snsTopicArn"`
}

// Properties for defining a `AWS::LookoutMetrics::Alert`.
//
// TODO: EXAMPLE
//
type CfnAlertProps struct {
	// `AWS::LookoutMetrics::Alert.Action`.
	Action interface{} `json:"action"`
	// `AWS::LookoutMetrics::Alert.AlertDescription`.
	AlertDescription *string `json:"alertDescription"`
	// `AWS::LookoutMetrics::Alert.AlertName`.
	AlertName *string `json:"alertName"`
	// `AWS::LookoutMetrics::Alert.AlertSensitivityThreshold`.
	AlertSensitivityThreshold *float64 `json:"alertSensitivityThreshold"`
	// `AWS::LookoutMetrics::Alert.AnomalyDetectorArn`.
	AnomalyDetectorArn *string `json:"anomalyDetectorArn"`
}

// A CloudFormation `AWS::LookoutMetrics::AnomalyDetector`.
//
// TODO: EXAMPLE
//
type CfnAnomalyDetector interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AnomalyDetectorConfig() interface{}
	SetAnomalyDetectorConfig(val interface{})
	AnomalyDetectorDescription() *string
	SetAnomalyDetectorDescription(val *string)
	AnomalyDetectorName() *string
	SetAnomalyDetectorName(val *string)
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	LogicalId() *string
	MetricSetList() interface{}
	SetMetricSetList(val interface{})
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

// The jsii proxy struct for CfnAnomalyDetector
type jsiiProxy_CfnAnomalyDetector struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAnomalyDetector) AnomalyDetectorConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anomalyDetectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) AnomalyDetectorDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) AnomalyDetectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
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

func (j *jsiiProxy_CfnAnomalyDetector) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetector) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
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

func (j *jsiiProxy_CfnAnomalyDetector) MetricSetList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricSetList",
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

func (j *jsiiProxy_CfnAnomalyDetector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
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


// Create a new `AWS::LookoutMetrics::AnomalyDetector`.
func NewCfnAnomalyDetector(scope constructs.Construct, id *string, props *CfnAnomalyDetectorProps) CfnAnomalyDetector {
	_init_.Initialize()

	j := jsiiProxy_CfnAnomalyDetector{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LookoutMetrics::AnomalyDetector`.
func NewCfnAnomalyDetector_Override(c CfnAnomalyDetector, scope constructs.Construct, id *string, props *CfnAnomalyDetectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetAnomalyDetectorConfig(val interface{}) {
	_jsii_.Set(
		j,
		"anomalyDetectorConfig",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetAnomalyDetectorDescription(val *string) {
	_jsii_.Set(
		j,
		"anomalyDetectorDescription",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetAnomalyDetectorName(val *string) {
	_jsii_.Set(
		j,
		"anomalyDetectorName",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnAnomalyDetector) SetMetricSetList(val interface{}) {
	_jsii_.Set(
		j,
		"metricSetList",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAnomalyDetector_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
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
func CfnAnomalyDetector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
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
		"aws-cdk-lib.aws_lookoutmetrics.CfnAnomalyDetector",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAnomalyDetector) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAnomalyDetector) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAnomalyDetector) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAnomalyDetector) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAnomalyDetector) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAnomalyDetector) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnAnomalyDetector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAnomalyDetector) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAnomalyDetector) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAnomalyDetector) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAnomalyDetector) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_AnomalyDetectorConfigProperty struct {
	// `CfnAnomalyDetector.AnomalyDetectorConfigProperty.AnomalyDetectorFrequency`.
	AnomalyDetectorFrequency *string `json:"anomalyDetectorFrequency"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_AppFlowConfigProperty struct {
	// `CfnAnomalyDetector.AppFlowConfigProperty.FlowName`.
	FlowName *string `json:"flowName"`
	// `CfnAnomalyDetector.AppFlowConfigProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_CloudwatchConfigProperty struct {
	// `CfnAnomalyDetector.CloudwatchConfigProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_CsvFormatDescriptorProperty struct {
	// `CfnAnomalyDetector.CsvFormatDescriptorProperty.Charset`.
	Charset *string `json:"charset"`
	// `CfnAnomalyDetector.CsvFormatDescriptorProperty.ContainsHeader`.
	ContainsHeader interface{} `json:"containsHeader"`
	// `CfnAnomalyDetector.CsvFormatDescriptorProperty.Delimiter`.
	Delimiter *string `json:"delimiter"`
	// `CfnAnomalyDetector.CsvFormatDescriptorProperty.FileCompression`.
	FileCompression *string `json:"fileCompression"`
	// `CfnAnomalyDetector.CsvFormatDescriptorProperty.HeaderList`.
	HeaderList *[]*string `json:"headerList"`
	// `CfnAnomalyDetector.CsvFormatDescriptorProperty.QuoteSymbol`.
	QuoteSymbol *string `json:"quoteSymbol"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_FileFormatDescriptorProperty struct {
	// `CfnAnomalyDetector.FileFormatDescriptorProperty.CsvFormatDescriptor`.
	CsvFormatDescriptor interface{} `json:"csvFormatDescriptor"`
	// `CfnAnomalyDetector.FileFormatDescriptorProperty.JsonFormatDescriptor`.
	JsonFormatDescriptor interface{} `json:"jsonFormatDescriptor"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_JsonFormatDescriptorProperty struct {
	// `CfnAnomalyDetector.JsonFormatDescriptorProperty.Charset`.
	Charset *string `json:"charset"`
	// `CfnAnomalyDetector.JsonFormatDescriptorProperty.FileCompression`.
	FileCompression *string `json:"fileCompression"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_MetricProperty struct {
	// `CfnAnomalyDetector.MetricProperty.AggregationFunction`.
	AggregationFunction *string `json:"aggregationFunction"`
	// `CfnAnomalyDetector.MetricProperty.MetricName`.
	MetricName *string `json:"metricName"`
	// `CfnAnomalyDetector.MetricProperty.Namespace`.
	Namespace *string `json:"namespace"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_MetricSetProperty struct {
	// `CfnAnomalyDetector.MetricSetProperty.DimensionList`.
	DimensionList *[]*string `json:"dimensionList"`
	// `CfnAnomalyDetector.MetricSetProperty.MetricList`.
	MetricList interface{} `json:"metricList"`
	// `CfnAnomalyDetector.MetricSetProperty.MetricSetDescription`.
	MetricSetDescription *string `json:"metricSetDescription"`
	// `CfnAnomalyDetector.MetricSetProperty.MetricSetFrequency`.
	MetricSetFrequency *string `json:"metricSetFrequency"`
	// `CfnAnomalyDetector.MetricSetProperty.MetricSetName`.
	MetricSetName *string `json:"metricSetName"`
	// `CfnAnomalyDetector.MetricSetProperty.MetricSource`.
	MetricSource interface{} `json:"metricSource"`
	// `CfnAnomalyDetector.MetricSetProperty.Offset`.
	Offset *float64 `json:"offset"`
	// `CfnAnomalyDetector.MetricSetProperty.TimestampColumn`.
	TimestampColumn interface{} `json:"timestampColumn"`
	// `CfnAnomalyDetector.MetricSetProperty.Timezone`.
	Timezone *string `json:"timezone"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_MetricSourceProperty struct {
	// `CfnAnomalyDetector.MetricSourceProperty.AppFlowConfig`.
	AppFlowConfig interface{} `json:"appFlowConfig"`
	// `CfnAnomalyDetector.MetricSourceProperty.CloudwatchConfig`.
	CloudwatchConfig interface{} `json:"cloudwatchConfig"`
	// `CfnAnomalyDetector.MetricSourceProperty.RDSSourceConfig`.
	RdsSourceConfig interface{} `json:"rdsSourceConfig"`
	// `CfnAnomalyDetector.MetricSourceProperty.RedshiftSourceConfig`.
	RedshiftSourceConfig interface{} `json:"redshiftSourceConfig"`
	// `CfnAnomalyDetector.MetricSourceProperty.S3SourceConfig`.
	S3SourceConfig interface{} `json:"s3SourceConfig"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_RDSSourceConfigProperty struct {
	// `CfnAnomalyDetector.RDSSourceConfigProperty.DatabaseHost`.
	DatabaseHost *string `json:"databaseHost"`
	// `CfnAnomalyDetector.RDSSourceConfigProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnAnomalyDetector.RDSSourceConfigProperty.DatabasePort`.
	DatabasePort *float64 `json:"databasePort"`
	// `CfnAnomalyDetector.RDSSourceConfigProperty.DBInstanceIdentifier`.
	DbInstanceIdentifier *string `json:"dbInstanceIdentifier"`
	// `CfnAnomalyDetector.RDSSourceConfigProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnAnomalyDetector.RDSSourceConfigProperty.SecretManagerArn`.
	SecretManagerArn *string `json:"secretManagerArn"`
	// `CfnAnomalyDetector.RDSSourceConfigProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnAnomalyDetector.RDSSourceConfigProperty.VpcConfiguration`.
	VpcConfiguration interface{} `json:"vpcConfiguration"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_RedshiftSourceConfigProperty struct {
	// `CfnAnomalyDetector.RedshiftSourceConfigProperty.ClusterIdentifier`.
	ClusterIdentifier *string `json:"clusterIdentifier"`
	// `CfnAnomalyDetector.RedshiftSourceConfigProperty.DatabaseHost`.
	DatabaseHost *string `json:"databaseHost"`
	// `CfnAnomalyDetector.RedshiftSourceConfigProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnAnomalyDetector.RedshiftSourceConfigProperty.DatabasePort`.
	DatabasePort *float64 `json:"databasePort"`
	// `CfnAnomalyDetector.RedshiftSourceConfigProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnAnomalyDetector.RedshiftSourceConfigProperty.SecretManagerArn`.
	SecretManagerArn *string `json:"secretManagerArn"`
	// `CfnAnomalyDetector.RedshiftSourceConfigProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnAnomalyDetector.RedshiftSourceConfigProperty.VpcConfiguration`.
	VpcConfiguration interface{} `json:"vpcConfiguration"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_S3SourceConfigProperty struct {
	// `CfnAnomalyDetector.S3SourceConfigProperty.FileFormatDescriptor`.
	FileFormatDescriptor interface{} `json:"fileFormatDescriptor"`
	// `CfnAnomalyDetector.S3SourceConfigProperty.HistoricalDataPathList`.
	HistoricalDataPathList *[]*string `json:"historicalDataPathList"`
	// `CfnAnomalyDetector.S3SourceConfigProperty.RoleArn`.
	RoleArn *string `json:"roleArn"`
	// `CfnAnomalyDetector.S3SourceConfigProperty.TemplatedPathList`.
	TemplatedPathList *[]*string `json:"templatedPathList"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_TimestampColumnProperty struct {
	// `CfnAnomalyDetector.TimestampColumnProperty.ColumnFormat`.
	ColumnFormat *string `json:"columnFormat"`
	// `CfnAnomalyDetector.TimestampColumnProperty.ColumnName`.
	ColumnName *string `json:"columnName"`
}

// TODO: EXAMPLE
//
type CfnAnomalyDetector_VpcConfigurationProperty struct {
	// `CfnAnomalyDetector.VpcConfigurationProperty.SecurityGroupIdList`.
	SecurityGroupIdList *[]*string `json:"securityGroupIdList"`
	// `CfnAnomalyDetector.VpcConfigurationProperty.SubnetIdList`.
	SubnetIdList *[]*string `json:"subnetIdList"`
}

// Properties for defining a `AWS::LookoutMetrics::AnomalyDetector`.
//
// TODO: EXAMPLE
//
type CfnAnomalyDetectorProps struct {
	// `AWS::LookoutMetrics::AnomalyDetector.AnomalyDetectorConfig`.
	AnomalyDetectorConfig interface{} `json:"anomalyDetectorConfig"`
	// `AWS::LookoutMetrics::AnomalyDetector.AnomalyDetectorDescription`.
	AnomalyDetectorDescription *string `json:"anomalyDetectorDescription"`
	// `AWS::LookoutMetrics::AnomalyDetector.AnomalyDetectorName`.
	AnomalyDetectorName *string `json:"anomalyDetectorName"`
	// `AWS::LookoutMetrics::AnomalyDetector.KmsKeyArn`.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// `AWS::LookoutMetrics::AnomalyDetector.MetricSetList`.
	MetricSetList interface{} `json:"metricSetList"`
}

