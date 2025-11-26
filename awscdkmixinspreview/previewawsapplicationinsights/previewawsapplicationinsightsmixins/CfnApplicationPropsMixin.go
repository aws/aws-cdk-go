package previewawsapplicationinsightsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsapplicationinsights/previewawsapplicationinsightsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApplicationInsights::Application` resource adds an application that is created from a resource group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationPropsMixin(&CfnApplicationMixinProps{
//   	AttachMissingPermission: jsii.Boolean(false),
//   	AutoConfigurationEnabled: jsii.Boolean(false),
//   	ComponentMonitoringSettings: []interface{}{
//   		&ComponentMonitoringSettingProperty{
//   			ComponentArn: jsii.String("componentArn"),
//   			ComponentConfigurationMode: jsii.String("componentConfigurationMode"),
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
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					JmxPrometheusExporter: &JMXPrometheusExporterProperty{
//   						HostPort: jsii.String("hostPort"),
//   						Jmxurl: jsii.String("jmxurl"),
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					Logs: []interface{}{
//   						&LogProperty{
//   							Encoding: jsii.String("encoding"),
//   							LogGroupName: jsii.String("logGroupName"),
//   							LogPath: jsii.String("logPath"),
//   							LogType: jsii.String("logType"),
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					NetWeaverPrometheusExporter: &NetWeaverPrometheusExporterProperty{
//   						InstanceNumbers: []*string{
//   							jsii.String("instanceNumbers"),
//   						},
//   						PrometheusPort: jsii.String("prometheusPort"),
//   						Sapsid: jsii.String("sapsid"),
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
//   									Encoding: jsii.String("encoding"),
//   									LogGroupName: jsii.String("logGroupName"),
//   									LogPath: jsii.String("logPath"),
//   									LogType: jsii.String("logType"),
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
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					JmxPrometheusExporter: &JMXPrometheusExporterProperty{
//   						HostPort: jsii.String("hostPort"),
//   						Jmxurl: jsii.String("jmxurl"),
//   						PrometheusPort: jsii.String("prometheusPort"),
//   					},
//   					Logs: []interface{}{
//   						&LogProperty{
//   							Encoding: jsii.String("encoding"),
//   							LogGroupName: jsii.String("logGroupName"),
//   							LogPath: jsii.String("logPath"),
//   							LogType: jsii.String("logType"),
//   							PatternSet: jsii.String("patternSet"),
//   						},
//   					},
//   					NetWeaverPrometheusExporter: &NetWeaverPrometheusExporterProperty{
//   						InstanceNumbers: []*string{
//   							jsii.String("instanceNumbers"),
//   						},
//   						PrometheusPort: jsii.String("prometheusPort"),
//   						Sapsid: jsii.String("sapsid"),
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
//   									Encoding: jsii.String("encoding"),
//   									LogGroupName: jsii.String("logGroupName"),
//   									LogPath: jsii.String("logPath"),
//   									LogType: jsii.String("logType"),
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
//   									PatternSet: jsii.String("patternSet"),
//   								},
//   							},
//   						},
//   						SubComponentType: jsii.String("subComponentType"),
//   					},
//   				},
//   			},
//   			Tier: jsii.String("tier"),
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
//   	ResourceGroupName: jsii.String("resourceGroupName"),
//   	SnsNotificationArn: jsii.String("snsNotificationArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html
//
type CfnApplicationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationPropsMixin
type jsiiProxy_CfnApplicationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Props() *CfnApplicationMixinProps {
	var returns *CfnApplicationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApplicationInsights::Application`.
func NewCfnApplicationPropsMixin(props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApplicationInsights::Application`.
func NewCfnApplicationPropsMixin_Override(c CfnApplicationPropsMixin, props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

