package awsapplicationinsights

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationinsights/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::ApplicationInsights::Application`.
//
// The `AWS::ApplicationInsights::Application` resource adds an application that is created from a resource group.
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnApplication) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnApplication(scope awscdk.Construct, id *string, props *CfnApplicationProps) CfnApplication {
	_init_.Initialize()

	j := jsiiProxy_CfnApplication{}

	_jsii_.Create(
		"monocdk.aws_applicationinsights.CfnApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApplicationInsights::Application`.
func NewCfnApplication_Override(c CfnApplication, scope awscdk.Construct, id *string, props *CfnApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationinsights.CfnApplication",
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
		"monocdk.aws_applicationinsights.CfnApplication",
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
		"monocdk.aws_applicationinsights.CfnApplication",
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
		"monocdk.aws_applicationinsights.CfnApplication",
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
		"monocdk.aws_applicationinsights.CfnApplication",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplication) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnApplication) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AWS::ApplicationInsights::Application AlarmMetric` property type defines a metric to monitor for the component.
//
// TODO: EXAMPLE
//
type CfnApplication_AlarmMetricProperty struct {
	// The name of the metric to be monitored for the component.
	//
	// For metrics supported by Application Insights, see [Logs and metrics supported by Amazon CloudWatch Application Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/appinsights-logs-and-metrics.html) .
	AlarmMetricName *string `json:"alarmMetricName" yaml:"alarmMetricName"`
}

// The `AWS::ApplicationInsights::Application Alarm` property type defines a CloudWatch alarm to be monitored for the component.
//
// TODO: EXAMPLE
//
type CfnApplication_AlarmProperty struct {
	// The name of the CloudWatch alarm to be monitored for the component.
	AlarmName *string `json:"alarmName" yaml:"alarmName"`
	// Indicates the degree of outage when the alarm goes off.
	Severity *string `json:"severity" yaml:"severity"`
}

// The `AWS::ApplicationInsights::Application ComponentConfiguration` property type defines the configuration settings of the component.
//
// TODO: EXAMPLE
//
type CfnApplication_ComponentConfigurationProperty struct {
	// The configuration settings.
	ConfigurationDetails interface{} `json:"configurationDetails" yaml:"configurationDetails"`
	// Sub-component configurations of the component.
	SubComponentTypeConfigurations interface{} `json:"subComponentTypeConfigurations" yaml:"subComponentTypeConfigurations"`
}

// The `AWS::ApplicationInsights::Application ComponentMonitoringSetting` property type defines the monitoring setting of the component.
//
// TODO: EXAMPLE
//
type CfnApplication_ComponentMonitoringSettingProperty struct {
	// Component monitoring can be configured in one of the following three modes:.
	//
	// - `DEFAULT` : The component will be configured with the recommended default monitoring settings of the selected `Tier` .
	// - `CUSTOM` : The component will be configured with the customized monitoring settings that are specified in `CustomComponentConfiguration` . If used, `CustomComponentConfiguration` must be provided.
	// - `DEFAULT_WITH_OVERWRITE` : The component will be configured with the recommended default monitoring settings of the selected `Tier` , and merged with customized overwrite settings that are specified in `DefaultOverwriteComponentConfiguration` . If used, `DefaultOverwriteComponentConfiguration` must be provided.
	ComponentConfigurationMode *string `json:"componentConfigurationMode" yaml:"componentConfigurationMode"`
	// The tier of the application component.
	//
	// Supported tiers include `DOT_NET_WORKER` , `DOT_NET_WEB` , `DOT_NET_CORE` , `SQL_SERVER` , and `DEFAULT` .
	Tier *string `json:"tier" yaml:"tier"`
	// The ARN of the component.
	ComponentArn *string `json:"componentArn" yaml:"componentArn"`
	// The name of the component.
	ComponentName *string `json:"componentName" yaml:"componentName"`
	// Customized monitoring settings.
	//
	// Required if CUSTOM mode is configured in `ComponentConfigurationMode` .
	CustomComponentConfiguration interface{} `json:"customComponentConfiguration" yaml:"customComponentConfiguration"`
	// Customized overwrite monitoring settings.
	//
	// Required if CUSTOM mode is configured in `ComponentConfigurationMode` .
	DefaultOverwriteComponentConfiguration interface{} `json:"defaultOverwriteComponentConfiguration" yaml:"defaultOverwriteComponentConfiguration"`
}

// The `AWS::ApplicationInsights::Application ConfigurationDetails` property type specifies the configuration settings.
//
// TODO: EXAMPLE
//
type CfnApplication_ConfigurationDetailsProperty struct {
	// A list of metrics to monitor for the component.
	//
	// All component types can use `AlarmMetrics` .
	AlarmMetrics interface{} `json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of alarms to monitor for the component.
	//
	// All component types can use `Alarm` .
	Alarms interface{} `json:"alarms" yaml:"alarms"`
	// A list of Java metrics to monitor for the component.
	JmxPrometheusExporter interface{} `json:"jmxPrometheusExporter" yaml:"jmxPrometheusExporter"`
	// A list of logs to monitor for the component.
	//
	// Only Amazon EC2 instances can use `Logs` .
	Logs interface{} `json:"logs" yaml:"logs"`
	// A list of Windows Events to monitor for the component.
	//
	// Only Amazon EC2 instances running on Windows can use `WindowsEvents` .
	WindowsEvents interface{} `json:"windowsEvents" yaml:"windowsEvents"`
}

// The `AWS::ApplicationInsights::Application CustomComponent` property type describes a custom component by grouping similar standalone instances to monitor.
//
// TODO: EXAMPLE
//
type CfnApplication_CustomComponentProperty struct {
	// The name of the component.
	ComponentName *string `json:"componentName" yaml:"componentName"`
	// The list of resource ARNs that belong to the component.
	ResourceList *[]*string `json:"resourceList" yaml:"resourceList"`
}

// The `AWS::ApplicationInsights::Application JMXPrometheusExporter` property type defines the JMXPrometheus Exporter configuration.
//
// For more information, see the [component configuration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config-sections.html#component-configuration-prometheus) in the CloudWatch Application Insights documentation.
//
// TODO: EXAMPLE
//
type CfnApplication_JMXPrometheusExporterProperty struct {
	// The host and port to connect to through remote JMX.
	//
	// Only one of `jmxURL` and `hostPort` can be specified.
	HostPort *string `json:"hostPort" yaml:"hostPort"`
	// The complete JMX URL to connect to.
	Jmxurl *string `json:"jmxurl" yaml:"jmxurl"`
	// The target port to send Prometheus metrics to.
	//
	// If not specified, the default port `9404` is used.
	PrometheusPort *string `json:"prometheusPort" yaml:"prometheusPort"`
}

// The `AWS::ApplicationInsights::Application LogPattern` property type specifies an object that defines the log patterns that belong to a `LogPatternSet` .
//
// TODO: EXAMPLE
//
type CfnApplication_LogPatternProperty struct {
	// A regular expression that defines the log pattern.
	//
	// A log pattern can contain up to 50 characters, and it cannot be empty.
	Pattern *string `json:"pattern" yaml:"pattern"`
	// The name of the log pattern.
	//
	// A log pattern name can contain up to 50 characters, and it cannot be empty. The characters can be Unicode letters, digits, or one of the following symbols: period, dash, underscore.
	PatternName *string `json:"patternName" yaml:"patternName"`
	// The rank of the log pattern.
	Rank *float64 `json:"rank" yaml:"rank"`
}

// The `AWS::ApplicationInsights::Application LogPatternSet` property type specifies the log pattern set.
//
// TODO: EXAMPLE
//
type CfnApplication_LogPatternSetProperty struct {
	// A list of objects that define the log patterns that belong to `LogPatternSet` .
	LogPatterns interface{} `json:"logPatterns" yaml:"logPatterns"`
	// The name of the log pattern.
	//
	// A log pattern name can contain up to 30 characters, and it cannot be empty. The characters can be Unicode letters, digits, or one of the following symbols: period, dash, underscore.
	PatternSetName *string `json:"patternSetName" yaml:"patternSetName"`
}

// The `AWS::ApplicationInsights::Application Log` property type specifies a log to monitor for the component.
//
// TODO: EXAMPLE
//
type CfnApplication_LogProperty struct {
	// The log type decides the log patterns against which Application Insights analyzes the log.
	//
	// The log type is selected from the following: `SQL_SERVER` , `IIS` , `APPLICATION` , and `DEFAULT` .
	LogType *string `json:"logType" yaml:"logType"`
	// The type of encoding of the logs to be monitored.
	//
	// The specified encoding should be included in the list of CloudWatch agent supported encodings. If not provided, CloudWatch Application Insights uses the default encoding type for the log type:
	//
	// - `APPLICATION/DEFAULT` : utf-8 encoding
	// - `SQL_SERVER` : utf-16 encoding
	// - `IIS` : ascii encoding
	Encoding *string `json:"encoding" yaml:"encoding"`
	// The CloudWatch log group name to be associated with the monitored log.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// The path of the logs to be monitored.
	//
	// The log path must be an absolute Windows or Linux system file path. For more information, see [CloudWatch Agent Configuration File: Logs Section](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-Configuration-File-Details.html#CloudWatch-Agent-Configuration-File-Logssection) .
	LogPath *string `json:"logPath" yaml:"logPath"`
	// The log pattern set.
	PatternSet *string `json:"patternSet" yaml:"patternSet"`
}

// The `AWS::ApplicationInsights::Application SubComponentConfigurationDetails` property type specifies the configuration settings of the sub-components.
//
// TODO: EXAMPLE
//
type CfnApplication_SubComponentConfigurationDetailsProperty struct {
	// A list of metrics to monitor for the component.
	//
	// All component types can use `AlarmMetrics` .
	AlarmMetrics interface{} `json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of logs to monitor for the component.
	//
	// Only Amazon EC2 instances can use `Logs` .
	Logs interface{} `json:"logs" yaml:"logs"`
	// A list of Windows Events to monitor for the component.
	//
	// Only Amazon EC2 instances running on Windows can use `WindowsEvents` .
	WindowsEvents interface{} `json:"windowsEvents" yaml:"windowsEvents"`
}

// The `AWS::ApplicationInsights::Application SubComponentTypeConfiguration` property type specifies the sub-component configurations for a component.
//
// TODO: EXAMPLE
//
type CfnApplication_SubComponentTypeConfigurationProperty struct {
	// The configuration settings of the sub-components.
	SubComponentConfigurationDetails interface{} `json:"subComponentConfigurationDetails" yaml:"subComponentConfigurationDetails"`
	// The sub-component type.
	SubComponentType *string `json:"subComponentType" yaml:"subComponentType"`
}

// The `AWS::ApplicationInsights::Application WindowsEvent` property type specifies a Windows Event to monitor for the component.
//
// TODO: EXAMPLE
//
type CfnApplication_WindowsEventProperty struct {
	// The levels of event to log.
	//
	// You must specify each level to log. Possible values include `INFORMATION` , `WARNING` , `ERROR` , `CRITICAL` , and `VERBOSE` . This field is required for each type of Windows Event to log.
	EventLevels *[]*string `json:"eventLevels" yaml:"eventLevels"`
	// The type of Windows Events to log, equivalent to the Windows Event log channel name.
	//
	// For example, System, Security, CustomEventName, and so on. This field is required for each type of Windows event to log.
	EventName *string `json:"eventName" yaml:"eventName"`
	// The CloudWatch log group name to be associated with the monitored log.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
	// The log pattern set.
	PatternSet *string `json:"patternSet" yaml:"patternSet"`
}

// Properties for defining a `CfnApplication`.
//
// TODO: EXAMPLE
//
type CfnApplicationProps struct {
	// The name of the resource group used for the application.
	ResourceGroupName *string `json:"resourceGroupName" yaml:"resourceGroupName"`
	// If set to `true` , the application components will be configured with the monitoring configuration recommended by Application Insights.
	AutoConfigurationEnabled interface{} `json:"autoConfigurationEnabled" yaml:"autoConfigurationEnabled"`
	// The monitoring settings of the components.
	ComponentMonitoringSettings interface{} `json:"componentMonitoringSettings" yaml:"componentMonitoringSettings"`
	// Describes a custom component by grouping similar standalone instances to monitor.
	CustomComponents interface{} `json:"customComponents" yaml:"customComponents"`
	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as `instance terminated` , `failed deployment` , and others.
	CweMonitorEnabled interface{} `json:"cweMonitorEnabled" yaml:"cweMonitorEnabled"`
	// The log pattern sets.
	LogPatternSets interface{} `json:"logPatternSets" yaml:"logPatternSets"`
	// Indicates whether Application Insights will create OpsItems for any problem that is detected by Application Insights for an application.
	OpsCenterEnabled interface{} `json:"opsCenterEnabled" yaml:"opsCenterEnabled"`
	// The SNS topic provided to Application Insights that is associated with the created OpsItems to receive SNS notifications for opsItem updates.
	OpsItemSnsTopicArn *string `json:"opsItemSnsTopicArn" yaml:"opsItemSnsTopicArn"`
	// An array of `Tags` .
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

