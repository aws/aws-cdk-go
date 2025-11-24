package mixinsawsapplicationinsights

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html
//
type CfnApplicationMixinProps struct {
	// If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-attachmissingpermission
	//
	AttachMissingPermission interface{} `field:"optional" json:"attachMissingPermission" yaml:"attachMissingPermission"`
	// If set to `true` , the application components will be configured with the monitoring configuration recommended by Application Insights.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-autoconfigurationenabled
	//
	AutoConfigurationEnabled interface{} `field:"optional" json:"autoConfigurationEnabled" yaml:"autoConfigurationEnabled"`
	// The monitoring settings of the components.
	//
	// Not required to set up default monitoring for all components. To set up default monitoring for all components, set `AutoConfigurationEnabled` to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-componentmonitoringsettings
	//
	ComponentMonitoringSettings interface{} `field:"optional" json:"componentMonitoringSettings" yaml:"componentMonitoringSettings"`
	// Describes a custom component by grouping similar standalone instances to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-customcomponents
	//
	CustomComponents interface{} `field:"optional" json:"customComponents" yaml:"customComponents"`
	// Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as `instance terminated` , `failed deployment` , and others.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-cwemonitorenabled
	//
	CweMonitorEnabled interface{} `field:"optional" json:"cweMonitorEnabled" yaml:"cweMonitorEnabled"`
	// Application Insights can create applications based on a resource group or on an account.
	//
	// To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-groupingtype
	//
	GroupingType *string `field:"optional" json:"groupingType" yaml:"groupingType"`
	// The log pattern sets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-logpatternsets
	//
	LogPatternSets interface{} `field:"optional" json:"logPatternSets" yaml:"logPatternSets"`
	// Indicates whether Application Insights will create OpsItems for any problem that is detected by Application Insights for an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-opscenterenabled
	//
	OpsCenterEnabled interface{} `field:"optional" json:"opsCenterEnabled" yaml:"opsCenterEnabled"`
	// The SNS topic provided to Application Insights that is associated with the created OpsItems to receive SNS notifications for opsItem updates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-opsitemsnstopicarn
	//
	OpsItemSnsTopicArn *string `field:"optional" json:"opsItemSnsTopicArn" yaml:"opsItemSnsTopicArn"`
	// The name of the resource group used for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-resourcegroupname
	//
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// The SNS topic ARN that is associated with SNS notifications for updates or issues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-snsnotificationarn
	//
	SnsNotificationArn *string `field:"optional" json:"snsNotificationArn" yaml:"snsNotificationArn"`
	// An array of `Tags` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationinsights-application.html#cfn-applicationinsights-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

