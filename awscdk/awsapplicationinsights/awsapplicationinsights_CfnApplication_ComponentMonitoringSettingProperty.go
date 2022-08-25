package awsapplicationinsights


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

