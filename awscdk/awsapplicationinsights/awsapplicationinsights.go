package awsapplicationinsights

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationinsights/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::ApplicationInsights::Application`.
//
// TODO: EXAMPLE
//
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrApplicationArn() *string
	AutoConfigurationEnabled() interface{}
	SetAutoConfigurationEnabled(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ComponentMonitoringSettings() interface{}
	SetComponentMonitoringSettings(val interface{})
	CreationStack() *[]*string
	CustomComponents() interface{}
	SetCustomComponents(val interface{})
	CweMonitorEnabled() interface{}
	SetCweMonitorEnabled(val interface{})
	LogicalId() *string
	LogPatternSets() interface{}
	SetLogPatternSets(val interface{})
	Node() constructs.Node
	OpsCenterEnabled() interface{}
	SetOpsCenterEnabled(val interface{})
	OpsItemSnsTopicArn() *string
	SetOpsItemSnsTopicArn(val *string)
	Ref() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApplication
type jsiiProxy_CfnApplication struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplication) AttrApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) AutoConfigurationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoConfigurationEnabled",
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

func (j *jsiiProxy_CfnApplication) ComponentMonitoringSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentMonitoringSettings",
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

func (j *jsiiProxy_CfnApplication) CustomComponents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CweMonitorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cweMonitorEnabled",
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

func (j *jsiiProxy_CfnApplication) LogPatternSets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPatternSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) OpsCenterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opsCenterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) OpsItemSnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsItemSnsTopicArn",
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

func (j *jsiiProxy_CfnApplication) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
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

func (j *jsiiProxy_CfnApplication) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
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


// Create a new `AWS::ApplicationInsights::Application`.
func NewCfnApplication(scope constructs.Construct, id *string, props *CfnApplicationProps) CfnApplication {
	_init_.Initialize()

	j := jsiiProxy_CfnApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApplicationInsights::Application`.
func NewCfnApplication_Override(c CfnApplication, scope constructs.Construct, id *string, props *CfnApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplication) SetAutoConfigurationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"autoConfigurationEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetComponentMonitoringSettings(val interface{}) {
	_jsii_.Set(
		j,
		"componentMonitoringSettings",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetCustomComponents(val interface{}) {
	_jsii_.Set(
		j,
		"customComponents",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetCweMonitorEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cweMonitorEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetLogPatternSets(val interface{}) {
	_jsii_.Set(
		j,
		"logPatternSets",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetOpsCenterEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"opsCenterEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetOpsItemSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"opsItemSnsTopicArn",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
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
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
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
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
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
func CfnApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
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
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
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
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

// Experimental.
func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnApplication_AlarmMetricProperty struct {
	// `CfnApplication.AlarmMetricProperty.AlarmMetricName`.
	AlarmMetricName *string `json:"alarmMetricName"`
}

// TODO: EXAMPLE
//
type CfnApplication_AlarmProperty struct {
	// `CfnApplication.AlarmProperty.AlarmName`.
	AlarmName *string `json:"alarmName"`
	// `CfnApplication.AlarmProperty.Severity`.
	Severity *string `json:"severity"`
}

// TODO: EXAMPLE
//
type CfnApplication_ComponentConfigurationProperty struct {
	// `CfnApplication.ComponentConfigurationProperty.ConfigurationDetails`.
	ConfigurationDetails interface{} `json:"configurationDetails"`
	// `CfnApplication.ComponentConfigurationProperty.SubComponentTypeConfigurations`.
	SubComponentTypeConfigurations interface{} `json:"subComponentTypeConfigurations"`
}

// TODO: EXAMPLE
//
type CfnApplication_ComponentMonitoringSettingProperty struct {
	// `CfnApplication.ComponentMonitoringSettingProperty.ComponentARN`.
	ComponentArn *string `json:"componentArn"`
	// `CfnApplication.ComponentMonitoringSettingProperty.ComponentConfigurationMode`.
	ComponentConfigurationMode *string `json:"componentConfigurationMode"`
	// `CfnApplication.ComponentMonitoringSettingProperty.ComponentName`.
	ComponentName *string `json:"componentName"`
	// `CfnApplication.ComponentMonitoringSettingProperty.CustomComponentConfiguration`.
	CustomComponentConfiguration interface{} `json:"customComponentConfiguration"`
	// `CfnApplication.ComponentMonitoringSettingProperty.DefaultOverwriteComponentConfiguration`.
	DefaultOverwriteComponentConfiguration interface{} `json:"defaultOverwriteComponentConfiguration"`
	// `CfnApplication.ComponentMonitoringSettingProperty.Tier`.
	Tier *string `json:"tier"`
}

// TODO: EXAMPLE
//
type CfnApplication_ConfigurationDetailsProperty struct {
	// `CfnApplication.ConfigurationDetailsProperty.AlarmMetrics`.
	AlarmMetrics interface{} `json:"alarmMetrics"`
	// `CfnApplication.ConfigurationDetailsProperty.Alarms`.
	Alarms interface{} `json:"alarms"`
	// `CfnApplication.ConfigurationDetailsProperty.JMXPrometheusExporter`.
	JmxPrometheusExporter interface{} `json:"jmxPrometheusExporter"`
	// `CfnApplication.ConfigurationDetailsProperty.Logs`.
	Logs interface{} `json:"logs"`
	// `CfnApplication.ConfigurationDetailsProperty.WindowsEvents`.
	WindowsEvents interface{} `json:"windowsEvents"`
}

// TODO: EXAMPLE
//
type CfnApplication_CustomComponentProperty struct {
	// `CfnApplication.CustomComponentProperty.ComponentName`.
	ComponentName *string `json:"componentName"`
	// `CfnApplication.CustomComponentProperty.ResourceList`.
	ResourceList *[]*string `json:"resourceList"`
}

// TODO: EXAMPLE
//
type CfnApplication_JMXPrometheusExporterProperty struct {
	// `CfnApplication.JMXPrometheusExporterProperty.HostPort`.
	HostPort *string `json:"hostPort"`
	// `CfnApplication.JMXPrometheusExporterProperty.JMXURL`.
	Jmxurl *string `json:"jmxurl"`
	// `CfnApplication.JMXPrometheusExporterProperty.PrometheusPort`.
	PrometheusPort *string `json:"prometheusPort"`
}

// TODO: EXAMPLE
//
type CfnApplication_LogPatternProperty struct {
	// `CfnApplication.LogPatternProperty.Pattern`.
	Pattern *string `json:"pattern"`
	// `CfnApplication.LogPatternProperty.PatternName`.
	PatternName *string `json:"patternName"`
	// `CfnApplication.LogPatternProperty.Rank`.
	Rank *float64 `json:"rank"`
}

// TODO: EXAMPLE
//
type CfnApplication_LogPatternSetProperty struct {
	// `CfnApplication.LogPatternSetProperty.LogPatterns`.
	LogPatterns interface{} `json:"logPatterns"`
	// `CfnApplication.LogPatternSetProperty.PatternSetName`.
	PatternSetName *string `json:"patternSetName"`
}

// TODO: EXAMPLE
//
type CfnApplication_LogProperty struct {
	// `CfnApplication.LogProperty.Encoding`.
	Encoding *string `json:"encoding"`
	// `CfnApplication.LogProperty.LogGroupName`.
	LogGroupName *string `json:"logGroupName"`
	// `CfnApplication.LogProperty.LogPath`.
	LogPath *string `json:"logPath"`
	// `CfnApplication.LogProperty.LogType`.
	LogType *string `json:"logType"`
	// `CfnApplication.LogProperty.PatternSet`.
	PatternSet *string `json:"patternSet"`
}

// TODO: EXAMPLE
//
type CfnApplication_SubComponentConfigurationDetailsProperty struct {
	// `CfnApplication.SubComponentConfigurationDetailsProperty.AlarmMetrics`.
	AlarmMetrics interface{} `json:"alarmMetrics"`
	// `CfnApplication.SubComponentConfigurationDetailsProperty.Logs`.
	Logs interface{} `json:"logs"`
	// `CfnApplication.SubComponentConfigurationDetailsProperty.WindowsEvents`.
	WindowsEvents interface{} `json:"windowsEvents"`
}

// TODO: EXAMPLE
//
type CfnApplication_SubComponentTypeConfigurationProperty struct {
	// `CfnApplication.SubComponentTypeConfigurationProperty.SubComponentConfigurationDetails`.
	SubComponentConfigurationDetails interface{} `json:"subComponentConfigurationDetails"`
	// `CfnApplication.SubComponentTypeConfigurationProperty.SubComponentType`.
	SubComponentType *string `json:"subComponentType"`
}

// TODO: EXAMPLE
//
type CfnApplication_WindowsEventProperty struct {
	// `CfnApplication.WindowsEventProperty.EventLevels`.
	EventLevels *[]*string `json:"eventLevels"`
	// `CfnApplication.WindowsEventProperty.EventName`.
	EventName *string `json:"eventName"`
	// `CfnApplication.WindowsEventProperty.LogGroupName`.
	LogGroupName *string `json:"logGroupName"`
	// `CfnApplication.WindowsEventProperty.PatternSet`.
	PatternSet *string `json:"patternSet"`
}

// Properties for defining a `AWS::ApplicationInsights::Application`.
//
// TODO: EXAMPLE
//
type CfnApplicationProps struct {
	// `AWS::ApplicationInsights::Application.AutoConfigurationEnabled`.
	AutoConfigurationEnabled interface{} `json:"autoConfigurationEnabled"`
	// `AWS::ApplicationInsights::Application.ComponentMonitoringSettings`.
	ComponentMonitoringSettings interface{} `json:"componentMonitoringSettings"`
	// `AWS::ApplicationInsights::Application.CustomComponents`.
	CustomComponents interface{} `json:"customComponents"`
	// `AWS::ApplicationInsights::Application.CWEMonitorEnabled`.
	CweMonitorEnabled interface{} `json:"cweMonitorEnabled"`
	// `AWS::ApplicationInsights::Application.LogPatternSets`.
	LogPatternSets interface{} `json:"logPatternSets"`
	// `AWS::ApplicationInsights::Application.OpsCenterEnabled`.
	OpsCenterEnabled interface{} `json:"opsCenterEnabled"`
	// `AWS::ApplicationInsights::Application.OpsItemSNSTopicArn`.
	OpsItemSnsTopicArn *string `json:"opsItemSnsTopicArn"`
	// `AWS::ApplicationInsights::Application.ResourceGroupName`.
	ResourceGroupName *string `json:"resourceGroupName"`
	// `AWS::ApplicationInsights::Application.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

