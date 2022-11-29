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

