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
// The `AWS::ApplicationInsights::Application` resource adds an application that is created from a resource group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplication := awscdk.Aws_applicationinsights.NewCfnApplication(this, jsii.String("MyCfnApplication"), &cfnApplicationProps{
//   	resourceGroupName: jsii.String("resourceGroupName"),
//
//   	// the properties below are optional
//   	autoConfigurationEnabled: jsii.Boolean(false),
//   	componentMonitoringSettings: []interface{}{
//   		&componentMonitoringSettingProperty{
//   			componentConfigurationMode: jsii.String("componentConfigurationMode"),
//   			tier: jsii.String("tier"),
//
//   			// the properties below are optional
//   			componentArn: jsii.String("componentArn"),
//   			componentName: jsii.String("componentName"),
//   			customComponentConfiguration: &componentConfigurationProperty{
//   				configurationDetails: &configurationDetailsProperty{
//   					alarmMetrics: []interface{}{
//   						&alarmMetricProperty{
//   							alarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					alarms: []interface{}{
//   						&alarmProperty{
//   							alarmName: jsii.String("alarmName"),
//
//   							// the properties below are optional
//   							severity: jsii.String("severity"),
//   						},
//   					},
//   					haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   						agreeToInstallHanadbClient: jsii.Boolean(false),
//   						hanaPort: jsii.String("hanaPort"),
//   						hanaSecretName: jsii.String("hanaSecretName"),
//   						hanasid: jsii.String("hanasid"),
//
//   						// the properties below are optional
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   						hostPort: jsii.String("hostPort"),
//   						jmxurl: jsii.String("jmxurl"),
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					logs: []interface{}{
//   						&logProperty{
//   							logType: jsii.String("logType"),
//
//   							// the properties below are optional
//   							encoding: jsii.String("encoding"),
//   							logGroupName: jsii.String("logGroupName"),
//   							logPath: jsii.String("logPath"),
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					windowsEvents: []interface{}{
//   						&windowsEventProperty{
//   							eventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							eventName: jsii.String("eventName"),
//   							logGroupName: jsii.String("logGroupName"),
//
//   							// the properties below are optional
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				subComponentTypeConfigurations: []interface{}{
//   					&subComponentTypeConfigurationProperty{
//   						subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   							alarmMetrics: []interface{}{
//   								&alarmMetricProperty{
//   									alarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							logs: []interface{}{
//   								&logProperty{
//   									logType: jsii.String("logType"),
//
//   									// the properties below are optional
//   									encoding: jsii.String("encoding"),
//   									logGroupName: jsii.String("logGroupName"),
//   									logPath: jsii.String("logPath"),
//   									patternSet: jsii.String("patternSet"),
//   								},
//   							},
//   							windowsEvents: []interface{}{
//   								&windowsEventProperty{
//   									eventLevels: []*string{
//   										jsii.String("eventLevels"),
//   									},
//   									eventName: jsii.String("eventName"),
//   									logGroupName: jsii.String("logGroupName"),
//
//   									// the properties below are optional
//   									patternSet: jsii.String("patternSet"),
//   								},
//   							},
//   						},
//   						subComponentType: jsii.String("subComponentType"),
//   					},
//   				},
//   			},
//   			defaultOverwriteComponentConfiguration: &componentConfigurationProperty{
//   				configurationDetails: &configurationDetailsProperty{
//   					alarmMetrics: []interface{}{
//   						&alarmMetricProperty{
//   							alarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					alarms: []interface{}{
//   						&alarmProperty{
//   							alarmName: jsii.String("alarmName"),
//
//   							// the properties below are optional
//   							severity: jsii.String("severity"),
//   						},
//   					},
//   					haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   						agreeToInstallHanadbClient: jsii.Boolean(false),
//   						hanaPort: jsii.String("hanaPort"),
//   						hanaSecretName: jsii.String("hanaSecretName"),
//   						hanasid: jsii.String("hanasid"),
//
//   						// the properties below are optional
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   						hostPort: jsii.String("hostPort"),
//   						jmxurl: jsii.String("jmxurl"),
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					logs: []interface{}{
//   						&logProperty{
//   							logType: jsii.String("logType"),
//
//   							// the properties below are optional
//   							encoding: jsii.String("encoding"),
//   							logGroupName: jsii.String("logGroupName"),
//   							logPath: jsii.String("logPath"),
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					windowsEvents: []interface{}{
//   						&windowsEventProperty{
//   							eventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							eventName: jsii.String("eventName"),
//   							logGroupName: jsii.String("logGroupName"),
//
//   							// the properties below are optional
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				subComponentTypeConfigurations: []interface{}{
//   					&subComponentTypeConfigurationProperty{
//   						subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   							alarmMetrics: []interface{}{
//   								&alarmMetricProperty{
//   									alarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							logs: []interface{}{
//   								&logProperty{
//   									logType: jsii.String("logType"),
//
//   									// the properties below are optional
//   									encoding: jsii.String("encoding"),
//   									logGroupName: jsii.String("logGroupName"),
//   									logPath: jsii.String("logPath"),
//   									patternSet: jsii.String("patternSet"),
//   								},
//   							},
//   							windowsEvents: []interface{}{
//   								&windowsEventProperty{
//   									eventLevels: []*string{
//   										jsii.String("eventLevels"),
//   									},
//   									eventName: jsii.String("eventName"),
//   									logGroupName: jsii.String("logGroupName"),
//
//   									// the properties below are optional
//   									patternSet: jsii.String("patternSet"),
//   								},
//   							},
//   						},
//   						subComponentType: jsii.String("subComponentType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	customComponents: []interface{}{
//   		&customComponentProperty{
//   			componentName: jsii.String("componentName"),
//   			resourceList: []*string{
//   				jsii.String("resourceList"),
//   			},
//   		},
//   	},
//   	cweMonitorEnabled: jsii.Boolean(false),
//   	groupingType: jsii.String("groupingType"),
//   	logPatternSets: []interface{}{
//   		&logPatternSetProperty{
//   			logPatterns: []interface{}{
//   				&logPatternProperty{
//   					pattern: jsii.String("pattern"),
//   					patternName: jsii.String("patternName"),
//   					rank: jsii.Number(123),
//   				},
//   			},
//   			patternSetName: jsii.String("patternSetName"),
//   		},
//   	},
//   	opsCenterEnabled: jsii.Boolean(false),
//   	opsItemSnsTopicArn: jsii.String("opsItemSnsTopicArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Returns the Amazon Resource Name (ARN) of the application, such as `arn:aws:applicationinsights:us-east-1:123456789012:application/resource-group/my_resource_group` .
	AttrApplicationArn() *string
	// If set to `true` , the application components will be configured with the monitoring configuration recommended by Application Insights.
	AutoConfigurationEnabled() interface{}
	SetAutoConfigurationEnabled(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The monitoring settings of the components.
	ComponentMonitoringSettings() interface{}
	SetComponentMonitoringSettings(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Describes a custom component by grouping similar standalone instances to monitor.
	CustomComponents() interface{}
	SetCustomComponents(val interface{})
	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as `instance terminated` , `failed deployment` , and others.
	CweMonitorEnabled() interface{}
	SetCweMonitorEnabled(val interface{})
	// `AWS::ApplicationInsights::Application.GroupingType`.
	GroupingType() *string
	SetGroupingType(val *string)
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
	// The log pattern sets.
	LogPatternSets() interface{}
	SetLogPatternSets(val interface{})
	// The tree node.
	Node() constructs.Node
	// Indicates whether Application Insights will create OpsItems for any problem that is detected by Application Insights for an application.
	OpsCenterEnabled() interface{}
	SetOpsCenterEnabled(val interface{})
	// The SNS topic provided to Application Insights that is associated with the created OpsItems to receive SNS notifications for opsItem updates.
	OpsItemSnsTopicArn() *string
	SetOpsItemSnsTopicArn(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The name of the resource group used for the application.
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of `Tags` .
	Tags() awscdk.TagManager
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

func (j *jsiiProxy_CfnApplication) GroupingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingType",
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

func (j *jsiiProxy_CfnApplication) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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

func (j *jsiiProxy_CfnApplication) SetGroupingType(val *string) {
	_jsii_.Set(
		j,
		"groupingType",
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

func (c *jsiiProxy_CfnApplication) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplication) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnApplication) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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

func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The `AWS::ApplicationInsights::Application AlarmMetric` property type defines a metric to monitor for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmMetricProperty := &alarmMetricProperty{
//   	alarmMetricName: jsii.String("alarmMetricName"),
//   }
//
type CfnApplication_AlarmMetricProperty struct {
	// The name of the metric to be monitored for the component.
	//
	// For metrics supported by Application Insights, see [Logs and metrics supported by Amazon CloudWatch Application Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/appinsights-logs-and-metrics.html) .
	AlarmMetricName *string `field:"required" json:"alarmMetricName" yaml:"alarmMetricName"`
}

// The `AWS::ApplicationInsights::Application Alarm` property type defines a CloudWatch alarm to be monitored for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmProperty := &alarmProperty{
//   	alarmName: jsii.String("alarmName"),
//
//   	// the properties below are optional
//   	severity: jsii.String("severity"),
//   }
//
type CfnApplication_AlarmProperty struct {
	// The name of the CloudWatch alarm to be monitored for the component.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// Indicates the degree of outage when the alarm goes off.
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
}

// The `AWS::ApplicationInsights::Application ComponentConfiguration` property type defines the configuration settings of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentConfigurationProperty := &componentConfigurationProperty{
//   	configurationDetails: &configurationDetailsProperty{
//   		alarmMetrics: []interface{}{
//   			&alarmMetricProperty{
//   				alarmMetricName: jsii.String("alarmMetricName"),
//   			},
//   		},
//   		alarms: []interface{}{
//   			&alarmProperty{
//   				alarmName: jsii.String("alarmName"),
//
//   				// the properties below are optional
//   				severity: jsii.String("severity"),
//   			},
//   		},
//   		haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   			prometheusPort: jsii.String("prometheusPort"),
//   		},
//   		hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   			agreeToInstallHanadbClient: jsii.Boolean(false),
//   			hanaPort: jsii.String("hanaPort"),
//   			hanaSecretName: jsii.String("hanaSecretName"),
//   			hanasid: jsii.String("hanasid"),
//
//   			// the properties below are optional
//   			prometheusPort: jsii.String("prometheusPort"),
//   		},
//   		jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   			hostPort: jsii.String("hostPort"),
//   			jmxurl: jsii.String("jmxurl"),
//   			prometheusPort: jsii.String("prometheusPort"),
//   		},
//   		logs: []interface{}{
//   			&logProperty{
//   				logType: jsii.String("logType"),
//
//   				// the properties below are optional
//   				encoding: jsii.String("encoding"),
//   				logGroupName: jsii.String("logGroupName"),
//   				logPath: jsii.String("logPath"),
//   				patternSet: jsii.String("patternSet"),
//   			},
//   		},
//   		windowsEvents: []interface{}{
//   			&windowsEventProperty{
//   				eventLevels: []*string{
//   					jsii.String("eventLevels"),
//   				},
//   				eventName: jsii.String("eventName"),
//   				logGroupName: jsii.String("logGroupName"),
//
//   				// the properties below are optional
//   				patternSet: jsii.String("patternSet"),
//   			},
//   		},
//   	},
//   	subComponentTypeConfigurations: []interface{}{
//   		&subComponentTypeConfigurationProperty{
//   			subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   				alarmMetrics: []interface{}{
//   					&alarmMetricProperty{
//   						alarmMetricName: jsii.String("alarmMetricName"),
//   					},
//   				},
//   				logs: []interface{}{
//   					&logProperty{
//   						logType: jsii.String("logType"),
//
//   						// the properties below are optional
//   						encoding: jsii.String("encoding"),
//   						logGroupName: jsii.String("logGroupName"),
//   						logPath: jsii.String("logPath"),
//   						patternSet: jsii.String("patternSet"),
//   					},
//   				},
//   				windowsEvents: []interface{}{
//   					&windowsEventProperty{
//   						eventLevels: []*string{
//   							jsii.String("eventLevels"),
//   						},
//   						eventName: jsii.String("eventName"),
//   						logGroupName: jsii.String("logGroupName"),
//
//   						// the properties below are optional
//   						patternSet: jsii.String("patternSet"),
//   					},
//   				},
//   			},
//   			subComponentType: jsii.String("subComponentType"),
//   		},
//   	},
//   }
//
type CfnApplication_ComponentConfigurationProperty struct {
	// The configuration settings.
	ConfigurationDetails interface{} `field:"optional" json:"configurationDetails" yaml:"configurationDetails"`
	// Sub-component configurations of the component.
	SubComponentTypeConfigurations interface{} `field:"optional" json:"subComponentTypeConfigurations" yaml:"subComponentTypeConfigurations"`
}

// The `AWS::ApplicationInsights::Application ComponentMonitoringSetting` property type defines the monitoring setting of the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentMonitoringSettingProperty := &componentMonitoringSettingProperty{
//   	componentConfigurationMode: jsii.String("componentConfigurationMode"),
//   	tier: jsii.String("tier"),
//
//   	// the properties below are optional
//   	componentArn: jsii.String("componentArn"),
//   	componentName: jsii.String("componentName"),
//   	customComponentConfiguration: &componentConfigurationProperty{
//   		configurationDetails: &configurationDetailsProperty{
//   			alarmMetrics: []interface{}{
//   				&alarmMetricProperty{
//   					alarmMetricName: jsii.String("alarmMetricName"),
//   				},
//   			},
//   			alarms: []interface{}{
//   				&alarmProperty{
//   					alarmName: jsii.String("alarmName"),
//
//   					// the properties below are optional
//   					severity: jsii.String("severity"),
//   				},
//   			},
//   			haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   				prometheusPort: jsii.String("prometheusPort"),
//   			},
//   			hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   				agreeToInstallHanadbClient: jsii.Boolean(false),
//   				hanaPort: jsii.String("hanaPort"),
//   				hanaSecretName: jsii.String("hanaSecretName"),
//   				hanasid: jsii.String("hanasid"),
//
//   				// the properties below are optional
//   				prometheusPort: jsii.String("prometheusPort"),
//   			},
//   			jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   				hostPort: jsii.String("hostPort"),
//   				jmxurl: jsii.String("jmxurl"),
//   				prometheusPort: jsii.String("prometheusPort"),
//   			},
//   			logs: []interface{}{
//   				&logProperty{
//   					logType: jsii.String("logType"),
//
//   					// the properties below are optional
//   					encoding: jsii.String("encoding"),
//   					logGroupName: jsii.String("logGroupName"),
//   					logPath: jsii.String("logPath"),
//   					patternSet: jsii.String("patternSet"),
//   				},
//   			},
//   			windowsEvents: []interface{}{
//   				&windowsEventProperty{
//   					eventLevels: []*string{
//   						jsii.String("eventLevels"),
//   					},
//   					eventName: jsii.String("eventName"),
//   					logGroupName: jsii.String("logGroupName"),
//
//   					// the properties below are optional
//   					patternSet: jsii.String("patternSet"),
//   				},
//   			},
//   		},
//   		subComponentTypeConfigurations: []interface{}{
//   			&subComponentTypeConfigurationProperty{
//   				subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   					alarmMetrics: []interface{}{
//   						&alarmMetricProperty{
//   							alarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					logs: []interface{}{
//   						&logProperty{
//   							logType: jsii.String("logType"),
//
//   							// the properties below are optional
//   							encoding: jsii.String("encoding"),
//   							logGroupName: jsii.String("logGroupName"),
//   							logPath: jsii.String("logPath"),
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					windowsEvents: []interface{}{
//   						&windowsEventProperty{
//   							eventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							eventName: jsii.String("eventName"),
//   							logGroupName: jsii.String("logGroupName"),
//
//   							// the properties below are optional
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				subComponentType: jsii.String("subComponentType"),
//   			},
//   		},
//   	},
//   	defaultOverwriteComponentConfiguration: &componentConfigurationProperty{
//   		configurationDetails: &configurationDetailsProperty{
//   			alarmMetrics: []interface{}{
//   				&alarmMetricProperty{
//   					alarmMetricName: jsii.String("alarmMetricName"),
//   				},
//   			},
//   			alarms: []interface{}{
//   				&alarmProperty{
//   					alarmName: jsii.String("alarmName"),
//
//   					// the properties below are optional
//   					severity: jsii.String("severity"),
//   				},
//   			},
//   			haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   				prometheusPort: jsii.String("prometheusPort"),
//   			},
//   			hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   				agreeToInstallHanadbClient: jsii.Boolean(false),
//   				hanaPort: jsii.String("hanaPort"),
//   				hanaSecretName: jsii.String("hanaSecretName"),
//   				hanasid: jsii.String("hanasid"),
//
//   				// the properties below are optional
//   				prometheusPort: jsii.String("prometheusPort"),
//   			},
//   			jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   				hostPort: jsii.String("hostPort"),
//   				jmxurl: jsii.String("jmxurl"),
//   				prometheusPort: jsii.String("prometheusPort"),
//   			},
//   			logs: []interface{}{
//   				&logProperty{
//   					logType: jsii.String("logType"),
//
//   					// the properties below are optional
//   					encoding: jsii.String("encoding"),
//   					logGroupName: jsii.String("logGroupName"),
//   					logPath: jsii.String("logPath"),
//   					patternSet: jsii.String("patternSet"),
//   				},
//   			},
//   			windowsEvents: []interface{}{
//   				&windowsEventProperty{
//   					eventLevels: []*string{
//   						jsii.String("eventLevels"),
//   					},
//   					eventName: jsii.String("eventName"),
//   					logGroupName: jsii.String("logGroupName"),
//
//   					// the properties below are optional
//   					patternSet: jsii.String("patternSet"),
//   				},
//   			},
//   		},
//   		subComponentTypeConfigurations: []interface{}{
//   			&subComponentTypeConfigurationProperty{
//   				subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   					alarmMetrics: []interface{}{
//   						&alarmMetricProperty{
//   							alarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					logs: []interface{}{
//   						&logProperty{
//   							logType: jsii.String("logType"),
//
//   							// the properties below are optional
//   							encoding: jsii.String("encoding"),
//   							logGroupName: jsii.String("logGroupName"),
//   							logPath: jsii.String("logPath"),
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					windowsEvents: []interface{}{
//   						&windowsEventProperty{
//   							eventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							eventName: jsii.String("eventName"),
//   							logGroupName: jsii.String("logGroupName"),
//
//   							// the properties below are optional
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				subComponentType: jsii.String("subComponentType"),
//   			},
//   		},
//   	},
//   }
//
type CfnApplication_ComponentMonitoringSettingProperty struct {
	// Component monitoring can be configured in one of the following three modes:.
	//
	// - `DEFAULT` : The component will be configured with the recommended default monitoring settings of the selected `Tier` .
	// - `CUSTOM` : The component will be configured with the customized monitoring settings that are specified in `CustomComponentConfiguration` . If used, `CustomComponentConfiguration` must be provided.
	// - `DEFAULT_WITH_OVERWRITE` : The component will be configured with the recommended default monitoring settings of the selected `Tier` , and merged with customized overwrite settings that are specified in `DefaultOverwriteComponentConfiguration` . If used, `DefaultOverwriteComponentConfiguration` must be provided.
	ComponentConfigurationMode *string `field:"required" json:"componentConfigurationMode" yaml:"componentConfigurationMode"`
	// The tier of the application component.
	//
	// Supported tiers include `DOT_NET_CORE` , `DOT_NET_WORKER` , `DOT_NET_WEB` , `SQL_SERVER` , `SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP` , `SQL_SERVER_FAILOVER_CLUSTER_INSTANCE` , `MYSQL` , `POSTGRESQL` , `JAVA_JMX` , `ORACLE` , `SAP_HANA_MULTI_NODE` , `SAP_HANA_SINGLE_NODE` , `SAP_HANA_HIGH_AVAILABILITY` , `SHAREPOINT` . `ACTIVE_DIRECTORY` , and `DEFAULT` .
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// The ARN of the component.
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// The name of the component.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// Customized monitoring settings.
	//
	// Required if CUSTOM mode is configured in `ComponentConfigurationMode` .
	CustomComponentConfiguration interface{} `field:"optional" json:"customComponentConfiguration" yaml:"customComponentConfiguration"`
	// Customized overwrite monitoring settings.
	//
	// Required if CUSTOM mode is configured in `ComponentConfigurationMode` .
	DefaultOverwriteComponentConfiguration interface{} `field:"optional" json:"defaultOverwriteComponentConfiguration" yaml:"defaultOverwriteComponentConfiguration"`
}

// The `AWS::ApplicationInsights::Application ConfigurationDetails` property type specifies the configuration settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationDetailsProperty := &configurationDetailsProperty{
//   	alarmMetrics: []interface{}{
//   		&alarmMetricProperty{
//   			alarmMetricName: jsii.String("alarmMetricName"),
//   		},
//   	},
//   	alarms: []interface{}{
//   		&alarmProperty{
//   			alarmName: jsii.String("alarmName"),
//
//   			// the properties below are optional
//   			severity: jsii.String("severity"),
//   		},
//   	},
//   	haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   		prometheusPort: jsii.String("prometheusPort"),
//   	},
//   	hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   		agreeToInstallHanadbClient: jsii.Boolean(false),
//   		hanaPort: jsii.String("hanaPort"),
//   		hanaSecretName: jsii.String("hanaSecretName"),
//   		hanasid: jsii.String("hanasid"),
//
//   		// the properties below are optional
//   		prometheusPort: jsii.String("prometheusPort"),
//   	},
//   	jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   		hostPort: jsii.String("hostPort"),
//   		jmxurl: jsii.String("jmxurl"),
//   		prometheusPort: jsii.String("prometheusPort"),
//   	},
//   	logs: []interface{}{
//   		&logProperty{
//   			logType: jsii.String("logType"),
//
//   			// the properties below are optional
//   			encoding: jsii.String("encoding"),
//   			logGroupName: jsii.String("logGroupName"),
//   			logPath: jsii.String("logPath"),
//   			patternSet: jsii.String("patternSet"),
//   		},
//   	},
//   	windowsEvents: []interface{}{
//   		&windowsEventProperty{
//   			eventLevels: []*string{
//   				jsii.String("eventLevels"),
//   			},
//   			eventName: jsii.String("eventName"),
//   			logGroupName: jsii.String("logGroupName"),
//
//   			// the properties below are optional
//   			patternSet: jsii.String("patternSet"),
//   		},
//   	},
//   }
//
type CfnApplication_ConfigurationDetailsProperty struct {
	// A list of metrics to monitor for the component.
	//
	// All component types can use `AlarmMetrics` .
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of alarms to monitor for the component.
	//
	// All component types can use `Alarm` .
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// The HA cluster Prometheus Exporter settings.
	HaClusterPrometheusExporter interface{} `field:"optional" json:"haClusterPrometheusExporter" yaml:"haClusterPrometheusExporter"`
	// The HANA DB Prometheus Exporter settings.
	HanaPrometheusExporter interface{} `field:"optional" json:"hanaPrometheusExporter" yaml:"hanaPrometheusExporter"`
	// A list of Java metrics to monitor for the component.
	JmxPrometheusExporter interface{} `field:"optional" json:"jmxPrometheusExporter" yaml:"jmxPrometheusExporter"`
	// A list of logs to monitor for the component.
	//
	// Only Amazon EC2 instances can use `Logs` .
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// A list of Windows Events to monitor for the component.
	//
	// Only Amazon EC2 instances running on Windows can use `WindowsEvents` .
	WindowsEvents interface{} `field:"optional" json:"windowsEvents" yaml:"windowsEvents"`
}

// The `AWS::ApplicationInsights::Application CustomComponent` property type describes a custom component by grouping similar standalone instances to monitor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customComponentProperty := &customComponentProperty{
//   	componentName: jsii.String("componentName"),
//   	resourceList: []*string{
//   		jsii.String("resourceList"),
//   	},
//   }
//
type CfnApplication_CustomComponentProperty struct {
	// The name of the component.
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The list of resource ARNs that belong to the component.
	ResourceList *[]*string `field:"required" json:"resourceList" yaml:"resourceList"`
}

// The `AWS::ApplicationInsights::Application HAClusterPrometheusExporter` property type defines the HA cluster Prometheus Exporter settings.
//
// For more information, see the [component configuration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config-sections.html#component-configuration-prometheus) in the CloudWatch Application Insights documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hAClusterPrometheusExporterProperty := &hAClusterPrometheusExporterProperty{
//   	prometheusPort: jsii.String("prometheusPort"),
//   }
//
type CfnApplication_HAClusterPrometheusExporterProperty struct {
	// The target port to which Prometheus sends metrics.
	//
	// If not specified, the default port 9668 is used.
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

// The `AWS::ApplicationInsights::Application HANAPrometheusExporter` property type defines the HANA DB Prometheus Exporter settings.
//
// For more information, see the [component configuration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config-sections.html#component-configuration-prometheus) in the CloudWatch Application Insights documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hANAPrometheusExporterProperty := &hANAPrometheusExporterProperty{
//   	agreeToInstallHanadbClient: jsii.Boolean(false),
//   	hanaPort: jsii.String("hanaPort"),
//   	hanaSecretName: jsii.String("hanaSecretName"),
//   	hanasid: jsii.String("hanasid"),
//
//   	// the properties below are optional
//   	prometheusPort: jsii.String("prometheusPort"),
//   }
//
type CfnApplication_HANAPrometheusExporterProperty struct {
	// Designates whether you agree to install the HANA DB client.
	AgreeToInstallHanadbClient interface{} `field:"required" json:"agreeToInstallHanadbClient" yaml:"agreeToInstallHanadbClient"`
	// The HANA database port by which the exporter will query HANA metrics.
	HanaPort *string `field:"required" json:"hanaPort" yaml:"hanaPort"`
	// The AWS Secrets Manager secret that stores HANA monitoring user credentials.
	//
	// The HANA Prometheus exporter uses these credentials to connect to the database and query HANA metrics.
	HanaSecretName *string `field:"required" json:"hanaSecretName" yaml:"hanaSecretName"`
	// The three-character SAP system ID (SID) of the SAP HANA system.
	Hanasid *string `field:"required" json:"hanasid" yaml:"hanasid"`
	// The target port to which Prometheus sends metrics.
	//
	// If not specified, the default port 9668 is used.
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

// The `AWS::ApplicationInsights::Application JMXPrometheusExporter` property type defines the JMXPrometheus Exporter configuration.
//
// For more information, see the [component configuration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/component-config-sections.html#component-configuration-prometheus) in the CloudWatch Application Insights documentation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jMXPrometheusExporterProperty := &jMXPrometheusExporterProperty{
//   	hostPort: jsii.String("hostPort"),
//   	jmxurl: jsii.String("jmxurl"),
//   	prometheusPort: jsii.String("prometheusPort"),
//   }
//
type CfnApplication_JMXPrometheusExporterProperty struct {
	// The host and port to connect to through remote JMX.
	//
	// Only one of `jmxURL` and `hostPort` can be specified.
	HostPort *string `field:"optional" json:"hostPort" yaml:"hostPort"`
	// The complete JMX URL to connect to.
	Jmxurl *string `field:"optional" json:"jmxurl" yaml:"jmxurl"`
	// The target port to send Prometheus metrics to.
	//
	// If not specified, the default port `9404` is used.
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

// The `AWS::ApplicationInsights::Application LogPattern` property type specifies an object that defines the log patterns that belong to a `LogPatternSet` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logPatternProperty := &logPatternProperty{
//   	pattern: jsii.String("pattern"),
//   	patternName: jsii.String("patternName"),
//   	rank: jsii.Number(123),
//   }
//
type CfnApplication_LogPatternProperty struct {
	// A regular expression that defines the log pattern.
	//
	// A log pattern can contain up to 50 characters, and it cannot be empty.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The name of the log pattern.
	//
	// A log pattern name can contain up to 50 characters, and it cannot be empty. The characters can be Unicode letters, digits, or one of the following symbols: period, dash, underscore.
	PatternName *string `field:"required" json:"patternName" yaml:"patternName"`
	// The rank of the log pattern.
	Rank *float64 `field:"required" json:"rank" yaml:"rank"`
}

// The `AWS::ApplicationInsights::Application LogPatternSet` property type specifies the log pattern set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logPatternSetProperty := &logPatternSetProperty{
//   	logPatterns: []interface{}{
//   		&logPatternProperty{
//   			pattern: jsii.String("pattern"),
//   			patternName: jsii.String("patternName"),
//   			rank: jsii.Number(123),
//   		},
//   	},
//   	patternSetName: jsii.String("patternSetName"),
//   }
//
type CfnApplication_LogPatternSetProperty struct {
	// A list of objects that define the log patterns that belong to `LogPatternSet` .
	LogPatterns interface{} `field:"required" json:"logPatterns" yaml:"logPatterns"`
	// The name of the log pattern.
	//
	// A log pattern name can contain up to 30 characters, and it cannot be empty. The characters can be Unicode letters, digits, or one of the following symbols: period, dash, underscore.
	PatternSetName *string `field:"required" json:"patternSetName" yaml:"patternSetName"`
}

// The `AWS::ApplicationInsights::Application Log` property type specifies a log to monitor for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logProperty := &logProperty{
//   	logType: jsii.String("logType"),
//
//   	// the properties below are optional
//   	encoding: jsii.String("encoding"),
//   	logGroupName: jsii.String("logGroupName"),
//   	logPath: jsii.String("logPath"),
//   	patternSet: jsii.String("patternSet"),
//   }
//
type CfnApplication_LogProperty struct {
	// The log type decides the log patterns against which Application Insights analyzes the log.
	//
	// The log type is selected from the following: `SQL_SERVER` , `MYSQL` , `MYSQL_SLOW_QUERY` , `POSTGRESQL` , `ORACLE_ALERT` , `ORACLE_LISTENER` , `IIS` , `APPLICATION` , `WINDOWS_EVENTS` , `WINDOWS_EVENTS_ACTIVE_DIRECTORY` , `WINDOWS_EVENTS_DNS` , `WINDOWS_EVENTS_IIS` , `WINDOWS_EVENTS_SHAREPOINT` , `SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP` , `SQL_SERVER_FAILOVER_CLUSTER_INSTANCE` , `STEP_FUNCTION` , `API_GATEWAY_ACCESS` , `API_GATEWAY_EXECUTION` , `SAP_HANA_LOGS` , `SAP_HANA_TRACE` , `SAP_HANA_HIGH_AVAILABILITY` , and `DEFAULT` .
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// The type of encoding of the logs to be monitored.
	//
	// The specified encoding should be included in the list of CloudWatch agent supported encodings. If not provided, CloudWatch Application Insights uses the default encoding type for the log type:
	//
	// - `APPLICATION/DEFAULT` : utf-8 encoding
	// - `SQL_SERVER` : utf-16 encoding
	// - `IIS` : ascii encoding.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The CloudWatch log group name to be associated with the monitored log.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The path of the logs to be monitored.
	//
	// The log path must be an absolute Windows or Linux system file path. For more information, see [CloudWatch Agent Configuration File: Logs Section](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Agent-Configuration-File-Details.html#CloudWatch-Agent-Configuration-File-Logssection) .
	LogPath *string `field:"optional" json:"logPath" yaml:"logPath"`
	// The log pattern set.
	PatternSet *string `field:"optional" json:"patternSet" yaml:"patternSet"`
}

// The `AWS::ApplicationInsights::Application SubComponentConfigurationDetails` property type specifies the configuration settings of the sub-components.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subComponentConfigurationDetailsProperty := &subComponentConfigurationDetailsProperty{
//   	alarmMetrics: []interface{}{
//   		&alarmMetricProperty{
//   			alarmMetricName: jsii.String("alarmMetricName"),
//   		},
//   	},
//   	logs: []interface{}{
//   		&logProperty{
//   			logType: jsii.String("logType"),
//
//   			// the properties below are optional
//   			encoding: jsii.String("encoding"),
//   			logGroupName: jsii.String("logGroupName"),
//   			logPath: jsii.String("logPath"),
//   			patternSet: jsii.String("patternSet"),
//   		},
//   	},
//   	windowsEvents: []interface{}{
//   		&windowsEventProperty{
//   			eventLevels: []*string{
//   				jsii.String("eventLevels"),
//   			},
//   			eventName: jsii.String("eventName"),
//   			logGroupName: jsii.String("logGroupName"),
//
//   			// the properties below are optional
//   			patternSet: jsii.String("patternSet"),
//   		},
//   	},
//   }
//
type CfnApplication_SubComponentConfigurationDetailsProperty struct {
	// A list of metrics to monitor for the component.
	//
	// All component types can use `AlarmMetrics` .
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of logs to monitor for the component.
	//
	// Only Amazon EC2 instances can use `Logs` .
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// A list of Windows Events to monitor for the component.
	//
	// Only Amazon EC2 instances running on Windows can use `WindowsEvents` .
	WindowsEvents interface{} `field:"optional" json:"windowsEvents" yaml:"windowsEvents"`
}

// The `AWS::ApplicationInsights::Application SubComponentTypeConfiguration` property type specifies the sub-component configurations for a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subComponentTypeConfigurationProperty := &subComponentTypeConfigurationProperty{
//   	subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   		alarmMetrics: []interface{}{
//   			&alarmMetricProperty{
//   				alarmMetricName: jsii.String("alarmMetricName"),
//   			},
//   		},
//   		logs: []interface{}{
//   			&logProperty{
//   				logType: jsii.String("logType"),
//
//   				// the properties below are optional
//   				encoding: jsii.String("encoding"),
//   				logGroupName: jsii.String("logGroupName"),
//   				logPath: jsii.String("logPath"),
//   				patternSet: jsii.String("patternSet"),
//   			},
//   		},
//   		windowsEvents: []interface{}{
//   			&windowsEventProperty{
//   				eventLevels: []*string{
//   					jsii.String("eventLevels"),
//   				},
//   				eventName: jsii.String("eventName"),
//   				logGroupName: jsii.String("logGroupName"),
//
//   				// the properties below are optional
//   				patternSet: jsii.String("patternSet"),
//   			},
//   		},
//   	},
//   	subComponentType: jsii.String("subComponentType"),
//   }
//
type CfnApplication_SubComponentTypeConfigurationProperty struct {
	// The configuration settings of the sub-components.
	SubComponentConfigurationDetails interface{} `field:"required" json:"subComponentConfigurationDetails" yaml:"subComponentConfigurationDetails"`
	// The sub-component type.
	SubComponentType *string `field:"required" json:"subComponentType" yaml:"subComponentType"`
}

// The `AWS::ApplicationInsights::Application WindowsEvent` property type specifies a Windows Event to monitor for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   windowsEventProperty := &windowsEventProperty{
//   	eventLevels: []*string{
//   		jsii.String("eventLevels"),
//   	},
//   	eventName: jsii.String("eventName"),
//   	logGroupName: jsii.String("logGroupName"),
//
//   	// the properties below are optional
//   	patternSet: jsii.String("patternSet"),
//   }
//
type CfnApplication_WindowsEventProperty struct {
	// The levels of event to log.
	//
	// You must specify each level to log. Possible values include `INFORMATION` , `WARNING` , `ERROR` , `CRITICAL` , and `VERBOSE` . This field is required for each type of Windows Event to log.
	EventLevels *[]*string `field:"required" json:"eventLevels" yaml:"eventLevels"`
	// The type of Windows Events to log, equivalent to the Windows Event log channel name.
	//
	// For example, System, Security, CustomEventName, and so on. This field is required for each type of Windows event to log.
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	// The CloudWatch log group name to be associated with the monitored log.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The log pattern set.
	PatternSet *string `field:"optional" json:"patternSet" yaml:"patternSet"`
}

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	resourceGroupName: jsii.String("resourceGroupName"),
//
//   	// the properties below are optional
//   	autoConfigurationEnabled: jsii.Boolean(false),
//   	componentMonitoringSettings: []interface{}{
//   		&componentMonitoringSettingProperty{
//   			componentConfigurationMode: jsii.String("componentConfigurationMode"),
//   			tier: jsii.String("tier"),
//
//   			// the properties below are optional
//   			componentArn: jsii.String("componentArn"),
//   			componentName: jsii.String("componentName"),
//   			customComponentConfiguration: &componentConfigurationProperty{
//   				configurationDetails: &configurationDetailsProperty{
//   					alarmMetrics: []interface{}{
//   						&alarmMetricProperty{
//   							alarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					alarms: []interface{}{
//   						&alarmProperty{
//   							alarmName: jsii.String("alarmName"),
//
//   							// the properties below are optional
//   							severity: jsii.String("severity"),
//   						},
//   					},
//   					haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   						agreeToInstallHanadbClient: jsii.Boolean(false),
//   						hanaPort: jsii.String("hanaPort"),
//   						hanaSecretName: jsii.String("hanaSecretName"),
//   						hanasid: jsii.String("hanasid"),
//
//   						// the properties below are optional
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   						hostPort: jsii.String("hostPort"),
//   						jmxurl: jsii.String("jmxurl"),
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					logs: []interface{}{
//   						&logProperty{
//   							logType: jsii.String("logType"),
//
//   							// the properties below are optional
//   							encoding: jsii.String("encoding"),
//   							logGroupName: jsii.String("logGroupName"),
//   							logPath: jsii.String("logPath"),
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					windowsEvents: []interface{}{
//   						&windowsEventProperty{
//   							eventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							eventName: jsii.String("eventName"),
//   							logGroupName: jsii.String("logGroupName"),
//
//   							// the properties below are optional
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				subComponentTypeConfigurations: []interface{}{
//   					&subComponentTypeConfigurationProperty{
//   						subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   							alarmMetrics: []interface{}{
//   								&alarmMetricProperty{
//   									alarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							logs: []interface{}{
//   								&logProperty{
//   									logType: jsii.String("logType"),
//
//   									// the properties below are optional
//   									encoding: jsii.String("encoding"),
//   									logGroupName: jsii.String("logGroupName"),
//   									logPath: jsii.String("logPath"),
//   									patternSet: jsii.String("patternSet"),
//   								},
//   							},
//   							windowsEvents: []interface{}{
//   								&windowsEventProperty{
//   									eventLevels: []*string{
//   										jsii.String("eventLevels"),
//   									},
//   									eventName: jsii.String("eventName"),
//   									logGroupName: jsii.String("logGroupName"),
//
//   									// the properties below are optional
//   									patternSet: jsii.String("patternSet"),
//   								},
//   							},
//   						},
//   						subComponentType: jsii.String("subComponentType"),
//   					},
//   				},
//   			},
//   			defaultOverwriteComponentConfiguration: &componentConfigurationProperty{
//   				configurationDetails: &configurationDetailsProperty{
//   					alarmMetrics: []interface{}{
//   						&alarmMetricProperty{
//   							alarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					alarms: []interface{}{
//   						&alarmProperty{
//   							alarmName: jsii.String("alarmName"),
//
//   							// the properties below are optional
//   							severity: jsii.String("severity"),
//   						},
//   					},
//   					haClusterPrometheusExporter: &hAClusterPrometheusExporterProperty{
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					hanaPrometheusExporter: &hANAPrometheusExporterProperty{
//   						agreeToInstallHanadbClient: jsii.Boolean(false),
//   						hanaPort: jsii.String("hanaPort"),
//   						hanaSecretName: jsii.String("hanaSecretName"),
//   						hanasid: jsii.String("hanasid"),
//
//   						// the properties below are optional
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					jmxPrometheusExporter: &jMXPrometheusExporterProperty{
//   						hostPort: jsii.String("hostPort"),
//   						jmxurl: jsii.String("jmxurl"),
//   						prometheusPort: jsii.String("prometheusPort"),
//   					},
//   					logs: []interface{}{
//   						&logProperty{
//   							logType: jsii.String("logType"),
//
//   							// the properties below are optional
//   							encoding: jsii.String("encoding"),
//   							logGroupName: jsii.String("logGroupName"),
//   							logPath: jsii.String("logPath"),
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					windowsEvents: []interface{}{
//   						&windowsEventProperty{
//   							eventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							eventName: jsii.String("eventName"),
//   							logGroupName: jsii.String("logGroupName"),
//
//   							// the properties below are optional
//   							patternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				subComponentTypeConfigurations: []interface{}{
//   					&subComponentTypeConfigurationProperty{
//   						subComponentConfigurationDetails: &subComponentConfigurationDetailsProperty{
//   							alarmMetrics: []interface{}{
//   								&alarmMetricProperty{
//   									alarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							logs: []interface{}{
//   								&logProperty{
//   									logType: jsii.String("logType"),
//
//   									// the properties below are optional
//   									encoding: jsii.String("encoding"),
//   									logGroupName: jsii.String("logGroupName"),
//   									logPath: jsii.String("logPath"),
//   									patternSet: jsii.String("patternSet"),
//   								},
//   							},
//   							windowsEvents: []interface{}{
//   								&windowsEventProperty{
//   									eventLevels: []*string{
//   										jsii.String("eventLevels"),
//   									},
//   									eventName: jsii.String("eventName"),
//   									logGroupName: jsii.String("logGroupName"),
//
//   									// the properties below are optional
//   									patternSet: jsii.String("patternSet"),
//   								},
//   							},
//   						},
//   						subComponentType: jsii.String("subComponentType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	customComponents: []interface{}{
//   		&customComponentProperty{
//   			componentName: jsii.String("componentName"),
//   			resourceList: []*string{
//   				jsii.String("resourceList"),
//   			},
//   		},
//   	},
//   	cweMonitorEnabled: jsii.Boolean(false),
//   	groupingType: jsii.String("groupingType"),
//   	logPatternSets: []interface{}{
//   		&logPatternSetProperty{
//   			logPatterns: []interface{}{
//   				&logPatternProperty{
//   					pattern: jsii.String("pattern"),
//   					patternName: jsii.String("patternName"),
//   					rank: jsii.Number(123),
//   				},
//   			},
//   			patternSetName: jsii.String("patternSetName"),
//   		},
//   	},
//   	opsCenterEnabled: jsii.Boolean(false),
//   	opsItemSnsTopicArn: jsii.String("opsItemSnsTopicArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// The name of the resource group used for the application.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// If set to `true` , the application components will be configured with the monitoring configuration recommended by Application Insights.
	AutoConfigurationEnabled interface{} `field:"optional" json:"autoConfigurationEnabled" yaml:"autoConfigurationEnabled"`
	// The monitoring settings of the components.
	ComponentMonitoringSettings interface{} `field:"optional" json:"componentMonitoringSettings" yaml:"componentMonitoringSettings"`
	// Describes a custom component by grouping similar standalone instances to monitor.
	CustomComponents interface{} `field:"optional" json:"customComponents" yaml:"customComponents"`
	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as `instance terminated` , `failed deployment` , and others.
	CweMonitorEnabled interface{} `field:"optional" json:"cweMonitorEnabled" yaml:"cweMonitorEnabled"`
	// `AWS::ApplicationInsights::Application.GroupingType`.
	GroupingType *string `field:"optional" json:"groupingType" yaml:"groupingType"`
	// The log pattern sets.
	LogPatternSets interface{} `field:"optional" json:"logPatternSets" yaml:"logPatternSets"`
	// Indicates whether Application Insights will create OpsItems for any problem that is detected by Application Insights for an application.
	OpsCenterEnabled interface{} `field:"optional" json:"opsCenterEnabled" yaml:"opsCenterEnabled"`
	// The SNS topic provided to Application Insights that is associated with the created OpsItems to receive SNS notifications for opsItem updates.
	OpsItemSnsTopicArn *string `field:"optional" json:"opsItemSnsTopicArn" yaml:"opsItemSnsTopicArn"`
	// An array of `Tags` .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

