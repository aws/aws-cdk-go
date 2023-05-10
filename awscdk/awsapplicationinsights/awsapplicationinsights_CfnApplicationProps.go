package awsapplicationinsights

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	ResourceGroupName: jsii.String("resourceGroupName"),
//
//   	// the properties below are optional
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
	// Application Insights can create applications based on a resource group or on an account.
	//
	// To create an account-based application using all of the resources in the account, set this parameter to `ACCOUNT_BASED` .
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

