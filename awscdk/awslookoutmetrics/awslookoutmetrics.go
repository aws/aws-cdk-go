package awslookoutmetrics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslookoutmetrics/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::LookoutMetrics::Alert`.
//
// The `AWS::LookoutMetrics::Alert` type creates an alert for an anomaly detector.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   cfnAlert := lookoutmetrics.NewCfnAlert(this, jsii.String("MyCfnAlert"), &cfnAlertProps{
//   	action: &actionProperty{
//   		lambdaConfiguration: &lambdaConfigurationProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		snsConfiguration: &sNSConfigurationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			snsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	alertSensitivityThreshold: jsii.Number(123),
//   	anomalyDetectorArn: jsii.String("anomalyDetectorArn"),
//
//   	// the properties below are optional
//   	alertDescription: jsii.String("alertDescription"),
//   	alertName: jsii.String("alertName"),
//   })
//
type CfnAlert interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Action that will be triggered when there is an alert.
	Action() interface{}
	SetAction(val interface{})
	// A description of the alert.
	AlertDescription() *string
	SetAlertDescription(val *string)
	// The name of the alert.
	AlertName() *string
	SetAlertName(val *string)
	// An integer from 0 to 100 specifying the alert sensitivity threshold.
	AlertSensitivityThreshold() *float64
	SetAlertSensitivityThreshold(val *float64)
	// The ARN of the detector to which the alert is attached.
	AnomalyDetectorArn() *string
	SetAnomalyDetectorArn(val *string)
	// The Amazon Resource Name (ARN) of the alert.
	//
	// For example, `arn:aws:lookoutmetrics:us-east-2:123456789012:Alert:my-alert`.
	AttrArn() *string
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

func (j *jsiiProxy_CfnAlert) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnAlert(scope awscdk.Construct, id *string, props *CfnAlertProps) CfnAlert {
	_init_.Initialize()

	j := jsiiProxy_CfnAlert{}

	_jsii_.Create(
		"monocdk.aws_lookoutmetrics.CfnAlert",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LookoutMetrics::Alert`.
func NewCfnAlert_Override(c CfnAlert, scope awscdk.Construct, id *string, props *CfnAlertProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lookoutmetrics.CfnAlert",
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
// Experimental.
func CfnAlert_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lookoutmetrics.CfnAlert",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAlert_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lookoutmetrics.CfnAlert",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAlert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lookoutmetrics.CfnAlert",
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
		"monocdk.aws_lookoutmetrics.CfnAlert",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlert) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAlert) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAlert) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAlert) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAlert) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAlert) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAlert) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnAlert) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAlert) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAlert) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAlert) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlert) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAlert) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnAlert) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnAlert) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

// A configuration that specifies the action to perform when anomalies are detected.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   actionProperty := &actionProperty{
//   	lambdaConfiguration: &lambdaConfigurationProperty{
//   		lambdaArn: jsii.String("lambdaArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	snsConfiguration: &sNSConfigurationProperty{
//   		roleArn: jsii.String("roleArn"),
//   		snsTopicArn: jsii.String("snsTopicArn"),
//   	},
//   }
//
type CfnAlert_ActionProperty struct {
	// A configuration for an AWS Lambda channel.
	LambdaConfiguration interface{} `json:"lambdaConfiguration" yaml:"lambdaConfiguration"`
	// A configuration for an Amazon SNS channel.
	SnsConfiguration interface{} `json:"snsConfiguration" yaml:"snsConfiguration"`
}

// Contains information about a Lambda configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   lambdaConfigurationProperty := &lambdaConfigurationProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnAlert_LambdaConfigurationProperty struct {
	// The ARN of the Lambda function.
	LambdaArn *string `json:"lambdaArn" yaml:"lambdaArn"`
	// The ARN of an IAM role that has permission to invoke the Lambda function.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// Contains information about the SNS topic to which you want to send your alerts and the IAM role that has access to that topic.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   sNSConfigurationProperty := &sNSConfigurationProperty{
//   	roleArn: jsii.String("roleArn"),
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   }
//
type CfnAlert_SNSConfigurationProperty struct {
	// The ARN of the IAM role that has access to the target SNS topic.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The ARN of the target SNS topic.
	SnsTopicArn *string `json:"snsTopicArn" yaml:"snsTopicArn"`
}

// Properties for defining a `CfnAlert`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   cfnAlertProps := &cfnAlertProps{
//   	action: &actionProperty{
//   		lambdaConfiguration: &lambdaConfigurationProperty{
//   			lambdaArn: jsii.String("lambdaArn"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		snsConfiguration: &sNSConfigurationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			snsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	alertSensitivityThreshold: jsii.Number(123),
//   	anomalyDetectorArn: jsii.String("anomalyDetectorArn"),
//
//   	// the properties below are optional
//   	alertDescription: jsii.String("alertDescription"),
//   	alertName: jsii.String("alertName"),
//   }
//
type CfnAlertProps struct {
	// Action that will be triggered when there is an alert.
	Action interface{} `json:"action" yaml:"action"`
	// An integer from 0 to 100 specifying the alert sensitivity threshold.
	AlertSensitivityThreshold *float64 `json:"alertSensitivityThreshold" yaml:"alertSensitivityThreshold"`
	// The ARN of the detector to which the alert is attached.
	AnomalyDetectorArn *string `json:"anomalyDetectorArn" yaml:"anomalyDetectorArn"`
	// A description of the alert.
	AlertDescription *string `json:"alertDescription" yaml:"alertDescription"`
	// The name of the alert.
	AlertName *string `json:"alertName" yaml:"alertName"`
}

// A CloudFormation `AWS::LookoutMetrics::AnomalyDetector`.
//
// The `AWS::LookoutMetrics::AnomalyDetector` type creates an anomaly detector.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   cfnAnomalyDetector := lookoutmetrics.NewCfnAnomalyDetector(this, jsii.String("MyCfnAnomalyDetector"), &cfnAnomalyDetectorProps{
//   	anomalyDetectorConfig: &anomalyDetectorConfigProperty{
//   		anomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   	},
//   	metricSetList: []interface{}{
//   		&metricSetProperty{
//   			metricList: []interface{}{
//   				&metricProperty{
//   					aggregationFunction: jsii.String("aggregationFunction"),
//   					metricName: jsii.String("metricName"),
//
//   					// the properties below are optional
//   					namespace: jsii.String("namespace"),
//   				},
//   			},
//   			metricSetName: jsii.String("metricSetName"),
//   			metricSource: &metricSourceProperty{
//   				appFlowConfig: &appFlowConfigProperty{
//   					flowName: jsii.String("flowName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				cloudwatchConfig: &cloudwatchConfigProperty{
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				rdsSourceConfig: &rDSSourceConfigProperty{
//   					databaseHost: jsii.String("databaseHost"),
//   					databaseName: jsii.String("databaseName"),
//   					databasePort: jsii.Number(123),
//   					dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   					roleArn: jsii.String("roleArn"),
//   					secretManagerArn: jsii.String("secretManagerArn"),
//   					tableName: jsii.String("tableName"),
//   					vpcConfiguration: &vpcConfigurationProperty{
//   						securityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						subnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				redshiftSourceConfig: &redshiftSourceConfigProperty{
//   					clusterIdentifier: jsii.String("clusterIdentifier"),
//   					databaseHost: jsii.String("databaseHost"),
//   					databaseName: jsii.String("databaseName"),
//   					databasePort: jsii.Number(123),
//   					roleArn: jsii.String("roleArn"),
//   					secretManagerArn: jsii.String("secretManagerArn"),
//   					tableName: jsii.String("tableName"),
//   					vpcConfiguration: &vpcConfigurationProperty{
//   						securityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						subnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				s3SourceConfig: &s3SourceConfigProperty{
//   					fileFormatDescriptor: &fileFormatDescriptorProperty{
//   						csvFormatDescriptor: &csvFormatDescriptorProperty{
//   							charset: jsii.String("charset"),
//   							containsHeader: jsii.Boolean(false),
//   							delimiter: jsii.String("delimiter"),
//   							fileCompression: jsii.String("fileCompression"),
//   							headerList: []*string{
//   								jsii.String("headerList"),
//   							},
//   							quoteSymbol: jsii.String("quoteSymbol"),
//   						},
//   						jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   							charset: jsii.String("charset"),
//   							fileCompression: jsii.String("fileCompression"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					historicalDataPathList: []*string{
//   						jsii.String("historicalDataPathList"),
//   					},
//   					templatedPathList: []*string{
//   						jsii.String("templatedPathList"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			dimensionList: []*string{
//   				jsii.String("dimensionList"),
//   			},
//   			metricSetDescription: jsii.String("metricSetDescription"),
//   			metricSetFrequency: jsii.String("metricSetFrequency"),
//   			offset: jsii.Number(123),
//   			timestampColumn: &timestampColumnProperty{
//   				columnFormat: jsii.String("columnFormat"),
//   				columnName: jsii.String("columnName"),
//   			},
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	anomalyDetectorDescription: jsii.String("anomalyDetectorDescription"),
//   	anomalyDetectorName: jsii.String("anomalyDetectorName"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   })
//
type CfnAnomalyDetector interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Contains information about the configuration of the anomaly detector.
	AnomalyDetectorConfig() interface{}
	SetAnomalyDetectorConfig(val interface{})
	// A description of the detector.
	AnomalyDetectorDescription() *string
	SetAnomalyDetectorDescription(val *string)
	// The name of the detector.
	AnomalyDetectorName() *string
	SetAnomalyDetectorName(val *string)
	// The Amazon Resource Name (ARN) of the detector.
	//
	// For example, `arn:aws:lookoutmetrics:us-east-2:123456789012:AnomalyDetector:my-detector`.
	AttrArn() *string
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
	// The ARN of the KMS key to use to encrypt your data.
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
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
	// The detector's dataset.
	MetricSetList() interface{}
	SetMetricSetList(val interface{})
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

func (j *jsiiProxy_CfnAnomalyDetector) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnAnomalyDetector(scope awscdk.Construct, id *string, props *CfnAnomalyDetectorProps) CfnAnomalyDetector {
	_init_.Initialize()

	j := jsiiProxy_CfnAnomalyDetector{}

	_jsii_.Create(
		"monocdk.aws_lookoutmetrics.CfnAnomalyDetector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LookoutMetrics::AnomalyDetector`.
func NewCfnAnomalyDetector_Override(c CfnAnomalyDetector, scope awscdk.Construct, id *string, props *CfnAnomalyDetectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lookoutmetrics.CfnAnomalyDetector",
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
// Experimental.
func CfnAnomalyDetector_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lookoutmetrics.CfnAnomalyDetector",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAnomalyDetector_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lookoutmetrics.CfnAnomalyDetector",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAnomalyDetector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lookoutmetrics.CfnAnomalyDetector",
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
		"monocdk.aws_lookoutmetrics.CfnAnomalyDetector",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnAnomalyDetector) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAnomalyDetector) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAnomalyDetector) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnAnomalyDetector) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnAnomalyDetector) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

// Contains information about a detector's configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   anomalyDetectorConfigProperty := &anomalyDetectorConfigProperty{
//   	anomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   }
//
type CfnAnomalyDetector_AnomalyDetectorConfigProperty struct {
	// The frequency at which the detector analyzes its source data.
	AnomalyDetectorFrequency *string `json:"anomalyDetectorFrequency" yaml:"anomalyDetectorFrequency"`
}

// Details about an Amazon AppFlow flow datasource.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   appFlowConfigProperty := &appFlowConfigProperty{
//   	flowName: jsii.String("flowName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnAnomalyDetector_AppFlowConfigProperty struct {
	// name of the flow.
	FlowName *string `json:"flowName" yaml:"flowName"`
	// An IAM role that gives Amazon Lookout for Metrics permission to access the flow.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// Details about an Amazon CloudWatch datasource.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   cloudwatchConfigProperty := &cloudwatchConfigProperty{
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnAnomalyDetector_CloudwatchConfigProperty struct {
	// An IAM role that gives Amazon Lookout for Metrics permission to access data in Amazon CloudWatch.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

// Contains information about how a source CSV data file should be analyzed.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   csvFormatDescriptorProperty := &csvFormatDescriptorProperty{
//   	charset: jsii.String("charset"),
//   	containsHeader: jsii.Boolean(false),
//   	delimiter: jsii.String("delimiter"),
//   	fileCompression: jsii.String("fileCompression"),
//   	headerList: []*string{
//   		jsii.String("headerList"),
//   	},
//   	quoteSymbol: jsii.String("quoteSymbol"),
//   }
//
type CfnAnomalyDetector_CsvFormatDescriptorProperty struct {
	// The character set in which the source CSV file is written.
	Charset *string `json:"charset" yaml:"charset"`
	// Whether or not the source CSV file contains a header.
	ContainsHeader interface{} `json:"containsHeader" yaml:"containsHeader"`
	// The character used to delimit the source CSV file.
	Delimiter *string `json:"delimiter" yaml:"delimiter"`
	// The level of compression of the source CSV file.
	FileCompression *string `json:"fileCompression" yaml:"fileCompression"`
	// A list of the source CSV file's headers, if any.
	HeaderList *[]*string `json:"headerList" yaml:"headerList"`
	// The character used as a quote character.
	QuoteSymbol *string `json:"quoteSymbol" yaml:"quoteSymbol"`
}

// Contains information about a source file's formatting.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   fileFormatDescriptorProperty := &fileFormatDescriptorProperty{
//   	csvFormatDescriptor: &csvFormatDescriptorProperty{
//   		charset: jsii.String("charset"),
//   		containsHeader: jsii.Boolean(false),
//   		delimiter: jsii.String("delimiter"),
//   		fileCompression: jsii.String("fileCompression"),
//   		headerList: []*string{
//   			jsii.String("headerList"),
//   		},
//   		quoteSymbol: jsii.String("quoteSymbol"),
//   	},
//   	jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   		charset: jsii.String("charset"),
//   		fileCompression: jsii.String("fileCompression"),
//   	},
//   }
//
type CfnAnomalyDetector_FileFormatDescriptorProperty struct {
	// Contains information about how a source CSV data file should be analyzed.
	CsvFormatDescriptor interface{} `json:"csvFormatDescriptor" yaml:"csvFormatDescriptor"`
	// Contains information about how a source JSON data file should be analyzed.
	JsonFormatDescriptor interface{} `json:"jsonFormatDescriptor" yaml:"jsonFormatDescriptor"`
}

// Contains information about how a source JSON data file should be analyzed.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   jsonFormatDescriptorProperty := &jsonFormatDescriptorProperty{
//   	charset: jsii.String("charset"),
//   	fileCompression: jsii.String("fileCompression"),
//   }
//
type CfnAnomalyDetector_JsonFormatDescriptorProperty struct {
	// The character set in which the source JSON file is written.
	Charset *string `json:"charset" yaml:"charset"`
	// The level of compression of the source CSV file.
	FileCompression *string `json:"fileCompression" yaml:"fileCompression"`
}

// A calculation made by contrasting a measure and a dimension from your source data.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   metricProperty := &metricProperty{
//   	aggregationFunction: jsii.String("aggregationFunction"),
//   	metricName: jsii.String("metricName"),
//
//   	// the properties below are optional
//   	namespace: jsii.String("namespace"),
//   }
//
type CfnAnomalyDetector_MetricProperty struct {
	// The function with which the metric is calculated.
	AggregationFunction *string `json:"aggregationFunction" yaml:"aggregationFunction"`
	// The name of the metric.
	MetricName *string `json:"metricName" yaml:"metricName"`
	// The namespace for the metric.
	Namespace *string `json:"namespace" yaml:"namespace"`
}

// Contains information about a dataset.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   metricSetProperty := &metricSetProperty{
//   	metricList: []interface{}{
//   		&metricProperty{
//   			aggregationFunction: jsii.String("aggregationFunction"),
//   			metricName: jsii.String("metricName"),
//
//   			// the properties below are optional
//   			namespace: jsii.String("namespace"),
//   		},
//   	},
//   	metricSetName: jsii.String("metricSetName"),
//   	metricSource: &metricSourceProperty{
//   		appFlowConfig: &appFlowConfigProperty{
//   			flowName: jsii.String("flowName"),
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		cloudwatchConfig: &cloudwatchConfigProperty{
//   			roleArn: jsii.String("roleArn"),
//   		},
//   		rdsSourceConfig: &rDSSourceConfigProperty{
//   			databaseHost: jsii.String("databaseHost"),
//   			databaseName: jsii.String("databaseName"),
//   			databasePort: jsii.Number(123),
//   			dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   			roleArn: jsii.String("roleArn"),
//   			secretManagerArn: jsii.String("secretManagerArn"),
//   			tableName: jsii.String("tableName"),
//   			vpcConfiguration: &vpcConfigurationProperty{
//   				securityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				subnetIdList: []*string{
//   					jsii.String("subnetIdList"),
//   				},
//   			},
//   		},
//   		redshiftSourceConfig: &redshiftSourceConfigProperty{
//   			clusterIdentifier: jsii.String("clusterIdentifier"),
//   			databaseHost: jsii.String("databaseHost"),
//   			databaseName: jsii.String("databaseName"),
//   			databasePort: jsii.Number(123),
//   			roleArn: jsii.String("roleArn"),
//   			secretManagerArn: jsii.String("secretManagerArn"),
//   			tableName: jsii.String("tableName"),
//   			vpcConfiguration: &vpcConfigurationProperty{
//   				securityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				subnetIdList: []*string{
//   					jsii.String("subnetIdList"),
//   				},
//   			},
//   		},
//   		s3SourceConfig: &s3SourceConfigProperty{
//   			fileFormatDescriptor: &fileFormatDescriptorProperty{
//   				csvFormatDescriptor: &csvFormatDescriptorProperty{
//   					charset: jsii.String("charset"),
//   					containsHeader: jsii.Boolean(false),
//   					delimiter: jsii.String("delimiter"),
//   					fileCompression: jsii.String("fileCompression"),
//   					headerList: []*string{
//   						jsii.String("headerList"),
//   					},
//   					quoteSymbol: jsii.String("quoteSymbol"),
//   				},
//   				jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   					charset: jsii.String("charset"),
//   					fileCompression: jsii.String("fileCompression"),
//   				},
//   			},
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			historicalDataPathList: []*string{
//   				jsii.String("historicalDataPathList"),
//   			},
//   			templatedPathList: []*string{
//   				jsii.String("templatedPathList"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	dimensionList: []*string{
//   		jsii.String("dimensionList"),
//   	},
//   	metricSetDescription: jsii.String("metricSetDescription"),
//   	metricSetFrequency: jsii.String("metricSetFrequency"),
//   	offset: jsii.Number(123),
//   	timestampColumn: &timestampColumnProperty{
//   		columnFormat: jsii.String("columnFormat"),
//   		columnName: jsii.String("columnName"),
//   	},
//   	timezone: jsii.String("timezone"),
//   }
//
type CfnAnomalyDetector_MetricSetProperty struct {
	// A list of metrics that the dataset will contain.
	MetricList interface{} `json:"metricList" yaml:"metricList"`
	// The name of the dataset.
	MetricSetName *string `json:"metricSetName" yaml:"metricSetName"`
	// Contains information about how the source data should be interpreted.
	MetricSource interface{} `json:"metricSource" yaml:"metricSource"`
	// A list of the fields you want to treat as dimensions.
	DimensionList *[]*string `json:"dimensionList" yaml:"dimensionList"`
	// A description of the dataset you are creating.
	MetricSetDescription *string `json:"metricSetDescription" yaml:"metricSetDescription"`
	// The frequency with which the source data will be analyzed for anomalies.
	MetricSetFrequency *string `json:"metricSetFrequency" yaml:"metricSetFrequency"`
	// After an interval ends, the amount of seconds that the detector waits before importing data.
	//
	// Offset is only supported for S3 and Redshift datasources.
	Offset *float64 `json:"offset" yaml:"offset"`
	// Contains information about the column used for tracking time in your source data.
	TimestampColumn interface{} `json:"timestampColumn" yaml:"timestampColumn"`
	// The time zone in which your source data was recorded.
	Timezone *string `json:"timezone" yaml:"timezone"`
}

// Contains information about how the source data should be interpreted.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   metricSourceProperty := &metricSourceProperty{
//   	appFlowConfig: &appFlowConfigProperty{
//   		flowName: jsii.String("flowName"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	cloudwatchConfig: &cloudwatchConfigProperty{
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	rdsSourceConfig: &rDSSourceConfigProperty{
//   		databaseHost: jsii.String("databaseHost"),
//   		databaseName: jsii.String("databaseName"),
//   		databasePort: jsii.Number(123),
//   		dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   		roleArn: jsii.String("roleArn"),
//   		secretManagerArn: jsii.String("secretManagerArn"),
//   		tableName: jsii.String("tableName"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			securityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			subnetIdList: []*string{
//   				jsii.String("subnetIdList"),
//   			},
//   		},
//   	},
//   	redshiftSourceConfig: &redshiftSourceConfigProperty{
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//   		databaseHost: jsii.String("databaseHost"),
//   		databaseName: jsii.String("databaseName"),
//   		databasePort: jsii.Number(123),
//   		roleArn: jsii.String("roleArn"),
//   		secretManagerArn: jsii.String("secretManagerArn"),
//   		tableName: jsii.String("tableName"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			securityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			subnetIdList: []*string{
//   				jsii.String("subnetIdList"),
//   			},
//   		},
//   	},
//   	s3SourceConfig: &s3SourceConfigProperty{
//   		fileFormatDescriptor: &fileFormatDescriptorProperty{
//   			csvFormatDescriptor: &csvFormatDescriptorProperty{
//   				charset: jsii.String("charset"),
//   				containsHeader: jsii.Boolean(false),
//   				delimiter: jsii.String("delimiter"),
//   				fileCompression: jsii.String("fileCompression"),
//   				headerList: []*string{
//   					jsii.String("headerList"),
//   				},
//   				quoteSymbol: jsii.String("quoteSymbol"),
//   			},
//   			jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   				charset: jsii.String("charset"),
//   				fileCompression: jsii.String("fileCompression"),
//   			},
//   		},
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		historicalDataPathList: []*string{
//   			jsii.String("historicalDataPathList"),
//   		},
//   		templatedPathList: []*string{
//   			jsii.String("templatedPathList"),
//   		},
//   	},
//   }
//
type CfnAnomalyDetector_MetricSourceProperty struct {
	// An object containing information about the AppFlow configuration.
	AppFlowConfig interface{} `json:"appFlowConfig" yaml:"appFlowConfig"`
	// An object containing information about the Amazon CloudWatch monitoring configuration.
	CloudwatchConfig interface{} `json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// An object containing information about the Amazon Relational Database Service (RDS) configuration.
	RdsSourceConfig interface{} `json:"rdsSourceConfig" yaml:"rdsSourceConfig"`
	// An object containing information about the Amazon Redshift database configuration.
	RedshiftSourceConfig interface{} `json:"redshiftSourceConfig" yaml:"redshiftSourceConfig"`
	// Contains information about the configuration of the S3 bucket that contains source files.
	S3SourceConfig interface{} `json:"s3SourceConfig" yaml:"s3SourceConfig"`
}

// Contains information about the Amazon Relational Database Service (RDS) configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   rDSSourceConfigProperty := &rDSSourceConfigProperty{
//   	databaseHost: jsii.String("databaseHost"),
//   	databaseName: jsii.String("databaseName"),
//   	databasePort: jsii.Number(123),
//   	dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	roleArn: jsii.String("roleArn"),
//   	secretManagerArn: jsii.String("secretManagerArn"),
//   	tableName: jsii.String("tableName"),
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		securityGroupIdList: []*string{
//   			jsii.String("securityGroupIdList"),
//   		},
//   		subnetIdList: []*string{
//   			jsii.String("subnetIdList"),
//   		},
//   	},
//   }
//
type CfnAnomalyDetector_RDSSourceConfigProperty struct {
	// The host name of the database.
	DatabaseHost *string `json:"databaseHost" yaml:"databaseHost"`
	// The name of the RDS database.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The port number where the database can be accessed.
	DatabasePort *float64 `json:"databasePort" yaml:"databasePort"`
	// A string identifying the database instance.
	DbInstanceIdentifier *string `json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The Amazon Resource Name (ARN) of the role.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
	SecretManagerArn *string `json:"secretManagerArn" yaml:"secretManagerArn"`
	// The name of the table in the database.
	TableName *string `json:"tableName" yaml:"tableName"`
	// An object containing information about the Amazon Virtual Private Cloud (VPC) configuration.
	VpcConfiguration interface{} `json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// Provides information about the Amazon Redshift database configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   redshiftSourceConfigProperty := &redshiftSourceConfigProperty{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	databaseHost: jsii.String("databaseHost"),
//   	databaseName: jsii.String("databaseName"),
//   	databasePort: jsii.Number(123),
//   	roleArn: jsii.String("roleArn"),
//   	secretManagerArn: jsii.String("secretManagerArn"),
//   	tableName: jsii.String("tableName"),
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		securityGroupIdList: []*string{
//   			jsii.String("securityGroupIdList"),
//   		},
//   		subnetIdList: []*string{
//   			jsii.String("subnetIdList"),
//   		},
//   	},
//   }
//
type CfnAnomalyDetector_RedshiftSourceConfigProperty struct {
	// A string identifying the Redshift cluster.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the database host.
	DatabaseHost *string `json:"databaseHost" yaml:"databaseHost"`
	// The Redshift database name.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The port number where the database can be accessed.
	DatabasePort *float64 `json:"databasePort" yaml:"databasePort"`
	// The Amazon Resource Name (ARN) of the role providing access to the database.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
	SecretManagerArn *string `json:"secretManagerArn" yaml:"secretManagerArn"`
	// The table name of the Redshift database.
	TableName *string `json:"tableName" yaml:"tableName"`
	// Contains information about the Amazon Virtual Private Cloud (VPC) configuration.
	VpcConfiguration interface{} `json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// Contains information about the configuration of the S3 bucket that contains source files.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   s3SourceConfigProperty := &s3SourceConfigProperty{
//   	fileFormatDescriptor: &fileFormatDescriptorProperty{
//   		csvFormatDescriptor: &csvFormatDescriptorProperty{
//   			charset: jsii.String("charset"),
//   			containsHeader: jsii.Boolean(false),
//   			delimiter: jsii.String("delimiter"),
//   			fileCompression: jsii.String("fileCompression"),
//   			headerList: []*string{
//   				jsii.String("headerList"),
//   			},
//   			quoteSymbol: jsii.String("quoteSymbol"),
//   		},
//   		jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   			charset: jsii.String("charset"),
//   			fileCompression: jsii.String("fileCompression"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	historicalDataPathList: []*string{
//   		jsii.String("historicalDataPathList"),
//   	},
//   	templatedPathList: []*string{
//   		jsii.String("templatedPathList"),
//   	},
//   }
//
type CfnAnomalyDetector_S3SourceConfigProperty struct {
	// Contains information about a source file's formatting.
	FileFormatDescriptor interface{} `json:"fileFormatDescriptor" yaml:"fileFormatDescriptor"`
	// The ARN of an IAM role that has read and write access permissions to the source S3 bucket.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// A list of paths to the historical data files.
	HistoricalDataPathList *[]*string `json:"historicalDataPathList" yaml:"historicalDataPathList"`
	// A list of templated paths to the source files.
	TemplatedPathList *[]*string `json:"templatedPathList" yaml:"templatedPathList"`
}

// Contains information about the column used to track time in a source data file.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   timestampColumnProperty := &timestampColumnProperty{
//   	columnFormat: jsii.String("columnFormat"),
//   	columnName: jsii.String("columnName"),
//   }
//
type CfnAnomalyDetector_TimestampColumnProperty struct {
	// The format of the timestamp column.
	ColumnFormat *string `json:"columnFormat" yaml:"columnFormat"`
	// The name of the timestamp column.
	ColumnName *string `json:"columnName" yaml:"columnName"`
}

// Contains configuration information about the Amazon Virtual Private Cloud (VPC).
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   vpcConfigurationProperty := &vpcConfigurationProperty{
//   	securityGroupIdList: []*string{
//   		jsii.String("securityGroupIdList"),
//   	},
//   	subnetIdList: []*string{
//   		jsii.String("subnetIdList"),
//   	},
//   }
//
type CfnAnomalyDetector_VpcConfigurationProperty struct {
	// An array of strings containing the list of security groups.
	SecurityGroupIdList *[]*string `json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// An array of strings containing the Amazon VPC subnet IDs (e.g., `subnet-0bb1c79de3EXAMPLE` .
	SubnetIdList *[]*string `json:"subnetIdList" yaml:"subnetIdList"`
}

// Properties for defining a `CfnAnomalyDetector`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lookoutmetrics "github.com/aws/aws-cdk-go/awscdk/aws_lookoutmetrics"
//   cfnAnomalyDetectorProps := &cfnAnomalyDetectorProps{
//   	anomalyDetectorConfig: &anomalyDetectorConfigProperty{
//   		anomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   	},
//   	metricSetList: []interface{}{
//   		&metricSetProperty{
//   			metricList: []interface{}{
//   				&metricProperty{
//   					aggregationFunction: jsii.String("aggregationFunction"),
//   					metricName: jsii.String("metricName"),
//
//   					// the properties below are optional
//   					namespace: jsii.String("namespace"),
//   				},
//   			},
//   			metricSetName: jsii.String("metricSetName"),
//   			metricSource: &metricSourceProperty{
//   				appFlowConfig: &appFlowConfigProperty{
//   					flowName: jsii.String("flowName"),
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				cloudwatchConfig: &cloudwatchConfigProperty{
//   					roleArn: jsii.String("roleArn"),
//   				},
//   				rdsSourceConfig: &rDSSourceConfigProperty{
//   					databaseHost: jsii.String("databaseHost"),
//   					databaseName: jsii.String("databaseName"),
//   					databasePort: jsii.Number(123),
//   					dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   					roleArn: jsii.String("roleArn"),
//   					secretManagerArn: jsii.String("secretManagerArn"),
//   					tableName: jsii.String("tableName"),
//   					vpcConfiguration: &vpcConfigurationProperty{
//   						securityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						subnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				redshiftSourceConfig: &redshiftSourceConfigProperty{
//   					clusterIdentifier: jsii.String("clusterIdentifier"),
//   					databaseHost: jsii.String("databaseHost"),
//   					databaseName: jsii.String("databaseName"),
//   					databasePort: jsii.Number(123),
//   					roleArn: jsii.String("roleArn"),
//   					secretManagerArn: jsii.String("secretManagerArn"),
//   					tableName: jsii.String("tableName"),
//   					vpcConfiguration: &vpcConfigurationProperty{
//   						securityGroupIdList: []*string{
//   							jsii.String("securityGroupIdList"),
//   						},
//   						subnetIdList: []*string{
//   							jsii.String("subnetIdList"),
//   						},
//   					},
//   				},
//   				s3SourceConfig: &s3SourceConfigProperty{
//   					fileFormatDescriptor: &fileFormatDescriptorProperty{
//   						csvFormatDescriptor: &csvFormatDescriptorProperty{
//   							charset: jsii.String("charset"),
//   							containsHeader: jsii.Boolean(false),
//   							delimiter: jsii.String("delimiter"),
//   							fileCompression: jsii.String("fileCompression"),
//   							headerList: []*string{
//   								jsii.String("headerList"),
//   							},
//   							quoteSymbol: jsii.String("quoteSymbol"),
//   						},
//   						jsonFormatDescriptor: &jsonFormatDescriptorProperty{
//   							charset: jsii.String("charset"),
//   							fileCompression: jsii.String("fileCompression"),
//   						},
//   					},
//   					roleArn: jsii.String("roleArn"),
//
//   					// the properties below are optional
//   					historicalDataPathList: []*string{
//   						jsii.String("historicalDataPathList"),
//   					},
//   					templatedPathList: []*string{
//   						jsii.String("templatedPathList"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			dimensionList: []*string{
//   				jsii.String("dimensionList"),
//   			},
//   			metricSetDescription: jsii.String("metricSetDescription"),
//   			metricSetFrequency: jsii.String("metricSetFrequency"),
//   			offset: jsii.Number(123),
//   			timestampColumn: &timestampColumnProperty{
//   				columnFormat: jsii.String("columnFormat"),
//   				columnName: jsii.String("columnName"),
//   			},
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	anomalyDetectorDescription: jsii.String("anomalyDetectorDescription"),
//   	anomalyDetectorName: jsii.String("anomalyDetectorName"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
type CfnAnomalyDetectorProps struct {
	// Contains information about the configuration of the anomaly detector.
	AnomalyDetectorConfig interface{} `json:"anomalyDetectorConfig" yaml:"anomalyDetectorConfig"`
	// The detector's dataset.
	MetricSetList interface{} `json:"metricSetList" yaml:"metricSetList"`
	// A description of the detector.
	AnomalyDetectorDescription *string `json:"anomalyDetectorDescription" yaml:"anomalyDetectorDescription"`
	// The name of the detector.
	AnomalyDetectorName *string `json:"anomalyDetectorName" yaml:"anomalyDetectorName"`
	// The ARN of the KMS key to use to encrypt your data.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

