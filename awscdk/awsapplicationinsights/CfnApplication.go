package awsapplicationinsights

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationinsights/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapplicationinsights"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApplicationInsights::Application` resource adds an application that is created from a resource group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplication := awscdk.Aws_applicationinsights.NewCfnApplication(this, jsii.String("MyCfnApplication"), &CfnApplicationProps{
//   	ResourceGroupName: jsii.String("resourceGroupName"),
//
//   	// the properties below are optional
//   	AttachMissingPermission: jsii.Boolean(false),
//   	AutoConfigurationEnabled: jsii.Boolean(false),
//   	ComponentMonitoringSettings: []interface{}{
//   		&ComponentMonitoringSettingProperty{
//   			ComponentConfigurationMode: jsii.String("componentConfigurationMode"),
//   			Tier: jsii.String("tier"),
//
//   			// the properties below are optional
//   			ComponentArn: jsii.String("componentArn"),
//   			ComponentName: jsii.String("componentName"),
//   			CustomComponentConfiguration: &ComponentConfigurationProperty{
//   				ConfigurationDetails: &ConfigurationDetailsProperty{
//   					AlarmMetrics: []interface{}{
//   						&AlarmMetricProperty{
//   							AlarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					Alarms: []interface{}{
//   						&AlarmProperty{
//   							AlarmName: jsii.String("alarmName"),
//
//   							// the properties below are optional
//   							Severity: jsii.String("severity"),
//   						},
//   					},
//   					HaClusterPrometheusExporter: &HAClusterPrometheusExporterProperty{
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					HanaPrometheusExporter: &HANAPrometheusExporterProperty{
//   						AgreeToInstallHanadbClient: jsii.Boolean(false),
//   						HanaPort: jsii.String("hanaPort"),
//   						HanaSecretName: jsii.String("hanaSecretName"),
//   						Hanasid: jsii.String("hanasid"),
//
//   						// the properties below are optional
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					JmxPrometheusExporter: &JMXPrometheusExporterProperty{
//   						HostPort: jsii.String("hostPort"),
//   						Jmxurl: jsii.String("jmxurl"),
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					Logs: []interface{}{
//   						&LogProperty{
//   							LogType: jsii.String("logType"),
//
//   							// the properties below are optional
//   							Encoding: jsii.String("encoding"),
//   							LogGroupName: jsii.String("logGroupName"),
//   							LogPath: jsii.String("logPath"),
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					NetWeaverPrometheusExporter: &NetWeaverPrometheusExporterProperty{
//   						InstanceNumbers: []*string{
//   							jsii.String("instanceNumbers"),
//   						},
//   						Sapsid: jsii.String("sapsid"),
//
//   						// the properties below are optional
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					Processes: []interface{}{
//   						&ProcessProperty{
//   							AlarmMetrics: []interface{}{
//   								&AlarmMetricProperty{
//   									AlarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							ProcessName: jsii.String("processName"),
//   						},
//   					},
//   					SqlServerPrometheusExporter: &SQLServerPrometheusExporterProperty{
//   						PrometheusPort: jsii.String("prometheusPort"),
//   						SqlSecretName: jsii.String("sqlSecretName"),
//   					},
//   					WindowsEvents: []interface{}{
//   						&WindowsEventProperty{
//   							EventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							EventName: jsii.String("eventName"),
//   							LogGroupName: jsii.String("logGroupName"),
//
//   							// the properties below are optional
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				SubComponentTypeConfigurations: []interface{}{
//   					&SubComponentTypeConfigurationProperty{
//   						SubComponentConfigurationDetails: &SubComponentConfigurationDetailsProperty{
//   							AlarmMetrics: []interface{}{
//   								&AlarmMetricProperty{
//   									AlarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							Logs: []interface{}{
//   								&LogProperty{
//   									LogType: jsii.String("logType"),
//
//   									// the properties below are optional
//   									Encoding: jsii.String("encoding"),
//   									LogGroupName: jsii.String("logGroupName"),
//   									LogPath: jsii.String("logPath"),
//   									PatternSet: jsii.String("patternSet"),
//   								},
//   							},
//   							Processes: []interface{}{
//   								&ProcessProperty{
//   									AlarmMetrics: []interface{}{
//   										&AlarmMetricProperty{
//   											AlarmMetricName: jsii.String("alarmMetricName"),
//   										},
//   									},
//   									ProcessName: jsii.String("processName"),
//   								},
//   							},
//   							WindowsEvents: []interface{}{
//   								&WindowsEventProperty{
//   									EventLevels: []*string{
//   										jsii.String("eventLevels"),
//   									},
//   									EventName: jsii.String("eventName"),
//   									LogGroupName: jsii.String("logGroupName"),
//
//   									// the properties below are optional
//   									PatternSet: jsii.String("patternSet"),
//   								},
//   							},
//   						},
//   						SubComponentType: jsii.String("subComponentType"),
//   					},
//   				},
//   			},
//   			DefaultOverwriteComponentConfiguration: &ComponentConfigurationProperty{
//   				ConfigurationDetails: &ConfigurationDetailsProperty{
//   					AlarmMetrics: []interface{}{
//   						&AlarmMetricProperty{
//   							AlarmMetricName: jsii.String("alarmMetricName"),
//   						},
//   					},
//   					Alarms: []interface{}{
//   						&AlarmProperty{
//   							AlarmName: jsii.String("alarmName"),
//
//   							// the properties below are optional
//   							Severity: jsii.String("severity"),
//   						},
//   					},
//   					HaClusterPrometheusExporter: &HAClusterPrometheusExporterProperty{
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					HanaPrometheusExporter: &HANAPrometheusExporterProperty{
//   						AgreeToInstallHanadbClient: jsii.Boolean(false),
//   						HanaPort: jsii.String("hanaPort"),
//   						HanaSecretName: jsii.String("hanaSecretName"),
//   						Hanasid: jsii.String("hanasid"),
//
//   						// the properties below are optional
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					JmxPrometheusExporter: &JMXPrometheusExporterProperty{
//   						HostPort: jsii.String("hostPort"),
//   						Jmxurl: jsii.String("jmxurl"),
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					Logs: []interface{}{
//   						&LogProperty{
//   							LogType: jsii.String("logType"),
//
//   							// the properties below are optional
//   							Encoding: jsii.String("encoding"),
//   							LogGroupName: jsii.String("logGroupName"),
//   							LogPath: jsii.String("logPath"),
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					NetWeaverPrometheusExporter: &NetWeaverPrometheusExporterProperty{
//   						InstanceNumbers: []*string{
//   							jsii.String("instanceNumbers"),
//   						},
//   						Sapsid: jsii.String("sapsid"),
//
//   						// the properties below are optional
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					Processes: []interface{}{
//   						&ProcessProperty{
//   							AlarmMetrics: []interface{}{
//   								&AlarmMetricProperty{
//   									AlarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							ProcessName: jsii.String("processName"),
//   						},
//   					},
//   					SqlServerPrometheusExporter: &SQLServerPrometheusExporterProperty{
//   						PrometheusPort: jsii.String("prometheusPort"),
//   						SqlSecretName: jsii.String("sqlSecretName"),
//   					},
//   					WindowsEvents: []interface{}{
//   						&WindowsEventProperty{
//   							EventLevels: []*string{
//   								jsii.String("eventLevels"),
//   							},
//   							EventName: jsii.String("eventName"),
//   							LogGroupName: jsii.String("logGroupName"),
//
//   							// the properties below are optional
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   				},
//   				SubComponentTypeConfigurations: []interface{}{
//   					&SubComponentTypeConfigurationProperty{
//   						SubComponentConfigurationDetails: &SubComponentConfigurationDetailsProperty{
//   							AlarmMetrics: []interface{}{
//   								&AlarmMetricProperty{
//   									AlarmMetricName: jsii.String("alarmMetricName"),
//   								},
//   							},
//   							Logs: []interface{}{
//   								&LogProperty{
//   									LogType: jsii.String("logType"),
//
//   									// the properties below are optional
//   									Encoding: jsii.String("encoding"),
//   									LogGroupName: jsii.String("logGroupName"),
//   									LogPath: jsii.String("logPath"),
//   									PatternSet: jsii.String("patternSet"),
//   								},
//   							},
//   							Processes: []interface{}{
//   								&ProcessProperty{
//   									AlarmMetrics: []interface{}{
//   										&AlarmMetricProperty{
//   											AlarmMetricName: jsii.String("alarmMetricName"),
//   										},
//   									},
//   									ProcessName: jsii.String("processName"),
//   								},
//   							},
//   							WindowsEvents: []interface{}{
//   								&WindowsEventProperty{
//   									EventLevels: []*string{
//   										jsii.String("eventLevels"),
//   									},
//   									EventName: jsii.String("eventName"),
//   									LogGroupName: jsii.String("logGroupName"),
//
//   									// the properties below are optional
//   									PatternSet: jsii.String("patternSet"),
//   								},
//   							},
//   						},
//   						SubComponentType: jsii.String("subComponentType"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	CustomComponents: []interface{}{
//   		&CustomComponentProperty{
//   			ComponentName: jsii.String("componentName"),
//   			ResourceList: []*string{
//   				jsii.String("resourceList"),
//   			},
//   		},
//   	},
//   	CweMonitorEnabled: jsii.Boolean(false),
//   	GroupingType: jsii.String("groupingType"),
//   	LogPatternSets: []interface{}{
//   		&LogPatternSetProperty{
//   			LogPatterns: []interface{}{
//   				&LogPatternProperty{
//   					Pattern: jsii.String("pattern"),
//   					PatternName: jsii.String("patternName"),
//   					Rank: jsii.Number(123),
//   				},
//   			},
//   			PatternSetName: jsii.String("patternSetName"),
//   		},
//   	},
//   	OpsCenterEnabled: jsii.Boolean(false),
//   	OpsItemSnsTopicArn: jsii.String("opsItemSnsTopicArn"),
//   	SnsNotificationArn: jsii.String("snsNotificationArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html
//
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsapplicationinsights.IApplicationRef
	awscdk.ITaggable
	// A reference to a Application resource.
	ApplicationRef() *interfacesawsapplicationinsights.ApplicationReference
	// If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing.
	AttachMissingPermission() interface{}
	SetAttachMissingPermission(val interface{})
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
	Env() *interfaces.ResourceEnvironment
	// Application Insights can create applications based on a resource group or on an account.
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
	// The SNS topic ARN that is associated with SNS notifications for updates or issues.
	SnsNotificationArn() *string
	SetSnsNotificationArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An array of `Tags` .
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnApplication
type jsiiProxy_CfnApplication struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsapplicationinsightsIApplicationRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnApplication) ApplicationRef() *interfacesawsapplicationinsights.ApplicationReference {
	var returns *interfacesawsapplicationinsights.ApplicationReference
	_jsii_.Get(
		j,
		"applicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) AttachMissingPermission() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachMissingPermission",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnApplication) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
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

func (j *jsiiProxy_CfnApplication) SnsNotificationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsNotificationArn",
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

func (j *jsiiProxy_CfnApplication) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
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

	if err := validateNewCfnApplicationParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnApplication)SetAttachMissingPermission(val interface{}) {
	if err := j.validateSetAttachMissingPermissionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachMissingPermission",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetAutoConfigurationEnabled(val interface{}) {
	if err := j.validateSetAutoConfigurationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoConfigurationEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetComponentMonitoringSettings(val interface{}) {
	if err := j.validateSetComponentMonitoringSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentMonitoringSettings",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetCustomComponents(val interface{}) {
	if err := j.validateSetCustomComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customComponents",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetCweMonitorEnabled(val interface{}) {
	if err := j.validateSetCweMonitorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cweMonitorEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetGroupingType(val *string) {
	_jsii_.Set(
		j,
		"groupingType",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetLogPatternSets(val interface{}) {
	if err := j.validateSetLogPatternSetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logPatternSets",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetOpsCenterEnabled(val interface{}) {
	if err := j.validateSetOpsCenterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsCenterEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetOpsItemSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"opsItemSnsTopicArn",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetSnsNotificationArn(val *string) {
	_jsii_.Set(
		j,
		"snsNotificationArn",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func CfnApplication_ArnForApplication(resource interfacesawsapplicationinsights.IApplicationRef) *string {
	_init_.Initialize()

	if err := validateCfnApplication_ArnForApplicationParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
		"arnForApplication",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnApplication.
func CfnApplication_IsCfnApplication(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsCfnApplicationParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
		"isCfnApplication",
		[]interface{}{x},
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
func CfnApplication_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnApplication_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationinsights.CfnApplication",
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
func CfnApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplication) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApplication) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnApplication) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnApplication) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplication) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApplication) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnApplication) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

