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
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The monitoring settings of the components.
	ComponentMonitoringSettings() interface{}
	SetComponentMonitoringSettings(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Describes a custom component by grouping similar standalone instances to monitor.
	CustomComponents() interface{}
	SetCustomComponents(val interface{})
	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as `instance terminated` , `failed deployment` , and others.
	CweMonitorEnabled() interface{}
	SetCweMonitorEnabled(val interface{})
	// Application Insights can create applications based on a resource group or on an account.
	//
	// To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED` .
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
	// Experimental.
	LogicalId() *string
	// The log pattern sets.
	LogPatternSets() interface{}
	SetLogPatternSets(val interface{})
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Experimental.
	Ref() *string
	// The name of the resource group used for the application.
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An array of `Tags` .
	Tags() awscdk.TagManager
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

	if err := validateNewCfnApplicationParameters(scope, id, props); err != nil {
		panic(err)
	}
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

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnApplication_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnApplication_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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

	if err := validateCfnApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnApplication) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnApplication) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApplication) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnApplication) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnApplication) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

